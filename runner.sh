#!/usr/bin/env bash

GO_VERSION=go1.4
GO_HOME=/data/workspace/goenv
GO_PLATFORM=$GO_HOME/src/github.com/gotravelydia/platform
USER_HOME=/home/ubuntu
GVM_EXE=$USER_HOME/.gvm/bin/gvm
GVM_SCRIPT=$USER_HOME/.gvm/scripts/gvm

if [ ! -e $GVM_SCRIPT ]
then
    wget https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer
    chmod +x gvm-installer
    ./gvm-installer
fi

echo "using go version $GO_VERSION from $GVM_SCRIPT"
echo "source $GVM_SCRIPT"
source $GVM_SCRIPT
gvm use $GO_VERSION

cd $GO_HOME
export GOPATH=$GO_HOME
export PATH=$PATH:$GOPATH/bin
cd $GO_PLATFORM

$GOPATH/bin/"$@"
