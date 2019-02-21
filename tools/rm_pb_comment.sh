#!/usr/bin/env bash

# This is a workaround for GoLand IDE which skips file indexing when file starts with long comment block.
# https://youtrack.jetbrains.com/issue/GO-6952

remove_comment_block()
{
    sed -i -r ':a; s%(.*)/\*.*\*/%\1%; ta; /\/\*/ !b; N; ba' $1
}
export -f remove_comment_block

find . -type f -name *.pb.* -not -path './vendor/*' -exec /bin/bash -c 'remove_comment_block "$0"' {} \;
