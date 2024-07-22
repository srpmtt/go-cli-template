GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
MAIN_PACKAGE=./cmd/myApp

clean:
	@echo "CLEAN"
	@$(GOCLEAN)
	@rm -rf ./bin

build:
	@echo "BUILD"
	clear
	@$(GOBUILD) -o bin/myApp $(MAIN_PACKAGE)

run: build
	@echo "RUN"
	clear
	@./bin/myApp

test: build
	clear
	@echo "TEST"
	@$(GOTEST) ./tests
