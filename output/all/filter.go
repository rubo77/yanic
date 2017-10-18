package all

import (
	"github.com/FreifunkBremen/yanic/data"
	"github.com/FreifunkBremen/yanic/runtime"
)

// Config Filter
type filterConfig map[string]interface{}

// Create Filter
func (f filterConfig) filtering(nodesOrigin *runtime.Nodes) *runtime.Nodes {
	nodes := runtime.NewNodes(&runtime.Config{})
	for nodeID, nodeOrigin := range nodesOrigin.List {
		//maybe cloning of this object is better?
		node := nodeOrigin

		if f.NoOwner() {
			node = filterNoOwner(node)
		}
		if ok := f.HasLocation(); node != nil && ok != nil {
			node = filterHasLocation(node, *ok)
		}
		if area := f.InArea(); node != nil && area != nil {
			node = filterLocationInArea(node, area)
		}
		if list := f.Blacklist(); node != nil && list != nil {
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
