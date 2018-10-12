#!/bin/bash

ensure_group()
{
	local expected_group='docker'
	local user
	user=$(id -nu)

	grep -qE "^$expected_group:" /etc/group || sudo groupadd "$expected_group" # ensure group exists
	groups | grep -q "$expected_group" || sudo usermod -aG "$expected_group" "$user" # ensure user is in that group

	if [ "$(id -gn)" != "$expected_group" ]; then
		exec sg "$expected_group" -c "$0" "$@"
	fi
}

