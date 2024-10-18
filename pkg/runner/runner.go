package runner

import "fmt"

var runners map[string]func()

func RegisterRunner(name string, r func()) {
	runners[name] = r
}

func RunAll() {
	if len(runners) == 0 {
		fmt.Println("NO RUNNERS!")
	}
	for _, r := range runners {
		r()
	}
}

func init() {
	runners = map[string]func(){}
}
