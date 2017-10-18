package all

import "github.com/FreifunkBremen/yanic/runtime"

func (f filterConfig) HasLocation() *bool {
	if v, ok := f["has_location"].(bool); ok {
		return &v
	}
	return nil
}

func filterHasLocation(node *runtime.Node, withLocation bool) *runtime.Node {
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
	} else if !withLocation {
		return node
	}
	return nil
}
