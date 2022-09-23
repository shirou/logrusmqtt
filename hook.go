package logrusmqtt

import (
	"encoding/json"
	"strings"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

// MQTTHook to send logs to the MQTT topic.
type MQTTHook struct {
	client MQTT.Client
	levels []logrus.Level
	topic  string
	qos    int
	retain bool
}

// MQTTMsg to create JSON to send MQTT.
type MQTTMsg struct {
	Time  time.Time     `json:"time"`
	Level string        `json:"level"`
	Msg   string        `json:"msg"`
	Data  logrus.Fields `json:"data"`
}

type Message struct {
	Name string
	Body string
	Time int64
}

// NewMQTTHook creates new MQTT Hook.
func NewMQTTHook(params MQTTHookParams, level logrus.Level) (*MQTTHook, error) {
	opts, err := setMQTTOpts(params)
	if err != nil {
		return nil, err
	}

	levels := []logrus.Level{}
	for _, l := range []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
		logrus.TraceLevel,
	} {
		if l <= level {
			levels = append(levels, l)
		}
	}

	hook := MQTTHook{
		client: MQTT.NewClient(opts),
		levels: levels,
		topic:  params.Topic,
		qos:    params.QoS,
		retain: params.Retain,
	}

	if token := hook.client.Connect(); token.Wait() && token.Error() != nil {
		return nil, err
	}

	return &hook, nil
}

// Fire sends mqtt msg.
func (hook *MQTTHook) Fire(entry *logrus.Entry) error {
	level := entry.Level.String()
	msg := MQTTMsg{
		Time:  entry.Time.UTC(),
		Level: level,
		Msg:   entry.Message,
		Data:  entry.Data,
	}

	payload, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	topic := strings.Join([]string{hook.topic, level}, "/")

	hook.client.Publish(topic, byte(hook.qos), hook.retain, string(payload))
	// no blocking here

	return nil
}

// Levels returns the list of logging levels that we want to send.
func (hook *MQTTHook) Levels() []logrus.Level {
	return hook.levels
}
