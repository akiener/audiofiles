fmt:
	@go fmt ./...

test:
	@go test ./...

lint:
	@golangci-lint run ./...

gen:
	@go generate ./...

build:
	@echo "building..."
	@go build -ldflags="-s -w" -o /tmp/audiofiles cmd/*
	@echo "packing..."
	@upx -q -9 /tmp/audiofiles > /dev/null
	@mv /tmp/audiofiles .ci/
	@echo "built .ci/audiofiles"

run:
	@docker-compose -f .ci/docker-compose.yml -p audiofiles up

clean:
	@docker-compose -f .ci/docker-compose.yml -p audiofiles down
	-docker rmi audiofiles_bot

clean-run:
	@$(MAKE) $(THIS_FILE) clean
	@$(MAKE) $(THIS_FILE) run