#!/bin/env bash

install_golang()
{
	pushd /tmp
	curl -o go.tar.gz https://dl.google.com/go/go1.12.6.linux-amd64.tar.gz
	sudo tar --overwrite -C /usr -xzf go.tar.gz
	sudo yum install -y wget unzip
	export PATH="$PATH:/usr/go/bin"
	hash -r
	popd
	go env
}

ensure_golang_installed()
{
    if [ -d /usr/go/bin ]; then
        echo "$PATH" | grep -q /usr/go/bin || export PATH="$PATH:/usr/go/bin"
    fi
    go env || install_golang
    [ -z "$GOPATH" ] && export GOPATH="$HOME/go"
    echo "$PATH" | grep -q "$GOPATH/bin" || export PATH="$PATH:$GOPATH/bin"
}
