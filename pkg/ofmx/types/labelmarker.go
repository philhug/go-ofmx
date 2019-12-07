package types

type LbmUid struct {
	RegionalUid
	CodeType string `xml:"codeType"`
	TxtName  string `xml:"txtName"`
}

func (uid *LbmUid) String() string {
	return uidString(*uid)
}

func (uid *LbmUid) Hash() string {
	return uidHash(*uid)
}

type Lbm struct {
	LbmUid        LbmUid      `xml:"LbmUid"`
	TxtValueLabel string      `xml:"txtValueLabel"`
	GeoLat        string      `xml:"geoLat"`
	GeoLong       string      `xml:"geoLong"`
	CodeDatum     string      `xml:"codeDatum"`
	ZoomLevel     []ZoomLevel `xml:"ZoomLevel"`
}

func (f *Lbm) Uid() FeatureUid {
	return &f.LbmUid
}

type ZoomLevel struct {
	ValZoomLevel  int    `xml:"valZoomLevel"`
	TxtValueLabel string `xml:"txtValueLabel"`
	GeoLat        string `xml:"geoLat"`
	GeoLong       string `xml:"geoLong"`
}
