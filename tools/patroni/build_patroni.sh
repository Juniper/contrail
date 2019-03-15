#!/bin/bash

patroni_version="1.5.1"

[[ "$(docker images -q patroni)" == "" ]] || { echo "Patroni image already exists. Skipping building docker image." ; exit 0; }

tmpdir=$(mktemp -d -t patroni-repository-XXXXXX) || { echo "Failed to create temporary directory" ; exit 1; }
echo "Downloading patroni"
(cd $tmpdir && curl -LO "https://github.com/zalando/patroni/archive/v$patroni_version.zip" --connect-timeout 60) || { echo "Failed to download patroni repository" ; exit 1; }
echo "Upacking repository archive"
unzip -q "$tmpdir/v$patroni_version.zip" -d $tmpdir && rm "$tmpdir/v$patroni_version.zip" || { echo "Failed to exctract repository archive" ; exit 1; }
echo "Archive unpacked"
sed -i 's/FROM postgres:10/FROM circleci\/postgres:10-ram/g' $tmpdir/patroni-$patroni_version/Dockerfile
docker build -t patroni "$tmpdir/patroni-$patroni_version" || { echo "Failed to build docker image" ; exit 1; }
rm -rf $tmpdir || { echo "Failed to remove temporary directory" ; exit 1; }
