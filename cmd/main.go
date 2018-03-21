package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/renjuju/go-parallel"
	"github.com/sirupsen/logrus"
)

type TerminalCommand struct {
	name string
	args string
}

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

func FileReader(filePath string) []TerminalCommand {
	file, err := os.Open(filePath)
	go_parallel.CheckError(err)

	data, err := ioutil.ReadAll(file)
	go_parallel.CheckError(err)

	str := string(data)
	commands := strings.Split(str, "\n")

	var tc []TerminalCommand
	for _, command := range commands {
		tokens := strings.Split(command, " ")

		tc = append(tc, TerminalCommand{
			name: tokens[0],
			args: strings.Join(tokens[1:], " "),
		})
	}

	return tc
}
