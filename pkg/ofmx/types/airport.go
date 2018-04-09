package types

import ()

type AhpUid struct {
	RegionalUid
	CodeId string `xml:"codeId"`
}

func (uid *AhpUid) String() string {
	return uidString(*uid)
}

func (uid *AhpUid) Hash() string {
	return uidHash(*uid)
}

type Ahp struct {
	AhpUid           AhpUid `xml:"AhpUid"`
	OrgUid           OrgUid `xml:"OrgUid"`
	TxtName          string `xml:"txtName"`
	CodeIcao         string `xml:"codeIcao"`
	CodeIata         string `xml:"codeIata"`
	CodeGps          string `xml:"codeGps"`
	CodeType         string `xml:"codeType"`
	GeoLat           string `xml:"geoLat"`
	GeoLong          string `xml:"geoLong"`
	CodeDatum        string `xml:"codeDatum"`
	ValElev          string `xml:"valElev"`
	UomDistVer       string `xml:"uomDistVer"`
	TxtNameCitySer   string `xml:"txtNameCitySer"`
	ValMagVar        string `xml:"valMagVar"`
	DateMagVar       string `xml:"dateMagVar"`
	ValMagVarChg     string `xml:"valMagVarChg"`
	ValRefT          string `xml:"valRefT"`
	UomRefT          string `xml:"uomRefT"`
	TxtNameAdmin     string `xml:"txtNameAdmin"`
	ValTransitionAlt string `xml:"valTransitionAlt"`
	UomTransitionAlt string `xml:"uomTransitionAlt"`
	TxtRmk           string `xml:"txtRmk"`
	// Aht
}

func (f *Ahp) Uid() FeatureUid {
	return &f.AhpUid
}

type RwyUid struct {
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *RwyUid) String() string {
	return uidString(*uid)
}

func (uid *RwyUid) Hash() string {
	return uidHash(*uid)
}

type SurfaceCharacteristics struct {
	CodeComposition         string `xml:"codeComposition,omitempty"`
	CodePreparation         string `xml:"codePreparation,omitempty"`
	ValPcnClass             string `xml:"valPcnClass,omitempty"`
	CodePcnPavementType     string `xml:"codePcnPavementType,omitempty"`
	CodePcnPavementSubgrade string `xml:"codePcnPavementSubgrade,omitempty"`
	CodePcnMaxTirePressure  string `xml:"codePcnMaxTirePressure,omitempty"`
	CodePcnEvalMethod       string `xml:"codePcnEvalMethod,omitempty"`
	TxtPcnNote              string `xml:"txtPcnNote,omitempty"`
}

type Rwy struct {
	RwyUid                  RwyUid `xml:"RwyUid"`
	ValLen                  string `xml:"valLen"`
	ValWid                  string `xml:"valWid"`
	UomDimRwy               string `xml:"uomDimRwy"`
	SurfaceCharacteristics
	ValLenStrip             string `xml:"valLenStrip"`
	ValWidStrip             string `xml:"valWidStrip"`
	UomDimStrip             string `xml:"uomDimStrip"`
	CodeSts                 string `xml:"codeSts"`
	TxtRmk                  string `xml:"txtRmk"`
}

func (f *Rwy) Uid() FeatureUid {
	return &f.RwyUid
}

type RdnUid struct {
	RwyUid   RwyUid `xml:"RwyUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *RdnUid) String() string {
	return uidString(*uid)
}

func (uid *RdnUid) Hash() string {
	return uidHash(*uid)
}

type Rdn struct {
	RdnUid  RdnUid `xml:"RdnUid"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
	// TODO, sometimes called geoLon
	BUGGeoLon string `xml:"geoLon"`

	ValTrueBrg string `xml:"valTrueBrg"`
	ValMagBrg  string `xml:"valMagBrg"`
	ValElevTdz string `xml:"valElevTdz"`
	UomElevTdz string `xml:"uomElevTdz"`
	TxtRmk     string `xml:"txtRmk"`
}

func (f *Rdn) Uid() FeatureUid {
	return &f.RdnUid
}

type TwyUid struct {
	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *TwyUid) String() string {
	return uidString(*uid)
}

func (uid *TwyUid) Hash() string {
	return uidHash(*uid)
}

type Twy struct {
	TwyUid TwyUid `xml:"TwyUid"`
	CodeType string `xml:"codeType"`
	ValWid   string `xml:"valWid"`
	UomWid   string `xml:"uomWid"`
	SurfaceCharacteristics
	CodeStrength     string `xml:"codeStrength"`     // DEPRECATED
	TxtDescrStrength string `xml:"txtDescrStrength"` // DEPRECATED
	CodeSts          string `xml:"codeSts"`
	TxtMarking       string `xml:"txtMarking"`
	TxtRmk           string `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`
}

func (f *Twy) Uid() FeatureUid {
	return &f.TwyUid
}

type TlyUid struct {
	TwyUid   TwyUid `xml:"TwyUid"`
	CodePsn string `xml:"codePsn"`
}

func (uid *TlyUid) String() string {
	return uidString(*uid)
}

func (uid *TlyUid) Hash() string {
	return uidHash(*uid)
}

type Tly struct {
	TlyUid TlyUid `xml:"TlyUid"`
	TxtRmk           string `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`
}

func (f *Tly) Uid() FeatureUid {
	return &f.TlyUid
}

// Apron Apn TODO

type ApnUid struct {
	AhpUid   AhpUid `xml:"AhpUid"`
	TextName string `xml:"textName"`
}

func (uid *ApnUid) String() string {
	return uidString(*uid)
}

func (uid *ApnUid) Hash() string {
	return uidHash(*uid)
}

type Apn struct {
	ApnUid ApnUid `xml:"ApnUid"`
	SurfaceCharacteristics
	TxtRmk           string `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`
}

func (f *Apn) Uid() FeatureUid {
	return &f.ApnUid
}

// Fuel Ful
type FulUid struct {
	//Uid
	AhpUid  AhpUid `xml:"AhpUid"`
	CodeCat string `xml:"codeCat"`
}

func (uid *FulUid) String() string {
	return uidString(uid)
}

func (uid *FulUid) Hash() string {
	return uidHash(uid)
}

type Ful struct {
	FulUid FulUid `xml:"FulUid"`

	TxtDescr string `xml:"txtDescr"`
	TxtRmk   string `xml:"txtRmk"`
}

func (f *Ful) Uid() FeatureUid {
	return &f.FulUid
}

type MiscUid struct {
	RegionalUid
	AhpUid  AhpUid `xml:"AhpUid"`
	TxtName string `xml:"txtName"`
}

func (uid *MiscUid) String() string {
	return uidString(*uid)
}

func (uid *MiscUid) Hash() string {
	return uidHash(*uid)
}

type XtSurface struct {
	GmlPosList string `xml:"gmlPosList"`
}

type Misc struct {
	MiscUid   MiscUid   `xml:"miscUid"`
	Type      string    `xml:"type"`
	TxtRmk    string    `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`
}

func (f *Misc) Uid() FeatureUid {
	return &f.MiscUid
}
