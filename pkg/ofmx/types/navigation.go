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
	DpnUid      DpnUid `xml:"DpnUid"`
	AhpUidAssoc AhpUid `xml:"AhpUidAssoc"`
	CodeDatum   string `xml:"codeDatum"`
	CodeType    string `xml:"codeType"`
	TxtName     string `xml:"txtName"`
	TxtRmk      string `xml:"txtRmk"`

	// TODO, drop, invalid
	XXvalCrs        string `xml:"valCrs" validate:"isdefault"`
	XXvalHgt        string `xml:"valHgt" validate:"isdefault"`
	XXuomDistVer    string `xml:"uomDistVer" validate:"isdefault"`
	XXuomFreq       string `xml:"uomFreq" validate:"isdefault"`
	XXcodeTypeNorth string `xml:"codeTypeNorth" validate:"isdefault"`
	XXAhpUid_codeId string `xml:"AhpUid_codeId" validate:"isdefault"`
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
	VorUid      VorUid `xml:"VorUid"`
	TxtName     string `xml:"txtName"`
	CodeChannel string `xml:"codeChannel"`
	CodeDatum   string `xml:"codeDatum"`
	ValElev     string `xml:"valElev"`
	UomDistVer  string `xml:"uomDistVer"`
	// Dtt TODO
	TxtRmk string `xml:"txtRmk"`

	// TODO, drop, invalid
	XXValFreq       string `xml:"valFreq" validate:"isdefault"`
	XXuomFreq       string `xml:"uomFreq" validate:"isdefault"`
	XXuomElev       string `xml:"uomElev" validate:"isdefault"`
	XXcodeTypeNorth string `xml:"codeTypeNorth" validate:"isdefault"`
	XXvalMagVar     string `xml:"valMagVar" validate:"isdefault"`
	XXValMagVarChg  string `xml:"valMagVarChg" validate:"isdefault"`
	XXDateMagVar    string `xml:"dateMagVar" validate:"isdefault"`
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

	// TODO, drop, NDB doesn't have
	XXCodeTypeNorth string `xml:"codeTypeNorth" validate:"isdefault"`
	XXValMagVar     string `xml:"valMagVar" validate:"isdefault"`
	XXValMagVarChg  string `xml:"valMagVarChg" validate:"isdefault"`
	XXDateMagVar    string `xml:"dateMagVar" validate:"isdefault"`
	XXuomElev       string `xml:"uomElev" validate:"isdefault"`
}

func (f *Ndb) Uid() FeatureUid {
	return &f.NdbUid
}

type TcnUid struct {
	RegionalUid
	CodeId  string `xml:"codeId"`
	GeoLat  string `xml:"geoLat"`
	GeoLong string `xml:"geoLong"`
}

func (uid *TcnUid) String() string {
	return uidString(*uid)
}

func (uid *TcnUid) Hash() string {
	return uidHash(*uid)
}

type Tcn struct {
	TcnUid      TcnUid `xml:"TcnUid"`
	OrgUid      OrgUid `xml:"OrgUid"`
	TxtName     string `xml:"txtName"`
	CodeChannel string `xml:"codeChannel"`
	CodeDatum   string `xml:"codeDatum"`
	ValElev     string `xml:"valElev"`
	UomDistVer  string `xml:"uomDistVer"`
	// Ttt TODO
	TxtRmk string `xml:"txtRmk"`
	// TODO, drop invalid
	XXXvalMagVar    string `xml:"valMagVar" validate:"isdefault"`
	XXXdateMagVar   string `xml:"dateMagVar" validate:"isdefault"`
	XXXvalMagVarChg string `xml:"valMagVarChg" validate:"isdefault"`
	XXXVorUid       VorUid `xml:"VorUid" validate:"isdefault"`
	XXXvalGhostFreq string `xml:"valGhostFreq" validate:"isdefault"`
	XXXuomGhostFreq string `xml:"uomGhostFreq" validate:"isdefault"`
}

func (f *Tcn) Uid() FeatureUid {
	return &f.TcnUid
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

	// TODO, drop invalid
	XXvalMagDecl string `xml:"valMagDecl" validate:"isdefault"`
	XXUomElev    string `xml:"uomElev" validate:"isdefault"`
}

func (f *Vor) Uid() FeatureUid {
	return &f.VorUid
}
