#!/usr/bin/env bash

DEPLOY_DIR=deploy/linux
DEBIAN_RELEASE=${DEPLOY_DIR}/release
DEBIAN_BIN=${DEBIAN_RELEASE}/usr/bin
DEBIAN_ICONS=${DEBIAN_RELEASE}/usr/share/icons/hicolor
DEBIAN_APPLICATIONS=${DEBIAN_RELEASE}/usr/share/applications
rm -rf ${DEPLOY_DIR}
qtdeploy build desktop

# debian dir
mkdir -p ${DEBIAN_RELEASE}/DEBIAN
cp debian-control ${DEBIAN_RELEASE}/DEBIAN/control
# bin
mkdir -p ${DEBIAN_BIN}
cp ${DEPLOY_DIR}/cmd ${DEBIAN_BIN}/lyricfier
# desktop file
mkdir -p ${DEBIAN_APPLICATIONS}
cp lyricfier.desktop ${DEBIAN_APPLICATIONS}
# icon

declare -a SIZES=("16x16" "24x24" "32x32" "48x48" "22x22" "128x128" "256x256")

for SIZE in ${SIZES[@]}; do
   mkdir -p ${DEBIAN_ICONS}/${SIZE}/apps/
   convert assets/icon.png -resize ${SIZE}\> ${DEBIAN_ICONS}/${SIZE}/apps/lyricfier.png
done

# Create release
dpkg-deb --build ${DEBIAN_RELEASE}
