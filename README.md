# Lyricfier 2

![Lyricfier2 Screenshot](screenshots/lyricfier2.jpg?raw=true "Lyricfier Screenshot")

Lyrics For Spotify App

Lyricfier2 is a rewrite of the old lyricfier using Golang and nuklear bindings for Go. Is planned to be cross platform but was tested only on Linux.

*Warning: We are in a very early stage!* 

## Building and running in Ubuntu 16.04


```bash
$ sudo apt go-dep install xorg-dev libgl1-mesa-dev
$ git clone git@github.com:emilioastarita/lyricfier2.git
$ cd lyricfier2
$ dep ensure 
$ cd cmd/
$ go build
$ ./cmd
```

