BIN_DIR := bin
BIN_FILE := $(BIN_DIR)/main
SRC_DIR := solutions/$(DIR)
SRC_FILES := $(shell find $(SRC_DIR) -name '*.go') $(SRC_DIR)/go.mod

all:
	go run $(SRC_DIR)/main.go

$(BIN_DIR):
	mkdir $@

clean:
	rm -rf bin

$(BIN_FILE): $(BIN_DIR) $(SRC_FILES)
	CGO_ENABLED=0 GOOS=linux go build -o $@ $(SRC_DIR)/main.go

rsync: $(BIN_FILE)
	rsync -avP $(BIN_FILE) $(SSH_USER)@$(SSH_FQDN):~/protohackers-go/
	
test:
	cd $(SRC_DIR) && ls -l && go test
