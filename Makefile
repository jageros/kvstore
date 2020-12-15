#export GOPATH=/home/server/goprojects3

plat ?= linux
plats = linux darwin

all: strsvc

define build_server
        @echo 'building $(1) ...'
        @GOOS=$(2) GOARCH=amd64 go build -o builder/$(1) ./$(1)
        @echo 'build $(1) done'
endef

strsvc:
	$(call build_server,app,$(plat))

.PHONY: strsvc
