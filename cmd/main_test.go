package main

import (
	"bytes"
	"io/ioutil"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
)

func Test_ReadFile(t *testing.T) {
	file, err := ioutil.TempFile("/tmp/", "paralleling")
	if err != nil {
		t.Error(err)
	}

	buf := bytes.NewBufferString("ls -la\npwd | pbcopy")

	file.Write(buf.Bytes())
	err = file.Close()
	if err != nil {
		t.Error(err)
	}

	commands := FileReader(file.Name())
	logrus.Info(commands)

	time.Sleep(100 * time.Millisecond) // why are these tests run as go routines?
}
