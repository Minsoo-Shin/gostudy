package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "defaultfile", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
}

// sum 함수
func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	var ret int
	for _, v := range arr {
		ret += v
	}
	return ret
}
