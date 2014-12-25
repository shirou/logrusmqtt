package logrusmqtt

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/Sirupsen/logrus"
)

func Test_MarshalEntry(t *testing.T) {
	entry := logrus.Entry{
		Time:    time.Now(),
		Level:   logrus.WarnLevel,
		Message: "test",
		Data: logrus.Fields{
			"key": "value",
		},
	}

	msg := MQTTMsg{
		Time:  entry.Time.UTC(),
		Level: entry.Level.String(),
		Msg:   entry.Message,
		Data:  entry.Data,
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		t.Error(err)
	}
	if string(payload) == "{}" {
		t.Error("could not marshal")
	}

}
