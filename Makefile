#export GOPATH=/home/server/goprojects3

plat ?= linux
plats = linux darwin

all: strsrv

define build_server
        @echo 'building $(1) ...'
        @GOOS=$(2) GOARCH=amd64 go build -o builder/$(1) ./apps/$(1)
        @echo 'build $(1) done'
endef

strsrv:
	$(call build_server,strsrv,$(plat))

.PHONY: strsrv
