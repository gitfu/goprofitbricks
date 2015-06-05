package gonap

import "encoding/json"

type ImageProperties struct {
	Name                string `json:"name,omitempty"`
	Description         string `json:"description,omitempty"`
	Location            string `json:"location"`
	Size                int    `json:"size"`
	Public              bool   `json:"public,omitempty"`
	ImageType           string `json:"imageType,omitempty"`
	CpuHotPlug          bool   `json:"cpuHotPlug,omitempty"`
	CpuHotUnplug        bool   `json:"cpuHotUnplug,omitempty"`
	RamHotPlug          bool   `json:"ramHotPlug,omitempty"`
	RamHotUnplug        bool   `json:"ramHotUnplug,omitempty"`
	NicHotPlug          bool   `json:"nicHotPlug,omitempty"`
	NicHotUnplug        bool   `json:"nicHotUnplug,omitempty"`
	DiscVirtioHotPlug   bool   `json:"discVirtioHotPlug,omitempty"`
	DiscVirtioHotUnplug bool   `json:"discVirtioHotUnplug,omitempty"`
	DiscScsiHotPlug     bool   `json:"discScsiHotPlug,omitempty"`
	DiscScsiHotUnplug   bool   `json:"discScsiHotUnplug,omitempty"`
	LicenceType         string `json:"licenceType,omitempty"`
}

// Image is thr struct for image data
type Image struct {
	Id         string          `json:"id,omitempty"`
	Type       string          `json:"type,omitempty"`
	Href       string          `json:"href,omitempty"`
	MetaData   MetaData        `json:"metadata,omitempty"`
	Properties ImageProperties `json:"properties,omitempty"`
	Resp       PBResp          `json:"-"`
}

func toImage(pbresp PBResp) Image {
	var IMG Image
	json.Unmarshal(pbresp.Body, &IMG)
	return IMG
}

// Images is a struct for Image collections
type Images struct {
	Id    string  `json:"id,omitempty"`
	Type  string  `json:"type,omitempty"`
	Href  string  `json:"href,omitempty"`
	Items []Image `json:"items,omitempty"`
	Resp  PBResp  `json:"-"`
}

func toImages(pbresp PBResp) Images {
	var IMGS Images
	json.Unmarshal(pbresp.Body, &IMGS)
	IMGS.Resp = pbresp
	return IMGS
}

// ListImages returns an Images struct
func ListImages() Images {
	path := image_col_path()
	return toImages(is_get(path))
}

// CreateImage creates an Image and returns an Image struct
func CreateImage(jason []byte) Image {
	path := image_col_path()
	return toImage(is_post(path, jason))
}

// GetImage returns an Image struct where id ==imageid
func GetImage(imageid string) Image {
	path := image_path(imageid)
	return toImage(is_get(path))
}

// UpdateImage updates all image properties from values in jason
//returns an Image struct where id ==imageid
func UpdateImage(imageid string, jason []byte) Image {
	path := image_path(imageid)
	return toImage(is_put(path, jason))
}

// PatchImage replaces any image properties from values in jason
//returns an Image struct where id ==imageid
func PatchImage(imageid string, jason []byte) Image {
	path := image_path(imageid)
	return toImage(is_patch(path, jason))
}

// Deletes an image where id==imageid
func DeleteImage(imageid string) Image {
	path := image_path(imageid)
	return toImage(is_delete(path))
}