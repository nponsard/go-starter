#!/bin/sh

MOD_URI="test"
MOD_NAME=$(basename $MOD_URI)


files=(./main.go,./.env,./internal/commands/commands.go, ./internal/commands/ping/ping.go, ./.gitignore, ./go.mod, ./Makefile )

for file in "${files[@]}"
do
    if [ -f $file ]; then
      sed -i "s/github.com/nilsponsard/go-starter/$OUT/g" ./template.go 
      sed -i "s/go-starter/$MOD_NAME/g" ./template.go
    fi
done

