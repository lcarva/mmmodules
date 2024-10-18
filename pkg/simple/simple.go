package simple

import (
	"fmt"

	"github.com/lcarva/mmmodules/pkg/runner"
)

func Run() {
	fmt.Println("run: simple")
}

func init() {
	runner.RegisterRunner("simple", Run)
}
