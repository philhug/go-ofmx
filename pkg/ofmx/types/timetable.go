package types

type Timetable struct {
	CodeWorkHr   string `xml:"codeWorkHr"`
	TxtRmkWorkHr string `xml:"txtRmkWorkHr"`
	Timsh   Timsh `xml:"Timsh"`
}

type Timsh struct {
	CodeTimeRef   string `xml:"codeTimeRef"`
	DateValidWef string `xml:"dateValidWef"`
	DateValidTil string `xml:"dateValidTil"`
	CodeDay string `xml:"codeDay"`
	CodeDayTil string `xml:"codeDayTil"`
	TimeWef string `xml:"timeWef"`
	CodeEventWef string `xml:"codeEventWef"`
	TimeRelEventWef string `xml:"timeRelEventWef"`
	CodeCombWef string `xml:"codeCombWef"`
	TimeTil string `xml:"timeTil"`
	CodeEventTil string `xml:"codeEventTil"`
	TimeRelEventTil string `xml:"timeRelEventTil"`
	CodeCombTil string `xml:"codeCombTil"`

	// TODO, REMOVE INVALID
	XXxt_AlignSummerSavingT string `xml:"xt_AlignSummerSavingT"`
}