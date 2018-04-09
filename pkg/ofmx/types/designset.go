package types

type AisMapDesignSet struct {
	XMLName                 string                    `xml:"aisMapDesignSet"`
	Name                    string                    `xml:"name,attr"`
	Origin                  string                    `xml:"origin,attr"`
	Created                 string                    `xml:"created,attr"`
	Projection              string                    `xml:"projection,attr"`
	FlightInformationRegion string                    `xml:"flightInformationRegion,attr"`
	General                 AisMapDesignSetGeneral    `xml:"general"`
	Airspaces               []AisMapDesignSetAirspace `xml:"airspace"`
	// Airports
	Airports       []AisMapDesignSetAirport `xml:"airport"`
	Heliports      []AisMapDesignSetAirport `xml:"heliport"`
	Ultralightport []AisMapDesignSetAirport `xml:"ultralightport"`

	DVor         []AisMapDesignSetNav `xml:"DVOR"`
	Vor          []AisMapDesignSetNav `xml:"VOR"`
	Ndb          []AisMapDesignSetNav `xml:"NDB"`
	Dme          []AisMapDesignSetNav `xml:"DME"`
	Fix          []AisMapDesignSetNav `xml:"FIX"`
	Tcn          []AisMapDesignSetNav `xml:"TCN"`
	VfrGldr      []AisMapDesignSetNav `xml:"VFR-GLDR"`
	VfrRp        []AisMapDesignSetNav `xml:"VFR-RP"`
	VfrMrp       []AisMapDesignSetNav `xml:"VFR-MRP"`
	VfrEnr       []AisMapDesignSetNav `xml:"VFR-ENR"`
	MountainTop  []AisMapDesignSetNav `xml:"MOUNTAIN-TOP"`
	MountainPass []AisMapDesignSetNav `xml:"MOUNTAIN-PASS"`
	Town         []AisMapDesignSetNav `xml:"TOWN"`
}

type AisMapDesignSetGeneral struct {
	Appearance          AisMapDesignSetAppearance          `xml:"appearance"`
	Grid                AisMapDesignSetGrid                `xml:"grid"`
	CountryBorders      AisMapDesignSetCountryBorders      `xml:"countryBorders"`
	AirspaceGroundShade AisMapDesignSetAirspaceGroundShade `xml:"airspaceGroundShade"`
}

type AisMapDesignSetAppearance struct {
	BackgroundColor string `xml:"backgroundcolor,attr"`
}

type AisMapDesignSetGrid struct {
	Color string `xml:"color,attr"`
	Width string `xml:"width,attr"`
}

type AisMapDesignSetAirspaceGroundShade struct {
	Show  string `xml:"show,attr"`
	Color string `xml:"color,attr"`
}

type AisMapDesignSetCountryBorders struct {
	Show                       string `xml:"show,attr"`
	DashStyle                  string `xml:"dashStyle,attr"`
	SimplifyUnlessCorridorOfNm string `xml:"simplifyUnlessCorridorOfNm,attr"`
	MaxElementLengthOfNm       string `xml:"maxElementLengthOfNm,attr"`
	Color                      string `xml:"color,attr"`
}

type AisMapDesignSetAirspace struct {
	CodeType   string                        `xml:"codeType,attr"`
	Descr      string                        `xml:"descr,attr,omitempty"`
	Outline    AisMapDesignSetAirspaceStyle  `xml:"outline"`
	Halo       *AisMapDesignSetAirspaceStyle `xml:"halo"`
	Area       AisMapDesignSetAirspaceStyle  `xml:"area"`
	Bend       *AisMapDesignSetAirspaceStyle `xml:"bend"`
	Designator *AisMapDesignSetDesignator    `xml:"designator"`
}

type AisMapDesignSetAirspaceStyle struct {
	Show      string `xml:"show,attr"`
	Width     string `xml:"width,attr,omitempty"`
	Color     string `xml:"color,attr"`
	DashStyle string `xml:"dashStyle,attr,omitempty"`
}

type AisMapDesignSetDesignator struct {
	// type="areaBox" show="1">
	Type                  string `xml:"type,attr,omitempty"`
	Show                  string `xml:"show,attr,omitempty"`
	BorderDockingDistance string `xml:"borderDockingDistance,attr,omitempty"`
	Mode                  string `xml:"mode,attr,omitempty"`

	// SVG HERE!
	Svg string `xml:",innerxml"`
}

type AisMapDesignSetAirport struct {
	Style string `xml:"style,attr"`

	Icon       *AisMapDesignIcon         `xml:"icon"`
	Designator AisMapDesignSetDesignator `xml:"designator"`
}

type AisMapDesignSetNav struct {
	Style string `xml:"style,attr"`

	Icon       *AisMapDesignIcon         `xml:"icon"`
	Designator AisMapDesignSetDesignator `xml:"designator"`
}

type AisMapDesignIcon struct {
	Svg string `xml:",innerxml"`
}

// Labels
type AisMapLabelPositioning struct {
	XMLName string        `xml:"aisMapLabelPositioning"`
	Label   []AisMapLabel `xml:"label"`
}

type AisMapLabel struct {
	Type  string `xml:"type,attr"`
	UName string `xml:"uName,attr"`

	DeclutteringLayer []AisMapLabelDeclutteringLayer `xml:"declutteringLayer"`
}

type AisMapLabelDeclutteringLayer struct {
	Id string `xml:"id,attr"`

	Placement []AisMapLabelPlacement `xml:"placement"`
}

type AisMapLabelPlacement struct {
	TrackOrient    string `xml:"trackOrient,attr"`
	X              string `xml:"x,attr"`
	Y              string `xml:"y,attr"`
	OutlinePolyIdx string `xml:"outlinePolyIdx,attr"`
	Rotation       string `xml:"rotation,attr"`
	YFocus         string `xml:"yFocus,attr"`
	XFocus         string `xml:"xFocus,attr"`
}
