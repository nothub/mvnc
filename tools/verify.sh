#!/usr/bin/env nix-shell
#! nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/a4d4fe8c5002202493e87ec8dbc91335ff55552c.tar.gz
#! nix-shell -p go_1_22
#! nix-shell -i sh --pure
# shellcheck shell=sh

set -eu
cd "$(dirname "$(realpath "$0")")/.."

set -x

go test -v -vet "" ./...
