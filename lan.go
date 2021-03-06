package goprofitbricks

// ListLan returns a Collection for lans in the Datacenter
func ListLans(dcid string) Collection {
	path := lan_col_path(dcid)
	return isList(path)
}

// CreateLan creates a lan in the datacenter
// from a jason []byte and returns a Instance struct
func CreateLan(dcid string, jason []byte) Instance {
	path := lan_col_path(dcid)
	return is_post(path, jason)
}

// GetLan pulls data for the lan where id = lanid returns an Instance struct
func GetLan(dcid, lanid string) Instance {
	path := lan_path(dcid, lanid)
	return isGet(path)
}

// UpdateLan does a complete update to a lan using json from []byte jason
// returns a Instance struct
func UpdateLan(dcid string, lanid string, jason []byte) Instance {
	path := lan_path(dcid, lanid)
	return isPut(path, jason)
}

// PatchLan does a partial update to a lan using json from []byte jason
// returns a Instance struct
func PatchLan(dcid string, lanid string, jason []byte) Instance {
	path := lan_path(dcid, lanid)
	return isPatch(path, jason)
}

// DeleteLan deletes a lan where id == lanid
func DeleteLan(dcid, lanid string) Resp {
	path := lan_path(dcid, lanid)
	return isDelete(path)
}

// ListLanMembers returns a Nic struct collection for the Lan
func ListLanMembers(dcid, lanid string) Collection {
	path := lan_nic_col(dcid, lanid)
	return isList(path)
}
