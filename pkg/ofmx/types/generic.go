package types

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"reflect"
	"strings"
)

const (
	REGION_NONE = ""
)

type Uid struct {
	Mid string `xml:"mid,attr" hash:"ignore"`
	DBUid  string `xml:"dbUid,attr" hash:"ignore"`

	// reference to object itself
	ref interface{} `hash:"ignore"`
}

type RegionalUid struct {
	Uid
	Region string `xml:"region,attr,omitempty" hash:"required"`
	XtFir  string `xml:"xt_fir,attr,omitempty" hash:"ignore"`
}

type FeatureUid interface {
	String() string
	Hash() string
	OriginalMid() string
}

type Feature interface {
	Uid() FeatureUid
}

func (uid *Uid) OriginalMid() string {
	return uid.Mid
}

func (uid *Uid) String() string {
	return uidString(*uid)
}

func (uid *Uid) Hash() string {
	return uidHash(*uid)
}

func (uid *RegionalUid) OriginalMid() string {
	return uid.Mid
}

func (uid *RegionalUid) String() string {
	return uidString(*uid)
}

func (uid *RegionalUid) Hash() string {
	return uidHash(*uid)
}

func (uid *RegionalUid) _Region() string {
	return uid.Region
}

func intHash(uid interface{}, all bool) []string {
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
		if h == "optional" && !all {
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
			val = append(val, intHash(f.Interface(), all)...)
		case reflect.Ptr:
			ptr := f.Elem().Interface()
			val = append(val, intHash(ptr, all)...)
		default:
			log.Println(t)
			log.Println("skipping")
		}
	}
	return val
}

func uidString(uid interface{}) string {
	val := intHash(uid, false)
	return fmt.Sprintf("%s|%s", reflect.ValueOf(uid).Type().Name(), strings.Join(val, "|"))
}

func uidHash(uid interface{}) string {
	val := intHash(uid, false)
	s := fmt.Sprintf("%s|%s", reflect.ValueOf(uid).Type().Name(), strings.Join(val, "|"))
	h := md5.New()
	io.WriteString(h, s)
	m := fmt.Sprintf("%x", h.Sum(nil))

	return fmt.Sprintf("%s-%s-%s-%s-%s", m[:8], m[8:12], m[12:16], m[16:20], m[20:32])
}
