# this makefile builds go api into an archive file (static library)

BUILD_FLAGS=--trimpath -gcflags "-N"

BUILD_ENVS=CGO_ENABLED=1 CGO_CFLAGS="-mmacosx-version-min=12.1" CGO_CXXFLAGS="-mmacosx-version-min=12.1" CC=clang

.PHONY: all clean

all: libwssocks_go_api.a

libwssocks_go_api_amd64.a: wssocks_client_wrapper.go
	$(BUILD_ENVS) GOOS=darwin GOARCH=amd64 go build $(BUILD_FLAGS) --buildmode=c-archive -o libwssocks_go_api_amd64.a
	mv libwssocks_go_api_amd64.h libwssocks_go_api.h

libwssocks_go_api_arm64.a: wssocks_client_wrapper.go
	$(BUILD_ENVS) GOOS=darwin GOARCH=arm64 go build $(BUILD_FLAGS) --buildmode=c-archive -o libwssocks_go_api_arm64.a
	mv libwssocks_go_api_arm64.h libwssocks_go_api.h

libwssocks_go_api.a: libwssocks_go_api_amd64.a libwssocks_go_api_arm64.a
	lipo libwssocks_go_api_amd64.a libwssocks_go_api_arm64.a -create -output libwssocks_go_api.a

clean:
	rm libwssocks_go_api.a libwssocks_go_api_arm64.a libwssocks_go_api_amd64.a libwssocks_go_api.h
