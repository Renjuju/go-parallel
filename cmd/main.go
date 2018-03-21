package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
)

type TerminalCommand struct {
	name string
	args string
}

type command func(cmd TerminalCommand) chan interface{}

func jobs(cmds ...TerminalCommand) {
	var chans []chan interface{}

	for _, cmd := range cmds {
		chans = append(chans, execCommand(cmd.name, cmd.args))
	}

	for _, ch := range chans {
		fmt.Sprint(<-ch)
	}
}

func main() {
	jobs(TerminalCommand{
		name: "docker",
		args: "build -f docker/Dockerfile .",
	})
}

func execCommand(name string, args string) chan interface{} {
	ch := make(chan interface{})
	arguments := strings.Split(args, " ")
	go func() {
		defer close(ch)
		cmd := exec.Command(name, arguments...)
		cmd.Dir = "/Users/radhakrishnanr1/go/src/github.com/renjuju/go-parallel"

		cmd.Stdout = os.Stdout
		err := cmd.Start()
		if err != nil {
			ch <- fmt.Sprintf("Finished with %v", err)
		}
		err = cmd.Wait()
		logrus.Infof("Finished with %v", err)
		ch <- fmt.Sprintf("Finished with %v", err)
	}()

	return ch
}
