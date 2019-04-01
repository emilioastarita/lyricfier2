# Lyricfier 2



![Lyricfier2 Screenshot](screenshots/lyricfier2_osx.jpg?raw=true "Lyricfier Screenshot")

Lyrics For Spotify App

Lyricfier2 is a rewrite of the old lyricfier using Golang and qt bindings for Go.

*Warning: We are in a very early stage!* 


## qt building dependencies

Lyricfier 2 is using qt so you need to install some deps for building. Follow the instructions for your platform in [qt golang binding](https://github.com/therecipe/qt#installation).

## Building and running in Ubuntu 16.04

```bash
# deps
$ sudo apt install golang-go go-dep go-bindata
# Create dir for clone
$ mkdir -p $GOPATH/src/github.com/emilioastarita/lyricfier2
# clone repo
$ git clone git@github.com:emilioastarita/lyricfier2.git $GOPATH/src/github.com/emilioastarita/lyricfier2
$ cd lyricfier2
# install deps
$ dep ensure 
$ cd cmd/
$ go run main.go
```

### Tested platforms

- windows
- darwin
- linux
