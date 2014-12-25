MQTT Hook for Logrus
========================================

.. image:: http://i.imgur.com/hTeVwmJ.png
   :target: https://github.com/Sirupsen/logrus


Usage
------------

::
   
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


And the output from MQTT server is,

::
   
   {"time":"2014-12-25T02:29:54.140874274Z","level":"info","msg":"Info message","data":{}}
   {"time":"2014-12-25T02:29:54.141082554Z","level":"error","msg":"Error Message with fields","data":{"age":42,"name":"joe"}}
