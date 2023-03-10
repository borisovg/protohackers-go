BIN_DIR := bin
BIN_FILE := $(BIN_DIR)/main
COVER_FILE := cover.out
SRC_DIR := solutions/$(DIR)
SRC_FILES := $(shell find $(SRC_DIR) -name '*.go') $(SRC_DIR)/go.mod

all:
	golangci-lint run $(SRC_DIR)
	go run $(SRC_DIR)/main.go

$(BIN_DIR):
	mkdir $@

clean:
	rm -rf bin
	find -name $(COVER_FILE) -delete

$(BIN_FILE): $(BIN_DIR) $(SRC_FILES)
	CGO_ENABLED=0 GOOS=linux go build -o $@ $(SRC_DIR)/main.go

$(SRC_DIR):
	mkdir $(SRC_DIR) && \
		cd $(SRC_DIR) && \
		go mod init protohackers/$(DIR) && \
		printf "%b\n" "package main\n\nfunc main() {\n}" > main.go
	go work use $(SRC_DIR)

new: $(SRC_DIR)

rsync: $(BIN_FILE)
	rsync -avP $(BIN_FILE) $(SSH_USER)@$(SSH_FQDN):~/protohackers-go/
	
test:
	cd $(SRC_DIR) && go test -cover -coverprofile=$(COVER_FILE) ./...
