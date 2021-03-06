# Lyricfier 2

Lyrics For Spotify App

Lyricfier2 is a rewrite of the old lyricfier using Golang and running as a web server.

*Warning: We are in a very early stage!* 

## Downloads

Go to releases page:

https://github.com/emilioastarita/lyricfier2/releases/latest


## Installation Linux

Install from snap store

```bash
sudo snap install lyricfier
sudo snap connect lyricfier:mpris spotify:spotify-mpris
```


[![Get it from the Snap Store](https://snapcraft.io/static/images/badges/en/snap-store-black.svg)](https://snapcraft.io/lyricfier)

## How to build

### Install golang 

In Ubuntu, you can use snap

```bash
sudo snap install --classic --channel=1.14/stable go
```

### Install build dependencies 


```bash
sudo apt-get install libgtk-3-dev libappindicator3-dev libwebkit2gtk-4.0-dev
```

Get `esc` utility 

```bash
go get -u github.com/mjibson/esc
```


### Clone repo and build

```bash
git clone git@github.com:emilioastarita/lyricfier2.git
cd lyricfier2/
# add go/bin directory to path
PATH=$PATH:~/go/bin/ make build
```

[Download latest release](https://github.com/emilioastarita/lyricfier2/releases/latest)

![Lyricfier 2 in Ubuntu](screenshots/screenshot-lyricfier.jpg?raw=true "Lyricfier 2")
