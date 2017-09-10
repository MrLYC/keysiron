GOPATH := $(shell pwd):${GOPATH}
GO15VENDOREXPERIMENT := 1

ROOTDIR := $(shell pwd)
SRCDIR := src/keysiron
GLIDE := cd ${SRCDIR} && glide

goenv:
	goenv init keysiron

bin:
	mkdir -p bin

${SRCDIR}:
	mkdir -p src
	ln -s ${ROOTDIR}/keysiron src/

${SRCDIR}/glide.yaml: ${SRCDIR}
	${GLIDE} init

.PHONY: init
init: ${SRCDIR}

.PHONY: update
update:
	${GLIDE} update

.PHONY: install
install:
	${GLIDE} install

.PHONY: build
build: bin
	go build -o ${ROOTDIR}/bin/keysiron keysiron

.PHONY: run
run:
	go run ${SRCDIR}/main.go

.PHONY: clean
clean:
	rm -rf bin || true
