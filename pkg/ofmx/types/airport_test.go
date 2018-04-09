package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAhpUidHash(t *testing.T) {
	uid := AhpUid{RegionalUid: RegionalUid{Region: "LF"}, CodeId: "LFMD"}

	assert.Equal(t, "AhpUid|LF|LFMD", uid.String())
	assert.Equal(t, "4b601836-9131-2bf6-6bd6-8cf5c48d367d", uid.Hash())
}
