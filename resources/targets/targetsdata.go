package targets

import (
	"github.com/johnmcdnl/darts/resources/data"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"sync"
	"time"
)

var once sync.Once
var collection *mgo.Collection

const dbTableName string = "targets"

func targetsDb() *mgo.Collection {
	once.Do(func() {
		collection = data.Database().C(dbTableName)
	})
	return collection
}

func createTarget(t *Target) (*Target, error) {
	now := time.Now()
	t.Date = &now

	percentage := float64(*t.Success) / float64(*t.Attempts) * float64(100)
	t.Percentage = &percentage

	err := targetsDb().Insert(t)

	return t, err
}

func retrieveAllTargets() (*Targets, error) {
	return findTargetsBy(nil)
}

func retrieveAllTargetsByTargetName(targetName string) (*Targets, error) {
	return findTargetsBy(bson.M{"targetname": targetName})
}

func retrieveAllTargetsByUsername(username string) (*Targets, error) {
	return findTargetsBy(bson.M{"username": username})
}

func retrieveAllTargetsByUsernameAndTargetName(username, targetName string) (*Targets, error) {
	return findTargetsBy(bson.M{"username": username, "targetname": targetName})
}

func findTargetsBy(m bson.M) (*Targets, error) {
	var results []Target
	if err := targetsDb().Find(m).All(&results); err != nil {
		return nil, err
	}
	return makeTargets(results)
}

func makeTargets(targets []Target) (*Targets, error) {
	var t Targets
	t.Targets = targets
	if err := t.generateTargetsAnalysis(); err != nil {
		return nil, err
	}
	return &t, nil
}
