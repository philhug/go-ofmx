package types

import (
	"errors"
	"log"
)

func NewFeature(name string) (Feature, error) {
	// Org/Frequency
	switch name {
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
	case "Aha":
		return &Aha{}, nil

	// Heliports
	case "Fto":
		return &Fto{}, nil
	case "Fdn":
		return &Fdn{}, nil
	case "Fdd":
		return &Fdd{}, nil
	case "Tla":
		return &Tla{}, nil
	case "Tls":
		return &Tls{}, nil
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

	// Obstacle Group
	case "Ogr":
		//return &Ogr{}, nil

	// Obstacle
	case "Obs":
		//return &Obs{}, nil

	// OFM Label Marker
	case "Lbm":
		return &Lbm{}, nil
	// OFM Plate Achive?
	case "Ppa":
		//TODO
		//return &Ppa{}, nil
	// OFM Platepackage
	case "Plp":
		//TODO
		//return &Plp{}, nil
	// OFM Procedure
	case "Prc":
		// TODO
		//return &Prc{}, nil
	// Misc structure
	case "Msc":
		return &Msc{}, nil
	default:
	}
	return nil, errors.New("unknown type " + name)
}

func UpdateReferences(fmap FeatureMap) {
	for _, f := range fmap {

		// loop through types
		switch v := f.(type) {
		case *Rdd:
			// set self-reference
			v.RddUid.ref = f

			rdn := fmap[v.RddUid.RdnUid.Hash()].(*Rdn)
			rdn.Rdd = append(rdn.Rdd, v)

		case *Rdn:
			// set self-reference
			v.RdnUid.ref = f

			// replace references
			rwy, ok := fmap[v.RdnUid.RwyUid.OriginalMid()].(*Rwy)
			if ok {
				v.RdnUid.RwyUid = &rwy.RwyUid
			} else {
				log.Printf("Rwy reference not found: %s\n", v.RdnUid.RwyUid.TxtDesig)
			}

		case *Rwy:
			// set self-reference
			v.RwyUid.ref = f
		case *Ahp:
			// set self-reference
			v.AhpUid.ref = f
		case *Ahu:
			// backlink
			ahp := fmap[v.AhuUid.AhpUid.Hash()].(*Ahp)
			ahp.Ahu = append(ahp.Ahu, v)
			// set self-reference
			v.AhuUid.ref = f
		case *Fqy:
			// set self-reference
			v.FqyUid.ref = f

			// replace references
			ser, ok := fmap[v.FqyUid.SerUid.OriginalMid()].(*Ser)
			if ok {
				v.FqyUid.SerUid = &ser.SerUid
			} else {
				log.Printf("Ser reference not found: %s\n", v.FqyUid.SerUid.UniUid.TxtName)
			}
		case *Ser:
			// set self-reference
			v.SerUid.ref = f

			// replace references
			uni, ok := fmap[v.SerUid.UniUid.OriginalMid()].(*Uni)
			if ok {
				v.SerUid.UniUid = &uni.UniUid
			} else {
				log.Printf("Uni reference not found: %s: %s/%s %s\n", v.SerUid.UniUid.TxtName, v.SerUid.UniUid.Hash(), v.SerUid.UniUid.OriginalMid(), v.SerUid.UniUid.String())
			}

		case *Uni:
			// set self-reference
			v.UniUid.ref = f

			// replace references
			if v.AhpUid == nil {
				log.Printf("Missing Ahp Reference\n")
			} else {
				ahp, ok := fmap[v.AhpUid.OriginalMid()].(*Ahp)
				if ok {
					v.AhpUid = &ahp.AhpUid
				} else {
					log.Printf("Ahp reference not found: %s\n", v.AhpUid.CodeId)
				}
			}

			org, ok := fmap[v.OrgUid.OriginalMid()].(*Org)
			if ok {
				v.OrgUid = &org.OrgUid
			} else {
				log.Printf("Org reference not found: %s\n", v.OrgUid.TxtName)
			}
		}

	}
}
