package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
)

const URL = "http://whatthecommit.com/index.txt"

var files = flag.String("f", ".", "files to git add")
var branch = flag.String("b", "", "push to remote branch")

func commit(msg string) {
	flag.Parse()
	switch runtime.GOOS {
	case "windows":
		err := exec.Command("cmd", "/C", "git", "add", *files).Run()
		if err != nil {
			log.Fatal("Something wrong happened. Check your files list", err)
		}
		log.Println("Added files: ", *files)
		err = exec.Command("cmd", "/C", "git", "commit", "-m", msg).Run()
		if err != nil {
			log.Fatal(err)
		}
		err = exec.Command("cmd","/C", "git", "push", "origin", *branch).Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Pushed to: ", *branch)
		if *branch == "" {
			log.Println("default branch")
		}
	case "linux":
		err := exec.Command("bash", "git", "add", *files).Run()
		if err != nil {
			log.Fatal("Something wrong happened. Check your files list", err)
		}
		log.Println("Added files: ", *files)
		err = exec.Command("bash", "git", "commit", "-m", msg).Run()
		if err != nil {
			log.Fatal(err)
		}
		err = exec.Command("bash", "git", "push", "origin", *branch).Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Pushed to: ", *branch)
		if *branch == "" {
			log.Println("default branch")
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
