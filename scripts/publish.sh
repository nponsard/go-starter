#!/bin/sh
. ./build-config

rm -rf publish
mkdir -p publish
cp package/*${VERSION}*.* publish