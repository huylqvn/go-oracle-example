package helpers

import (
	"time"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func TrackTime(s string, startTime *time.Time) {
	now := time.Now()
	t := now.Sub(*startTime)
	log.Info("TimeUsed: ", s+" ", t)
	*startTime = now
}
