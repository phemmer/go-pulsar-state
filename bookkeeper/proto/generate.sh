#!/bin/sh

set -e

ref=release-4.10.0
curl -L https://github.com/apache/bookkeeper/archive/$ref.tar.gz | tar xz --strip-components=6 bookkeeper-$ref/stream/proto/src/main/proto

for file in *.proto; do
  pkg_name="$(grep "^package " "$file" | sed "s/^package bookkeeper\.proto\.\(.*\);/\1/; s/\./\//g")"
  sed -i "/^package/a option go_package = \"github.com/phemmer/go-pulsar-state/bookkeeper/proto/$pkg_name\";" "$file"
done
for file in *.proto; do
  protoc -I=. --go_out=plugins=grpc:../../../../../ "$file"
done
