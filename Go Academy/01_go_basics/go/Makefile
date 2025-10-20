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

# Cert paths
CERT_FOLDER := certs
CA_CERT := $(CERT_FOLDER)/ca.pem
MOVIE_CERT_FOLDER := $(CERT_FOLDER)/movies
CHARACTER_CERT_FOLDER := $(CERT_FOLDER)/characters

# Go CLI entrypoint
CERT_GEN := go run cert/generator.go

.PHONY: create-ca
create-ca:
	$(CERT_GEN) -mode=ca -out=$(CA_CERT)

.PHONY: create-movie
create-movie:
	$(CERT_GEN) -mode=movie -name=$(NAME) -ca=$(CA_CERT) -out=$(MOVIE_CERT_FOLDER)/$(NAME).pem

.PHONY: create-character
create-character:
	$(CERT_GEN) -mode=character -name=$(NAME) -movie=$(MOVIE_CERT_FOLDER)/$(MOVIE) -out=$(CHARACTER_CERT_FOLDER)/$(NAME).pem

PORT ?= 8080
APP_NAME ?= movie-character-api

.PHONY: docker-build
docker-build:
	docker build -f ../Dockerfile -t $(APP_NAME):latest ..

.PHONY: docker-run
docker-run:
	docker run -p $(PORT):$(PORT) $(APP_NAME):latest
