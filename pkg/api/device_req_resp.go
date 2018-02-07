package api

import "github.com/pborman/uuid"

// AddDeviceReq structure
type AddDeviceReq struct {
	PeerID uuid.UUID `json:"peer-id"`
	Names  []string  `json:"names"`
}

// Info structure
type Info struct {
	PeerID uuid.UUID `json:"peer-id"`
	Names  []string  `json:"names"`
	State  string    `json:"state"`
}
