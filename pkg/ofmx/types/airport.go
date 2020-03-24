package types

const (
	CODE_USAGE_LIMITATION_PERMIT = "PERMIT"
	CODE_USAGE_LIMITATION_FORBID = "FORBID"
)

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

func (uid *AhpUid) Ref() *Ahp {
	if uid.ref == nil {
		//TODO, needed? BUGON
		return nil
	}
	return uid.ref.(*Ahp)
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

	// References
	Ahu []*Ahu

	// TODO, drop, document
	XXTxtDescrSite       string `xml:"txtDescrSite" validate:"isdefault"`
	XXvalGeoidUndulation string `xml:"valGeoidUndulation" validate:"isdefault"`
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

func (uid *RwyUid) Ref() *Rwy {
	if uid.ref == nil {
		//TODO, needed? BUGON
		return nil
	}
	return uid.ref.(*Rwy)
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

	RwyUid   *RwyUid `xml:"RwyUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *RdnUid) String() string {
	return uidString(*uid)
}

func (uid *RdnUid) Hash() string {
	return uidHash(*uid)
}

func (uid *RdnUid) Ref() *Rdn {
	if uid.ref == nil {
		//TODO, needed? BUGON
		return nil
	}
	return uid.ref.(*Rdn)
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

	// References
	Rdd []*Rdd
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
	ValDist string `xml:"valDist"`
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

func (uid *RlsUid) Region() string {
	return uid.RdnUid.RwyUid.AhpUid.Region
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
	XXMid                 string      `xml:"mid,attr" validate:"isdefault"`
	XXusageLimitation     interface{} `xml:"usageLimitation" validate:"isdefault"`
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
	XtLabel interface{} `xml:"xt_label"`
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
	XtLabel interface{} `xml:"xt_label"`
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
	XtLabel interface{} `xml:"xt_label"`
}

func (f *Msc) Uid() FeatureUid {
	return &f.MscUid
}

// Aha - Airport Address

type AhaUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid   AhpUid `xml:"AhpUid"`
	CodeType string `xml:"codeType"`
	NoSeq    int    `xml:"noSeq"`
}

func (uid *AhaUid) String() string {
	return uidString(*uid)
}

func (uid *AhaUid) Hash() string {
	return uidHash(*uid)
}

type Aha struct {
	AhaUid     AhaUid `xml:"AhaUid"`
	TxtAddress string `xml:"txtAddress"`
	TxtRmk     string `xml:"txtRmk"`
}

func (f *Aha) Uid() FeatureUid {
	return &f.AhaUid
}

// Fato - Final approach and takeoff area

type FtoUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *FtoUid) String() string {
	return uidString(*uid)
}

func (uid *FtoUid) Hash() string {
	return uidHash(*uid)
}

type Fto struct {
	FtoUid FtoUid `xml:"FtoUid"`

	ValLen                  string `xml:"valLen"`
	ValWid                  string `xml:"valWid"`
	UomDim                  string `xml:"uomDim"`
	CodeComposition         string `xml:"codeComposition"`
	CodePreparation         string `xml:"codePreparation"`
	CodeCondSfc             string `xml:"codeCondSfc"`
	ValPcnClass             string `xml:"valPcnClass"`
	CodePcnPavementType     string `xml:"codePcnPavementType"`
	CodePcnPavementSubgrade string `xml:"codePcnPavementSubgrade"`
	CodePcnMaxTirePressure  string `xml:"codePcnMaxTirePressure"`
	CodePcnEvalMethod       string `xml:"codePcnEvalMethod"`
	TxtPcnNote              string `xml:"txtPcnNote"`
	ValSiwlWeight           string `xml:"valSiwlWeight"`
	UomSiwlWeight           string `xml:"uomSiwlWeight"`
	ValSiwlTirePressure     string `xml:"valSiwlTirePressure"`
	UomSiwlTirePressure     string `xml:"uomSiwlTirePressure"`
	ValAuwWeight            string `xml:"valAuwWeight"`
	TomAuwWeight            string `xml:"uomAuwWeight"`
	ZxtProfile              string `xml:"txtProfile"`
	TxtMarking              string `xml:"txtMarking"`
	CodeSts                 string `xml:"codeSts"`
	TxtRmk                  string `xml:"txtRmk"`

	// TODO, drop, document
	XXXmaxRotorDia string `xml:"maxRotorDia" validate:"isdefault"`
}

func (f *Fto) Uid() FeatureUid {
	return &f.FtoUid
}

// Fato Direction

type FdnUid struct {
	// TODO, temp allow mid
	Uid

	FtoUid   FtoUid `xml:"FtoUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *FdnUid) String() string {
	return uidString(*uid)
}

func (uid *FdnUid) Hash() string {
	return uidHash(*uid)
}

type Fdn struct {
	FdnUid FdnUid `xml:"FdnUid"`

	ValTrueBrg           string `xml:"valTrueBrg"`
	ValMagBrg            string `xml:"valMagBrg"`
	CodeTypeVasis        string `xml:"codeTypeVasis"`
	CodePsnVasis         string `xml:"codePsnVasis"`
	NoBoxVasis           string `xml:"noBoxVasis"`
	CodePortableVasis    string `xml:"codePortableVasis"`
	ValSlopeAngleGpVasis string `xml:"valSlopeAngleGpVasis"`
	ValMeht              string `xml:"valMeht"`
	UomMeht              string `xml:"uomMeht"`
	TxtRmk               string `xml:"txtRmk"`

	// TODO, drop, document
	XXXrearTakeOffSectorAvail      string             `xml:"rearTakeOffSectorAvail" validate:"isdefault"`
	XXXrearTakeoffSectorBrg        string             `xml:"rearTakeoffSectorBrg" validate:"isdefault"`
	XXXrearTakeoffSectorSlopeAngle string             `xml:"rearTakeoffSectorSlopeAngle" validate:"isdefault"`
	XXXcodingMode                  string             `xml:"codingMode" validate:"isdefault"`
	XXXnonIcaoValSlopeAngle        string             `xml:"nonIcaoValSlopeAngle" validate:"isdefault"`
	XXXnonIcaoValOpeningAngle      string             `xml:"nonIcaoValOpeningAngle" validate:"isdefault"`
	XXXnonIcaoMaxSectorLen         string             `xml:"nonIcaoMaxSectorLen" validate:"isdefault"`
	XXXiapInitialAlt               string             `xml:"iapInitialAlt" validate:"isdefault"`
	XXXgeoLat                      string             `xml:"geoLat" validate:"isdefault"`
	XXXgeoLong                     string             `xml:"geoLong" validate:"isdefault"`
	XXXgmlSectorExtend             XXXgmlSectorExtend `xml:"gmlSectorExtend" validate:"isdefault"`
	XXXgmlSectorCenterline         XXXgmlSectorExtend `xml:"gmlSectorCenterline" validate:"isdefault"`
}

type XXXgmlSectorExtend struct {
	XXXgmlPosList string `xml:"gmlPosList" validate:"isdefault"`
}

func (f *Fdn) Uid() FeatureUid {
	return &f.FdnUid
}

// Fato Direction Declared Distance

type FddUid struct {
	// TODO, temp allow mid
	Uid

	FdnUid        FdnUid `xml:"FdnUid"`
	CodeType      string `xml:"codeType"`
	CodeDayPeriod string `xml:"codeDayPeriod"`
}

func (uid *FddUid) String() string {
	return uidString(*uid)
}

func (uid *FddUid) Hash() string {
	return uidHash(*uid)
}

type Fdd struct {
	FddUid FddUid `xml:"FddUid"`

	ValDist string `xml:"valDist"`
	UomDist string `xml:"uomDist"`
	TxtRmk  string `xml:"txtRmk"`
}

func (f *Fdd) Uid() FeatureUid {
	return &f.FddUid
}

// Helipad (TLOF)

type TlaUid struct {
	// TODO, temp allow mid
	Uid

	AhpUid   AhpUid `xml:"AhpUid"`
	TxtDesig string `xml:"txtDesig"`
}

func (uid *TlaUid) String() string {
	return uidString(*uid)
}

func (uid *TlaUid) Hash() string {
	return uidHash(*uid)
}

type Tla struct {
	TlaUid TlaUid `xml:"TlaUid"`
	FtoUid FtoUid `xml:"FtoUid"`
	// TODO
	TxtDesig                string `xml:"txtDesig"`
	GeoLat                  string `xml:"geoLat"`
	GeoLong                 string `xml:"geoLong"`
	CodeDatum               string `xml:"codeDatum"`
	ValElev                 string `xml:"valElev"`
	UomDistVer              string `xml:"uomDistVer"`
	ValLen                  string `xml:"valLen"`
	ValWid                  string `xml:"valWid"`
	UomDim                  string `xml:"uomDim"`
	CodeComposition         string `xml:"codeComposition"`
	CodePreparation         string `xml:"codePreparation"`
	CodeCondSfc             string `xml:"codeCondSfc"`
	ValPcnClass             string `xml:"valPcnClass"`
	CodePcnPavementType     string `xml:"codePcnPavementType"`
	CodePcnPavementSubgrade string `xml:"codePcnPavementSubgrade"`
	CodePcnMaxTirePressure  string `xml:"codePcnMaxTirePressure"`
	CodePcnEvalMethod       string `xml:"codePcnEvalMethod"`
	TxtPcnNote              string `xml:"txtPcnNote"`
	ValSiwlWeight           string `xml:"valSiwlWeight"`
	UomSiwlWeight           string `xml:"uomSiwlWeight"`
	ValSiwlTirePressure     string `xml:"valSiwlTirePressure"`
	UomSiwlTirePressure     string `xml:"uomSiwlTirePressure"`
	ValAuwWeight            string `xml:"valAuwWeight"`
	UomAuwWeight            string `xml:"uomAuwWeight"`
	CodeClassHel            string `xml:"codeClassHel"`
	TxtMarking              string `xml:"txtMarking"`
	CodeSts                 string `xml:"codeSts"`
	TxtRmk                  string `xml:"txtRmk"`

	XXXvalTrueBrg string `xml:"valTrueBrg"`
}

func (f *Tla) Uid() FeatureUid {
	return &f.TlaUid
}

// Helipad (TLOF) Lightning

type TlsUid struct {
	// TODO, temp allow mid
	Uid

	TlaUid  TlaUid `xml:"TlaUid"`
	CodePsn string `xml:"codePsn"`
}

func (uid *TlsUid) String() string {
	return uidString(*uid)
}

func (uid *TlsUid) Hash() string {
	return uidHash(*uid)
}

type Tls struct {
	TlsUid TlsUid `xml:"TlsUid"`

	TxtDescr   string `xml:"txtDescr"`
	CodeIntst  string `xml:"codeIntst"`
	CodeColour string `xml:"codeColour"`
	TxtRmk     string `xml:"txtRmk"`
}

func (f *Tls) Uid() FeatureUid {
	return &f.TlsUid
}
