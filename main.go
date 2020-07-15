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
	fmt.Println(website)
	res, err := uploadFunction(cfg.Filename)
	assertNonNil(err)
	fmt.Printf("particeps: successfully uploaded \"%s\" to %s/\n", cfg.Filename, website)
	fmt.Printf("particeps: full-length link: %s\n", res.FullURL)
	fmt.Printf("particeps: short link: %s\n", res.ShortURL)
}

func main() {
	cfg := cliargs.ParseCLIArgs(os.Args)
	fmt.Printf("config folder: %s\n", particeps.GetPrefFolder())
	fileSize, err := particeps.CheckFile(cfg.Filename)
	assertNonNil(err)
	fmt.Printf("particeps: file \"%s\" has size %s\n", cfg.Filename, fileSize)
	fmt.Printf("particeps: uploading to ")
	switch cfg.Destination {
	case particeps.AnonFiles:
		helperAnonFiles(cfg, true)
	case particeps.BayFiles:
		helperAnonFiles(cfg, false)
	case particeps.Filebin:
		fmt.Println("https://filebin.com")
		res, err := particeps.FilebinUpload(cfg.Filename)
		assertNonNil(err)
		fmt.Printf("particeps: successfully uploaded \"%s\" to https://filebin.com\n", cfg.Filename)
		fmt.Printf("particeps: full-length link: %s\n", res.FullURL)
		fmt.Println("particeps: bear in mind that Filebin only stores the files for a week.")
	case particeps.Imagebin:
		fmt.Println("http://imagebin.ca")
		fmt.Println("particeps: warning - Imagebin support is unstable and experimental")
		res, err := particeps.ImagebinUpload(cfg.Filename)
		assertNonNil(err)
		fmt.Printf("particeps: full-length link: %s\n", res.FullURL)
	}
}
