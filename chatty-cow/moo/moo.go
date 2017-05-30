package moo

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/pkg/errors"
	"github.com/rancher/go-rancher-metadata/metadata"
)

const (
	mdURL  = "http://rancher-metadata/2016-07-29"
	mooURL = "http://%v/moo"
)

func CreateCow(herd string, mooInterval int64) (*Mooer, error) {
	logrus.Infof("Creating cow")
	client, err := metadata.NewClientAndWait(mdURL)
	if err != nil {
		return nil, errors.Wrap(err, "Couldn't create metadata client")
	}
	return &Mooer{
		herd:     herd,
		interval: mooInterval,
		client:   client,
	}, nil
}

type Mooer struct {
	herd         string // Subtring to match in the name of other containers to moo at (ping)
	interval     int64
	lastMoo      bool
	previousMoos []bool
	client       metadata.Client
}

func (m *Mooer) IsHappyCow() bool {
	return m.lastMoo
}

func (m *Mooer) StartMooing() error {
	logrus.Infof("Starting moo")
	var otherCow metadata.Container
	me, err := m.client.GetSelfContainer()
	if err != nil {
		return errors.Wrap(err, "Couldn't get self")
	}

	for {
		time.Sleep(time.Second * time.Duration(m.interval))

		otherContainers, err := m.client.GetContainers()
		if err != nil {
			logrus.Errorf("Couldn't get containers %v", err)
			m.lastMoo = false
		}

		length := len(otherContainers)
		if length == 0 {
			continue
		}

		found := false
		for i := 0; i < 10; i++ {
			otherCow = otherContainers[rand.Int()%length]
			if strings.Contains(otherCow.Name, m.herd) && otherCow.Name != me.Name {
				found = true
				break
			}
		}

		if !found {
			continue
		}

		url := fmt.Sprintf(mooURL, otherCow.Name)
		logrus.Debugf("Mooing at %v", url)
		r, err := http.Get(url)
		if err != nil {
			logrus.Errorf("Couldn't hear other cow moo: %v", err)
			m.lastMoo = false
			continue
		}

		logrus.Debugf("Got response moo %v", r.StatusCode)
		if r.StatusCode >= 300 {
			logrus.Errorf("Bad response from other cow: %v %v", otherCow.PrimaryIp, r.StatusCode)
			m.lastMoo = false
			continue
		}

		m.lastMoo = true
	}
}
