#!/usr/bin/env nix-shell
#! nix-shell -I nixpkgs=https://github.com/NixOS/nixpkgs/archive/b0d36bd0a420ecee3bc916c91886caca87c894e9.tar.gz
#! nix-shell -p go_1_21 jdk21_headless cacert curl screen
#! nix-shell -i sh --pure
# shellcheck shell=sh

set -eu
cd "$(dirname "$(realpath "$0")")/.."
set -x

url="https://maven.reposilite.com/releases/com/reposilite/reposilite/3.5.2/reposilite-3.5.2-all.jar"
sha="2cb341fc35d1e53446db37072bc381e8c73d496da39c27a494c5fd7d20ecd505"
file="reposilite.jar"
token="testing:correcthorsebatterystaple"

start_reposilite() (

    # workdir
    mkdir -p "reposilite"
    cd "reposilite"

    # download server jar
    if test ! -f ${file}; then
        curl -fsSLo "${file}" "${url}"
        echo "${sha} ${file}" | sha256sum -c
    fi

    # run in background
    screen -dmS "reposilite" -- java -Xmx32M -jar "${file}" \
        --hostname "127.0.0.1" \
        --port "8080" \
        --token "${token}"

    # wait for startup
    if ! curl \
        --retry 10 \
        --retry-delay 1 \
        --retry-connrefused \
        --silent \
        --output /dev/null \
        "http://127.0.0.1:8080/"; then
        echo >&2 "reposilite server unreachable, killing screen..."
        screen -S "reposilite" -X quit
        exit 1
    fi
    sleep 1

)

start_reposilite

# run vet and tests
go test -vet "" ./...

# stop reposilite server
screen -S "reposilite" -X stuff "stop^M"
