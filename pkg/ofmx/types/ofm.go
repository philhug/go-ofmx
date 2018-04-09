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
	XtImageGroup string   `xml:"xt_imageGroup"`
}

func (f *XtPpa) Uid() FeatureUid {
	return &f.XtPpaUid
}

type XtPlpUid struct {
	XtPpaUid XtPpaUid `xml:"PpaUid"`
}

func (uid *XtPlpUid) String() string {
	return uidString(*uid)
}

func (uid *XtPlpUid) Hash() string {
	return uidHash(*uid)
}

type BrStrip struct {
	Type            XtPlpUid `xml:"type"`
	BoxWidth        XtPlpUid `xml:"boxWidth"`
	BoxHeight       XtPlpUid `xml:"boxHeight"`
	TxtValue        XtPlpUid `xml:"txtValue"`
	TxtValueLocLang XtPlpUid `xml:"txtValueLocLang"`
	Orientation     XtPlpUid `xml:"orientation"`
	GmlPosList      XtPlpUid `xml:"gmlPosList"`
}

type XtPlp struct {
	XtPlpUid               XtPlpUid  `xml:"XtPlpUid"`
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
	VisGridMora            string    `xml:"visGridMora"`
	MoraNbrOfSecPerDegrLat string    `xml:"moraNbrOfSecPerDegrLat"`
	MoraNbrOfSecPerDegrLon string    `xml:"moraNbrOfSecPerDegrLon"`
	Projection             string    `xml:"projection"`
	StandardParallel1      string    `xml:"standardParallel1"`
	StandardParallel2      string    `xml:"standardParallel2"`
	FrameShot              string    `xml:"frameShot"`
	BrStrip                []BrStrip `xml:"brStrip"`
}

func (f *XtPlp) Uid() FeatureUid {
	return &f.XtPlpUid
}
