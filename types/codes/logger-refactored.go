package main

import (
	"log"
	"os"
)

type job struct {
	Command string
	*log.Logger
}

func main() {
	j := &job{"demo", log.New(os.Stderr, "Job: ", log.Ldate)}
	j.Print("starting now...")
}