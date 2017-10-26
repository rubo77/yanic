package data

// NodeInfo struct
type NodeInfo struct {
	NodeID   string    `json:"node_id"`
	Network  Network   `json:"network"`
	Owner    *Owner    `json:"owner"`
	System   System    `json:"system"`
	Hostname string    `json:"hostname"`
	Location *Location `json:"location,omitempty"`
	Software Software  `json:"software"`
	Hardware Hardware  `json:"hardware"`
	VPN      bool      `json:"vpn"`
	Wireless *Wireless `json:"wireless,omitempty"`
}

// BatInterface struct
type BatInterface struct {
	Interfaces struct {
		Wireless []string `json:"wireless,omitempty"`
		Other    []string `json:"other,omitempty"`
		Tunnel   []string `json:"tunnel,omitempty"`
	} `json:"interfaces"`
}

// Addresses returns a flat list of all MAC addresses
func (iface *BatInterface) Addresses() []string {
	return append(append(iface.Interfaces.Other, iface.Interfaces.Tunnel...), iface.Interfaces.Wireless...)
}

// Network struct
type Network struct {
	Mac            string                   `json:"mac"`
	Addresses      []string                 `json:"addresses"`
	Mesh           map[string]*BatInterface `json:"mesh"`
	MeshInterfaces []string                 `json:"mesh_interfaces"`
}

// Owner struct
type Owner struct {
	Contact string `json:"contact"`
}

// System struct
type System struct {
	SiteCode string `json:"site_code,omitempty"`
}

// Location struct
type Location struct {
	Longtitude float64 `json:"longitude,omitempty"`
	Latitude   float64 `json:"latitude,omitempty"`
	Altitude   float64 `json:"altitude,omitempty"`
}

// Software struct
type Software struct {
	Autoupdater struct {
		Enabled bool   `json:"enabled,omitempty"`
		Branch  string `json:"branch,omitempty"`
	} `json:"autoupdater,omitempty"`
	BatmanAdv struct {
		Version string `json:"version,omitempty"`
		Compat  int    `json:"compat,omitempty"`
	} `json:"batman-adv,omitempty"`
	Fastd struct {
		Enabled bool   `json:"enabled,omitempty"`
		Version string `json:"version,omitempty"`
	} `json:"fastd,omitempty"`
	Firmware struct {
		Base    string `json:"base,omitempty"`
		Release string `json:"release,omitempty"`
	} `json:"firmware,omitempty"`
	StatusPage struct {
		API int `json:"api"`
	} `json:"status-page,omitempty"`
}

// Hardware struct
type Hardware struct {
	Nproc int    `json:"nproc"`
	Model string `json:"model,omitempty"`
}
