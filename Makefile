IMAGE_NAME = kramuenke/platform-test

ifndef CIRCLE_BUILD_NUM
	override CIRCLE_BUILD_NUM = 0
endif

.PHONY: test
test:  ## Create the docker image.
	@echo "--- test"
	docker run \
		--rm \
		--workdir=/app \
		-v `pwd`:/app \
		golang:1.16-bullseye \
		scripts/test.sh

.PHONY: local-test
local-test:
	scripts/test.sh

.PHONY: start
start:
	@echo "--- :fire: start :fire:"
	cd app; go run ./cmd/...

.PHONY: build
build:
	docker build --build-arg CIRCLE_SHA1 -t ${IMAGE_NAME}:${CIRCLE_BUILD_NUM} .

.PHONY: push
push:
	docker push ${IMAGE_NAME}:${CIRCLE_BUILD_NUM}
