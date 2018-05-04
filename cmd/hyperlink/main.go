package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/haya14busa/hyperlink"
)

func usage() {
	fmt.Fprintln(os.Stderr, `USAGE:	hyperlink <link> [<text>]
	hyperlink given link and text with OSC8 feature.

	GitHub: https://github.com/haya14busa/hyperlink`)
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if err := run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func run() error {
	args := flag.Args()
	if len(args) == 0 || len(args) > 2 {
		usage()
		os.Exit(2)
	}
	link := args[0]
	text := link
	if len(args) == 2 {
		text = args[1]
	}
	return hyperlink.Write(os.Stdout, link, text)
}
