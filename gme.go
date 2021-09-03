package gmeUtil

import (
	"fmt"
	"strconv"
	"net/http"
	"os/exec"
	"time"
)

func Run(cmd string) string {
	_, err := os.Command(cmd)
	if err != nil {
		fmt.Println(err)
		output := "There was a error running the command: "+ cmd +", the error will be outputed to the terminal above"
		return output
	}
	return "command was executed successfully"
}

func RunCmds(cmds []string) string {
	for k, v := range cmds {
		os.Command(cmd)
		fmt.Printf("outputted %d command: %s\n", k, v)
	}
	return "finished"
}

func Add(a, b int) int {
	s := a + b
	return s
}

func Multiply(a, b int) int {
	s := a * b
	return s
}

func GetMultiple(urls []string) string {
	for _, v := range urls {
		http.Get(v)
	}
	return "sent requests"
}

func Get(url string) {
	go func() {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println("there was a error sending the request")
		}
	}()
	fmt.Println("sent request successfully")
}

