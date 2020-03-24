package types

type PpaUid struct {
	RegionalUid
	PackageType            string `xml:"packageType"`
	EntityOfInterestCodeId string `xml:"entityOfInterestCodeId"`
	TxtPurpose             string `xml:"txtPurpose"`
	TxtPurposeLocLang      string `xml:"txtPurposeLocLang"`
	StatusRelease          string `xml:"statusRelease"`
}

func (uid *PpaUid) String() string {
	return uidString(*uid)
}

func (uid *PpaUid) Hash() string {
	return uidHash(*uid)
}

type Ppa struct {
	PpaUid       PpaUid       `xml:"PpaUid"`

	// TODO
	XXXtxtPurpose    string    `xml:"txtPurpose" validate:"isdefault"`
	XXXtxtPurposeLocLang    string    `xml:"txtPurposeLocLang" validate:"isdefault"`
	XtImageGroup XtImageGroup `xml:"xt_imageGroup" validate:"isdefault"`
	XXXstatusRelease    string    `xml:"statusRelease" validate:"isdefault"`
	XXXdisplayType    string    `xml:"displayType" validate:"isdefault"`
}

func (f *Ppa) Uid() FeatureUid {
	return &f.PpaUid
}

type PlpUid struct {
	// TODO, temp allow mid
	Uid
	PlateIdentifier        OrgUid    `xml:"plateIdentifier"`

	XtPpaUid PpaUid `xml:"PpaUid"`
}

func (uid *PlpUid) String() string {
	return uidString(*uid)
}

func (uid *PlpUid) Hash() string {
	return uidHash(*uid)
}

type XtImageGroup struct {
	Img             string `xml:"img"`
}

type BrStrip struct {
	Type            string `xml:"type"`
	BoxWidth        string `xml:"boxWidth"`
	BoxHeight       string `xml:"boxHeight"`
	TxtValue        string `xml:"txtValue"`
	TxtValueLocLang string `xml:"txtValueLocLang"`
	Orientation     string `xml:"orientation"`
	PointerGeoLat   string `xml:"pointerGeoLat"`
	PointerGeoLong  string `xml:"pointerGeoLong"`
	GmlPosList      string `xml:"gmlPosList"`
	Value           string `xml:"value"`
}

type FrameShot struct {
	GmlPosList      string `xml:"gmlPosList"`
}

type Plp struct {
	PlpUid                 PlpUid    `xml:"PlpUid"`
	TxtDesig               string    `xml:"txtDesig"`
	CodeType               string    `xml:"codeType"`
	Cat                    string    `xml:"cat"`
	VisScale               string    `xml:"visScale"`
	AspectRatio            string    `xml:"aspectRatio"`
	MaxPageHeight_mm       string    `xml:"maxPageHeight_mm"`
	MaxPageWidth_mm        string    `xml:"maxPageWidth_mm"`
	NbrFoldingMarksX       string    `xml:"NbrFoldingMarksX"`
	NbrFoldingMarksY       string    `xml:"NbrFoldingMarksY"`
	Frame                  string    `xml:"frame"`
	TxtDesignSet           string    `xml:"txtDesignSet"`
	BasemapTransparency    string    `xml:"basemapTransparency"`
	StaticLabelSource      string    `xml:"staticLabelSource"`
	VisSectionalLabels     string    `xml:"VisSectionalLabels"`
	FrameShot  			   FrameShot `xml:"frameShot"`
	VisGridMora            string    `xml:"visGridMora"`
	MoraNbrOfSecPerDegrLat string    `xml:"moraNbrOfSecPerDegrLat"`
	MoraNbrOfSecPerDegrLon string    `xml:"moraNbrOfSecPerDegrLon"`
	Projection             string    `xml:"projection"`
	StandardParallel1      string    `xml:"standardParallel1"`
	StandardParallel2      string    `xml:"standardParallel2"`
	BrStrip                []BrStrip `xml:"brStrip"`

	// TODO, document ??
	XXXBasemapOpacity    string    `xml:"basemapOpacity" validate:"isdefault"`
	XXXTerrainOpacity    string    `xml:"terrainOpacity" validate:"isdefault"`
	XXXzoomLevel    string    `xml:"zoomLevel" validate:"isdefault"`
	XXXbriefingStripPattern    string    `xml:"briefingStripPattern" validate:"isdefault"`
}

func (f *Plp) Uid() FeatureUid {
	return &f.PlpUid
}
