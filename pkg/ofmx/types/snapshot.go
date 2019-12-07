package types

import (
	"encoding/xml"

	"fmt"
	"reflect"
)

// <AIXM-Snapshot xsi="http://www.aixm.aero/schema/4.5/AIXM-Snapshot.xsd" version="4.5 + OFM extensions of version 0.1 " effective="2014-01-07T13:17:32" origin="F-OPS data distribution client" created="2014-01-07T12:17:32">

type AixmSnapshot struct {
	//XMLName string
	Xsi string `xml:"xsi,attr"`
	NoNamespaceSchemaLocation string `xml:"noNamespaceSchemaLocation,attr"`

	Version   string `xml:"version,attr"`
	Origin    string `xml:"origin,attr"`
	Namespace string `xml:"namespace,attr"`
	Created   string `xml:"created,attr"`
	Effective string `xml:"effective,attr"`

	Features FeatureList `xml:",any"`
}

type FeatureList []Feature

func (s FeatureList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	for _, f := range s {
		start.Name.Local = reflect.TypeOf(f).Elem().Name()
		e.EncodeElement(f, start)
	}

	return nil
}

func (s *FeatureList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	feature, err := NewFeature(start.Name.Local)
	if err == nil {
		err := d.DecodeElement(&feature, &start)
		if err != nil {
			//fmt.Printf("error in %s: %v\n", feature.Uid().String(), err)
			return err
		}
		*s = append(*s, feature)
	} else {
		fmt.Println("Unknown feature: " + start.Name.Local)
		d.Skip()
	}
	return nil
}
