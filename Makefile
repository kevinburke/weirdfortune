.PHONY: install venv clean

venv:
	virtualenv venv

install: venv
	. venv/bin/activate; python setup.py install

clean: 
	rm -rf venv
