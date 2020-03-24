package types

import (
	"encoding/xml"
)

type OfmxShapeExtensions struct {
	//XMLName string
	Xsi                       string `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string `xml:"noNamespaceSchemaLocation,attr"`

	Version   string `xml:"version,attr"`
	Origin    string `xml:"origin,attr"`
	Namespace string `xml:"namespace,attr"`
	Created   string `xml:"created,attr"`
	Effective string `xml:"effective,attr"`

	Shapes ShapeMap `xml:",any"`
}

type Shape struct {
	GmlPosList string `xml:"gmlPosList"`
}

type ShapeMap map[string]Shape

func (s ShapeMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	shape := Shape{}
	err := d.DecodeElement(&shape, &start)
	if err != nil {
		return err
	}
	s[start.Attr[0].Value] = shape
	return nil
}
