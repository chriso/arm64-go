import sys
from dataclasses import dataclass
from os import path
from typing import Optional

from lxml import etree
from lxml.etree import _ElementTree as ElementTree


@dataclass
class IndexNode:
    groupname: Optional[str]
    iclass: Optional[str]
    unallocated: bool
    header: Optional[str]
    regdiagram: Optional[list["Box"]]
    decode: Optional[list["Box"]]
    nodes: list["IndexNode"]

    @property
    def name(self) -> str:
        name = self.groupname if self.groupname is not None else self.iclass
        return name.replace("-", "")


@dataclass
class Box:
    hibit: int
    width: int
    name: Optional[str]
    value: Optional[str]

    @property
    def mask(self) -> int:
        return (1 << self.width) - 1


@dataclass
class Iclass:
    name: str
    title: str
    groups: list[str]
    regdiagram: list["Box"]
    instructiontable: "InstructionTable"


@dataclass
class InstructionTable:
    bitfields: list[str]
    encodings: list["InstructionTableEncoding"]


@dataclass
class InstructionTableEncoding:
    encname: str
    undef: bool
    iformname: str
    bitfields: list[Optional[str]]


def parse_index(root):
    index = root.xpath("//encodingindex[@instructionset='A64']/hierarchy")[0]
    return parse_index_node(index)


def parse_index_node(root):
    groupname = root.get("groupname")
    iclass = root.get("iclass")
    unallocated = root.get("unallocated") == "1"
    header = ''.join(e.text for e in root.xpath("header") if e.text is not None)
    decode = parse_boxes(root, "decode")
    regdiagram = parse_boxes(root, "regdiagram")
    nodes = [parse_index_node(e) for e in root.xpath("node")]
    return IndexNode(groupname, iclass, unallocated, header, regdiagram, decode, nodes)


def parse_boxes(root: ElementTree, name: str) -> Optional[list[Box]]:
    container = root.xpath(name)
    return None if not container else [parse_box(e) for e in container[0].xpath("box")]


def parse_box(root: ElementTree) -> Box:
    hibit = int(root.attrib["hibit"])
    width = int(root.get("width", 1))
    name = root.get("name")
    c = [e.text for e in root.xpath("c")]
    value = ''.join(c) if len(c) != 1 or c[0] is not None else None
    return Box(hibit, width, name, value)


def parse_iclass(root: ElementTree, name: str, groups: list[str]) -> Iclass:
    sect = root.xpath(f"//iclass_sect[@id='{name}']")[0]
    title = sect.attrib["title"]
    regdiagram = parse_boxes(sect, "regdiagram")
    instructiontable = parse_iclass_instructiontable(sect.xpath("instructiontable")[0])
    return Iclass(name, title, groups, regdiagram, instructiontable)


def parse_iclass_instructiontable(node: ElementTree) -> InstructionTable:
    bitfields = [e.text for e in node.xpath("thead/tr[@id='heading2']/th[@class='bitfields']")]
    encodings = [parse_iclass_instructiontable_encoding(e) for e in
                 node.xpath("tbody/tr[@class='instructiontable']")]
    return InstructionTable(bitfields, encodings)


def parse_iclass_instructiontable_encoding(node: ElementTree) -> InstructionTableEncoding:
    encname = node.attrib["encname"]
    undef = node.get("undef") == "1"
    bitfields = [e.text for e in node.xpath("td[@class='bitfield']")]
    iformname = node.xpath("td[@class='iformname']")[0].text
    return InstructionTableEncoding(encname, undef, iformname, bitfields)


def main():
    if len(sys.argv) != 3:
        raise RuntimeError(f"usage: {sys.argv[0]} <encodingindex.xml> <output-dir>")

    xml_file = sys.argv[1]
    output_dir = sys.argv[2]

    prelude = "// This file was generated. DO NOT EDIT.\npackage arm64\n"

    doc = etree.parse(xml_file).getroot()

    index = parse_index(doc)

    iclasses = [parse_iclass(doc, name, groups) for name, groups in collect_iclasses(index)]
    iclass_index = {iclass.name: iclass for iclass in iclasses}

    with open(path.join(output_dir, "iclass.go"), "w") as f:
        print(prelude, file=f)
        print_iclasses(iclasses, f=f)
        print(file=f)
        print_encodings(iclasses, f=f)

    with open(path.join(output_dir, "decode.go"), "w") as f:
        print(prelude, file=f)
        print_decode_errors(f=f)
        print(file=f)
        print_decode_struct(iclasses, f=f)
        print(file=f)
        print_decode(index, "decode", iclass_index, f=f)

    with open(path.join(output_dir, "encode.go"), "w") as f:
        print(prelude, file=f)
        print_encode_errors(f=f)
        print(file=f)
        print_encode(index, iclass_index, f=f)


def collect_iclasses(index: IndexNode) -> list[tuple[str, list[str]]]:
    collect_iclass2(index, [], iclasses := [])
    return iclasses


def collect_iclass2(node: IndexNode, groups: list[str], iclasses: list[tuple[str, list[str]]]):
    if node.unallocated:
        return
    if node.iclass is not None:
        iclasses.append((node.iclass, groups))
    if node.groupname is not None:
        groups = [] + groups + [node.groupname]
    for child in node.nodes:
        collect_iclass2(child, groups, iclasses)


def print_iclasses(iclasses: list[Iclass], f=sys.stdout):
    print("type iclass int", file=f)
    print(file=f)
    print("const (", file=f)
    print("\ticlass_none iclass = iota", file=f)
    last_groups = None
    for iclass in iclasses:
        groups = ' > '.join(iclass.groups)
        if last_groups is None or groups != last_groups:
            print(file=f)
            print(f"\t// {groups}", file=f)
            last_groups = groups
        print(f"\ticlass_{iclass.name} // {iclass.title}", file=f)
    print(")", file=f)


def print_encodings(iclasses: list[Iclass], f=sys.stdout):
    print("type encoding int", file=f)
    print(file=f)
    print("const (", file=f)
    print("\tencoding_none encoding = iota", file=f)
    for iclass in iclasses:
        print(file=f)
        print(f"\t// iclass_{iclass.name}", file=f)
        for encoding in iclass.instructiontable.encodings:
            if encoding.undef:
                continue
            print(f"\tencoding_{encoding.encname} // {encoding.iformname}", file=f)
    print(")", file=f)


def print_decode_errors(f=sys.stdout):
    print("import \"errors\"", file=f)
    print(file=f)
    print("var errUnallocated = errors.New(\"unallocated\")", file=f)
    print("var errUnmatched = errors.New(\"unmatched\")", file=f)


def print_decode_struct(iclasses: list[Iclass], f=sys.stdout):
    print("type decoded struct {", file=f)
    print("\ticlass iclass", file=f)
    print("\tencoding encoding", file=f)
    print(file=f)
    vars = set()
    for iclass in iclasses:
        vars = vars | set(box.name for box in iclass.regdiagram if box.name is not None)
    for var in sorted(vars, key=lambda v: (v.lower(), v)):
        print(f"\t{var} uint32", file=f)
    print("}", file=f)


def print_decode(node: IndexNode, name: str, iclasses: dict[str, Iclass], f=sys.stdout):
    print(f"func {name}(ins uint32, d *decoded) (err error) {{", file=f)

    if node.iclass is None:
        print_decode_group_body(node, name, iclasses, f)
    else:
        print_decode_iclass_body(node, iclasses[node.iclass], f)

    print("}", file=f)

    if node.iclass is None:
        for child in node.nodes:
            if child.unallocated:
                continue
            print(file=f)
            print_decode(child, f"{name}_{child.name}", iclasses, f)


def print_decode_group_body(node, name: str, iclasses: dict[str, Iclass], f=sys.stdout):
    if node.regdiagram is not None:
        for child in node.regdiagram:
            if child.name is None:
                continue
            print(f"\t{child.name} := (ins >> {child.hibit}) & {hex(child.mask)}", file=f)
        print(file=f)

    for i, child in enumerate(node.nodes):
        conditions = []
        for box in child.decode:
            if box.name is None:
                raise RuntimeError("missing name")
            if box.value is None:
                continue
            conditions.append(predicate(box.name, box.value))
        if len(conditions) == 0:
            raise RuntimeError("no conditions")
        all_conditions = " && ".join(conditions)
        if i == 0:
            print("\tswitch {", file=f)
        print(f"\tcase {all_conditions}:", file=f)
        if child.unallocated:
            print(f"\t\terr = errUnallocated", file=f)
        else:
            # print(f"\t\t// {child.header}", file=f)
            print(f"\t\terr = {name}_{child.name}(ins, d)", file=f)

    print("\tdefault:", file=f)
    print("\t\terr = errUnmatched", file=f)
    print("\t}", file=f)
    print("\treturn", file=f)


def print_decode_iclass_body(node: IndexNode, iclass: Iclass, f=sys.stdout):
    print(f"\td.iclass = iclass_{iclass.name}", file=f)
    for child in iclass.regdiagram:
        if child.name is None:
            continue
        print(f"\td.{child.name} = (ins >> {child.hibit}) & {hex(child.mask)}", file=f)

    table = iclass.instructiontable
    if len(table.encodings) == 0:
        raise RuntimeError(f"no encodings found for iclass {node.iclass}")
    elif len(table.encodings) == 1:
        child = table.encodings[0]
        if len(child.bitfields) != 0:
            raise RuntimeError("found single encoding with bitfields")
        if child.undef:
            print(f"\terr = errUnallocated", file=f)
        else:
            print(f"\td.encoding = encoding_{child.encname}", file=f)
        print("\treturn", file=f)
    else:
        print(file=f)
        catchall = None
        count = 0

        for i, child in enumerate(table.encodings):
            conditions = []
            for j, value in enumerate(child.bitfields):
                if value is None:
                    continue
                name = table.bitfields[j]
                conditions.append(predicate(f"d.{name}", value))

            if len(conditions) == 0:
                if catchall is not None:
                    raise RuntimeError("more than one child had no conditions")
                catchall = child
                continue

            all_conditions = " && ".join(conditions)
            if count == 0:
                print("\tswitch {", file=f)
            print(f"\tcase {all_conditions}:", file=f)

            count += 1

            if child.undef:
                print(f"\t\terr = errUnallocated", file=f)
            else:
                print(f"\t\td.encoding = encoding_{child.encname}", file=f)

        print("\tdefault:", file=f)
        if catchall is not None:
            print(f"\t\td.encoding = encoding_{catchall.encname}", file=f)
        else:
            print("\t\terr = errUnmatched", file=f)
        print("\t}", file=f)
        print("\treturn", file=f)


def predicate(name: str, value: str):
    op = "=="
    value = value
    if value.startswith("!= "):
        value = value[3:]
        op = "!="

    if len(set(value) - {"1", "0", "x"}) > 0:
        raise RuntimeError(f"value has unexpected chars: {value}")

    mask = 0
    match = 0
    for i, c in enumerate(reversed(value)):
        if c == "x":
            continue
        mask += 1 << i
        match += int(c) << i

    if "x" not in value:
        return f"{name} {op} {hex(match)}"
    return f"({name} & {hex(mask)}) {op} {hex(match)}"


def print_encode_errors(f=sys.stdout):
    print("import \"errors\"", file=f)
    print(file=f)
    print("var errOverflow    = errors.New(\"overflow\")", file=f)


def print_encode(node: IndexNode, iclasses: dict[str, Iclass], f=sys.stdout):
    if node.iclass is not None:
        iclass = iclasses[node.iclass]
        print_encode_iclass(iclass, f)
    for child in node.nodes:
        if child.unallocated:
            continue
        if child.groupname is None:
            print(file=f)
        print_encode(child, iclasses, f)


def print_encode_iclass(iclass: Iclass, f=sys.stdout):
    names = [box.name for box in iclass.regdiagram if box.name is not None]
    overflow_checks = ' || '.join(f"{box.name} > {hex(box.mask)}" for box in iclass.regdiagram if box.name is not None)
    input_bits = ' | '.join(
        shift_left(box.name, box.hibit - box.width + 1) for box in iclass.regdiagram if box.name is not None)
    iclass_bits = sum(int(box.value, 2) << (box.hibit - box.width + 1) for box in iclass.regdiagram if box.name is None)
    params = ", ".join(names) + " uint32" if names else ""

    print(f"func encode_{iclass.name}({params}) (ins uint32, err error) {{", file=f)

    count = 0
    if overflow_checks:
        print(f"\tswitch {{", file=f)
        count += 1
        print(f"\tcase {overflow_checks}:", file=f)
        print("\t\terr = errOverflow", file=f)

    catchall = None
    for i, child in enumerate(iclass.instructiontable.encodings):
        conditions = []
        for j, value in enumerate(child.bitfields):
            if value is None:
                continue
            name = iclass.instructiontable.bitfields[j]
            conditions.append(predicate(f"{name}", value))

        if len(conditions) == 0:
            if catchall is not None:
                raise RuntimeError("more than one child had no conditions")
            catchall = child
            continue

        all_conditions = " && ".join(conditions)
        if count == 0:
            print("\tswitch {", file=f)
        print(f"\tcase {all_conditions}:", file=f)

        count += 1

        if child.undef:
            print(f"\t\terr = errUnallocated", file=f)
        else:
            print(f"\t\t// encoding_{child.encname}", file=f)

    if count or catchall is not None:
        if not count and catchall:
            print(f"\t// encoding_{catchall.encname}", file=f)
        else:
            print("\tdefault:", file=f)
            if catchall is not None:
                print(f"\t\t// encoding_{catchall.encname}", file=f)
            else:
                print("\t\terr = errUnmatched", file=f)

            print("\t}", file=f)
        print(file=f)

    if input_bits:
        print(f"\tins |= {input_bits}", file=f)
    if iclass_bits:
        print(f"\tins |= {iclass_bits:#0{8}x}", file=f)
    if input_bits or iclass_bits:
        print(file=f)

    print("\treturn", file=f)
    print("}", file=f)


def shift_left(var: str, shift: int) -> str:
    return var if shift == 0 else f"({var} << {shift})"


if __name__ == "__main__":
    main()
