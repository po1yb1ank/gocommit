package profiler

import (
	"fmt"
	"log"
)

//ExecProfile will execute profile...kinda....yep
func ExecProfile(args ...string) error {

	p := initProfs()
	if _, ok := p[args[0]]; !ok {
		return fmt.Errorf("no such profile")
	}
	log.Println("profile:", args[0])
	p[args[0]](args[1:]...)
	return nil
}
