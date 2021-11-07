.PHONY: test
test:  ## Create the docker image.
	@echo "--- test"
	docker run \
		--rm \
		--workdir=/app \
		-v `pwd`:/app \
		golang:1.16-bullseye \
		scripts/test.sh

.PHONY: start
start:
	@echo "--- :fire: start :fire:"
	cd app; go run ./cmd/...