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

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
PACKAGE_NAME = archiver-linux
endif
ifeq ($(UNAME_S),Darwin)
PACKAGE_NAME = archiver-macos
endif
PACKAGE_BIN_PATH = $(BIN_PATH)

LIBS_MAKEFILES_PATH:=$(CURDIR)/external/libs/makefiles
include $(LIBS_MAKEFILES_PATH)/package.mk 

