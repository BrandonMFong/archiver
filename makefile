## 
# author: brando
# date: 11/6/24
#

BIN_PATH = bin
BIN_NAME = archive

build:
	@go build -o $(BIN_PATH)/$(BIN_NAME) .

clean:
	rm -rfv $(BIN_PATH)

