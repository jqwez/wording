package util_test

import (
	"testing"

	util "github.com/jqwez/wording/utils"
	"github.com/stretchr/testify/assert"
)

func TestGetPublicPath(t *testing.T) {
	public, err := util.GetPublicPath()
	assert.Nil(t, err)
	assert.NotNil(t, public)
}
