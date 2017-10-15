package meshviewerFFRGB

import (
	"github.com/FreifunkBremen/yanic/jsontime"
	"github.com/FreifunkBremen/yanic/runtime"
)

func transform(nodes *runtime.Nodes) *Meshviewer {

	meshviewer := &Meshviewer{
		Timestamp: jsontime.Now(),
	}

	links := make(map[string]*Link)

	for _, nodeOrigin := range nodes.List {
		node := NewNode(nodeOrigin)
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

		for _, link := range nodes.NodeLinks(nodeOrigin) {
			if links[link.SourceMAC] != nil {
				continue
			}
			linkType := typeList[link.SourceMAC]
			if linkType == "" {
				linkType = "other"
			}
			link := &Link{
				Type:      linkType,
				Source:    link.SourceID,
				SourceMAC: link.SourceMAC,
				Target:    link.TargetID,
				TargetMAC: link.TargetMAC,
				SourceTQ:  float32(link.TQ) / 255.0,
			}
			links[link.TargetMAC] = link
			meshviewer.Links = append(meshviewer.Links, link)
		}
	}
	for _, nodeOrigin := range nodes.List {
		for _, linkOrigin := range nodes.NodeLinks(nodeOrigin) {
			link := links[linkOrigin.SourceMAC]
			if link != nil {
				link.TargetTQ = float32(linkOrigin.TQ) / 255.0
			}
		}
	}

	return meshviewer
}
