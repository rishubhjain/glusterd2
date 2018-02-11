package device

// This file contains helper functions facilitate easier interaction with the
// peer information stores in the store

import (
	"fmt"
	"context"
        "encoding/json"

        "github.com/gluster/glusterd2/glusterd2/store"
	"github.com/gluster/glusterd2/pkg/api"
)

const (
	devicePrefix string = "devices/"
)


// GetDevice returns devices of specified peer from the store
func GetDevice(peerid string) (*api.Device, error) {
	resp, err := store.Store.Get(context.TODO(), devicePrefix+peerid)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Printing response %s", resp)
	fmt.Printf("Printing details %s", resp.Kvs)
	if len(resp.Kvs) > 0 {
		var deviceDetail api.Device
        	if err := json.Unmarshal(resp.Kvs[0].Value, &deviceDetail); err != nil {
                	return nil, err
        	}
        	fmt.Printf("Printing Device info %s", deviceDetail)
        	return &deviceDetail, nil
	} else {
		fmt.Printf("Printing Data")
		return nil, nil
	}
}

func AddOrUpdateDevice(d *api.Device) error {
	json, err := json.Marshal(d)
	if err != nil {
		return err
	}

	idStr := d.PeerID.String()

	if _, err := store.Store.Put(context.TODO(), devicePrefix+idStr, string(json)); err != nil {
		return err
	}

	return nil
}
