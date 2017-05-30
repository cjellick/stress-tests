package main

import (
	"net/http"
	"os"

	"math/rand"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/rancher/stress-tests/chatty-cow/moo"
	"github.com/urfave/cli"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "chatty-cow"
	app.Version = VERSION
	app.Usage = "Talks to other chatty cows and reports the health of its conversations."
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	rand.Seed(time.Now().UnixNano())

	mooInterval := c.Int64("moo-interval")
	if mooInterval == 0 {
		mooInterval = 1
	}

	if c.Bool("debug") {
		logrus.SetLevel(logrus.DebugLevel)
	}

	herd := c.String("herd")
	if herd == "" {
		herd = "chatty-cow"
	}

	cow, err := moo.CreateCow(herd, mooInterval)
	if err != nil {
		logrus.Fatalf("Unrecoverable error: %v", err)
	}

	s := server{
		cow: cow,
	}

	http.HandleFunc("/healthcheck", s.healthcheck)
	http.HandleFunc("/moo", s.mooHandler)
	go http.ListenAndServe(":80", nil)

	err = cow.StartMooing()
	return err
}

type server struct {
	cow *moo.Mooer
}

func (s *server) mooHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (s *server) healthcheck(w http.ResponseWriter, r *http.Request) {
	if s.cow.IsHappyCow() {
		logrus.Infof("Healthcheck reply 200")
		w.WriteHeader(200)
		return
	}

	logrus.Infof("Healthcheck reply 500")
	w.WriteHeader(500)
}
