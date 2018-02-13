package volumecommands


import (
	"fmt"
	"sync"
	peer "github.com/gluster/glusterd2/glusterd2/peer"
)


// Simple Devices so that we have no pointers and no race conditions
type SimpleDevices []SimpleDevice

// Elements in the balanced list
type SimpleDevice struct {
	zone            int
	nodeId		string
	deviceid	string
}

type SimpleAllocatorRing struct {

	// Map [zone] to [node] to slice of SimpleDevices
	ring         map[int]map[string][]*SimpleDevice
	balancedList SimpleDevices
}

// Create a new simple ring
func NewSimpleAllocatorRing() *SimpleAllocatorRing {
	s := &SimpleAllocatorRing{}
	s.ring = make(map[int]map[string][]*SimpleDevice)

	return s
}

// Simple allocator contains a map to rings of clusters
type SimpleAllocator struct {
        rings map[string]*SimpleAllocatorRing
        lock  sync.Mutex
}

// Create a new simple allocator
func NewSimpleAllocator() *SimpleAllocator {
	s := &SimpleAllocator{}
	s.rings = make(map[string]*SimpleAllocatorRing)
	return s
}


// Create a new simple allocator and initialize it with data
func NewSimpleAllocatortest() *SimpleAllocator {

	s := NewSimpleAllocator()
	peerIds, _ := peer.GetPeerIDs()
	for index, peerid := range peerIds {
		fmt.Printf("Peer id ^^^^^^^^^ %s", peerid)
		fmt.Printf("Index %s $$$$$$$$ ",index)
				/*node, err := NewNodeEntryFromId(tx, nodeId)
				if err != nil {
					return err
				}

				// Check node is online
				if !node.isOnline() {
					continue
				}

				for _, deviceId := range node.Devices {
					device, err := NewDeviceEntryFromId(tx, deviceId)
					if err != nil {
						return err
					}

					// Check device is online
					if !device.isOnline() {
						continue
					}

					// Add device to ring
					err = s.AddDevice(cluster, node, device)
					if err != nil {
						return err
					}

				}
			}
		}
		return nil
	})
	if err != nil {
		return nil
	}

	return s
*/
}
return s
}

