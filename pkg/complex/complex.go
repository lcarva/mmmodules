package complex

import (
	"fmt"

	"github.com/lcarva/mmmodules/pkg/runner"
	_ "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/tools/clientcmd"
)

func Run() {
	fmt.Println("run: complex")
}

func init() {
	runner.RegisterRunner("complex", Run)
}
