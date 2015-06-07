package gonap

import "encoding/json"

type Nic Instance

func toNic(pbresp PBResp) Nic {
	var nic Nic
	json.Unmarshal(pbresp.Body, &nic)
	nic.Resp = pbresp
	return nic
}

type Nics Collection

func toNics(pbresp PBResp) Nics {
	var nics Nics
	json.Unmarshal(pbresp.Body, &nics)
	nics.Resp = pbresp
	return nics
}

// ListNics returns a Nics struct collection
func ListNics(dcid, srvid string) Nics {
	path := nic_col_path(dcid, srvid)
	return toNics(is_get(path))
}

// CreateNic creates a nic on a server
// from a jason []byte and returns a Nic struct
func CreateNic(dcid string, srvid string, jason []byte) Nic {
	path := nic_col_path(dcid, srvid)
	return toNic(is_post(path, jason))
}

// GetNic pulls data for the nic where id = srvid returns a Nic struct
func GetNic(dcid, srvid, nicid string) Nic {
	path := nic_path(dcid, srvid, nicid)
	return toNic(is_get(path))
}

// UpdateNic is a full update of nic properties passed in as jason []byte
// Returns Nic struct
func UpdateNic(dcid string, srvid string, nicid string, jason []byte) Nic {
	path := nic_path(dcid, srvid, nicid)
	return toNic(is_put(path, jason))
}

// PatchNic partial update of nic properties passed in as jason []byte
// Returns Nic struct
func PatchNic(dcid string, srvid string, nicid string, jason []byte) Nic {
	path := nic_path(dcid, srvid, nicid)
	return toNic(is_patch(path, jason))
}

// DeleteNic deletes the nic where id=nicid and returns a PBResp struct
func DeleteNic(dcid, srvid, nicid string) Nic {
	path := nic_path(dcid, srvid, nicid)
	return toNic(is_delete(path))
}
