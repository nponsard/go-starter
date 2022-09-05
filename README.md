# go-starter project

You will need to change the url in `go.mod` and the names in `build-config` and `Makefile` to match your project name & VCS hoster

## Dependencies

- make
- go
- pandoc for manual generation

## Boostraping

Use `./bootsrap.sh <project-uri>` to rename every import automatically.
Example project uri : `github.com/nilsponsard/go-starter`

## make commands

- `make all` : builds for Windows, linux generic and Ubuntu/debian (deb), builds the manuals and put everything in the `publish` folder
- `make` : builds for your current platform

## Folder structure

### assets

Non source files used in the program

### configs

config file

### DEBIAN

Configuration file for packaging a .deb package.
Update dependencies in DEBIAN/control.

### internal

Code that won’t be reusable in other projects

### manuals

Sources files to build a manual
TODO : document how to write and build a documentation

### pkg

Code that can be reused in other projects like log system

### scripts

Scripts used by the Makefile to do different build actions
