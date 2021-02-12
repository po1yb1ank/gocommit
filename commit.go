package main

import (
	"flag"
	"github.com/po1yb1ank/gocommit/profiler"
	"log"
)

var profile = flag.String("p", "committer", "auto committer profile")
var files = flag.String("f", ".", "files for committer profile")
var branch = flag.String("b", "", "branch for committer profile")
var remote = flag.String("r", "", "remote repo for reminit")

func main() {
	flag.Parse()
	err := profiler.ExecProfile(*profile, *files, *branch, *remote)
	if err != nil{
		log.Fatal("Something wrong happened while executing ", *profile, "with args:",
			"\nFiles: ", *files,
			"\nBranch: ", *branch,
			"\nRemote: ", *remote)
	}
}
