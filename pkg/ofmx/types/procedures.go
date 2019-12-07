package types

type PrcUid struct {
	RegionalUid
	AhpUid AhpUid `xml:"AhpUid"`
	CodeId string `xml:"codeId"`
}

func (uid *PrcUid) String() string {
	return uidString(*uid)
}

func (uid *PrcUid) Hash() string {
	return uidHash(*uid)
}

type BezTrajectory struct {
	GmlPosList string `xml:"gmlPosList"`
}

type Leg struct {
}

type Prc struct {
	PrcUid                 PrcUid        `xml:"PrcUid"`
	TxtName                string        `xml:"txtName"`
	CodeType               string        `xml:"codeType"`
	UsageType              string        `xml:"usageType"`
	BezTrajectory          BezTrajectory `xml:"beztrajectory"`
	BeztrajectoryAlternate BezTrajectory `xml:"beztrajectoryAlternate"`
	SceletonPath           BezTrajectory `xml:"sceletonPath"`
	Leg                    []Leg         `xml:"Leg"`

	// TODO, FIX incorrect data
	XXXtFir  string `xml:"xt_fir,attr,omitempty" hash:"ignore" validate:"isdefault"`
}

func (f *Prc) Uid() FeatureUid {
	return &f.PrcUid
}
