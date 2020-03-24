package types

import (
	"encoding/xml"

	"log"
	"reflect"
)

// <OFMX-Snapshot version="1.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="https://schema.openflightmaps.org/0/OFMX-Snapshot.xsd" effective="2019-11-28T13:32:23" origin="ofmx editor" created="2019-11-28T13:32:23" namespace="210444d1-4576-e92d-0983-4669182a8c04">

type OfmxSnapshot struct {
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
			log.Printf("error in %s: %v\n", feature.Uid().String(), err)
			return err
		}
		*s = append(*s, feature)
	} else {
		log.Println("Unknown feature: " + start.Name.Local)
		d.Skip()
	}
	return nil
}
