venv:
	virtualenv venv

install: venv
	. venv/bin/activate; python setup.py install
