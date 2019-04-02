#!/usr/bin/env bash

DEPLOY_DIR=deploy/linux
RELEASE_DIR=deploy/lyricfier
ICONS=${DEPLOY_DIR}/icons/
rm -rf ${DEPLOY_DIR}
qtdeploy build desktop

# bin
mv ${DEPLOY_DIR}/cmd ${DEPLOY_DIR}/lyricfier
# desktop file
cp lyricfier.desktop ${DEPLOY_DIR}
# icon
declare -a SIZES=("16x16" "24x24" "32x32" "48x48" "22x22" "128x128" "256x256")

for SIZE in ${SIZES[@]}; do
   mkdir -p ${ICONS}/${SIZE}/
   convert assets/icon.png -resize ${SIZE}\> ${ICONS}/${SIZE}/lyricfier.png
done

mv ${DEPLOY_DIR}  ${RELEASE_DIR}
tar -zcvf release.tar.gz -C ${RELEASE_DIR}/../ lyricfier

