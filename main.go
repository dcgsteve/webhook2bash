package main

import (
	"fmt"
	"os"
)

const API_VERSION = "0.1"
const DATA_DIR = "/opt/wtb"

var envTag string
var envValidToken string
var envPort string
var envBashfile string

func main() {

	// get required service info from environment
	envTag = os.Getenv("WTB_HEADER_TAG")
	envValidToken = os.Getenv("WTB_TOKEN")
	envPort = os.Getenv("WTB_PORT")
	envBashfile = os.Getenv("WTB_BASHFILE")

	// Any info we can default?
	if envPort == "" {
		envPort = "80"
	}
	if envTag == "" {
		envTag = "X-Gitlab-Token"
	}

	// Any other missing mandatory info?
	if envValidToken == "" {
		fmt.Println("Listener not started - no WTB_TOKEN variable specified!")
		os.Exit(1)
	}
	if envBashfile == "" {
		fmt.Println("Listener not started - no WTB_BASHFILE variable specified!")
		os.Exit(1)
	}

	// Ok, go ahead
	handleRequests()
}
