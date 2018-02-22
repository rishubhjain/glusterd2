// Package device stores device information in the store
package device

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gluster/glusterd2/glusterd2/store"
	"github.com/gluster/glusterd2/pkg/api"
)

const (
	devicePrefix string = "devices/"
	peerPrefix string = "peers/"
)

// GetDevice returns devices of specified peer from the store
func GetDevice(peerid string) (*api.DeviceInfo, error) {
	resp, err := store.Store.Get(context.TODO(), peerPrefix+peerid+devicePrefix)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Printing Get Device %s", resp)
	if len(resp.Kvs) > 0 {
		var deviceInfo api.DeviceInfo
		fmt.Printf("Printing awesome %s", resp.Kvs[0].Value)
		if err := json.Unmarshal(resp.Kvs[0].Value, &deviceInfo); err != nil {
			return nil, err
		}
		return &deviceInfo, nil
	}

	return nil, nil
}

// AddOrUpdateDevice adds device to specific peer
func AddOrUpdateDevice(d *api.Device) error {
	json, err := json.Marshal(d.Detail)
	if err != nil {
		return err
	}

	idStr := d.PeerID.String()

	if _, err := store.Store.Put(context.TODO(), peerPrefix+idStr+devicePrefix, string(json)); err != nil {
		return err
	}

	return nil
}
