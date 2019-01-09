package main

import (
	"flag"
	"log"
	"os"

	"github.com/hashicorp/go-getter"
)

func main() {
	modeRaw := flag.String("mode", "any", "get mode (any, file, dir)")
	progress := flag.Bool("progress", false, "display terminal progress")
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		log.Fatalf("Expected two args: URL and dst")
		os.Exit(1)
	}

	// Get the mode
	var mode getter.ClientMode
	switch *modeRaw {
	case "any":
		mode = getter.ClientModeAny
	case "file":
		mode = getter.ClientModeFile
	case "dir":
		mode = getter.ClientModeDir
	default:
		log.Fatalf("Invalid client mode, must be 'any', 'file', or 'dir': %s", *modeRaw)
		os.Exit(1)
	}

	// Get the pwd
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting wd: %s", err)
		os.Exit(1)
	}

	// Build the client
	client := &getter.Client{
		Src:  args[0],
		Dst:  args[1],
		Pwd:  pwd,
		Mode: mode,
	}
	var opts []getter.ClientOption
	if *progress {
		opts = append(opts, getter.WithProgress(defaultProgressBar))
	}

	if err := client.Configure(opts...); err != nil {
		log.Fatalf("Configure: %s", err)
	}

	if err := client.Get(); err != nil {
		log.Fatalf("Error downloading: %s", err)
	}

	log.Println("Success!")
}
