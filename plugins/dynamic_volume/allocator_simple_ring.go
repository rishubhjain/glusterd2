package dynamic_volume


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

// Elements in the balanced list
type SimpleDevice struct {
	zone             int
	peerId, deviceId string
}

// Pretty pring a SimpleDevice
func (s *SimpleDevice) String() string {
	return fmt.Sprintf("{Z:%v P:%v D:%v}",
		s.zone,
		s.peerId,
		s.deviceId)
}

// Simple Devices so that we have no pointers and no race conditions
type SimpleDevices []SimpleDevice

// A peer is a collection of devices
type SimplePeer []*SimpleDevice

// A zone is a collection of peers
type SimpleZone []SimplePeer

// The allocation ring will contain a map composed of all
// the devices available in the cluster.  Call Rebalance()
// for it to create a balanced list.
type SimpleAllocatorRing struct {

	// Map [zone] to [peer] to slice of SimpleDevices
	ring         map[int]map[string][]*SimpleDevice
	balancedList SimpleDevices
}

// Create a new simple ring
func NewSimpleAllocatorRing() *SimpleAllocatorRing {
	s := &SimpleAllocatorRing{}
	s.ring = make(map[int]map[string][]*SimpleDevice)

	return s
}
