package types

import (
	"errors"
)

func NewFeature(name string) (Feature, error) {
	switch name {
	// Org/Frequency
	case "Fqy":
		return &Fqy{}, nil
	case "Ser":
		return &Ser{}, nil
	case "Uni":
		return &Uni{}, nil
	// Airport
	case "Ahp":
		return &Ahp{}, nil
	case "Ful":
		return &Ful{}, nil
	case "Rwy":
		return &Rwy{}, nil
	case "Rdn":
		return &Rdn{}, nil
	case "Twy":
		return &Twy{}, nil
	case "Tly":
		return &Tly{}, nil
	case "Apn":
		return &Apn{}, nil
	// TODO
	case "Gsd":

	// navigation
	case "Dpn":
		return &Dpn{}, nil
	case "Dme":
		return &Dme{}, nil
	case "Mkr":
		return &Mkr{}, nil
	case "Ndb":
		return &Ndb{}, nil
	case "Tac":
		return &Tac{}, nil
	case "Vor":
		return &Vor{}, nil

	// airspace
	case "Ase":
		return &Ase{}, nil
	case "Abd":
		return &Abd{}, nil
	// missing Avx
	case "Adg":
		return &Adg{}, nil
	case "Sae":
		return &Sae{}, nil

	// Geographical border
	case "Gbr":
		return &Gbr{}, nil
	// Geographical border vertex // not a standalone feature
	case "Gbv":

	// OFM Label Marker
	case "Lbm":
	// OFM Label Marker Zoom level
	case "Lbz":
	// OFM Plate Achive?
	case "xt_Ppa":
		return &XtPpa{}, nil
	// OFM Platepackage
	case "xt_Plp":
		return &XtPlp{}, nil
	// OFM Procedure
	case "Prc":
	// Misc structure
	case "misc", "Misc": // TODO FIX
		return &Misc{}, nil
	default:
	}
	return nil, errors.New("unknown type " + name)
}
