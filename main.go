package main

import (
	"context"
	"fmt"
	"github.com/Lebeau-Alexandre/oxygen-cs-grp01-eq08/config"
	"github.com/philippseith/signalr"
)

type SrConfig struct {
	config.OxygenConfig
	receiver receiver
	client   signalr.Client
}
type receiver struct {
	signalr.Hub
}

type HVACResponse struct {
	// Define the structure of the response JSON here
	// For example, assuming the response has a field named "Result"
	Result string `json:"Response"`
	// Add more fields as needed
}

func (m *SrConfig) setup() {
	conf := config.GetOxygenConfig()
	m.Host = conf.Host
	m.Token = conf.Token
	m.TMax = conf.TMax
	m.TMin = conf.TMin
}

var srContext *SrConfig

func main() {
	config.InitConfig()
	go func() {
		srContext = &SrConfig{}
		srContext.setup()
		srContext.setSensorHub()
	}()
	select {}
}

func (m *SrConfig) setSensorHub() {
	route := fmt.Sprintf("%s/SensorHub?token=%s", m.Host, m.Token)
	m.receiver = receiver{}

	conn, _ := signalr.NewHTTPConnection(context.TODO(), route)
	m.client, _ = signalr.NewClient(
		context.TODO(),
		signalr.WithConnection(conn),
		signalr.WithReceiver(&m.receiver),
	)

	// Start the client loop
	m.client.Start()
}
