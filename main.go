package main

import (
	"fmt"
	"log"
	"os"

	"./cliargs"
	"./particeps"
)

func assertNonNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg := cliargs.ParseCLIArgs(os.Args)
	fileSize, err := particeps.CheckFile(cfg.Filename)
	assertNonNil(err)
	fmt.Printf("particeps: file \"%s\" has size %s\n", cfg.Filename, fileSize)

	if cfg.Destination == particeps.AnonFiles {
		fmt.Println("particeps: uploading to https://anonfiles.com/")
		res, err := particeps.AnonFilesUpload(cfg.Filename)
		assertNonNil(err)
		fmt.Printf("particeps: successfully uploaded \"%s\" to https://anonfiles.com/\n", cfg.Filename)
		fmt.Printf("particeps: full-length link: %s\n", res.FullURL)
		fmt.Printf("particeps: short link: %s\n", res.ShortURL)
	}
}
