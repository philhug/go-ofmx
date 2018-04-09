package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVorUidHash(t *testing.T) {
	uid := VorUid{RegionalUid: RegionalUid{Region: "LF"}, CodeId: "AGN", GeoLat: "435316.90N", GeoLong: "0005222.30E"}

	assert.Equal(t, "VorUid|LF|AGN|435316.90N|0005222.30E", uid.String())
	assert.Equal(t, "fe6e3186-80aa-7c48-4557-edc84d9419f8", uid.Hash())
}
