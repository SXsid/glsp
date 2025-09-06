package main

import (
	"log"
	"os"

	dummycompiler "github.com/SXsid/glsp/DummyCompiler"
	"github.com/SXsid/glsp/server"
)

func main() {
	logger := getLogger("/home/shekhar/Personal/glsp/log.txt")
	state := dummycompiler.NewState()
	server := server.NewServer(state, logger, os.Stdin, os.Stdout)
	logger.Println("glsp is started")
	if err := server.Start(); err != nil {
		log.Fatalf("server failed %s", err)
	}
}

func getLogger(fileName string) *log.Logger {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("error while creating the log file ")
	}
	return log.New(file, "[ glsp ]", log.Ldate|log.Ltime|log.Lshortfile)
}
