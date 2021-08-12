# 说明：
# 可加参数指定编译平台，plat=linux
# 可指定编译其中某些程序，详情见all：列表

#export GOPATH=$(shell pwd)

plat ?= darwin
plats = linux darwin

arch ?= amd64
archs = amd64 arm arm64

all: kvs

define build_app
        @echo 'building $(1) ...'
        @GOOS=$(2) GOARCH=$(3) go build -o builder/$(1) ./cmd/$(1)
        @echo 'build $(1) done'
endef

kvs:
	$(call build_app,kvs,$(plat),$(arch))

.PHONY: kvs

clean:
	rm -f builder/*