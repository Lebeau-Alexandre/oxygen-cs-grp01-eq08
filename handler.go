package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Lebeau-Alexandre/oxygen-cs-grp01-eq08/config"
	"github.com/Lebeau-Alexandre/oxygen-cs-grp01-eq08/models"
	"log"
	"net/http"
	"time"
)

func (c *receiver) ReceiveSensorData(msg models.SensorData) error {
	temp, err := msg.GetData()
	if err != nil {
		return err
	}
	srContext.takeAction(temp)

	db := config.GetPostgresConfig().Context

	var id int
	// Insert data into the "oxygen.temperature" table

	return db.QueryRow(`INSERT INTO oxygen.temperature (tempValue, tempDate) VALUES ($1, $2) RETURNING id`,
		temp,
		time.Now(),
	).Scan(&id)
}

func (m *SrConfig) takeAction(temperature float64) {
	var err error
	if temperature >= m.TMax {
		if err := m.sendActionToHVAC("TurnOnAc"); err == nil {
			err = saveEventToDB(-1)
		}
	} else if temperature <= m.TMin {
		if err := m.sendActionToHVAC("TurnOnHeater"); err == nil {
			err = saveEventToDB(1)
		}
	}
	if err != nil {
		log.Print(err.Error())
	}
}

func (m *SrConfig) sendActionToHVAC(action string) error {
	url := fmt.Sprintf("%s/api/hvac/%s/%s/%d", m.Host, m.Token, action, 1)

	// Send GET request to the specified URL
	response, err := http.Get(url)
	if err != nil {
		return errors.New("Couldn't get resource :(")
	}
	defer response.Body.Close()

	// Decode the JSON response
	var hvacResponse HVACResponse
	if err := json.NewDecoder(response.Body).Decode(&hvacResponse); err != nil {
		return errors.New("Oh shit, Mikey YACKson")
	}

	// Print the response details
	fmt.Printf("%+v\n", hvacResponse)
	return nil
}

func saveEventToDB(eventType int) error {
	db := config.GetPostgresConfig().Context

	var id int
	// Insert event data into the "oxygen.hvacEvent" table

	return db.QueryRow(`INSERT INTO oxygen.hvacEvent (eventType, eventDate) VALUES ($1, $2) RETURNING id`,
		eventType,
		time.Now(),
	).Scan(&id)

}
