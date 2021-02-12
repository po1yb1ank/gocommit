package profiler

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func Add(files string)error{
	if files == ""{
		files = "."
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		err := exec.Command("bash", "git", "add", files).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	case "windows":
		err := exec.Command("cmd", "/C", "git", "add", files).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	}
	return nil
}
func Commit(msg string)error{
	switch runtime.GOOS {
	case "linux", "darwin":
		err := exec.Command("bash", "git", "commit", "-m", msg).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	case "windows":
		err := exec.Command("cmd", "/C", "git", "commit", "-m", msg).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	}
	return nil
}
func Init(branch string)error{
	if branch == ""{
		branch = "master"
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		err := exec.Command("bash", "git", "init", "-b", branch).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	case "windows":
		err := exec.Command("cmd", "/C", "git", "init", "-b", branch).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	}
	return nil
}
func Push(branch string) error{
	if branch == ""{
		branch = "master"
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		err := exec.Command("bash", "git", "push", "origin", branch).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	case "windows":
		err := exec.Command("cmd", "/C", "git", "push", "origin", branch).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	}
	return nil
}
func Remote(URL string) error{
	if URL == ""{
		return fmt.Errorf("empty remote url")
	}
	switch runtime.GOOS {
	case "linux", "darwin":
		err := exec.Command("bash", "git", "remote", "add", "origin", URL).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	case "windows":
		err := exec.Command("cmd", "/C", "git", "remote", "add", "origin", URL).Run()
		if err != nil {
			log.Fatal("Something wrong happened while git init.", err)
			return err
		}
	}
	return nil
}
