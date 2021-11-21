ARCH=A64
VERSION=v9
EDITION=2021-09

ISA=ISA_$(ARCH)_xml_$(VERSION)A-$(EDITION)
ARCHIVE=$(ISA).tar.gz
URL=https://developer.arm.com/-/media/developer/products/architecture/arm$(VERSION)-a-architecture/$(EDITION)/$(ARCHIVE)
ENCODINGINDEX=$(ISA)/encodingindex.xml

PYTHON=python3
VENV=env
ACTIVATE=$(VENV)/bin/activate

gen: $(ACTIVATE) $(ENCODINGINDEX)
	source $(ACTIVATE) && $(PYTHON) gen.py $(ENCODINGINDEX) ../
	go fmt ../

$(ACTIVATE):
	$(PYTHON) -m venv $(VENV)
	source $(ACTIVATE) && $(PYTHON) -m pip install -r requirements.txt

$(ENCODINGINDEX): $(ARCHIVE)
	tar xfm $^

$(ARCHIVE):
	@ which wget &>/dev/null || (echo "wget is required"; exit 1)
	wget "$(URL)"

.PHONY: clean
clean:
	rm -rf $(VENV) *.html *.pdf ISA_$(ARCH)*
