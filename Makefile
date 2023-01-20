
clean:
	rm -rf pb/*
	rm -f plugins/host/host.wasm
	rm -f ./host
#uses v0.5.0 of protoc-gen-go-plugin
compile:
	mkdir -p pb/greenfinch
	protoc --go-plugin_out=./pb/greenfinch --go-plugin_opt=paths=source_relative --proto_path=./proto ./proto/greenfinch.proto

plugin:
	 tinygo build -o plugins/demo/demo.wasm -scheduler=none -target=wasi --no-debug plugins/demo/demo.go

.PHONY: host
host:
	go build -o host

run:
	./host
#
#
#MYDIR = proto
#ls:
#	$(foreach file, $(wildcard $(MYDIR)/*), echo ${file%/*})
#
#something=fdff/aaa/bbb/ccc/ddd/eee/fff.txt
#test:
#	tail="${$(something)#/*/*/}"
#	head="${$(something)%/$(tail)}"
#	echo $(head) $(tail) $(something)
