package api

import "github.com/pborman/uuid"

// AddDeviceReq structure
type AddDeviceReq struct {
	PeerID uuid.UUID `json:"peer-id"`
	Names  []string  `json:"names"`
}

// Device is the added device info
type Device struct {
        PeerID uuid.UUID `json:"peer-id"`
        Detail   []Info  `json:"device-details"`
        Names   string   `json:"name"`
        State   string   `json:"state"`
}

// Info is the info of each device
type Info struct {
        Name   string   `json:"name"`
        State  string   `json:"state"`
}
