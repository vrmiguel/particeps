/*
 * particeps
 * https://github.com/vrmiguel/particeps
 *
 * Copyright (c) 2020 Vinícius R. Miguel <vinicius.miguel at unifesp.br>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package cliargs

import (
	"fmt"
	"os"

	"../particeps"
)

const usage = "Usage: ./particeps [-h, --help] [-a, --anonfiles] [-F, --filebin] [-b, --bayfiles] -f, --filename path-to-file"

// CLIArgs stores the passed command-line options
type CLIArgs struct {
	Destination int
	Filename    string
}

func printHelp() {
	fmt.Printf("%-16s\tShow this help message and exit.\n", "-h, --help")
	fmt.Printf("%-16s\tUpload the file to anonfiles.com\n", "-a, --anonfiles")
	fmt.Printf("%-16s\tUpload the file to bayfiles.com\n", "-b, --bayfiles")
	fmt.Printf("%-16s\tUpload the file to filebin.net\n", "-F, --filebin")
	fmt.Printf("%-16s\tUpload the image to imagebin.net\n", "-F, --imagebin")
	fmt.Printf("%-16s\tIndicates the file to be uploaded.\n", "-f, --filename")
	fmt.Println(usage)
}

// ParseCLIArgs reads through the given CLI arg. list and builds a CliArgs
func ParseCLIArgs(args []string) CLIArgs {
	if len(args) == 1 {
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(1)
	}
	// SaveUserInfo, SaveFileInfo, SaveCPUInfo, SaveOutput, SaveInTomlFormat, TimeFormat
	var cfg CLIArgs
	for i := 1; i < len(args); i++ {
		arg := args[i]
		if arg == "-h" || arg == "--help" {
			printHelp()
			os.Exit(0)
		} else if arg == "-a" || arg == "--anonfiles" {
			cfg.Destination = particeps.AnonFiles
		} else if arg == "-b" || arg == "--bayfiles" {
			cfg.Destination = particeps.BayFiles
		} else if arg == "-F" || arg == "--filebin" {
			cfg.Destination = particeps.Filebin
		} else if arg == "-I" || arg == "--imagebin" {
			cfg.Destination = particeps.Imagebin
		} else if arg == "-f" || arg == "--filename" {
			if i+1 >= len(args) || args[i+1] == "" {
				fmt.Println("error: missing value to -f, --filename")
				fmt.Println(usage)
				os.Exit(1)
			}
			i++
			cfg.Filename = args[i]
		} else {
			fmt.Printf("error: unknown option %s\n", arg)
			os.Exit(1)
		}
	}
	if cfg.Destination == 0 {
		fmt.Fprintln(os.Stderr, "error: destination not supplied")
		fmt.Println(usage)
		os.Exit(1)
	}
	return cfg
}
