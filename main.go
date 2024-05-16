package main

import (
	"fmt"
	"github.com/zxlzhd/temporal_md_a/t_start"
	"github.com/zxlzhd/temporal_md_a/t_worker"
	"os"
	"strings"
)

func determineRole() string {
	hostname := os.Getenv("HOSTNAME")
	if strings.HasSuffix(hostname, "-0") {
		return "leader"
	}
	return "follower"
}
func main() {
	fmt.Println("runing")
	if determineRole() == "leader" {
		fmt.Println("start run")
		t_start.StartRun()
	}
	//go func() {
	fmt.Println("worker run")
	t_worker.WorkerRun()
	for {
	}
	//}()
}
