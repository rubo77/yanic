package meshviewerFFRGB

import (
	"fmt"
	"strings"

	"github.com/FreifunkBremen/yanic/jsontime"
	"github.com/FreifunkBremen/yanic/runtime"
)

func transform(nodes *runtime.Nodes) *Meshviewer {

	meshviewer := &Meshviewer{
		Timestamp: jsontime.Now(),
	}

	links := make(map[string]*Link)

	nodes.RLock()
	defer nodes.RUnlock()

	for _, nodeOrigin := range nodes.List {
		node := NewNode(nodes, nodeOrigin)
		meshviewer.Nodes = append(meshviewer.Nodes, node)

		typeList := make(map[string]string)

		if nodeinfo := nodeOrigin.Nodeinfo; nodeinfo != nil {
			if meshes := nodeinfo.Network.Mesh; meshes != nil {
				for _, mesh := range meshes {
					for _, mac := range mesh.Interfaces.Wireless {
						typeList[mac] = "wifi"
					}
					for _, mac := range mesh.Interfaces.Tunnel {
						typeList[mac] = "vpn"
					}
				}
			}
		}

		for _, linkOrigin := range nodes.NodeLinks(nodeOrigin) {
			var key string
			if strings.Compare(linkOrigin.SourceMAC, linkOrigin.TargetMAC) > 0 {
				key = fmt.Sprintf("%s-%s", linkOrigin.SourceMAC, linkOrigin.TargetMAC)
			} else {
				key = fmt.Sprintf("%s-%s", linkOrigin.TargetMAC, linkOrigin.SourceMAC)
			}
			if link := links[key]; link != nil {
				link.TargetTQ = float32(linkOrigin.TQ) / 255.0
				continue
			}
			linkType := typeList[linkOrigin.SourceMAC]
			if linkType == "" {
				linkType = "other"
			}
			tq := float32(linkOrigin.TQ) / 255.0
			link := &Link{
				Type:      linkType,
				Source:    linkOrigin.SourceID,
				SourceMAC: linkOrigin.SourceMAC,
				Target:    linkOrigin.TargetID,
				TargetMAC: linkOrigin.TargetMAC,
				SourceTQ:  tq,
				TargetTQ:  tq,
			}
			links[key] = link
			meshviewer.Links = append(meshviewer.Links, link)
		}
	}

	return meshviewer
}
