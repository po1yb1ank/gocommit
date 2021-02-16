package profiler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type executor func(args ...string)

func initProfs() map[string]executor {
	profs := make(map[string]executor)
	profs["i"] = func(args ...string) {

		/* init local repo */
		err := Init(args[2])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		/* current path */
		path, err := os.Getwd()
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		log.Println("Successfully inited git repo in ", path)
		/* git add files */
		err = Add(args[0])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		/* git commit with message */
		err = Commit("init commit")
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		/* git remote add origin */
		err = Remote(args[2])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		/* git push origin branch */
		err = Push(args[1])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		log.Printf("Pushed to: %s", args[1])
		if args[1] == "" {
			log.Printf("default branch\n")
		}
		log.Println("Successfully committed and pushed changes for", args[0], "in branch", args[1])
	}
	profs["c"] = func(args ...string) {
		/* get commit message */
		msg := getCommitMsg()

		/* git add files
		err := Checkout(args[1])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		*/
		log.Println("Switched to: ", args[1])
		if args[1] == "" {
			log.Println("default branch")
		}
		/* git add files */
		err := Add(args[0])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		log.Println("Added files: ", args[0])

		/* git commit with message */
		err = Commit(msg)
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		/* push to remote */
		err = Push(args[1])
		if err != nil {
			log.Fatal("Something wrong happened.", err)
		}
		log.Printf("Pushed to: %s", args[1])
		if args[1] == "" {
			log.Println("default branch")
		}
		log.Println("Successfully committed and pushed changes for", args[0], "in branch", args[1])
	}

	return profs
}
func getCommitMsg() string {
	req, err := http.NewRequest("GET", "http://whatthecommit.com/index.txt", nil)
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
	return string(msg)
}
