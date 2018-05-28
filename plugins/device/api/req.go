package device

const (
	// DeviceEnabled represents enabled
	DeviceEnabled = "Enabled"

	// DeviceDisabled represents disabled
	DeviceDisabled = "Disabled"
)

// DeviceStates represents states of devices
var DeviceStates = map[string]string{
	"enabled":  DeviceEnabled,
	"disabled": DeviceDisabled,
}

// AddDeviceReq structure
type AddDeviceReq struct {
	Device string `json:"device"`
}

// EditDeviceStateReq structure
type EditDeviceStateReq struct {
	DeviceName string `json:"device-name"`
	State      string `json:"state"`
}
