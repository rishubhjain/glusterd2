package peer

import (
	"github.com/gluster/glusterd2/glusterd2/gdctx"

	config "github.com/spf13/viper"
)

// AddSelfDetails results in the peer adding its own details into etcd
func AddSelfDetails(group int) error {
	p := &Peer{
		ID:        gdctx.MyUUID,
		Name:      gdctx.HostName,
		Addresses: []string{config.GetString("peeraddress")},
		Group:     group,
	}

	return AddOrUpdatePeer(p)
}
