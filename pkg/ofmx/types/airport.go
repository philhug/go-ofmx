package types

// Ahp - Airport

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
	// Aht TODO

	// TODO, drop, document
	XXTxtDescrSite     string `xml:"txtDescrSite" validate:"isdefault"`
	XXvalGeoidUndulation     string `xml:"valGeoidUndulation" validate:"isdefault"`
}

func (f *Ahp) Uid() FeatureUid {
	return &f.AhpUid
}

// Rwy - Runway

type RwyUid struct {
	// TODO, temp allow mid
	Uid

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
	RwyUid    RwyUid `xml:"RwyUid"`
	ValLen    string `xml:"valLen"`
	ValWid    string `xml:"valWid"`
	UomDimRwy string `xml:"uomDimRwy"`
	SurfaceCharacteristics
	ValLenStrip string `xml:"valLenStrip"`
	ValWidStrip string `xml:"valWidStrip"`
	UomDimStrip string `xml:"uomDimStrip"`
	CodeSts     string `xml:"codeSts"`
	TxtRmk      string `xml:"txtRmk"`
}

func (f *Rwy) Uid() FeatureUid {
	return &f.RwyUid
}

// Rdn - Runway Direction

type RdnUid struct {
	// TODO, temp allow mid
	Uid

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
	BUGGeoLon string `xml:"geoLon" validate:"isdefault"`

	ValTrueBrg string `xml:"valTrueBrg"`
	ValMagBrg  string `xml:"valMagBrg"`
	ValElevTdz string `xml:"valElevTdz"`
	UomElevTdz string `xml:"uomElevTdz"`
	TxtRmk     string `xml:"txtRmk"`
}

func (f *Rdn) Uid() FeatureUid {
	return &f.RdnUid
}

// Rdd - Runway Direction Declared Distance

type RddUid struct {
	// TODO, temp allow mid
	Uid

	RdnUid        RdnUid `xml:"RdnUid"`
	CodeType      string `xml:"codeType"`
	CodeDayPeriod string `xml:"codeDayPeriod"`
}

func (uid *RddUid) String() string {
	return uidString(*uid)
}

func (uid *RddUid) Hash() string {
	return uidHash(*uid)
}

type Rdd struct {
	RddUid  RddUid `xml:"RddUid"`
	ValDist int    `xml:"valDist"`
	UomDist string `xml:"uomDist"`
	TxtRmk  string `xml:"txtRmk"`
}

func (f *Rdd) Uid() FeatureUid {
	return &f.RddUid
}

// Rls - Runway Direction Lighting

type RlsUid struct {
	// TODO, temp allow mid
	Uid

	RdnUid  RdnUid `xml:"RdnUid"`
	CodePsn string `xml:"codePsn"`
}

func (uid *RlsUid) String() string {
	return uidString(*uid)
}

func (uid *RlsUid) Hash() string {
	return uidHash(*uid)
}

type Rls struct {
	RlsUid     RlsUid `xml:"RlsUid"`
	TxtDescr   string `xml:"txtDescr"`
	CodeColour string `xml:"codeColour"`
	TxtRmk     string `xml:"txtRmk"`
}

func (f *Rls) Uid() FeatureUid {
	return &f.RlsUid
}

// Ahu - Airport Usage

type AhuUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid AhpUid `xml:"AhpUid"`
}

func (uid *AhuUid) String() string {
	return uidString(*uid)
}

func (uid *AhuUid) Hash() string {
	return uidHash(*uid)
}

type Ahu struct {
	AhuUid          AhuUid            `xml:"AhuUid"`
	UsageLimitation []UsageLimitation `xml:"UsageLimitation"`

	// TODO, THIS IS WRONG!!!
	XXMid string `xml:"mid,attr" validate:"isdefault"`
	XXusageLimitation interface{} `xml:"usageLimitation" validate:"isdefault"`
	XXcodeUsageLimitation interface{} `xml:"codeUsageLimitation" validate:"isdefault"`
}

func (f *Ahu) Uid() FeatureUid {
	return &f.AhuUid
}

type UsageLimitation struct {
	CodeUsageLimitation string           `xml:"codeUsageLimitation"`
	UsageCondition      []UsageCondition `xml:"UsageCondition"`
	// TODO
	//Timetable     Timetable `xml:"Timetable"`
	TxtRmk string `xml:"txtRmk"`
}

type UsageCondition struct {
	AircraftClass []AircraftClass `xml:"AircraftClass"`
	FlightClass   []FlightClass   `xml:"FlightClass"`
}

type AircraftClass struct {
	CodeType string `xml:"codeType"`
}

type FlightClass struct {
	CodeRule    string `xml:"codeRule"`
	CodeMil     string `xml:"codeMil"`
	CodeOrigin  string `xml:"codeOrigin"`
	CodePurpose string `xml:"codePurpose"`
}

// Ahs - Airport Ground Service
type AhsUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid   AhpUid `xml:"AhpUid"`
	CodeType string `xml:"codeType"`
}

func (uid *AhsUid) String() string {
	return uidString(*uid)
}

func (uid *AhsUid) Hash() string {
	return uidHash(*uid)
}

type Ahs struct {
	AhsUid      AhsUid `xml:"AhsUid"`
	TxtDescrFac string `xml:"txtDescrFac"`
	// TODO, timetable
	//Ast         Timetable `xml:"Ast"`
	TxtRmk string `xml:"txtRmk"`
}

func (f *Ahs) Uid() FeatureUid {
	return &f.AhsUid
}

// Aga - Airport Ground Service Address

type AgaUid struct {
	// TODO, temp allow mid
	Uid

	AhsUid   AhsUid `xml:"AhsUid"`
	CodeType string `xml:"codeType"`
	NoSeq    string `xml:"noSeq"`
}

func (uid *AgaUid) String() string {
	return uidString(*uid)
}

func (uid *AgaUid) Hash() string {
	return uidHash(*uid)
}

type Aga struct {
	AgaUid     AgaUid `xml:"AgaUid"`
	TxtAddress string `xml:"txtAddress"`
	TxtRmk     string `xml:"txtRmk"`
}

func (f *Aga) Uid() FeatureUid {
	return &f.AgaUid
}

// Twy - Taxiway

type TwyUid struct {
	// TODO, temp allow mid
	Uid

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
	TwyUid   TwyUid `xml:"TwyUid"`
	CodeType string `xml:"codeType"`
	ValWid   string `xml:"valWid"`
	UomWid   string `xml:"uomWid"`
	SurfaceCharacteristics
	CodeStrength     string    `xml:"codeStrength"`     // DEPRECATED
	TxtDescrStrength string    `xml:"txtDescrStrength"` // DEPRECATED
	CodeSts          string    `xml:"codeSts"`
	TxtMarking       string    `xml:"txtMarking"`
	TxtRmk           string    `xml:"txtRmk"`
	XtSurface        XtSurface `xml:"xt_surface"`

	// TODO document
	XtLabel        interface{} `xml:"xt_label"`
}

func (f *Twy) Uid() FeatureUid {
	return &f.TwyUid
}

// Tly - Taxiway Lighting

type TlyUid struct {
	// TODO, temp allow mid
	Uid

	TwyUid  TwyUid `xml:"TwyUid"`
	CodePsn string `xml:"codePsn"`
}

func (uid *TlyUid) String() string {
	return uidString(*uid)
}

func (uid *TlyUid) Hash() string {
	return uidHash(*uid)
}

type Tly struct {
	TlyUid    TlyUid    `xml:"TlyUid"`
	TxtRmk    string    `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`
}

func (f *Tly) Uid() FeatureUid {
	return &f.TlyUid
}

// Apn - Apron TODO

type ApnUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid  AhpUid `xml:"AhpUid"`
	TxtName string `xml:"txtName"`
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
	TxtRmk    string    `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`

	// TODO document
	XtLabel        interface{} `xml:"xt_label"`
}

func (f *Apn) Uid() FeatureUid {
	return &f.ApnUid
}

// Ful - Fuel

type FulUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid  AhpUid `xml:"AhpUid"`
	CodeCat string `xml:"codeCat"`
}

func (uid *FulUid) String() string {
	return uidString(*uid)
}

func (uid *FulUid) Hash() string {
	return uidHash(*uid)
}

type Ful struct {
	FulUid FulUid `xml:"FulUid"`

	TxtDescr string `xml:"txtDescr"`
	TxtRmk   string `xml:"txtRmk"`
}

func (f *Ful) Uid() FeatureUid {
	return &f.FulUid
}

// Msc - Miscellaneous

type MscUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid  AhpUid `xml:"AhpUid"`
	TxtName string `xml:"txtName"`
}

func (uid *MscUid) String() string {
	return uidString(*uid)
}

func (uid *MscUid) Hash() string {
	return uidHash(*uid)
}

type XtSurface struct {
	GmlPosList string `xml:"gmlPosList"`
}

type Msc struct {
	MscUid    MscUid    `xml:"MscUid"`
	Type      string    `xml:"type"`
	TxtRmk    string    `xml:"txtRmk"`
	XtSurface XtSurface `xml:"xt_surface"`

	// TODO document
	XtLabel        interface{} `xml:"xt_label"`
}

func (f *Msc) Uid() FeatureUid {
	return &f.MscUid
}
