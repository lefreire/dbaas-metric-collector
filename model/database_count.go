package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type DatabaseCount struct {
	Moment time.Time
	Count  int
}

func DatabaseCounterAdd(connection Connection, moment time.Time, count int) {
	collection := connection.Database.C("DatabaseCounter")
	err := collection.Insert(&DatabaseCount{moment, count})
	if err != nil {
		panic(err)
	}
}

func DatabaseCounterGet(connection Connection, size int) []DatabaseCount {
	counters := []DatabaseCount{}
	collection := connection.Database.C("DatabaseCounter")
	err := collection.Find(bson.M{}).All(&counters)
	if err != nil {
		panic(err)
	}

	diff := len(counters) - size
	if diff > 0 {
		counters = counters[diff:]
	}
	return counters
}
