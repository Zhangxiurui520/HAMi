GO=go
GO111MODULE=on
CMDS=scheduler vGPUmonitor
DEVICES=nvidia
OUTPUT_DIR=bin
TARGET_ARCH=amd64
GOLANG_IMAGE=registry.nscc-tj.cn/hami/golang:1.24.6-bullseye
NVIDIA_IMAGE=registry.nscc-tj.cn/hami/nvidia/cuda:12.3.2-devel-ubuntu20.04
DEST_DIR=/usr/local/vgpu/

VERSION = v2.7.0
IMG_NAME =registry.nscc-tj.cn/hami/hami-mars-filter
IMG_TAG="${IMG_NAME}:${VERSION}"

