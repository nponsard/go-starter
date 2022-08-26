#!/bin/sh

MOD_URI="$1"
MOD_NAME=$(basename $MOD_URI)

if [ -z "$MOD_URI" ]; then
    echo "Usage: $0 <mod-uri>"
    exit 1
fi

files=(./main.go ./.env ./internal/commands/commands.go  ./internal/commands/ping/ping.go  ./.gitignore  ./go.mod  ./Makefile )

for file in "${files[@]}"
do
    if [ -f $file ]; then
      sed -i 's|github.com/nilsponsard/go-starter|'$MOD_URI'|g' $file
      sed -i 's|go-starter|'$MOD_NAME'|g' $file
    fi
done

rm ./bootstrap.sh