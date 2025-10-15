# Makefile for running Echo server and k6 tests

K6_FOLDER := k6
K6 := k6
GO := go
MAIN := main.go

.PHONY: help
help:
	@echo "Usage:"
	@echo "  make run               Start Echo server"
	@echo "  make test-get          Run GET /movies test"
	@echo "  make test-post         Run POST /movies test"
	@echo "  make test-all          Run all k6 tests"
	@echo "  make test FILE=your.js Run custom test file"

.PHONY: run
run:
	$(GO) run $(MAIN)

.PHONY: test-get
test-get:
	$(K6) run $(K6_FOLDER)/get_movies_test.js

.PHONY: test-post
test-post:
	$(K6) run $(K6_FOLDER)/create_movie_test.js

.PHONY: test-all
test-all: test-get test-post

.PHONY: test
test:
	$(K6) run $(K6_FOLDER)/$(FILE)