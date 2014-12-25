MQTT Publish Hook for Logrus
========================================

.. image:: http://i.imgur.com/hTeVwmJ.png
   :target: https://github.com/Sirupsen/logrus

This is an `Logrus <https://github.com/Sirupsen/logrus>`_ Hook which can send log message to MQTT server with
operator friendly topic manner.

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

Topic
-------

messages are published to ``topic + / + level`` topic.

If you set topic ``logrusmqtt/log`` and pass to NewMQTTHook,

- info level log be sent to ``logrusmqtt/log/info``
- error level log will be sent to ``logrusmqtt/log/error``

You can subscribe any specific log level. Or if you want to get any of
log message, just subscribt ``logrusmqtt/log/#``.

Parameters
--------------

To create MQTTHook, use MQTTHookParams{}.

::

   type MQTTHookParams struct {
        Hostname   string
        Port       int
        Username   string
        Password   string
        QoS        int
        Topic      string
        Retain     bool
        ClientId   string
        CAFilepath string
        Insecure   bool
   }

Only Topic and Hostname is required. Other is optional.
