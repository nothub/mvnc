#!/usr/bin/env sh

url="https://maven.reposilite.com/releases/com/reposilite/reposilite/3.5.2/reposilite-3.5.2-all.jar"
sha="2cb341fc35d1e53446db37072bc381e8c73d496da39c27a494c5fd7d20ecd505"
file="reposilite.jar"
token="testing:correcthorsebatterystaple"

set -eu
cd "$(dirname "$(realpath "$0")")/.."

mkdir -p reposilite
cd reposilite

if test ! -f ${file}; then
    curl -fsSLo "${file}" "${url}"
    echo "${sha} ${file}" | sha256sum -c
fi

java -Xmx32M -jar "${file}" --token "${token}"
