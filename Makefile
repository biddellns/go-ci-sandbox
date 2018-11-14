REGISTRY ?= kohonski
IMAGE ?= go-ci-sandbox
TAG ?= latest

all: container

container:
	docker build -t $(REGISTRY)/$(IMAGE):$(TAG) .

push:
	docker push $(REGISTRY)/$(IMAGE):$(TAG)
