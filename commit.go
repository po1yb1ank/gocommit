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

const url = "http://whatthecommit.com/index.txt"
var files = flag.String("f", "nil", "files to add")
func commit(msg string){
	flag.Parse()
	if *files == "nil"{
		log.Println("Usage: -f main.go foo.go (files to git add)")
		os.Exit(1)
	}
	switch runtime.GOOS{
	case "windows":
		err := exec.Command("cmd", "/C","git", "add", *files).Run()
		if err != nil{
			log.Fatal("Something wrong happened. Check your files list", err)
		}
		err = exec.Command("cmd", "/C","git", "commit", "-m", msg).Run()
		if err != nil{
			log.Fatal(err)
		}
	case "linux":
		err := exec.Command("bash", "git", "add", *files).Run()
		if err != nil{
			log.Fatal("Something wrong happened. Check your files list", err)
		}
		err = exec.Command("bash", "git", "commit", "-m", msg).Run()
		if err != nil{
			log.Fatal(err)
		}
	}
}
func main(){
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatal(err)
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer res.Body.Close()
	msg, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatal(err)
	}
	commit(string(msg))
}



