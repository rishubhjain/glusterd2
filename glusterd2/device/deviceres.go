package device

import "github.com/pborman/uuid"

const (
	// DeviceEnabled represents enabled
	DeviceEnabled = "Enabled"

	// DeviceFrozen represents frozen
	DeviceFrozen = "Frozen"

	// DeviceEvacuated represents evacuated
	DeviceEvacuated = "Evacuated"
)

// Device is the added device info
type Device struct {
        PeerID uuid.UUID `json:"peer-id"`
        Detail   []Info  `json:"device-details"`
}

// Info is the info of each device
type Info struct {
        Name   string   `json:"name"`
        State  string   `json:"state"`
}
