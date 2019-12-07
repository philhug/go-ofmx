package types

type XtPpaUid struct {
	RegionalUid
	PackageType            string `xml:"packageType"`
	EntityOfInterestCodeId string `xml:"entityOfInterestCodeId"`
	TxtPurpose             string `xml:"txtPurpose"`
	TxtPurposeLocLang      string `xml:"txtPurposeLocLang"`
	StatusRelease          string `xml:"statusRelease"`
}

func (uid *XtPpaUid) String() string {
	return uidString(*uid)
}

func (uid *XtPpaUid) Hash() string {
	return uidHash(*uid)
}

type XtPpa struct {
	XtPpaUid     XtPpaUid `xml:"PpaUid"`
	XtImageGroup XtImageGroup   `xml:"xt_imageGroup"`
}

func (f *XtPpa) Uid() FeatureUid {
	return &f.XtPpaUid
}

type PlpUid struct {
	// TODO, temp allow mid
	Uid

	XtPpaUid XtPpaUid `xml:"PpaUid"`
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

type XtPlp struct {
	PlpUid                 PlpUid    `xml:"PlpUid"`
	PlateIdentifier        OrgUid    `xml:"plateIdentifier"`
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
	XXBasemapOpacity    string    `xml:"basemapOpacity" validate:"isdefault"`
	XXTerrainOpacity    string    `xml:"terrainOpacity"`
}

func (f *XtPlp) Uid() FeatureUid {
	return &f.PlpUid
}
