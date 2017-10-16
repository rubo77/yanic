package all

import (
	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
)

// Config Filter
type filterConfig map[string]interface{}

func (f filterConfig) Blacklist() *map[string]interface{} {
	if v, ok := f["blacklist"]; ok {
		list := make(map[string]interface{})
		for _, nodeid := range v.([]interface{}) {
			list[nodeid.(string)] = true
		}
		return &list
	}
	return nil
}

func (f filterConfig) NoOwner() bool {
	if v, ok := f["no_owner"]; ok {
		return v.(bool)
	}
	return true
}
func (f filterConfig) HasLocation() *bool {
	if v, ok := f["has_location"].(bool); ok {
		return &v
	}
	return nil
}

type area struct {
	xA float64
	xB float64
	yA float64
	yB float64
}

func (f filterConfig) InArea() *area {
	if areaConfigInt, ok := f["in_area"]; ok {
		areaConfig := areaConfigInt.(map[string]interface{})
		a := area{}
		if v, ok := areaConfig["latitude_min"]; ok {
			a.xA = v.(float64)
		}
		if v, ok := areaConfig["latitude_max"]; ok {
			a.xB = v.(float64)
		}
		if v, ok := areaConfig["longitude_min"]; ok {
			a.yA = v.(float64)
		}
		if v, ok := areaConfig["longitude_max"]; ok {
			a.yB = v.(float64)
		}
		return &a
	}
	return nil
}

// Create Filter
func (f filterConfig) filtering(nodesOrigin *runtime.Nodes) *runtime.Nodes {
	nodes := runtime.NewNodes(&runtime.Config{})
	for nodeID, nodeOrigin := range nodesOrigin.List {
		//maybe cloning of this object is better?
		node := nodeOrigin

		if f.NoOwner() {
			node = filterNoOwner(node)
		}
		if ok := f.HasLocation(); ok != nil {
			node = filterHasLocation(node, *ok)
		}
		if area := f.InArea(); area != nil {
			node = filterLocationInArea(node, *area)
		}
		if list := f.Blacklist(); list != nil {
			node = filterBlacklist(node, *list)
		}
		if node != nil {
			nodes.Update(nodeID, &data.ResponseData{
				NodeInfo:   node.Nodeinfo,
				Statistics: node.Statistics,
				Neighbours: node.Neighbours,
			})
		}
	}
	return nodes
}

func filterBlacklist(node *runtime.Node, list map[string]interface{}) *runtime.Node {
	if node != nil {
		if nodeinfo := node.Nodeinfo; nodeinfo != nil {
			if _, ok := list[nodeinfo.NodeID]; !ok {
				return node
			}
		}
	}
	return nil
}

func filterNoOwner(node *runtime.Node) *runtime.Node {
	if node == nil {
		return nil
	}
	return &runtime.Node{
		Address:    node.Address,
		Firstseen:  node.Firstseen,
		Lastseen:   node.Lastseen,
		Online:     node.Online,
		Statistics: node.Statistics,
		Nodeinfo: &data.NodeInfo{
			NodeID:   node.Nodeinfo.NodeID,
			Network:  node.Nodeinfo.Network,
			System:   node.Nodeinfo.System,
			Owner:    nil,
			Hostname: node.Nodeinfo.Hostname,
			Location: node.Nodeinfo.Location,
			Software: node.Nodeinfo.Software,
			Hardware: node.Nodeinfo.Hardware,
			VPN:      node.Nodeinfo.VPN,
			Wireless: node.Nodeinfo.Wireless,
		},
		Neighbours: node.Neighbours,
	}
}

func filterHasLocation(node *runtime.Node, withLocation bool) *runtime.Node {
	if node != nil {
		if nodeinfo := node.Nodeinfo; nodeinfo != nil {
			if withLocation {
				if location := nodeinfo.Location; location != nil {
					return node
				}
			} else {
				if location := nodeinfo.Location; location == nil {
					return node
				}
			}
		}
	}
	return nil
}

func filterLocationInArea(node *runtime.Node, a area) *runtime.Node {
	if node != nil {
		if nodeinfo := node.Nodeinfo; nodeinfo != nil {
			if location := nodeinfo.Location; location != nil {
				if location.Latitude >= a.xA && location.Latitude <= a.xB {
					if location.Longtitude >= a.yA && location.Longtitude <= a.yB {
						return node
					}
				}
			} else {
				return node
			}
		}
	}
	return nil
}
