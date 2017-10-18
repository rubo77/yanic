package all

import (
	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
)

func (f filterConfig) NoOwner() bool {
	if v, ok := f["no_owner"]; ok {
		return v.(bool)
	}
	return true
}

func filterNoOwner(node *runtime.Node) *runtime.Node {
	if nodeinfo := node.Nodeinfo; nodeinfo != nil {
		return &runtime.Node{
			Address:    node.Address,
			Firstseen:  node.Firstseen,
			Lastseen:   node.Lastseen,
			Online:     node.Online,
			Statistics: node.Statistics,
			Nodeinfo: &data.NodeInfo{
				NodeID:   nodeinfo.NodeID,
				Network:  nodeinfo.Network,
				System:   nodeinfo.System,
				Owner:    nil,
				Hostname: nodeinfo.Hostname,
				Location: nodeinfo.Location,
				Software: nodeinfo.Software,
				Hardware: nodeinfo.Hardware,
				VPN:      nodeinfo.VPN,
				Wireless: nodeinfo.Wireless,
			},
			Neighbours: node.Neighbours,
		}
	}
	return node
}
