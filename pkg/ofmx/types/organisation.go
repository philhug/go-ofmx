package types

import ()

type OrgUid struct {
	RegionalUid
	TxtName string `xml:"txtName"`
}

func (uid *OrgUid) String() string {
	return uidString(*uid)
}

func (uid *OrgUid) Hash() string {
	return uidHash(*uid)
}

type Org struct {
	OrgUid   OrgUid
	CodeId   string `xml:"codeId"`
	CodeType string `xml:"codeType"`
	TxtRmk   string `xml:"txtRmk"`
}

type UniUid struct {
	RegionalUid
	TxtName string `xml:"txtName"`
}

func (uid *UniUid) String() string {
	return uidString(*uid)
}

func (uid *UniUid) Hash() string {
	return uidHash(*uid)
}

type Uni struct {
	// Attributes
	Source string `xml:"source,attr,omitempty"`
	XtFir  string `xml:"xt_fir,attr,omitempty"`

	UniUid    UniUid `xml:"UniUid"`
	OrgUid    UniUid `xml:"OrgUid"`
	AhpUid    AhpUid `xml:"AhpUid"`
	CodeType  string `xml:"codeType"`
	CodeClass string `xml:"codeClass"`
	TxtRmk    string `xml:"txtRmk"`
}

func (f *Uni) Uid() FeatureUid {
	return &f.UniUid
}

type SerUid struct {
	//Uid
	UniUid   UniUid `xml:"UniUid"`
	CodeType string `xml:"codeType"`
	NoSeq    int    `xml:"noSeq"`
}

func (uid *SerUid) String() string {
	return uidString(*uid)
}

func (uid *SerUid) Hash() string {
	return uidHash(*uid)
}

type Ser struct {
	SerUid SerUid `xml:"SerUid"`
	// Stt
	TxtRmk string `xml:"txtRmk"`
}

func (f *Ser) Uid() FeatureUid {
	return &f.SerUid
}

type FqyUid struct {
	//Uid
	SerUid       SerUid `xml:"SerUid"`
	ValFreqTrans string `xml:"valFreqTrans"`
}

func (uid *FqyUid) String() string {
	return uidString(*uid)
}

func (uid *FqyUid) Hash() string {
	return uidHash(*uid)
}

type Cdl struct {
	TxtCallSign string `xml:"txtCallSign"`
	CodeLang    string `xml:"codeLang"`
}

type Fqy struct {
	FqyUid     FqyUid `xml:"FqyUid"`
	ValFreqRec string `xml:"valFreqRec"`
	UomFreq    string `xml:"uomFreq"`
	CodeType   string `xml:"codeType"`
	// Ftt
	TxtRmk string `xml:"txtRmk"`
	Cdl    Cdl    `xml:"Cdl"`
}

func (f *Fqy) Uid() FeatureUid {
	return &f.FqyUid
}
