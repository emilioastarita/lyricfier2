# Lyricfier 2

![Lyricfier2 Screenshot](screenshots/lyricfier2.jpg?raw=true "Lyricfier Screenshot")

Lyrics For Spotify App

Lyricfier2 is a rewrite of the old lyricfier using Golang and nuklear bindings for Go. Is planned to be cross platform but was tested only on Linux.

*Warning: We are in a very early stage!* 

## Building and running in Ubuntu 16.04


```bash
# deps
$ sudo apt install golang-go go-dep xorg-dev libgl1-mesa-dev
# Create dir for clone
$ mkdir -p $GOPATH/src/github.com/emilioastarita/lyricfier2
# clone repo
$ git clone git@github.com:emilioastarita/lyricfier2.git $GOPATH/src/github.com/emilioastarita/lyricfier2
$ cd lyricfier2
# install deps
$ dep ensure 
$ cd cmd/
# build
$ go build
# test build
$ ./cmd
```

