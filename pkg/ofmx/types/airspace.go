package types

// TODO, still inconsistencies with documentation

type AseUid struct {
	RegionalUid
	CodeType string `xml:"codeType"`
	CodeId string `xml:"codeId"`
}

func (uid *AseUid) String() string {
	return uidString(*uid)
}

func (uid *AseUid) Hash() string {
	return uidHash(*uid)
}

type Ase struct {
	AseUid           AseUid `xml:"AseUid"`
	OrgUid           OrgUid `xml:"OrgUid"`
	TxtName          string `xml:"txtName"`
	CodeClass        string `xml:"codeClass"`
	CodeDistVerUpper string `xml:"codeDistVerUpper"`
	ValDistVerUpper  string `xml:"valDistVerUpper"`
	UomDistVerUpper  string `xml:"uomDistVerUpper"`
	CodeDistVerLower string `xml:"codeDistVerLower"`
	ValDistVerLower  string `xml:"valDistVerLower"`
	UomDistVerLower  string `xml:"uomDistVerLower"`
	CodeDistVerMax   string `xml:"codeDistVerMax"`
	ValDistVerMax    string `xml:"valDistVerMax"`
	UomDistVerMax    string `xml:"uomDistVerMax"`
	CodeDistVerMnm   string `xml:"codeDistVerMnm"`
	ValDistVerMnm    string `xml:"valDistVerMnm"`
	UomDistVerMnm    string `xml:"uomDistVerMnm"`
	//Att
	CodeSelAvbl string `xml:"codeSelAvbl"`
	TxtRmk      string `xml:"txtRmk"`
}

func (f *Ase) Uid() FeatureUid {
	return &f.AseUid
}

type AbdUid struct {
	AseUid AseUid `xml:"AseUid"`
}

func (uid *AbdUid) String() string {
	return uidString(*uid)
}

func (uid *AbdUid) Hash() string {
	return uidHash(*uid)
}

type Abd struct {
	AbdUid AbdUid `xml:"AbdUid"`
	Avx    []Avx    `xml:"Avx"`
}

func (f *Abd) Uid() FeatureUid {
	return &f.AbdUid
}


type Avx struct { // TODO
	GbrUid GbrUid `xml:"GbrUid"`
	CodeType   string `xml:"codeType"`
	GeoLat     string `xml:"geoLat"`
	GeoLong    string `xml:"geoLong"`
	CodeDatum  string `xml:"codeDatum"`
	ValGeoAccuracy  string `xml:"valGeoAccuracy"`
	UomGeoAccuracy  string `xml:"uomGeoAccuracy"`
	GeoLatArc  string `xml:"geoLatArc"`
	GeoLongArc string `xml:"geoLongArc"`
	ValRadiusArc  string `xml:"valRadiusArc"`
	UomRadiusArc  string `xml:"uomRadiusArc"`
	ValHex  string `xml:"valHex"`
	TxtRmk  string `xml:"txtRmk"`
}

type AdgUid struct {
	AseUid AseUid `xml:"AseUid"`
}

func (uid *AdgUid) String() string {
	return uidString(*uid)
}

func (uid *AdgUid) Hash() string {
	return uidHash(*uid)
}

type Adg struct {
	AdgUid           AdgUid `xml:"AdgUid"`
	AseUidSameExtent AseUid `xml:"AseUid"`
}

func (f *Adg) Uid() FeatureUid {
	return &f.AdgUid
}

type SaeUid struct {
	SerUid SerUid `xml:"SerUid"`
	AdgUid AdgUid `xml:"AdgUid"`
}

func (uid *SaeUid) String() string {
	return uidString(*uid)
}

func (uid *SaeUid) Hash() string {
	return uidHash(*uid)
}

type Sae struct {
	SaeUid SaeUid `xml:"SaeUid"`
}

func (f *Sae) Uid() FeatureUid {
	return &f.SaeUid
}
