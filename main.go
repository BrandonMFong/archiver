/**
 * author: brando
 * date: 11/6/24
 */

package main

import (
	"os"
	"log"
	"fmt"
	"path/filepath"
	"strings"
	"flag"
)

const VERSION = "0.1"

// the archive directory name
const ARCHIVE_DIR_NAME = "archive"

// version of program
const ARG_VERSION = "version"

// reads arg
//
// conditionally exits programs
func ArgumentsRead() {
	flag.Usage = func() {
		fmt.Println("usage: " + os.Args[0] + " [ <args> ] <path>")
		fmt.Println()
		fmt.Println("args:")
		flag.PrintDefaults()
		fmt.Println("")
		fmt.Println("Copyright © 2024 Brando. All rights reserved.")
	}

	var exit_program bool = false
	flag.BoolFunc("version", "print version", func(s string) error {
		fmt.Println(VERSION)
		exit_program = true
		return nil
	})

	flag.Parse()

	if exit_program {
		os.Exit(0)
	}
}

// main:
// sets up the directory we are archive then moves everything to 
// the archive folder
func main() {
	ArgumentsRead()
	SetupEnv()
	DoArchive()
}

// returns the target directory
//
// either can be the last argument or current directory
func GetTargetDir() string {
	if len(os.Args) < 2 {
		return "."
	} else {
		return os.Args[len(os.Args)-1]
	}
}

// creates the archive folder if not exists
// exits program on error
func SetupEnv() {
	archive_path := filepath.Join(GetTargetDir(), ARCHIVE_DIR_NAME)
    folderInfo, err := os.Stat(archive_path)
    if !os.IsNotExist(err) {
		// if it does exist, it might be a file
		if !folderInfo.IsDir() {
			log.Fatalln(" '" + archive_path + "' " + "exists, but isn't a directory")
		}
	} else {
		err := os.Mkdir(archive_path, 0750)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Pushes all items into the archive folder
func DoArchive() {
	paths := TargetDirGetPaths()

	for _, path := range paths {
		PathArchive(path)
	}
}

// pushes path into the archive folder
func PathArchive(path string) {
	newpath := PathGetNewName(path)

	err := os.Rename(path, newpath) // move
	if err != nil {
		fmt.Println(err)
	}
}

// gets the new path as if it were in the archive folder already
// 
// resolves naming conflicts
//
// exits program on error
func PathGetNewName(path string) string {
	var res string
	archive_path := filepath.Join(GetTargetDir(), ARCHIVE_DIR_NAME) // archive folder
	ext := filepath.Ext(path)
	base_name := filepath.Base(path) // file name w/ ext
	file_name := strings.TrimSuffix(base_name, ext) // file name w/o ext
	res = filepath.Join(archive_path, base_name) // temp new path

	// resolve name conflicts
	i := 1
	for PathExists(res) {
		str := fmt.Sprint(file_name, "_", i, ext)
		res = filepath.Join(archive_path, str)
		i++

		if i >= 1000 {
			log.Fatalln("name conflict resolution exhausted")
		}
	}

	return res
}

// checks if path exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil { return true }
	return false
}

// gets a list of paths that are subject to be archived
func TargetDirGetPaths() []string {
	items, err := os.ReadDir(GetTargetDir())
	if err != nil {
		log.Fatal(err)
	}

	// craft current path
	var res []string
	for _, item := range items {
		if item.Name() != ARCHIVE_DIR_NAME {
			res = append(res, filepath.Join(GetTargetDir(), item.Name()))
		}
	}

	return res
}

