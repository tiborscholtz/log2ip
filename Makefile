APP_NAME := log2ip
CMD_PATH := ./cmd/log2ip
BIN_DIR  := ./bin
SCRIPTS_DIR  := ./scripts

.PHONY: run build clean tidy test

run:
	$(SCRIPTS_DIR)/start.sh

build:
	$(SCRIPTS_DIR)/build.sh

clean:
	rm -rf $(BIN_DIR)

tidy:
	go mod tidy
