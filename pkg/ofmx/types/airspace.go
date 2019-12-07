package types

// TODO, still inconsistencies with documentation

type AseUid struct {
	RegionalUid
	CodeType string `xml:"codeType"`
	CodeId   string `xml:"codeId"`

	// TODO, drop again
	NewEntity string `xml:"newEntity,attr" validate:"isdefault"`
}

func (uid *AseUid) String() string {
	return uidString(*uid)
}

func (uid *AseUid) Hash() string {
	return uidHash(*uid)
}

type Ase struct {
	AseUid           AseUid    `xml:"AseUid"`
	OrgUid           OrgUid    `xml:"OrgUid"`
	TxtLocalType     string    `xml:"txtLocalType"`
	TxtName          string    `xml:"txtName"`
	CodeClass        string    `xml:"codeClass"`
	CodeLocInd       string    `xml:"codeLocInd"`
	CodeDistVerUpper string    `xml:"codeDistVerUpper"`
	ValDistVerUpper  string    `xml:"valDistVerUpper"`
	UomDistVerUpper  string    `xml:"uomDistVerUpper"`
	CodeDistVerLower string    `xml:"codeDistVerLower"`
	ValDistVerLower  string    `xml:"valDistVerLower"`
	UomDistVerLower  string    `xml:"uomDistVerLower"`
	CodeDistVerMax   string    `xml:"codeDistVerMax"`
	ValDistVerMax    string    `xml:"valDistVerMax"`
	UomDistVerMax    string    `xml:"uomDistVerMax"`
	CodeDistVerMnm   string    `xml:"codeDistVerMnm"`
	ValDistVerMnm    string    `xml:"valDistVerMnm"`
	UomDistVerMnm    string    `xml:"uomDistVerMnm"`
	Att              Timetable `xml:"Att"`
	CodeSelAvbl      string    `xml:"codeSelAvbl"`
	TxtRmk           string    `xml:"txtRmk"`

	// TODO, TODO, remove
	XtSelAvail         string `xml:"xt_selAvail" validate:"isdefault"`
	XtClassLayersAvail string `xml:"xt_classLayersAvail,attr" validate:"isdefault"`
	XXtxtRmk           string `xml:"xt_txtRmk" validate:"isdefault"`
}

func (f *Ase) Uid() FeatureUid {
	return &f.AseUid
}

type AbdUid struct {
	// TODO, temp allow mid
	Uid

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
	Avx    []Avx  `xml:"Avx"`
}

func (f *Abd) Uid() FeatureUid {
	return &f.AbdUid
}

// Avx - Airspace Border Vertex
type Avx struct { // TODO
	GbrUid         GbrUid `xml:"GbrUid"`
	CodeType       string `xml:"codeType"`
	GeoLat         string `xml:"geoLat"`
	GeoLong        string `xml:"geoLong"`
	CodeDatum      string `xml:"codeDatum"`
	ValGeoAccuracy string `xml:"valGeoAccuracy"`
	UomGeoAccuracy string `xml:"uomGeoAccuracy"`
	GeoLatArc      string `xml:"geoLatArc"`
	GeoLongArc     string `xml:"geoLongArc"`
	ValRadiusArc   string `xml:"valRadiusArc"`
	UomRadiusArc   string `xml:"uomRadiusArc"`
	ValHex         string `xml:"valHex"`
	TxtRmk         string `xml:"txtRmk"`

	// TODO, invalid
	XXValCrc string `xml:"valCrc" validate:"isdefault"`
}

type AdgUid struct {
	// TODO, temp allow mid
	Uid

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
	AseUidSameExtent AseUid `xml:"AseUidSameExtent"`
}

func (f *Adg) Uid() FeatureUid {
	return &f.AdgUid
}

type SaeUid struct {
	// TODO, temp allow mid
	Uid

	SerUid SerUid `xml:"SerUid"`
	AseUid AseUid `xml:"AseUid"`
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
