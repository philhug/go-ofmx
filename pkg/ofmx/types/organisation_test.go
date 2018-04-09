package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniUidHash(t *testing.T) {
	exampleRegion := RegionalUid{Region: "LF"}
	exampleUni := UniUid{RegionalUid: exampleRegion, TxtName: "STRASBOURG APP"}
	assert.Equal(t, "UniUid|LF|STRASBOURG APP", exampleUni.String())
	assert.Equal(t, "8e2aed9c-d02b-a22f-a91e-98bc2888f9f9", exampleUni.Hash())
}

func TestSerSerHash(t *testing.T) {
	exampleRegion := RegionalUid{Region: "LF"}
	exampleUni := UniUid{RegionalUid: exampleRegion, TxtName: "STRASBOURG APP"}
	exampleSer := SerUid{UniUid: exampleUni, CodeType: "APP", NoSeq: 1}

	assert.Equal(t, "SerUid|LF|STRASBOURG APP|APP|1", exampleSer.String())
	assert.Equal(t, "9d858c0b-3317-9a16-8e62-213fbaa85f76", exampleSer.Hash())
}

func TestFqyUidHash(t *testing.T) {
	exampleRegion := RegionalUid{Region: "LF"}
	exampleUni := UniUid{RegionalUid: exampleRegion, TxtName: "STRASBOURG APP"}
	exampleSer := SerUid{UniUid: exampleUni, CodeType: "APP", NoSeq: 1}
	exampleFqy := FqyUid{SerUid: exampleSer, ValFreqTrans: "134.575"}

	assert.Equal(t, "FqyUid|LF|STRASBOURG APP|APP|1|134.575", exampleFqy.String())
	assert.Equal(t, "2a1276d3-f704-2967-171b-de573e5dac47", exampleFqy.Hash())
}
