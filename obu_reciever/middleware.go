package main

import (
	"github.com/sirupsen/logrus"
	"hotel_train_antonyGG/types"
	"time"
)

type LogMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {

	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuId": data.OBUID,
			"lat":   data.Lat,
			"long":  data.Long,
			"took":  time.Since(start),
		}).Info("producing to kafka")
	}(time.Now())

	return l.next.ProduceData(data)
}
