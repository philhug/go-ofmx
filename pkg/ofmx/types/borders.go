package types

// Geographical border types
type GbrUid struct {
	Uid
	TxtName  string `xml:"txtName"`

}

type Gbr struct {
	GbrUid GbrUid  `xml:"GbrUid"`
	CodeType  string `xml:"codeType"`
	TxtRmk  string `xml:"txtRmk"`
	Gbv  []Gbv `xml:"Gbv"`
}

type Gbv struct {
	GbrUid GbrUid  `xml:"GbrUid"`
	CodeType  string `xml:"codeType"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
	CodeDatum  string `xml:"codeDatum"`
	ValGeoAccuracy  string `xml:"valGeoAccuracy"`
	UomGeoAccuracy  string `xml:"uomGeoAccuracy"`
	GeoLatArc  string `xml:"geoLatArc"`
	GeoLongArc  string `xml:"geoLongArc"`
	ValRadiusArc  string `xml:"valRadiusArc"`
	UomRadiusArc  string `xml:"uomRadiusArc"`
	ValHex  string `xml:"valHex"`
	TxtRmk  string `xml:"txtRmk"`
	}


func (uid *GbrUid) String() string {
	return uidString(*uid)
}

func (uid *GbrUid) Hash() string {
	return uidHash(*uid)
}

func (f *Gbr) Uid() FeatureUid {
	return &f.GbrUid
}
