MAKEFLAGS += --no-print-directory
MASTER_BRANCH=main
EXEC=tradeFetcher
SRCDIR=./
TST_FILES=$(wildcard $(SRCDIR)/*/*/*_test.go $(SRCDIR)/*/*_test.go $(SRCDIR)/*_test.go)
SRC_FILES=$(wildcard $(SRCDIR)/*/*/*.go $(SRCDIR)/*/*.go $(SRCDIR)/*.go)
SRC1= $(filter-out $(wildcard $(SRCDIR)/*_test.go), $(SRC_FILES))
SRC2= $(filter-out $(wildcard $(SRCDIR)/*/*_test.go), $(SRC1))
SRC= $(filter-out $(wildcard $(SRCDIR)/*/*/*_test.go), $(SRC2))
INTERFACE_FILES=$(wildcard $(SRCDIR)/*/*/i[A-Z]*.go $(SRCDIR)/*/i[A-Z]*.go $(SRCDIR)/i[A-Z]*.go)
TEST_PACKAGE_WITHOUT_COVERAGE=tradeFetcher/integrationTests
EXCLUDED_PACKAGE_FOR_BENCHMARK=$(shell go list ./... | grep -v generatedMocks | grep -v model/)
EXCLUDED_PACKAGE_FOR_COVERAGE=$(filter-out $(TEST_PACKAGE_WITHOUT_COVERAGE), $(EXCLUDED_PACKAGE_FOR_BENCHMARK))
COVERAGE_FILE=coverOutput.cov
TARGET_COVERAGE_PERCENT=99.9

.PHONY: all clear 

all: displayCompilation $(EXEC) run-tests code-analysis list-todo 

$(EXEC): $(SRC) makefile
	go mod tidy
	go fmt ./...
	go build -o . ./...

displayCompilation:
	@echo "${_RED}  --COMPILATION ${_END}"

code-analysis:
	go vet ./...
	golangci-lint run -D exportloopref

clear: clean
	rm -rf $(EXEC)

mock-generator: $(INTERFACE_FILES)
	$(foreach interface, $^, mockgen -package=generatedMocks -source=./$(interface) -destination=./internal/generatedMocks/$(shell basename $(shell dirname $(interface)))-$(basename $(notdir $(interface)))Mock.go;)

run-tests: mock-generator $(TST_FILES)
	@echo "${_RED}  --LAUNCH TESTS ${_END}"
	
	go mod tidy
	go test -cover -coverprofile=$(COVERAGE_FILE) $(EXCLUDED_PACKAGE_FOR_COVERAGE) 
	go test $(TEST_PACKAGE_WITHOUT_COVERAGE)
	@ACTUAL_COVERAGE=`go tool cover -func $(COVERAGE_FILE) | tail -n 1 | awk -F '[ \t]+' '{print $$3}' | cut -d '%' -f1`; \
	COVERAGE_COMPARAISON_RESULT=`echo "$$ACTUAL_COVERAGE <= $(TARGET_COVERAGE_PERCENT)" | bc`; \
	if [ $$COVERAGE_COMPARAISON_RESULT -eq 1 ]; then exit "Insufisent coverage $$ACTUAL_COVERAGE%"; fi;

	go test -bench=. $(EXCLUDED_PACKAGE_FOR_BENCHMARK)

list-todo:
	@echo "${_RED}  --LOOKING FOR TODO ${_END}"
	@grep -I -ri 'todo ' $(SRCDIR) | grep -v 'makefile:' || true
	
release: test-branch-master $(EXEC) test-release-version
	
test-branch-master:
	@if [ `git branch --show-current` != "$(MASTER_BRANCH)" ]; then exit "release can only build on master"; fi; \
	HASH_COMMIT=`git log -n 1 $(MASTER_BRANCH) --pretty=format:"%H"`; \
	TAG_ON_COMMIT=`git tag --contains $$HASH_COMMIT | grep -E "^V[0-9]+\.[0-9]+\.[0-9]+$\"`; \
	echo $$TAG_ON_COMMIT; \
	if [ -z $$TAG_ON_COMMIT ]; then exit "Last commit missed a version tag"; fi; \
	VERSION_UPDATED=`git diff-tree --no-commit-id --name-only -m -r $$HASH_COMMIT | grep tradeFetcher.go`; \
	if [ -z $$VERSION_UPDATED ]; then exit "No new verson on last commit"; fi; \
	
test-release-version: $(EXEC)
	@LAST_TAG_VERSION=`git describe --tags --abbrev=0 | grep -E "^V[0-9]+\.[0-9]+\.[0-9]+$\"`; \
	if [ -z $$LAST_TAG_VERSION ]; then exit "Last tag not match policy"; fi; \
	if [ `./$(EXEC) -v` != $$LAST_TAG_VERSION ]; then exit "Last tag not match application version"; fi