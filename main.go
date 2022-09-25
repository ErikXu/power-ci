package main

import (
	"bufio"
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("yum", "update", "-y")
	cmdReader, _ := cmd.StdoutPipe()
	scanner := bufio.NewScanner(cmdReader)
	done := make(chan bool)
	go func() {
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		done <- true
	}()
	cmd.Start()
	<-done
	err := cmd.Wait()
	if err != nil {
		fmt.Println(err.Error())
	}
}
