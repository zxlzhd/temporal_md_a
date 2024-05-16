package main

import (
	"github.com/zxlzhd/temporal_md_a/t_start"
	"github.com/zxlzhd/temporal_md_a/t_worker"
	"os"
	"strings"
	"sync"
)

func determineRole() string {
	hostname := os.Getenv("HOSTNAME")
	if strings.HasSuffix(hostname, "-0") {
		return "leader"
	}
	return "follower"
}
func main() {
	var ww = sync.WaitGroup{}
	ww.Add(1)
	go func() {
		defer ww.Done()
		t_worker.MainW()
	}()
	if determineRole() == "leader" {
		ww.Add(1)
		go func() {
			defer ww.Done()
			t_start.MainS()
		}()
	}
	ww.Wait()
}
