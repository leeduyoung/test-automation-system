.DEFAULT_GOAL := start

.PHONY: gen
gen: generate

.PHONY: generate
generate: 
	cd pkg/graph && go run github.com/99designs/gqlgen generate

.PHONY: setup
setup: 
	go mod tidy

.PHONY: start	
start: setup
	cd src && go run main.go

.PHONY: test
test: test_setup test_run test_clean

.PHONY: test_setup
test_setup:
	@echo "building the environment.."
	docker-compose -f docker/docker-compose.yml up --build -d
	@sleep 5
	@echo "environment build is done."

.PHONY: test_run
test_run:
	@echo "started run the all tests."
	cd test && go test -v -count 1 -p 1 -timeout 120s ./...
	@echo "all tests were completed."

.PHONY: test_clean
test_clean:
	@echo "cleaning the environment.."
	docker-compose -f docker/docker-compose.yml down
	@echo "environment cleaned up."