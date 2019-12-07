package types

import (
	"errors"
)

func NewFeature(name string) (Feature, error) {
	switch name {
	// Org/Frequency
	case "Org":
		return &Org{}, nil
	case "Uni":
		return &Uni{}, nil
	case "Ser":
		return &Ser{}, nil
	case "Fqy":
		return &Fqy{}, nil

	// Airport
	case "Ahp":
		return &Ahp{}, nil
	case "Ful":
		return &Ful{}, nil
	case "Rwy":
		return &Rwy{}, nil
	case "Rdn":
		return &Rdn{}, nil
	case "Rdd":
		return &Rdd{}, nil
	case "Rls":
		return &Rls{}, nil
	case "Ahu":
		return &Ahu{}, nil
	case "Ahs":
		return &Ahs{}, nil
	case "Aga":
		return &Aga{}, nil
	case "Twy":
		return &Twy{}, nil
	case "Tly":
		return &Tly{}, nil
	case "Apn":
		return &Apn{}, nil
	// TODO
	case "Gsd":

	// Navigation
	case "Dpn":
		return &Dpn{}, nil
	case "Dme":
		return &Dme{}, nil
	case "Mkr":
		return &Mkr{}, nil
	case "Ndb":
		return &Ndb{}, nil
	case "Tcn":
		return &Tcn{}, nil
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
		return &Lbm{}, nil
	// OFM Plate Achive?
	case "xt_Ppa":
		//TODO
		return &XtPpa{}, nil
	// OFM Platepackage
	case "xt_Plp":
		//TODO
		return &XtPlp{}, nil
	// OFM Procedure
	case "Prc":
		// TODO
		return &Prc{}, nil
	// Misc structure
	case "Msc":
		return &Msc{}, nil
	default:
	}
	return nil, errors.New("unknown type " + name)
}
