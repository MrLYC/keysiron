GOPATH := $(shell pwd):${GOPATH}
GO15VENDOREXPERIMENT := 1

ROOTDIR := $(shell pwd)
SRCDIR := src/keysiron
GLIDE := cd ${SRCDIR} && glide

TARGET := bin/keysiron

goenv:
	goenv init keysiron

bin:
	mkdir -p bin

${SRCDIR}:
	mkdir -p src
	ln -s ${ROOTDIR}/keysiron src/

${SRCDIR}/glide.yaml: ${SRCDIR}
	${GLIDE} init

src/%: keysiron/vendor/%
	@ln -s "$<" "$@" >/dev/null 2>&1 || true

.PHONY: vendor
vendor:$(addprefix src/, $(notdir $(wildcard ${SRCDIR}/vendor/*)))

.PHONY: init
init: ${SRCDIR}

.PHONY: update
update: vendor
	${GLIDE} update

.PHONY: install
install:
	${GLIDE} install

.PHONY: build
build: bin
	go build -o ${TARGET} keysiron

.PHONY: run
run:
	go run ${SRCDIR}/main.go

.PHONY: clean
clean:
	rm -rf bin || true
