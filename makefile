## 
# author: brando
# date: 11/6/24
#

CONFIG = release
BIN_PATH = bin
BIN_NAME = archive

help:
	@echo "Usage:"
	@echo "	make <target> <variables>"
	@echo ""
	@echo "Target(s):"
	@echo "	clean			cleans build and bin folder"
	@echo "	build 			builds release verions"
	@echo "	package			compresses build"
	@echo ""
	@echo "Variable(s):"
	@echo "	CONFIG		use this to change the build config. Accepts \"release\" (default), \"debug\", or \"test\""
	@echo "	IDENTITY	(macos only) \"Developer ID Application\" common name"
	@echo "	TEAMID 		(macos only) Organizational Unit"
	@echo "	EMAIL 		(macos only) Developer account email"
	@echo "	PW		(macos only) Developer account password"
	@echo ""
	@echo "Example(s):"
	@echo "	Build for release for macOS distribution"
	@echo "		make clean build codesign package notarize staple IDENTITY=\"\" TEAMID=\"\" EMAIL=\"\" PW=\"\""
	@echo "	Build for release for Linux distribution"
	@echo "		make clean build package"

build:
	@go build -o $(BIN_PATH)/$(CONFIG)/$(BIN_NAME) .

clean:
	rm -rfv $(BIN_PATH)

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
PACKAGE_NAME = archiver-linux
endif
ifeq ($(UNAME_S),Darwin)
PACKAGE_NAME = archiver-macos
endif
PACKAGE_BIN_PATH = $(BIN_PATH)/$(CONFIG)
PACKAGE_BIN_TARGET = $(BIN_NAME)

LIBS_MAKEFILES_PATH:=$(CURDIR)/external/libs/makefiles
include $(LIBS_MAKEFILES_PATH)/package.mk 

