package commitstorage

import (
	"github.com/sdslabs/kiwi/stdkiwi"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func MakeStorage(){
	store := stdkiwi.NewStore()
	for i := 0; i < 10; i++ {
		res, err := http.Get("http://whatthecommit.com/index.txt")
		if err != nil {
			log.Fatal()
			return
		}
		if err := store.AddKey(strconv.Itoa(i), "str"); err != nil {
			panic(err)
		}
		comm := store.Str(strconv.Itoa(i))
		msg, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = comm.Update(string(msg))
		if err != nil {
			log.Fatal(err)
			return
		}
		err = res.Body.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
		log.Println("written", i, "to file")
	}
	js, err :=store.Export()
	if err != nil {
		log.Fatal(err)
		return
	}
	err = ioutil.WriteFile("commits.json", js, os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
}
