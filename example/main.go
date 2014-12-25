package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/shirou/logrusmqtt"
)

func main() {
	log := logrus.New()

	p := logrusmqtt.MQTTHookParams{
		Hostname: "test.mosquitto.org",
		Topic:    "logrusmqtt/test",
	}

	hook, err := logrusmqtt.NewMQTTHook(p, logrus.DebugLevel)
	if err != nil {
		panic(err)
	}
	log.Hooks.Add(hook)

	// publish to "logrusmqtt/test/info"
	log.Info("Info message")

	// publish to "logrusmqtt/test/error"
	log.WithFields(logrus.Fields{
		"name": "joe",
		"age":  42,
	}).Error("Error Message with fields")
}
