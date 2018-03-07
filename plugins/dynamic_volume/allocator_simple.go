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
