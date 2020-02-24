#!/usr/bin/env bash

if [ -z "$1" ]
  then
    echo "Usage ./deploy-debian.bash LYRICFIER_VERSION"
    exit 1
fi

LYRICFIER_VERSION=$1
RELEASE_DIR=releases/debian/lyricfier_${LYRICFIER_VERSION}
DEBIAN_SOURCE=lyricfier/debian
LYRICFIER_SOURCE=lyricfier/

echo $RELEASE_DIR;
if [ -d $RELEASE_DIR ]
  then
    echo "Directory $RELEASE_DIR already exists remove before continue."
    exit 1
fi
mkdir -p $RELEASE_DIR

cp -R $DEBIAN_SOURCE/* $RELEASE_DIR/
go build  -o $RELEASE_DIR/usr/lib/lyricfier/lyricfier $LYRICFIER_SOURCE/*.go
cp -R $LYRICFIER_SOURCE/static/ $LYRICFIER_SOURCE/views/ $RELEASE_DIR/usr/lib/lyricfier/
dpkg-deb --build $RELEASE_DIR/




