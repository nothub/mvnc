#!/usr/bin/env nix-shell
#! nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/b0d36bd0a420ecee3bc916c91886caca87c894e9.tar.gz
#! nix-shell -p go_1_21 upx
#! nix-shell -i sh --pure
# shellcheck shell=sh

set -eu
cd "$(dirname "$(realpath "$0")")"
set -x

rm -rf "dist"
mkdir -p "dist"

go vet ./...
go test -v ./...

build() {
    file="dist/mvnc_${1}_${2}"

    GOOS="$1" GOARCH="$2" go build \
        -ldflags "-s -w" \
        -o "${file}" \
        ./cmd

    if test "$1" != "darwin"; then
        upx --best --lzma \
            --no-color \
            --no-progress \
            --no-time \
            "${file}"
    fi

    if test "$1" = "windows"; then
        mv "${file}" "${file}.exe"
    fi
}

build linux amd64
build linux arm64
build darwin amd64
build darwin arm64
build windows amd64
