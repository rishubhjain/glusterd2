package peer

import (
        "fmt"
        "encoding/json"
        //restutils "github.com/gluster/glusterd2/glusterd2/servers/rest/utils"
	"github.com/gluster/glusterd2/glusterd2/gdctx"
        "github.com/gluster/glusterd2/pkg/api"
	config "github.com/spf13/viper"
)

// AddSelfDetails results in the peer adding its own details into etcd
func AddSelfDetails(req string) error {
        var v api.PeerAddReq
        fmt.Printf("Request %s", req)
        _ = json.Unmarshal([]byte(req), &v)
	p := &Peer{
		ID:        gdctx.MyUUID,
		Name:      gdctx.HostName,
		Addresses: []string{config.GetString("peeraddress")},
                PeerMetadata: v.PeerMetadata,
	}

	return AddOrUpdatePeer(p)
}
