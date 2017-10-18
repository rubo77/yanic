package all

import "github.com/FreifunkBremen/yanic/runtime"

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

func filterBlacklist(node *runtime.Node, list map[string]interface{}) *runtime.Node {
	if nodeinfo := node.Nodeinfo; nodeinfo != nil {
		if _, ok := list[nodeinfo.NodeID]; ok {
			return nil
		}
	}
	return node
}
