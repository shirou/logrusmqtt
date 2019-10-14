package logrusmqtt

import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

const MaxClientIdLen = 8

// MQTTHookParams set MQTTHook instance.
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

// setMQTTOpts create MQTT client options from passed parameters.
func setMQTTOpts(p MQTTHookParams) (*MQTT.ClientOptions, error) {
	opts := MQTT.NewClientOptions()

	if p.Topic == "" {
		return nil, fmt.Errorf("topic required")
	}
	if p.Hostname == "" {
		return nil, fmt.Errorf("hostname required")
	}

	if p.Port == 0 {
		p.Port = 1883
	}

	// clientId
	if p.ClientId == "" {
		p.ClientId = getRandomClientId()
	}
	opts.SetClientID(p.ClientId)

	if p.Username != "" {
		opts.SetUsername(p.Username)
	}
	if p.Password != "" {
		opts.SetPassword(p.Password)
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: false}
	cafile := p.CAFilepath
	scheme := "tcp"
	if cafile != "" {
		scheme = "ssl"
		certPool, err := getCertPool(cafile)
		if err != nil {
			return nil, err
		}
		tlsConfig.RootCAs = certPool
	}
	if p.Insecure {
		tlsConfig.InsecureSkipVerify = true
	}
	opts.SetTLSConfig(tlsConfig)

	brokerUri := fmt.Sprintf("%s://%s:%d", scheme, p.Hostname, p.Port)

	opts.AddBroker(brokerUri)

	return opts, nil
}

// getRandomClientId generates random ClientId.
func getRandomClientId() string {
	const alphanum = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, MaxClientIdLen)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return "logrusmqtt-" + string(bytes)
}

func getCertPool(pemPath string) (*x509.CertPool, error) {
	certs := x509.NewCertPool()

	pemData, err := ioutil.ReadFile(pemPath)
	if err != nil {
		return nil, err
	}
	certs.AppendCertsFromPEM(pemData)
	return certs, nil
}
