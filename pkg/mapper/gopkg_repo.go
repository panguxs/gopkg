package mapper

import (
	lru "github.com/hashicorp/golang-lru"
	"github.com/jinzhu/gorm"
	"pgxs.io/chassis"
)

var pkgCache *lru.Cache

//Package package ORM table struct
type Package struct {
	gorm.Model
	Name        string
	Source      string
	Description string
}

//TableName db table name
func (pkg Package) TableName() string {
	return "gopkg_package"
}

//PackageMapper pkg mapper
type PackageMapper struct{}

//NewPackageMapper new pkg mapper
func NewPackageMapper() *PackageMapper {
	return new(PackageMapper)
}

//Create pkg to db
func (pm PackageMapper) Create(pkg *Package) {
	chassis.DB().Create(pkg)
	log.Infof("%+v", pkg)
}

// FindByName find Package by name
func (pm PackageMapper) FindByName(name string) *Package {
	var pkg Package

	chassis.DB().Where("name = ?", name).First(&pkg)
	log.Infof("%+v", pkg)
	return &pkg
}
