package models

import (
	"fmt"
	"strconv"
	"time"
)

type SensorData struct {
	Date time.Time `json:"date"`
	Data string    `json:"data"`
}

func (data *SensorData) String() string {
	return fmt.Sprintf("[%s]: %s", data.Date, data.Data)
}

func (data SensorData) GetData() (float64, error) {
	return strconv.ParseFloat(data.Data, 64)
}
