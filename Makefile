.DEFAULT_GOAL := start

.PHONY: setup
setup: 
	go mod tidy

.PHONY: start	
start: setup
	cd src && go run main.go

.PHONY: test
test:
	@echo "started run the all tests."
	cd src/lib && go test -v -count 1 -p 1 -timeout 120s ./...
	@echo "all tests were completed."

.PHONY: build	
build: 
	docker build -t ldy/test-automation-system:latest . -f Dockerfile