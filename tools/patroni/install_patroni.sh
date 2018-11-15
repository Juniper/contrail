#/bin/bash

patroni_version="1.5.1"
echo "Downloading patroni"
wget "https://github.com/zalando/patroni/archive/v$patroni_version.zip"
echo "Upacking repository archive"
unzip -q "v$patroni_version.zip" && rm "v$patroni_version.zip"
echo "Archive unpacked"
docker build -t patroni "./patroni-$patroni_version"

