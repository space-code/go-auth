all: bootstrap

bootstrap: hook lint

hook:
	pre-commit install

lint:
	npm install -g @commitlint/cli @commitlint/config-conventional

addlicense:
	addlicense -c 'space-code' -l mit ./**/*.go
