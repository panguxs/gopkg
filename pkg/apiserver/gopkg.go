package apiserver

import (
	"html/template"

	"github.com/emicklei/go-restful"
	lru "github.com/hashicorp/golang-lru"
	"pgxs.io/chassis"

	"pgxs.io/gopkg/pkg/apiserver/tmpl"
	"pgxs.io/gopkg/pkg/mapper"

	"pgxs.io/gopkg/pkg/config"
)

var pkgCache *lru.Cache

//PackageService package service
type PackageService struct {
	cache *lru.Cache
}

//NewPackageService new package service
func NewPackageService() *PackageService {

	ps := new(PackageService)
	var err error
	ps.cache, err = lru.New(50)
	if err != nil {
		log.Error("Create LUR cache failed. error=", err)
	}
	return ps
}

func (ps PackageService) getPackageFromCache(pkgName string) *mapper.Package {
	var pkg mapper.Package
	//get from cache
	if iPkg, ok := pkgCache.Get(pkgName); ok {
		pkg = iPkg.(mapper.Package)
		log.Debug("Get from cache")
	} else {
		pkgP := findPkg(pkgName)
		pkg = *pkgP
		//Set package to cache
		pkgCache.Add(pkgName, pkg)
	}
	return &pkg
}

// WebService  xx
func (ps PackageService) WebService() *restful.WebService {

	ws := new(restful.WebService)
	tags := []string{"gopkg"}
	ws.Path("/package")
	ws.Route(ws.GET("/{*}").To(ps.render).Doc("Go package").
		Metadata(chassis.KeyOpenAPITags, tags).
		Param(ws.PathParameter("*", "package name").
			Required(true).DataType("string")).
		Param(ws.QueryParameter("go-get", "go get").
			DataType("integer").DefaultValue("1")).
		Produces("text/plain"))
	return ws
}

//PackageView package view struct
type PackageView struct {
	Host   string
	Source string
	Name   string
}

// Render render pkg
func (ps PackageService) render(req *restful.Request, resp *restful.Response) {
	if req.QueryParameter("go-get") != "1" {
		resp.WriteHeader(400)
		return
	}
	pkgName := req.PathParameter("*")

	pkg := ps.getPackageFromCache(pkgName)

	log.Debugf("pkg %+v", pkg)
	if &pkg == nil || pkg.Name == "" || pkg.Source == "" {
		resp.WriteHeader(404)
		return
	}
	p := &PackageView{
		Host:   config.Pkg().Host,
		Source: pkg.Source,
		Name:   pkg.Name,
	}
	// you might want to cache compiled templates
	text, err := tmpl.Asset("pkg.html")
	if err != nil {
		log.Error(err)
	}
	tmpl, err := template.New("pkgTemplate").Parse(string(text))
	log.Debugf("tmpl text: %s", text)
	if err != nil {
		log.Fatalf("Template gave: %s", err)
	}
	resp.AddHeader("Content-Type", "text/plain")
	tmpl.Execute(resp.ResponseWriter, p)
}

// 从数据库查询包信息
func findPkg(iPath string) *mapper.Package {
	log.Debug(iPath)
	pkg := mapper.NewPackageMapper().FindByName(iPath)
	return pkg
}
