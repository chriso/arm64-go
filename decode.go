// This file was generated. DO NOT EDIT.
package arm64

import "errors"

var errUnallocated = errors.New("unallocated")
var errUnmatched = errors.New("unmatched")

type decoded struct {
	iclass   iclass
	encoding encoding

	A       uint32
	a       uint32
	B       uint32
	b       uint32
	b40     uint32
	b5      uint32
	c       uint32
	cmode   uint32
	cond    uint32
	CRm     uint32
	CRn     uint32
	D       uint32
	d       uint32
	dtype   uint32
	dtypeh  uint32
	dtypel  uint32
	e       uint32
	eq      uint32
	f       uint32
	ff      uint32
	g       uint32
	H       uint32
	h       uint32
	hw      uint32
	i1      uint32
	i2      uint32
	i3h     uint32
	i3l     uint32
	il      uint32
	imm12   uint32
	imm13   uint32
	imm14   uint32
	imm16   uint32
	imm19   uint32
	imm2    uint32
	imm26   uint32
	imm3    uint32
	imm4    uint32
	imm5    uint32
	imm5b   uint32
	imm6    uint32
	imm7    uint32
	imm8    uint32
	imm8h   uint32
	imm8l   uint32
	imm9    uint32
	imm9h   uint32
	imm9l   uint32
	immb    uint32
	immh    uint32
	immhi   uint32
	immlo   uint32
	immr    uint32
	imms    uint32
	L       uint32
	len     uint32
	LL      uint32
	lt      uint32
	M       uint32
	mask    uint32
	msz     uint32
	N       uint32
	ne      uint32
	nzcv    uint32
	O       uint32
	o0      uint32
	o1      uint32
	o2      uint32
	o3      uint32
	op      uint32
	Op0     uint32
	op1     uint32
	op2     uint32
	op21    uint32
	op3     uint32
	op31    uint32
	op4     uint32
	op54    uint32
	opc     uint32
	opc2    uint32
	opcode  uint32
	opcode2 uint32
	opt     uint32
	option  uint32
	pattern uint32
	Pd      uint32
	Pdm     uint32
	Pdn     uint32
	Pg      uint32
	Pm      uint32
	Pn      uint32
	prfop   uint32
	Pt      uint32
	ptype   uint32
	Q       uint32
	R       uint32
	Ra      uint32
	Rd      uint32
	Rdn     uint32
	Rm      uint32
	rmode   uint32
	Rn      uint32
	rot     uint32
	Rs      uint32
	Rt      uint32
	Rt2     uint32
	Rv      uint32
	rw      uint32
	S       uint32
	scale   uint32
	sf      uint32
	sh      uint32
	shift   uint32
	size    uint32
	ssz     uint32
	sz      uint32
	T       uint32
	tb      uint32
	tsz     uint32
	tszh    uint32
	tszl    uint32
	U       uint32
	u0      uint32
	u1      uint32
	uimm4   uint32
	uimm6   uint32
	uns     uint32
	V       uint32
	Vd      uint32
	Vdn     uint32
	Vm      uint32
	Vn      uint32
	W       uint32
	xs      uint32
	Za      uint32
	ZAda    uint32
	ZAt     uint32
	Zd      uint32
	Zda     uint32
	Zdn     uint32
	Zk      uint32
	Zm      uint32
	Zn      uint32
	Zt      uint32
}

func decode(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 31) & 0x1
	op1 := (ins >> 25) & 0xf

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_reserved(ins, d)
	case op0 == 0x1 && op1 == 0x0:
		err = decode_sme(ins, d)
	case op1 == 0x1:
		err = errUnallocated
	case op1 == 0x2:
		err = decode_sve(ins, d)
	case op1 == 0x3:
		err = errUnallocated
	case (op1 & 0xe) == 0x8:
		err = decode_dpimm(ins, d)
	case (op1 & 0xe) == 0xa:
		err = decode_control(ins, d)
	case (op1 & 0x5) == 0x4:
		err = decode_ldst(ins, d)
	case (op1 & 0x7) == 0x5:
		err = decode_dpreg(ins, d)
	case (op1 & 0x7) == 0x7:
		err = decode_simddp(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_reserved(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 29) & 0x3
	op1 := (ins >> 16) & 0x1ff

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_reserved_perm_undef(ins, d)
	case op1 != 0x0:
		err = errUnallocated
	case op0 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_reserved_perm_undef(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_perm_undef
	d.imm16 = (ins >> 0) & 0xffff
	d.encoding = encoding_UDF_only_perm_undef
	return
}

func decode_sme(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 29) & 0x3
	op1 := (ins >> 19) & 0x3f
	op2 := (ins >> 17) & 0x1
	op3 := (ins >> 9) & 0x1
	op4 := (ins >> 2) & 0x7

	switch {
	case (op0&0x2) == 0x0 && (op1&0x10) == 0x0:
		err = errUnallocated
	case (op0&0x2) == 0x0 && (op1&0x18) == 0x10 && (op4&0x1) == 0x0:
		err = decode_sme_mortlach_32bit_prod(ins, d)
	case (op0&0x2) == 0x0 && (op1&0x18) == 0x10 && (op4&0x1) == 0x1:
		err = errUnallocated
	case (op0&0x2) == 0x0 && (op1&0x18) == 0x18 && (op4&0x2) == 0x0:
		err = decode_sme_mortlach_64bit_prod(ins, d)
	case (op0&0x2) == 0x0 && (op1&0x18) == 0x18 && (op4&0x2) == 0x2:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x27) == 0x0 && op2 == 0x0 && (op4&0x4) == 0x0:
		err = decode_sme_mortlach_ins(ins, d)
	case op0 == 0x2 && (op1&0x27) == 0x0 && op2 == 0x0 && (op4&0x4) == 0x4:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x27) == 0x0 && op2 == 0x1 && op3 == 0x0:
		err = decode_sme_mortlach_ext(ins, d)
	case op0 == 0x2 && (op1&0x27) == 0x0 && op2 == 0x1 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x27) == 0x1:
		err = decode_sme_mortlach_misc(ins, d)
	case op0 == 0x2 && (op1&0x27) == 0x2 && (op4&0x2) == 0x0:
		err = decode_sme_mortlach_hvadd(ins, d)
	case op0 == 0x2 && (op1&0x27) == 0x2 && (op4&0x2) == 0x2:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x27) == 0x3:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x24) == 0x4:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x20) == 0x20:
		err = errUnallocated
	case op0 == 0x3:
		err = decode_sme_mortlach_mem(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_32bit_prod(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 29) & 0x1
	op1 := (ins >> 24) & 0x1
	op2 := (ins >> 21) & 0x1
	op3 := (ins >> 3) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0 && op3 == 0x0:
		err = decode_sme_mortlach_32bit_prod_mortlach_f32f32_prod(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x1 && op2 == 0x0 && op3 == 0x0:
		err = decode_sme_mortlach_32bit_prod_mortlach_b16f32_prod(ins, d)
	case op0 == 0x0 && op1 == 0x1 && op2 == 0x1 && op3 == 0x0:
		err = decode_sme_mortlach_32bit_prod_mortlach_f16f32_prod(ins, d)
	case op0 == 0x0 && op1 == 0x1 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x1 && op3 == 0x0:
		err = decode_sme_mortlach_32bit_prod_mortlach_i8i32_prod(ins, d)
	case op0 == 0x1 && op3 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_32bit_prod_mortlach_f32f32_prod(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_f32f32_prod
	d.Zm = (ins >> 16) & 0x1f
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.S = (ins >> 4) & 0x1
	d.ZAda = (ins >> 0) & 0x3

	switch {
	case d.S == 0x0:
		d.encoding = encoding_fmopa_za_pp_zz_32
	case d.S == 0x1:
		d.encoding = encoding_fmops_za_pp_zz_32
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_32bit_prod_mortlach_b16f32_prod(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_b16f32_prod
	d.Zm = (ins >> 16) & 0x1f
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.S = (ins >> 4) & 0x1
	d.ZAda = (ins >> 0) & 0x3

	switch {
	case d.S == 0x0:
		d.encoding = encoding_bfmopa_za32_pp_zz_
	case d.S == 0x1:
		d.encoding = encoding_bfmops_za32_pp_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_32bit_prod_mortlach_f16f32_prod(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_f16f32_prod
	d.Zm = (ins >> 16) & 0x1f
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.S = (ins >> 4) & 0x1
	d.ZAda = (ins >> 0) & 0x3

	switch {
	case d.S == 0x0:
		d.encoding = encoding_fmopa_za32_pp_zz_16
	case d.S == 0x1:
		d.encoding = encoding_fmops_za32_pp_zz_16
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_32bit_prod_mortlach_i8i32_prod(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_i8i32_prod
	d.u0 = (ins >> 24) & 0x1
	d.u1 = (ins >> 21) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.S = (ins >> 4) & 0x1
	d.ZAda = (ins >> 0) & 0x3

	switch {
	case d.u0 == 0x0 && d.u1 == 0x0 && d.S == 0x0:
		d.encoding = encoding_smopa_za_pp_zz_32
	case d.u0 == 0x0 && d.u1 == 0x0 && d.S == 0x1:
		d.encoding = encoding_smops_za_pp_zz_32
	case d.u0 == 0x0 && d.u1 == 0x1 && d.S == 0x0:
		d.encoding = encoding_sumopa_za_pp_zz_32
	case d.u0 == 0x0 && d.u1 == 0x1 && d.S == 0x1:
		d.encoding = encoding_sumops_za_pp_zz_32
	case d.u0 == 0x1 && d.u1 == 0x0 && d.S == 0x0:
		d.encoding = encoding_usmopa_za_pp_zz_32
	case d.u0 == 0x1 && d.u1 == 0x0 && d.S == 0x1:
		d.encoding = encoding_usmops_za_pp_zz_32
	case d.u0 == 0x1 && d.u1 == 0x1 && d.S == 0x0:
		d.encoding = encoding_umopa_za_pp_zz_32
	case d.u0 == 0x1 && d.u1 == 0x1 && d.S == 0x1:
		d.encoding = encoding_umops_za_pp_zz_32
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_64bit_prod(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 29) & 0x1
	op1 := (ins >> 24) & 0x1
	op2 := (ins >> 21) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0:
		err = decode_sme_mortlach_64bit_prod_mortlach_f64f64_prod(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x1:
		err = errUnallocated
	case op0 == 0x1:
		err = decode_sme_mortlach_64bit_prod_mortlach_i16i64_prod(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_64bit_prod_mortlach_f64f64_prod(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_f64f64_prod
	d.Zm = (ins >> 16) & 0x1f
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.S = (ins >> 4) & 0x1
	d.ZAda = (ins >> 0) & 0x7

	switch {
	case d.S == 0x0:
		d.encoding = encoding_fmopa_za_pp_zz_64
	case d.S == 0x1:
		d.encoding = encoding_fmops_za_pp_zz_64
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_64bit_prod_mortlach_i16i64_prod(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_i16i64_prod
	d.u0 = (ins >> 24) & 0x1
	d.u1 = (ins >> 21) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.S = (ins >> 4) & 0x1
	d.ZAda = (ins >> 0) & 0x7

	switch {
	case d.u0 == 0x0 && d.u1 == 0x0 && d.S == 0x0:
		d.encoding = encoding_smopa_za_pp_zz_64
	case d.u0 == 0x0 && d.u1 == 0x0 && d.S == 0x1:
		d.encoding = encoding_smops_za_pp_zz_64
	case d.u0 == 0x0 && d.u1 == 0x1 && d.S == 0x0:
		d.encoding = encoding_sumopa_za_pp_zz_64
	case d.u0 == 0x0 && d.u1 == 0x1 && d.S == 0x1:
		d.encoding = encoding_sumops_za_pp_zz_64
	case d.u0 == 0x1 && d.u1 == 0x0 && d.S == 0x0:
		d.encoding = encoding_usmopa_za_pp_zz_64
	case d.u0 == 0x1 && d.u1 == 0x0 && d.S == 0x1:
		d.encoding = encoding_usmops_za_pp_zz_64
	case d.u0 == 0x1 && d.u1 == 0x1 && d.S == 0x0:
		d.encoding = encoding_umopa_za_pp_zz_64
	case d.u0 == 0x1 && d.u1 == 0x1 && d.S == 0x1:
		d.encoding = encoding_umops_za_pp_zz_64
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_ins(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sme_mortlach_ins_mortlach_insert_pred(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_ins_mortlach_insert_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_insert_pred
	d.size = (ins >> 22) & 0x3
	d.Q = (ins >> 16) & 0x1
	d.V = (ins >> 15) & 0x1
	d.Rs = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.opc = (ins >> 0) & 0xf

	switch {
	case (d.size&0x2) == 0x0 && d.Q == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.Q == 0x0:
		d.encoding = encoding_mova_za_p_rz_b
	case d.size == 0x1 && d.Q == 0x0:
		d.encoding = encoding_mova_za_p_rz_h
	case d.size == 0x2 && d.Q == 0x0:
		d.encoding = encoding_mova_za_p_rz_w
	case d.size == 0x2 && d.Q == 0x1:
		err = errUnallocated
	case d.size == 0x3 && d.Q == 0x0:
		d.encoding = encoding_mova_za_p_rz_d
	case d.size == 0x3 && d.Q == 0x1:
		d.encoding = encoding_mova_za_p_rz_q
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_ext(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sme_mortlach_ext_mortlach_extract_pred(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_ext_mortlach_extract_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_extract_pred
	d.size = (ins >> 22) & 0x3
	d.Q = (ins >> 16) & 0x1
	d.V = (ins >> 15) & 0x1
	d.Rs = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.opc = (ins >> 5) & 0xf
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.Q == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.Q == 0x0:
		d.encoding = encoding_mova_z_p_rza_b
	case d.size == 0x1 && d.Q == 0x0:
		d.encoding = encoding_mova_z_p_rza_h
	case d.size == 0x2 && d.Q == 0x0:
		d.encoding = encoding_mova_z_p_rza_w
	case d.size == 0x2 && d.Q == 0x1:
		err = errUnallocated
	case d.size == 0x3 && d.Q == 0x0:
		d.encoding = encoding_mova_z_p_rza_d
	case d.size == 0x3 && d.Q == 0x1:
		d.encoding = encoding_mova_z_p_rza_q
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_misc(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 22) & 0x3
	op1 := (ins >> 8) & 0x7ff

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_sme_mortlach_misc_mortlach_zero(ins, d)
	case op0 == 0x0 && op1 != 0x0:
		err = errUnallocated
	case op0 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_misc_mortlach_zero(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_zero
	d.imm8 = (ins >> 0) & 0xff
	d.encoding = encoding_zero_za_i_
	return
}

func decode_sme_mortlach_hvadd(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 17) & 0x3
	op2 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x0:
		err = decode_sme_mortlach_hvadd_mortlach_addhv(ins, d)
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x1 && op1 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_hvadd_mortlach_addhv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_addhv
	d.op = (ins >> 22) & 0x1
	d.V = (ins >> 16) & 0x1
	d.Pm = (ins >> 13) & 0x7
	d.Pn = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.opc2 = (ins >> 0) & 0x7

	switch {
	case d.op == 0x0 && (d.opc2&0x4) == 0x4:
		err = errUnallocated
	case d.op == 0x0 && d.V == 0x0 && (d.opc2&0x4) == 0x0:
		d.encoding = encoding_addha_za_pp_z_32
	case d.op == 0x0 && d.V == 0x1 && (d.opc2&0x4) == 0x0:
		d.encoding = encoding_addva_za_pp_z_32
	case d.op == 0x1 && d.V == 0x0:
		d.encoding = encoding_addha_za_pp_z_64
	case d.op == 0x1 && d.V == 0x1:
		d.encoding = encoding_addva_za_pp_z_64
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_mem(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 21) & 0xf
	op1 := (ins >> 15) & 0x3f
	op2 := (ins >> 10) & 0x7
	op3 := (ins >> 4) & 0x1

	switch {
	case (op0&0x9) == 0x0 && op3 == 0x0:
		err = decode_sme_mortlach_mem_mortlach_contig_load(ins, d)
	case (op0&0x9) == 0x1 && op3 == 0x0:
		err = decode_sme_mortlach_mem_mortlach_contig_store(ins, d)
	case (op0&0x8) == 0x0 && op3 == 0x1:
		err = errUnallocated
	case (op0&0xe) == 0x8 && op1 == 0x0 && op2 == 0x0 && op3 == 0x0:
		err = decode_sme_mortlach_mem_mortlach_ctxt_ldst(ins, d)
	case (op0&0xe) == 0x8 && op1 == 0x0 && op2 == 0x0 && op3 == 0x1:
		err = errUnallocated
	case (op0&0xe) == 0x8 && op1 == 0x0 && op2 != 0x0:
		err = errUnallocated
	case (op0&0xe) == 0x8 && op1 != 0x0:
		err = errUnallocated
	case (op0 & 0xe) == 0xa:
		err = errUnallocated
	case (op0 & 0xe) == 0xc:
		err = errUnallocated
	case op0 == 0xe && op3 == 0x0:
		err = decode_sme_mortlach_mem_mortlach_contig_qload(ins, d)
	case op0 == 0xf && op3 == 0x0:
		err = decode_sme_mortlach_mem_mortlach_contig_qstore(ins, d)
	case (op0&0xe) == 0xe && op3 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_mem_mortlach_contig_load(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_contig_load
	d.msz = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.V = (ins >> 15) & 0x1
	d.Rs = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.opc = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_ld1b_za_p_rrr_
	case d.msz == 0x1:
		d.encoding = encoding_ld1h_za_p_rrr_
	case d.msz == 0x2:
		d.encoding = encoding_ld1w_za_p_rrr_
	case d.msz == 0x3:
		d.encoding = encoding_ld1d_za_p_rrr_
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_mem_mortlach_contig_store(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_contig_store
	d.msz = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.V = (ins >> 15) & 0x1
	d.Rs = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.opc = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_za_p_rrr_
	case d.msz == 0x1:
		d.encoding = encoding_st1h_za_p_rrr_
	case d.msz == 0x2:
		d.encoding = encoding_st1w_za_p_rrr_
	case d.msz == 0x3:
		d.encoding = encoding_st1d_za_p_rrr_
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_mem_mortlach_ctxt_ldst(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_ctxt_ldst
	d.op = (ins >> 21) & 0x1
	d.Rv = (ins >> 13) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.imm4 = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0:
		d.encoding = encoding_ldr_za_ri_
	case d.op == 0x1:
		d.encoding = encoding_str_za_ri_
	default:
		err = errUnmatched
	}
	return
}

func decode_sme_mortlach_mem_mortlach_contig_qload(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_contig_qload
	d.Rm = (ins >> 16) & 0x1f
	d.V = (ins >> 15) & 0x1
	d.Rs = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.ZAt = (ins >> 0) & 0xf
	d.encoding = encoding_ld1q_za_p_rrr_
	return
}

func decode_sme_mortlach_mem_mortlach_contig_qstore(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_mortlach_contig_qstore
	d.Rm = (ins >> 16) & 0x1f
	d.V = (ins >> 15) & 0x1
	d.Rs = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.ZAt = (ins >> 0) & 0xf
	d.encoding = encoding_st1q_za_p_rrr_
	return
}

func decode_sve(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 29) & 0x7
	op1 := (ins >> 23) & 0x3
	op2 := (ins >> 17) & 0x1f
	op3 := (ins >> 10) & 0x3f
	op4 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x10) == 0x10:
		err = decode_sve_sve_int_muladd_pred(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x38) == 0x0:
		err = decode_sve_sve_int_pred_bin(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x38) == 0x8:
		err = decode_sve_sve_int_pred_red(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x38) == 0x20:
		err = decode_sve_sve_int_pred_shift(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x38) == 0x28:
		err = decode_sve_sve_int_pred_un(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x38) == 0x0:
		err = decode_sve_sve_int_bin_cons_arit_0(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x38) == 0x8:
		err = decode_sve_sve_int_unpred_logical(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0x10:
		err = decode_sve_sve_index(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0x14:
		err = decode_sve_sve_alloca(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x38) == 0x18:
		err = decode_sve_sve_int_unpred_arit_b(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x38) == 0x20:
		err = decode_sve_sve_int_unpred_shift(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0x28:
		err = decode_sve_sve_int_bin_cons_misc_0_a(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0x2c:
		err = decode_sve_sve_int_unpred_misc(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x30) == 0x30:
		err = decode_sve_sve_countelt(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x18) == 0x0:
		err = decode_sve_sve_maskimm(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x18) == 0x8:
		err = decode_sve_sve_wideimm_pred(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && op3 == 0x8:
		err = decode_sve_sve_int_perm_dup_i(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && op3 == 0x9:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x3e) == 0xa:
		err = decode_sve_sve_int_perm_tbl_3src(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x3d) == 0xd:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && op3 == 0xc:
		err = decode_sve_sve_int_perm_tbl(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && op3 == 0xe:
		err = decode_sve_sve_perm_unpred_d(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x10:
		err = decode_sve_sve_perm_predicates(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x18:
		err = decode_sve_sve_int_perm_bin_perm_zz(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x30) == 0x20:
		err = decode_sve_sve_perm_pred(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x30) == 0x30:
		err = decode_sve_sve_int_sel_vvv(ins, d)
	case op0 == 0x0 && op1 == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x0:
		err = decode_sve_sve_perm_extract(ins, d)
	case op0 == 0x0 && op1 == 0x3 && (op2&0x10) == 0x10 && (op3&0x38) == 0x0:
		err = decode_sve_sve_perm_inter_long(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0:
		err = decode_sve_sve_cmpvec(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10:
		err = decode_sve_sve_int_ucmp_vi(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x10) == 0x0:
		err = decode_sve_sve_int_scmp_vi(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x18) == 0x0 && (op3&0x30) == 0x10:
		err = decode_sve_sve_int_pred_log(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x18) == 0x0 && (op3&0x30) == 0x30:
		err = decode_sve_sve_pred_gen_b(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x18) == 0x8 && (op3&0x30) == 0x10:
		err = decode_sve_sve_pred_gen_c(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x18) == 0x8 && (op3&0x30) == 0x30:
		err = decode_sve_sve_pred_gen_d(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x30) == 0x0:
		err = decode_sve_sve_cmpgpr(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x30) == 0x10 && op4 == 0x0:
		err = decode_sve_sve_int_pred_dup(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x30) == 0x10 && op4 == 0x1:
		err = errUnallocated
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x30) == 0x30:
		err = decode_sve_sve_wideimm_unpred(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x10 && (op3&0x30) == 0x20:
		err = decode_sve_sve_pred_count_a(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x14 && (op3&0x3c) == 0x20:
		err = decode_sve_sve_pred_count_b(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x14 && (op3&0x3c) == 0x24:
		err = decode_sve_sve_pred_wrffr(ins, d)
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x14 && (op3&0x38) == 0x28:
		err = errUnallocated
	case op0 == 0x1 && (op1&0x2) == 0x2 && (op2&0x18) == 0x18 && (op3&0x30) == 0x20:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x20) == 0x0:
		err = decode_sve_sve_intx_muladd_unpred(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x30) == 0x20:
		err = decode_sve_sve_intx_predicated(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x3e) == 0x30:
		err = decode_sve_sve_intx_clamp(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x3e) == 0x32:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x3c) == 0x34:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x38) == 0x38:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10:
		err = decode_sve_sve_intx_by_indexed_elem(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x20) == 0x0:
		err = decode_sve_sve_intx_cons_widening(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x30) == 0x20:
		err = decode_sve_sve_intx_constructive(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x30) == 0x30:
		err = decode_sve_sve_intx_acc(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x20) == 0x0:
		err = decode_sve_sve_intx_narrowing(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x20:
		err = decode_sve_sve_intx_match(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x28:
		err = decode_sve_sve_intx_histseg(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x30:
		err = decode_sve_sve_intx_histcnt(ins, d)
	case op0 == 0x2 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10 && (op3&0x38) == 0x38:
		err = decode_sve_sve_intx_crypto(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x0 && (op3&0x20) == 0x0:
		err = decode_sve_sve_fp_fcmla(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1a) == 0x2 && (op3&0x20) == 0x20:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x0 && (op3&0x38) == 0x20:
		err = decode_sve_sve_fp_fcadd(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x0 && (op3&0x38) == 0x28:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x0 && (op3&0x30) == 0x30:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x1 && (op3&0x20) == 0x20:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1e) == 0x4 && (op3&0x38) == 0x20:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1e) == 0x4 && (op3&0x38) == 0x28:
		err = decode_sve_sve_fp_fcvt2(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1e) == 0x4 && (op3&0x30) == 0x30:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1c) == 0x8 && (op3&0x38) == 0x20:
		err = decode_sve_sve_fp_pairwise(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1c) == 0x8 && (op3&0x38) == 0x28:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1c) == 0x8 && (op3&0x30) == 0x30:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x1c) == 0xc && (op3&0x20) == 0x20:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x16) == 0x2:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3e) == 0x0:
		err = decode_sve_sve_fp_fma_by_indexed_elem(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0x4:
		err = decode_sve_sve_fp_fcmla_by_indexed_elem(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && op3 == 0x8:
		err = decode_sve_sve_fp_fmul_by_indexed_elem(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && op3 == 0x9:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0xc:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x34) == 0x10:
		err = decode_sve_sve_fp_fma_w_by_indexed_elem(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x34) == 0x14:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x36) == 0x20:
		err = decode_sve_sve_fp_fma_w(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x34) == 0x24:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x38) == 0x30:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && op3 == 0x38:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && op3 == 0x39:
		err = decode_sve_sve_fp_fmmla(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3e) == 0x3a:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x10) == 0x10 && (op3&0x3c) == 0x3c:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x10) == 0x10:
		err = decode_sve_sve_fp_3op_p_pd(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x38) == 0x0:
		err = decode_sve_sve_fp_3op_u_zd(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x38) == 0x20:
		err = decode_sve_sve_fp_pred(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x10) == 0x0 && (op3&0x38) == 0x28:
		err = decode_sve_sve_fp_unary(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x0 && (op3&0x38) == 0x8:
		err = decode_sve_sve_fp_fast_red(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x4 && (op3&0x3c) == 0x8:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x4 && (op3&0x3c) == 0xc:
		err = decode_sve_sve_fp_unary_unpred(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x1c) == 0x8 && (op3&0x38) == 0x8:
		err = decode_sve_sve_fp_cmpzero(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x1c) == 0xc && (op3&0x38) == 0x8:
		err = decode_sve_sve_fp_slowreduce(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x10) == 0x10:
		err = decode_sve_sve_fp_fma(ins, d)
	case op0 == 0x4:
		err = decode_sve_sve_mem32(ins, d)
	case op0 == 0x5:
		err = decode_sve_sve_memcld(ins, d)
	case op0 == 0x6:
		err = decode_sve_sve_mem64(ins, d)
	case op0 == 0x7 && (op3&0x28) == 0x0:
		err = decode_sve_sve_memst_cs(ins, d)
	case op0 == 0x7 && (op3&0x28) == 0x8:
		err = decode_sve_sve_memst_nt(ins, d)
	case op0 == 0x7 && (op3&0x28) == 0x20:
		err = decode_sve_sve_memst_ss(ins, d)
	case op0 == 0x7 && (op3&0x38) == 0x28:
		err = decode_sve_sve_memst_ss2(ins, d)
	case op0 == 0x7 && (op3&0x38) == 0x38:
		err = decode_sve_sve_memst_si(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_muladd_pred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 15) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_int_muladd_pred_sve_int_mlas_vvv_pred(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_int_muladd_pred_sve_int_mladdsub_vvv_pred(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_muladd_pred_sve_int_mlas_vvv_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_mlas_vvv_pred
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_mla_z_p_zzz_
	case d.op == 0x1:
		d.encoding = encoding_mls_z_p_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_muladd_pred_sve_int_mladdsub_vvv_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_mladdsub_vvv_pred
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Za = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_mad_z_p_zzz_
	case d.op == 0x1:
		d.encoding = encoding_msb_z_p_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_bin(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x7

	switch {
	case (op0 & 0x6) == 0x0:
		err = decode_sve_sve_int_pred_bin_sve_int_bin_pred_arit_0(ins, d)
	case (op0 & 0x6) == 0x2:
		err = decode_sve_sve_int_pred_bin_sve_int_bin_pred_arit_1(ins, d)
	case op0 == 0x4:
		err = decode_sve_sve_int_pred_bin_sve_int_bin_pred_arit_2(ins, d)
	case op0 == 0x5:
		err = decode_sve_sve_int_pred_bin_sve_int_bin_pred_div(ins, d)
	case (op0 & 0x6) == 0x6:
		err = decode_sve_sve_int_pred_bin_sve_int_bin_pred_log(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_bin_sve_int_bin_pred_arit_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_arit_0
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_add_z_p_zz_
	case d.opc == 0x1:
		d.encoding = encoding_sub_z_p_zz_
	case d.opc == 0x2:
		err = errUnallocated
	case d.opc == 0x3:
		d.encoding = encoding_subr_z_p_zz_
	case (d.opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_bin_sve_int_bin_pred_arit_1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_arit_1
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 17) & 0x3
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.U == 0x0:
		d.encoding = encoding_smax_z_p_zz_
	case d.opc == 0x0 && d.U == 0x1:
		d.encoding = encoding_umax_z_p_zz_
	case d.opc == 0x1 && d.U == 0x0:
		d.encoding = encoding_smin_z_p_zz_
	case d.opc == 0x1 && d.U == 0x1:
		d.encoding = encoding_umin_z_p_zz_
	case d.opc == 0x2 && d.U == 0x0:
		d.encoding = encoding_sabd_z_p_zz_
	case d.opc == 0x2 && d.U == 0x1:
		d.encoding = encoding_uabd_z_p_zz_
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_bin_sve_int_bin_pred_arit_2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_arit_2
	d.size = (ins >> 22) & 0x3
	d.H = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.H == 0x0 && d.U == 0x0:
		d.encoding = encoding_mul_z_p_zz_
	case d.H == 0x0 && d.U == 0x1:
		err = errUnallocated
	case d.H == 0x1 && d.U == 0x0:
		d.encoding = encoding_smulh_z_p_zz_
	case d.H == 0x1 && d.U == 0x1:
		d.encoding = encoding_umulh_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_bin_sve_int_bin_pred_div(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_div
	d.size = (ins >> 22) & 0x3
	d.R = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.R == 0x0 && d.U == 0x0:
		d.encoding = encoding_sdiv_z_p_zz_
	case d.R == 0x0 && d.U == 0x1:
		d.encoding = encoding_udiv_z_p_zz_
	case d.R == 0x1 && d.U == 0x0:
		d.encoding = encoding_sdivr_z_p_zz_
	case d.R == 0x1 && d.U == 0x1:
		d.encoding = encoding_udivr_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_bin_sve_int_bin_pred_log(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_log
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_orr_z_p_zz_
	case d.opc == 0x1:
		d.encoding = encoding_eor_z_p_zz_
	case d.opc == 0x2:
		d.encoding = encoding_and_z_p_zz_
	case d.opc == 0x3:
		d.encoding = encoding_bic_z_p_zz_
	case (d.opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_red(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x7

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_int_pred_red_sve_int_reduce_0(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_int_pred_red_sve_int_reduce_1(ins, d)
	case (op0 & 0x5) == 0x1:
		err = errUnallocated
	case (op0 & 0x6) == 0x4:
		err = decode_sve_sve_int_pred_red_sve_int_movprfx_pred(ins, d)
	case op0 == 0x6:
		err = decode_sve_sve_int_pred_red_sve_int_reduce_2(ins, d)
	case op0 == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_red_sve_int_reduce_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_reduce_0
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Vd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.U == 0x0:
		d.encoding = encoding_saddv_r_p_z_
	case d.op == 0x0 && d.U == 0x1:
		d.encoding = encoding_uaddv_r_p_z_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_red_sve_int_reduce_1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_reduce_1
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Vd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.U == 0x0:
		d.encoding = encoding_smaxv_r_p_z_
	case d.op == 0x0 && d.U == 0x1:
		d.encoding = encoding_umaxv_r_p_z_
	case d.op == 0x1 && d.U == 0x0:
		d.encoding = encoding_sminv_r_p_z_
	case d.op == 0x1 && d.U == 0x1:
		d.encoding = encoding_uminv_r_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_red_sve_int_movprfx_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_movprfx_pred
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 17) & 0x3
	d.M = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_movprfx_z_p_z_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_red_sve_int_reduce_2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_reduce_2
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Vd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_orv_r_p_z_
	case d.opc == 0x1:
		d.encoding = encoding_eorv_r_p_z_
	case d.opc == 0x2:
		d.encoding = encoding_andv_r_p_z_
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_shift(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 19) & 0x3

	switch {
	case (op0 & 0x2) == 0x0:
		err = decode_sve_sve_int_pred_shift_sve_int_bin_pred_shift_0(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_int_pred_shift_sve_int_bin_pred_shift_1(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_int_pred_shift_sve_int_bin_pred_shift_2(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_shift_sve_int_bin_pred_shift_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_shift_0
	d.tszh = (ins >> 22) & 0x3
	d.opc = (ins >> 18) & 0x3
	d.L = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.tszl = (ins >> 8) & 0x3
	d.imm3 = (ins >> 5) & 0x7
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.L == 0x0 && d.U == 0x0:
		d.encoding = encoding_asr_z_p_zi_
	case d.opc == 0x0 && d.L == 0x0 && d.U == 0x1:
		d.encoding = encoding_lsr_z_p_zi_
	case d.opc == 0x0 && d.L == 0x1 && d.U == 0x0:
		err = errUnallocated
	case d.opc == 0x0 && d.L == 0x1 && d.U == 0x1:
		d.encoding = encoding_lsl_z_p_zi_
	case d.opc == 0x1 && d.L == 0x0 && d.U == 0x0:
		d.encoding = encoding_asrd_z_p_zi_
	case d.opc == 0x1 && d.L == 0x0 && d.U == 0x1:
		err = errUnallocated
	case d.opc == 0x1 && d.L == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqshl_z_p_zi_
	case d.opc == 0x1 && d.L == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqshl_z_p_zi_
	case d.opc == 0x2:
		err = errUnallocated
	case d.opc == 0x3 && d.L == 0x0 && d.U == 0x0:
		d.encoding = encoding_srshr_z_p_zi_
	case d.opc == 0x3 && d.L == 0x0 && d.U == 0x1:
		d.encoding = encoding_urshr_z_p_zi_
	case d.opc == 0x3 && d.L == 0x1 && d.U == 0x0:
		err = errUnallocated
	case d.opc == 0x3 && d.L == 0x1 && d.U == 0x1:
		d.encoding = encoding_sqshlu_z_p_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_shift_sve_int_bin_pred_shift_1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_shift_1
	d.size = (ins >> 22) & 0x3
	d.R = (ins >> 18) & 0x1
	d.L = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x1 && d.U == 0x0:
		err = errUnallocated
	case d.R == 0x0 && d.L == 0x0 && d.U == 0x0:
		d.encoding = encoding_asr_z_p_zz_
	case d.R == 0x0 && d.L == 0x0 && d.U == 0x1:
		d.encoding = encoding_lsr_z_p_zz_
	case d.R == 0x0 && d.L == 0x1 && d.U == 0x1:
		d.encoding = encoding_lsl_z_p_zz_
	case d.R == 0x1 && d.L == 0x0 && d.U == 0x0:
		d.encoding = encoding_asrr_z_p_zz_
	case d.R == 0x1 && d.L == 0x0 && d.U == 0x1:
		d.encoding = encoding_lsrr_z_p_zz_
	case d.R == 0x1 && d.L == 0x1 && d.U == 0x1:
		d.encoding = encoding_lslr_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_shift_sve_int_bin_pred_shift_2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_pred_shift_2
	d.size = (ins >> 22) & 0x3
	d.R = (ins >> 18) & 0x1
	d.L = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.R == 0x0 && d.L == 0x0 && d.U == 0x0:
		d.encoding = encoding_asr_z_p_zw_
	case d.R == 0x0 && d.L == 0x0 && d.U == 0x1:
		d.encoding = encoding_lsr_z_p_zw_
	case d.R == 0x0 && d.L == 0x1 && d.U == 0x0:
		err = errUnallocated
	case d.R == 0x0 && d.L == 0x1 && d.U == 0x1:
		d.encoding = encoding_lsl_z_p_zw_
	case d.R == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_un(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 19) & 0x3

	switch {
	case (op0 & 0x2) == 0x0:
		err = errUnallocated
	case op0 == 0x2:
		err = decode_sve_sve_int_pred_un_sve_int_un_pred_arit_0(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_int_pred_un_sve_int_un_pred_arit_1(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_un_sve_int_un_pred_arit_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_un_pred_arit_0
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_sxtb_z_p_z_
	case d.opc == 0x1:
		d.encoding = encoding_uxtb_z_p_z_
	case d.opc == 0x2:
		d.encoding = encoding_sxth_z_p_z_
	case d.opc == 0x3:
		d.encoding = encoding_uxth_z_p_z_
	case d.opc == 0x4:
		d.encoding = encoding_sxtw_z_p_z_
	case d.opc == 0x5:
		d.encoding = encoding_uxtw_z_p_z_
	case d.opc == 0x6:
		d.encoding = encoding_abs_z_p_z_
	case d.opc == 0x7:
		d.encoding = encoding_neg_z_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_un_sve_int_un_pred_arit_1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_un_pred_arit_1
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_cls_z_p_z_
	case d.opc == 0x1:
		d.encoding = encoding_clz_z_p_z_
	case d.opc == 0x2:
		d.encoding = encoding_cnt_z_p_z_
	case d.opc == 0x3:
		d.encoding = encoding_cnot_z_p_z_
	case d.opc == 0x4:
		d.encoding = encoding_fabs_z_p_z_
	case d.opc == 0x5:
		d.encoding = encoding_fneg_z_p_z_
	case d.opc == 0x6:
		d.encoding = encoding_not_z_p_z_
	case d.opc == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_bin_cons_arit_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_arit_0
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_add_z_zz_
	case d.opc == 0x1:
		d.encoding = encoding_sub_z_zz_
	case (d.opc & 0x6) == 0x2:
		err = errUnallocated
	case d.opc == 0x4:
		d.encoding = encoding_sqadd_z_zz_
	case d.opc == 0x5:
		d.encoding = encoding_uqadd_z_zz_
	case d.opc == 0x6:
		d.encoding = encoding_sqsub_z_zz_
	case d.opc == 0x7:
		d.encoding = encoding_uqsub_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_logical(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x7

	switch {
	case (op0 & 0x4) == 0x0:
		err = errUnallocated
	case op0 == 0x4:
		err = decode_sve_sve_int_unpred_logical_sve_int_bin_cons_log(ins, d)
	case op0 == 0x5:
		err = decode_sve_sve_int_unpred_logical_sve_int_rotate_imm(ins, d)
	case (op0 & 0x6) == 0x6:
		err = decode_sve_sve_int_unpred_logical_sve_int_tern_log(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_logical_sve_int_bin_cons_log(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_log
	d.opc = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_and_z_zz_
	case d.opc == 0x1:
		d.encoding = encoding_orr_z_zz_
	case d.opc == 0x2:
		d.encoding = encoding_eor_z_zz_
	case d.opc == 0x3:
		d.encoding = encoding_bic_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_logical_sve_int_rotate_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_rotate_imm
	d.tszh = (ins >> 22) & 0x3
	d.tszl = (ins >> 19) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_xar_z_zzi_
	return
}

func decode_sve_sve_int_unpred_logical_sve_int_tern_log(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_tern_log
	d.opc = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.o2 = (ins >> 10) & 0x1
	d.Zk = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_eor3_z_zzz_
	case d.opc == 0x0 && d.o2 == 0x1:
		d.encoding = encoding_bsl_z_zzz_
	case d.opc == 0x1 && d.o2 == 0x0:
		d.encoding = encoding_bcax_z_zzz_
	case d.opc == 0x1 && d.o2 == 0x1:
		d.encoding = encoding_bsl1n_z_zzz_
	case (d.opc&0x2) == 0x2 && d.o2 == 0x0:
		err = errUnallocated
	case d.opc == 0x2 && d.o2 == 0x1:
		d.encoding = encoding_bsl2n_z_zzz_
	case d.opc == 0x3 && d.o2 == 0x1:
		d.encoding = encoding_nbsl_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_index(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x3

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_index_sve_int_index_ii(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_index_sve_int_index_ri(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_index_sve_int_index_ir(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_index_sve_int_index_rr(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_index_sve_int_index_ii(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_index_ii
	d.size = (ins >> 22) & 0x3
	d.imm5b = (ins >> 16) & 0x1f
	d.imm5 = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_index_z_ii_
	return
}

func decode_sve_sve_index_sve_int_index_ri(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_index_ri
	d.size = (ins >> 22) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_index_z_ri_
	return
}

func decode_sve_sve_index_sve_int_index_ir(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_index_ir
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.imm5 = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_index_z_ir_
	return
}

func decode_sve_sve_index_sve_int_index_rr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_index_rr
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_index_z_rr_
	return
}

func decode_sve_sve_alloca(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 11) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_sve_sve_alloca_sve_int_arith_vl(ins, d)
	case op0 == 0x0 && op1 == 0x1:
		err = decode_sve_sve_alloca_sve_int_arith_svl(ins, d)
	case op0 == 0x1 && op1 == 0x0:
		err = decode_sve_sve_alloca_sve_int_read_vl_a(ins, d)
	case op0 == 0x1 && op1 == 0x1:
		err = decode_sve_sve_alloca_sve_int_read_svl_a(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_alloca_sve_int_arith_vl(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_arith_vl
	d.op = (ins >> 22) & 0x1
	d.Rn = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 5) & 0x3f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_addvl_r_ri_
	case d.op == 0x1:
		d.encoding = encoding_addpl_r_ri_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_alloca_sve_int_arith_svl(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_arith_svl
	d.op = (ins >> 22) & 0x1
	d.Rn = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 5) & 0x3f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_addsvl_r_ri_
	case d.op == 0x1:
		d.encoding = encoding_addspl_r_ri_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_alloca_sve_int_read_vl_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_read_vl_a
	d.op = (ins >> 22) & 0x1
	d.opc2 = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 5) & 0x3f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && (d.opc2&0x10) == 0x0:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x18) == 0x10:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x1c) == 0x18:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x1e) == 0x1c:
		err = errUnallocated
	case d.op == 0x0 && d.opc2 == 0x1e:
		err = errUnallocated
	case d.op == 0x0 && d.opc2 == 0x1f:
		d.encoding = encoding_rdvl_r_i_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_alloca_sve_int_read_svl_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_read_svl_a
	d.op = (ins >> 22) & 0x1
	d.opc2 = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 5) & 0x3f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && (d.opc2&0x10) == 0x0:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x18) == 0x10:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x1c) == 0x18:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x1e) == 0x1c:
		err = errUnallocated
	case d.op == 0x0 && d.opc2 == 0x1e:
		err = errUnallocated
	case d.op == 0x0 && d.opc2 == 0x1f:
		d.encoding = encoding_rdsvl_r_i_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_arit_b(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 11) & 0x3

	switch {
	case (op0 & 0x2) == 0x0:
		err = decode_sve_sve_int_unpred_arit_b_sve_int_mul_b(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_int_unpred_arit_b_sve_int_sqdmulh(ins, d)
	case op0 == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_arit_b_sve_int_mul_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_mul_b
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_mul_z_zz_
	case d.opc == 0x2:
		d.encoding = encoding_smulh_z_zz_
	case d.opc == 0x3:
		d.encoding = encoding_umulh_z_zz_
	case d.size == 0x0 && d.opc == 0x1:
		d.encoding = encoding_pmul_z_zz_
	case d.size == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opc == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_arit_b_sve_int_sqdmulh(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_sqdmulh
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.R = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.R == 0x0:
		d.encoding = encoding_sqdmulh_z_zz_
	case d.R == 0x1:
		d.encoding = encoding_sqrdmulh_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_shift(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 12) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_int_unpred_shift_sve_int_bin_cons_shift_a(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_int_unpred_shift_sve_int_bin_cons_shift_b(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_shift_sve_int_bin_cons_shift_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_shift_a
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_asr_z_zw_
	case d.opc == 0x1:
		d.encoding = encoding_lsr_z_zw_
	case d.opc == 0x2:
		err = errUnallocated
	case d.opc == 0x3:
		d.encoding = encoding_lsl_z_zw_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_shift_sve_int_bin_cons_shift_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_shift_b
	d.tszh = (ins >> 22) & 0x3
	d.tszl = (ins >> 19) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.opc = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_asr_z_zi_
	case d.opc == 0x1:
		d.encoding = encoding_lsr_z_zi_
	case d.opc == 0x2:
		err = errUnallocated
	case d.opc == 0x3:
		d.encoding = encoding_lsl_z_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_bin_cons_misc_0_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_misc_0_a
	d.opc = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.msz = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_adr_z_az_d_s32_scaled
	case d.opc == 0x1:
		d.encoding = encoding_adr_z_az_d_u32_scaled
	case (d.opc & 0x2) == 0x2:
		d.encoding = encoding_adr_z_az_sd_same_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_misc(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x3

	switch {
	case (op0 & 0x2) == 0x0:
		err = decode_sve_sve_int_unpred_misc_sve_int_bin_cons_misc_0_b(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_int_unpred_misc_sve_int_bin_cons_misc_0_c(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_int_unpred_misc_sve_int_bin_cons_misc_0_d(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_misc_sve_int_bin_cons_misc_0_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_misc_0_b
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_ftssel_z_zz_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_misc_sve_int_bin_cons_misc_0_c(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_misc_0_c
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fexpa_z_z_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x1e) == 0x2:
		err = errUnallocated
	case (d.opc & 0x1c) == 0x4:
		err = errUnallocated
	case (d.opc & 0x18) == 0x8:
		err = errUnallocated
	case (d.opc & 0x10) == 0x10:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_unpred_misc_sve_int_bin_cons_misc_0_d(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_bin_cons_misc_0_d
	d.opc = (ins >> 22) & 0x3
	d.opc2 = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.opc2 == 0x0:
		d.encoding = encoding_movprfx_z_z_
	case d.opc == 0x0 && d.opc2 == 0x1:
		err = errUnallocated
	case d.opc == 0x0 && (d.opc2&0x1e) == 0x2:
		err = errUnallocated
	case d.opc == 0x0 && (d.opc2&0x1c) == 0x4:
		err = errUnallocated
	case d.opc == 0x0 && (d.opc2&0x18) == 0x8:
		err = errUnallocated
	case d.opc == 0x0 && (d.opc2&0x10) == 0x10:
		err = errUnallocated
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_countelt(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 20) & 0x1
	op1 := (ins >> 11) & 0x7

	switch {
	case op0 == 0x0 && (op1&0x6) == 0x0:
		err = decode_sve_sve_countelt_sve_int_countvlv0(ins, d)
	case op0 == 0x0 && op1 == 0x4:
		err = decode_sve_sve_countelt_sve_int_count(ins, d)
	case op0 == 0x0 && op1 == 0x5:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0:
		err = decode_sve_sve_countelt_sve_int_countvlv1(ins, d)
	case op0 == 0x1 && op1 == 0x4:
		err = decode_sve_sve_countelt_sve_int_pred_pattern_a(ins, d)
	case op0 == 0x1 && (op1&0x3) == 0x1:
		err = errUnallocated
	case (op1 & 0x6) == 0x2:
		err = errUnallocated
	case (op1 & 0x6) == 0x6:
		err = decode_sve_sve_countelt_sve_int_pred_pattern_b(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_countelt_sve_int_countvlv0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_countvlv0
	d.size = (ins >> 22) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.D = (ins >> 11) & 0x1
	d.U = (ins >> 10) & 0x1
	d.pattern = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0:
		err = errUnallocated
	case d.size == 0x1 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqinch_z_zs_
	case d.size == 0x1 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqinch_z_zs_
	case d.size == 0x1 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdech_z_zs_
	case d.size == 0x1 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdech_z_zs_
	case d.size == 0x2 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincw_z_zs_
	case d.size == 0x2 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincw_z_zs_
	case d.size == 0x2 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecw_z_zs_
	case d.size == 0x2 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecw_z_zs_
	case d.size == 0x3 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincd_z_zs_
	case d.size == 0x3 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincd_z_zs_
	case d.size == 0x3 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecd_z_zs_
	case d.size == 0x3 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecd_z_zs_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_countelt_sve_int_count(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_count
	d.size = (ins >> 22) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.op = (ins >> 10) & 0x1
	d.pattern = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.op == 0x0:
		d.encoding = encoding_cntb_r_s_
	case d.size == 0x1 && d.op == 0x0:
		d.encoding = encoding_cnth_r_s_
	case d.size == 0x2 && d.op == 0x0:
		d.encoding = encoding_cntw_r_s_
	case d.size == 0x3 && d.op == 0x0:
		d.encoding = encoding_cntd_r_s_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_countelt_sve_int_countvlv1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_countvlv1
	d.size = (ins >> 22) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.D = (ins >> 10) & 0x1
	d.pattern = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0:
		err = errUnallocated
	case d.size == 0x1 && d.D == 0x0:
		d.encoding = encoding_inch_z_zs_
	case d.size == 0x1 && d.D == 0x1:
		d.encoding = encoding_dech_z_zs_
	case d.size == 0x2 && d.D == 0x0:
		d.encoding = encoding_incw_z_zs_
	case d.size == 0x2 && d.D == 0x1:
		d.encoding = encoding_decw_z_zs_
	case d.size == 0x3 && d.D == 0x0:
		d.encoding = encoding_incd_z_zs_
	case d.size == 0x3 && d.D == 0x1:
		d.encoding = encoding_decd_z_zs_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_countelt_sve_int_pred_pattern_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pred_pattern_a
	d.size = (ins >> 22) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.D = (ins >> 10) & 0x1
	d.pattern = (ins >> 5) & 0x1f
	d.Rdn = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.D == 0x0:
		d.encoding = encoding_incb_r_rs_
	case d.size == 0x0 && d.D == 0x1:
		d.encoding = encoding_decb_r_rs_
	case d.size == 0x1 && d.D == 0x0:
		d.encoding = encoding_inch_r_rs_
	case d.size == 0x1 && d.D == 0x1:
		d.encoding = encoding_dech_r_rs_
	case d.size == 0x2 && d.D == 0x0:
		d.encoding = encoding_incw_r_rs_
	case d.size == 0x2 && d.D == 0x1:
		d.encoding = encoding_decw_r_rs_
	case d.size == 0x3 && d.D == 0x0:
		d.encoding = encoding_incd_r_rs_
	case d.size == 0x3 && d.D == 0x1:
		d.encoding = encoding_decd_r_rs_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_countelt_sve_int_pred_pattern_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pred_pattern_b
	d.size = (ins >> 22) & 0x3
	d.sf = (ins >> 20) & 0x1
	d.imm4 = (ins >> 16) & 0xf
	d.D = (ins >> 11) & 0x1
	d.U = (ins >> 10) & 0x1
	d.pattern = (ins >> 5) & 0x1f
	d.Rdn = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincb_r_rs_sx
	case d.size == 0x0 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincb_r_rs_uw
	case d.size == 0x0 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecb_r_rs_sx
	case d.size == 0x0 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecb_r_rs_uw
	case d.size == 0x0 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincb_r_rs_x
	case d.size == 0x0 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincb_r_rs_x
	case d.size == 0x0 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecb_r_rs_x
	case d.size == 0x0 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecb_r_rs_x
	case d.size == 0x1 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqinch_r_rs_sx
	case d.size == 0x1 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqinch_r_rs_uw
	case d.size == 0x1 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdech_r_rs_sx
	case d.size == 0x1 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdech_r_rs_uw
	case d.size == 0x1 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqinch_r_rs_x
	case d.size == 0x1 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqinch_r_rs_x
	case d.size == 0x1 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdech_r_rs_x
	case d.size == 0x1 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdech_r_rs_x
	case d.size == 0x2 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincw_r_rs_sx
	case d.size == 0x2 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincw_r_rs_uw
	case d.size == 0x2 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecw_r_rs_sx
	case d.size == 0x2 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecw_r_rs_uw
	case d.size == 0x2 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincw_r_rs_x
	case d.size == 0x2 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincw_r_rs_x
	case d.size == 0x2 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecw_r_rs_x
	case d.size == 0x2 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecw_r_rs_x
	case d.size == 0x3 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincd_r_rs_sx
	case d.size == 0x3 && d.sf == 0x0 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincd_r_rs_uw
	case d.size == 0x3 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecd_r_rs_sx
	case d.size == 0x3 && d.sf == 0x0 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecd_r_rs_uw
	case d.size == 0x3 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqincd_r_rs_x
	case d.size == 0x3 && d.sf == 0x1 && d.D == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqincd_r_rs_x
	case d.size == 0x3 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqdecd_r_rs_x
	case d.size == 0x3 && d.sf == 0x1 && d.D == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqdecd_r_rs_x
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_maskimm(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 22) & 0x3
	op1 := (ins >> 18) & 0x3

	switch {
	case op0 == 0x3 && op1 == 0x0:
		err = decode_sve_sve_maskimm_sve_int_dup_mask_imm(ins, d)
	case op0 != 0x3 && op1 == 0x0:
		err = decode_sve_sve_maskimm_sve_int_log_imm(ins, d)
	case op1 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_maskimm_sve_int_dup_mask_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_dup_mask_imm
	d.imm13 = (ins >> 5) & 0x1fff
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_dupm_z_i_
	return
}

func decode_sve_sve_maskimm_sve_int_log_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_log_imm
	d.opc = (ins >> 22) & 0x3
	d.imm13 = (ins >> 5) & 0x1fff
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_orr_z_zi_
	case d.opc == 0x1:
		d.encoding = encoding_eor_z_zi_
	case d.opc == 0x2:
		d.encoding = encoding_and_z_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_pred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 13) & 0x7

	switch {
	case (op0 & 0x4) == 0x0:
		err = decode_sve_sve_wideimm_pred_sve_int_dup_imm_pred(ins, d)
	case (op0 & 0x6) == 0x4:
		err = errUnallocated
	case op0 == 0x6:
		err = decode_sve_sve_wideimm_pred_sve_int_dup_fpimm_pred(ins, d)
	case op0 == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_pred_sve_int_dup_imm_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_dup_imm_pred
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 16) & 0xf
	d.M = (ins >> 14) & 0x1
	d.sh = (ins >> 13) & 0x1
	d.imm8 = (ins >> 5) & 0xff
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.M == 0x0:
		d.encoding = encoding_cpy_z_o_i_
	case d.M == 0x1:
		d.encoding = encoding_cpy_z_p_i_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_pred_sve_int_dup_fpimm_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_dup_fpimm_pred
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 16) & 0xf
	d.imm8 = (ins >> 5) & 0xff
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_fcpy_z_p_i_
	return
}

func decode_sve_sve_int_perm_dup_i(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_dup_i
	d.imm2 = (ins >> 22) & 0x3
	d.tsz = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_dup_z_zi_
	return
}

func decode_sve_sve_int_perm_tbl_3src(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_tbl_3src
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_tbl_z_zz_2
	case d.op == 0x1:
		d.encoding = encoding_tbx_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_perm_tbl(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_tbl
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_tbl_z_zz_1
	return
}

func decode_sve_sve_perm_unpred_d(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 19) & 0x3
	op1 := (ins >> 16) & 0x7

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_sve_sve_perm_unpred_d_sve_int_perm_dup_r(ins, d)
	case op0 == 0x0 && op1 == 0x4:
		err = decode_sve_sve_perm_unpred_d_sve_int_perm_insrs(ins, d)
	case op0 == 0x0 && (op1&0x3) == 0x2:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x1) == 0x1:
		err = errUnallocated
	case op0 == 0x1:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x4) == 0x0:
		err = decode_sve_sve_perm_unpred_d_sve_int_perm_unpk(ins, d)
	case op0 == 0x2 && op1 == 0x4:
		err = decode_sve_sve_perm_unpred_d_sve_int_perm_insrv(ins, d)
	case op0 == 0x2 && op1 == 0x6:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x5) == 0x5:
		err = errUnallocated
	case op0 == 0x3 && op1 == 0x0:
		err = decode_sve_sve_perm_unpred_d_sve_int_perm_reverse_z(ins, d)
	case op0 == 0x3 && op1 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_unpred_d_sve_int_perm_dup_r(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_dup_r
	d.size = (ins >> 22) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_dup_z_r_
	return
}

func decode_sve_sve_perm_unpred_d_sve_int_perm_insrs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_insrs
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_insr_z_r_
	return
}

func decode_sve_sve_perm_unpred_d_sve_int_perm_unpk(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_unpk
	d.size = (ins >> 22) & 0x3
	d.U = (ins >> 17) & 0x1
	d.H = (ins >> 16) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0 && d.H == 0x0:
		d.encoding = encoding_sunpklo_z_z_
	case d.U == 0x0 && d.H == 0x1:
		d.encoding = encoding_sunpkhi_z_z_
	case d.U == 0x1 && d.H == 0x0:
		d.encoding = encoding_uunpklo_z_z_
	case d.U == 0x1 && d.H == 0x1:
		d.encoding = encoding_uunpkhi_z_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_unpred_d_sve_int_perm_insrv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_insrv
	d.size = (ins >> 22) & 0x3
	d.Vm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_insr_z_v_
	return
}

func decode_sve_sve_perm_unpred_d_sve_int_perm_reverse_z(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_reverse_z
	d.size = (ins >> 22) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_rev_z_z_
	return
}

func decode_sve_sve_perm_predicates(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 22) & 0x3
	op1 := (ins >> 16) & 0x1f
	op2 := (ins >> 9) & 0xf
	op3 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0 && (op1&0x1e) == 0x10 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_perm_predicates_sve_int_perm_punpk(ins, d)
	case op0 == 0x1 && (op1&0x1e) == 0x10 && op2 == 0x0 && op3 == 0x0:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x1e) == 0x10 && op2 == 0x0 && op3 == 0x0:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x1e) == 0x10 && op2 == 0x0 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x10) == 0x0 && (op2&0x1) == 0x0 && op3 == 0x0:
		err = decode_sve_sve_perm_predicates_sve_int_perm_bin_perm_pp(ins, d)
	case (op1&0x10) == 0x0 && (op2&0x1) == 0x1 && op3 == 0x0:
		err = errUnallocated
	case op1 == 0x14 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_perm_predicates_sve_int_perm_reverse_p(ins, d)
	case op1 == 0x15 && op2 == 0x0 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x1a) == 0x10 && op2 == 0x8 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x1a) == 0x10 && (op2&0x7) == 0x4 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x1a) == 0x10 && (op2&0x3) == 0x2 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x1a) == 0x10 && (op2&0x1) == 0x1 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x1a) == 0x12 && op3 == 0x0:
		err = errUnallocated
	case (op1&0x18) == 0x18 && op3 == 0x0:
		err = errUnallocated
	case op3 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_predicates_sve_int_perm_punpk(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_punpk
	d.H = (ins >> 16) & 0x1
	d.Pn = (ins >> 5) & 0xf
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.H == 0x0:
		d.encoding = encoding_punpklo_p_p_
	case d.H == 0x1:
		d.encoding = encoding_punpkhi_p_p_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_predicates_sve_int_perm_bin_perm_pp(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_bin_perm_pp
	d.size = (ins >> 22) & 0x3
	d.Pm = (ins >> 16) & 0xf
	d.opc = (ins >> 11) & 0x3
	d.H = (ins >> 10) & 0x1
	d.Pn = (ins >> 5) & 0xf
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.opc == 0x0 && d.H == 0x0:
		d.encoding = encoding_zip1_p_pp_
	case d.opc == 0x0 && d.H == 0x1:
		d.encoding = encoding_zip2_p_pp_
	case d.opc == 0x1 && d.H == 0x0:
		d.encoding = encoding_uzp1_p_pp_
	case d.opc == 0x1 && d.H == 0x1:
		d.encoding = encoding_uzp2_p_pp_
	case d.opc == 0x2 && d.H == 0x0:
		d.encoding = encoding_trn1_p_pp_
	case d.opc == 0x2 && d.H == 0x1:
		d.encoding = encoding_trn2_p_pp_
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_predicates_sve_int_perm_reverse_p(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_reverse_p
	d.size = (ins >> 22) & 0x3
	d.Pn = (ins >> 5) & 0xf
	d.Pd = (ins >> 0) & 0xf
	d.encoding = encoding_rev_p_p_
	return
}

func decode_sve_sve_int_perm_bin_perm_zz(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_bin_perm_zz
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_zip1_z_zz_
	case d.opc == 0x1:
		d.encoding = encoding_zip2_z_zz_
	case d.opc == 0x2:
		d.encoding = encoding_uzp1_z_zz_
	case d.opc == 0x3:
		d.encoding = encoding_uzp2_z_zz_
	case d.opc == 0x4:
		d.encoding = encoding_trn1_z_zz_
	case d.opc == 0x5:
		d.encoding = encoding_trn2_z_zz_
	case (d.opc & 0x6) == 0x6:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 20) & 0x1
	op1 := (ins >> 17) & 0x7
	op2 := (ins >> 16) & 0x1
	op3 := (ins >> 13) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_cpy_v(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x1 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_compact(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op3 == 0x1:
		err = decode_sve_sve_perm_pred_sve_int_perm_last_r(ins, d)
	case op0 == 0x0 && op1 == 0x1 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_last_v(ins, d)
	case op0 == 0x0 && (op1&0x6) == 0x2 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_rev(ins, d)
	case op0 == 0x0 && (op1&0x6) == 0x2 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x4 && op2 == 0x0 && op3 == 0x1:
		err = decode_sve_sve_perm_pred_sve_int_perm_cpy_r(ins, d)
	case op0 == 0x0 && op1 == 0x4 && op2 == 0x1 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x4 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_clast_zz(ins, d)
	case op0 == 0x0 && op1 == 0x5 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_clast_vz(ins, d)
	case op0 == 0x0 && op1 == 0x6 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_splice(ins, d)
	case op0 == 0x0 && op1 == 0x6 && op2 == 0x1 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_intx_perm_splice(ins, d)
	case op0 == 0x0 && op1 == 0x6 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x7 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_perm_pred_sve_int_perm_revd(ins, d)
	case op0 == 0x0 && op1 == 0x7 && op2 == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x7 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x3) == 0x1 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op3 == 0x0:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op3 == 0x1:
		err = decode_sve_sve_perm_pred_sve_int_perm_clast_rz(ins, d)
	case op0 == 0x1 && op1 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_cpy_v(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_cpy_v
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Vn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_cpy_z_p_v_
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_compact(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_compact
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_compact_z_p_z_
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_last_r(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_last_r
	d.size = (ins >> 22) & 0x3
	d.B = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.B == 0x0:
		d.encoding = encoding_lasta_r_p_z_
	case d.B == 0x1:
		d.encoding = encoding_lastb_r_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_last_v(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_last_v
	d.size = (ins >> 22) & 0x3
	d.B = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Vd = (ins >> 0) & 0x1f

	switch {
	case d.B == 0x0:
		d.encoding = encoding_lasta_v_p_z_
	case d.B == 0x1:
		d.encoding = encoding_lastb_v_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_rev(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_rev
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_revb_z_z_
	case d.opc == 0x1:
		d.encoding = encoding_revh_z_z_
	case d.opc == 0x2:
		d.encoding = encoding_revw_z_z_
	case d.opc == 0x3:
		d.encoding = encoding_rbit_z_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_cpy_r(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_cpy_r
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_cpy_z_p_r_
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_clast_zz(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_clast_zz
	d.size = (ins >> 22) & 0x3
	d.B = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.B == 0x0:
		d.encoding = encoding_clasta_z_p_zz_
	case d.B == 0x1:
		d.encoding = encoding_clastb_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_clast_vz(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_clast_vz
	d.size = (ins >> 22) & 0x3
	d.B = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Vdn = (ins >> 0) & 0x1f

	switch {
	case d.B == 0x0:
		d.encoding = encoding_clasta_v_p_z_
	case d.B == 0x1:
		d.encoding = encoding_clastb_v_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_splice(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_splice
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_splice_z_p_zz_des
	return
}

func decode_sve_sve_perm_pred_sve_intx_perm_splice(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_perm_splice
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_splice_z_p_zz_con
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_revd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_revd
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0:
		d.encoding = encoding_revd_z_p_z_
	case d.size == 0x1:
		err = errUnallocated
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_pred_sve_int_perm_clast_rz(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_clast_rz
	d.size = (ins >> 22) & 0x3
	d.B = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Rdn = (ins >> 0) & 0x1f

	switch {
	case d.B == 0x0:
		d.encoding = encoding_clasta_r_p_z_
	case d.B == 0x1:
		d.encoding = encoding_clastb_r_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_sel_vvv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_sel_vvv
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0xf
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_sel_z_p_zz_
	return
}

func decode_sve_sve_perm_extract(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 22) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_perm_extract_sve_int_perm_extract_i(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_perm_extract_sve_intx_perm_extract_i(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_extract_sve_int_perm_extract_i(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_extract_i
	d.imm8h = (ins >> 16) & 0x1f
	d.imm8l = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_ext_z_zi_des
	return
}

func decode_sve_sve_perm_extract_sve_intx_perm_extract_i(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_perm_extract_i
	d.imm8h = (ins >> 16) & 0x1f
	d.imm8l = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_ext_z_zi_con
	return
}

func decode_sve_sve_perm_inter_long(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 22) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_perm_inter_long_sve_int_perm_bin_long_perm_zz(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_perm_inter_long_sve_int_perm_bin_long_perm_zz(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_perm_bin_long_perm_zz
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 11) & 0x3
	d.H = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.H == 0x0:
		d.encoding = encoding_zip1_z_zz_q
	case d.opc == 0x0 && d.H == 0x1:
		d.encoding = encoding_zip2_z_zz_q
	case d.opc == 0x1 && d.H == 0x0:
		d.encoding = encoding_uzp1_z_zz_q
	case d.opc == 0x1 && d.H == 0x1:
		d.encoding = encoding_uzp2_z_zz_q
	case d.opc == 0x2:
		err = errUnallocated
	case d.opc == 0x3 && d.H == 0x0:
		d.encoding = encoding_trn1_z_zz_q
	case d.opc == 0x3 && d.H == 0x1:
		d.encoding = encoding_trn2_z_zz_q
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpvec(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 14) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_cmpvec_sve_int_cmp_0(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_cmpvec_sve_int_cmp_1(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpvec_sve_int_cmp_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_cmp_0
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 15) & 0x1
	d.o2 = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.ne = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.o2 == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmphs_p_p_zz_
	case d.op == 0x0 && d.o2 == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmphi_p_p_zz_
	case d.op == 0x0 && d.o2 == 0x1 && d.ne == 0x0:
		d.encoding = encoding_cmpeq_p_p_zw_
	case d.op == 0x0 && d.o2 == 0x1 && d.ne == 0x1:
		d.encoding = encoding_cmpne_p_p_zw_
	case d.op == 0x1 && d.o2 == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmpge_p_p_zz_
	case d.op == 0x1 && d.o2 == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmpgt_p_p_zz_
	case d.op == 0x1 && d.o2 == 0x1 && d.ne == 0x0:
		d.encoding = encoding_cmpeq_p_p_zz_
	case d.op == 0x1 && d.o2 == 0x1 && d.ne == 0x1:
		d.encoding = encoding_cmpne_p_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpvec_sve_int_cmp_1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_cmp_1
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 15) & 0x1
	d.lt = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.ne = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.U == 0x0 && d.lt == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmpge_p_p_zw_
	case d.U == 0x0 && d.lt == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmpgt_p_p_zw_
	case d.U == 0x0 && d.lt == 0x1 && d.ne == 0x0:
		d.encoding = encoding_cmplt_p_p_zw_
	case d.U == 0x0 && d.lt == 0x1 && d.ne == 0x1:
		d.encoding = encoding_cmple_p_p_zw_
	case d.U == 0x1 && d.lt == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmphs_p_p_zw_
	case d.U == 0x1 && d.lt == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmphi_p_p_zw_
	case d.U == 0x1 && d.lt == 0x1 && d.ne == 0x0:
		d.encoding = encoding_cmplo_p_p_zw_
	case d.U == 0x1 && d.lt == 0x1 && d.ne == 0x1:
		d.encoding = encoding_cmpls_p_p_zw_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_ucmp_vi(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_ucmp_vi
	d.size = (ins >> 22) & 0x3
	d.imm7 = (ins >> 14) & 0x7f
	d.lt = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.ne = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.lt == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmphs_p_p_zi_
	case d.lt == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmphi_p_p_zi_
	case d.lt == 0x1 && d.ne == 0x0:
		d.encoding = encoding_cmplo_p_p_zi_
	case d.lt == 0x1 && d.ne == 0x1:
		d.encoding = encoding_cmpls_p_p_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_scmp_vi(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_scmp_vi
	d.size = (ins >> 22) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.op = (ins >> 15) & 0x1
	d.o2 = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.ne = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.o2 == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmpge_p_p_zi_
	case d.op == 0x0 && d.o2 == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmpgt_p_p_zi_
	case d.op == 0x0 && d.o2 == 0x1 && d.ne == 0x0:
		d.encoding = encoding_cmplt_p_p_zi_
	case d.op == 0x0 && d.o2 == 0x1 && d.ne == 0x1:
		d.encoding = encoding_cmple_p_p_zi_
	case d.op == 0x1 && d.o2 == 0x0 && d.ne == 0x0:
		d.encoding = encoding_cmpeq_p_p_zi_
	case d.op == 0x1 && d.o2 == 0x0 && d.ne == 0x1:
		d.encoding = encoding_cmpne_p_p_zi_
	case d.op == 0x1 && d.o2 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_log(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pred_log
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pm = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0xf
	d.o2 = (ins >> 9) & 0x1
	d.Pn = (ins >> 5) & 0xf
	d.o3 = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_and_p_p_pp_z
	case d.op == 0x0 && d.S == 0x0 && d.o2 == 0x0 && d.o3 == 0x1:
		d.encoding = encoding_bic_p_p_pp_z
	case d.op == 0x0 && d.S == 0x0 && d.o2 == 0x1 && d.o3 == 0x0:
		d.encoding = encoding_eor_p_p_pp_z
	case d.op == 0x0 && d.S == 0x0 && d.o2 == 0x1 && d.o3 == 0x1:
		d.encoding = encoding_sel_p_p_pp_
	case d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_ands_p_p_pp_z
	case d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x1:
		d.encoding = encoding_bics_p_p_pp_z
	case d.op == 0x0 && d.S == 0x1 && d.o2 == 0x1 && d.o3 == 0x0:
		d.encoding = encoding_eors_p_p_pp_z
	case d.op == 0x0 && d.S == 0x1 && d.o2 == 0x1 && d.o3 == 0x1:
		err = errUnallocated
	case d.op == 0x1 && d.S == 0x0 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_orr_p_p_pp_z
	case d.op == 0x1 && d.S == 0x0 && d.o2 == 0x0 && d.o3 == 0x1:
		d.encoding = encoding_orn_p_p_pp_z
	case d.op == 0x1 && d.S == 0x0 && d.o2 == 0x1 && d.o3 == 0x0:
		d.encoding = encoding_nor_p_p_pp_z
	case d.op == 0x1 && d.S == 0x0 && d.o2 == 0x1 && d.o3 == 0x1:
		d.encoding = encoding_nand_p_p_pp_z
	case d.op == 0x1 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_orrs_p_p_pp_z
	case d.op == 0x1 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x1:
		d.encoding = encoding_orns_p_p_pp_z
	case d.op == 0x1 && d.S == 0x1 && d.o2 == 0x1 && d.o3 == 0x0:
		d.encoding = encoding_nors_p_p_pp_z
	case d.op == 0x1 && d.S == 0x1 && d.o2 == 0x1 && d.o3 == 0x1:
		d.encoding = encoding_nands_p_p_pp_z
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_b(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 9) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_pred_gen_b_sve_int_brkp(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_b_sve_int_brkp(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_brkp
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pm = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0xf
	d.Pn = (ins >> 5) & 0xf
	d.B = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0 && d.B == 0x0:
		d.encoding = encoding_brkpa_p_p_pp_
	case d.op == 0x0 && d.S == 0x0 && d.B == 0x1:
		d.encoding = encoding_brkpb_p_p_pp_
	case d.op == 0x0 && d.S == 0x1 && d.B == 0x0:
		d.encoding = encoding_brkpas_p_p_pp_
	case d.op == 0x0 && d.S == 0x1 && d.B == 0x1:
		d.encoding = encoding_brkpbs_p_p_pp_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_c(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 16) & 0xf
	op2 := (ins >> 9) & 0x1
	op3 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x8 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_pred_gen_c_sve_int_brkn(ins, d)
	case op0 == 0x0 && op1 == 0x8 && op2 == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x7) == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x4) == 0x4:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x2) == 0x2:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x1) == 0x1:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x1 && op1 != 0x0:
		err = errUnallocated
	case op1 == 0x0 && op2 == 0x0:
		err = decode_sve_sve_pred_gen_c_sve_int_break(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_c_sve_int_brkn(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_brkn
	d.S = (ins >> 22) & 0x1
	d.Pg = (ins >> 10) & 0xf
	d.Pn = (ins >> 5) & 0xf
	d.Pdm = (ins >> 0) & 0xf

	switch {
	case d.S == 0x0:
		d.encoding = encoding_brkn_p_p_pp_
	case d.S == 0x1:
		d.encoding = encoding_brkns_p_p_pp_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_c_sve_int_break(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_break
	d.B = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pg = (ins >> 10) & 0xf
	d.Pn = (ins >> 5) & 0xf
	d.M = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.S == 0x1 && d.M == 0x1:
		err = errUnallocated
	case d.B == 0x0 && d.S == 0x0:
		d.encoding = encoding_brka_p_p_p_
	case d.B == 0x0 && d.S == 0x1 && d.M == 0x0:
		d.encoding = encoding_brkas_p_p_p_z
	case d.B == 0x1 && d.S == 0x0:
		d.encoding = encoding_brkb_p_p_p_
	case d.B == 0x1 && d.S == 0x1 && d.M == 0x0:
		d.encoding = encoding_brkbs_p_p_p_z
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 16) & 0xf
	op1 := (ins >> 11) & 0x7
	op2 := (ins >> 9) & 0x3
	op3 := (ins >> 5) & 0xf
	op4 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0 && (op2&0x1) == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_ptest(ins, d)
	case op0 == 0x4 && (op2&0x1) == 0x0 && op4 == 0x0:
		err = errUnallocated
	case (op0&0xb) == 0x2 && (op2&0x1) == 0x0 && op4 == 0x0:
		err = errUnallocated
	case (op0&0x9) == 0x1 && (op2&0x1) == 0x0 && op4 == 0x0:
		err = errUnallocated
	case (op0&0x8) == 0x0 && (op2&0x1) == 0x1 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x8 && op1 == 0x0 && op2 == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_pfirst(ins, d)
	case op0 == 0x8 && op1 == 0x0 && op2 != 0x0 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x8 && op1 == 0x4 && op2 == 0x2 && op3 == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_pfalse(ins, d)
	case op0 == 0x8 && op1 == 0x4 && op2 == 0x2 && op3 != 0x0 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x8 && op1 == 0x6 && op2 == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_rdffr(ins, d)
	case op0 == 0x9 && op1 == 0x0 && (op2&0x2) == 0x0 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x9 && op1 == 0x0 && op2 == 0x2 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_pnext(ins, d)
	case op0 == 0x9 && op1 == 0x0 && op2 == 0x3 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x9 && op1 == 0x4 && op2 == 0x2 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x9 && op1 == 0x6 && op2 == 0x0 && op3 == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_rdffr_2(ins, d)
	case op0 == 0x9 && op1 == 0x6 && op2 == 0x0 && op3 != 0x0 && op4 == 0x0:
		err = errUnallocated
	case (op0&0xe) == 0x8 && op1 == 0x2 && op4 == 0x0:
		err = errUnallocated
	case (op0&0xe) == 0x8 && op1 == 0x4 && (op2&0x2) == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_gen_d_sve_int_ptrue(ins, d)
	case (op0&0xe) == 0x8 && op1 == 0x4 && op2 == 0x3 && op4 == 0x0:
		err = errUnallocated
	case (op0&0xe) == 0x8 && op1 == 0x6 && op2 != 0x0 && op4 == 0x0:
		err = errUnallocated
	case (op0&0xe) == 0x8 && (op1&0x1) == 0x1 && op4 == 0x0:
		err = errUnallocated
	case (op0&0xe) == 0xc && op4 == 0x0:
		err = errUnallocated
	case (op0&0xa) == 0xa && op4 == 0x0:
		err = errUnallocated
	case op4 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d_sve_int_ptest(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_ptest
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pg = (ins >> 10) & 0xf
	d.Pn = (ins >> 5) & 0xf
	d.opc2 = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0:
		err = errUnallocated
	case d.op == 0x0 && d.S == 0x1 && d.opc2 == 0x0:
		d.encoding = encoding_ptest_p_p_
	case d.op == 0x0 && d.S == 0x1 && d.opc2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && d.S == 0x1 && (d.opc2&0xe) == 0x2:
		err = errUnallocated
	case d.op == 0x0 && d.S == 0x1 && (d.opc2&0xc) == 0x4:
		err = errUnallocated
	case d.op == 0x0 && d.S == 0x1 && (d.opc2&0x8) == 0x8:
		err = errUnallocated
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d_sve_int_pfirst(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pfirst
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pg = (ins >> 5) & 0xf
	d.Pdn = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0:
		err = errUnallocated
	case d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_pfirst_p_p_p_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d_sve_int_pfalse(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pfalse
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_pfalse_p_
	case d.op == 0x0 && d.S == 0x1:
		err = errUnallocated
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d_sve_int_rdffr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_rdffr
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pg = (ins >> 5) & 0xf
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_rdffr_p_p_f_
	case d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_rdffrs_p_p_f_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d_sve_int_pnext(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pnext
	d.size = (ins >> 22) & 0x3
	d.Pg = (ins >> 5) & 0xf
	d.Pdn = (ins >> 0) & 0xf
	d.encoding = encoding_pnext_p_p_p_
	return
}

func decode_sve_sve_pred_gen_d_sve_int_rdffr_2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_rdffr_2
	d.op = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_rdffr_p_f_
	case d.op == 0x0 && d.S == 0x1:
		err = errUnallocated
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_gen_d_sve_int_ptrue(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_ptrue
	d.size = (ins >> 22) & 0x3
	d.S = (ins >> 16) & 0x1
	d.pattern = (ins >> 5) & 0x1f
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.S == 0x0:
		d.encoding = encoding_ptrue_p_s_
	case d.S == 0x1:
		d.encoding = encoding_ptrues_p_s_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpgpr(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 12) & 0x3
	op1 := (ins >> 10) & 0x3
	op2 := (ins >> 0) & 0xf

	switch {
	case (op0 & 0x2) == 0x0:
		err = decode_sve_sve_cmpgpr_sve_int_while_rr(ins, d)
	case op0 == 0x2 && op1 == 0x0 && op2 == 0x0:
		err = decode_sve_sve_cmpgpr_sve_int_cterm(ins, d)
	case op0 == 0x2 && op1 == 0x0 && op2 != 0x0:
		err = errUnallocated
	case op0 == 0x3 && op1 == 0x0:
		err = decode_sve_sve_cmpgpr_sve_int_whilenc(ins, d)
	case (op0&0x2) == 0x2 && op1 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpgpr_sve_int_while_rr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_while_rr
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.sf = (ins >> 12) & 0x1
	d.U = (ins >> 11) & 0x1
	d.lt = (ins >> 10) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.eq = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.U == 0x0 && d.lt == 0x0 && d.eq == 0x0:
		d.encoding = encoding_whilege_p_p_rr_
	case d.U == 0x0 && d.lt == 0x0 && d.eq == 0x1:
		d.encoding = encoding_whilegt_p_p_rr_
	case d.U == 0x0 && d.lt == 0x1 && d.eq == 0x0:
		d.encoding = encoding_whilelt_p_p_rr_
	case d.U == 0x0 && d.lt == 0x1 && d.eq == 0x1:
		d.encoding = encoding_whilele_p_p_rr_
	case d.U == 0x1 && d.lt == 0x0 && d.eq == 0x0:
		d.encoding = encoding_whilehs_p_p_rr_
	case d.U == 0x1 && d.lt == 0x0 && d.eq == 0x1:
		d.encoding = encoding_whilehi_p_p_rr_
	case d.U == 0x1 && d.lt == 0x1 && d.eq == 0x0:
		d.encoding = encoding_whilelo_p_p_rr_
	case d.U == 0x1 && d.lt == 0x1 && d.eq == 0x1:
		d.encoding = encoding_whilels_p_p_rr_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpgpr_sve_int_cterm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_cterm
	d.op = (ins >> 23) & 0x1
	d.sz = (ins >> 22) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.ne = (ins >> 4) & 0x1

	switch {
	case d.op == 0x0:
		err = errUnallocated
	case d.op == 0x1 && d.ne == 0x0:
		d.encoding = encoding_ctermeq_rr_
	case d.op == 0x1 && d.ne == 0x1:
		d.encoding = encoding_ctermne_rr_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_cmpgpr_sve_int_whilenc(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_whilenc
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.rw = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.rw == 0x0:
		d.encoding = encoding_whilewr_p_rr_
	case d.rw == 0x1:
		d.encoding = encoding_whilerw_p_rr_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_int_pred_dup(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pred_dup
	d.i1 = (ins >> 23) & 0x1
	d.tszh = (ins >> 22) & 0x1
	d.tszl = (ins >> 18) & 0x7
	d.Rv = (ins >> 16) & 0x3
	d.Pn = (ins >> 10) & 0xf
	d.S = (ins >> 9) & 0x1
	d.Pm = (ins >> 5) & 0xf
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.S == 0x0:
		d.encoding = encoding_psel_p_ppi_
	case d.S == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_unpred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 19) & 0x3
	op1 := (ins >> 16) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_wideimm_unpred_sve_int_arith_imm0(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_wideimm_unpred_sve_int_arith_imm1(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_wideimm_unpred_sve_int_arith_imm2(ins, d)
	case op0 == 0x3 && op1 == 0x0:
		err = decode_sve_sve_wideimm_unpred_sve_int_dup_imm(ins, d)
	case op0 == 0x3 && op1 == 0x1:
		err = decode_sve_sve_wideimm_unpred_sve_int_dup_fpimm(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_unpred_sve_int_arith_imm0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_arith_imm0
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.sh = (ins >> 13) & 0x1
	d.imm8 = (ins >> 5) & 0xff
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_add_z_zi_
	case d.opc == 0x1:
		d.encoding = encoding_sub_z_zi_
	case d.opc == 0x2:
		err = errUnallocated
	case d.opc == 0x3:
		d.encoding = encoding_subr_z_zi_
	case d.opc == 0x4:
		d.encoding = encoding_sqadd_z_zi_
	case d.opc == 0x5:
		d.encoding = encoding_uqadd_z_zi_
	case d.opc == 0x6:
		d.encoding = encoding_sqsub_z_zi_
	case d.opc == 0x7:
		d.encoding = encoding_uqsub_z_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_unpred_sve_int_arith_imm1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_arith_imm1
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.o2 = (ins >> 13) & 0x1
	d.imm8 = (ins >> 5) & 0xff
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case (d.opc&0x4) == 0x0 && d.o2 == 0x1:
		err = errUnallocated
	case d.opc == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_smax_z_zi_
	case d.opc == 0x1 && d.o2 == 0x0:
		d.encoding = encoding_umax_z_zi_
	case d.opc == 0x2 && d.o2 == 0x0:
		d.encoding = encoding_smin_z_zi_
	case d.opc == 0x3 && d.o2 == 0x0:
		d.encoding = encoding_umin_z_zi_
	case (d.opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_unpred_sve_int_arith_imm2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_arith_imm2
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.o2 = (ins >> 13) & 0x1
	d.imm8 = (ins >> 5) & 0xff
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_mul_z_zi_
	case d.opc == 0x0 && d.o2 == 0x1:
		err = errUnallocated
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x6) == 0x2:
		err = errUnallocated
	case (d.opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_unpred_sve_int_dup_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_dup_imm
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 17) & 0x3
	d.sh = (ins >> 13) & 0x1
	d.imm8 = (ins >> 5) & 0xff
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_dup_z_i_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_wideimm_unpred_sve_int_dup_fpimm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_dup_fpimm
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 17) & 0x3
	d.o2 = (ins >> 13) & 0x1
	d.imm8 = (ins >> 5) & 0xff
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_fdup_z_i_
	case d.opc == 0x0 && d.o2 == 0x1:
		err = errUnallocated
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_a(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 9) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_pred_count_a_sve_int_pcount_pred(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_a_sve_int_pcount_pred(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_pcount_pred
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0xf
	d.Pn = (ins >> 5) & 0xf
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_cntp_r_p_p_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x6) == 0x2:
		err = errUnallocated
	case (d.opc & 0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_b(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x1
	op1 := (ins >> 11) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_sve_sve_pred_count_b_sve_int_count_v_sat(ins, d)
	case op0 == 0x0 && op1 == 0x1:
		err = decode_sve_sve_pred_count_b_sve_int_count_r_sat(ins, d)
	case op0 == 0x1 && op1 == 0x0:
		err = decode_sve_sve_pred_count_b_sve_int_count_v(ins, d)
	case op0 == 0x1 && op1 == 0x1:
		err = decode_sve_sve_pred_count_b_sve_int_count_r(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_b_sve_int_count_v_sat(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_count_v_sat
	d.size = (ins >> 22) & 0x3
	d.D = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.opc = (ins >> 9) & 0x3
	d.Pm = (ins >> 5) & 0xf
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	case d.D == 0x0 && d.U == 0x0 && d.opc == 0x0:
		d.encoding = encoding_sqincp_z_p_z_
	case d.D == 0x0 && d.U == 0x1 && d.opc == 0x0:
		d.encoding = encoding_uqincp_z_p_z_
	case d.D == 0x1 && d.U == 0x0 && d.opc == 0x0:
		d.encoding = encoding_sqdecp_z_p_z_
	case d.D == 0x1 && d.U == 0x1 && d.opc == 0x0:
		d.encoding = encoding_uqdecp_z_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_b_sve_int_count_r_sat(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_count_r_sat
	d.size = (ins >> 22) & 0x3
	d.D = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.sf = (ins >> 10) & 0x1
	d.op = (ins >> 9) & 0x1
	d.Pm = (ins >> 5) & 0xf
	d.Rdn = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x1:
		err = errUnallocated
	case d.D == 0x0 && d.U == 0x0 && d.sf == 0x0 && d.op == 0x0:
		d.encoding = encoding_sqincp_r_p_r_sx
	case d.D == 0x0 && d.U == 0x0 && d.sf == 0x1 && d.op == 0x0:
		d.encoding = encoding_sqincp_r_p_r_x
	case d.D == 0x0 && d.U == 0x1 && d.sf == 0x0 && d.op == 0x0:
		d.encoding = encoding_uqincp_r_p_r_uw
	case d.D == 0x0 && d.U == 0x1 && d.sf == 0x1 && d.op == 0x0:
		d.encoding = encoding_uqincp_r_p_r_x
	case d.D == 0x1 && d.U == 0x0 && d.sf == 0x0 && d.op == 0x0:
		d.encoding = encoding_sqdecp_r_p_r_sx
	case d.D == 0x1 && d.U == 0x0 && d.sf == 0x1 && d.op == 0x0:
		d.encoding = encoding_sqdecp_r_p_r_x
	case d.D == 0x1 && d.U == 0x1 && d.sf == 0x0 && d.op == 0x0:
		d.encoding = encoding_uqdecp_r_p_r_uw
	case d.D == 0x1 && d.U == 0x1 && d.sf == 0x1 && d.op == 0x0:
		d.encoding = encoding_uqdecp_r_p_r_x
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_b_sve_int_count_v(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_count_v
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 17) & 0x1
	d.D = (ins >> 16) & 0x1
	d.opc2 = (ins >> 9) & 0x3
	d.Pm = (ins >> 5) & 0xf
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.opc2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x2) == 0x2:
		err = errUnallocated
	case d.op == 0x0 && d.D == 0x0 && d.opc2 == 0x0:
		d.encoding = encoding_incp_z_p_z_
	case d.op == 0x0 && d.D == 0x1 && d.opc2 == 0x0:
		d.encoding = encoding_decp_z_p_z_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_count_b_sve_int_count_r(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_count_r
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 17) & 0x1
	d.D = (ins >> 16) & 0x1
	d.opc2 = (ins >> 9) & 0x3
	d.Pm = (ins >> 5) & 0xf
	d.Rdn = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.opc2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && (d.opc2&0x2) == 0x2:
		err = errUnallocated
	case d.op == 0x0 && d.D == 0x0 && d.opc2 == 0x0:
		d.encoding = encoding_incp_r_p_r_
	case d.op == 0x0 && d.D == 0x1 && d.opc2 == 0x0:
		d.encoding = encoding_decp_r_p_r_
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_wrffr(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x1
	op1 := (ins >> 16) & 0x3
	op2 := (ins >> 9) & 0x7
	op3 := (ins >> 5) & 0xf
	op4 := (ins >> 0) & 0x1f

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_wrffr_sve_int_wrffr(ins, d)
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x0 && op3 == 0x0 && op4 == 0x0:
		err = decode_sve_sve_pred_wrffr_sve_int_setffr(ins, d)
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x0 && (op3&0x8) == 0x8 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x0 && (op3&0x4) == 0x4 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x0 && (op3&0x2) == 0x2 && op4 == 0x0:
		err = errUnallocated
	case op0 == 0x1 && op1 == 0x0 && op2 == 0x0 && (op3&0x1) == 0x1 && op4 == 0x0:
		err = errUnallocated
	case op1 == 0x0 && op2 == 0x0 && op4 != 0x0:
		err = errUnallocated
	case op1 == 0x0 && op2 != 0x0:
		err = errUnallocated
	case op1 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_wrffr_sve_int_wrffr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_wrffr
	d.opc = (ins >> 22) & 0x3
	d.Pn = (ins >> 5) & 0xf

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_wrffr_f_p_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_pred_wrffr_sve_int_setffr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_int_setffr
	d.opc = (ins >> 22) & 0x3

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_setffr_f_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x1f

	switch {
	case (op0 & 0x1e) == 0x0:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_dot(ins, d)
	case (op0 & 0x1e) == 0x2:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_qdmlalbt(ins, d)
	case (op0 & 0x1c) == 0x4:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_cdot(ins, d)
	case (op0 & 0x18) == 0x8:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_cmla(ins, d)
	case (op0 & 0x18) == 0x10:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_mlal_long(ins, d)
	case (op0 & 0x1c) == 0x18:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_qdmlal_long(ins, d)
	case (op0 & 0x1e) == 0x1c:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_qrdmlah(ins, d)
	case op0 == 0x1e:
		err = decode_sve_sve_intx_muladd_unpred_sve_intx_mixed_dot(ins, d)
	case op0 == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_dot(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_dot
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0:
		d.encoding = encoding_sdot_z_zzz_
	case d.U == 0x1:
		d.encoding = encoding_udot_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_qdmlalbt(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qdmlalbt
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0:
		d.encoding = encoding_sqdmlalbt_z_zzz_
	case d.S == 0x1:
		d.encoding = encoding_sqdmlslbt_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_cdot(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cdot
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.rot = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f
	d.encoding = encoding_cdot_z_zzz_
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_cmla(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cmla
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 12) & 0x1
	d.rot = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_cmla_z_zzz_
	case d.op == 0x1:
		d.encoding = encoding_sqrdcmlah_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_mlal_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mlal_long
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 12) & 0x1
	d.U = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smlalb_z_zzz_
	case d.S == 0x0 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smlalt_z_zzz_
	case d.S == 0x0 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umlalb_z_zzz_
	case d.S == 0x0 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umlalt_z_zzz_
	case d.S == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smlslb_z_zzz_
	case d.S == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smlslt_z_zzz_
	case d.S == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umlslb_z_zzz_
	case d.S == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umlslt_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_qdmlal_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qdmlal_long
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqdmlalb_z_zzz_
	case d.S == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqdmlalt_z_zzz_
	case d.S == 0x1 && d.T == 0x0:
		d.encoding = encoding_sqdmlslb_z_zzz_
	case d.S == 0x1 && d.T == 0x1:
		d.encoding = encoding_sqdmlslt_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_qrdmlah(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qrdmlah
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0:
		d.encoding = encoding_sqrdmlah_z_zzz_
	case d.S == 0x1:
		d.encoding = encoding_sqrdmlsh_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_muladd_unpred_sve_intx_mixed_dot(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mixed_dot
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2:
		d.encoding = encoding_usdot_z_zzz_s
	case d.size == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 17) & 0xf
	op1 := (ins >> 13) & 0x1

	switch {
	case op0 == 0x2 && op1 == 0x1:
		err = decode_sve_sve_intx_predicated_sve_intx_accumulate_long_pairs(ins, d)
	case op0 == 0x3 && op1 == 0x1:
		err = errUnallocated
	case (op0&0xe) == 0x6 && op1 == 0x1:
		err = errUnallocated
	case (op0&0xa) == 0x0 && op1 == 0x1:
		err = decode_sve_sve_intx_predicated_sve_intx_pred_arith_unary(ins, d)
	case (op0&0x8) == 0x0 && op1 == 0x0:
		err = decode_sve_sve_intx_predicated_sve_intx_bin_pred_shift_sat_round(ins, d)
	case (op0&0xc) == 0x8 && op1 == 0x0:
		err = decode_sve_sve_intx_predicated_sve_intx_pred_arith_binary(ins, d)
	case (op0&0xc) == 0x8 && op1 == 0x1:
		err = decode_sve_sve_intx_predicated_sve_intx_arith_binary_pairs(ins, d)
	case (op0&0xc) == 0xc && op1 == 0x0:
		err = decode_sve_sve_intx_predicated_sve_intx_pred_arith_binary_sat(ins, d)
	case (op0&0xc) == 0xc && op1 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated_sve_intx_accumulate_long_pairs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_accumulate_long_pairs
	d.size = (ins >> 22) & 0x3
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0:
		d.encoding = encoding_sadalp_z_p_z_
	case d.U == 0x1:
		d.encoding = encoding_uadalp_z_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated_sve_intx_pred_arith_unary(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_pred_arith_unary
	d.size = (ins >> 22) & 0x3
	d.Q = (ins >> 19) & 0x1
	d.opc = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	case d.Q == 0x0 && d.opc == 0x0:
		d.encoding = encoding_urecpe_z_p_z_
	case d.Q == 0x0 && d.opc == 0x1:
		d.encoding = encoding_ursqrte_z_p_z_
	case d.Q == 0x1 && d.opc == 0x0:
		d.encoding = encoding_sqabs_z_p_z_
	case d.Q == 0x1 && d.opc == 0x1:
		d.encoding = encoding_sqneg_z_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated_sve_intx_bin_pred_shift_sat_round(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_bin_pred_shift_sat_round
	d.size = (ins >> 22) & 0x3
	d.Q = (ins >> 19) & 0x1
	d.R = (ins >> 18) & 0x1
	d.N = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.Q == 0x0 && d.N == 0x0:
		err = errUnallocated
	case d.Q == 0x0 && d.R == 0x0 && d.N == 0x1 && d.U == 0x0:
		d.encoding = encoding_srshl_z_p_zz_
	case d.Q == 0x0 && d.R == 0x0 && d.N == 0x1 && d.U == 0x1:
		d.encoding = encoding_urshl_z_p_zz_
	case d.Q == 0x0 && d.R == 0x1 && d.N == 0x1 && d.U == 0x0:
		d.encoding = encoding_srshlr_z_p_zz_
	case d.Q == 0x0 && d.R == 0x1 && d.N == 0x1 && d.U == 0x1:
		d.encoding = encoding_urshlr_z_p_zz_
	case d.Q == 0x1 && d.R == 0x0 && d.N == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqshl_z_p_zz_
	case d.Q == 0x1 && d.R == 0x0 && d.N == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqshl_z_p_zz_
	case d.Q == 0x1 && d.R == 0x0 && d.N == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqrshl_z_p_zz_
	case d.Q == 0x1 && d.R == 0x0 && d.N == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqrshl_z_p_zz_
	case d.Q == 0x1 && d.R == 0x1 && d.N == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqshlr_z_p_zz_
	case d.Q == 0x1 && d.R == 0x1 && d.N == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqshlr_z_p_zz_
	case d.Q == 0x1 && d.R == 0x1 && d.N == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqrshlr_z_p_zz_
	case d.Q == 0x1 && d.R == 0x1 && d.N == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqrshlr_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated_sve_intx_pred_arith_binary(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_pred_arith_binary
	d.size = (ins >> 22) & 0x3
	d.R = (ins >> 18) & 0x1
	d.S = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.R == 0x0 && d.S == 0x0 && d.U == 0x0:
		d.encoding = encoding_shadd_z_p_zz_
	case d.R == 0x0 && d.S == 0x0 && d.U == 0x1:
		d.encoding = encoding_uhadd_z_p_zz_
	case d.R == 0x0 && d.S == 0x1 && d.U == 0x0:
		d.encoding = encoding_shsub_z_p_zz_
	case d.R == 0x0 && d.S == 0x1 && d.U == 0x1:
		d.encoding = encoding_uhsub_z_p_zz_
	case d.R == 0x1 && d.S == 0x0 && d.U == 0x0:
		d.encoding = encoding_srhadd_z_p_zz_
	case d.R == 0x1 && d.S == 0x0 && d.U == 0x1:
		d.encoding = encoding_urhadd_z_p_zz_
	case d.R == 0x1 && d.S == 0x1 && d.U == 0x0:
		d.encoding = encoding_shsubr_z_p_zz_
	case d.R == 0x1 && d.S == 0x1 && d.U == 0x1:
		d.encoding = encoding_uhsubr_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated_sve_intx_arith_binary_pairs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_arith_binary_pairs
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 17) & 0x3
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.U == 0x0:
		err = errUnallocated
	case d.opc == 0x0 && d.U == 0x1:
		d.encoding = encoding_addp_z_p_zz_
	case d.opc == 0x1:
		err = errUnallocated
	case d.opc == 0x2 && d.U == 0x0:
		d.encoding = encoding_smaxp_z_p_zz_
	case d.opc == 0x2 && d.U == 0x1:
		d.encoding = encoding_umaxp_z_p_zz_
	case d.opc == 0x3 && d.U == 0x0:
		d.encoding = encoding_sminp_z_p_zz_
	case d.opc == 0x3 && d.U == 0x1:
		d.encoding = encoding_uminp_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_predicated_sve_intx_pred_arith_binary_sat(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_pred_arith_binary_sat
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 18) & 0x1
	d.S = (ins >> 17) & 0x1
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.S == 0x0 && d.U == 0x0:
		d.encoding = encoding_sqadd_z_p_zz_
	case d.op == 0x0 && d.S == 0x0 && d.U == 0x1:
		d.encoding = encoding_uqadd_z_p_zz_
	case d.op == 0x0 && d.S == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqsub_z_p_zz_
	case d.op == 0x0 && d.S == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqsub_z_p_zz_
	case d.op == 0x1 && d.S == 0x0 && d.U == 0x0:
		d.encoding = encoding_suqadd_z_p_zz_
	case d.op == 0x1 && d.S == 0x0 && d.U == 0x1:
		d.encoding = encoding_usqadd_z_p_zz_
	case d.op == 0x1 && d.S == 0x1 && d.U == 0x0:
		d.encoding = encoding_sqsubr_z_p_zz_
	case d.op == 0x1 && d.S == 0x1 && d.U == 0x1:
		d.encoding = encoding_uqsubr_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_clamp(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_clamp
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0:
		d.encoding = encoding_sclamp_z_zz_
	case d.U == 0x1:
		d.encoding = encoding_uclamp_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x3f

	switch {
	case (op0 & 0x3e) == 0x0:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_dot_by_indexed_elem(ins, d)
	case (op0 & 0x3e) == 0x2:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_mla_by_indexed_elem(ins, d)
	case (op0 & 0x3e) == 0x4:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_qrdmlah_by_indexed_elem(ins, d)
	case (op0 & 0x3e) == 0x6:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_mixed_dot_by_indexed_elem(ins, d)
	case (op0 & 0x38) == 0x8:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_qdmla_long_by_indexed_elem(ins, d)
	case (op0 & 0x3c) == 0x10:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_cdot_by_indexed_elem(ins, d)
	case (op0 & 0x3c) == 0x14:
		err = errUnallocated
	case (op0 & 0x3c) == 0x18:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_cmla_by_indexed_elem(ins, d)
	case (op0 & 0x3c) == 0x1c:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_qrdcmla_by_indexed_elem(ins, d)
	case (op0 & 0x30) == 0x20:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_mla_long_by_indexed_elem(ins, d)
	case (op0 & 0x38) == 0x30:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_mul_long_by_indexed_elem(ins, d)
	case (op0 & 0x3c) == 0x38:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_qdmul_long_by_indexed_elem(ins, d)
	case (op0 & 0x3e) == 0x3c:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_qdmulh_by_indexed_elem(ins, d)
	case op0 == 0x3e:
		err = decode_sve_sve_intx_by_indexed_elem_sve_intx_mul_by_indexed_elem(ins, d)
	case op0 == 0x3f:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_dot_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_dot_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.U = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2 && d.U == 0x0:
		d.encoding = encoding_sdot_z_zzzi_s
	case d.size == 0x2 && d.U == 0x1:
		d.encoding = encoding_udot_z_zzzi_s
	case d.size == 0x3 && d.U == 0x0:
		d.encoding = encoding_sdot_z_zzzi_d
	case d.size == 0x3 && d.U == 0x1:
		d.encoding = encoding_udot_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_mla_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mla_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.S = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.S == 0x0:
		d.encoding = encoding_mla_z_zzzi_h
	case (d.size&0x2) == 0x0 && d.S == 0x1:
		d.encoding = encoding_mls_z_zzzi_h
	case d.size == 0x2 && d.S == 0x0:
		d.encoding = encoding_mla_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1:
		d.encoding = encoding_mls_z_zzzi_s
	case d.size == 0x3 && d.S == 0x0:
		d.encoding = encoding_mla_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1:
		d.encoding = encoding_mls_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_qrdmlah_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qrdmlah_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.S = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.S == 0x0:
		d.encoding = encoding_sqrdmlah_z_zzzi_h
	case (d.size&0x2) == 0x0 && d.S == 0x1:
		d.encoding = encoding_sqrdmlsh_z_zzzi_h
	case d.size == 0x2 && d.S == 0x0:
		d.encoding = encoding_sqrdmlah_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1:
		d.encoding = encoding_sqrdmlsh_z_zzzi_s
	case d.size == 0x3 && d.S == 0x0:
		d.encoding = encoding_sqrdmlah_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1:
		d.encoding = encoding_sqrdmlsh_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_mixed_dot_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mixed_dot_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.U = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2 && d.U == 0x0:
		d.encoding = encoding_usdot_z_zzzi_s
	case d.size == 0x2 && d.U == 0x1:
		d.encoding = encoding_sudot_z_zzzi_s
	case d.size == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_qdmla_long_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qdmla_long_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.S = (ins >> 12) & 0x1
	d.il = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2 && d.S == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqdmlalb_z_zzzi_s
	case d.size == 0x2 && d.S == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqdmlalt_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1 && d.T == 0x0:
		d.encoding = encoding_sqdmlslb_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1 && d.T == 0x1:
		d.encoding = encoding_sqdmlslt_z_zzzi_s
	case d.size == 0x3 && d.S == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqdmlalb_z_zzzi_d
	case d.size == 0x3 && d.S == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqdmlalt_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1 && d.T == 0x0:
		d.encoding = encoding_sqdmlslb_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1 && d.T == 0x1:
		d.encoding = encoding_sqdmlslt_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_cdot_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cdot_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.rot = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2:
		d.encoding = encoding_cdot_z_zzzi_s
	case d.size == 0x3:
		d.encoding = encoding_cdot_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_cmla_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cmla_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.rot = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2:
		d.encoding = encoding_cmla_z_zzzi_h
	case d.size == 0x3:
		d.encoding = encoding_cmla_z_zzzi_s
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_qrdcmla_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qrdcmla_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.rot = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2:
		d.encoding = encoding_sqrdcmlah_z_zzzi_h
	case d.size == 0x3:
		d.encoding = encoding_sqrdcmlah_z_zzzi_s
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_mla_long_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mla_long_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.S = (ins >> 13) & 0x1
	d.U = (ins >> 12) & 0x1
	d.il = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2 && d.S == 0x0 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smlalb_z_zzzi_s
	case d.size == 0x2 && d.S == 0x0 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smlalt_z_zzzi_s
	case d.size == 0x2 && d.S == 0x0 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umlalb_z_zzzi_s
	case d.size == 0x2 && d.S == 0x0 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umlalt_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smlslb_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smlslt_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umlslb_z_zzzi_s
	case d.size == 0x2 && d.S == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umlslt_z_zzzi_s
	case d.size == 0x3 && d.S == 0x0 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smlalb_z_zzzi_d
	case d.size == 0x3 && d.S == 0x0 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smlalt_z_zzzi_d
	case d.size == 0x3 && d.S == 0x0 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umlalb_z_zzzi_d
	case d.size == 0x3 && d.S == 0x0 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umlalt_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smlslb_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smlslt_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umlslb_z_zzzi_d
	case d.size == 0x3 && d.S == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umlslt_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_mul_long_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mul_long_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.U = (ins >> 12) & 0x1
	d.il = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smullb_z_zzi_s
	case d.size == 0x2 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smullt_z_zzi_s
	case d.size == 0x2 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umullb_z_zzi_s
	case d.size == 0x2 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umullt_z_zzi_s
	case d.size == 0x3 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smullb_z_zzi_d
	case d.size == 0x3 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smullt_z_zzi_d
	case d.size == 0x3 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umullb_z_zzi_d
	case d.size == 0x3 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umullt_z_zzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_qdmul_long_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qdmul_long_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.il = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2 && d.T == 0x0:
		d.encoding = encoding_sqdmullb_z_zzi_s
	case d.size == 0x2 && d.T == 0x1:
		d.encoding = encoding_sqdmullt_z_zzi_s
	case d.size == 0x3 && d.T == 0x0:
		d.encoding = encoding_sqdmullb_z_zzi_d
	case d.size == 0x3 && d.T == 0x1:
		d.encoding = encoding_sqdmullt_z_zzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_qdmulh_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_qdmulh_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.R = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.R == 0x0:
		d.encoding = encoding_sqdmulh_z_zzi_h
	case (d.size&0x2) == 0x0 && d.R == 0x1:
		d.encoding = encoding_sqrdmulh_z_zzi_h
	case d.size == 0x2 && d.R == 0x0:
		d.encoding = encoding_sqdmulh_z_zzi_s
	case d.size == 0x2 && d.R == 0x1:
		d.encoding = encoding_sqrdmulh_z_zzi_s
	case d.size == 0x3 && d.R == 0x0:
		d.encoding = encoding_sqdmulh_z_zzi_d
	case d.size == 0x3 && d.R == 0x1:
		d.encoding = encoding_sqrdmulh_z_zzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_by_indexed_elem_sve_intx_mul_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mul_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		d.encoding = encoding_mul_z_zzi_h
	case d.size == 0x2:
		d.encoding = encoding_mul_z_zzi_s
	case d.size == 0x3:
		d.encoding = encoding_mul_z_zzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_cons_widening(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 13) & 0x3

	switch {
	case (op0 & 0x2) == 0x0:
		err = decode_sve_sve_intx_cons_widening_sve_intx_cons_arith_long(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_intx_cons_widening_sve_intx_cons_arith_wide(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_intx_cons_widening_sve_intx_cons_mul_long(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_cons_widening_sve_intx_cons_arith_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cons_arith_long
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 13) & 0x1
	d.S = (ins >> 12) & 0x1
	d.U = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.S == 0x0 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_saddlb_z_zz_
	case d.op == 0x0 && d.S == 0x0 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_saddlt_z_zz_
	case d.op == 0x0 && d.S == 0x0 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_uaddlb_z_zz_
	case d.op == 0x0 && d.S == 0x0 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_uaddlt_z_zz_
	case d.op == 0x0 && d.S == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_ssublb_z_zz_
	case d.op == 0x0 && d.S == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_ssublt_z_zz_
	case d.op == 0x0 && d.S == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_usublb_z_zz_
	case d.op == 0x0 && d.S == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_usublt_z_zz_
	case d.op == 0x1 && d.S == 0x0:
		err = errUnallocated
	case d.op == 0x1 && d.S == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_sabdlb_z_zz_
	case d.op == 0x1 && d.S == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_sabdlt_z_zz_
	case d.op == 0x1 && d.S == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_uabdlb_z_zz_
	case d.op == 0x1 && d.S == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_uabdlt_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_cons_widening_sve_intx_cons_arith_wide(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cons_arith_wide
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 12) & 0x1
	d.U = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_saddwb_z_zz_
	case d.S == 0x0 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_saddwt_z_zz_
	case d.S == 0x0 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_uaddwb_z_zz_
	case d.S == 0x0 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_uaddwt_z_zz_
	case d.S == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_ssubwb_z_zz_
	case d.S == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_ssubwt_z_zz_
	case d.S == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_usubwb_z_zz_
	case d.S == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_usubwt_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_cons_widening_sve_intx_cons_mul_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cons_mul_long
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 12) & 0x1
	d.U = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqdmullb_z_zz_
	case d.op == 0x0 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqdmullt_z_zz_
	case d.op == 0x0 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_pmullb_z_zz_
	case d.op == 0x0 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_pmullt_z_zz_
	case d.op == 0x1 && d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_smullb_z_zz_
	case d.op == 0x1 && d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_smullt_z_zz_
	case d.op == 0x1 && d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_umullb_z_zz_
	case d.op == 0x1 && d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_umullt_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_constructive(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 10) & 0xf

	switch {
	case op0 == 0x0 && (op1&0xc) == 0x8:
		err = decode_sve_sve_intx_constructive_sve_intx_shift_long(ins, d)
	case op0 == 0x1 && (op1&0xc) == 0x8:
		err = errUnallocated
	case (op1 & 0xc) == 0x0:
		err = decode_sve_sve_intx_constructive_sve_intx_clong(ins, d)
	case (op1 & 0xe) == 0x4:
		err = decode_sve_sve_intx_constructive_sve_intx_eorx(ins, d)
	case op1 == 0x6:
		err = decode_sve_sve_intx_constructive_sve_intx_mmla(ins, d)
	case op1 == 0x7:
		err = errUnallocated
	case (op1 & 0xc) == 0xc:
		err = decode_sve_sve_intx_constructive_sve_intx_perm_bit(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_constructive_sve_intx_shift_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_shift_long
	d.tszh = (ins >> 22) & 0x1
	d.tszl = (ins >> 19) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.U = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_sshllb_z_zi_
	case d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_sshllt_z_zi_
	case d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_ushllb_z_zi_
	case d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_ushllt_z_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_constructive_sve_intx_clong(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_clong
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 11) & 0x1
	d.tb = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0 && d.tb == 0x0:
		d.encoding = encoding_saddlbt_z_zz_
	case d.S == 0x0 && d.tb == 0x1:
		err = errUnallocated
	case d.S == 0x1 && d.tb == 0x0:
		d.encoding = encoding_ssublbt_z_zz_
	case d.S == 0x1 && d.tb == 0x1:
		d.encoding = encoding_ssubltb_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_constructive_sve_intx_eorx(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_eorx
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.tb = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.tb == 0x0:
		d.encoding = encoding_eorbt_z_zz_
	case d.tb == 0x1:
		d.encoding = encoding_eortb_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_constructive_sve_intx_mmla(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_mmla
	d.uns = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.uns == 0x0:
		d.encoding = encoding_smmla_z_zzz_
	case d.uns == 0x1:
		err = errUnallocated
	case d.uns == 0x2:
		d.encoding = encoding_usmmla_z_zzz_
	case d.uns == 0x3:
		d.encoding = encoding_ummla_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_constructive_sve_intx_perm_bit(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_perm_bit
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_bext_z_zz_
	case d.opc == 0x1:
		d.encoding = encoding_bdep_z_zz_
	case d.opc == 0x2:
		d.encoding = encoding_bgrp_z_zz_
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 17) & 0xf
	op1 := (ins >> 11) & 0x7

	switch {
	case op0 == 0x0 && op1 == 0x3:
		err = decode_sve_sve_intx_acc_sve_intx_cadd(ins, d)
	case op0 != 0x0 && op1 == 0x3:
		err = errUnallocated
	case (op1 & 0x6) == 0x0:
		err = decode_sve_sve_intx_acc_sve_intx_aba_long(ins, d)
	case op1 == 0x2:
		err = decode_sve_sve_intx_acc_sve_intx_adc_long(ins, d)
	case (op1 & 0x6) == 0x4:
		err = decode_sve_sve_intx_acc_sve_intx_sra(ins, d)
	case op1 == 0x6:
		err = decode_sve_sve_intx_acc_sve_intx_shift_insert(ins, d)
	case op1 == 0x7:
		err = decode_sve_sve_intx_acc_sve_intx_aba(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc_sve_intx_cadd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_cadd
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 16) & 0x1
	d.rot = (ins >> 10) & 0x1
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_cadd_z_zz_
	case d.op == 0x1:
		d.encoding = encoding_sqcadd_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc_sve_intx_aba_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_aba_long
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0 && d.T == 0x0:
		d.encoding = encoding_sabalb_z_zzz_
	case d.U == 0x0 && d.T == 0x1:
		d.encoding = encoding_sabalt_z_zzz_
	case d.U == 0x1 && d.T == 0x0:
		d.encoding = encoding_uabalb_z_zzz_
	case d.U == 0x1 && d.T == 0x1:
		d.encoding = encoding_uabalt_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc_sve_intx_adc_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_adc_long
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.T == 0x0:
		d.encoding = encoding_adclb_z_zzz_
	case (d.size&0x2) == 0x0 && d.T == 0x1:
		d.encoding = encoding_adclt_z_zzz_
	case (d.size&0x2) == 0x2 && d.T == 0x0:
		d.encoding = encoding_sbclb_z_zzz_
	case (d.size&0x2) == 0x2 && d.T == 0x1:
		d.encoding = encoding_sbclt_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc_sve_intx_sra(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_sra
	d.tszh = (ins >> 22) & 0x3
	d.tszl = (ins >> 19) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.R = (ins >> 11) & 0x1
	d.U = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.R == 0x0 && d.U == 0x0:
		d.encoding = encoding_ssra_z_zi_
	case d.R == 0x0 && d.U == 0x1:
		d.encoding = encoding_usra_z_zi_
	case d.R == 0x1 && d.U == 0x0:
		d.encoding = encoding_srsra_z_zi_
	case d.R == 0x1 && d.U == 0x1:
		d.encoding = encoding_ursra_z_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc_sve_intx_shift_insert(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_shift_insert
	d.tszh = (ins >> 22) & 0x3
	d.tszl = (ins >> 19) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.op = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_sri_z_zzi_
	case d.op == 0x1:
		d.encoding = encoding_sli_z_zzi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_acc_sve_intx_aba(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_aba
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0:
		d.encoding = encoding_saba_z_zzz_
	case d.U == 0x1:
		d.encoding = encoding_uaba_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_narrowing(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 16) & 0x7
	op2 := (ins >> 13) & 0x3

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x2:
		err = decode_sve_sve_intx_narrowing_sve_intx_extract_narrow(ins, d)
	case op0 == 0x0 && op1 != 0x0 && op2 == 0x2:
		err = errUnallocated
	case op0 == 0x0 && (op2&0x2) == 0x0:
		err = decode_sve_sve_intx_narrowing_sve_intx_shift_narrow(ins, d)
	case op0 == 0x1 && (op2&0x2) == 0x0:
		err = errUnallocated
	case op0 == 0x1 && op2 == 0x2:
		err = errUnallocated
	case op2 == 0x3:
		err = decode_sve_sve_intx_narrowing_sve_intx_arith_narrow(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_narrowing_sve_intx_extract_narrow(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_extract_narrow
	d.tszh = (ins >> 22) & 0x1
	d.tszl = (ins >> 19) & 0x3
	d.opc = (ins >> 11) & 0x3
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqxtnb_z_zz_
	case d.opc == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqxtnt_z_zz_
	case d.opc == 0x1 && d.T == 0x0:
		d.encoding = encoding_uqxtnb_z_zz_
	case d.opc == 0x1 && d.T == 0x1:
		d.encoding = encoding_uqxtnt_z_zz_
	case d.opc == 0x2 && d.T == 0x0:
		d.encoding = encoding_sqxtunb_z_zz_
	case d.opc == 0x2 && d.T == 0x1:
		d.encoding = encoding_sqxtunt_z_zz_
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_narrowing_sve_intx_shift_narrow(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_shift_narrow
	d.tszh = (ins >> 22) & 0x1
	d.tszl = (ins >> 19) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.op = (ins >> 13) & 0x1
	d.U = (ins >> 12) & 0x1
	d.R = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && d.U == 0x0 && d.R == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqshrunb_z_zi_
	case d.op == 0x0 && d.U == 0x0 && d.R == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqshrunt_z_zi_
	case d.op == 0x0 && d.U == 0x0 && d.R == 0x1 && d.T == 0x0:
		d.encoding = encoding_sqrshrunb_z_zi_
	case d.op == 0x0 && d.U == 0x0 && d.R == 0x1 && d.T == 0x1:
		d.encoding = encoding_sqrshrunt_z_zi_
	case d.op == 0x0 && d.U == 0x1 && d.R == 0x0 && d.T == 0x0:
		d.encoding = encoding_shrnb_z_zi_
	case d.op == 0x0 && d.U == 0x1 && d.R == 0x0 && d.T == 0x1:
		d.encoding = encoding_shrnt_z_zi_
	case d.op == 0x0 && d.U == 0x1 && d.R == 0x1 && d.T == 0x0:
		d.encoding = encoding_rshrnb_z_zi_
	case d.op == 0x0 && d.U == 0x1 && d.R == 0x1 && d.T == 0x1:
		d.encoding = encoding_rshrnt_z_zi_
	case d.op == 0x1 && d.U == 0x0 && d.R == 0x0 && d.T == 0x0:
		d.encoding = encoding_sqshrnb_z_zi_
	case d.op == 0x1 && d.U == 0x0 && d.R == 0x0 && d.T == 0x1:
		d.encoding = encoding_sqshrnt_z_zi_
	case d.op == 0x1 && d.U == 0x0 && d.R == 0x1 && d.T == 0x0:
		d.encoding = encoding_sqrshrnb_z_zi_
	case d.op == 0x1 && d.U == 0x0 && d.R == 0x1 && d.T == 0x1:
		d.encoding = encoding_sqrshrnt_z_zi_
	case d.op == 0x1 && d.U == 0x1 && d.R == 0x0 && d.T == 0x0:
		d.encoding = encoding_uqshrnb_z_zi_
	case d.op == 0x1 && d.U == 0x1 && d.R == 0x0 && d.T == 0x1:
		d.encoding = encoding_uqshrnt_z_zi_
	case d.op == 0x1 && d.U == 0x1 && d.R == 0x1 && d.T == 0x0:
		d.encoding = encoding_uqrshrnb_z_zi_
	case d.op == 0x1 && d.U == 0x1 && d.R == 0x1 && d.T == 0x1:
		d.encoding = encoding_uqrshrnt_z_zi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_narrowing_sve_intx_arith_narrow(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_arith_narrow
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.S = (ins >> 12) & 0x1
	d.R = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.S == 0x0 && d.R == 0x0 && d.T == 0x0:
		d.encoding = encoding_addhnb_z_zz_
	case d.S == 0x0 && d.R == 0x0 && d.T == 0x1:
		d.encoding = encoding_addhnt_z_zz_
	case d.S == 0x0 && d.R == 0x1 && d.T == 0x0:
		d.encoding = encoding_raddhnb_z_zz_
	case d.S == 0x0 && d.R == 0x1 && d.T == 0x1:
		d.encoding = encoding_raddhnt_z_zz_
	case d.S == 0x1 && d.R == 0x0 && d.T == 0x0:
		d.encoding = encoding_subhnb_z_zz_
	case d.S == 0x1 && d.R == 0x0 && d.T == 0x1:
		d.encoding = encoding_subhnt_z_zz_
	case d.S == 0x1 && d.R == 0x1 && d.T == 0x0:
		d.encoding = encoding_rsubhnb_z_zz_
	case d.S == 0x1 && d.R == 0x1 && d.T == 0x1:
		d.encoding = encoding_rsubhnt_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_match(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_match
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.op = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0:
		d.encoding = encoding_match_p_p_zz_
	case d.op == 0x1:
		d.encoding = encoding_nmatch_p_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_histseg(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x7

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_intx_histseg_sve_intx_histseg(ins, d)
	case op0 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_histseg_sve_intx_histseg(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_histseg
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_histseg_z_zz_
	return
}

func decode_sve_sve_intx_histcnt(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_intx_histcnt
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f
	d.encoding = encoding_histcnt_z_p_zz_
	return
}

func decode_sve_sve_intx_crypto(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x7
	op1 := (ins >> 16) & 0x3
	op2 := (ins >> 11) & 0x3
	op3 := (ins >> 5) & 0x1f

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_intx_crypto_sve_crypto_unary(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0 && op3 != 0x0:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x0 && (op2&0x1) == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x1 && (op2&0x2) == 0x0:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x1 && op2 == 0x3:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x2) == 0x2 && op2 == 0x0:
		err = decode_sve_sve_intx_crypto_sve_crypto_binary_dest(ins, d)
	case op0 == 0x0 && (op1&0x2) == 0x2 && (op2&0x1) == 0x1:
		err = errUnallocated
	case op0 != 0x0 && (op2&0x2) == 0x0:
		err = errUnallocated
	case op0 != 0x0 && op2 == 0x3:
		err = errUnallocated
	case op2 == 0x2:
		err = decode_sve_sve_intx_crypto_sve_crypto_binary_const(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_crypto_sve_crypto_unary(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_crypto_unary
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 10) & 0x1
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.op == 0x0:
		d.encoding = encoding_aesmc_z_z_
	case d.size == 0x0 && d.op == 0x1:
		d.encoding = encoding_aesimc_z_z_
	case d.size == 0x1:
		err = errUnallocated
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_crypto_sve_crypto_binary_dest(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_crypto_binary_dest
	d.size = (ins >> 22) & 0x3
	d.op = (ins >> 16) & 0x1
	d.o2 = (ins >> 10) & 0x1
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.op == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_aese_z_zz_
	case d.size == 0x0 && d.op == 0x0 && d.o2 == 0x1:
		d.encoding = encoding_aesd_z_zz_
	case d.size == 0x0 && d.op == 0x1 && d.o2 == 0x0:
		d.encoding = encoding_sm4e_z_zz_
	case d.size == 0x0 && d.op == 0x1 && d.o2 == 0x1:
		err = errUnallocated
	case d.size == 0x1:
		err = errUnallocated
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_intx_crypto_sve_crypto_binary_const(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_crypto_binary_const
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.op == 0x0:
		d.encoding = encoding_sm4ekey_z_zz_
	case d.size == 0x0 && d.op == 0x1:
		d.encoding = encoding_rax1_z_zz_
	case d.size == 0x1:
		err = errUnallocated
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fcmla(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fcmla
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.rot = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f
	d.encoding = encoding_fcmla_z_p_zzz_
	return
}

func decode_sve_sve_fp_fcadd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fcadd
	d.size = (ins >> 22) & 0x3
	d.rot = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_fcadd_z_p_zz_
	return
}

func decode_sve_sve_fp_fcvt2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fcvt2
	d.opc = (ins >> 22) & 0x3
	d.opc2 = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.opc&0x1) == 0x0 && d.opc2 == 0x3:
		err = errUnallocated
	case d.opc == 0x0 && (d.opc2&0x2) == 0x0:
		err = errUnallocated
	case d.opc == 0x0 && d.opc2 == 0x2:
		d.encoding = encoding_fcvtxnt_z_p_z_d2s
	case d.opc == 0x1:
		err = errUnallocated
	case d.opc == 0x2 && d.opc2 == 0x0:
		d.encoding = encoding_fcvtnt_z_p_z_s2h
	case d.opc == 0x2 && d.opc2 == 0x1:
		d.encoding = encoding_fcvtlt_z_p_z_h2s
	case d.opc == 0x2 && d.opc2 == 0x2:
		d.encoding = encoding_bfcvtnt_z_p_z_s2bf
	case d.opc == 0x3 && (d.opc2&0x2) == 0x0:
		err = errUnallocated
	case d.opc == 0x3 && d.opc2 == 0x2:
		d.encoding = encoding_fcvtnt_z_p_z_d2s
	case d.opc == 0x3 && d.opc2 == 0x3:
		d.encoding = encoding_fcvtlt_z_p_z_s2d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_pairwise(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_pairwise
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_faddp_z_p_zz_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x6) == 0x2:
		err = errUnallocated
	case d.opc == 0x4:
		d.encoding = encoding_fmaxnmp_z_p_zz_
	case d.opc == 0x5:
		d.encoding = encoding_fminnmp_z_p_zz_
	case d.opc == 0x6:
		d.encoding = encoding_fmaxp_z_p_zz_
	case d.opc == 0x7:
		d.encoding = encoding_fminp_z_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fma_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.op = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.op == 0x0:
		d.encoding = encoding_fmla_z_zzzi_h
	case (d.size&0x2) == 0x0 && d.op == 0x1:
		d.encoding = encoding_fmls_z_zzzi_h
	case d.size == 0x2 && d.op == 0x0:
		d.encoding = encoding_fmla_z_zzzi_s
	case d.size == 0x2 && d.op == 0x1:
		d.encoding = encoding_fmls_z_zzzi_s
	case d.size == 0x3 && d.op == 0x0:
		d.encoding = encoding_fmla_z_zzzi_d
	case d.size == 0x3 && d.op == 0x1:
		d.encoding = encoding_fmls_z_zzzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fcmla_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fcmla_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.rot = (ins >> 10) & 0x3
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		err = errUnallocated
	case d.size == 0x2:
		d.encoding = encoding_fcmla_z_zzzi_h
	case d.size == 0x3:
		d.encoding = encoding_fcmla_z_zzzi_s
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fmul_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fmul_by_indexed_elem
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.size & 0x2) == 0x0:
		d.encoding = encoding_fmul_z_zzi_h
	case d.size == 0x2:
		d.encoding = encoding_fmul_z_zzi_s
	case d.size == 0x3:
		d.encoding = encoding_fmul_z_zzi_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_w_by_indexed_elem(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 13) & 0x1
	op2 := (ins >> 10) & 0x3

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0:
		err = decode_sve_sve_fp_fma_w_by_indexed_elem_sve_fp_fdot_by_indexed_elem(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op2 != 0x0:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x1:
		err = errUnallocated
	case op0 == 0x1:
		err = decode_sve_sve_fp_fma_w_by_indexed_elem_sve_fp_fma_long_by_indexed_elem(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_w_by_indexed_elem_sve_fp_fdot_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fdot_by_indexed_elem
	d.op = (ins >> 22) & 0x1
	d.i2 = (ins >> 19) & 0x3
	d.Zm = (ins >> 16) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		err = errUnallocated
	case d.op == 0x1:
		d.encoding = encoding_bfdot_z_zzzi_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_w_by_indexed_elem_sve_fp_fma_long_by_indexed_elem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fma_long_by_indexed_elem
	d.o2 = (ins >> 22) & 0x1
	d.i3h = (ins >> 19) & 0x3
	d.Zm = (ins >> 16) & 0x7
	d.op = (ins >> 13) & 0x1
	d.i3l = (ins >> 11) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.o2 == 0x0 && d.op == 0x0 && d.T == 0x0:
		d.encoding = encoding_fmlalb_z_zzzi_s
	case d.o2 == 0x0 && d.op == 0x0 && d.T == 0x1:
		d.encoding = encoding_fmlalt_z_zzzi_s
	case d.o2 == 0x0 && d.op == 0x1 && d.T == 0x0:
		d.encoding = encoding_fmlslb_z_zzzi_s
	case d.o2 == 0x0 && d.op == 0x1 && d.T == 0x1:
		d.encoding = encoding_fmlslt_z_zzzi_s
	case d.o2 == 0x1 && d.op == 0x0 && d.T == 0x0:
		d.encoding = encoding_bfmlalb_z_zzzi_
	case d.o2 == 0x1 && d.op == 0x0 && d.T == 0x1:
		d.encoding = encoding_bfmlalt_z_zzzi_
	case d.o2 == 0x1 && d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_w(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x1
	op1 := (ins >> 13) & 0x1
	op2 := (ins >> 10) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x0:
		err = decode_sve_sve_fp_fma_w_sve_fp_fdot(ins, d)
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x1:
		err = errUnallocated
	case op0 == 0x1:
		err = decode_sve_sve_fp_fma_w_sve_fp_fma_long(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_w_sve_fp_fdot(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fdot
	d.op = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		err = errUnallocated
	case d.op == 0x1:
		d.encoding = encoding_bfdot_z_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_w_sve_fp_fma_long(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fma_long
	d.o2 = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 13) & 0x1
	d.T = (ins >> 10) & 0x1
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.o2 == 0x0 && d.op == 0x0 && d.T == 0x0:
		d.encoding = encoding_fmlalb_z_zzz_
	case d.o2 == 0x0 && d.op == 0x0 && d.T == 0x1:
		d.encoding = encoding_fmlalt_z_zzz_
	case d.o2 == 0x0 && d.op == 0x1 && d.T == 0x0:
		d.encoding = encoding_fmlslb_z_zzz_
	case d.o2 == 0x0 && d.op == 0x1 && d.T == 0x1:
		d.encoding = encoding_fmlslt_z_zzz_
	case d.o2 == 0x1 && d.op == 0x0 && d.T == 0x0:
		d.encoding = encoding_bfmlalb_z_zzz_
	case d.o2 == 0x1 && d.op == 0x0 && d.T == 0x1:
		d.encoding = encoding_bfmlalt_z_zzz_
	case d.o2 == 0x1 && d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fmmla(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fmmla
	d.opc = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		err = errUnallocated
	case d.opc == 0x1:
		d.encoding = encoding_bfmmla_z_zzz_
	case d.opc == 0x2:
		d.encoding = encoding_fmmla_z_zzz_s
	case d.opc == 0x3:
		d.encoding = encoding_fmmla_z_zzz_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_3op_p_pd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_3op_p_pd
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.op = (ins >> 15) & 0x1
	d.o2 = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.o3 = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.op == 0x0 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_fcmge_p_p_zz_
	case d.op == 0x0 && d.o2 == 0x0 && d.o3 == 0x1:
		d.encoding = encoding_fcmgt_p_p_zz_
	case d.op == 0x0 && d.o2 == 0x1 && d.o3 == 0x0:
		d.encoding = encoding_fcmeq_p_p_zz_
	case d.op == 0x0 && d.o2 == 0x1 && d.o3 == 0x1:
		d.encoding = encoding_fcmne_p_p_zz_
	case d.op == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_fcmuo_p_p_zz_
	case d.op == 0x1 && d.o2 == 0x0 && d.o3 == 0x1:
		d.encoding = encoding_facge_p_p_zz_
	case d.op == 0x1 && d.o2 == 0x1 && d.o3 == 0x0:
		err = errUnallocated
	case d.op == 0x1 && d.o2 == 0x1 && d.o3 == 0x1:
		d.encoding = encoding_facgt_p_p_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_3op_u_zd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_3op_u_zd
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fadd_z_zz_
	case d.opc == 0x1:
		d.encoding = encoding_fsub_z_zz_
	case d.opc == 0x2:
		d.encoding = encoding_fmul_z_zz_
	case d.opc == 0x3:
		d.encoding = encoding_ftsmul_z_zz_
	case (d.opc & 0x6) == 0x4:
		err = errUnallocated
	case d.opc == 0x6:
		d.encoding = encoding_frecps_z_zz_
	case d.opc == 0x7:
		d.encoding = encoding_frsqrts_z_zz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_pred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 19) & 0x3
	op1 := (ins >> 10) & 0x7
	op2 := (ins >> 6) & 0xf

	switch {
	case (op0 & 0x2) == 0x0:
		err = decode_sve_sve_fp_pred_sve_fp_2op_p_zds(ins, d)
	case op0 == 0x2 && op1 == 0x0:
		err = decode_sve_sve_fp_pred_sve_fp_ftmad(ins, d)
	case op0 == 0x2 && op1 != 0x0:
		err = errUnallocated
	case op0 == 0x3 && op2 == 0x0:
		err = decode_sve_sve_fp_pred_sve_fp_2op_i_p_zds(ins, d)
	case op0 == 0x3 && op2 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_pred_sve_fp_2op_p_zds(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_zds
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fadd_z_p_zz_
	case d.opc == 0x1:
		d.encoding = encoding_fsub_z_p_zz_
	case d.opc == 0x2:
		d.encoding = encoding_fmul_z_p_zz_
	case d.opc == 0x3:
		d.encoding = encoding_fsubr_z_p_zz_
	case d.opc == 0x4:
		d.encoding = encoding_fmaxnm_z_p_zz_
	case d.opc == 0x5:
		d.encoding = encoding_fminnm_z_p_zz_
	case d.opc == 0x6:
		d.encoding = encoding_fmax_z_p_zz_
	case d.opc == 0x7:
		d.encoding = encoding_fmin_z_p_zz_
	case d.opc == 0x8:
		d.encoding = encoding_fabd_z_p_zz_
	case d.opc == 0x9:
		d.encoding = encoding_fscale_z_p_zz_
	case d.opc == 0xa:
		d.encoding = encoding_fmulx_z_p_zz_
	case d.opc == 0xb:
		err = errUnallocated
	case d.opc == 0xc:
		d.encoding = encoding_fdivr_z_p_zz_
	case d.opc == 0xd:
		d.encoding = encoding_fdiv_z_p_zz_
	case (d.opc & 0xe) == 0xe:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_pred_sve_fp_ftmad(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_ftmad
	d.size = (ins >> 22) & 0x3
	d.imm3 = (ins >> 16) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f
	d.encoding = encoding_ftmad_z_zzi_
	return
}

func decode_sve_sve_fp_pred_sve_fp_2op_i_p_zds(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_i_p_zds
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.i1 = (ins >> 5) & 0x1
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fadd_z_p_zs_
	case d.opc == 0x1:
		d.encoding = encoding_fsub_z_p_zs_
	case d.opc == 0x2:
		d.encoding = encoding_fmul_z_p_zs_
	case d.opc == 0x3:
		d.encoding = encoding_fsubr_z_p_zs_
	case d.opc == 0x4:
		d.encoding = encoding_fmaxnm_z_p_zs_
	case d.opc == 0x5:
		d.encoding = encoding_fminnm_z_p_zs_
	case d.opc == 0x6:
		d.encoding = encoding_fmax_z_p_zs_
	case d.opc == 0x7:
		d.encoding = encoding_fmin_z_p_zs_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x7

	switch {
	case (op0 & 0x6) == 0x0:
		err = decode_sve_sve_fp_unary_sve_fp_2op_p_zd_a(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_fp_unary_sve_fp_2op_p_zd_b_0(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_fp_unary_sve_fp_2op_p_zd_b_1(ins, d)
	case (op0 & 0x6) == 0x4:
		err = decode_sve_sve_fp_unary_sve_fp_2op_p_zd_c(ins, d)
	case (op0 & 0x6) == 0x6:
		err = decode_sve_sve_fp_unary_sve_fp_2op_p_zd_d(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_sve_fp_2op_p_zd_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_zd_a
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_frintn_z_p_z_
	case d.opc == 0x1:
		d.encoding = encoding_frintp_z_p_z_
	case d.opc == 0x2:
		d.encoding = encoding_frintm_z_p_z_
	case d.opc == 0x3:
		d.encoding = encoding_frintz_z_p_z_
	case d.opc == 0x4:
		d.encoding = encoding_frinta_z_p_z_
	case d.opc == 0x5:
		err = errUnallocated
	case d.opc == 0x6:
		d.encoding = encoding_frintx_z_p_z_
	case d.opc == 0x7:
		d.encoding = encoding_frinti_z_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_sve_fp_2op_p_zd_b_0(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_zd_b_0
	d.opc = (ins >> 22) & 0x3
	d.opc2 = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.opc&0x1) == 0x0 && d.opc2 == 0x3:
		err = errUnallocated
	case d.opc == 0x0 && (d.opc2&0x2) == 0x0:
		err = errUnallocated
	case d.opc == 0x0 && d.opc2 == 0x2:
		d.encoding = encoding_fcvtx_z_p_z_d2s
	case d.opc == 0x1:
		err = errUnallocated
	case d.opc == 0x2 && d.opc2 == 0x0:
		d.encoding = encoding_fcvt_z_p_z_s2h
	case d.opc == 0x2 && d.opc2 == 0x1:
		d.encoding = encoding_fcvt_z_p_z_h2s
	case d.opc == 0x2 && d.opc2 == 0x2:
		d.encoding = encoding_bfcvt_z_p_z_s2bf
	case d.opc == 0x3 && d.opc2 == 0x0:
		d.encoding = encoding_fcvt_z_p_z_d2h
	case d.opc == 0x3 && d.opc2 == 0x1:
		d.encoding = encoding_fcvt_z_p_z_h2d
	case d.opc == 0x3 && d.opc2 == 0x2:
		d.encoding = encoding_fcvt_z_p_z_d2s
	case d.opc == 0x3 && d.opc2 == 0x3:
		d.encoding = encoding_fcvt_z_p_z_s2d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_sve_fp_2op_p_zd_b_1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_zd_b_1
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_frecpx_z_p_z_
	case d.opc == 0x1:
		d.encoding = encoding_fsqrt_z_p_z_
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_sve_fp_2op_p_zd_c(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_zd_c
	d.opc = (ins >> 22) & 0x3
	d.opc2 = (ins >> 17) & 0x3
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		err = errUnallocated
	case d.opc == 0x1 && d.opc2 == 0x0:
		err = errUnallocated
	case d.opc == 0x1 && d.opc2 == 0x1 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_h2fp16
	case d.opc == 0x1 && d.opc2 == 0x1 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_h2fp16
	case d.opc == 0x1 && d.opc2 == 0x2 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_w2fp16
	case d.opc == 0x1 && d.opc2 == 0x2 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_w2fp16
	case d.opc == 0x1 && d.opc2 == 0x3 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_x2fp16
	case d.opc == 0x1 && d.opc2 == 0x3 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_x2fp16
	case d.opc == 0x2 && (d.opc2&0x2) == 0x0:
		err = errUnallocated
	case d.opc == 0x2 && d.opc2 == 0x2 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_w2s
	case d.opc == 0x2 && d.opc2 == 0x2 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_w2s
	case d.opc == 0x2 && d.opc2 == 0x3:
		err = errUnallocated
	case d.opc == 0x3 && d.opc2 == 0x0 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_w2d
	case d.opc == 0x3 && d.opc2 == 0x0 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_w2d
	case d.opc == 0x3 && d.opc2 == 0x1:
		err = errUnallocated
	case d.opc == 0x3 && d.opc2 == 0x2 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_x2s
	case d.opc == 0x3 && d.opc2 == 0x2 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_x2s
	case d.opc == 0x3 && d.opc2 == 0x3 && d.U == 0x0:
		d.encoding = encoding_scvtf_z_p_z_x2d
	case d.opc == 0x3 && d.opc2 == 0x3 && d.U == 0x1:
		d.encoding = encoding_ucvtf_z_p_z_x2d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_sve_fp_2op_p_zd_d(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_zd_d
	d.opc = (ins >> 22) & 0x3
	d.opc2 = (ins >> 17) & 0x3
	d.U = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.U == 0x0:
		d.encoding = encoding_flogb_z_p_z_
	case d.opc == 0x0 && d.U == 0x1:
		err = errUnallocated
	case d.opc == 0x1 && d.opc2 == 0x0:
		err = errUnallocated
	case d.opc == 0x1 && d.opc2 == 0x1 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_fp162h
	case d.opc == 0x1 && d.opc2 == 0x1 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_fp162h
	case d.opc == 0x1 && d.opc2 == 0x2 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_fp162w
	case d.opc == 0x1 && d.opc2 == 0x2 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_fp162w
	case d.opc == 0x1 && d.opc2 == 0x3 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_fp162x
	case d.opc == 0x1 && d.opc2 == 0x3 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_fp162x
	case d.opc == 0x2 && (d.opc2&0x2) == 0x0:
		err = errUnallocated
	case d.opc == 0x2 && d.opc2 == 0x2 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_s2w
	case d.opc == 0x2 && d.opc2 == 0x2 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_s2w
	case d.opc == 0x2 && d.opc2 == 0x3:
		err = errUnallocated
	case d.opc == 0x3 && d.opc2 == 0x0 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_d2w
	case d.opc == 0x3 && d.opc2 == 0x0 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_d2w
	case d.opc == 0x3 && d.opc2 == 0x1:
		err = errUnallocated
	case d.opc == 0x3 && d.opc2 == 0x2 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_s2x
	case d.opc == 0x3 && d.opc2 == 0x2 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_s2x
	case d.opc == 0x3 && d.opc2 == 0x3 && d.U == 0x0:
		d.encoding = encoding_fcvtzs_z_p_z_d2x
	case d.opc == 0x3 && d.opc2 == 0x3 && d.U == 0x1:
		d.encoding = encoding_fcvtzu_z_p_z_d2x
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fast_red(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_fast_red
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Vd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_faddv_v_p_z_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x6) == 0x2:
		err = errUnallocated
	case d.opc == 0x4:
		d.encoding = encoding_fmaxnmv_v_p_z_
	case d.opc == 0x5:
		d.encoding = encoding_fminnmv_v_p_z_
	case d.opc == 0x6:
		d.encoding = encoding_fmaxv_v_p_z_
	case d.opc == 0x7:
		d.encoding = encoding_fminv_v_p_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_unpred(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 10) & 0x3

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_fp_unary_unpred_sve_fp_2op_u_zd(ins, d)
	case op0 != 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_unary_unpred_sve_fp_2op_u_zd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_u_zd
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zd = (ins >> 0) & 0x1f

	switch {
	case (d.opc & 0x4) == 0x0:
		err = errUnallocated
	case (d.opc & 0x6) == 0x4:
		err = errUnallocated
	case d.opc == 0x6:
		d.encoding = encoding_frecpe_z_z_
	case d.opc == 0x7:
		d.encoding = encoding_frsqrte_z_z_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_cmpzero(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_fp_cmpzero_sve_fp_2op_p_pd(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_cmpzero_sve_fp_2op_p_pd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_pd
	d.size = (ins >> 22) & 0x3
	d.eq = (ins >> 17) & 0x1
	d.lt = (ins >> 16) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.ne = (ins >> 4) & 0x1
	d.Pd = (ins >> 0) & 0xf

	switch {
	case d.eq == 0x0 && d.lt == 0x0 && d.ne == 0x0:
		d.encoding = encoding_fcmge_p_p_z0_
	case d.eq == 0x0 && d.lt == 0x0 && d.ne == 0x1:
		d.encoding = encoding_fcmgt_p_p_z0_
	case d.eq == 0x0 && d.lt == 0x1 && d.ne == 0x0:
		d.encoding = encoding_fcmlt_p_p_z0_
	case d.eq == 0x0 && d.lt == 0x1 && d.ne == 0x1:
		d.encoding = encoding_fcmle_p_p_z0_
	case d.eq == 0x1 && d.ne == 0x1:
		err = errUnallocated
	case d.eq == 0x1 && d.lt == 0x0 && d.ne == 0x0:
		d.encoding = encoding_fcmeq_p_p_z0_
	case d.eq == 0x1 && d.lt == 0x1 && d.ne == 0x0:
		d.encoding = encoding_fcmne_p_p_z0_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_slowreduce(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 18) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_fp_slowreduce_sve_fp_2op_p_vd(ins, d)
	case op0 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_slowreduce_sve_fp_2op_p_vd(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_2op_p_vd
	d.size = (ins >> 22) & 0x3
	d.opc = (ins >> 16) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Vdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fadda_v_p_z_
	case d.opc == 0x1:
		err = errUnallocated
	case (d.opc & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 15) & 0x1

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_fp_fma_sve_fp_3op_p_zds_a(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_fp_fma_sve_fp_3op_p_zds_b(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_sve_fp_3op_p_zds_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_3op_p_zds_a
	d.size = (ins >> 22) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.opc = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zda = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fmla_z_p_zzz_
	case d.opc == 0x1:
		d.encoding = encoding_fmls_z_p_zzz_
	case d.opc == 0x2:
		d.encoding = encoding_fnmla_z_p_zzz_
	case d.opc == 0x3:
		d.encoding = encoding_fnmls_z_p_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_fp_fma_sve_fp_3op_p_zds_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_fp_3op_p_zds_b
	d.size = (ins >> 22) & 0x3
	d.Za = (ins >> 16) & 0x1f
	d.opc = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Zm = (ins >> 5) & 0x1f
	d.Zdn = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0:
		d.encoding = encoding_fmad_z_p_zzz_
	case d.opc == 0x1:
		d.encoding = encoding_fmsb_z_p_zzz_
	case d.opc == 0x2:
		d.encoding = encoding_fnmad_z_p_zzz_
	case d.opc == 0x3:
		d.encoding = encoding_fnmsb_z_p_zzz_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x3
	op1 := (ins >> 21) & 0x3
	op2 := (ins >> 13) & 0x7
	op3 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0 && (op1&0x1) == 0x1 && (op2&0x4) == 0x0 && op3 == 0x0:
		err = decode_sve_sve_mem32_sve_mem_32b_prfm_sv(ins, d)
	case op0 == 0x0 && (op1&0x1) == 0x1 && (op2&0x4) == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x1 && (op1&0x1) == 0x1 && (op2&0x4) == 0x0:
		err = decode_sve_sve_mem32_sve_mem_32b_gld_sv_a(ins, d)
	case op0 == 0x2 && (op1&0x1) == 0x1 && (op2&0x4) == 0x0:
		err = decode_sve_sve_mem32_sve_mem_32b_gld_sv_b(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x0 && op3 == 0x0:
		err = decode_sve_sve_mem32_sve_mem_32b_pfill(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x0 && op2 == 0x2:
		err = decode_sve_sve_mem32_sve_mem_32b_fill(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x0 && (op2&0x5) == 0x1:
		err = errUnallocated
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x4) == 0x0 && op3 == 0x0:
		err = decode_sve_sve_mem32_sve_mem_prfm_si(ins, d)
	case op0 == 0x3 && (op1&0x2) == 0x2 && (op2&0x4) == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 != 0x3 && (op1&0x1) == 0x0 && (op2&0x4) == 0x0:
		err = decode_sve_sve_mem32_sve_mem_32b_gld_vs(ins, d)
	case op1 == 0x0 && (op2&0x6) == 0x4:
		err = decode_sve_sve_mem32_sve_mem_32b_gldnt_vs(ins, d)
	case op1 == 0x0 && op2 == 0x6 && op3 == 0x0:
		err = decode_sve_sve_mem32_sve_mem_prfm_ss(ins, d)
	case op1 == 0x0 && op2 == 0x7 && op3 == 0x0:
		err = decode_sve_sve_mem32_sve_mem_32b_prfm_vi(ins, d)
	case op1 == 0x0 && (op2&0x6) == 0x6 && op3 == 0x1:
		err = errUnallocated
	case op1 == 0x1 && (op2&0x4) == 0x4:
		err = decode_sve_sve_mem32_sve_mem_32b_gld_vi(ins, d)
	case (op1&0x2) == 0x2 && (op2&0x4) == 0x4:
		err = decode_sve_sve_mem32_sve_mem_ld_dup(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_prfm_sv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_prfm_sv
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.msz = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_bz_s_x32_scaled
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_bz_s_x32_scaled
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_bz_s_x32_scaled
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_bz_s_x32_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_gld_sv_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_gld_sv_a
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_bz_s_x32_scaled
	case d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_bz_s_x32_scaled
	case d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_bz_s_x32_scaled
	case d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_bz_s_x32_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_gld_sv_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_gld_sv_b
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0:
		err = errUnallocated
	case d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_bz_s_x32_scaled
	case d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_bz_s_x32_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_pfill(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_pfill
	d.imm9h = (ins >> 16) & 0x3f
	d.imm9l = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Pt = (ins >> 0) & 0xf
	d.encoding = encoding_ldr_p_bi_
	return
}

func decode_sve_sve_mem32_sve_mem_32b_fill(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_fill
	d.imm9h = (ins >> 16) & 0x3f
	d.imm9l = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f
	d.encoding = encoding_ldr_z_bi_
	return
}

func decode_sve_sve_mem32_sve_mem_prfm_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_prfm_si
	d.imm6 = (ins >> 16) & 0x3f
	d.msz = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_bi_s
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_bi_s
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_bi_s
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_bi_s
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_gld_vs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_gld_vs
	d.opc = (ins >> 23) & 0x3
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sb_z_p_bz_s_x32_unscaled
	case d.opc == 0x0 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sb_z_p_bz_s_x32_unscaled
	case d.opc == 0x0 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1b_z_p_bz_s_x32_unscaled
	case d.opc == 0x0 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1b_z_p_bz_s_x32_unscaled
	case d.opc == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_bz_s_x32_unscaled
	case d.opc == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_bz_s_x32_unscaled
	case d.opc == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_bz_s_x32_unscaled
	case d.opc == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_bz_s_x32_unscaled
	case d.opc == 0x2 && d.U == 0x0:
		err = errUnallocated
	case d.opc == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_bz_s_x32_unscaled
	case d.opc == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_bz_s_x32_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_gldnt_vs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_gldnt_vs
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.U = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.U == 0x0:
		d.encoding = encoding_ldnt1sb_z_p_ar_s_x32_unscaled
	case d.msz == 0x0 && d.U == 0x1:
		d.encoding = encoding_ldnt1b_z_p_ar_s_x32_unscaled
	case d.msz == 0x1 && d.U == 0x0:
		d.encoding = encoding_ldnt1sh_z_p_ar_s_x32_unscaled
	case d.msz == 0x1 && d.U == 0x1:
		d.encoding = encoding_ldnt1h_z_p_ar_s_x32_unscaled
	case d.msz == 0x2 && d.U == 0x0:
		err = errUnallocated
	case d.msz == 0x2 && d.U == 0x1:
		d.encoding = encoding_ldnt1w_z_p_ar_s_x32_unscaled
	case d.msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_prfm_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_prfm_ss
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_br_s
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_br_s
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_br_s
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_br_s
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_prfm_vi(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_prfm_vi
	d.msz = (ins >> 23) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_ai_s
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_ai_s
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_ai_s
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_ai_s
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_32b_gld_vi(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_32b_gld_vi
	d.msz = (ins >> 23) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sb_z_p_ai_s
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sb_z_p_ai_s
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1b_z_p_ai_s
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1b_z_p_ai_s
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_ai_s
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_ai_s
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_ai_s
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_ai_s
	case d.msz == 0x2 && d.U == 0x0:
		err = errUnallocated
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_ai_s
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_ai_s
	case d.msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem32_sve_mem_ld_dup(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_ld_dup
	d.dtypeh = (ins >> 23) & 0x3
	d.imm6 = (ins >> 16) & 0x3f
	d.dtypel = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.dtypeh == 0x0 && d.dtypel == 0x0:
		d.encoding = encoding_ld1rb_z_p_bi_u8
	case d.dtypeh == 0x0 && d.dtypel == 0x1:
		d.encoding = encoding_ld1rb_z_p_bi_u16
	case d.dtypeh == 0x0 && d.dtypel == 0x2:
		d.encoding = encoding_ld1rb_z_p_bi_u32
	case d.dtypeh == 0x0 && d.dtypel == 0x3:
		d.encoding = encoding_ld1rb_z_p_bi_u64
	case d.dtypeh == 0x1 && d.dtypel == 0x0:
		d.encoding = encoding_ld1rsw_z_p_bi_s64
	case d.dtypeh == 0x1 && d.dtypel == 0x1:
		d.encoding = encoding_ld1rh_z_p_bi_u16
	case d.dtypeh == 0x1 && d.dtypel == 0x2:
		d.encoding = encoding_ld1rh_z_p_bi_u32
	case d.dtypeh == 0x1 && d.dtypel == 0x3:
		d.encoding = encoding_ld1rh_z_p_bi_u64
	case d.dtypeh == 0x2 && d.dtypel == 0x0:
		d.encoding = encoding_ld1rsh_z_p_bi_s64
	case d.dtypeh == 0x2 && d.dtypel == 0x1:
		d.encoding = encoding_ld1rsh_z_p_bi_s32
	case d.dtypeh == 0x2 && d.dtypel == 0x2:
		d.encoding = encoding_ld1rw_z_p_bi_u32
	case d.dtypeh == 0x2 && d.dtypel == 0x3:
		d.encoding = encoding_ld1rw_z_p_bi_u64
	case d.dtypeh == 0x3 && d.dtypel == 0x0:
		d.encoding = encoding_ld1rsb_z_p_bi_s64
	case d.dtypeh == 0x3 && d.dtypel == 0x1:
		d.encoding = encoding_ld1rsb_z_p_bi_s32
	case d.dtypeh == 0x3 && d.dtypel == 0x2:
		d.encoding = encoding_ld1rsb_z_p_bi_s16
	case d.dtypeh == 0x3 && d.dtypel == 0x3:
		d.encoding = encoding_ld1rd_z_p_bi_u64
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 21) & 0x3
	op1 := (ins >> 20) & 0x1
	op2 := (ins >> 13) & 0x7

	switch {
	case op0 == 0x0 && op1 == 0x0 && op2 == 0x7:
		err = decode_sve_sve_memcld_sve_mem_cldnt_si(ins, d)
	case op0 == 0x0 && op2 == 0x6:
		err = decode_sve_sve_memcld_sve_mem_cldnt_ss(ins, d)
	case op0 != 0x0 && op1 == 0x0 && op2 == 0x7:
		err = decode_sve_sve_memcld_sve_mem_eld_si(ins, d)
	case op0 != 0x0 && op2 == 0x6:
		err = decode_sve_sve_memcld_sve_mem_eld_ss(ins, d)
	case op1 == 0x0 && op2 == 0x1:
		err = decode_sve_sve_memcld_sve_mem_ldqr_si(ins, d)
	case op1 == 0x0 && op2 == 0x5:
		err = decode_sve_sve_memcld_sve_mem_cld_si(ins, d)
	case op1 == 0x1 && op2 == 0x1:
		err = errUnallocated
	case op1 == 0x1 && op2 == 0x5:
		err = decode_sve_sve_memcld_sve_mem_cldnf_si(ins, d)
	case op1 == 0x1 && op2 == 0x7:
		err = errUnallocated
	case op2 == 0x0:
		err = decode_sve_sve_memcld_sve_mem_ldqr_ss(ins, d)
	case op2 == 0x2:
		err = decode_sve_sve_memcld_sve_mem_cld_ss(ins, d)
	case op2 == 0x3:
		err = decode_sve_sve_memcld_sve_mem_cldff_ss(ins, d)
	case op2 == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_cldnt_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cldnt_si
	d.msz = (ins >> 23) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_ldnt1b_z_p_bi_contiguous
	case d.msz == 0x1:
		d.encoding = encoding_ldnt1h_z_p_bi_contiguous
	case d.msz == 0x2:
		d.encoding = encoding_ldnt1w_z_p_bi_contiguous
	case d.msz == 0x3:
		d.encoding = encoding_ldnt1d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_cldnt_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cldnt_ss
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_ldnt1b_z_p_br_contiguous
	case d.msz == 0x1:
		d.encoding = encoding_ldnt1h_z_p_br_contiguous
	case d.msz == 0x2:
		d.encoding = encoding_ldnt1w_z_p_br_contiguous
	case d.msz == 0x3:
		d.encoding = encoding_ldnt1d_z_p_br_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_eld_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_eld_si
	d.msz = (ins >> 23) & 0x3
	d.opc = (ins >> 21) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.opc == 0x1:
		d.encoding = encoding_ld2b_z_p_bi_contiguous
	case d.msz == 0x0 && d.opc == 0x2:
		d.encoding = encoding_ld3b_z_p_bi_contiguous
	case d.msz == 0x0 && d.opc == 0x3:
		d.encoding = encoding_ld4b_z_p_bi_contiguous
	case d.msz == 0x1 && d.opc == 0x1:
		d.encoding = encoding_ld2h_z_p_bi_contiguous
	case d.msz == 0x1 && d.opc == 0x2:
		d.encoding = encoding_ld3h_z_p_bi_contiguous
	case d.msz == 0x1 && d.opc == 0x3:
		d.encoding = encoding_ld4h_z_p_bi_contiguous
	case d.msz == 0x2 && d.opc == 0x1:
		d.encoding = encoding_ld2w_z_p_bi_contiguous
	case d.msz == 0x2 && d.opc == 0x2:
		d.encoding = encoding_ld3w_z_p_bi_contiguous
	case d.msz == 0x2 && d.opc == 0x3:
		d.encoding = encoding_ld4w_z_p_bi_contiguous
	case d.msz == 0x3 && d.opc == 0x1:
		d.encoding = encoding_ld2d_z_p_bi_contiguous
	case d.msz == 0x3 && d.opc == 0x2:
		d.encoding = encoding_ld3d_z_p_bi_contiguous
	case d.msz == 0x3 && d.opc == 0x3:
		d.encoding = encoding_ld4d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_eld_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_eld_ss
	d.msz = (ins >> 23) & 0x3
	d.opc = (ins >> 21) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.opc == 0x1:
		d.encoding = encoding_ld2b_z_p_br_contiguous
	case d.msz == 0x0 && d.opc == 0x2:
		d.encoding = encoding_ld3b_z_p_br_contiguous
	case d.msz == 0x0 && d.opc == 0x3:
		d.encoding = encoding_ld4b_z_p_br_contiguous
	case d.msz == 0x1 && d.opc == 0x1:
		d.encoding = encoding_ld2h_z_p_br_contiguous
	case d.msz == 0x1 && d.opc == 0x2:
		d.encoding = encoding_ld3h_z_p_br_contiguous
	case d.msz == 0x1 && d.opc == 0x3:
		d.encoding = encoding_ld4h_z_p_br_contiguous
	case d.msz == 0x2 && d.opc == 0x1:
		d.encoding = encoding_ld2w_z_p_br_contiguous
	case d.msz == 0x2 && d.opc == 0x2:
		d.encoding = encoding_ld3w_z_p_br_contiguous
	case d.msz == 0x2 && d.opc == 0x3:
		d.encoding = encoding_ld4w_z_p_br_contiguous
	case d.msz == 0x3 && d.opc == 0x1:
		d.encoding = encoding_ld2d_z_p_br_contiguous
	case d.msz == 0x3 && d.opc == 0x2:
		d.encoding = encoding_ld3d_z_p_br_contiguous
	case d.msz == 0x3 && d.opc == 0x3:
		d.encoding = encoding_ld4d_z_p_br_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_ldqr_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_ldqr_si
	d.msz = (ins >> 23) & 0x3
	d.ssz = (ins >> 21) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case (d.ssz & 0x2) == 0x2:
		err = errUnallocated
	case d.msz == 0x0 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqb_z_p_bi_u8
	case d.msz == 0x0 && d.ssz == 0x1:
		d.encoding = encoding_ld1rob_z_p_bi_u8
	case d.msz == 0x1 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqh_z_p_bi_u16
	case d.msz == 0x1 && d.ssz == 0x1:
		d.encoding = encoding_ld1roh_z_p_bi_u16
	case d.msz == 0x2 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqw_z_p_bi_u32
	case d.msz == 0x2 && d.ssz == 0x1:
		d.encoding = encoding_ld1row_z_p_bi_u32
	case d.msz == 0x3 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqd_z_p_bi_u64
	case d.msz == 0x3 && d.ssz == 0x1:
		d.encoding = encoding_ld1rod_z_p_bi_u64
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_cld_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cld_si
	d.dtype = (ins >> 21) & 0xf
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.dtype == 0x0:
		d.encoding = encoding_ld1b_z_p_bi_u8
	case d.dtype == 0x1:
		d.encoding = encoding_ld1b_z_p_bi_u16
	case d.dtype == 0x2:
		d.encoding = encoding_ld1b_z_p_bi_u32
	case d.dtype == 0x3:
		d.encoding = encoding_ld1b_z_p_bi_u64
	case d.dtype == 0x4:
		d.encoding = encoding_ld1sw_z_p_bi_s64
	case d.dtype == 0x5:
		d.encoding = encoding_ld1h_z_p_bi_u16
	case d.dtype == 0x6:
		d.encoding = encoding_ld1h_z_p_bi_u32
	case d.dtype == 0x7:
		d.encoding = encoding_ld1h_z_p_bi_u64
	case d.dtype == 0x8:
		d.encoding = encoding_ld1sh_z_p_bi_s64
	case d.dtype == 0x9:
		d.encoding = encoding_ld1sh_z_p_bi_s32
	case d.dtype == 0xa:
		d.encoding = encoding_ld1w_z_p_bi_u32
	case d.dtype == 0xb:
		d.encoding = encoding_ld1w_z_p_bi_u64
	case d.dtype == 0xc:
		d.encoding = encoding_ld1sb_z_p_bi_s64
	case d.dtype == 0xd:
		d.encoding = encoding_ld1sb_z_p_bi_s32
	case d.dtype == 0xe:
		d.encoding = encoding_ld1sb_z_p_bi_s16
	case d.dtype == 0xf:
		d.encoding = encoding_ld1d_z_p_bi_u64
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_cldnf_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cldnf_si
	d.dtype = (ins >> 21) & 0xf
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.dtype == 0x0:
		d.encoding = encoding_ldnf1b_z_p_bi_u8
	case d.dtype == 0x1:
		d.encoding = encoding_ldnf1b_z_p_bi_u16
	case d.dtype == 0x2:
		d.encoding = encoding_ldnf1b_z_p_bi_u32
	case d.dtype == 0x3:
		d.encoding = encoding_ldnf1b_z_p_bi_u64
	case d.dtype == 0x4:
		d.encoding = encoding_ldnf1sw_z_p_bi_s64
	case d.dtype == 0x5:
		d.encoding = encoding_ldnf1h_z_p_bi_u16
	case d.dtype == 0x6:
		d.encoding = encoding_ldnf1h_z_p_bi_u32
	case d.dtype == 0x7:
		d.encoding = encoding_ldnf1h_z_p_bi_u64
	case d.dtype == 0x8:
		d.encoding = encoding_ldnf1sh_z_p_bi_s64
	case d.dtype == 0x9:
		d.encoding = encoding_ldnf1sh_z_p_bi_s32
	case d.dtype == 0xa:
		d.encoding = encoding_ldnf1w_z_p_bi_u32
	case d.dtype == 0xb:
		d.encoding = encoding_ldnf1w_z_p_bi_u64
	case d.dtype == 0xc:
		d.encoding = encoding_ldnf1sb_z_p_bi_s64
	case d.dtype == 0xd:
		d.encoding = encoding_ldnf1sb_z_p_bi_s32
	case d.dtype == 0xe:
		d.encoding = encoding_ldnf1sb_z_p_bi_s16
	case d.dtype == 0xf:
		d.encoding = encoding_ldnf1d_z_p_bi_u64
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_ldqr_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_ldqr_ss
	d.msz = (ins >> 23) & 0x3
	d.ssz = (ins >> 21) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case (d.ssz & 0x2) == 0x2:
		err = errUnallocated
	case d.msz == 0x0 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqb_z_p_br_contiguous
	case d.msz == 0x0 && d.ssz == 0x1:
		d.encoding = encoding_ld1rob_z_p_br_contiguous
	case d.msz == 0x1 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqh_z_p_br_contiguous
	case d.msz == 0x1 && d.ssz == 0x1:
		d.encoding = encoding_ld1roh_z_p_br_contiguous
	case d.msz == 0x2 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqw_z_p_br_contiguous
	case d.msz == 0x2 && d.ssz == 0x1:
		d.encoding = encoding_ld1row_z_p_br_contiguous
	case d.msz == 0x3 && d.ssz == 0x0:
		d.encoding = encoding_ld1rqd_z_p_br_contiguous
	case d.msz == 0x3 && d.ssz == 0x1:
		d.encoding = encoding_ld1rod_z_p_br_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_cld_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cld_ss
	d.dtype = (ins >> 21) & 0xf
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.dtype == 0x0:
		d.encoding = encoding_ld1b_z_p_br_u8
	case d.dtype == 0x1:
		d.encoding = encoding_ld1b_z_p_br_u16
	case d.dtype == 0x2:
		d.encoding = encoding_ld1b_z_p_br_u32
	case d.dtype == 0x3:
		d.encoding = encoding_ld1b_z_p_br_u64
	case d.dtype == 0x4:
		d.encoding = encoding_ld1sw_z_p_br_s64
	case d.dtype == 0x5:
		d.encoding = encoding_ld1h_z_p_br_u16
	case d.dtype == 0x6:
		d.encoding = encoding_ld1h_z_p_br_u32
	case d.dtype == 0x7:
		d.encoding = encoding_ld1h_z_p_br_u64
	case d.dtype == 0x8:
		d.encoding = encoding_ld1sh_z_p_br_s64
	case d.dtype == 0x9:
		d.encoding = encoding_ld1sh_z_p_br_s32
	case d.dtype == 0xa:
		d.encoding = encoding_ld1w_z_p_br_u32
	case d.dtype == 0xb:
		d.encoding = encoding_ld1w_z_p_br_u64
	case d.dtype == 0xc:
		d.encoding = encoding_ld1sb_z_p_br_s64
	case d.dtype == 0xd:
		d.encoding = encoding_ld1sb_z_p_br_s32
	case d.dtype == 0xe:
		d.encoding = encoding_ld1sb_z_p_br_s16
	case d.dtype == 0xf:
		d.encoding = encoding_ld1d_z_p_br_u64
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memcld_sve_mem_cldff_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cldff_ss
	d.dtype = (ins >> 21) & 0xf
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.dtype == 0x0:
		d.encoding = encoding_ldff1b_z_p_br_u8
	case d.dtype == 0x1:
		d.encoding = encoding_ldff1b_z_p_br_u16
	case d.dtype == 0x2:
		d.encoding = encoding_ldff1b_z_p_br_u32
	case d.dtype == 0x3:
		d.encoding = encoding_ldff1b_z_p_br_u64
	case d.dtype == 0x4:
		d.encoding = encoding_ldff1sw_z_p_br_s64
	case d.dtype == 0x5:
		d.encoding = encoding_ldff1h_z_p_br_u16
	case d.dtype == 0x6:
		d.encoding = encoding_ldff1h_z_p_br_u32
	case d.dtype == 0x7:
		d.encoding = encoding_ldff1h_z_p_br_u64
	case d.dtype == 0x8:
		d.encoding = encoding_ldff1sh_z_p_br_s64
	case d.dtype == 0x9:
		d.encoding = encoding_ldff1sh_z_p_br_s32
	case d.dtype == 0xa:
		d.encoding = encoding_ldff1w_z_p_br_u32
	case d.dtype == 0xb:
		d.encoding = encoding_ldff1w_z_p_br_u64
	case d.dtype == 0xc:
		d.encoding = encoding_ldff1sb_z_p_br_s64
	case d.dtype == 0xd:
		d.encoding = encoding_ldff1sb_z_p_br_s32
	case d.dtype == 0xe:
		d.encoding = encoding_ldff1sb_z_p_br_s16
	case d.dtype == 0xf:
		d.encoding = encoding_ldff1d_z_p_br_u64
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x3
	op1 := (ins >> 21) & 0x3
	op2 := (ins >> 13) & 0x7
	op3 := (ins >> 4) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x1 && (op2&0x4) == 0x0 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && op1 == 0x3 && (op2&0x4) == 0x4 && op3 == 0x0:
		err = decode_sve_sve_mem64_sve_mem_64b_prfm_sv2(ins, d)
	case op0 == 0x0 && op1 == 0x3 && op3 == 0x1:
		err = errUnallocated
	case op0 == 0x0 && (op1&0x1) == 0x1 && (op2&0x4) == 0x0 && op3 == 0x0:
		err = decode_sve_sve_mem64_sve_mem_64b_prfm_sv(ins, d)
	case op0 != 0x0 && op1 == 0x3 && (op2&0x4) == 0x4:
		err = decode_sve_sve_mem64_sve_mem_64b_gld_sv2(ins, d)
	case op0 != 0x0 && (op1&0x1) == 0x1 && (op2&0x4) == 0x0:
		err = decode_sve_sve_mem64_sve_mem_64b_gld_sv(ins, d)
	case op1 == 0x0 && op2 == 0x5:
		err = errUnallocated
	case op1 == 0x0 && op2 == 0x7 && op3 == 0x0:
		err = decode_sve_sve_mem64_sve_mem_64b_prfm_vi(ins, d)
	case op1 == 0x0 && op2 == 0x7 && op3 == 0x1:
		err = errUnallocated
	case op1 == 0x0 && (op2&0x5) == 0x4:
		err = decode_sve_sve_mem64_sve_mem_64b_gldnt_vs(ins, d)
	case op1 == 0x1 && (op2&0x4) == 0x4:
		err = decode_sve_sve_mem64_sve_mem_64b_gld_vi(ins, d)
	case op1 == 0x2 && (op2&0x4) == 0x4:
		err = decode_sve_sve_mem64_sve_mem_64b_gld_vs2(ins, d)
	case (op1&0x1) == 0x0 && (op2&0x4) == 0x0:
		err = decode_sve_sve_mem64_sve_mem_64b_gld_vs(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_prfm_sv2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_prfm_sv2
	d.Zm = (ins >> 16) & 0x1f
	d.msz = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_bz_d_64_scaled
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_bz_d_64_scaled
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_bz_d_64_scaled
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_bz_d_64_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_prfm_sv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_prfm_sv
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.msz = (ins >> 13) & 0x3
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_bz_d_x32_scaled
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_bz_d_x32_scaled
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_bz_d_x32_scaled
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_bz_d_x32_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_gld_sv2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_gld_sv2
	d.opc = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_bz_d_64_scaled
	case d.opc == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_bz_d_64_scaled
	case d.opc == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_bz_d_64_scaled
	case d.opc == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_bz_d_64_scaled
	case d.opc == 0x2 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sw_z_p_bz_d_64_scaled
	case d.opc == 0x2 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sw_z_p_bz_d_64_scaled
	case d.opc == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_bz_d_64_scaled
	case d.opc == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_bz_d_64_scaled
	case d.opc == 0x3 && d.U == 0x0:
		err = errUnallocated
	case d.opc == 0x3 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1d_z_p_bz_d_64_scaled
	case d.opc == 0x3 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1d_z_p_bz_d_64_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_gld_sv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_gld_sv
	d.opc = (ins >> 23) & 0x3
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_bz_d_x32_scaled
	case d.opc == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_bz_d_x32_scaled
	case d.opc == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_bz_d_x32_scaled
	case d.opc == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_bz_d_x32_scaled
	case d.opc == 0x2 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sw_z_p_bz_d_x32_scaled
	case d.opc == 0x2 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sw_z_p_bz_d_x32_scaled
	case d.opc == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_bz_d_x32_scaled
	case d.opc == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_bz_d_x32_scaled
	case d.opc == 0x3 && d.U == 0x0:
		err = errUnallocated
	case d.opc == 0x3 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1d_z_p_bz_d_x32_scaled
	case d.opc == 0x3 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1d_z_p_bz_d_x32_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_prfm_vi(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_prfm_vi
	d.msz = (ins >> 23) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.prfop = (ins >> 0) & 0xf

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_prfb_i_p_ai_d
	case d.msz == 0x1:
		d.encoding = encoding_prfh_i_p_ai_d
	case d.msz == 0x2:
		d.encoding = encoding_prfw_i_p_ai_d
	case d.msz == 0x3:
		d.encoding = encoding_prfd_i_p_ai_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_gldnt_vs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_gldnt_vs
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.U == 0x0:
		d.encoding = encoding_ldnt1sb_z_p_ar_d_64_unscaled
	case d.msz == 0x0 && d.U == 0x1:
		d.encoding = encoding_ldnt1b_z_p_ar_d_64_unscaled
	case d.msz == 0x1 && d.U == 0x0:
		d.encoding = encoding_ldnt1sh_z_p_ar_d_64_unscaled
	case d.msz == 0x1 && d.U == 0x1:
		d.encoding = encoding_ldnt1h_z_p_ar_d_64_unscaled
	case d.msz == 0x2 && d.U == 0x0:
		d.encoding = encoding_ldnt1sw_z_p_ar_d_64_unscaled
	case d.msz == 0x2 && d.U == 0x1:
		d.encoding = encoding_ldnt1w_z_p_ar_d_64_unscaled
	case d.msz == 0x3 && d.U == 0x0:
		err = errUnallocated
	case d.msz == 0x3 && d.U == 0x1:
		d.encoding = encoding_ldnt1d_z_p_ar_d_64_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_gld_vi(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_gld_vi
	d.msz = (ins >> 23) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sb_z_p_ai_d
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sb_z_p_ai_d
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1b_z_p_ai_d
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1b_z_p_ai_d
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_ai_d
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_ai_d
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_ai_d
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_ai_d
	case d.msz == 0x2 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sw_z_p_ai_d
	case d.msz == 0x2 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sw_z_p_ai_d
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_ai_d
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_ai_d
	case d.msz == 0x3 && d.U == 0x0:
		err = errUnallocated
	case d.msz == 0x3 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1d_z_p_ai_d
	case d.msz == 0x3 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1d_z_p_ai_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_gld_vs2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_gld_vs2
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sb_z_p_bz_d_64_unscaled
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sb_z_p_bz_d_64_unscaled
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1b_z_p_bz_d_64_unscaled
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1b_z_p_bz_d_64_unscaled
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_bz_d_64_unscaled
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_bz_d_64_unscaled
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_bz_d_64_unscaled
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_bz_d_64_unscaled
	case d.msz == 0x2 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sw_z_p_bz_d_64_unscaled
	case d.msz == 0x2 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sw_z_p_bz_d_64_unscaled
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_bz_d_64_unscaled
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_bz_d_64_unscaled
	case d.msz == 0x3 && d.U == 0x0:
		err = errUnallocated
	case d.msz == 0x3 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1d_z_p_bz_d_64_unscaled
	case d.msz == 0x3 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1d_z_p_bz_d_64_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_mem64_sve_mem_64b_gld_vs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_64b_gld_vs
	d.msz = (ins >> 23) & 0x3
	d.xs = (ins >> 22) & 0x1
	d.Zm = (ins >> 16) & 0x1f
	d.U = (ins >> 14) & 0x1
	d.ff = (ins >> 13) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sb_z_p_bz_d_x32_unscaled
	case d.msz == 0x0 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sb_z_p_bz_d_x32_unscaled
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1b_z_p_bz_d_x32_unscaled
	case d.msz == 0x0 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1b_z_p_bz_d_x32_unscaled
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sh_z_p_bz_d_x32_unscaled
	case d.msz == 0x1 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sh_z_p_bz_d_x32_unscaled
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1h_z_p_bz_d_x32_unscaled
	case d.msz == 0x1 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1h_z_p_bz_d_x32_unscaled
	case d.msz == 0x2 && d.U == 0x0 && d.ff == 0x0:
		d.encoding = encoding_ld1sw_z_p_bz_d_x32_unscaled
	case d.msz == 0x2 && d.U == 0x0 && d.ff == 0x1:
		d.encoding = encoding_ldff1sw_z_p_bz_d_x32_unscaled
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1w_z_p_bz_d_x32_unscaled
	case d.msz == 0x2 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1w_z_p_bz_d_x32_unscaled
	case d.msz == 0x3 && d.U == 0x0:
		err = errUnallocated
	case d.msz == 0x3 && d.U == 0x1 && d.ff == 0x0:
		d.encoding = encoding_ld1d_z_p_bz_d_x32_unscaled
	case d.msz == 0x3 && d.U == 0x1 && d.ff == 0x1:
		d.encoding = encoding_ldff1d_z_p_bz_d_x32_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_cs(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 22) & 0x7
	op1 := (ins >> 14) & 0x1
	op2 := (ins >> 4) & 0x1

	switch {
	case (op0&0x4) == 0x0 && op1 == 0x0:
		err = errUnallocated
	case (op0&0x6) == 0x4 && op1 == 0x0:
		err = errUnallocated
	case op0 == 0x6 && op1 == 0x0 && op2 == 0x0:
		err = decode_sve_sve_memst_cs_sve_mem_pspill(ins, d)
	case op0 == 0x6 && op1 == 0x0 && op2 == 0x1:
		err = errUnallocated
	case op0 == 0x6 && op1 == 0x1:
		err = decode_sve_sve_memst_cs_sve_mem_spill(ins, d)
	case op0 == 0x7 && op1 == 0x0:
		err = errUnallocated
	case op0 != 0x6 && op1 == 0x1:
		err = decode_sve_sve_memst_cs_sve_mem_cst_ss(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_cs_sve_mem_pspill(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_pspill
	d.imm9h = (ins >> 16) & 0x3f
	d.imm9l = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Pt = (ins >> 0) & 0xf
	d.encoding = encoding_str_p_bi_
	return
}

func decode_sve_sve_memst_cs_sve_mem_spill(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_spill
	d.imm9h = (ins >> 16) & 0x3f
	d.imm9l = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f
	d.encoding = encoding_str_z_bi_
	return
}

func decode_sve_sve_memst_cs_sve_mem_cst_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cst_ss
	d.opc = (ins >> 22) & 0x7
	d.o2 = (ins >> 21) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case (d.opc & 0x6) == 0x0:
		d.encoding = encoding_st1b_z_p_br_
	case (d.opc & 0x6) == 0x2:
		d.encoding = encoding_st1h_z_p_br_
	case (d.opc & 0x6) == 0x4:
		d.encoding = encoding_st1w_z_p_br_
	case d.opc == 0x7 && d.o2 == 0x0:
		err = errUnallocated
	case d.opc == 0x7 && d.o2 == 0x1:
		d.encoding = encoding_st1d_z_p_br_
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_nt(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 21) & 0x3
	op1 := (ins >> 14) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x0:
		err = decode_sve_sve_memst_nt_sve_mem_sstnt_64b_vs(ins, d)
	case op0 == 0x0 && op1 == 0x1:
		err = decode_sve_sve_memst_nt_sve_mem_cstnt_ss(ins, d)
	case op0 == 0x2 && op1 == 0x0:
		err = decode_sve_sve_memst_nt_sve_mem_sstnt_32b_vs(ins, d)
	case op0 != 0x0 && op1 == 0x1:
		err = decode_sve_sve_memst_nt_sve_mem_est_ss(ins, d)
	case (op0&0x1) == 0x1 && op1 == 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_nt_sve_mem_sstnt_64b_vs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sstnt_64b_vs
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_stnt1b_z_p_ar_d_64_unscaled
	case d.msz == 0x1:
		d.encoding = encoding_stnt1h_z_p_ar_d_64_unscaled
	case d.msz == 0x2:
		d.encoding = encoding_stnt1w_z_p_ar_d_64_unscaled
	case d.msz == 0x3:
		d.encoding = encoding_stnt1d_z_p_ar_d_64_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_nt_sve_mem_cstnt_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cstnt_ss
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_stnt1b_z_p_br_contiguous
	case d.msz == 0x1:
		d.encoding = encoding_stnt1h_z_p_br_contiguous
	case d.msz == 0x2:
		d.encoding = encoding_stnt1w_z_p_br_contiguous
	case d.msz == 0x3:
		d.encoding = encoding_stnt1d_z_p_br_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_nt_sve_mem_sstnt_32b_vs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sstnt_32b_vs
	d.msz = (ins >> 23) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_stnt1b_z_p_ar_s_x32_unscaled
	case d.msz == 0x1:
		d.encoding = encoding_stnt1h_z_p_ar_s_x32_unscaled
	case d.msz == 0x2:
		d.encoding = encoding_stnt1w_z_p_ar_s_x32_unscaled
	case d.msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_nt_sve_mem_est_ss(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_est_ss
	d.msz = (ins >> 23) & 0x3
	d.opc = (ins >> 21) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.opc == 0x1:
		d.encoding = encoding_st2b_z_p_br_contiguous
	case d.msz == 0x0 && d.opc == 0x2:
		d.encoding = encoding_st3b_z_p_br_contiguous
	case d.msz == 0x0 && d.opc == 0x3:
		d.encoding = encoding_st4b_z_p_br_contiguous
	case d.msz == 0x1 && d.opc == 0x1:
		d.encoding = encoding_st2h_z_p_br_contiguous
	case d.msz == 0x1 && d.opc == 0x2:
		d.encoding = encoding_st3h_z_p_br_contiguous
	case d.msz == 0x1 && d.opc == 0x3:
		d.encoding = encoding_st4h_z_p_br_contiguous
	case d.msz == 0x2 && d.opc == 0x1:
		d.encoding = encoding_st2w_z_p_br_contiguous
	case d.msz == 0x2 && d.opc == 0x2:
		d.encoding = encoding_st3w_z_p_br_contiguous
	case d.msz == 0x2 && d.opc == 0x3:
		d.encoding = encoding_st4w_z_p_br_contiguous
	case d.msz == 0x3 && d.opc == 0x1:
		d.encoding = encoding_st2d_z_p_br_contiguous
	case d.msz == 0x3 && d.opc == 0x2:
		d.encoding = encoding_st3d_z_p_br_contiguous
	case d.msz == 0x3 && d.opc == 0x3:
		d.encoding = encoding_st4d_z_p_br_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 21) & 0x3

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_memst_ss_sve_mem_sst_vs_a(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_memst_ss_sve_mem_sst_sv_a(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_memst_ss_sve_mem_sst_vs_b(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_memst_ss_sve_mem_sst_sv_b(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss_sve_mem_sst_vs_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_vs_a
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.xs = (ins >> 14) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_z_p_bz_d_x32_unscaled
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bz_d_x32_unscaled
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bz_d_x32_unscaled
	case d.msz == 0x3:
		d.encoding = encoding_st1d_z_p_bz_d_x32_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss_sve_mem_sst_sv_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_sv_a
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.xs = (ins >> 14) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		err = errUnallocated
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bz_d_x32_scaled
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bz_d_x32_scaled
	case d.msz == 0x3:
		d.encoding = encoding_st1d_z_p_bz_d_x32_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss_sve_mem_sst_vs_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_vs_b
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.xs = (ins >> 14) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_z_p_bz_s_x32_unscaled
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bz_s_x32_unscaled
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bz_s_x32_unscaled
	case d.msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss_sve_mem_sst_sv_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_sv_b
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.xs = (ins >> 14) & 0x1
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		err = errUnallocated
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bz_s_x32_scaled
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bz_s_x32_scaled
	case d.msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss2(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 21) & 0x3

	switch {
	case op0 == 0x0:
		err = decode_sve_sve_memst_ss2_sve_mem_sst_vs2(ins, d)
	case op0 == 0x1:
		err = decode_sve_sve_memst_ss2_sve_mem_sst_sv2(ins, d)
	case op0 == 0x2:
		err = decode_sve_sve_memst_ss2_sve_mem_sst_vi_a(ins, d)
	case op0 == 0x3:
		err = decode_sve_sve_memst_ss2_sve_mem_sst_vi_b(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss2_sve_mem_sst_vs2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_vs2
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_z_p_bz_d_64_unscaled
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bz_d_64_unscaled
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bz_d_64_unscaled
	case d.msz == 0x3:
		d.encoding = encoding_st1d_z_p_bz_d_64_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss2_sve_mem_sst_sv2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_sv2
	d.msz = (ins >> 23) & 0x3
	d.Zm = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		err = errUnallocated
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bz_d_64_scaled
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bz_d_64_scaled
	case d.msz == 0x3:
		d.encoding = encoding_st1d_z_p_bz_d_64_scaled
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss2_sve_mem_sst_vi_a(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_vi_a
	d.msz = (ins >> 23) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_z_p_ai_d
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_ai_d
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_ai_d
	case d.msz == 0x3:
		d.encoding = encoding_st1d_z_p_ai_d
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_ss2_sve_mem_sst_vi_b(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_sst_vi_b
	d.msz = (ins >> 23) & 0x3
	d.imm5 = (ins >> 16) & 0x1f
	d.Pg = (ins >> 10) & 0x7
	d.Zn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_z_p_ai_s
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_ai_s
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_ai_s
	case d.msz == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_si(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 21) & 0x3
	op1 := (ins >> 20) & 0x1

	switch {
	case op0 == 0x0 && op1 == 0x1:
		err = decode_sve_sve_memst_si_sve_mem_cstnt_si(ins, d)
	case op0 != 0x0 && op1 == 0x1:
		err = decode_sve_sve_memst_si_sve_mem_est_si(ins, d)
	case op1 == 0x0:
		err = decode_sve_sve_memst_si_sve_mem_cst_si(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_si_sve_mem_cstnt_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cstnt_si
	d.msz = (ins >> 23) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_stnt1b_z_p_bi_contiguous
	case d.msz == 0x1:
		d.encoding = encoding_stnt1h_z_p_bi_contiguous
	case d.msz == 0x2:
		d.encoding = encoding_stnt1w_z_p_bi_contiguous
	case d.msz == 0x3:
		d.encoding = encoding_stnt1d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_si_sve_mem_est_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_est_si
	d.msz = (ins >> 23) & 0x3
	d.opc = (ins >> 21) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0 && d.opc == 0x1:
		d.encoding = encoding_st2b_z_p_bi_contiguous
	case d.msz == 0x0 && d.opc == 0x2:
		d.encoding = encoding_st3b_z_p_bi_contiguous
	case d.msz == 0x0 && d.opc == 0x3:
		d.encoding = encoding_st4b_z_p_bi_contiguous
	case d.msz == 0x1 && d.opc == 0x1:
		d.encoding = encoding_st2h_z_p_bi_contiguous
	case d.msz == 0x1 && d.opc == 0x2:
		d.encoding = encoding_st3h_z_p_bi_contiguous
	case d.msz == 0x1 && d.opc == 0x3:
		d.encoding = encoding_st4h_z_p_bi_contiguous
	case d.msz == 0x2 && d.opc == 0x1:
		d.encoding = encoding_st2w_z_p_bi_contiguous
	case d.msz == 0x2 && d.opc == 0x2:
		d.encoding = encoding_st3w_z_p_bi_contiguous
	case d.msz == 0x2 && d.opc == 0x3:
		d.encoding = encoding_st4w_z_p_bi_contiguous
	case d.msz == 0x3 && d.opc == 0x1:
		d.encoding = encoding_st2d_z_p_bi_contiguous
	case d.msz == 0x3 && d.opc == 0x2:
		d.encoding = encoding_st3d_z_p_bi_contiguous
	case d.msz == 0x3 && d.opc == 0x3:
		d.encoding = encoding_st4d_z_p_bi_contiguous
	default:
		err = errUnmatched
	}
	return
}

func decode_sve_sve_memst_si_sve_mem_cst_si(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_sve_mem_cst_si
	d.msz = (ins >> 23) & 0x3
	d.size = (ins >> 21) & 0x3
	d.imm4 = (ins >> 16) & 0xf
	d.Pg = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Zt = (ins >> 0) & 0x1f

	switch {
	case d.msz == 0x0:
		d.encoding = encoding_st1b_z_p_bi_
	case d.msz == 0x1:
		d.encoding = encoding_st1h_z_p_bi_
	case d.msz == 0x2:
		d.encoding = encoding_st1w_z_p_bi_
	case d.msz == 0x3:
		d.encoding = encoding_st1d_z_p_bi_
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 23) & 0x7

	switch {
	case (op0 & 0x6) == 0x0:
		err = decode_dpimm_pcreladdr(ins, d)
	case op0 == 0x2:
		err = decode_dpimm_addsub_imm(ins, d)
	case op0 == 0x3:
		err = decode_dpimm_addsub_immtags(ins, d)
	case op0 == 0x4:
		err = decode_dpimm_log_imm(ins, d)
	case op0 == 0x5:
		err = decode_dpimm_movewide(ins, d)
	case op0 == 0x6:
		err = decode_dpimm_bitfield(ins, d)
	case op0 == 0x7:
		err = decode_dpimm_extract(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_pcreladdr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_pcreladdr
	d.op = (ins >> 31) & 0x1
	d.immlo = (ins >> 29) & 0x3
	d.immhi = (ins >> 5) & 0x7ffff
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_ADR_only_pcreladdr
	case d.op == 0x1:
		d.encoding = encoding_ADRP_only_pcreladdr
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_addsub_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_addsub_imm
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.sh = (ins >> 22) & 0x1
	d.imm12 = (ins >> 10) & 0xfff
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_ADD_32_addsub_imm
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_ADDS_32S_addsub_imm
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x0:
		d.encoding = encoding_SUB_32_addsub_imm
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x1:
		d.encoding = encoding_SUBS_32S_addsub_imm
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_ADD_64_addsub_imm
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_ADDS_64S_addsub_imm
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0:
		d.encoding = encoding_SUB_64_addsub_imm
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x1:
		d.encoding = encoding_SUBS_64S_addsub_imm
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_addsub_immtags(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_addsub_immtags
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.o2 = (ins >> 22) & 0x1
	d.uimm6 = (ins >> 16) & 0x3f
	d.op3 = (ins >> 14) & 0x3
	d.uimm4 = (ins >> 10) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.o2 == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.o2 == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x1 && d.o2 == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_ADDG_64_addsub_immtags
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_SUBG_64_addsub_immtags
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_log_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_log_imm
	d.sf = (ins >> 31) & 0x1
	d.opc = (ins >> 29) & 0x3
	d.N = (ins >> 22) & 0x1
	d.immr = (ins >> 16) & 0x3f
	d.imms = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.sf == 0x0 && d.N == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.opc == 0x0 && d.N == 0x0:
		d.encoding = encoding_AND_32_log_imm
	case d.sf == 0x0 && d.opc == 0x1 && d.N == 0x0:
		d.encoding = encoding_ORR_32_log_imm
	case d.sf == 0x0 && d.opc == 0x2 && d.N == 0x0:
		d.encoding = encoding_EOR_32_log_imm
	case d.sf == 0x0 && d.opc == 0x3 && d.N == 0x0:
		d.encoding = encoding_ANDS_32S_log_imm
	case d.sf == 0x1 && d.opc == 0x0:
		d.encoding = encoding_AND_64_log_imm
	case d.sf == 0x1 && d.opc == 0x1:
		d.encoding = encoding_ORR_64_log_imm
	case d.sf == 0x1 && d.opc == 0x2:
		d.encoding = encoding_EOR_64_log_imm
	case d.sf == 0x1 && d.opc == 0x3:
		d.encoding = encoding_ANDS_64S_log_imm
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_movewide(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_movewide
	d.sf = (ins >> 31) & 0x1
	d.opc = (ins >> 29) & 0x3
	d.hw = (ins >> 21) & 0x3
	d.imm16 = (ins >> 5) & 0xffff
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && (d.hw&0x2) == 0x2:
		err = errUnallocated
	case d.sf == 0x0 && d.opc == 0x0 && (d.hw&0x2) == 0x0:
		d.encoding = encoding_MOVN_32_movewide
	case d.sf == 0x0 && d.opc == 0x2 && (d.hw&0x2) == 0x0:
		d.encoding = encoding_MOVZ_32_movewide
	case d.sf == 0x0 && d.opc == 0x3 && (d.hw&0x2) == 0x0:
		d.encoding = encoding_MOVK_32_movewide
	case d.sf == 0x1 && d.opc == 0x0:
		d.encoding = encoding_MOVN_64_movewide
	case d.sf == 0x1 && d.opc == 0x2:
		d.encoding = encoding_MOVZ_64_movewide
	case d.sf == 0x1 && d.opc == 0x3:
		d.encoding = encoding_MOVK_64_movewide
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_bitfield(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_bitfield
	d.sf = (ins >> 31) & 0x1
	d.opc = (ins >> 29) & 0x3
	d.N = (ins >> 22) & 0x1
	d.immr = (ins >> 16) & 0x3f
	d.imms = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x3:
		err = errUnallocated
	case d.sf == 0x0 && d.N == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.opc == 0x0 && d.N == 0x0:
		d.encoding = encoding_SBFM_32M_bitfield
	case d.sf == 0x0 && d.opc == 0x1 && d.N == 0x0:
		d.encoding = encoding_BFM_32M_bitfield
	case d.sf == 0x0 && d.opc == 0x2 && d.N == 0x0:
		d.encoding = encoding_UBFM_32M_bitfield
	case d.sf == 0x1 && d.N == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.opc == 0x0 && d.N == 0x1:
		d.encoding = encoding_SBFM_64M_bitfield
	case d.sf == 0x1 && d.opc == 0x1 && d.N == 0x1:
		d.encoding = encoding_BFM_64M_bitfield
	case d.sf == 0x1 && d.opc == 0x2 && d.N == 0x1:
		d.encoding = encoding_UBFM_64M_bitfield
	default:
		err = errUnmatched
	}
	return
}

func decode_dpimm_extract(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_extract
	d.sf = (ins >> 31) & 0x1
	d.op21 = (ins >> 29) & 0x3
	d.N = (ins >> 22) & 0x1
	d.o0 = (ins >> 21) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.imms = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.op21 & 0x1) == 0x1:
		err = errUnallocated
	case d.op21 == 0x0 && d.o0 == 0x1:
		err = errUnallocated
	case (d.op21 & 0x2) == 0x2:
		err = errUnallocated
	case d.sf == 0x0 && (d.imms&0x20) == 0x20:
		err = errUnallocated
	case d.sf == 0x0 && d.N == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.op21 == 0x0 && d.N == 0x0 && d.o0 == 0x0 && (d.imms&0x20) == 0x0:
		d.encoding = encoding_EXTR_32_extract
	case d.sf == 0x1 && d.N == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.op21 == 0x0 && d.N == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_EXTR_64_extract
	default:
		err = errUnmatched
	}
	return
}

func decode_control(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 29) & 0x7
	op1 := (ins >> 12) & 0x3fff
	op2 := (ins >> 0) & 0x1f

	switch {
	case op0 == 0x2 && (op1&0x2000) == 0x0:
		err = decode_control_condbranch(ins, d)
	case op0 == 0x6 && (op1&0x3000) == 0x0:
		err = decode_control_exception(ins, d)
	case op0 == 0x6 && op1 == 0x1031:
		err = decode_control_systeminstrswithreg(ins, d)
	case op0 == 0x6 && op1 == 0x1032 && op2 == 0x1f:
		err = decode_control_hints(ins, d)
	case op0 == 0x6 && op1 == 0x1033:
		err = decode_control_barriers(ins, d)
	case op0 == 0x6 && (op1&0x3f8f) == 0x1004:
		err = decode_control_pstate(ins, d)
	case op0 == 0x6 && (op1&0x3f80) == 0x1200:
		err = decode_control_systemresult(ins, d)
	case op0 == 0x6 && (op1&0x3d80) == 0x1080:
		err = decode_control_systeminstrs(ins, d)
	case op0 == 0x6 && (op1&0x3d00) == 0x1100:
		err = decode_control_systemmove(ins, d)
	case op0 == 0x6 && (op1&0x2000) == 0x2000:
		err = decode_control_branch_reg(ins, d)
	case (op0 & 0x3) == 0x0:
		err = decode_control_branch_imm(ins, d)
	case (op0&0x3) == 0x1 && (op1&0x2000) == 0x0:
		err = decode_control_compbranch(ins, d)
	case (op0&0x3) == 0x1 && (op1&0x2000) == 0x2000:
		err = decode_control_testbranch(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_control_condbranch(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_condbranch
	d.o1 = (ins >> 24) & 0x1
	d.imm19 = (ins >> 5) & 0x7ffff
	d.o0 = (ins >> 4) & 0x1
	d.cond = (ins >> 0) & 0xf

	switch {
	case d.o1 == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_B_only_condbranch
	case d.o1 == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_BC_only_condbranch
	case d.o1 == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_control_exception(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_exception
	d.opc = (ins >> 21) & 0x7
	d.imm16 = (ins >> 5) & 0xffff
	d.op2 = (ins >> 2) & 0x7
	d.LL = (ins >> 0) & 0x3

	switch {
	case d.op2 == 0x1:
		err = errUnallocated
	case (d.op2 & 0x6) == 0x2:
		err = errUnallocated
	case (d.op2 & 0x4) == 0x4:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x0 && d.LL == 0x0:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x0 && d.LL == 0x1:
		d.encoding = encoding_SVC_EX_exception
	case d.opc == 0x0 && d.op2 == 0x0 && d.LL == 0x2:
		d.encoding = encoding_HVC_EX_exception
	case d.opc == 0x0 && d.op2 == 0x0 && d.LL == 0x3:
		d.encoding = encoding_SMC_EX_exception
	case d.opc == 0x1 && d.op2 == 0x0 && (d.LL&0x1) == 0x1:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x0 && d.LL == 0x0:
		d.encoding = encoding_BRK_EX_exception
	case d.opc == 0x1 && d.op2 == 0x0 && (d.LL&0x2) == 0x2:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x0 && (d.LL&0x1) == 0x1:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x0 && d.LL == 0x0:
		d.encoding = encoding_HLT_EX_exception
	case d.opc == 0x2 && d.op2 == 0x0 && (d.LL&0x2) == 0x2:
		err = errUnallocated
	case d.opc == 0x3 && d.op2 == 0x0 && d.LL == 0x0:
		d.encoding = encoding_TCANCEL_EX_exception
	case d.opc == 0x3 && d.op2 == 0x0 && d.LL == 0x1:
		err = errUnallocated
	case d.opc == 0x3 && d.op2 == 0x0 && (d.LL&0x2) == 0x2:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x0:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x0 && d.LL == 0x0:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x0 && d.LL == 0x1:
		d.encoding = encoding_DCPS1_DC_exception
	case d.opc == 0x5 && d.op2 == 0x0 && d.LL == 0x2:
		d.encoding = encoding_DCPS2_DC_exception
	case d.opc == 0x5 && d.op2 == 0x0 && d.LL == 0x3:
		d.encoding = encoding_DCPS3_DC_exception
	case d.opc == 0x6 && d.op2 == 0x0:
		err = errUnallocated
	case d.opc == 0x7 && d.op2 == 0x0:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_control_systeminstrswithreg(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_systeminstrswithreg
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.CRm != 0x0:
		err = errUnallocated
	case d.CRm == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_WFET_only_systeminstrswithreg
	case d.CRm == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_WFIT_only_systeminstrswithreg
	case d.CRm == 0x0 && (d.op2&0x6) == 0x2:
		err = errUnallocated
	case d.CRm == 0x0 && (d.op2&0x4) == 0x4:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_control_hints(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_hints
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7

	switch {
	case d.CRm == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_NOP_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_YIELD_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x2:
		d.encoding = encoding_WFE_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x3:
		d.encoding = encoding_WFI_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x4:
		d.encoding = encoding_SEV_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x5:
		d.encoding = encoding_SEVL_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x6:
		d.encoding = encoding_DGH_HI_hints
	case d.CRm == 0x0 && d.op2 == 0x7:
		d.encoding = encoding_XPACLRI_HI_hints
	case d.CRm == 0x1 && d.op2 == 0x0:
		d.encoding = encoding_PACIA1716_HI_hints
	case d.CRm == 0x1 && d.op2 == 0x2:
		d.encoding = encoding_PACIB1716_HI_hints
	case d.CRm == 0x1 && d.op2 == 0x4:
		d.encoding = encoding_AUTIA1716_HI_hints
	case d.CRm == 0x1 && d.op2 == 0x6:
		d.encoding = encoding_AUTIB1716_HI_hints
	case d.CRm == 0x2 && d.op2 == 0x0:
		d.encoding = encoding_ESB_HI_hints
	case d.CRm == 0x2 && d.op2 == 0x1:
		d.encoding = encoding_PSB_HC_hints
	case d.CRm == 0x2 && d.op2 == 0x2:
		d.encoding = encoding_TSB_HC_hints
	case d.CRm == 0x2 && d.op2 == 0x4:
		d.encoding = encoding_CSDB_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x0:
		d.encoding = encoding_PACIAZ_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x1:
		d.encoding = encoding_PACIASP_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x2:
		d.encoding = encoding_PACIBZ_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x3:
		d.encoding = encoding_PACIBSP_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x4:
		d.encoding = encoding_AUTIAZ_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x5:
		d.encoding = encoding_AUTIASP_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x6:
		d.encoding = encoding_AUTIBZ_HI_hints
	case d.CRm == 0x3 && d.op2 == 0x7:
		d.encoding = encoding_AUTIBSP_HI_hints
	case d.CRm == 0x4 && (d.op2&0x1) == 0x0:
		d.encoding = encoding_BTI_HB_hints
	default:
		d.encoding = encoding_HINT_HM_hints
	}
	return
}

func decode_control_barriers(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_barriers
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.op2 == 0x0:
		err = errUnallocated
	case d.op2 == 0x1 && d.Rt != 0x1f:
		err = errUnallocated
	case d.op2 == 0x2 && d.Rt == 0x1f:
		d.encoding = encoding_CLREX_BN_barriers
	case d.op2 == 0x4 && d.Rt == 0x1f:
		d.encoding = encoding_DSB_BO_barriers
	case d.op2 == 0x5 && d.Rt == 0x1f:
		d.encoding = encoding_DMB_BO_barriers
	case d.op2 == 0x6 && d.Rt == 0x1f:
		d.encoding = encoding_ISB_BI_barriers
	case d.op2 == 0x7 && d.Rt != 0x1f:
		err = errUnallocated
	case d.op2 == 0x7 && d.Rt == 0x1f:
		d.encoding = encoding_SB_only_barriers
	case (d.CRm&0x2) == 0x0 && d.op2 == 0x1 && d.Rt == 0x1f:
		err = errUnallocated
	case (d.CRm&0x3) == 0x2 && d.op2 == 0x1 && d.Rt == 0x1f:
		d.encoding = encoding_DSB_BOn_barriers
	case (d.CRm&0x3) == 0x3 && d.op2 == 0x1 && d.Rt == 0x1f:
		err = errUnallocated
	case d.CRm == 0x0 && d.op2 == 0x3 && d.Rt == 0x1f:
		d.encoding = encoding_TCOMMIT_only_barriers
	case d.CRm == 0x1 && d.op2 == 0x3:
		err = errUnallocated
	case (d.CRm&0xe) == 0x2 && d.op2 == 0x3:
		err = errUnallocated
	case (d.CRm&0xc) == 0x4 && d.op2 == 0x3:
		err = errUnallocated
	case (d.CRm&0x8) == 0x8 && d.op2 == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_control_pstate(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_pstate
	d.op1 = (ins >> 16) & 0x7
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.Rt != 0x1f:
		err = errUnallocated
	case d.Rt == 0x1f:
		d.encoding = encoding_MSR_SI_pstate
	case d.op1 == 0x0 && d.op2 == 0x0 && d.Rt == 0x1f:
		d.encoding = encoding_CFINV_M_pstate
	case d.op1 == 0x0 && d.op2 == 0x1 && d.Rt == 0x1f:
		d.encoding = encoding_XAFLAG_M_pstate
	case d.op1 == 0x0 && d.op2 == 0x2 && d.Rt == 0x1f:
		d.encoding = encoding_AXFLAG_M_pstate
	default:
		err = errUnmatched
	}
	return
}

func decode_control_systemresult(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_systemresult
	d.op1 = (ins >> 16) & 0x7
	d.CRn = (ins >> 12) & 0xf
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.op1 != 0x3:
		err = errUnallocated
	case d.op1 == 0x3 && d.CRn != 0x3:
		err = errUnallocated
	case d.op1 == 0x3 && d.CRn == 0x3 && d.op2 != 0x3:
		err = errUnallocated
	case d.op1 == 0x3 && d.CRn == 0x3 && (d.CRm&0xe) != 0x0 && d.op2 == 0x3:
		err = errUnallocated
	case d.op1 == 0x3 && d.CRn == 0x3 && d.CRm == 0x0 && d.op2 == 0x3:
		d.encoding = encoding_TSTART_BR_systemresult
	case d.op1 == 0x3 && d.CRn == 0x3 && d.CRm == 0x1 && d.op2 == 0x3:
		d.encoding = encoding_TTEST_BR_systemresult
	default:
		err = errUnmatched
	}
	return
}

func decode_control_systeminstrs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_systeminstrs
	d.L = (ins >> 21) & 0x1
	d.op1 = (ins >> 16) & 0x7
	d.CRn = (ins >> 12) & 0xf
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x0:
		d.encoding = encoding_SYS_CR_systeminstrs
	case d.L == 0x1:
		d.encoding = encoding_SYSL_RC_systeminstrs
	default:
		err = errUnmatched
	}
	return
}

func decode_control_systemmove(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_systemmove
	d.L = (ins >> 21) & 0x1
	d.o0 = (ins >> 19) & 0x1
	d.op1 = (ins >> 16) & 0x7
	d.CRn = (ins >> 12) & 0xf
	d.CRm = (ins >> 8) & 0xf
	d.op2 = (ins >> 5) & 0x7
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x0:
		d.encoding = encoding_MSR_SR_systemmove
	case d.L == 0x1:
		d.encoding = encoding_MRS_RS_systemmove
	default:
		err = errUnmatched
	}
	return
}

func decode_control_branch_reg(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_branch_reg
	d.opc = (ins >> 21) & 0xf
	d.op2 = (ins >> 16) & 0x1f
	d.op3 = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.op4 = (ins >> 0) & 0x1f

	switch {
	case d.op2 != 0x1f:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x0 && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x0 && d.op4 == 0x0:
		d.encoding = encoding_BR_64_branch_reg
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x1:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x2 && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x2 && d.op4 == 0x1f:
		d.encoding = encoding_BRAAZ_64_branch_reg
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x3 && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && d.op3 == 0x3 && d.op4 == 0x1f:
		d.encoding = encoding_BRABZ_64_branch_reg
	case d.opc == 0x0 && d.op2 == 0x1f && (d.op3&0x3c) == 0x4:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && (d.op3&0x38) == 0x8:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && (d.op3&0x30) == 0x10:
		err = errUnallocated
	case d.opc == 0x0 && d.op2 == 0x1f && (d.op3&0x20) == 0x20:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x0 && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x0 && d.op4 == 0x0:
		d.encoding = encoding_BLR_64_branch_reg
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x1:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x2 && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x2 && d.op4 == 0x1f:
		d.encoding = encoding_BLRAAZ_64_branch_reg
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x3 && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && d.op3 == 0x3 && d.op4 == 0x1f:
		d.encoding = encoding_BLRABZ_64_branch_reg
	case d.opc == 0x1 && d.op2 == 0x1f && (d.op3&0x3c) == 0x4:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && (d.op3&0x38) == 0x8:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && (d.op3&0x30) == 0x10:
		err = errUnallocated
	case d.opc == 0x1 && d.op2 == 0x1f && (d.op3&0x20) == 0x20:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x0 && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x0 && d.op4 == 0x0:
		d.encoding = encoding_RET_64R_branch_reg
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x1:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn != 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn != 0x1f && d.op4 == 0x1f:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn == 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn == 0x1f && d.op4 == 0x1f:
		d.encoding = encoding_RETAA_64E_branch_reg
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn != 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn != 0x1f && d.op4 == 0x1f:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn == 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn == 0x1f && d.op4 == 0x1f:
		d.encoding = encoding_RETAB_64E_branch_reg
	case d.opc == 0x2 && d.op2 == 0x1f && (d.op3&0x3c) == 0x4:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && (d.op3&0x38) == 0x8:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && (d.op3&0x30) == 0x10:
		err = errUnallocated
	case d.opc == 0x2 && d.op2 == 0x1f && (d.op3&0x20) == 0x20:
		err = errUnallocated
	case d.opc == 0x3 && d.op2 == 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn != 0x1f && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn != 0x1f && d.op4 == 0x0:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn == 0x1f && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn == 0x1f && d.op4 == 0x0:
		d.encoding = encoding_ERET_64E_branch_reg
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x1:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn != 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn != 0x1f && d.op4 == 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn == 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x2 && d.Rn == 0x1f && d.op4 == 0x1f:
		d.encoding = encoding_ERETAA_64E_branch_reg
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn != 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn != 0x1f && d.op4 == 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn == 0x1f && d.op4 != 0x1f:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && d.op3 == 0x3 && d.Rn == 0x1f && d.op4 == 0x1f:
		d.encoding = encoding_ERETAB_64E_branch_reg
	case d.opc == 0x4 && d.op2 == 0x1f && (d.op3&0x3c) == 0x4:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && (d.op3&0x38) == 0x8:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && (d.op3&0x30) == 0x10:
		err = errUnallocated
	case d.opc == 0x4 && d.op2 == 0x1f && (d.op3&0x20) == 0x20:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x1f && d.op3 != 0x0:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn != 0x1f && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn != 0x1f && d.op4 == 0x0:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn == 0x1f && d.op4 != 0x0:
		err = errUnallocated
	case d.opc == 0x5 && d.op2 == 0x1f && d.op3 == 0x0 && d.Rn == 0x1f && d.op4 == 0x0:
		d.encoding = encoding_DRPS_64E_branch_reg
	case (d.opc&0xe) == 0x6 && d.op2 == 0x1f:
		err = errUnallocated
	case d.opc == 0x8 && d.op2 == 0x1f && (d.op3&0x3e) == 0x0:
		err = errUnallocated
	case d.opc == 0x8 && d.op2 == 0x1f && d.op3 == 0x2:
		d.encoding = encoding_BRAA_64P_branch_reg
	case d.opc == 0x8 && d.op2 == 0x1f && d.op3 == 0x3:
		d.encoding = encoding_BRAB_64P_branch_reg
	case d.opc == 0x8 && d.op2 == 0x1f && (d.op3&0x3c) == 0x4:
		err = errUnallocated
	case d.opc == 0x8 && d.op2 == 0x1f && (d.op3&0x38) == 0x8:
		err = errUnallocated
	case d.opc == 0x8 && d.op2 == 0x1f && (d.op3&0x30) == 0x10:
		err = errUnallocated
	case d.opc == 0x8 && d.op2 == 0x1f && (d.op3&0x20) == 0x20:
		err = errUnallocated
	case d.opc == 0x9 && d.op2 == 0x1f && (d.op3&0x3e) == 0x0:
		err = errUnallocated
	case d.opc == 0x9 && d.op2 == 0x1f && d.op3 == 0x2:
		d.encoding = encoding_BLRAA_64P_branch_reg
	case d.opc == 0x9 && d.op2 == 0x1f && d.op3 == 0x3:
		d.encoding = encoding_BLRAB_64P_branch_reg
	case d.opc == 0x9 && d.op2 == 0x1f && (d.op3&0x3c) == 0x4:
		err = errUnallocated
	case d.opc == 0x9 && d.op2 == 0x1f && (d.op3&0x38) == 0x8:
		err = errUnallocated
	case d.opc == 0x9 && d.op2 == 0x1f && (d.op3&0x30) == 0x10:
		err = errUnallocated
	case d.opc == 0x9 && d.op2 == 0x1f && (d.op3&0x20) == 0x20:
		err = errUnallocated
	case (d.opc&0xe) == 0xa && d.op2 == 0x1f:
		err = errUnallocated
	case (d.opc&0xc) == 0xc && d.op2 == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_control_branch_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_branch_imm
	d.op = (ins >> 31) & 0x1
	d.imm26 = (ins >> 0) & 0x3ffffff

	switch {
	case d.op == 0x0:
		d.encoding = encoding_B_only_branch_imm
	case d.op == 0x1:
		d.encoding = encoding_BL_only_branch_imm
	default:
		err = errUnmatched
	}
	return
}

func decode_control_compbranch(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_compbranch
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 24) & 0x1
	d.imm19 = (ins >> 5) & 0x7ffff
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.sf == 0x0 && d.op == 0x0:
		d.encoding = encoding_CBZ_32_compbranch
	case d.sf == 0x0 && d.op == 0x1:
		d.encoding = encoding_CBNZ_32_compbranch
	case d.sf == 0x1 && d.op == 0x0:
		d.encoding = encoding_CBZ_64_compbranch
	case d.sf == 0x1 && d.op == 0x1:
		d.encoding = encoding_CBNZ_64_compbranch
	default:
		err = errUnmatched
	}
	return
}

func decode_control_testbranch(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_testbranch
	d.b5 = (ins >> 31) & 0x1
	d.op = (ins >> 24) & 0x1
	d.b40 = (ins >> 19) & 0x1f
	d.imm14 = (ins >> 5) & 0x3fff
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0:
		d.encoding = encoding_TBZ_only_testbranch
	case d.op == 0x1:
		d.encoding = encoding_TBNZ_only_testbranch
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 28) & 0xf
	op1 := (ins >> 26) & 0x1
	op2 := (ins >> 23) & 0x3
	op3 := (ins >> 16) & 0x3f
	op4 := (ins >> 10) & 0x3

	switch {
	case (op0&0xb) == 0x0 && op1 == 0x0 && op2 == 0x0 && (op3&0x20) == 0x20:
		err = decode_ldst_comswappr(ins, d)
	case (op0&0xb) == 0x0 && op1 == 0x1 && op2 == 0x0 && op3 == 0x0:
		err = decode_ldst_asisdlse(ins, d)
	case (op0&0xb) == 0x0 && op1 == 0x1 && op2 == 0x1 && (op3&0x20) == 0x0:
		err = decode_ldst_asisdlsep(ins, d)
	case (op0&0xb) == 0x0 && op1 == 0x1 && (op2&0x2) == 0x0 && (op3&0x20) == 0x20:
		err = errUnallocated
	case (op0&0xb) == 0x0 && op1 == 0x1 && op2 == 0x2 && (op3&0x1f) == 0x0:
		err = decode_ldst_asisdlso(ins, d)
	case (op0&0xb) == 0x0 && op1 == 0x1 && op2 == 0x3:
		err = decode_ldst_asisdlsop(ins, d)
	case (op0&0xb) == 0x0 && op1 == 0x1 && (op2&0x1) == 0x0 && (op3&0x10) == 0x10:
		err = errUnallocated
	case (op0&0xb) == 0x0 && op1 == 0x1 && (op2&0x1) == 0x0 && (op3&0x8) == 0x8:
		err = errUnallocated
	case (op0&0xb) == 0x0 && op1 == 0x1 && (op2&0x1) == 0x0 && (op3&0x4) == 0x4:
		err = errUnallocated
	case (op0&0xb) == 0x0 && op1 == 0x1 && (op2&0x1) == 0x0 && (op3&0x2) == 0x2:
		err = errUnallocated
	case (op0&0xb) == 0x0 && op1 == 0x1 && (op2&0x1) == 0x0 && (op3&0x1) == 0x1:
		err = errUnallocated
	case op0 == 0xd && op1 == 0x0 && (op2&0x2) == 0x2 && (op3&0x20) == 0x20:
		err = decode_ldst_ldsttags(ins, d)
	case (op0&0xb) == 0x8 && op1 == 0x0 && op2 == 0x0 && (op3&0x20) == 0x20:
		err = decode_ldst_ldstexclp(ins, d)
	case (op0&0xb) == 0x8 && op1 == 0x1:
		err = errUnallocated
	case (op0&0x3) == 0x0 && op1 == 0x0 && op2 == 0x0 && (op3&0x20) == 0x0:
		err = decode_ldst_ldstexclr(ins, d)
	case (op0&0x3) == 0x0 && op1 == 0x0 && op2 == 0x1 && (op3&0x20) == 0x0:
		err = decode_ldst_ldstord(ins, d)
	case (op0&0x3) == 0x0 && op1 == 0x0 && op2 == 0x1 && (op3&0x20) == 0x20:
		err = decode_ldst_comswap(ins, d)
	case (op0&0x3) == 0x1 && op1 == 0x0 && (op2&0x2) == 0x2 && (op3&0x20) == 0x0 && op4 == 0x0:
		err = decode_ldst_ldapstl_unscaled(ins, d)
	case (op0&0x3) == 0x1 && (op2&0x2) == 0x0:
		err = decode_ldst_loadlit(ins, d)
	case (op0&0x3) == 0x1 && (op2&0x2) == 0x2 && (op3&0x20) == 0x0 && op4 == 0x1:
		err = decode_ldst_memcms(ins, d)
	case (op0&0x3) == 0x2 && op2 == 0x0:
		err = decode_ldst_ldstnapair_offs(ins, d)
	case (op0&0x3) == 0x2 && op2 == 0x1:
		err = decode_ldst_ldstpair_post(ins, d)
	case (op0&0x3) == 0x2 && op2 == 0x2:
		err = decode_ldst_ldstpair_off(ins, d)
	case (op0&0x3) == 0x2 && op2 == 0x3:
		err = decode_ldst_ldstpair_pre(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x0 && op4 == 0x0:
		err = decode_ldst_ldst_unscaled(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x0 && op4 == 0x1:
		err = decode_ldst_ldst_immpost(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x0 && op4 == 0x2:
		err = decode_ldst_ldst_unpriv(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x0 && op4 == 0x3:
		err = decode_ldst_ldst_immpre(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x20 && op4 == 0x0:
		err = decode_ldst_memop(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x20 && op4 == 0x2:
		err = decode_ldst_ldst_regoff(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x0 && (op3&0x20) == 0x20 && (op4&0x1) == 0x1:
		err = decode_ldst_ldst_pac(ins, d)
	case (op0&0x3) == 0x3 && (op2&0x2) == 0x2:
		err = decode_ldst_ldst_pos(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_comswappr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_comswappr
	d.sz = (ins >> 30) & 0x1
	d.L = (ins >> 22) & 0x1
	d.Rs = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.Rt2 != 0x1f:
		err = errUnallocated
	case d.sz == 0x0 && d.L == 0x0 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASP_CP32_comswappr
	case d.sz == 0x0 && d.L == 0x0 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASPL_CP32_comswappr
	case d.sz == 0x0 && d.L == 0x1 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASPA_CP32_comswappr
	case d.sz == 0x0 && d.L == 0x1 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASPAL_CP32_comswappr
	case d.sz == 0x1 && d.L == 0x0 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASP_CP64_comswappr
	case d.sz == 0x1 && d.L == 0x0 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASPL_CP64_comswappr
	case d.sz == 0x1 && d.L == 0x1 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASPA_CP64_comswappr
	case d.sz == 0x1 && d.L == 0x1 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASPAL_CP64_comswappr
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_asisdlse(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdlse
	d.Q = (ins >> 30) & 0x1
	d.L = (ins >> 22) & 0x1
	d.opcode = (ins >> 12) & 0xf
	d.size = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_ST4_asisdlse_R4
	case d.L == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_ST1_asisdlse_R4_4v
	case d.L == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_ST3_asisdlse_R3
	case d.L == 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_ST1_asisdlse_R3_3v
	case d.L == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_ST1_asisdlse_R1_1v
	case d.L == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_ST2_asisdlse_R2
	case d.L == 0x0 && d.opcode == 0x9:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_ST1_asisdlse_R2_2v
	case d.L == 0x0 && d.opcode == 0xb:
		err = errUnallocated
	case d.L == 0x0 && (d.opcode&0xc) == 0xc:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_LD4_asisdlse_R4
	case d.L == 0x1 && d.opcode == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_LD1_asisdlse_R4_4v
	case d.L == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_LD3_asisdlse_R3
	case d.L == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_LD1_asisdlse_R3_3v
	case d.L == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_LD1_asisdlse_R1_1v
	case d.L == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_LD2_asisdlse_R2
	case d.L == 0x1 && d.opcode == 0x9:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_LD1_asisdlse_R2_2v
	case d.L == 0x1 && d.opcode == 0xb:
		err = errUnallocated
	case d.L == 0x1 && (d.opcode&0xc) == 0xc:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_asisdlsep(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdlsep
	d.Q = (ins >> 30) & 0x1
	d.L = (ins >> 22) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 12) & 0xf
	d.size = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0x9:
		err = errUnallocated
	case d.L == 0x0 && d.opcode == 0xb:
		err = errUnallocated
	case d.L == 0x0 && (d.opcode&0xc) == 0xc:
		err = errUnallocated
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0x0:
		d.encoding = encoding_ST4_asisdlsep_R4_r
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0x2:
		d.encoding = encoding_ST1_asisdlsep_R4_r4
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0x4:
		d.encoding = encoding_ST3_asisdlsep_R3_r
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0x6:
		d.encoding = encoding_ST1_asisdlsep_R3_r3
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0x7:
		d.encoding = encoding_ST1_asisdlsep_R1_r1
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0x8:
		d.encoding = encoding_ST2_asisdlsep_R2_r
	case d.L == 0x0 && d.Rm != 0x1f && d.opcode == 0xa:
		d.encoding = encoding_ST1_asisdlsep_R2_r2
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0x0:
		d.encoding = encoding_ST4_asisdlsep_I4_i
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0x2:
		d.encoding = encoding_ST1_asisdlsep_I4_i4
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0x4:
		d.encoding = encoding_ST3_asisdlsep_I3_i
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0x6:
		d.encoding = encoding_ST1_asisdlsep_I3_i3
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0x7:
		d.encoding = encoding_ST1_asisdlsep_I1_i1
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0x8:
		d.encoding = encoding_ST2_asisdlsep_I2_i
	case d.L == 0x0 && d.Rm == 0x1f && d.opcode == 0xa:
		d.encoding = encoding_ST1_asisdlsep_I2_i2
	case d.L == 0x1 && d.opcode == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0x9:
		err = errUnallocated
	case d.L == 0x1 && d.opcode == 0xb:
		err = errUnallocated
	case d.L == 0x1 && (d.opcode&0xc) == 0xc:
		err = errUnallocated
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0x0:
		d.encoding = encoding_LD4_asisdlsep_R4_r
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0x2:
		d.encoding = encoding_LD1_asisdlsep_R4_r4
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0x4:
		d.encoding = encoding_LD3_asisdlsep_R3_r
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0x6:
		d.encoding = encoding_LD1_asisdlsep_R3_r3
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0x7:
		d.encoding = encoding_LD1_asisdlsep_R1_r1
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0x8:
		d.encoding = encoding_LD2_asisdlsep_R2_r
	case d.L == 0x1 && d.Rm != 0x1f && d.opcode == 0xa:
		d.encoding = encoding_LD1_asisdlsep_R2_r2
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0x0:
		d.encoding = encoding_LD4_asisdlsep_I4_i
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0x2:
		d.encoding = encoding_LD1_asisdlsep_I4_i4
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0x4:
		d.encoding = encoding_LD3_asisdlsep_I3_i
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0x6:
		d.encoding = encoding_LD1_asisdlsep_I3_i3
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0x7:
		d.encoding = encoding_LD1_asisdlsep_I1_i1
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0x8:
		d.encoding = encoding_LD2_asisdlsep_I2_i
	case d.L == 0x1 && d.Rm == 0x1f && d.opcode == 0xa:
		d.encoding = encoding_LD1_asisdlsep_I2_i2
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_asisdlso(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdlso
	d.Q = (ins >> 30) & 0x1
	d.L = (ins >> 22) & 0x1
	d.R = (ins >> 21) & 0x1
	d.opcode = (ins >> 13) & 0x7
	d.S = (ins >> 12) & 0x1
	d.size = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x0 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_ST1_asisdlso_B1_1b
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_ST3_asisdlso_B3_3b
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST1_asisdlso_H1_1h
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST3_asisdlso_H3_3h
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_ST1_asisdlso_S1_1s
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x4 && (d.size&0x2) == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST1_asisdlso_D1_1d
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x4 && d.S == 0x1 && d.size == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_ST3_asisdlso_S3_3s
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST3_asisdlso_D3_3d
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_ST2_asisdlso_B2_2b
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_ST4_asisdlso_B4_4b
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST2_asisdlso_H2_2h
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST4_asisdlso_H4_4h
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_ST2_asisdlso_S2_2s
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST2_asisdlso_D2_2d
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_ST4_asisdlso_S4_4s
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST4_asisdlso_D4_4d
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_LD1_asisdlso_B1_1b
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_LD3_asisdlso_B3_3b
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD1_asisdlso_H1_1h
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD3_asisdlso_H3_3h
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_LD1_asisdlso_S1_1s
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x4 && (d.size&0x2) == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD1_asisdlso_D1_1d
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x4 && d.S == 0x1 && d.size == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_LD3_asisdlso_S3_3s
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD3_asisdlso_D3_3d
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x6 && d.S == 0x0:
		d.encoding = encoding_LD1R_asisdlso_R1
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x6 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x7 && d.S == 0x0:
		d.encoding = encoding_LD3R_asisdlso_R3
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x7 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_LD2_asisdlso_B2_2b
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_LD4_asisdlso_B4_4b
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD2_asisdlso_H2_2h
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD4_asisdlso_H4_4h
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_LD2_asisdlso_S2_2s
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD2_asisdlso_D2_2d
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_LD4_asisdlso_S4_4s
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD4_asisdlso_D4_4d
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x6 && d.S == 0x0:
		d.encoding = encoding_LD2R_asisdlso_R2
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x6 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x7 && d.S == 0x0:
		d.encoding = encoding_LD4R_asisdlso_R4
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x7 && d.S == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_asisdlsop(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdlsop
	d.Q = (ins >> 30) & 0x1
	d.L = (ins >> 22) & 0x1
	d.R = (ins >> 21) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 13) & 0x7
	d.S = (ins >> 12) & 0x1
	d.size = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.L == 0x0 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x4 && (d.size&0x2) == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x4 && d.S == 0x1 && d.size == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x0:
		d.encoding = encoding_ST1_asisdlsop_BX1_r1b
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x1:
		d.encoding = encoding_ST3_asisdlsop_BX3_r3b
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST1_asisdlsop_HX1_r1h
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST3_asisdlsop_HX3_r3h
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_ST1_asisdlsop_SX1_r1s
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST1_asisdlsop_DX1_r1d
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_ST3_asisdlsop_SX3_r3s
	case d.L == 0x0 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST3_asisdlsop_DX3_r3d
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x0:
		d.encoding = encoding_ST1_asisdlsop_B1_i1b
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x1:
		d.encoding = encoding_ST3_asisdlsop_B3_i3b
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST1_asisdlsop_H1_i1h
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST3_asisdlsop_H3_i3h
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_ST1_asisdlsop_S1_i1s
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST1_asisdlsop_D1_i1d
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_ST3_asisdlsop_S3_i3s
	case d.L == 0x0 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST3_asisdlsop_D3_i3d
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x0:
		d.encoding = encoding_ST2_asisdlsop_BX2_r2b
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x1:
		d.encoding = encoding_ST4_asisdlsop_BX4_r4b
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST2_asisdlsop_HX2_r2h
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST4_asisdlsop_HX4_r4h
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_ST2_asisdlsop_SX2_r2s
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST2_asisdlsop_DX2_r2d
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_ST4_asisdlsop_SX4_r4s
	case d.L == 0x0 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST4_asisdlsop_DX4_r4d
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x0:
		d.encoding = encoding_ST2_asisdlsop_B2_i2b
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x1:
		d.encoding = encoding_ST4_asisdlsop_B4_i4b
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST2_asisdlsop_H2_i2h
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_ST4_asisdlsop_H4_i4h
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_ST2_asisdlsop_S2_i2s
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST2_asisdlsop_D2_i2d
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_ST4_asisdlsop_S4_i4s
	case d.L == 0x0 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_ST4_asisdlsop_D4_i4d
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x4 && (d.size&0x2) == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x4 && d.S == 0x1 && d.size == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x6 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.opcode == 0x7 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x0:
		d.encoding = encoding_LD1_asisdlsop_BX1_r1b
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x1:
		d.encoding = encoding_LD3_asisdlsop_BX3_r3b
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD1_asisdlsop_HX1_r1h
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD3_asisdlsop_HX3_r3h
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_LD1_asisdlsop_SX1_r1s
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD1_asisdlsop_DX1_r1d
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_LD3_asisdlsop_SX3_r3s
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD3_asisdlsop_DX3_r3d
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x6 && d.S == 0x0:
		d.encoding = encoding_LD1R_asisdlsop_RX1_r
	case d.L == 0x1 && d.R == 0x0 && d.Rm != 0x1f && d.opcode == 0x7 && d.S == 0x0:
		d.encoding = encoding_LD3R_asisdlsop_RX3_r
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x0:
		d.encoding = encoding_LD1_asisdlsop_B1_i1b
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x1:
		d.encoding = encoding_LD3_asisdlsop_B3_i3b
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD1_asisdlsop_H1_i1h
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD3_asisdlsop_H3_i3h
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_LD1_asisdlsop_S1_i1s
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD1_asisdlsop_D1_i1d
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_LD3_asisdlsop_S3_i3s
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD3_asisdlsop_D3_i3d
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x6 && d.S == 0x0:
		d.encoding = encoding_LD1R_asisdlsop_R1_i
	case d.L == 0x1 && d.R == 0x0 && d.Rm == 0x1f && d.opcode == 0x7 && d.S == 0x0:
		d.encoding = encoding_LD3R_asisdlsop_R3_i
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x2 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x3 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x4 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.size == 0x2:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x3:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x5 && d.S == 0x1 && (d.size&0x1) == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x6 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.opcode == 0x7 && d.S == 0x1:
		err = errUnallocated
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x0:
		d.encoding = encoding_LD2_asisdlsop_BX2_r2b
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x1:
		d.encoding = encoding_LD4_asisdlsop_BX4_r4b
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD2_asisdlsop_HX2_r2h
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD4_asisdlsop_HX4_r4h
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_LD2_asisdlsop_SX2_r2s
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD2_asisdlsop_DX2_r2d
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_LD4_asisdlsop_SX4_r4s
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD4_asisdlsop_DX4_r4d
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x6 && d.S == 0x0:
		d.encoding = encoding_LD2R_asisdlsop_RX2_r
	case d.L == 0x1 && d.R == 0x1 && d.Rm != 0x1f && d.opcode == 0x7 && d.S == 0x0:
		d.encoding = encoding_LD4R_asisdlsop_RX4_r
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x0:
		d.encoding = encoding_LD2_asisdlsop_B2_i2b
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x1:
		d.encoding = encoding_LD4_asisdlsop_B4_i4b
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x2 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD2_asisdlsop_H2_i2h
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x3 && (d.size&0x1) == 0x0:
		d.encoding = encoding_LD4_asisdlsop_H4_i4h
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x4 && d.size == 0x0:
		d.encoding = encoding_LD2_asisdlsop_S2_i2s
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x4 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD2_asisdlsop_D2_i2d
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x5 && d.size == 0x0:
		d.encoding = encoding_LD4_asisdlsop_S4_i4s
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x5 && d.S == 0x0 && d.size == 0x1:
		d.encoding = encoding_LD4_asisdlsop_D4_i4d
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x6 && d.S == 0x0:
		d.encoding = encoding_LD2R_asisdlsop_R2_i
	case d.L == 0x1 && d.R == 0x1 && d.Rm == 0x1f && d.opcode == 0x7 && d.S == 0x0:
		d.encoding = encoding_LD4R_asisdlsop_R4_i
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldsttags(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldsttags
	d.opc = (ins >> 22) & 0x3
	d.imm9 = (ins >> 12) & 0x1ff
	d.op2 = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_STG_64Spost_ldsttags
	case d.opc == 0x0 && d.op2 == 0x2:
		d.encoding = encoding_STG_64Soffset_ldsttags
	case d.opc == 0x0 && d.op2 == 0x3:
		d.encoding = encoding_STG_64Spre_ldsttags
	case d.opc == 0x0 && d.imm9 == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_STZGM_64bulk_ldsttags
	case d.opc == 0x1 && d.op2 == 0x0:
		d.encoding = encoding_LDG_64Loffset_ldsttags
	case d.opc == 0x1 && d.op2 == 0x1:
		d.encoding = encoding_STZG_64Spost_ldsttags
	case d.opc == 0x1 && d.op2 == 0x2:
		d.encoding = encoding_STZG_64Soffset_ldsttags
	case d.opc == 0x1 && d.op2 == 0x3:
		d.encoding = encoding_STZG_64Spre_ldsttags
	case d.opc == 0x2 && d.op2 == 0x1:
		d.encoding = encoding_ST2G_64Spost_ldsttags
	case d.opc == 0x2 && d.op2 == 0x2:
		d.encoding = encoding_ST2G_64Soffset_ldsttags
	case d.opc == 0x2 && d.op2 == 0x3:
		d.encoding = encoding_ST2G_64Spre_ldsttags
	case d.opc == 0x2 && d.imm9 != 0x0 && d.op2 == 0x0:
		err = errUnallocated
	case d.opc == 0x2 && d.imm9 == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_STGM_64bulk_ldsttags
	case d.opc == 0x3 && d.op2 == 0x1:
		d.encoding = encoding_STZ2G_64Spost_ldsttags
	case d.opc == 0x3 && d.op2 == 0x2:
		d.encoding = encoding_STZ2G_64Soffset_ldsttags
	case d.opc == 0x3 && d.op2 == 0x3:
		d.encoding = encoding_STZ2G_64Spre_ldsttags
	case d.opc == 0x3 && d.imm9 != 0x0 && d.op2 == 0x0:
		err = errUnallocated
	case d.opc == 0x3 && d.imm9 == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_LDGM_64bulk_ldsttags
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstexclp(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstexclp
	d.sz = (ins >> 30) & 0x1
	d.L = (ins >> 22) & 0x1
	d.Rs = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.sz == 0x0 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STXP_SP32_ldstexclp
	case d.sz == 0x0 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLXP_SP32_ldstexclp
	case d.sz == 0x0 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDXP_LP32_ldstexclp
	case d.sz == 0x0 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAXP_LP32_ldstexclp
	case d.sz == 0x1 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STXP_SP64_ldstexclp
	case d.sz == 0x1 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLXP_SP64_ldstexclp
	case d.sz == 0x1 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDXP_LP64_ldstexclp
	case d.sz == 0x1 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAXP_LP64_ldstexclp
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstexclr(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstexclr
	d.size = (ins >> 30) & 0x3
	d.L = (ins >> 22) & 0x1
	d.Rs = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STXRB_SR32_ldstexclr
	case d.size == 0x0 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLXRB_SR32_ldstexclr
	case d.size == 0x0 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDXRB_LR32_ldstexclr
	case d.size == 0x0 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAXRB_LR32_ldstexclr
	case d.size == 0x1 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STXRH_SR32_ldstexclr
	case d.size == 0x1 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLXRH_SR32_ldstexclr
	case d.size == 0x1 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDXRH_LR32_ldstexclr
	case d.size == 0x1 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAXRH_LR32_ldstexclr
	case d.size == 0x2 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STXR_SR32_ldstexclr
	case d.size == 0x2 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLXR_SR32_ldstexclr
	case d.size == 0x2 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDXR_LR32_ldstexclr
	case d.size == 0x2 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAXR_LR32_ldstexclr
	case d.size == 0x3 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STXR_SR64_ldstexclr
	case d.size == 0x3 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLXR_SR64_ldstexclr
	case d.size == 0x3 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDXR_LR64_ldstexclr
	case d.size == 0x3 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAXR_LR64_ldstexclr
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstord(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstord
	d.size = (ins >> 30) & 0x3
	d.L = (ins >> 22) & 0x1
	d.Rs = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STLLRB_SL32_ldstord
	case d.size == 0x0 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLRB_SL32_ldstord
	case d.size == 0x0 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDLARB_LR32_ldstord
	case d.size == 0x0 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDARB_LR32_ldstord
	case d.size == 0x1 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STLLRH_SL32_ldstord
	case d.size == 0x1 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLRH_SL32_ldstord
	case d.size == 0x1 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDLARH_LR32_ldstord
	case d.size == 0x1 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDARH_LR32_ldstord
	case d.size == 0x2 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STLLR_SL32_ldstord
	case d.size == 0x2 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLR_SL32_ldstord
	case d.size == 0x2 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDLAR_LR32_ldstord
	case d.size == 0x2 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAR_LR32_ldstord
	case d.size == 0x3 && d.L == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_STLLR_SL64_ldstord
	case d.size == 0x3 && d.L == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_STLR_SL64_ldstord
	case d.size == 0x3 && d.L == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_LDLAR_LR64_ldstord
	case d.size == 0x3 && d.L == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_LDAR_LR64_ldstord
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_comswap(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_comswap
	d.size = (ins >> 30) & 0x3
	d.L = (ins >> 22) & 0x1
	d.Rs = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.Rt2 != 0x1f:
		err = errUnallocated
	case d.size == 0x0 && d.L == 0x0 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASB_C32_comswap
	case d.size == 0x0 && d.L == 0x0 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASLB_C32_comswap
	case d.size == 0x0 && d.L == 0x1 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASAB_C32_comswap
	case d.size == 0x0 && d.L == 0x1 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASALB_C32_comswap
	case d.size == 0x1 && d.L == 0x0 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASH_C32_comswap
	case d.size == 0x1 && d.L == 0x0 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASLH_C32_comswap
	case d.size == 0x1 && d.L == 0x1 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASAH_C32_comswap
	case d.size == 0x1 && d.L == 0x1 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASALH_C32_comswap
	case d.size == 0x2 && d.L == 0x0 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CAS_C32_comswap
	case d.size == 0x2 && d.L == 0x0 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASL_C32_comswap
	case d.size == 0x2 && d.L == 0x1 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASA_C32_comswap
	case d.size == 0x2 && d.L == 0x1 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASAL_C32_comswap
	case d.size == 0x3 && d.L == 0x0 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CAS_C64_comswap
	case d.size == 0x3 && d.L == 0x0 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASL_C64_comswap
	case d.size == 0x3 && d.L == 0x1 && d.o0 == 0x0 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASA_C64_comswap
	case d.size == 0x3 && d.L == 0x1 && d.o0 == 0x1 && d.Rt2 == 0x1f:
		d.encoding = encoding_CASAL_C64_comswap
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldapstl_unscaled(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldapstl_unscaled
	d.size = (ins >> 30) & 0x3
	d.opc = (ins >> 22) & 0x3
	d.imm9 = (ins >> 12) & 0x1ff
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STLURB_32_ldapstl_unscaled
	case d.size == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDAPURB_32_ldapstl_unscaled
	case d.size == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDAPURSB_64_ldapstl_unscaled
	case d.size == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDAPURSB_32_ldapstl_unscaled
	case d.size == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STLURH_32_ldapstl_unscaled
	case d.size == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDAPURH_32_ldapstl_unscaled
	case d.size == 0x1 && d.opc == 0x2:
		d.encoding = encoding_LDAPURSH_64_ldapstl_unscaled
	case d.size == 0x1 && d.opc == 0x3:
		d.encoding = encoding_LDAPURSH_32_ldapstl_unscaled
	case d.size == 0x2 && d.opc == 0x0:
		d.encoding = encoding_STLUR_32_ldapstl_unscaled
	case d.size == 0x2 && d.opc == 0x1:
		d.encoding = encoding_LDAPUR_32_ldapstl_unscaled
	case d.size == 0x2 && d.opc == 0x2:
		d.encoding = encoding_LDAPURSW_64_ldapstl_unscaled
	case d.size == 0x2 && d.opc == 0x3:
		err = errUnallocated
	case d.size == 0x3 && d.opc == 0x0:
		d.encoding = encoding_STLUR_64_ldapstl_unscaled
	case d.size == 0x3 && d.opc == 0x1:
		d.encoding = encoding_LDAPUR_64_ldapstl_unscaled
	case d.size == 0x3 && d.opc == 0x2:
		err = errUnallocated
	case d.size == 0x3 && d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_loadlit(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_loadlit
	d.opc = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.imm19 = (ins >> 5) & 0x7ffff
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.V == 0x0:
		d.encoding = encoding_LDR_32_loadlit
	case d.opc == 0x0 && d.V == 0x1:
		d.encoding = encoding_LDR_S_loadlit
	case d.opc == 0x1 && d.V == 0x0:
		d.encoding = encoding_LDR_64_loadlit
	case d.opc == 0x1 && d.V == 0x1:
		d.encoding = encoding_LDR_D_loadlit
	case d.opc == 0x2 && d.V == 0x0:
		d.encoding = encoding_LDRSW_64_loadlit
	case d.opc == 0x2 && d.V == 0x1:
		d.encoding = encoding_LDR_Q_loadlit
	case d.opc == 0x3 && d.V == 0x0:
		d.encoding = encoding_PRFM_P_loadlit
	case d.opc == 0x3 && d.V == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_memcms(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_memcms
	d.size = (ins >> 30) & 0x3
	d.o0 = (ins >> 26) & 0x1
	d.op1 = (ins >> 22) & 0x3
	d.Rs = (ins >> 16) & 0x1f
	d.op2 = (ins >> 12) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_CPYFP_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_CPYFPWT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x2:
		d.encoding = encoding_CPYFPRT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x3:
		d.encoding = encoding_CPYFPT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x4:
		d.encoding = encoding_CPYFPWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x5:
		d.encoding = encoding_CPYFPWTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x6:
		d.encoding = encoding_CPYFPRTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x7:
		d.encoding = encoding_CPYFPTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x8:
		d.encoding = encoding_CPYFPRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0x9:
		d.encoding = encoding_CPYFPWTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0xa:
		d.encoding = encoding_CPYFPRTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0xb:
		d.encoding = encoding_CPYFPTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0xc:
		d.encoding = encoding_CPYFPN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0xd:
		d.encoding = encoding_CPYFPWTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0xe:
		d.encoding = encoding_CPYFPRTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x0 && d.op2 == 0xf:
		d.encoding = encoding_CPYFPTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x0:
		d.encoding = encoding_CPYFM_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x1:
		d.encoding = encoding_CPYFMWT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x2:
		d.encoding = encoding_CPYFMRT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x3:
		d.encoding = encoding_CPYFMT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x4:
		d.encoding = encoding_CPYFMWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x5:
		d.encoding = encoding_CPYFMWTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x6:
		d.encoding = encoding_CPYFMRTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x7:
		d.encoding = encoding_CPYFMTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x8:
		d.encoding = encoding_CPYFMRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0x9:
		d.encoding = encoding_CPYFMWTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0xa:
		d.encoding = encoding_CPYFMRTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0xb:
		d.encoding = encoding_CPYFMTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0xc:
		d.encoding = encoding_CPYFMN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0xd:
		d.encoding = encoding_CPYFMWTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0xe:
		d.encoding = encoding_CPYFMRTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x1 && d.op2 == 0xf:
		d.encoding = encoding_CPYFMTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x0:
		d.encoding = encoding_CPYFE_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x1:
		d.encoding = encoding_CPYFEWT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x2:
		d.encoding = encoding_CPYFERT_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x3:
		d.encoding = encoding_CPYFET_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x4:
		d.encoding = encoding_CPYFEWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x5:
		d.encoding = encoding_CPYFEWTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x6:
		d.encoding = encoding_CPYFERTWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x7:
		d.encoding = encoding_CPYFETWN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x8:
		d.encoding = encoding_CPYFERN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0x9:
		d.encoding = encoding_CPYFEWTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0xa:
		d.encoding = encoding_CPYFERTRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0xb:
		d.encoding = encoding_CPYFETRN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0xc:
		d.encoding = encoding_CPYFEN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0xd:
		d.encoding = encoding_CPYFEWTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0xe:
		d.encoding = encoding_CPYFERTN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x2 && d.op2 == 0xf:
		d.encoding = encoding_CPYFETN_CPY_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x0:
		d.encoding = encoding_SETP_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x1:
		d.encoding = encoding_SETPT_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x2:
		d.encoding = encoding_SETPN_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x3:
		d.encoding = encoding_SETPTN_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x4:
		d.encoding = encoding_SETM_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x5:
		d.encoding = encoding_SETMT_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x6:
		d.encoding = encoding_SETMN_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x7:
		d.encoding = encoding_SETMTN_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x8:
		d.encoding = encoding_SETE_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0x9:
		d.encoding = encoding_SETET_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0xa:
		d.encoding = encoding_SETEN_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && d.op2 == 0xb:
		d.encoding = encoding_SETETN_SET_memcms
	case d.o0 == 0x0 && d.op1 == 0x3 && (d.op2&0xc) == 0xc:
		err = errUnallocated
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_CPYP_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_CPYPWT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x2:
		d.encoding = encoding_CPYPRT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x3:
		d.encoding = encoding_CPYPT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x4:
		d.encoding = encoding_CPYPWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x5:
		d.encoding = encoding_CPYPWTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x6:
		d.encoding = encoding_CPYPRTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x7:
		d.encoding = encoding_CPYPTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x8:
		d.encoding = encoding_CPYPRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0x9:
		d.encoding = encoding_CPYPWTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0xa:
		d.encoding = encoding_CPYPRTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0xb:
		d.encoding = encoding_CPYPTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0xc:
		d.encoding = encoding_CPYPN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0xd:
		d.encoding = encoding_CPYPWTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0xe:
		d.encoding = encoding_CPYPRTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x0 && d.op2 == 0xf:
		d.encoding = encoding_CPYPTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x0:
		d.encoding = encoding_CPYM_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x1:
		d.encoding = encoding_CPYMWT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x2:
		d.encoding = encoding_CPYMRT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x3:
		d.encoding = encoding_CPYMT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x4:
		d.encoding = encoding_CPYMWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x5:
		d.encoding = encoding_CPYMWTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x6:
		d.encoding = encoding_CPYMRTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x7:
		d.encoding = encoding_CPYMTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x8:
		d.encoding = encoding_CPYMRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0x9:
		d.encoding = encoding_CPYMWTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0xa:
		d.encoding = encoding_CPYMRTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0xb:
		d.encoding = encoding_CPYMTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0xc:
		d.encoding = encoding_CPYMN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0xd:
		d.encoding = encoding_CPYMWTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0xe:
		d.encoding = encoding_CPYMRTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x1 && d.op2 == 0xf:
		d.encoding = encoding_CPYMTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x0:
		d.encoding = encoding_CPYE_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x1:
		d.encoding = encoding_CPYEWT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x2:
		d.encoding = encoding_CPYERT_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x3:
		d.encoding = encoding_CPYET_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x4:
		d.encoding = encoding_CPYEWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x5:
		d.encoding = encoding_CPYEWTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x6:
		d.encoding = encoding_CPYERTWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x7:
		d.encoding = encoding_CPYETWN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x8:
		d.encoding = encoding_CPYERN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0x9:
		d.encoding = encoding_CPYEWTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0xa:
		d.encoding = encoding_CPYERTRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0xb:
		d.encoding = encoding_CPYETRN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0xc:
		d.encoding = encoding_CPYEN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0xd:
		d.encoding = encoding_CPYEWTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0xe:
		d.encoding = encoding_CPYERTN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x2 && d.op2 == 0xf:
		d.encoding = encoding_CPYETN_CPY_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x0:
		d.encoding = encoding_SETGP_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x1:
		d.encoding = encoding_SETGPT_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x2:
		d.encoding = encoding_SETGPN_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x3:
		d.encoding = encoding_SETGPTN_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x4:
		d.encoding = encoding_SETGM_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x5:
		d.encoding = encoding_SETGMT_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x6:
		d.encoding = encoding_SETGMN_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x7:
		d.encoding = encoding_SETGMTN_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x8:
		d.encoding = encoding_SETGE_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0x9:
		d.encoding = encoding_SETGET_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0xa:
		d.encoding = encoding_SETGEN_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && d.op2 == 0xb:
		d.encoding = encoding_SETGETN_SET_memcms
	case d.o0 == 0x1 && d.op1 == 0x3 && (d.op2&0xc) == 0xc:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstnapair_offs(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstnapair_offs
	d.opc = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.L = (ins >> 22) & 0x1
	d.imm7 = (ins >> 15) & 0x7f
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STNP_32_ldstnapair_offs
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDNP_32_ldstnapair_offs
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STNP_S_ldstnapair_offs
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDNP_S_ldstnapair_offs
	case d.opc == 0x1 && d.V == 0x0:
		err = errUnallocated
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STNP_D_ldstnapair_offs
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDNP_D_ldstnapair_offs
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STNP_64_ldstnapair_offs
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDNP_64_ldstnapair_offs
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STNP_Q_ldstnapair_offs
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDNP_Q_ldstnapair_offs
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstpair_post(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstpair_post
	d.opc = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.L = (ins >> 22) & 0x1
	d.imm7 = (ins >> 15) & 0x7f
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STP_32_ldstpair_post
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDP_32_ldstpair_post
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_S_ldstpair_post
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_S_ldstpair_post
	case d.opc == 0x1 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STGP_64_ldstpair_post
	case d.opc == 0x1 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDPSW_64_ldstpair_post
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_D_ldstpair_post
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_D_ldstpair_post
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STP_64_ldstpair_post
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDP_64_ldstpair_post
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_Q_ldstpair_post
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_Q_ldstpair_post
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstpair_off(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstpair_off
	d.opc = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.L = (ins >> 22) & 0x1
	d.imm7 = (ins >> 15) & 0x7f
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STP_32_ldstpair_off
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDP_32_ldstpair_off
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_S_ldstpair_off
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_S_ldstpair_off
	case d.opc == 0x1 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STGP_64_ldstpair_off
	case d.opc == 0x1 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDPSW_64_ldstpair_off
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_D_ldstpair_off
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_D_ldstpair_off
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STP_64_ldstpair_off
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDP_64_ldstpair_off
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_Q_ldstpair_off
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_Q_ldstpair_off
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldstpair_pre(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldstpair_pre
	d.opc = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.L = (ins >> 22) & 0x1
	d.imm7 = (ins >> 15) & 0x7f
	d.Rt2 = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STP_32_ldstpair_pre
	case d.opc == 0x0 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDP_32_ldstpair_pre
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_S_ldstpair_pre
	case d.opc == 0x0 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_S_ldstpair_pre
	case d.opc == 0x1 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STGP_64_ldstpair_pre
	case d.opc == 0x1 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDPSW_64_ldstpair_pre
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_D_ldstpair_pre
	case d.opc == 0x1 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_D_ldstpair_pre
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x0:
		d.encoding = encoding_STP_64_ldstpair_pre
	case d.opc == 0x2 && d.V == 0x0 && d.L == 0x1:
		d.encoding = encoding_LDP_64_ldstpair_pre
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x0:
		d.encoding = encoding_STP_Q_ldstpair_pre
	case d.opc == 0x2 && d.V == 0x1 && d.L == 0x1:
		d.encoding = encoding_LDP_Q_ldstpair_pre
	case d.opc == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_unscaled(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_unscaled
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.opc = (ins >> 22) & 0x3
	d.imm9 = (ins >> 12) & 0x1ff
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x1) == 0x1 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STURB_32_ldst_unscaled
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDURB_32_ldst_unscaled
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDURSB_64_ldst_unscaled
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDURSB_32_ldst_unscaled
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STUR_B_ldst_unscaled
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDUR_B_ldst_unscaled
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x2:
		d.encoding = encoding_STUR_Q_ldst_unscaled
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x3:
		d.encoding = encoding_LDUR_Q_ldst_unscaled
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STURH_32_ldst_unscaled
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDURH_32_ldst_unscaled
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDURSH_64_ldst_unscaled
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDURSH_32_ldst_unscaled
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STUR_H_ldst_unscaled
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDUR_H_ldst_unscaled
	case (d.size&0x2) == 0x2 && d.V == 0x0 && d.opc == 0x3:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STUR_32_ldst_unscaled
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDUR_32_ldst_unscaled
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDURSW_64_ldst_unscaled
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STUR_S_ldst_unscaled
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDUR_S_ldst_unscaled
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STUR_64_ldst_unscaled
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDUR_64_ldst_unscaled
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_PRFUM_P_ldst_unscaled
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STUR_D_ldst_unscaled
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDUR_D_ldst_unscaled
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_immpost(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_immpost
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.opc = (ins >> 22) & 0x3
	d.imm9 = (ins >> 12) & 0x1ff
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x1) == 0x1 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRB_32_ldst_immpost
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRB_32_ldst_immpost
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSB_64_ldst_immpost
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSB_32_ldst_immpost
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_B_ldst_immpost
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_B_ldst_immpost
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x2:
		d.encoding = encoding_STR_Q_ldst_immpost
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x3:
		d.encoding = encoding_LDR_Q_ldst_immpost
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRH_32_ldst_immpost
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRH_32_ldst_immpost
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSH_64_ldst_immpost
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSH_32_ldst_immpost
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_H_ldst_immpost
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_H_ldst_immpost
	case (d.size&0x2) == 0x2 && d.V == 0x0 && d.opc == 0x3:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_32_ldst_immpost
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_32_ldst_immpost
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSW_64_ldst_immpost
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_S_ldst_immpost
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_S_ldst_immpost
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_64_ldst_immpost
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_64_ldst_immpost
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x2:
		err = errUnallocated
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_D_ldst_immpost
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_D_ldst_immpost
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_unpriv(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_unpriv
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.opc = (ins >> 22) & 0x3
	d.imm9 = (ins >> 12) & 0x1ff
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.V == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STTRB_32_ldst_unpriv
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDTRB_32_ldst_unpriv
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDTRSB_64_ldst_unpriv
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDTRSB_32_ldst_unpriv
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STTRH_32_ldst_unpriv
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDTRH_32_ldst_unpriv
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDTRSH_64_ldst_unpriv
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDTRSH_32_ldst_unpriv
	case (d.size&0x2) == 0x2 && d.V == 0x0 && d.opc == 0x3:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STTR_32_ldst_unpriv
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDTR_32_ldst_unpriv
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDTRSW_64_ldst_unpriv
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STTR_64_ldst_unpriv
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDTR_64_ldst_unpriv
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_immpre(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_immpre
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.opc = (ins >> 22) & 0x3
	d.imm9 = (ins >> 12) & 0x1ff
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x1) == 0x1 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRB_32_ldst_immpre
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRB_32_ldst_immpre
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSB_64_ldst_immpre
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSB_32_ldst_immpre
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_B_ldst_immpre
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_B_ldst_immpre
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x2:
		d.encoding = encoding_STR_Q_ldst_immpre
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x3:
		d.encoding = encoding_LDR_Q_ldst_immpre
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRH_32_ldst_immpre
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRH_32_ldst_immpre
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSH_64_ldst_immpre
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSH_32_ldst_immpre
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_H_ldst_immpre
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_H_ldst_immpre
	case (d.size&0x2) == 0x2 && d.V == 0x0 && d.opc == 0x3:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_32_ldst_immpre
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_32_ldst_immpre
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSW_64_ldst_immpre
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_S_ldst_immpre
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_S_ldst_immpre
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_64_ldst_immpre
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_64_ldst_immpre
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x2:
		err = errUnallocated
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_D_ldst_immpre
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_D_ldst_immpre
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_memop(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_memop
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.A = (ins >> 23) & 0x1
	d.R = (ins >> 22) & 0x1
	d.Rs = (ins >> 16) & 0x1f
	d.o3 = (ins >> 15) & 0x1
	d.opc = (ins >> 12) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.V == 0x0 && d.o3 == 0x1 && (d.opc&0x6) == 0x6:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x0 && d.o3 == 0x1 && d.opc == 0x4:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x2:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x3:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x5:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x2:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x3:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x5:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x2:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x3:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x4:
		err = errUnallocated
	case d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x5:
		err = errUnallocated
	case d.V == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x2:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x3:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x5:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPLB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPAB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x4:
		d.encoding = encoding_LDAPRB_32L_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINALB_32_memop
	case d.size == 0x0 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPALB_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x2:
		err = errUnallocated
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x3:
		err = errUnallocated
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x5:
		err = errUnallocated
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPLH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPAH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x4:
		d.encoding = encoding_LDAPRH_32L_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINALH_32_memop
	case d.size == 0x1 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPALH_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADD_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLR_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEOR_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSET_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAX_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMIN_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAX_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMIN_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWP_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x1:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x2:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x3:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x5:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPA_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x4:
		d.encoding = encoding_LDAPR_32L_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINAL_32_memop
	case d.size == 0x2 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPAL_32_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADD_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLR_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEOR_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSET_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAX_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMIN_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAX_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMIN_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWP_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x2:
		d.encoding = encoding_ST64BV0_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x3:
		d.encoding = encoding_ST64BV_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.Rs == 0x1f && d.o3 == 0x1 && d.opc == 0x1:
		d.encoding = encoding_ST64B_64L_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x0 && d.Rs == 0x1f && d.o3 == 0x1 && d.opc == 0x5:
		d.encoding = encoding_LD64B_64L_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x0 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPA_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x0 && d.o3 == 0x1 && d.opc == 0x4:
		d.encoding = encoding_LDAPR_64L_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x0:
		d.encoding = encoding_LDADDAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDCLRAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDEORAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDSETAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x4:
		d.encoding = encoding_LDSMAXAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x5:
		d.encoding = encoding_LDSMINAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x6:
		d.encoding = encoding_LDUMAXAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x0 && d.opc == 0x7:
		d.encoding = encoding_LDUMINAL_64_memop
	case d.size == 0x3 && d.V == 0x0 && d.A == 0x1 && d.R == 0x1 && d.o3 == 0x1 && d.opc == 0x0:
		d.encoding = encoding_SWPAL_64_memop
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_regoff(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_regoff
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.opc = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.option = (ins >> 13) & 0x7
	d.S = (ins >> 12) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x1) == 0x1 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0 && d.option != 0x3:
		d.encoding = encoding_STRB_32B_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0 && d.option == 0x3:
		d.encoding = encoding_STRB_32BL_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1 && d.option != 0x3:
		d.encoding = encoding_LDRB_32B_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1 && d.option == 0x3:
		d.encoding = encoding_LDRB_32BL_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2 && d.option != 0x3:
		d.encoding = encoding_LDRSB_64B_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2 && d.option == 0x3:
		d.encoding = encoding_LDRSB_64BL_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3 && d.option != 0x3:
		d.encoding = encoding_LDRSB_32B_ldst_regoff
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3 && d.option == 0x3:
		d.encoding = encoding_LDRSB_32BL_ldst_regoff
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x0 && d.option != 0x3:
		d.encoding = encoding_STR_B_ldst_regoff
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x0 && d.option == 0x3:
		d.encoding = encoding_STR_BL_ldst_regoff
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x1 && d.option != 0x3:
		d.encoding = encoding_LDR_B_ldst_regoff
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x1 && d.option == 0x3:
		d.encoding = encoding_LDR_BL_ldst_regoff
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x2:
		d.encoding = encoding_STR_Q_ldst_regoff
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x3:
		d.encoding = encoding_LDR_Q_ldst_regoff
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRH_32_ldst_regoff
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRH_32_ldst_regoff
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSH_64_ldst_regoff
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSH_32_ldst_regoff
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_H_ldst_regoff
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_H_ldst_regoff
	case (d.size&0x2) == 0x2 && d.V == 0x0 && d.opc == 0x3:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_32_ldst_regoff
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_32_ldst_regoff
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSW_64_ldst_regoff
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_S_ldst_regoff
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_S_ldst_regoff
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_64_ldst_regoff
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_64_ldst_regoff
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_PRFM_P_ldst_regoff
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_D_ldst_regoff
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_D_ldst_regoff
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_pac(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_pac
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.M = (ins >> 23) & 0x1
	d.S = (ins >> 22) & 0x1
	d.imm9 = (ins >> 12) & 0x1ff
	d.W = (ins >> 11) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case d.size != 0x3:
		err = errUnallocated
	case d.size == 0x3 && d.V == 0x0 && d.M == 0x0 && d.W == 0x0:
		d.encoding = encoding_LDRAA_64_ldst_pac
	case d.size == 0x3 && d.V == 0x0 && d.M == 0x0 && d.W == 0x1:
		d.encoding = encoding_LDRAA_64W_ldst_pac
	case d.size == 0x3 && d.V == 0x0 && d.M == 0x1 && d.W == 0x0:
		d.encoding = encoding_LDRAB_64_ldst_pac
	case d.size == 0x3 && d.V == 0x0 && d.M == 0x1 && d.W == 0x1:
		d.encoding = encoding_LDRAB_64W_ldst_pac
	case d.size == 0x3 && d.V == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_ldst_ldst_pos(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_ldst_pos
	d.size = (ins >> 30) & 0x3
	d.V = (ins >> 26) & 0x1
	d.opc = (ins >> 22) & 0x3
	d.imm12 = (ins >> 10) & 0xfff
	d.Rn = (ins >> 5) & 0x1f
	d.Rt = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x1) == 0x1 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRB_32_ldst_pos
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRB_32_ldst_pos
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSB_64_ldst_pos
	case d.size == 0x0 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSB_32_ldst_pos
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_B_ldst_pos
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_B_ldst_pos
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x2:
		d.encoding = encoding_STR_Q_ldst_pos
	case d.size == 0x0 && d.V == 0x1 && d.opc == 0x3:
		d.encoding = encoding_LDR_Q_ldst_pos
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STRH_32_ldst_pos
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDRH_32_ldst_pos
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSH_64_ldst_pos
	case d.size == 0x1 && d.V == 0x0 && d.opc == 0x3:
		d.encoding = encoding_LDRSH_32_ldst_pos
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_H_ldst_pos
	case d.size == 0x1 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_H_ldst_pos
	case (d.size&0x2) == 0x2 && d.V == 0x0 && d.opc == 0x3:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.V == 0x1 && (d.opc&0x2) == 0x2:
		err = errUnallocated
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_32_ldst_pos
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_32_ldst_pos
	case d.size == 0x2 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_LDRSW_64_ldst_pos
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_S_ldst_pos
	case d.size == 0x2 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_S_ldst_pos
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x0:
		d.encoding = encoding_STR_64_ldst_pos
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x1:
		d.encoding = encoding_LDR_64_ldst_pos
	case d.size == 0x3 && d.V == 0x0 && d.opc == 0x2:
		d.encoding = encoding_PRFM_P_ldst_pos
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x0:
		d.encoding = encoding_STR_D_ldst_pos
	case d.size == 0x3 && d.V == 0x1 && d.opc == 0x1:
		d.encoding = encoding_LDR_D_ldst_pos
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 30) & 0x1
	op1 := (ins >> 28) & 0x1
	op2 := (ins >> 21) & 0xf
	op3 := (ins >> 10) & 0x3f

	switch {
	case op0 == 0x0 && op1 == 0x1 && op2 == 0x6:
		err = decode_dpreg_dp_2src(ins, d)
	case op0 == 0x1 && op1 == 0x1 && op2 == 0x6:
		err = decode_dpreg_dp_1src(ins, d)
	case op1 == 0x0 && (op2&0x8) == 0x0:
		err = decode_dpreg_log_shift(ins, d)
	case op1 == 0x0 && (op2&0x9) == 0x8:
		err = decode_dpreg_addsub_shift(ins, d)
	case op1 == 0x0 && (op2&0x9) == 0x9:
		err = decode_dpreg_addsub_ext(ins, d)
	case op1 == 0x1 && op2 == 0x0 && op3 == 0x0:
		err = decode_dpreg_addsub_carry(ins, d)
	case op1 == 0x1 && op2 == 0x0 && (op3&0x1f) == 0x1:
		err = decode_dpreg_rmif(ins, d)
	case op1 == 0x1 && op2 == 0x0 && (op3&0xf) == 0x2:
		err = decode_dpreg_setf(ins, d)
	case op1 == 0x1 && op2 == 0x2 && (op3&0x2) == 0x0:
		err = decode_dpreg_condcmp_reg(ins, d)
	case op1 == 0x1 && op2 == 0x2 && (op3&0x2) == 0x2:
		err = decode_dpreg_condcmp_imm(ins, d)
	case op1 == 0x1 && op2 == 0x4:
		err = decode_dpreg_condsel(ins, d)
	case op1 == 0x1 && (op2&0x8) == 0x8:
		err = decode_dpreg_dp_3src(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_dp_2src(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_dp_2src
	d.sf = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x1:
		err = errUnallocated
	case (d.opcode & 0x38) == 0x18:
		err = errUnallocated
	case (d.opcode & 0x20) == 0x20:
		err = errUnallocated
	case d.S == 0x0 && (d.opcode&0x3e) == 0x6:
		err = errUnallocated
	case d.S == 0x0 && d.opcode == 0xd:
		err = errUnallocated
	case d.S == 0x0 && (d.opcode&0x3e) == 0xe:
		err = errUnallocated
	case d.S == 0x1 && (d.opcode&0x3e) == 0x2:
		err = errUnallocated
	case d.S == 0x1 && (d.opcode&0x3c) == 0x4:
		err = errUnallocated
	case d.S == 0x1 && (d.opcode&0x38) == 0x8:
		err = errUnallocated
	case d.S == 0x1 && (d.opcode&0x30) == 0x10:
		err = errUnallocated
	case d.sf == 0x0 && d.opcode == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_UDIV_32_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SDIV_32_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && (d.opcode&0x3e) == 0x4:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_LSLV_32_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_LSRV_32_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_ASRV_32_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_RORV_32_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0xc:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && (d.opcode&0x3b) == 0x13:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x10:
		d.encoding = encoding_CRC32B_32C_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x11:
		d.encoding = encoding_CRC32H_32C_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x12:
		d.encoding = encoding_CRC32W_32C_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x14:
		d.encoding = encoding_CRC32CB_32C_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x15:
		d.encoding = encoding_CRC32CH_32C_dp_2src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode == 0x16:
		d.encoding = encoding_CRC32CW_32C_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SUBP_64S_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_UDIV_64_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SDIV_64_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_IRG_64I_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_GMI_64G_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_LSLV_64_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_LSRV_64_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_ASRV_64_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_RORV_64_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_PACGA_64P_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && (d.opcode&0x39) == 0x10:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && (d.opcode&0x3a) == 0x10:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x13:
		d.encoding = encoding_CRC32X_64C_dp_2src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode == 0x17:
		d.encoding = encoding_CRC32CX_64C_dp_2src
	case d.sf == 0x1 && d.S == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_SUBPS_64S_dp_2src
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_dp_1src(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_dp_1src
	d.sf = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.opcode2 = (ins >> 16) & 0x1f
	d.opcode = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x20) == 0x20:
		err = errUnallocated
	case (d.opcode2 & 0x2) == 0x2:
		err = errUnallocated
	case (d.opcode2 & 0x4) == 0x4:
		err = errUnallocated
	case (d.opcode2 & 0x8) == 0x8:
		err = errUnallocated
	case (d.opcode2 & 0x10) == 0x10:
		err = errUnallocated
	case d.S == 0x0 && d.opcode2 == 0x0 && (d.opcode&0x3e) == 0x6:
		err = errUnallocated
	case d.S == 0x0 && d.opcode2 == 0x0 && (d.opcode&0x38) == 0x8:
		err = errUnallocated
	case d.S == 0x0 && d.opcode2 == 0x0 && (d.opcode&0x30) == 0x10:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.opcode2 == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_RBIT_32_dp_1src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_REV16_32_dp_1src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_REV_32_dp_1src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_CLZ_32_dp_1src
	case d.sf == 0x0 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_CLS_32_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_RBIT_64_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_REV16_64_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_REV32_64_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_REV_64_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_CLZ_64_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_CLS_64_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_PACIA_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_PACIB_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_PACDA_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_PACDB_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_AUTIA_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_AUTIB_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_AUTDA_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_AUTDB_64P_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x8 && d.Rn == 0x1f:
		d.encoding = encoding_PACIZA_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x9 && d.Rn == 0x1f:
		d.encoding = encoding_PACIZB_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0xa && d.Rn == 0x1f:
		d.encoding = encoding_PACDZA_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0xb && d.Rn == 0x1f:
		d.encoding = encoding_PACDZB_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0xc && d.Rn == 0x1f:
		d.encoding = encoding_AUTIZA_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0xd && d.Rn == 0x1f:
		d.encoding = encoding_AUTIZB_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0xe && d.Rn == 0x1f:
		d.encoding = encoding_AUTDZA_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0xf && d.Rn == 0x1f:
		d.encoding = encoding_AUTDZB_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x10 && d.Rn == 0x1f:
		d.encoding = encoding_XPACI_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && d.opcode == 0x11 && d.Rn == 0x1f:
		d.encoding = encoding_XPACD_64Z_dp_1src
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && (d.opcode&0x3e) == 0x12:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && (d.opcode&0x3c) == 0x14:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.opcode2 == 0x1 && (d.opcode&0x38) == 0x18:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_log_shift(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_log_shift
	d.sf = (ins >> 31) & 0x1
	d.opc = (ins >> 29) & 0x3
	d.shift = (ins >> 22) & 0x3
	d.N = (ins >> 21) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.sf == 0x0 && (d.imm6&0x20) == 0x20:
		err = errUnallocated
	case d.sf == 0x0 && d.opc == 0x0 && d.N == 0x0:
		d.encoding = encoding_AND_32_log_shift
	case d.sf == 0x0 && d.opc == 0x0 && d.N == 0x1:
		d.encoding = encoding_BIC_32_log_shift
	case d.sf == 0x0 && d.opc == 0x1 && d.N == 0x0:
		d.encoding = encoding_ORR_32_log_shift
	case d.sf == 0x0 && d.opc == 0x1 && d.N == 0x1:
		d.encoding = encoding_ORN_32_log_shift
	case d.sf == 0x0 && d.opc == 0x2 && d.N == 0x0:
		d.encoding = encoding_EOR_32_log_shift
	case d.sf == 0x0 && d.opc == 0x2 && d.N == 0x1:
		d.encoding = encoding_EON_32_log_shift
	case d.sf == 0x0 && d.opc == 0x3 && d.N == 0x0:
		d.encoding = encoding_ANDS_32_log_shift
	case d.sf == 0x0 && d.opc == 0x3 && d.N == 0x1:
		d.encoding = encoding_BICS_32_log_shift
	case d.sf == 0x1 && d.opc == 0x0 && d.N == 0x0:
		d.encoding = encoding_AND_64_log_shift
	case d.sf == 0x1 && d.opc == 0x0 && d.N == 0x1:
		d.encoding = encoding_BIC_64_log_shift
	case d.sf == 0x1 && d.opc == 0x1 && d.N == 0x0:
		d.encoding = encoding_ORR_64_log_shift
	case d.sf == 0x1 && d.opc == 0x1 && d.N == 0x1:
		d.encoding = encoding_ORN_64_log_shift
	case d.sf == 0x1 && d.opc == 0x2 && d.N == 0x0:
		d.encoding = encoding_EOR_64_log_shift
	case d.sf == 0x1 && d.opc == 0x2 && d.N == 0x1:
		d.encoding = encoding_EON_64_log_shift
	case d.sf == 0x1 && d.opc == 0x3 && d.N == 0x0:
		d.encoding = encoding_ANDS_64_log_shift
	case d.sf == 0x1 && d.opc == 0x3 && d.N == 0x1:
		d.encoding = encoding_BICS_64_log_shift
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_addsub_shift(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_addsub_shift
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.shift = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.shift == 0x3:
		err = errUnallocated
	case d.sf == 0x0 && (d.imm6&0x20) == 0x20:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_ADD_32_addsub_shift
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_ADDS_32_addsub_shift
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x0:
		d.encoding = encoding_SUB_32_addsub_shift
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x1:
		d.encoding = encoding_SUBS_32_addsub_shift
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_ADD_64_addsub_shift
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_ADDS_64_addsub_shift
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0:
		d.encoding = encoding_SUB_64_addsub_shift
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x1:
		d.encoding = encoding_SUBS_64_addsub_shift
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_addsub_ext(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_addsub_ext
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.opt = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.option = (ins >> 13) & 0x7
	d.imm3 = (ins >> 10) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.imm3 & 0x5) == 0x5:
		err = errUnallocated
	case (d.imm3 & 0x6) == 0x6:
		err = errUnallocated
	case (d.opt & 0x1) == 0x1:
		err = errUnallocated
	case (d.opt & 0x2) == 0x2:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0 && d.opt == 0x0:
		d.encoding = encoding_ADD_32_addsub_ext
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.opt == 0x0:
		d.encoding = encoding_ADDS_32S_addsub_ext
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x0 && d.opt == 0x0:
		d.encoding = encoding_SUB_32_addsub_ext
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x1 && d.opt == 0x0:
		d.encoding = encoding_SUBS_32S_addsub_ext
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0 && d.opt == 0x0:
		d.encoding = encoding_ADD_64_addsub_ext
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1 && d.opt == 0x0:
		d.encoding = encoding_ADDS_64S_addsub_ext
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0 && d.opt == 0x0:
		d.encoding = encoding_SUB_64_addsub_ext
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x1 && d.opt == 0x0:
		d.encoding = encoding_SUBS_64S_addsub_ext
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_addsub_carry(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_addsub_carry
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_ADC_32_addsub_carry
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_ADCS_32_addsub_carry
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x0:
		d.encoding = encoding_SBC_32_addsub_carry
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x1:
		d.encoding = encoding_SBCS_32_addsub_carry
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0:
		d.encoding = encoding_ADC_64_addsub_carry
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1:
		d.encoding = encoding_ADCS_64_addsub_carry
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0:
		d.encoding = encoding_SBC_64_addsub_carry
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x1:
		d.encoding = encoding_SBCS_64_addsub_carry
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_rmif(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_rmif
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.imm6 = (ins >> 15) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.o2 = (ins >> 4) & 0x1
	d.mask = (ins >> 0) & 0xf

	switch {
	case d.sf == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0:
		d.encoding = encoding_RMIF_only_rmif
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1 && d.o2 == 0x1:
		err = errUnallocated
	case d.sf == 0x1 && d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_setf(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_setf
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.opcode2 = (ins >> 15) & 0x3f
	d.sz = (ins >> 14) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.o3 = (ins >> 4) & 0x1
	d.mask = (ins >> 0) & 0xf

	switch {
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.opcode2 != 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.opcode2 == 0x0 && d.o3 == 0x0 && d.mask != 0xd:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.opcode2 == 0x0 && d.o3 == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.opcode2 == 0x0 && d.sz == 0x0 && d.o3 == 0x0 && d.mask == 0xd:
		d.encoding = encoding_SETF8_only_setf
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.opcode2 == 0x0 && d.sz == 0x1 && d.o3 == 0x0 && d.mask == 0xd:
		d.encoding = encoding_SETF16_only_setf
	case d.sf == 0x0 && d.op == 0x1:
		err = errUnallocated
	case d.sf == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_condcmp_reg(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_condcmp_reg
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.cond = (ins >> 12) & 0xf
	d.o2 = (ins >> 10) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.o3 = (ins >> 4) & 0x1
	d.nzcv = (ins >> 0) & 0xf

	switch {
	case d.o3 == 0x1:
		err = errUnallocated
	case d.o2 == 0x1:
		err = errUnallocated
	case d.S == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMN_32_condcmp_reg
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMP_32_condcmp_reg
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMN_64_condcmp_reg
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMP_64_condcmp_reg
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_condcmp_imm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_condcmp_imm
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.imm5 = (ins >> 16) & 0x1f
	d.cond = (ins >> 12) & 0xf
	d.o2 = (ins >> 10) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.o3 = (ins >> 4) & 0x1
	d.nzcv = (ins >> 0) & 0xf

	switch {
	case d.o3 == 0x1:
		err = errUnallocated
	case d.o2 == 0x1:
		err = errUnallocated
	case d.S == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMN_32_condcmp_imm
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMP_32_condcmp_imm
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMN_64_condcmp_imm
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x1 && d.o2 == 0x0 && d.o3 == 0x0:
		d.encoding = encoding_CCMP_64_condcmp_imm
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_condsel(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_condsel
	d.sf = (ins >> 31) & 0x1
	d.op = (ins >> 30) & 0x1
	d.S = (ins >> 29) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.cond = (ins >> 12) & 0xf
	d.op2 = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.op2 & 0x2) == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_CSEL_32_condsel
	case d.sf == 0x0 && d.op == 0x0 && d.S == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_CSINC_32_condsel
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_CSINV_32_condsel
	case d.sf == 0x0 && d.op == 0x1 && d.S == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_CSNEG_32_condsel
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_CSEL_64_condsel
	case d.sf == 0x1 && d.op == 0x0 && d.S == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_CSINC_64_condsel
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0 && d.op2 == 0x0:
		d.encoding = encoding_CSINV_64_condsel
	case d.sf == 0x1 && d.op == 0x1 && d.S == 0x0 && d.op2 == 0x1:
		d.encoding = encoding_CSNEG_64_condsel
	default:
		err = errUnmatched
	}
	return
}

func decode_dpreg_dp_3src(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_dp_3src
	d.sf = (ins >> 31) & 0x1
	d.op54 = (ins >> 29) & 0x3
	d.op31 = (ins >> 21) & 0x7
	d.Rm = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Ra = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op54 == 0x0 && d.op31 == 0x2 && d.o0 == 0x1:
		err = errUnallocated
	case d.op54 == 0x0 && d.op31 == 0x3:
		err = errUnallocated
	case d.op54 == 0x0 && d.op31 == 0x4:
		err = errUnallocated
	case d.op54 == 0x0 && d.op31 == 0x6 && d.o0 == 0x1:
		err = errUnallocated
	case d.op54 == 0x0 && d.op31 == 0x7:
		err = errUnallocated
	case d.op54 == 0x1:
		err = errUnallocated
	case (d.op54 & 0x2) == 0x2:
		err = errUnallocated
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_MADD_32A_dp_3src
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_MSUB_32A_dp_3src
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x1 && d.o0 == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x1 && d.o0 == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x2 && d.o0 == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x5 && d.o0 == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x5 && d.o0 == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.op54 == 0x0 && d.op31 == 0x6 && d.o0 == 0x0:
		err = errUnallocated
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_MADD_64A_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_MSUB_64A_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_SMADDL_64WA_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_SMSUBL_64WA_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x2 && d.o0 == 0x0:
		d.encoding = encoding_SMULH_64_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x5 && d.o0 == 0x0:
		d.encoding = encoding_UMADDL_64WA_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x5 && d.o0 == 0x1:
		d.encoding = encoding_UMSUBL_64WA_dp_3src
	case d.sf == 0x1 && d.op54 == 0x0 && d.op31 == 0x6 && d.o0 == 0x0:
		d.encoding = encoding_UMULH_64_dp_3src
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp(ins uint32, d *decoded) (err error) {
	op0 := (ins >> 28) & 0xf
	op1 := (ins >> 23) & 0x3
	op2 := (ins >> 19) & 0xf
	op3 := (ins >> 10) & 0x1ff

	switch {
	case op0 == 0x0 && (op1&0x2) == 0x0 && (op2&0x7) == 0x5 && (op3&0x183) == 0x2:
		err = errUnallocated
	case op0 == 0x2 && (op1&0x2) == 0x0 && (op2&0x7) == 0x5 && (op3&0x183) == 0x2:
		err = errUnallocated
	case op0 == 0x4 && (op1&0x2) == 0x0 && (op2&0x7) == 0x5 && (op3&0x183) == 0x2:
		err = decode_simddp_cryptoaes(ins, d)
	case op0 == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x23) == 0x0:
		err = decode_simddp_cryptosha3(ins, d)
	case op0 == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x23) == 0x2:
		err = errUnallocated
	case op0 == 0x5 && (op1&0x2) == 0x0 && (op2&0x7) == 0x5 && (op3&0x183) == 0x2:
		err = decode_simddp_cryptosha2(ins, d)
	case op0 == 0x6 && (op1&0x2) == 0x0 && (op2&0x7) == 0x5 && (op3&0x183) == 0x2:
		err = errUnallocated
	case op0 == 0x7 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x21) == 0x0:
		err = errUnallocated
	case op0 == 0x7 && (op1&0x2) == 0x0 && (op2&0x7) == 0x5 && (op3&0x183) == 0x2:
		err = errUnallocated
	case (op0&0xd) == 0x5 && op1 == 0x0 && (op2&0xc) == 0x0 && (op3&0x21) == 0x1:
		err = decode_simddp_asisdone(ins, d)
	case (op0&0xd) == 0x5 && op1 == 0x1 && (op2&0xc) == 0x0 && (op3&0x21) == 0x1:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && op2 == 0x7 && (op3&0x183) == 0x2:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0xc) == 0x8 && (op3&0x31) == 0x1:
		err = decode_simddp_asisdsamefp16(ins, d)
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0xc) == 0x8 && (op3&0x31) == 0x11:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && op2 == 0xf && (op3&0x183) == 0x2:
		err = decode_simddp_asisdmiscfp16(ins, d)
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x21) == 0x20:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x21) == 0x21:
		err = decode_simddp_asisdsame2(ins, d)
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x7) == 0x4 && (op3&0x183) == 0x2:
		err = decode_simddp_asisdmisc(ins, d)
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x7) == 0x6 && (op3&0x183) == 0x2:
		err = decode_simddp_asisdpair(ins, d)
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x103) == 0x102:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x83) == 0x82:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x3) == 0x0:
		err = decode_simddp_asisddiff(ins, d)
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x1) == 0x1:
		err = decode_simddp_asisdsame(ins, d)
	case (op0&0xd) == 0x5 && op1 == 0x2 && (op3&0x1) == 0x1:
		err = decode_simddp_asisdshf(ins, d)
	case (op0&0xd) == 0x5 && op1 == 0x3 && (op3&0x1) == 0x1:
		err = errUnallocated
	case (op0&0xd) == 0x5 && (op1&0x2) == 0x2 && (op3&0x1) == 0x0:
		err = decode_simddp_asisdelem(ins, d)
	case (op0&0xb) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x23) == 0x0:
		err = decode_simddp_asimdtbl(ins, d)
	case (op0&0xb) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x23) == 0x2:
		err = decode_simddp_asimdperm(ins, d)
	case (op0&0xb) == 0x2 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x21) == 0x0:
		err = decode_simddp_asimdext(ins, d)
	case (op0&0x9) == 0x0 && op1 == 0x0 && (op2&0xc) == 0x0 && (op3&0x21) == 0x1:
		err = decode_simddp_asimdins(ins, d)
	case (op0&0x9) == 0x0 && op1 == 0x1 && (op2&0xc) == 0x0 && (op3&0x21) == 0x1:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && op2 == 0x7 && (op3&0x183) == 0x2:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0xc) == 0x8 && (op3&0x31) == 0x1:
		err = decode_simddp_asimdsamefp16(ins, d)
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0xc) == 0x8 && (op3&0x31) == 0x11:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && op2 == 0xf && (op3&0x183) == 0x2:
		err = decode_simddp_asimdmiscfp16(ins, d)
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x21) == 0x20:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0 && (op3&0x21) == 0x21:
		err = decode_simddp_asimdsame2(ins, d)
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x7) == 0x4 && (op3&0x183) == 0x2:
		err = decode_simddp_asimdmisc(ins, d)
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x7) == 0x6 && (op3&0x183) == 0x2:
		err = decode_simddp_asimdall(ins, d)
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x103) == 0x102:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x83) == 0x82:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x3) == 0x0:
		err = decode_simddp_asimddiff(ins, d)
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x1) == 0x1:
		err = decode_simddp_asimdsame(ins, d)
	case (op0&0x9) == 0x0 && op1 == 0x2 && op2 == 0x0 && (op3&0x1) == 0x1:
		err = decode_simddp_asimdimm(ins, d)
	case (op0&0x9) == 0x0 && op1 == 0x2 && op2 != 0x0 && (op3&0x1) == 0x1:
		err = decode_simddp_asimdshf(ins, d)
	case (op0&0x9) == 0x0 && op1 == 0x3 && (op3&0x1) == 0x1:
		err = errUnallocated
	case (op0&0x9) == 0x0 && (op1&0x2) == 0x2 && (op3&0x1) == 0x0:
		err = decode_simddp_asimdelem(ins, d)
	case op0 == 0xc && op1 == 0x0 && (op2&0xc) == 0x8 && (op3&0x30) == 0x20:
		err = decode_simddp_crypto3_imm2(ins, d)
	case op0 == 0xc && op1 == 0x0 && (op2&0xc) == 0xc && (op3&0x2c) == 0x20:
		err = decode_simddp_cryptosha512_3(ins, d)
	case op0 == 0xc && op1 == 0x0 && (op3&0x20) == 0x0:
		err = decode_simddp_crypto4(ins, d)
	case op0 == 0xc && op1 == 0x1 && (op2&0xc) == 0x0:
		err = decode_simddp_crypto3_imm6(ins, d)
	case op0 == 0xc && op1 == 0x1 && op2 == 0x8 && (op3&0x1fc) == 0x20:
		err = decode_simddp_cryptosha512_2(ins, d)
	case (op0&0x9) == 0x8 && (op1&0x2) == 0x2:
		err = errUnallocated
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x0:
		err = decode_simddp_float2fix(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x3f) == 0x0:
		err = decode_simddp_float2int(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x1f) == 0x10:
		err = decode_simddp_floatdp1(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0xf) == 0x8:
		err = decode_simddp_floatcmp(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x7) == 0x4:
		err = decode_simddp_floatimm(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x3) == 0x1:
		err = decode_simddp_floatccmp(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x3) == 0x2:
		err = decode_simddp_floatdp2(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x0 && (op2&0x4) == 0x4 && (op3&0x3) == 0x3:
		err = decode_simddp_floatsel(ins, d)
	case (op0&0x5) == 0x1 && (op1&0x2) == 0x2:
		err = decode_simddp_floatdp3(ins, d)
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_cryptoaes(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_cryptoaes
	d.size = (ins >> 22) & 0x3
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x8) == 0x8:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x0:
		err = errUnallocated
	case (d.opcode & 0x10) == 0x10:
		err = errUnallocated
	case (d.size & 0x1) == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_AESE_B_cryptoaes
	case d.size == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_AESD_B_cryptoaes
	case d.size == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_AESMC_B_cryptoaes
	case d.size == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_AESIMC_B_cryptoaes
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_cryptosha3(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_cryptosha3
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 12) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x7:
		err = errUnallocated
	case (d.size & 0x1) == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SHA1C_QSV_cryptosha3
	case d.size == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_SHA1P_QSV_cryptosha3
	case d.size == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SHA1M_QSV_cryptosha3
	case d.size == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SHA1SU0_VVV_cryptosha3
	case d.size == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_SHA256H_QQV_cryptosha3
	case d.size == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_SHA256H2_QQV_cryptosha3
	case d.size == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_SHA256SU1_VVV_cryptosha3
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_cryptosha2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_cryptosha2
	d.size = (ins >> 22) & 0x3
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x4) == 0x4:
		err = errUnallocated
	case (d.opcode & 0x8) == 0x8:
		err = errUnallocated
	case (d.opcode & 0x10) == 0x10:
		err = errUnallocated
	case (d.size & 0x1) == 0x1:
		err = errUnallocated
	case d.size == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SHA1H_SS_cryptosha2
	case d.size == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_SHA1SU1_VV_cryptosha2
	case d.size == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SHA256SU0_VV_cryptosha2
	case d.size == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case (d.size & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdone(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdone
	d.op = (ins >> 29) & 0x1
	d.imm5 = (ins >> 16) & 0x1f
	d.imm4 = (ins >> 11) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && (d.imm4&0x1) == 0x1:
		err = errUnallocated
	case d.op == 0x0 && (d.imm4&0x2) == 0x2:
		err = errUnallocated
	case d.op == 0x0 && (d.imm4&0x4) == 0x4:
		err = errUnallocated
	case d.op == 0x0 && d.imm4 == 0x0:
		d.encoding = encoding_DUP_asisdone_only
	case d.op == 0x0 && (d.imm4&0x8) == 0x8:
		err = errUnallocated
	case d.op == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdsamefp16(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdsamefp16
	d.U = (ins >> 29) & 0x1
	d.a = (ins >> 23) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 11) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x6:
		err = errUnallocated
	case d.a == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_FMULX_asisdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCMEQ_asisdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FRECPS_asisdsamefp16_only
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x4:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_FRSQRTS_asisdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCMGE_asisdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FACGE_asisdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x7:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_FABD_asisdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_FCMGT_asisdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_FACGT_asisdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdmiscfp16(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdmiscfp16
	d.U = (ins >> 29) & 0x1
	d.a = (ins >> 23) & 0x1
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x18) == 0x0:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x8:
		err = errUnallocated
	case (d.opcode & 0x18) == 0x10:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x18:
		err = errUnallocated
	case d.opcode == 0x1e:
		err = errUnallocated
	case d.a == 0x0 && (d.opcode&0x1c) == 0xc:
		err = errUnallocated
	case d.a == 0x0 && d.opcode == 0x1f:
		err = errUnallocated
	case d.a == 0x1 && d.opcode == 0xf:
		err = errUnallocated
	case d.a == 0x1 && d.opcode == 0x1c:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNS_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMS_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAS_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_SCVTF_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_FCMGT_asisdmiscfp16_FZ
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_FCMEQ_asisdmiscfp16_FZ
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xe:
		d.encoding = encoding_FCMLT_asisdmiscfp16_FZ
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPS_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZS_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1d:
		d.encoding = encoding_FRECPE_asisdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1f:
		d.encoding = encoding_FRECPX_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNU_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMU_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAU_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_UCVTF_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_FCMGE_asisdmiscfp16_FZ
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_FCMLE_asisdmiscfp16_FZ
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xe:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPU_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZU_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1d:
		d.encoding = encoding_FRSQRTE_asisdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdsame2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdsame2
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 11) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0xe) == 0x2:
		err = errUnallocated
	case (d.opcode & 0xc) == 0x4:
		err = errUnallocated
	case (d.opcode & 0x8) == 0x8:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x0:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_SQRDMLAH_asisdsame2_only
	case d.U == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_SQRDMLSH_asisdsame2_only
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdmisc(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdmisc
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x1e) == 0x0:
		err = errUnallocated
	case d.opcode == 0x2:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x4:
		err = errUnallocated
	case d.opcode == 0x6:
		err = errUnallocated
	case d.opcode == 0xf:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x10:
		err = errUnallocated
	case d.opcode == 0x13:
		err = errUnallocated
	case d.opcode == 0x15:
		err = errUnallocated
	case d.opcode == 0x17:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x18:
		err = errUnallocated
	case d.opcode == 0x1e:
		err = errUnallocated
	case (d.size&0x2) == 0x0 && (d.opcode&0x1c) == 0xc:
		err = errUnallocated
	case (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opcode == 0x16:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SUQADD_asisdmisc_R
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_SQABS_asisdmisc_R
	case d.U == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_CMGT_asisdmisc_Z
	case d.U == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_CMEQ_asisdmisc_Z
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_CMLT_asisdmisc_Z
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_ABS_asisdmisc_R
	case d.U == 0x0 && d.opcode == 0x12:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x14:
		d.encoding = encoding_SQXTN_asisdmisc_N
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x16:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNS_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMS_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAS_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_SCVTF_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FCMGT_asisdmisc_FZ
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xd:
		d.encoding = encoding_FCMEQ_asisdmisc_FZ
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xe:
		d.encoding = encoding_FCMLT_asisdmisc_FZ
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPS_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZS_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FRECPE_asisdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		d.encoding = encoding_FRECPX_asisdmisc_R
	case d.U == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_USQADD_asisdmisc_R
	case d.U == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_SQNEG_asisdmisc_R
	case d.U == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_CMGE_asisdmisc_Z
	case d.U == 0x1 && d.opcode == 0x9:
		d.encoding = encoding_CMLE_asisdmisc_Z
	case d.U == 0x1 && d.opcode == 0xa:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xb:
		d.encoding = encoding_NEG_asisdmisc_R
	case d.U == 0x1 && d.opcode == 0x12:
		d.encoding = encoding_SQXTUN_asisdmisc_N
	case d.U == 0x1 && d.opcode == 0x14:
		d.encoding = encoding_UQXTN_asisdmisc_N
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x16:
		d.encoding = encoding_FCVTXN_asisdmisc_N
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNU_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMU_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAU_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_UCVTF_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FCMGE_asisdmisc_FZ
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xd:
		d.encoding = encoding_FCMLE_asisdmisc_FZ
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xe:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPU_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZU_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FRSQRTE_asisdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdpair(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdpair
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x18) == 0x0:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x8:
		err = errUnallocated
	case d.opcode == 0xe:
		err = errUnallocated
	case (d.opcode & 0x18) == 0x10:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x18:
		err = errUnallocated
	case d.opcode == 0x1a:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x1c:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opcode == 0xd:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_ADDP_asisdpair_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_FMAXNMP_asisdpair_only_H
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_FADDP_asisdpair_only_H
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_FMAXP_asisdpair_only_H
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FMINNMP_asisdpair_only_H
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_FMINP_asisdpair_only_H
	case d.U == 0x1 && d.opcode == 0x1b:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_FMAXNMP_asisdpair_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_FADDP_asisdpair_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_FMAXP_asisdpair_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FMINNMP_asisdpair_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_FMINP_asisdpair_only_SD
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisddiff(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisddiff
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 12) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0xc) == 0x0:
		err = errUnallocated
	case (d.opcode & 0xc) == 0x4:
		err = errUnallocated
	case d.opcode == 0x8:
		err = errUnallocated
	case d.opcode == 0xa:
		err = errUnallocated
	case d.opcode == 0xc:
		err = errUnallocated
	case (d.opcode & 0xe) == 0xe:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_SQDMLAL_asisddiff_only
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_SQDMLSL_asisddiff_only
	case d.U == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_SQDMULL_asisddiff_only
	case d.U == 0x1 && d.opcode == 0x9:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xb:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xd:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdsame(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdsame
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 11) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x0:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x2:
		err = errUnallocated
	case d.opcode == 0x4:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0xc:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x12:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_SQADD_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_SQSUB_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_CMGT_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_CMGE_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_SSHL_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_SQSHL_asisdsame_only
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SRSHL_asisdsame_only
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_SQRSHL_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x10:
		d.encoding = encoding_ADD_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x11:
		d.encoding = encoding_CMTST_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x14:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x15:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x16:
		d.encoding = encoding_SQDMULH_asisdsame_only
	case d.U == 0x0 && d.opcode == 0x17:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x18:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x19:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FMULX_asisdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCMEQ_asisdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1e:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FRECPS_asisdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x18:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x19:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1e:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		d.encoding = encoding_FRSQRTS_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_UQADD_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_UQSUB_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_CMHI_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_CMHS_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_USHL_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x9:
		d.encoding = encoding_UQSHL_asisdsame_only
	case d.U == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_URSHL_asisdsame_only
	case d.U == 0x1 && d.opcode == 0xb:
		d.encoding = encoding_UQRSHL_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x10:
		d.encoding = encoding_SUB_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x11:
		d.encoding = encoding_CMEQ_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x14:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x15:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x16:
		d.encoding = encoding_SQRDMULH_asisdsame_only
	case d.U == 0x1 && d.opcode == 0x17:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x18:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x19:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCMGE_asisdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_FACGE_asisdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1e:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x18:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x19:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FABD_asisdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		d.encoding = encoding_FCMGT_asisdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FACGT_asisdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1e:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdshf(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdshf
	d.U = (ins >> 29) & 0x1
	d.immh = (ins >> 19) & 0xf
	d.immb = (ins >> 16) & 0x7
	d.opcode = (ins >> 11) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.immh != 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0x7:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0x9:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0xb:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0xd:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0xf:
		err = errUnallocated
	case d.immh != 0x0 && (d.opcode&0x1c) == 0x14:
		err = errUnallocated
	case d.immh != 0x0 && (d.opcode&0x1c) == 0x18:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0x1d:
		err = errUnallocated
	case d.immh != 0x0 && d.opcode == 0x1e:
		err = errUnallocated
	case d.immh == 0x0:
		err = errUnallocated
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SSHR_asisdshf_R
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SSRA_asisdshf_R
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x4:
		d.encoding = encoding_SRSHR_asisdshf_R
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x6:
		d.encoding = encoding_SRSRA_asisdshf_R
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x8:
		err = errUnallocated
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SHL_asisdshf_R
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0xc:
		err = errUnallocated
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0xe:
		d.encoding = encoding_SQSHL_asisdshf_R
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x10:
		err = errUnallocated
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x11:
		err = errUnallocated
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x12:
		d.encoding = encoding_SQSHRN_asisdshf_N
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x13:
		d.encoding = encoding_SQRSHRN_asisdshf_N
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_SCVTF_asisdshf_C
	case d.U == 0x0 && d.immh != 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FCVTZS_asisdshf_C
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x0:
		d.encoding = encoding_USHR_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x2:
		d.encoding = encoding_USRA_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x4:
		d.encoding = encoding_URSHR_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x6:
		d.encoding = encoding_URSRA_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x8:
		d.encoding = encoding_SRI_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SLI_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0xc:
		d.encoding = encoding_SQSHLU_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0xe:
		d.encoding = encoding_UQSHL_asisdshf_R
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x10:
		d.encoding = encoding_SQSHRUN_asisdshf_N
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x11:
		d.encoding = encoding_SQRSHRUN_asisdshf_N
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x12:
		d.encoding = encoding_UQSHRN_asisdshf_N
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x13:
		d.encoding = encoding_UQRSHRN_asisdshf_N
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_UCVTF_asisdshf_C
	case d.U == 0x1 && d.immh != 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FCVTZU_asisdshf_C
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asisdelem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asisdelem
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.L = (ins >> 21) & 0x1
	d.M = (ins >> 20) & 0x1
	d.Rm = (ins >> 16) & 0xf
	d.opcode = (ins >> 12) & 0xf
	d.H = (ins >> 11) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x0:
		err = errUnallocated
	case d.opcode == 0x2:
		err = errUnallocated
	case d.opcode == 0x4:
		err = errUnallocated
	case d.opcode == 0x6:
		err = errUnallocated
	case d.opcode == 0x8:
		err = errUnallocated
	case d.opcode == 0xa:
		err = errUnallocated
	case d.opcode == 0xe:
		err = errUnallocated
	case d.size == 0x1 && d.opcode == 0x1:
		err = errUnallocated
	case d.size == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.size == 0x1 && d.opcode == 0x9:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SQDMLAL_asisdelem_L
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_SQDMLSL_asisdelem_L
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_SQDMULL_asisdelem_L
	case d.U == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_SQDMULH_asisdelem_R
	case d.U == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_SQRDMULH_asisdelem_R
	case d.U == 0x0 && d.opcode == 0xf:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FMLA_asisdelem_RH_H
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FMLS_asisdelem_RH_H
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_FMUL_asisdelem_RH_H
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FMLA_asisdelem_R_SD
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x5:
		d.encoding = encoding_FMLS_asisdelem_R_SD
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x9:
		d.encoding = encoding_FMUL_asisdelem_R_SD
	case d.U == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x7:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xb:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xc:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_SQRDMLAH_asisdelem_R
	case d.U == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_SQRDMLSH_asisdelem_R
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_FMULX_asisdelem_RH_H
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x9:
		d.encoding = encoding_FMULX_asisdelem_R_SD
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdtbl(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdtbl
	d.Q = (ins >> 30) & 0x1
	d.op2 = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.len = (ins >> 13) & 0x3
	d.op = (ins >> 12) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.op2 & 0x1) == 0x1:
		err = errUnallocated
	case d.op2 == 0x0 && d.len == 0x0 && d.op == 0x0:
		d.encoding = encoding_TBL_asimdtbl_L1_1
	case d.op2 == 0x0 && d.len == 0x0 && d.op == 0x1:
		d.encoding = encoding_TBX_asimdtbl_L1_1
	case d.op2 == 0x0 && d.len == 0x1 && d.op == 0x0:
		d.encoding = encoding_TBL_asimdtbl_L2_2
	case d.op2 == 0x0 && d.len == 0x1 && d.op == 0x1:
		d.encoding = encoding_TBX_asimdtbl_L2_2
	case d.op2 == 0x0 && d.len == 0x2 && d.op == 0x0:
		d.encoding = encoding_TBL_asimdtbl_L3_3
	case d.op2 == 0x0 && d.len == 0x2 && d.op == 0x1:
		d.encoding = encoding_TBX_asimdtbl_L3_3
	case d.op2 == 0x0 && d.len == 0x3 && d.op == 0x0:
		d.encoding = encoding_TBL_asimdtbl_L4_4
	case d.op2 == 0x0 && d.len == 0x3 && d.op == 0x1:
		d.encoding = encoding_TBX_asimdtbl_L4_4
	case (d.op2 & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdperm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdperm
	d.Q = (ins >> 30) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 12) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x0:
		err = errUnallocated
	case d.opcode == 0x1:
		d.encoding = encoding_UZP1_asimdperm_only
	case d.opcode == 0x2:
		d.encoding = encoding_TRN1_asimdperm_only
	case d.opcode == 0x3:
		d.encoding = encoding_ZIP1_asimdperm_only
	case d.opcode == 0x4:
		err = errUnallocated
	case d.opcode == 0x5:
		d.encoding = encoding_UZP2_asimdperm_only
	case d.opcode == 0x6:
		d.encoding = encoding_TRN2_asimdperm_only
	case d.opcode == 0x7:
		d.encoding = encoding_ZIP2_asimdperm_only
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdext(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdext
	d.Q = (ins >> 30) & 0x1
	d.op2 = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.imm4 = (ins >> 11) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.op2 & 0x1) == 0x1:
		err = errUnallocated
	case d.op2 == 0x0:
		d.encoding = encoding_EXT_asimdext_only
	case (d.op2 & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdins(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdins
	d.Q = (ins >> 30) & 0x1
	d.op = (ins >> 29) & 0x1
	d.imm5 = (ins >> 16) & 0x1f
	d.imm4 = (ins >> 11) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.imm5 & 0xf) == 0x0:
		err = errUnallocated
	case d.op == 0x0 && d.imm4 == 0x0:
		d.encoding = encoding_DUP_asimdins_DV_v
	case d.op == 0x0 && d.imm4 == 0x1:
		d.encoding = encoding_DUP_asimdins_DR_r
	case d.op == 0x0 && d.imm4 == 0x2:
		err = errUnallocated
	case d.op == 0x0 && d.imm4 == 0x4:
		err = errUnallocated
	case d.op == 0x0 && d.imm4 == 0x6:
		err = errUnallocated
	case d.op == 0x0 && (d.imm4&0x8) == 0x8:
		err = errUnallocated
	case d.Q == 0x0 && d.op == 0x0 && d.imm4 == 0x3:
		err = errUnallocated
	case d.Q == 0x0 && d.op == 0x0 && d.imm4 == 0x5:
		d.encoding = encoding_SMOV_asimdins_W_w
	case d.Q == 0x0 && d.op == 0x0 && d.imm4 == 0x7:
		d.encoding = encoding_UMOV_asimdins_W_w
	case d.Q == 0x0 && d.op == 0x1:
		err = errUnallocated
	case d.Q == 0x1 && d.op == 0x0 && d.imm4 == 0x3:
		d.encoding = encoding_INS_asimdins_IR_r
	case d.Q == 0x1 && d.op == 0x0 && d.imm4 == 0x5:
		d.encoding = encoding_SMOV_asimdins_X_x
	case d.Q == 0x1 && d.op == 0x0 && (d.imm5&0xf) == 0x8 && d.imm4 == 0x7:
		d.encoding = encoding_UMOV_asimdins_X_x
	case d.Q == 0x1 && d.op == 0x1:
		d.encoding = encoding_INS_asimdins_IV_v
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdsamefp16(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdsamefp16
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.a = (ins >> 23) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 11) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FMAXNM_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FMLA_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_FADD_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_FMULX_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCMEQ_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMAX_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FRECPS_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FMINNM_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FMLS_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_FSUB_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x4:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_FMIN_asimdsamefp16_only
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_FRSQRTS_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FMAXNMP_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_FADDP_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_FMUL_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCMGE_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FACGE_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMAXP_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FDIV_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FMINNMP_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_FABD_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_FCMGT_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_FACGT_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_FMINP_asimdsamefp16_only
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x7:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdmiscfp16(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdmiscfp16
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.a = (ins >> 23) & 0x1
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x18) == 0x0:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x8:
		err = errUnallocated
	case (d.opcode & 0x18) == 0x10:
		err = errUnallocated
	case d.opcode == 0x1e:
		err = errUnallocated
	case d.a == 0x0 && (d.opcode&0x1c) == 0xc:
		err = errUnallocated
	case d.a == 0x0 && d.opcode == 0x1f:
		err = errUnallocated
	case d.a == 0x1 && d.opcode == 0x1c:
		err = errUnallocated
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x18:
		d.encoding = encoding_FRINTN_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x19:
		d.encoding = encoding_FRINTM_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNS_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMS_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAS_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_SCVTF_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_FCMGT_asimdmiscfp16_FZ
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_FCMEQ_asimdmiscfp16_FZ
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xe:
		d.encoding = encoding_FCMLT_asimdmiscfp16_FZ
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_FABS_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x18:
		d.encoding = encoding_FRINTP_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x19:
		d.encoding = encoding_FRINTZ_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPS_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZS_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1d:
		d.encoding = encoding_FRECPE_asimdmiscfp16_R
	case d.U == 0x0 && d.a == 0x1 && d.opcode == 0x1f:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x18:
		d.encoding = encoding_FRINTA_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x19:
		d.encoding = encoding_FRINTX_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNU_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMU_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAU_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_UCVTF_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_FCMGE_asimdmiscfp16_FZ
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_FCMLE_asimdmiscfp16_FZ
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xe:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_FNEG_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x18:
		err = errUnallocated
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x19:
		d.encoding = encoding_FRINTI_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPU_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZU_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1d:
		d.encoding = encoding_FRSQRTE_asimdmiscfp16_R
	case d.U == 0x1 && d.a == 0x1 && d.opcode == 0x1f:
		d.encoding = encoding_FSQRT_asimdmiscfp16_R
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdsame2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdsame2
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 11) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.size&0x2) == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.size == 0x3 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x0:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SDOT_asimdsame2_D
	case d.U == 0x0 && (d.opcode&0x8) == 0x8:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0x3:
		d.encoding = encoding_USDOT_asimdsame2_D
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_SQRDMLAH_asimdsame2_only
	case d.U == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_SQRDMLSH_asimdsame2_only
	case d.U == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_UDOT_asimdsame2_D
	case d.U == 0x1 && (d.opcode&0xc) == 0x8:
		d.encoding = encoding_FCMLA_asimdsame2_C
	case d.U == 0x1 && (d.opcode&0xd) == 0xc:
		d.encoding = encoding_FCADD_asimdsame2_C
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0xd:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0xf:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_BFDOT_asimdsame2_D
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xd:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0xf:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0xf:
		d.encoding = encoding_BFMLAL_asimdsame2_F_
	case d.Q == 0x0 && (d.opcode&0xc) == 0x4:
		err = errUnallocated
	case d.Q == 0x0 && d.U == 0x1 && d.size == 0x1 && d.opcode == 0xd:
		err = errUnallocated
	case d.Q == 0x1 && (d.size&0x2) == 0x0 && (d.opcode&0xc) == 0x4:
		err = errUnallocated
	case d.Q == 0x1 && (d.size&0x2) == 0x2 && (d.opcode&0xe) == 0x6:
		err = errUnallocated
	case d.Q == 0x1 && d.U == 0x0 && d.size == 0x2 && d.opcode == 0x4:
		d.encoding = encoding_SMMLA_asimdsame2_G
	case d.Q == 0x1 && d.U == 0x0 && d.size == 0x2 && d.opcode == 0x5:
		d.encoding = encoding_USMMLA_asimdsame2_G
	case d.Q == 0x1 && d.U == 0x1 && d.size == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_BFMMLA_asimdsame2_E
	case d.Q == 0x1 && d.U == 0x1 && d.size == 0x2 && d.opcode == 0x4:
		d.encoding = encoding_UMMLA_asimdsame2_G
	case d.Q == 0x1 && d.U == 0x1 && d.size == 0x2 && d.opcode == 0x5:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdmisc(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdmisc
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x1e) == 0x10:
		err = errUnallocated
	case d.opcode == 0x15:
		err = errUnallocated
	case (d.size&0x2) == 0x0 && (d.opcode&0x1c) == 0xc:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opcode == 0x17:
		err = errUnallocated
	case (d.size&0x2) == 0x2 && d.opcode == 0x1e:
		err = errUnallocated
	case d.size == 0x3 && d.opcode == 0x16:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_REV64_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_REV16_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SADDLP_asimdmisc_P
	case d.U == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SUQADD_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_CLS_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_CNT_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_SADALP_asimdmisc_P
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_SQABS_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_CMGT_asimdmisc_Z
	case d.U == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_CMEQ_asimdmisc_Z
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_CMLT_asimdmisc_Z
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_ABS_asimdmisc_R
	case d.U == 0x0 && d.opcode == 0x12:
		d.encoding = encoding_XTN_asimdmisc_N
	case d.U == 0x0 && d.opcode == 0x13:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x14:
		d.encoding = encoding_SQXTN_asimdmisc_N
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x16:
		d.encoding = encoding_FCVTN_asimdmisc_N
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x17:
		d.encoding = encoding_FCVTL_asimdmisc_L
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x18:
		d.encoding = encoding_FRINTN_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x19:
		d.encoding = encoding_FRINTM_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNS_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMS_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAS_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_SCVTF_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1e:
		d.encoding = encoding_FRINT32Z_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FRINT64Z_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FCMGT_asimdmisc_FZ
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xd:
		d.encoding = encoding_FCMEQ_asimdmisc_FZ
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xe:
		d.encoding = encoding_FCMLT_asimdmisc_FZ
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_FABS_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x18:
		d.encoding = encoding_FRINTP_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x19:
		d.encoding = encoding_FRINTZ_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPS_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZS_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		d.encoding = encoding_URECPE_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FRECPE_asimdmisc_R
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0x16:
		d.encoding = encoding_BFCVTN_asimdmisc_4S
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_REV32_asimdmisc_R
	case d.U == 0x1 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_UADDLP_asimdmisc_P
	case d.U == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_USQADD_asimdmisc_R
	case d.U == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_CLZ_asimdmisc_R
	case d.U == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_UADALP_asimdmisc_P
	case d.U == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_SQNEG_asimdmisc_R
	case d.U == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_CMGE_asimdmisc_Z
	case d.U == 0x1 && d.opcode == 0x9:
		d.encoding = encoding_CMLE_asimdmisc_Z
	case d.U == 0x1 && d.opcode == 0xa:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xb:
		d.encoding = encoding_NEG_asimdmisc_R
	case d.U == 0x1 && d.opcode == 0x12:
		d.encoding = encoding_SQXTUN_asimdmisc_N
	case d.U == 0x1 && d.opcode == 0x13:
		d.encoding = encoding_SHLL_asimdmisc_S
	case d.U == 0x1 && d.opcode == 0x14:
		d.encoding = encoding_UQXTN_asimdmisc_N
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x16:
		d.encoding = encoding_FCVTXN_asimdmisc_N
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x17:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x18:
		d.encoding = encoding_FRINTA_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x19:
		d.encoding = encoding_FRINTX_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTNU_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTMU_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCVTAU_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_UCVTF_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1e:
		d.encoding = encoding_FRINT32X_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FRINT64X_asimdmisc_R
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_NOT_asimdmisc_R
	case d.U == 0x1 && d.size == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_RBIT_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FCMGE_asimdmisc_FZ
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xd:
		d.encoding = encoding_FCMLE_asimdmisc_FZ
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xe:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_FNEG_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x18:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x19:
		d.encoding = encoding_FRINTI_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FCVTPU_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		d.encoding = encoding_FCVTZU_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		d.encoding = encoding_URSQRTE_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FRSQRTE_asimdmisc_R
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		d.encoding = encoding_FSQRT_asimdmisc_R
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0x16:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdall(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdall
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.opcode = (ins >> 12) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x1e) == 0x0:
		err = errUnallocated
	case d.opcode == 0x2:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x4:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x8:
		err = errUnallocated
	case d.opcode == 0xb:
		err = errUnallocated
	case d.opcode == 0xd:
		err = errUnallocated
	case d.opcode == 0xe:
		err = errUnallocated
	case (d.opcode & 0x18) == 0x10:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x18:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x1c:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SADDLV_asimdall_only
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SMAXV_asimdall_only
	case d.U == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_SMINV_asimdall_only
	case d.U == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_ADDV_asimdall_only
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_FMAXNMV_asimdall_only_H
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_FMAXV_asimdall_only_H
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0xc:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0xf:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FMINNMV_asimdall_only_H
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_FMINV_asimdall_only_H
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0xc:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0xf:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_UADDLV_asimdall_only
	case d.U == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_UMAXV_asimdall_only
	case d.U == 0x1 && d.opcode == 0x1a:
		d.encoding = encoding_UMINV_asimdall_only
	case d.U == 0x1 && d.opcode == 0x1b:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_FMAXNMV_asimdall_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_FMAXV_asimdall_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FMINNMV_asimdall_only_SD
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_FMINV_asimdall_only_SD
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimddiff(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimddiff
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 12) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0xf:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SADDL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_SADDW_asimddiff_W
	case d.U == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SSUBL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SSUBW_asimddiff_W
	case d.U == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_ADDHN_asimddiff_N
	case d.U == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_SABAL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_SUBHN_asimddiff_N
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_SABDL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_SMLAL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_SQDMLAL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SMLSL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_SQDMLSL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_SMULL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_SQDMULL_asimddiff_L
	case d.U == 0x0 && d.opcode == 0xe:
		d.encoding = encoding_PMULL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_UADDL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_UADDW_asimddiff_W
	case d.U == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_USUBL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_USUBW_asimddiff_W
	case d.U == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_RADDHN_asimddiff_N
	case d.U == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_UABAL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_RSUBHN_asimddiff_N
	case d.U == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_UABDL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_UMLAL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0x9:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_UMLSL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0xb:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_UMULL_asimddiff_L
	case d.U == 0x1 && d.opcode == 0xd:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xe:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdsame(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdsame
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 11) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.U == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SHADD_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_SQADD_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SRHADD_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_SHSUB_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_SQSUB_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_CMGT_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_CMGE_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_SSHL_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_SQSHL_asimdsame_only
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SRSHL_asimdsame_only
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_SQRSHL_asimdsame_only
	case d.U == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_SMAX_asimdsame_only
	case d.U == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_SMIN_asimdsame_only
	case d.U == 0x0 && d.opcode == 0xe:
		d.encoding = encoding_SABD_asimdsame_only
	case d.U == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_SABA_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x10:
		d.encoding = encoding_ADD_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x11:
		d.encoding = encoding_CMTST_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x12:
		d.encoding = encoding_MLA_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x13:
		d.encoding = encoding_MUL_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x14:
		d.encoding = encoding_SMAXP_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x15:
		d.encoding = encoding_SMINP_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x16:
		d.encoding = encoding_SQDMULH_asimdsame_only
	case d.U == 0x0 && d.opcode == 0x17:
		d.encoding = encoding_ADDP_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x18:
		d.encoding = encoding_FMAXNM_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x19:
		d.encoding = encoding_FMLA_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FADD_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FMULX_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCMEQ_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1e:
		d.encoding = encoding_FMAX_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FRECPS_asimdsame_only
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_AND_asimdsame_only
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_FMLAL_asimdsame_F
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_BIC_asimdsame_only
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0x1d:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x18:
		d.encoding = encoding_FMINNM_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x19:
		d.encoding = encoding_FMLS_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FSUB_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1e:
		d.encoding = encoding_FMIN_asimdsame_only
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		d.encoding = encoding_FRSQRTS_asimdsame_only
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0x3:
		d.encoding = encoding_ORR_asimdsame_only
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FMLSL_asimdsame_F
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0x3:
		d.encoding = encoding_ORN_asimdsame_only
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0x1d:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_UHADD_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_UQADD_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_URHADD_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_UHSUB_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_UQSUB_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_CMHI_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_CMHS_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_USHL_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x9:
		d.encoding = encoding_UQSHL_asimdsame_only
	case d.U == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_URSHL_asimdsame_only
	case d.U == 0x1 && d.opcode == 0xb:
		d.encoding = encoding_UQRSHL_asimdsame_only
	case d.U == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_UMAX_asimdsame_only
	case d.U == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_UMIN_asimdsame_only
	case d.U == 0x1 && d.opcode == 0xe:
		d.encoding = encoding_UABD_asimdsame_only
	case d.U == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_UABA_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x10:
		d.encoding = encoding_SUB_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x11:
		d.encoding = encoding_CMEQ_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x12:
		d.encoding = encoding_MLS_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x13:
		d.encoding = encoding_PMUL_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x14:
		d.encoding = encoding_UMAXP_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x15:
		d.encoding = encoding_UMINP_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x16:
		d.encoding = encoding_SQRDMULH_asimdsame_only
	case d.U == 0x1 && d.opcode == 0x17:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x18:
		d.encoding = encoding_FMAXNMP_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1a:
		d.encoding = encoding_FADDP_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1b:
		d.encoding = encoding_FMUL_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_FCMGE_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1d:
		d.encoding = encoding_FACGE_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1e:
		d.encoding = encoding_FMAXP_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FDIV_asimdsame_only
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_EOR_asimdsame_only
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x19:
		d.encoding = encoding_FMLAL2_asimdsame_F
	case d.U == 0x1 && d.size == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_BSL_asimdsame_only
	case d.U == 0x1 && d.size == 0x1 && d.opcode == 0x19:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x18:
		d.encoding = encoding_FMINNMP_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1a:
		d.encoding = encoding_FABD_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1b:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1c:
		d.encoding = encoding_FCMGT_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1d:
		d.encoding = encoding_FACGT_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1e:
		d.encoding = encoding_FMINP_asimdsame_only
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x1f:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0x3:
		d.encoding = encoding_BIT_asimdsame_only
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0x19:
		d.encoding = encoding_FMLSL2_asimdsame_F
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x3:
		d.encoding = encoding_BIF_asimdsame_only
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x19:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdimm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdimm
	d.Q = (ins >> 30) & 0x1
	d.op = (ins >> 29) & 0x1
	d.a = (ins >> 18) & 0x1
	d.b = (ins >> 17) & 0x1
	d.c = (ins >> 16) & 0x1
	d.cmode = (ins >> 12) & 0xf
	d.o2 = (ins >> 11) & 0x1
	d.d = (ins >> 9) & 0x1
	d.e = (ins >> 8) & 0x1
	d.f = (ins >> 7) & 0x1
	d.g = (ins >> 6) & 0x1
	d.h = (ins >> 5) & 0x1
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.op == 0x0 && (d.cmode&0x8) == 0x0 && d.o2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && (d.cmode&0x9) == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_MOVI_asimdimm_L_sl
	case d.op == 0x0 && (d.cmode&0x9) == 0x1 && d.o2 == 0x0:
		d.encoding = encoding_ORR_asimdimm_L_sl
	case d.op == 0x0 && (d.cmode&0xc) == 0x8 && d.o2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && (d.cmode&0xd) == 0x8 && d.o2 == 0x0:
		d.encoding = encoding_MOVI_asimdimm_L_hl
	case d.op == 0x0 && (d.cmode&0xd) == 0x9 && d.o2 == 0x0:
		d.encoding = encoding_ORR_asimdimm_L_hl
	case d.op == 0x0 && (d.cmode&0xe) == 0xc && d.o2 == 0x0:
		d.encoding = encoding_MOVI_asimdimm_M_sm
	case d.op == 0x0 && (d.cmode&0xe) == 0xc && d.o2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && d.cmode == 0xe && d.o2 == 0x0:
		d.encoding = encoding_MOVI_asimdimm_N_b
	case d.op == 0x0 && d.cmode == 0xe && d.o2 == 0x1:
		err = errUnallocated
	case d.op == 0x0 && d.cmode == 0xf && d.o2 == 0x0:
		d.encoding = encoding_FMOV_asimdimm_S_s
	case d.op == 0x0 && d.cmode == 0xf && d.o2 == 0x1:
		d.encoding = encoding_FMOV_asimdimm_H_h
	case d.op == 0x1 && d.o2 == 0x1:
		err = errUnallocated
	case d.op == 0x1 && (d.cmode&0x9) == 0x0 && d.o2 == 0x0:
		d.encoding = encoding_MVNI_asimdimm_L_sl
	case d.op == 0x1 && (d.cmode&0x9) == 0x1 && d.o2 == 0x0:
		d.encoding = encoding_BIC_asimdimm_L_sl
	case d.op == 0x1 && (d.cmode&0xd) == 0x8 && d.o2 == 0x0:
		d.encoding = encoding_MVNI_asimdimm_L_hl
	case d.op == 0x1 && (d.cmode&0xd) == 0x9 && d.o2 == 0x0:
		d.encoding = encoding_BIC_asimdimm_L_hl
	case d.op == 0x1 && (d.cmode&0xe) == 0xc && d.o2 == 0x0:
		d.encoding = encoding_MVNI_asimdimm_M_sm
	case d.Q == 0x0 && d.op == 0x1 && d.cmode == 0xe && d.o2 == 0x0:
		d.encoding = encoding_MOVI_asimdimm_D_ds
	case d.Q == 0x0 && d.op == 0x1 && d.cmode == 0xf && d.o2 == 0x0:
		err = errUnallocated
	case d.Q == 0x1 && d.op == 0x1 && d.cmode == 0xe && d.o2 == 0x0:
		d.encoding = encoding_MOVI_asimdimm_D2_d
	case d.Q == 0x1 && d.op == 0x1 && d.cmode == 0xf && d.o2 == 0x0:
		d.encoding = encoding_FMOV_asimdimm_D2_d
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdshf(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdshf
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.immh = (ins >> 19) & 0xf
	d.immb = (ins >> 16) & 0x7
	d.opcode = (ins >> 11) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x1:
		err = errUnallocated
	case d.opcode == 0x3:
		err = errUnallocated
	case d.opcode == 0x5:
		err = errUnallocated
	case d.opcode == 0x7:
		err = errUnallocated
	case d.opcode == 0x9:
		err = errUnallocated
	case d.opcode == 0xb:
		err = errUnallocated
	case d.opcode == 0xd:
		err = errUnallocated
	case d.opcode == 0xf:
		err = errUnallocated
	case d.opcode == 0x15:
		err = errUnallocated
	case (d.opcode & 0x1e) == 0x16:
		err = errUnallocated
	case (d.opcode & 0x1c) == 0x18:
		err = errUnallocated
	case d.opcode == 0x1d:
		err = errUnallocated
	case d.opcode == 0x1e:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SSHR_asimdshf_R
	case d.U == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SSRA_asimdshf_R
	case d.U == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_SRSHR_asimdshf_R
	case d.U == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_SRSRA_asimdshf_R
	case d.U == 0x0 && d.opcode == 0x8:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SHL_asimdshf_R
	case d.U == 0x0 && d.opcode == 0xc:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0xe:
		d.encoding = encoding_SQSHL_asimdshf_R
	case d.U == 0x0 && d.opcode == 0x10:
		d.encoding = encoding_SHRN_asimdshf_N
	case d.U == 0x0 && d.opcode == 0x11:
		d.encoding = encoding_RSHRN_asimdshf_N
	case d.U == 0x0 && d.opcode == 0x12:
		d.encoding = encoding_SQSHRN_asimdshf_N
	case d.U == 0x0 && d.opcode == 0x13:
		d.encoding = encoding_SQRSHRN_asimdshf_N
	case d.U == 0x0 && d.opcode == 0x14:
		d.encoding = encoding_SSHLL_asimdshf_L
	case d.U == 0x0 && d.opcode == 0x1c:
		d.encoding = encoding_SCVTF_asimdshf_C
	case d.U == 0x0 && d.opcode == 0x1f:
		d.encoding = encoding_FCVTZS_asimdshf_C
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_USHR_asimdshf_R
	case d.U == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_USRA_asimdshf_R
	case d.U == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_URSHR_asimdshf_R
	case d.U == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_URSRA_asimdshf_R
	case d.U == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_SRI_asimdshf_R
	case d.U == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_SLI_asimdshf_R
	case d.U == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_SQSHLU_asimdshf_R
	case d.U == 0x1 && d.opcode == 0xe:
		d.encoding = encoding_UQSHL_asimdshf_R
	case d.U == 0x1 && d.opcode == 0x10:
		d.encoding = encoding_SQSHRUN_asimdshf_N
	case d.U == 0x1 && d.opcode == 0x11:
		d.encoding = encoding_SQRSHRUN_asimdshf_N
	case d.U == 0x1 && d.opcode == 0x12:
		d.encoding = encoding_UQSHRN_asimdshf_N
	case d.U == 0x1 && d.opcode == 0x13:
		d.encoding = encoding_UQRSHRN_asimdshf_N
	case d.U == 0x1 && d.opcode == 0x14:
		d.encoding = encoding_USHLL_asimdshf_L
	case d.U == 0x1 && d.opcode == 0x1c:
		d.encoding = encoding_UCVTF_asimdshf_C
	case d.U == 0x1 && d.opcode == 0x1f:
		d.encoding = encoding_FCVTZU_asimdshf_C
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_asimdelem(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_asimdelem
	d.Q = (ins >> 30) & 0x1
	d.U = (ins >> 29) & 0x1
	d.size = (ins >> 22) & 0x3
	d.L = (ins >> 21) & 0x1
	d.M = (ins >> 20) & 0x1
	d.Rm = (ins >> 16) & 0xf
	d.opcode = (ins >> 12) & 0xf
	d.H = (ins >> 11) & 0x1
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.size == 0x1 && d.opcode == 0x9:
		err = errUnallocated
	case d.U == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SMLAL_asimdelem_L
	case d.U == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_SQDMLAL_asimdelem_L
	case d.U == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_SMLSL_asimdelem_L
	case d.U == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_SQDMLSL_asimdelem_L
	case d.U == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_MUL_asimdelem_R
	case d.U == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_SMULL_asimdelem_L
	case d.U == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_SQDMULL_asimdelem_L
	case d.U == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_SQDMULH_asimdelem_R
	case d.U == 0x0 && d.opcode == 0xd:
		d.encoding = encoding_SQRDMULH_asimdelem_R
	case d.U == 0x0 && d.opcode == 0xe:
		d.encoding = encoding_SDOT_asimdelem_D
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x0:
		err = errUnallocated
	case d.U == 0x0 && (d.size&0x2) == 0x0 && d.opcode == 0x4:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FMLA_asimdelem_RH_H
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FMLS_asimdelem_RH_H
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_FMUL_asimdelem_RH_H
	case d.U == 0x0 && d.size == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_SUDOT_asimdelem_D
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_BFDOT_asimdelem_E
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FMLA_asimdelem_R_SD
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x5:
		d.encoding = encoding_FMLS_asimdelem_R_SD
	case d.U == 0x0 && (d.size&0x2) == 0x2 && d.opcode == 0x9:
		d.encoding = encoding_FMUL_asimdelem_R_SD
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FMLAL_asimdelem_LH
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0x4:
		d.encoding = encoding_FMLSL_asimdelem_LH
	case d.U == 0x0 && d.size == 0x2 && d.opcode == 0xf:
		d.encoding = encoding_USDOT_asimdelem_D
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0x0:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0x4:
		err = errUnallocated
	case d.U == 0x0 && d.size == 0x3 && d.opcode == 0xf:
		d.encoding = encoding_BFMLAL_asimdelem_F
	case d.U == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_MLA_asimdelem_R
	case d.U == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_UMLAL_asimdelem_L
	case d.U == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_MLS_asimdelem_R
	case d.U == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_UMLSL_asimdelem_L
	case d.U == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_UMULL_asimdelem_L
	case d.U == 0x1 && d.opcode == 0xb:
		err = errUnallocated
	case d.U == 0x1 && d.opcode == 0xd:
		d.encoding = encoding_SQRDMLAH_asimdelem_R
	case d.U == 0x1 && d.opcode == 0xe:
		d.encoding = encoding_UDOT_asimdelem_D
	case d.U == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_SQRDMLSH_asimdelem_R
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0x8:
		err = errUnallocated
	case d.U == 0x1 && (d.size&0x2) == 0x0 && d.opcode == 0xc:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x7:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_FMULX_asimdelem_RH_H
	case d.U == 0x1 && d.size == 0x1 && (d.opcode&0x9) == 0x1:
		d.encoding = encoding_FCMLA_asimdelem_C_H
	case d.U == 0x1 && (d.size&0x2) == 0x2 && d.opcode == 0x9:
		d.encoding = encoding_FMULX_asimdelem_R_SD
	case d.U == 0x1 && d.size == 0x2 && (d.opcode&0x9) == 0x1:
		d.encoding = encoding_FCMLA_asimdelem_C_S
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0x8:
		d.encoding = encoding_FMLAL2_asimdelem_LH
	case d.U == 0x1 && d.size == 0x2 && d.opcode == 0xc:
		d.encoding = encoding_FMLSL2_asimdelem_LH
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x1:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x3:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x5:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x7:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0x8:
		err = errUnallocated
	case d.U == 0x1 && d.size == 0x3 && d.opcode == 0xc:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_crypto3_imm2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_crypto3_imm2
	d.Rm = (ins >> 16) & 0x1f
	d.imm2 = (ins >> 12) & 0x3
	d.opcode = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x0:
		d.encoding = encoding_SM3TT1A_VVV4_crypto3_imm2
	case d.opcode == 0x1:
		d.encoding = encoding_SM3TT1B_VVV4_crypto3_imm2
	case d.opcode == 0x2:
		d.encoding = encoding_SM3TT2A_VVV4_crypto3_imm2
	case d.opcode == 0x3:
		d.encoding = encoding_SM3TT2B_VVV_crypto3_imm2
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_cryptosha512_3(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_cryptosha512_3
	d.Rm = (ins >> 16) & 0x1f
	d.O = (ins >> 14) & 0x1
	d.opcode = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.O == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_SHA512H_QQV_cryptosha512_3
	case d.O == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_SHA512H2_QQV_cryptosha512_3
	case d.O == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SHA512SU1_VVV2_cryptosha512_3
	case d.O == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_RAX1_VVV2_cryptosha512_3
	case d.O == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_SM3PARTW1_VVV4_cryptosha512_3
	case d.O == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_SM3PARTW2_VVV4_cryptosha512_3
	case d.O == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_SM4EKEY_VVV4_cryptosha512_3
	case d.O == 0x1 && d.opcode == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_crypto4(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_crypto4
	d.Op0 = (ins >> 21) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.Ra = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.Op0 == 0x0:
		d.encoding = encoding_EOR3_VVV16_crypto4
	case d.Op0 == 0x1:
		d.encoding = encoding_BCAX_VVV16_crypto4
	case d.Op0 == 0x2:
		d.encoding = encoding_SM3SS1_VVV4_crypto4
	case d.Op0 == 0x3:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_crypto3_imm6(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_crypto3_imm6
	d.Rm = (ins >> 16) & 0x1f
	d.imm6 = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f
	d.encoding = encoding_XAR_VVV2_crypto3_imm6
	return
}

func decode_simddp_cryptosha512_2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_cryptosha512_2
	d.opcode = (ins >> 10) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.opcode == 0x0:
		d.encoding = encoding_SHA512SU0_VV2_cryptosha512_2
	case d.opcode == 0x1:
		d.encoding = encoding_SM4E_VV4_cryptosha512_2
	case (d.opcode & 0x2) == 0x2:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_float2fix(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_float2fix
	d.sf = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.rmode = (ins >> 19) & 0x3
	d.opcode = (ins >> 16) & 0x7
	d.scale = (ins >> 10) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x4) == 0x4:
		err = errUnallocated
	case (d.rmode&0x1) == 0x0 && (d.opcode&0x6) == 0x0:
		err = errUnallocated
	case (d.rmode&0x1) == 0x1 && (d.opcode&0x6) == 0x2:
		err = errUnallocated
	case (d.rmode&0x2) == 0x0 && (d.opcode&0x6) == 0x0:
		err = errUnallocated
	case (d.rmode&0x2) == 0x2 && (d.opcode&0x6) == 0x2:
		err = errUnallocated
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && (d.scale&0x20) == 0x0:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_S32_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_S32_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_32S_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_32S_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_D32_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_D32_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_32D_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_32D_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_H32_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_H32_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_32H_float2fix
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_32H_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_S64_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_S64_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_64S_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_64S_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_D64_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_D64_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_64D_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_64D_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_H64_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_H64_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_64H_float2fix
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_64H_float2fix
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_float2int(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_float2int
	d.sf = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.rmode = (ins >> 19) & 0x3
	d.opcode = (ins >> 16) & 0x7
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.rmode&0x1) == 0x1 && (d.opcode&0x6) == 0x2:
		err = errUnallocated
	case (d.rmode&0x1) == 0x1 && (d.opcode&0x6) == 0x4:
		err = errUnallocated
	case (d.rmode&0x2) == 0x2 && (d.opcode&0x6) == 0x2:
		err = errUnallocated
	case (d.rmode&0x2) == 0x2 && (d.opcode&0x6) == 0x4:
		err = errUnallocated
	case d.S == 0x0 && d.ptype == 0x2 && (d.opcode&0x4) == 0x0:
		err = errUnallocated
	case d.S == 0x0 && d.ptype == 0x2 && (d.opcode&0x6) == 0x4:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && (d.rmode&0x1) == 0x1 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FCVTNS_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FCVTNU_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_S32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_S32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCVTAS_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVTAU_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMOV_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FMOV_S32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FCVTPS_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FCVTPU_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && (d.rmode&0x2) == 0x2 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FCVTMS_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FCVTMU_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_32S_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && (d.rmode&0x2) == 0x0 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FCVTNS_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FCVTNU_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_D32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_D32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCVTAS_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVTAU_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FCVTPS_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FCVTPU_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FCVTMS_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FCVTMU_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x2 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x6:
		d.encoding = encoding_FJCVTZS_32D_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x7:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x2 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FCVTNS_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FCVTNU_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_H32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_H32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCVTAS_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVTAU_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMOV_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FMOV_H32_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FCVTPS_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FCVTPU_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FCVTMS_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FCVTMU_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_32H_float2int
	case d.sf == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_32H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FCVTNS_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FCVTNU_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_S64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_S64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCVTAS_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVTAU_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FCVTPS_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FCVTPU_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FCVTMS_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FCVTMU_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x0 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_64S_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && (d.rmode&0x1) == 0x1 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FCVTNS_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FCVTNU_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_D64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_D64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCVTAS_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVTAU_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMOV_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FMOV_D64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FCVTPS_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FCVTPU_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && (d.rmode&0x2) == 0x2 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FCVTMS_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FCVTMU_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x1 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_64D_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x2 && (d.rmode&0x1) == 0x0 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x2 && d.rmode == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_FMOV_64VX_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x2 && d.rmode == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_FMOV_V64I_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x2 && (d.rmode&0x2) == 0x2 && (d.opcode&0x6) == 0x6:
		err = errUnallocated
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FCVTNS_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FCVTNU_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_SCVTF_H64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_UCVTF_H64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FCVTAS_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVTAU_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMOV_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FMOV_H64_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FCVTPS_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FCVTPU_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x2 && d.opcode == 0x0:
		d.encoding = encoding_FCVTMS_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x2 && d.opcode == 0x1:
		d.encoding = encoding_FCVTMU_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FCVTZS_64H_float2int
	case d.sf == 0x1 && d.S == 0x0 && d.ptype == 0x3 && d.rmode == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FCVTZU_64H_float2int
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatdp1(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatdp1
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.opcode = (ins >> 15) & 0x3f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x20) == 0x20:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FMOV_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FABS_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_FNEG_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_FSQRT_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x4:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FCVT_DS_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x6:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FCVT_HS_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_FRINTN_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x9:
		d.encoding = encoding_FRINTP_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0xa:
		d.encoding = encoding_FRINTM_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0xb:
		d.encoding = encoding_FRINTZ_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0xc:
		d.encoding = encoding_FRINTA_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0xd:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0xe:
		d.encoding = encoding_FRINTX_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0xf:
		d.encoding = encoding_FRINTI_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x10:
		d.encoding = encoding_FRINT32Z_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x11:
		d.encoding = encoding_FRINT32X_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x12:
		d.encoding = encoding_FRINT64Z_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x13:
		d.encoding = encoding_FRINT64X_S_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && (d.opcode&0x3c) == 0x14:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && (d.opcode&0x38) == 0x18:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FMOV_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FABS_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_FNEG_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_FSQRT_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_FCVT_SD_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x5:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_BFCVT_BS_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_FCVT_HD_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_FRINTN_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x9:
		d.encoding = encoding_FRINTP_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0xa:
		d.encoding = encoding_FRINTM_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0xb:
		d.encoding = encoding_FRINTZ_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0xc:
		d.encoding = encoding_FRINTA_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0xd:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0xe:
		d.encoding = encoding_FRINTX_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0xf:
		d.encoding = encoding_FRINTI_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x10:
		d.encoding = encoding_FRINT32Z_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x11:
		d.encoding = encoding_FRINT32X_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x12:
		d.encoding = encoding_FRINT64Z_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x13:
		d.encoding = encoding_FRINT64X_D_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && (d.opcode&0x3c) == 0x14:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && (d.opcode&0x38) == 0x18:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x2 && (d.opcode&0x20) == 0x0:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FMOV_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FABS_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x2:
		d.encoding = encoding_FNEG_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x3:
		d.encoding = encoding_FSQRT_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x4:
		d.encoding = encoding_FCVT_SH_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x5:
		d.encoding = encoding_FCVT_DH_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && (d.opcode&0x3e) == 0x6:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x8:
		d.encoding = encoding_FRINTN_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x9:
		d.encoding = encoding_FRINTP_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0xa:
		d.encoding = encoding_FRINTM_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0xb:
		d.encoding = encoding_FRINTZ_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0xc:
		d.encoding = encoding_FRINTA_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0xd:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0xe:
		d.encoding = encoding_FRINTX_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0xf:
		d.encoding = encoding_FRINTI_H_floatdp1
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && (d.opcode&0x30) == 0x10:
		err = errUnallocated
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatcmp(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatcmp
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.op = (ins >> 14) & 0x3
	d.Rn = (ins >> 5) & 0x1f
	d.opcode2 = (ins >> 0) & 0x1f

	switch {
	case (d.opcode2 & 0x1) == 0x1:
		err = errUnallocated
	case (d.opcode2 & 0x2) == 0x2:
		err = errUnallocated
	case (d.opcode2 & 0x4) == 0x4:
		err = errUnallocated
	case (d.op & 0x1) == 0x1:
		err = errUnallocated
	case (d.op & 0x2) == 0x2:
		err = errUnallocated
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.op == 0x0 && d.opcode2 == 0x0:
		d.encoding = encoding_FCMP_S_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.op == 0x0 && d.opcode2 == 0x8:
		d.encoding = encoding_FCMP_SZ_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.op == 0x0 && d.opcode2 == 0x10:
		d.encoding = encoding_FCMPE_S_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.op == 0x0 && d.opcode2 == 0x18:
		d.encoding = encoding_FCMPE_SZ_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.op == 0x0 && d.opcode2 == 0x0:
		d.encoding = encoding_FCMP_D_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.op == 0x0 && d.opcode2 == 0x8:
		d.encoding = encoding_FCMP_DZ_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.op == 0x0 && d.opcode2 == 0x10:
		d.encoding = encoding_FCMPE_D_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.op == 0x0 && d.opcode2 == 0x18:
		d.encoding = encoding_FCMPE_DZ_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.op == 0x0 && d.opcode2 == 0x0:
		d.encoding = encoding_FCMP_H_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.op == 0x0 && d.opcode2 == 0x8:
		d.encoding = encoding_FCMP_HZ_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.op == 0x0 && d.opcode2 == 0x10:
		d.encoding = encoding_FCMPE_H_floatcmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.op == 0x0 && d.opcode2 == 0x18:
		d.encoding = encoding_FCMPE_HZ_floatcmp
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatimm(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatimm
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.imm8 = (ins >> 13) & 0xff
	d.imm5 = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.imm5 & 0x1) == 0x1:
		err = errUnallocated
	case (d.imm5 & 0x2) == 0x2:
		err = errUnallocated
	case (d.imm5 & 0x4) == 0x4:
		err = errUnallocated
	case (d.imm5 & 0x8) == 0x8:
		err = errUnallocated
	case (d.imm5 & 0x10) == 0x10:
		err = errUnallocated
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.imm5 == 0x0:
		d.encoding = encoding_FMOV_S_floatimm
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.imm5 == 0x0:
		d.encoding = encoding_FMOV_D_floatimm
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.imm5 == 0x0:
		d.encoding = encoding_FMOV_H_floatimm
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatccmp(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatccmp
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.cond = (ins >> 12) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.op = (ins >> 4) & 0x1
	d.nzcv = (ins >> 0) & 0xf

	switch {
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.op == 0x0:
		d.encoding = encoding_FCCMP_S_floatccmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.op == 0x1:
		d.encoding = encoding_FCCMPE_S_floatccmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.op == 0x0:
		d.encoding = encoding_FCCMP_D_floatccmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.op == 0x1:
		d.encoding = encoding_FCCMPE_D_floatccmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.op == 0x0:
		d.encoding = encoding_FCCMP_H_floatccmp
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.op == 0x1:
		d.encoding = encoding_FCCMPE_H_floatccmp
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatdp2(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatdp2
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.opcode = (ins >> 12) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case (d.opcode & 0x9) == 0x9:
		err = errUnallocated
	case (d.opcode & 0xa) == 0xa:
		err = errUnallocated
	case (d.opcode & 0xc) == 0xc:
		err = errUnallocated
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x0:
		d.encoding = encoding_FMUL_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x1:
		d.encoding = encoding_FDIV_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x2:
		d.encoding = encoding_FADD_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x3:
		d.encoding = encoding_FSUB_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x4:
		d.encoding = encoding_FMAX_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x5:
		d.encoding = encoding_FMIN_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x6:
		d.encoding = encoding_FMAXNM_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x7:
		d.encoding = encoding_FMINNM_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.opcode == 0x8:
		d.encoding = encoding_FNMUL_S_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x0:
		d.encoding = encoding_FMUL_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x1:
		d.encoding = encoding_FDIV_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x2:
		d.encoding = encoding_FADD_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x3:
		d.encoding = encoding_FSUB_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x4:
		d.encoding = encoding_FMAX_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x5:
		d.encoding = encoding_FMIN_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x6:
		d.encoding = encoding_FMAXNM_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x7:
		d.encoding = encoding_FMINNM_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.opcode == 0x8:
		d.encoding = encoding_FNMUL_D_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x0:
		d.encoding = encoding_FMUL_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x1:
		d.encoding = encoding_FDIV_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x2:
		d.encoding = encoding_FADD_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x3:
		d.encoding = encoding_FSUB_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x4:
		d.encoding = encoding_FMAX_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x5:
		d.encoding = encoding_FMIN_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x6:
		d.encoding = encoding_FMAXNM_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x7:
		d.encoding = encoding_FMINNM_H_floatdp2
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.opcode == 0x8:
		d.encoding = encoding_FNMUL_H_floatdp2
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatsel(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatsel
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.Rm = (ins >> 16) & 0x1f
	d.cond = (ins >> 12) & 0xf
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0:
		d.encoding = encoding_FCSEL_S_floatsel
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1:
		d.encoding = encoding_FCSEL_D_floatsel
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3:
		d.encoding = encoding_FCSEL_H_floatsel
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}

func decode_simddp_floatdp3(ins uint32, d *decoded) (err error) {
	d.iclass = iclass_floatdp3
	d.M = (ins >> 31) & 0x1
	d.S = (ins >> 29) & 0x1
	d.ptype = (ins >> 22) & 0x3
	d.o1 = (ins >> 21) & 0x1
	d.Rm = (ins >> 16) & 0x1f
	d.o0 = (ins >> 15) & 0x1
	d.Ra = (ins >> 10) & 0x1f
	d.Rn = (ins >> 5) & 0x1f
	d.Rd = (ins >> 0) & 0x1f

	switch {
	case d.ptype == 0x2:
		err = errUnallocated
	case d.S == 0x1:
		err = errUnallocated
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.o1 == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_FMADD_S_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.o1 == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_FMSUB_S_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.o1 == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_FNMADD_S_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x0 && d.o1 == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_FNMSUB_S_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.o1 == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_FMADD_D_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.o1 == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_FMSUB_D_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.o1 == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_FNMADD_D_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x1 && d.o1 == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_FNMSUB_D_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.o1 == 0x0 && d.o0 == 0x0:
		d.encoding = encoding_FMADD_H_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.o1 == 0x0 && d.o0 == 0x1:
		d.encoding = encoding_FMSUB_H_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.o1 == 0x1 && d.o0 == 0x0:
		d.encoding = encoding_FNMADD_H_floatdp3
	case d.M == 0x0 && d.S == 0x0 && d.ptype == 0x3 && d.o1 == 0x1 && d.o0 == 0x1:
		d.encoding = encoding_FNMSUB_H_floatdp3
	case d.M == 0x1:
		err = errUnallocated
	default:
		err = errUnmatched
	}
	return
}
