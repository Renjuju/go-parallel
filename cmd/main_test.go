package main

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test_ReadFile(t *testing.T) {
	commands := FileReader("/Users/radhakrishnanr1/go/src/github.com/renjuju/go-parallel/test.txt")
	logrus.Info(commands)
}
