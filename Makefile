.PHONY: run
.SILENT: node go

VALUES=1/3 1.5 -2.001953125 -2.000000238418579 -2.0000000000000004 0.00000095367431640625 0.0000000000000000000000000000000000000000000014012984643248170709237295832899161312802619418765157718

all: run

node_modules/@ipld/block/package.json:
	npm install

node: node_modules/@ipld/block/package.json
	@echo
	@echo Running Node.js ...
	node floats.js $(VALUES)

go:
	@echo
	@echo Running Go ...
	go run floats.go $(VALUES)

run: node go
