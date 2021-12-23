#!/bin/bash

if [ -z "${OPERATOR_SDK_VERSION}" ]
then
    echo "OPERATOR_SDK_VERSION should be defined as an env variable"
    exit 1
fi

if [ -z "${OPERATOR_SDK}" ]
then
    echo "OPERATOR_SDK should be defined as an env variable"
    exit 1
fi

mkdir -p "$(dirname "${OPERATOR_SDK}")"
if [ -x "${OPERATOR_SDK}" ]
then
    "${OPERATOR_SDK}" version | grep -q "\"${OPERATOR_SDK_VERSION}\""
    if [ $? -eq 0 ]
    then
        exit 0
    fi
fi

PLATFORM="$(uname)"
ARCHITECTURE="$(uname -m)"
if [ "${PLATFORM}" == "Darwin" ]
then 
    PLAT="darwin"
else # Assuming that if not darwin then running on linux
    PLAT="linux"
fi
if [ "${ARCHITECTURE}" == "x86_64" ]
then 
    ARCH="amd64"
else # Assuming that if not x86 then arm. TODO ppc64le, s390x
    ARCH="arm64"
fi

SDK_RELEASE="https://github.com/operator-framework/operator-sdk/releases/download/${OPERATOR_SDK_VERSION}/operator-sdk_${PLAT}_${ARCH}"

echo "installing version ${OPERATOR_SDK_VERSION}"
curl -f "${SDK_RELEASE}" -Lo "${OPERATOR_SDK}"
if [ $? -ne 0 ]
then
    echo "could not download and install ${SDK_RELEASE}"
    exit 1
fi
chmod +x "${OPERATOR_SDK}"
