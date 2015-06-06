package gonap

import "encoding/json"

// Ipblock_Properties is just that
type Ipblock_Properties struct {
	Location string `json:"location"`
	Size     int    `json:"size"`
}

// Ipblock is the struct for Ipblock data
type Ipblock struct {
	Id_Type_Href
	MetaData   MetaData           `json:"metadata,omitempty"`
	Properties Ipblock_Properties `json:"properties"`
	Resp       PBResp             `json:"-"`
}

func toIpblock(pbresp PBResp) Ipblock {
	var ipb Ipblock
	json.Unmarshal(pbresp.Body, &ipb)
	ipb.Resp = pbresp
	return ipb
}

// Ipblocks is the struct for an Ipblock collection
type Ipblocks struct {
	Id_Type_Href
	Items []Ipblock `json:"items,omitempty"`
	Resp  PBResp    `json:"-"`
}

func toIpblocks(pbresp PBResp) Ipblocks {
	var ipbs Ipblocks
	json.Unmarshal(pbresp.Body, &ipbs)
	ipbs.Resp = pbresp
	return ipbs
}

// ListIpBlocks
func ListIpBlocks() Ipblocks {
	path := ipblock_col_path()
	return toIpblocks(is_get(path))
}
func ReserveIpBlock(jason []byte) Ipblock {
	path := ipblock_col_path()
	return toIpblock(is_post(path, jason))

}
func GetIpBlock(ipblockid string) Ipblock {
	path := ipblock_path(ipblockid)
	return toIpblock(is_get(path))
}

func ReleaseIpBlock(ipblockid string) Ipblock {
	path := ipblock_path(ipblockid)
	return toIpblock(is_delete(path))
}
