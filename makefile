## 
# author: brando
# date: 11/6/24
#

CONFIG = release
BIN_PATH = bin
BIN_NAME = archive

build:
	@go build -o $(BIN_PATH)/$(CONFIG)/$(BIN_NAME) .

clean:
	rm -rfv $(BIN_PATH)

dependencies:

