package devicecommands

import (
	"os/exec"
	"strings"
	"fmt"

	"github.com/gluster/glusterd2/glusterd2/transaction"
	device "github.com/gluster/glusterd2/glusterd2/device"
	"github.com/gluster/glusterd2/pkg/api"
)

func txnPrepareDevice(c transaction.TxnCtx) error {
	var deviceinfo api.Device

	if err := c.Get("peerid", &deviceinfo.PeerID); err != nil {
		fmt.Printf("Failed Peerid")
		c.Logger().WithError(err).Error("Failed transaction, cannot find peer-id")
		return err
	}
	if err := c.Get("device-details", &deviceinfo.Detail); err != nil {
		fmt.Printf("Failed adding device")
		c.Logger().WithError(err).Error("Failed transaction, cannot find devicename")
		return err
	}
	fmt.Printf("Trdting 1 %s", deviceinfo.Detail)
	for _, element := range deviceinfo.Detail {
		fmt.Printf("Printing element %s", element)
		pvcreateCmd := exec.Command("pvcreate", "--metadatasize=128M", "--dataalignment=256K", string(element.Name))
		if err := pvcreateCmd.Run(); err != nil {
			c.Logger().WithError(err).Error("Failed transaction, pvcreate failed")
			return err
		}
		vgcreateCmd := exec.Command("vgcreate", strings.Replace("vg"+string(element.Name), "/", "-", -1), string(element.Name))
		if err := vgcreateCmd.Run(); err != nil {
			c.Logger().WithError(err).Error("Failed transaction, vgcreate failed")
			return err
		}
		element.State = device.DeviceEnabled
	}
	return nil
}
