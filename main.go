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

// helper function for AnonFiles & BayFiles
// anonfiles == true  => anonfiles
// anonfiles == false => bayfiles
func helperAnonFiles(cfg cliargs.CLIArgs, anonfiles bool) {
	uploadFunction := particeps.AnonFilesUpload
	var website string
	if anonfiles {
		website = "https://anonfiles.com"
	} else {
		website = "https://bayfiles.com"
		uploadFunction = particeps.BayFilesUpload
	}
	fmt.Printf("particeps: uploading to %s\n", website)
	res, err := uploadFunction(cfg.Filename)
	assertNonNil(err)
	fmt.Printf("particeps: successfully uploaded \"%s\" to %s/\n", cfg.Filename, website)
	fmt.Printf("particeps: full-length link: %s\n", res.FullURL)
	fmt.Printf("particeps: short link: %s\n", res.ShortURL)
}

func main() {
	cfg := cliargs.ParseCLIArgs(os.Args)
	fileSize, err := particeps.CheckFile(cfg.Filename)
	assertNonNil(err)
	fmt.Printf("particeps: file \"%s\" has size %s\n", cfg.Filename, fileSize)
	if cfg.Destination == particeps.AnonFiles {
		helperAnonFiles(cfg, true)
	} else if cfg.Destination == particeps.BayFiles {
		helperAnonFiles(cfg, false)
	} else if cfg.Destination == particeps.Filebin {
		fmt.Println("particeps: uploading to https://filebin.com")
		res, err := particeps.FilebinUpload(cfg.Filename)
		assertNonNil(err)
		fmt.Printf("particeps: successfully uploaded \"%s\" to https://filebin.com\n", cfg.Filename)
		fmt.Printf("particeps: full-length link: %s\n", res.FullURL)
	}
}
