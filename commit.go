package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

const URL = "http://whatthecommit.com/index.txt"

var files = flag.String("f", "nil", "files to add")
var branch = flag.String("b", "nil", "push to remote branch")

func commit(msg string) {
	flag.Parse()
	if *files == "nil" || *branch == "nil" {
		log.Println("Usage: -f main.go foo.go (files to git add)")
		log.Println("-b master (remote branch)")
		os.Exit(1)
	}
	switch runtime.GOOS {
	case "windows":
		err := exec.Command("cmd", "/C", "git", "add", *files).Run()
		if err != nil {
			log.Fatal("Something wrong happened. Check your files list", err)
		}
		err = exec.Command("cmd", "/C", "git", "commit", "-m", msg).Run()
		if err != nil {
			log.Fatal(err)
		}
		err = exec.Command("cmd","/C", "git", "push", "origin", *branch).Run()
		if err != nil {
			log.Fatal(err)
		}
	case "linux":
		err := exec.Command("bash", "git", "add", *files).Run()
		if err != nil {
			log.Fatal("Something wrong happened. Check your files list", err)
		}
		err = exec.Command("bash", "git", "commit", "-m", msg).Run()
		if err != nil {
			log.Fatal(err)
		}
		err = exec.Command("bash", "git", "push", "origin", *branch).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Println("Successfully committed changes for", *files)
}
func main() {
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Your commit message: ", string(msg))
	commit(string(msg))
}
