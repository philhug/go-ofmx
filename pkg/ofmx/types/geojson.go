package types

import (
	flatten "github.com/doublerebel/bellows"
	"github.com/paulmach/go.geo"
	"github.com/paulmach/go.geojson"

	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func convertFromDMS(f float64) float64 {
	d := math.Floor(f / 10000)
	m := math.Mod(math.Floor(f/100), 100)
	s := math.Mod(f, 100)

	r := math.Round((d + m/60 + s/3600) * 1000)
	return r / 1000
}

func parseCoord(in string) (float64, error) {
	v := in[:len(in)-1]
	d := string(in[len(in)-1])

	var dir = 1.0
	if d == "W" || d == "S" {
		dir = -1.0
	}
	f, _ := strconv.ParseFloat(v, 64)

	// TODO HACK!

	if f > 180 {
		f = convertFromDMS(f)
	}
	if f > 180 {
		fmt.Println("invalid coord:" + in)
		return f * dir, errors.New("invalid coord:" + in)
	}
	return f * dir, nil
}

func parseLongLat(long, lat string, codeDatum string) (*geo.Point, error) {
	if codeDatum != "" && codeDatum != "WGE" {
		return nil, errors.New("Unsupported codeDatum: " + codeDatum)
	}
	longp, err := parseCoord(long)
	if err != nil {
		return nil, err
	}
	latp, err := parseCoord(lat)
	if err != nil {
		return nil, err
	}
	p := geo.NewPoint(longp, latp)
	return p, nil
}

// navigation
func (f *Dpn) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.DpnUid.GeoLong, f.DpnUid.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}
	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Dpn"
	p.Properties["CodeId"] = f.DpnUid.CodeId
	return p, nil
}
func (f *Dme) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.DmeUid.GeoLong, f.DmeUid.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}

	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Dme"
	p.Properties["CodeId"] = f.DmeUid.CodeId
	return p, nil
}
func (f *Mkr) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.MkrUid.GeoLong, f.MkrUid.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}
	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Mkr"
	p.Properties["CodeId"] = f.MkrUid.CodeId
	return p, nil
}
func (f *Ndb) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.NdbUid.GeoLong, f.NdbUid.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}
	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Ndb"
	p.Properties["CodeId"] = f.NdbUid.CodeId
	return p, nil
}
func (f *Tac) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.TacUid.GeoLong, f.TacUid.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}
	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Tac"
	p.Properties["CodeId"] = f.TacUid.CodeId
	return p, nil
}
func (f *Vor) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.VorUid.GeoLong, f.VorUid.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}
	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Vor"
	p.Properties["CodeId"] = f.VorUid.CodeId
	return p, nil
}

// AD
func (f *Ahp) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	pt, err := parseLongLat(f.GeoLong, f.GeoLat, f.CodeDatum)
	if err != nil {
		return nil, err
	}
	p := pt.ToGeoJSON()
	p.Properties["Type"] = "Ahp"
	p.Properties["CodeId"] = f.AhpUid.CodeId
	return p, nil
}

func (f *Twy) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	return f.XtSurface.GeoJson(fmap)
}

func (f *Tly) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	return f.XtSurface.GeoJson(fmap)
}

func (f *Apn) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	return f.XtSurface.GeoJson(fmap)
}

func (f *Misc) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	return f.XtSurface.GeoJson(fmap)
}

func (f *XtSurface) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	psn := strings.Split(f.GmlPosList, " ")
	path := geo.NewPath()
	for i := 0; i < len(psn); i++ {
		longlat := strings.Split(psn[i], ",")
		if len(longlat) == 2 {
			pt, err := parseLongLat(longlat[0], longlat[1], "WGE")
			if err != nil {
				return nil, err
			}
			path.Push(pt)
		}
	}
	p := path.ToGeoJSON()
	return p, nil
}

// helper
func GetNearestPointOffset(pt *geo.Point, points *geo.Path) int {
	var nd = math.Inf(1)
	//var np geo.Point
	var ni int = 0
	for i, p := range points.PointSet {
		d := p.DistanceFrom(pt)
		if d < nd {
			nd = d
			//np = p
			ni = i
		}
	}
	return ni
}

// Airspace

func Reverse(a geo.PointSet) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}
func GetSlice(pts geo.PointSet, start, end int) geo.PointSet {
	if start <= end {
		return pts[start:end]
	} else {
		b := append(geo.PointSet(nil), pts[end:start]...)
		Reverse(b)
		return b
	}
}

func (f *Abd) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	poly := geo.NewPath()
	var queued *geo.Path
	var start int

	for _, avx := range f.Avx {
		psn, err := parseLongLat(avx.GeoLong, avx.GeoLat, avx.CodeDatum)
		if err != nil {
			fmt.Println(psn)
			return nil, err
		}
		if queued != nil {
			end := GetNearestPointOffset(psn, queued)
			poly.PointSet = append(poly.PointSet, GetSlice(queued.PointSet, start, end)...)
			// TODO
			//*poly = append(poly, geom.LineString...)
			queued = nil
		}

		switch avx.CodeType {
		case "GRC":
			poly.Push(psn)

		case "FNT": // BORDER
			f := fmap[avx.GbrUid.String()]
			if f == nil {
				return nil, errors.New("GBR not found:" + avx.GbrUid.String())
			} else {
				gbr := f.(*Gbr)
				queued, err = gbr.GeoJsonPath(fmap)
				if err != nil {
					return nil, err
				}
				start = GetNearestPointOffset(psn, queued)
			}
		case "CWA", "CCA":
			fmt.Println("UNKNOWN CodeType: " + avx.CodeType)
			fmt.Println(avx)
			// TODO, silly circle
			poly.Push(psn)
		default:
			fmt.Println("UNKNOWN CodeType: " + avx.CodeType)

		}
	}

	//p := geojson.NewLineStringFeature(poly)
	//p := poly.ToGeoJSON()
	p := ToGeoJSONPolygon(poly)
	p.Properties["Type"] = "Abd"
	p.Properties["Region"] = f.AbdUid.AseUid.Region
	p.Properties["CodeType"] = f.AbdUid.AseUid.CodeType
	p.Properties["CodeId"] = f.AbdUid.AseUid.CodeId

	p.Properties["fill"] = "#ffbaba"

	ase := fmap[f.AbdUid.AseUid.String()].(*Ase)
	p.Properties["TxtName"] = ase.TxtName

	return p, nil
}

// Geographical Borders

func (f *Gbr) GeoJsonPath(fmap FeatureMap) (*geo.Path, error) {
	poly := geo.NewPath()
	for _, gbv := range f.Gbv {
		switch gbv.CodeType {
		case "GRC":
			psn, err := parseLongLat(gbv.GeoLong, gbv.GeoLat, gbv.CodeDatum)
			if err != nil {
				return nil, err
			}
			poly.Push(psn)
		case "END":
			// TODO: verify if it matches start
			//fmt.Printf("CodeType End: %s\n", gbv)
		default:
			fmt.Println("UNKNOWN CodeType: " + gbv.CodeType)
		}
	}
	return poly, nil
}

// ToGeoJSON creates a new geojson feature with a linestring geometry
// containing all the points.
func ToGeoJSONPolygon(p *geo.Path) *geojson.Feature {
	coords := make([][]float64, 0, len(p.PointSet))

	for _, p := range p.PointSet {
		coords = append(coords, []float64{p[0], p[1]})
	}

	return geojson.NewPolygonFeature([][][]float64{coords})
}

func (f *Gbr) GeoJson(fmap FeatureMap) (*geojson.Feature, error) {
	g, err := f.GeoJsonPath(fmap)
	gf := ToGeoJSONPolygon(g)
	gf.Properties["Type"] = "Gbr"
	gf.Properties["TxtName"] = f.GbrUid.TxtName

	return gf, err
}

func FillProperties(f GeoFeature, gf *geojson.Feature) {
	gf.Properties = flatten.Flatten(f)
}

type GeoFeature interface {
	Uid() FeatureUid // From Feature
	GeoJson(fmap FeatureMap) (*geojson.Feature, error)
}

type FeatureMap map[string]Feature
