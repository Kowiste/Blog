package main

import (
	"flag"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	configuration := readParameter()
	mqtt.NewClient(createFunctOptMQTT(configuration))
}

func createFunctOptMQTT(conf Config) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions().AddBroker(conf.IP)
	opts.SetKeepAlive(2 * time.Second)
	opts.SetAutoReconnect(true)
	return opts
}

func readParameter() Config {
	conf := Config{}
	broker := flag.String("ip", "tcp://127.0.0.1:1883", "IP of the ")
	interval := flag.Int("i", 1000, "interval between keepalive in millisecond")

	flag.Parse()
	conf.IP = *broker
	conf.interval = *interval
	return conf
}

type Config struct {
	IP       string
	interval int
}
