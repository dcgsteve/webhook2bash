package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func handleRequests() {
	APIRouter := mux.NewRouter().StrictSlash(true)
	APIRouter.HandleFunc("/", defaultResponse).Methods("GET")
	APIRouter.HandleFunc("/trigger", cmdTrigger).Methods("GET")
	APIRouter.HandleFunc("/version", cmdVersion).Methods("GET")
	log.Fatal(http.ListenAndServe(":"+envPort, APIRouter))
}

func defaultResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ok")
}

func cmdVersion(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Webhook-To-Bash API Version: %s", API_VERSION)
}

func cmdTrigger(w http.ResponseWriter, r *http.Request) {
	if isAuthorised(r) {
		// create job ref
		job := strings.Replace(uuid.New().String(), "-", "", -1)
		fmt.Fprintf(w, "%s", job)
		go executeScript(job)
	} else {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Authorisation failed")
	}
}

func executeScript(job string) {
	// set up command
	cmd := exec.Command("/bin/bash", envBashfile, job)
	// set up output files
	outfile, err := os.Create(DATA_DIR + "/output/" + job + ".output")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()
	cmd.Stdout = outfile
	errfile, err := os.Create(DATA_DIR + "/output/" + job + ".err")
	if err != nil {
		panic(err)
	}
	defer errfile.Close()
	cmd.Stderr = outfile
	// run
	if err := cmd.Run(); err != nil {
		os.Create(DATA_DIR + "/output/" + job + ".failed")
	} else {
		os.Create(DATA_DIR + "/output/" + job + ".completed")
	}
}
