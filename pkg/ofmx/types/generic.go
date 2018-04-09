package types

import (
	"crypto/md5"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Uid struct {
	Mid string `xml:"mid,attr" hash:"ignore"`
}

type RegionalUid struct {
	Uid
	Region string `xml:"region,attr,omitempty" hash:"sure"`
	XtFir  string `xml:"xt_fir,attr,omitempty" hash:"ignore"`
}

type FeatureUid interface {
	String() string
	Hash() string
}

type Feature interface {
	Uid() FeatureUid
}

func intHash(uid interface{}) []string {
	val := make([]string, 0)
	s := reflect.ValueOf(uid)
	t := s.Type()
	for i := 0; i < s.NumField(); i++ {
		t := t.Field(i)
		f := s.Field(i)
		h := t.Tag.Get("hash")
		if h == "ignore" {
			continue
		}
		switch f.Kind() {
		case reflect.String:
			val = append(val, f.String())
			//fmt.Printf("%d: %s %s = %v\n", i,
			//	t.Name, f.Type(), f.String())
		case reflect.Int:
			val = append(val, fmt.Sprintf("%d", f.Int()))
		case reflect.Struct:
			val = append(val, intHash(f.Interface())...)
		default:
			fmt.Println(t)
			fmt.Println("skipping")
		}
	}
	return val
}

func uidString(uid interface{}) string {
	val := intHash(uid)
	return fmt.Sprintf("%s|%s", reflect.ValueOf(uid).Type().Name(), strings.Join(val, "|"))
}

func uidHash(uid interface{}) string {
	s := uidString(uid)
	h := md5.New()
	io.WriteString(h, s)
	m := fmt.Sprintf("%x", h.Sum(nil))

	return fmt.Sprintf("%s-%s-%s-%s-%s", m[:8], m[8:12], m[12:16], m[16:20], m[20:32])
}
