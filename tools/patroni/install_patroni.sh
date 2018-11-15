#/bin/bash

set -e

patroni_version="1.5.1"

tmpdir=$(mktemp -t -d patroni-repository-XXXXXX)
echo "Downloading patroni"
wget "https://github.com/zalando/patroni/archive/v$patroni_version.zip" -P $tmpdir
echo "Upacking repository archive"
unzip -q "$tmpdir/v$patroni_version.zip" -d $tmpdir && rm "$tmpdir/v$patroni_version.zip"
echo "Archive unpacked"
docker build -t patroni "$tmpdir/patroni-$patroni_version"
rm -rf $tmpdir

