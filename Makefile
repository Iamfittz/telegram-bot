APP_NAME := telegram-bot

REGISTRY := ghcr.io
OWNER := iamfittz
IMAGE := $(REGISTRY)/$(OWNER)/$(APP_NAME)

VERSION := v1.0.0
GIT_SHA := $(shell git rev-parse --short=7 HEAD)

OS := linux
ARCH := amd64

TAG := $(VERSION)-$(GIT_SHA)
FULL_TAG := $(TAG)-$(OS)-$(ARCH)

.PHONY: build image push

build:
	docker build -t $(APP_NAME):local .

image:
	docker buildx build --platform $(OS)/$(ARCH) -t $(IMAGE):$(FULL_TAG) .

push:
	docker push $(IMAGE):$(FULL_TAG)
