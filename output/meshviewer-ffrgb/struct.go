package meshviewerFFRGB

import (
	"time"

	"github.com/FreifunkBremen/yanic/jsontime"
	"github.com/FreifunkBremen/yanic/runtime"
)

type Meshviewer struct {
	Timestamp jsontime.Time `json:"timestamp"`
	Nodes     []*Node       `json:"nodes"`
	Links     []*Link       `json:"links"`
}

type Node struct {
	Firstseen     jsontime.Time `json:"firstseen"`
	Lastseen      jsontime.Time `json:"lastseen"`
	IsOnline      bool          `json:"is_online"`
	IsGateway     bool          `json:"is_gateway"`
	Clients       uint32        `json:"clients"`
	ClientsWifi24 uint32        `json:"clients_wifi24"`
	ClientsWifi5  uint32        `json:"clients_wifi5"`
	ClientsOthers uint32        `json:"clients_other"`
	RootFSUsage   float64       `json:"rootfs_usage,omitempty"`
	LoadAverage   float64       `json:"loadavg,omitempty"`
	MemoryUsage   *float64      `json:"memory_usage,omitempty"`
	Uptime        jsontime.Time `json:"uptime,omitempty"`
	GatewayIPv4   string        `json:"gateway,omitempty"`
	GatewayIPv6   string        `json:"gateway6,omitempty"`
	NodeID        string        `json:"node_id"` // duplicated, ja bzw. nein ?
	Network       Network       `json:"network"`
	SiteCode      string        `json:"site_code,omitempty"`
	Hostname      string        `json:"hostname"`
	Location      *Location     `json:"location,omitempty"`
	Firmware      Firmware      `json:"firmware,omitempty"`
	Autoupdater   Autoupdater   `json:"autoupdater,omitempty"`
	Nproc         int           `json:"nproc"`
	Model         string        `json:"model,omitempty"`
	VPN           bool          `json:"vpn"`
}

// Firmware out of software
type Firmware struct {
	Base    string `json:"base,omitempty"`
	Release string `json:"release,omitempty"`
}

// Autoupdater
type Autoupdater struct {
	Enabled bool   `json:"enabled,omitempty"`
	Branch  string `json:"branch,omitempty"`
}

// Network struct
type Network struct {
	MAC       string   `json:"mac"`
	Addresses []string `json:"addresses"`
}

// Location struct
type Location struct {
	Longtitude float64 `json:"longitude,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
}

// Link
type Link struct {
	Type      string  `json:"type"`
	Source    string  `json:"source"`
	Target    string  `json:"target"`
	SourceTQ  float32 `json:"source_tq"`
	TargetTQ  float32 `json:"target_tq"`
	SourceMAC string  `json:"-"`
	TargetMAC string  `json:"-"`
}

func NewNode(n *runtime.Node) *Node {
	node := &Node{
		Firstseen: n.Firstseen,
		Lastseen:  n.Lastseen,
		IsOnline:  n.Online,
		IsGateway: n.IsGateway(),
	}

	if nodeinfo := n.Nodeinfo; nodeinfo != nil {
		node.NodeID = nodeinfo.NodeID
		node.Network = Network{
			MAC:       nodeinfo.Network.Mac,
			Addresses: nodeinfo.Network.Addresses,
		}
		node.SiteCode = nodeinfo.System.SiteCode
		node.Hostname = nodeinfo.Hostname
		if location := nodeinfo.Location; location != nil {
			node.Location = &Location{
				Longtitude: location.Longtitude,
				Latitude:   location.Latitude,
			}
		}
		node.Firmware = nodeinfo.Software.Firmware
		node.Autoupdater = nodeinfo.Software.Autoupdater
		node.Nproc = nodeinfo.Hardware.Nproc
		node.Model = nodeinfo.Hardware.Model
		node.VPN = nodeinfo.VPN
	}
	if statistic := n.Statistics; statistic != nil {
		node.Clients = statistic.Clients.Total
		if node.Clients == 0 {
			node.Clients = statistic.Clients.Wifi24 + statistic.Clients.Wifi5
		}
		node.ClientsWifi24 = statistic.Clients.Wifi24
		node.ClientsWifi5 = statistic.Clients.Wifi5

		node.ClientsOthers = node.Clients - node.ClientsWifi24 - node.ClientsWifi5

		node.RootFSUsage = statistic.RootFsUsage
		node.LoadAverage = statistic.LoadAverage

		/* The Meshviewer could not handle absolute memory output
		 * calc the used memory as a float which 100% equal 1.0
		 * calc is coppied from node statuspage (look discussion:
		 * https://github.com/FreifunkBremen/yanic/issues/35)
		 */
		if statistic.Memory.Total > 0 {
			usage := 1 - (float64(statistic.Memory.Free)+float64(statistic.Memory.Buffers)+float64(statistic.Memory.Cached))/float64(statistic.Memory.Total)
			node.MemoryUsage = &usage
		}

		node.Uptime = jsontime.Now().Add(time.Duration(statistic.Uptime) * -time.Second)
		node.GatewayIPv4 = statistic.GatewayIPv4
		node.GatewayIPv6 = statistic.GatewayIPv6
	}

	return node
}
