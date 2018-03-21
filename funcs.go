package go_parallel

import "github.com/sirupsen/logrus"

var (
	Err = make(chan error)
)

func init() {
	ReadErrors()
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func CheckError(err error) {
	if err != nil {
		logrus.Errorf("error found, exiting!!")
		Err <- err
	}
}

func ReadErrors() {
	go func() {
		defer close(Err)
		for {
			logrus.Fatal(<-Err)
		}
	}()
}
