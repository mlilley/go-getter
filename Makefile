GO_SRC = cmd/go-getter/go-getter.go
BINARY_NAME = go-getter
BUILD_DIR = bin

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(GO_SRC)

clean:
	@rm -rf $(BUILD_DIR)

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: all build clean run
