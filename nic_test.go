// nic_test.go

package goprofitbricks

import "testing"

var nic_dcid = mkdcid()
var nic_srvid = mksrvid(nic_dcid)

func mknicid(nic_dcid, nic_srvid string) string {
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	nic := CreateNic(nic_dcid, nic_srvid, jason)
	nicid := nic.Id
	return nicid
}

func TestListNics(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListNics(nic_dcid, nic_srvid)
	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestCreateNic(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Original Nic",
					"lan":1
					}`)

	resp := CreateNic(nic_dcid, nic_srvid, jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestGetNic(t *testing.T) {
	//t.Parallel()
	shouldbe := "nic"
	want := 200
	nicid := mknicid(nic_dcid, nic_srvid)
	resp := GetNic(nic_dcid, nic_srvid, nicid)
	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestPatchNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jasonPatch := []byte(`{
					"name":"Patched Nic",
					"lan":1
					}`)

	nicid := mknicid(nic_dcid, nic_srvid)
	resp := PatchNic(nic_dcid, nic_srvid, nicid, jasonPatch)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestUpdateNic(t *testing.T) {
	//t.Parallel()
	want := 202
	jasonUpdate := []byte(`{
					"name":"Update Nic",
					"lan":1
					}`)

	nicid := mknicid(nic_dcid, nic_srvid)
	resp := UpdateNic(nic_dcid, nic_srvid, nicid, jasonUpdate)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestDeleteNic(t *testing.T) {
	//t.Parallel()
	want := 202
	nicid := mknicid(nic_dcid, nic_srvid)
	resp := DeleteNic(nic_dcid, nic_srvid, nicid)
	if resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.StatusCode))
	}
}
