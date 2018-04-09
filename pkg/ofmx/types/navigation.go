package types

type DpnUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *DpnUid) String() string {
	return uidString(*uid)
}

func (uid *DpnUid) Hash() string {
	return uidHash(*uid)
}

type Dpn struct {
	DpnUid    DpnUid `xml:"DpnUid"`
	CodeDatum string `xml:"codeDatum"`
	CodeType  string `xml:"codeType"`
	TxtName   string `xml:"txtName"`
	TxtRmk    string `xml:"txtRmk"`
	// AhpUid_codeId TODO ofm
}

func (f *Dpn) Uid() FeatureUid {
	return &f.DpnUid
}

type DmeUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *DmeUid) String() string {
	return uidString(*uid)
}

func (uid *DmeUid) Hash() string {
	return uidHash(*uid)
}

type Dme struct {
	DmeUid      DmeUid `xml:"DmeUid"`
	OrgUid      OrgUid `xml:"OrgUid"`
	TxtName     string `xml:"txtName"`
	CodeChannel string `xml:"codeChannel"`
	CodeDatum   string `xml:"codeDatum"`
	ValElev     string `xml:"valElev"`
	UomDistVer  string `xml:"uomDistVer"`
	// Dtt TODO
	TxtRmk string `xml:"txtRmk"`
}

func (f *Dme) Uid() FeatureUid {
	return &f.DmeUid
}

type MkrUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *MkrUid) String() string {
	return uidString(*uid)
}

func (uid *MkrUid) Hash() string {
	return uidHash(*uid)
}

type Mkr struct {
	MkrUid MkrUid `xml:"MkrUid"`
	OrgUid OrgUid `xml:"OrgUid"`
	//IlsUid    OrgUid `xml:"IlsUid"` // TODO
	CodePsnIls string `xml:"codePsnIls"`
	ValFreq    string `xml:"valFreq"`
	UomFreq    string `xml:"uomFreq"`
	TxtName    string `xml:"txtName"`
	CodeDatum  string `xml:"codeDatum"`
	ValElev    string `xml:"valElev"`
	UomDistVer string `xml:"uomDistVer"`
	// Mtt TODO
	TxtRmk string `xml:"txtRmk"`
}

func (f *Mkr) Uid() FeatureUid {
	return &f.MkrUid
}

type NdbUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *NdbUid) String() string {
	return uidString(*uid)
}

func (uid *NdbUid) Hash() string {
	return uidHash(*uid)
}

type Ndb struct {
	NdbUid     NdbUid `xml:"NdbUid"`
	OrgUid     OrgUid `xml:"OrgUid"`
	TxtName    string `xml:"txtName"`
	ValFreq    string `xml:"valFreq"`
	UomFreq    string `xml:"uomFreq"`
	CodeDatum  string `xml:"codeDatum"`
	ValElev    string `xml:"valElev"`
	UomDistVer string `xml:"uomDistVer"`
	// Ntt TODO
	TxtRmk string `xml:"txtRmk"`
}

func (f *Ndb) Uid() FeatureUid {
	return &f.NdbUid
}

type TacUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *TacUid) String() string {
	return uidString(*uid)
}

func (uid *TacUid) Hash() string {
	return uidHash(*uid)
}

type Tac struct {
	TacUid      TacUid `xml:"TacUid"`
	OrgUid      OrgUid `xml:"OrgUid"`
	TxtName     string `xml:"txtName"`
	CodeChannel string `xml:"codeChannel"`
	CodeDatum   string `xml:"codeDatum"`
	ValElev     string `xml:"valElev"`
	UomDistVer  string `xml:"uomDistVer"`
	// Ttt TODO
	TxtRmk string `xml:"txtRmk"`
}

func (f *Tac) Uid() FeatureUid {
	return &f.TacUid
}

type VorUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *VorUid) String() string {
	return uidString(*uid)
}

func (uid *VorUid) Hash() string {
	return uidHash(*uid)
}

type Vor struct {
	VorUid         VorUid `xml:"VorUid"`
	OrgUid         OrgUid `xml:"OrgUid"`
	TxtName        string `xml:"txtName"`
	CodeType       string `xml:"codeType"`
	ValFreq        string `xml:"valFreq"`
	UomFreq        string `xml:"uomFreq"`
	CodeTypeNorth  string `xml:"codeTypeNorth"`
	ValDeclination string `xml:"valDeclination"`
	ValMagVar      string `xml:"valMagVar"`
	DateMagVar     string `xml:"dateMagVar"`
	ValMagVarChg   string `xml:"valMagVarChg"`
	CodeDatum      string `xml:"codeDatum"`
	ValElev        string `xml:"valElev"`
	UomDistVer     string `xml:"uomDistVer"`
	// vtt
	TxtRmk string `xml:"txtRmk"`
}

func (f *Vor) Uid() FeatureUid {
	return &f.VorUid
}