package profiler

import "fmt"

func ExecProfile(args ...string)error{

	p := initProfs()
	if _,ok := p[args[0]]; !ok{
		return fmt.Errorf("no such profile")
	}
	p[args[0]](args[1:]...)
	return nil
}
