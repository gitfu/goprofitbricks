// loadbalancer_test.go
package goprofitbricks

import "testing"
import "fmt"

var lbal_dcid = mkdcid()

func mklbalid(lbal_dcid string) string {
	var jason = []byte(`{
					"name":"Original Loadbalancer",
					}`)
	lbal := CreateLoadbalancer(lbal_dcid, jason)

	lbalid := lbal.Id
	return lbalid
}

func ExampleListLoadbalancers() {
	s := ListLoadbalancers(lbal_dcid)
	fmt.Println(s.Resp.StatusCode)
	// Output: 200

}

func TestListLoadbalancers(t *testing.T) {
	//t.Parallel()
	shouldbe := "collection"
	want := 200
	resp := ListLoadbalancers(lbal_dcid)

	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestCreateLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	var jason = []byte(`{
					"name":"Goat",
					}`)
	resp := CreateLoadbalancer(lbal_dcid, jason)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}

}

func TestGetLoadbalancer(t *testing.T) {
	//t.Parallel()
	shouldbe := "loadbalancer"
	want := 200
	lbalid := mklbalid(lbal_dcid)
	resp := GetLoadbalancer(lbal_dcid, lbalid)

	if resp.Type != shouldbe {
		t.Errorf(badType(shouldbe, resp.Type))
	}
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestPatchLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	jasonPatch := []byte(`{
					"name":"Renamed Loadbalancer",
					}`)
	lbalid := mklbalid(lbal_dcid)
	resp := PatchLoadbalancer(lbal_dcid, lbalid, jasonPatch)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}

func TestUpdateLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	jasonUpdate := []byte(`{
					"name":"Renamed Loadbalancer",
				
					}`)

	lbalid := mklbalid(lbal_dcid)
	resp := UpdateLoadbalancer(lbal_dcid, lbalid, jasonUpdate)
	if resp.Resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.Resp.StatusCode))
	}
}
func TestDeleteLoadbalancer(t *testing.T) {
	//t.Parallel()
	want := 202
	lbalid := mklbalid(lbal_dcid)
	resp := DeleteLoadbalancer(lbal_dcid, lbalid)
	if resp.StatusCode != want {
		t.Errorf(badStatus(want, resp.StatusCode))
	}
}
