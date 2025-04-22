package model

type BatteryInfo struct {
	Level       int  `json:"level"`
	Charging    bool `json:"charging"`
	Temperature int  `json:"temperature"`
}
