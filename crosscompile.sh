#!/bin/bash

set -e

VERSION=${TRAVIS_TAG:-version}
PACKAGE=${PACKAGE:-package}

function build-and-zip() {
  echo "Building for $1 $2"

  dir=/tmp/${PACKAGE}-$1-$2
  cwd=`pwd`

  mkdir -p ${dir}
  GOOS=$1 GOARCH=$2 go build -a -tags netgo -installsuffix netgo -ldflags '-extldflags "-static"' -o ${dir}/${PACKAGE}
  cd ${dir}
  zip ${cwd}/${PACKAGE}-$1-$2.zip ${PACKAGE}
  tar czfv ${cwd}/${PACKAGE}-$1-$2.tgz ${PACKAGE}

  cd ${cwd}
}

build-and-zip linux 386
build-and-zip linux amd64
build-and-zip linux arm
build-and-zip linux arm64
build-and-zip darwin 386
build-and-zip darwin amd64
