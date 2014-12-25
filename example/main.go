package main

import (
	"github.com/Sirupsen/logrus"

	"github.com/shirou/logrusmqtt"
)

func main() {
	log := logrus.New()

	p := logrusmqtt.MQTTHookParams{
		Hostname: "test.mosquitto.org",
		Topic:    "logrusmqtt/test/error",
	}

	hook, err := logrusmqtt.NewMQTTHook(p, logrus.DebugLevel)
	if err != nil {
		panic(err)
	}
	log.Hooks.Add(hook)

	log.Info("Info message")

	log.WithFields(logrus.Fields{
		"name": "joe",
		"age":  42,
	}).Error("Error Message with fields")
}
