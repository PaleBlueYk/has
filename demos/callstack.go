package main

import (
	"fmt"
	"path"
	"runtime"
	"strconv"
)

func main() {

	fmt.Printf("%016x\r\n\n", 15)

	_, file, line, ok := runtime.Caller(0)
	if !ok {
		file = "???"
		line = 0
	}
	_, filename := path.Split(file)
	msg := "hello"
	msg = "[" + filename + ":" + strconv.Itoa(line) + "] " + msg
	fmt.Println(msg)
}