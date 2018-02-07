package devicecommands

import (
	"os/exec"
	"strings"

	"github.com/gluster/glusterd2/glusterd2/transaction"
	"github.com/gluster/glusterd2/pkg/api"
)

func txnPrepareDevice(c transaction.TxnCtx) error {
	var deviceinfo api.Info

	if err := c.Get("peerid", &deviceinfo.PeerID); err != nil {
		c.Logger().WithError(err).Error("Failed transaction, cannot find peer-id")
		return err
	}
	if err := c.Get("names", &deviceinfo.Names); err != nil {
		c.Logger().WithError(err).Error("Failed transaction, cannot find devicename")
		return err
	}
	for _, element := range deviceinfo.Names {
		pvcreateCmd := exec.Command("pvcreate", "--metadatasize=128M", "--dataalignment=256K", element)
		if err := pvcreateCmd.Run(); err != nil {
			c.Logger().WithError(err).Error("Failed transaction, pvcreate failed")
			return err
		}
		vgcreateCmd := exec.Command("vgcreate", strings.Replace("vg"+element, "/", "-", -1), element)
		if err := vgcreateCmd.Run(); err != nil {
			c.Logger().WithError(err).Error("Failed transaction, vgcreate failed")
			return err
		}
	}
	return nil
}
