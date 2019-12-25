package mapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PackageMapper_Create(t *testing.T) {
	pm := NewPackageMapper()
	pkg := &Package{
		Name:   "test4",
		Source: "test",
	}

	pm.Create(pkg)
}

func Test_PackageMapper_FindByName(t *testing.T) {
	pm := NewPackageMapper()

	pkg := pm.FindByName("test4")

	assert.NotNil(t, pkg)
}
