package types

// Obstacle Group

type OgrUid struct {
	RegionalUid

	TxtName string `xml:"txtName"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *OgrUid) String() string {
	return uidString(*uid)
}

func (uid *OgrUid) Hash() string {
	return uidHash(*uid)
}

type Ogr struct {
	OgrUid OgrUid `xml:"OgrUid"`
	// TODO
}

func (f *Ogr) Uid() FeatureUid {
	return &f.OgrUid
}

// Obstacle

type ObsUid struct {
	// TODO, temp allow mid
	Uid

	OgrUid OgrUid `xml:"OgrUid"`
	GeoLat string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`

	// TODO, verify and correct
	XXXCodeId string `xml:"codeId" validate:"isdefault"`
	XXXregion string `xml:"region" validate:"isdefault"`

}

func (uid *ObsUid) String() string {
	return uidString(*uid)
}

func (uid *ObsUid) Hash() string {
	return uidHash(*uid)
}

type Obs struct {
	ObsUid ObsUid `xml:"ObsUid"`

	TxtName          string `xml:"txtName"`
	CodeType         string `xml:"codeType"`
	CodeGroup        string `xml:"codeGroup"`
	CodeLgt          string `xml:"codeLgt"`
	CodeMarking      string `xml:"codeMarking"`
	TxtDescrLgt      string `xml:"txtDescrLgt"`
	TxtDescrMarking  string `xml:"txtDescrMarking"`
	CodeDatum        string `xml:"codeDatum"`
	ValElev          string `xml:"valElev"`
	ValHgt           string `xml:"valHgt"`
	UomDistVer       string `xml:"uomDistVer"`
	CodeHgtAccuracy  string `xml:"codeHgtAccuracy"`
	ValRadius        string `xml:"valRadius"`
	UomRadius        string `xml:"uomRadius"`
	ObsUidLink       ObsUid `xml:"ObsUidLink"`
	CodeLinkType     string `xml:"codeLinkType"`
	DatetimeValidWef string `xml:"datetimeValidWef"`
	DatetimeValidTil string `xml:"datetimeValidTil"`
	TxtRmk           string `xml:"txtRmk"`

	// TODO, verify and correct
	XXOrgUid OgrUid `xml:"OrgUid" validate:"isdefault"`
}

func (f *Obs) Uid() FeatureUid {
	return &f.ObsUid
}
