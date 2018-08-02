package main

import (
	"punching/logger"
	"punching/server"
	"time"
)

func main() {
	go startMoniter()
	server.Main()
	finishMoniter()
}

var (
	flushStart = 1
	flushStop  = 0
)

func startMoniter() {
	go logger.ThreadFlush()
	for {
		time.Sleep(time.Second)
		logger.FlushChannel <- flushStart
	}
}
func finishMoniter() {

	logger.FlushChannel <- flushStop
	_ = <-logger.ReturnChannel
}
