package apiserver

import "github.com/emicklei/go-restful"

//RegistPackageWebservice regist usercenter webservices
func RegistPackageWebservice() {
	restful.DefaultContainer.Add(NewPackageService().WebService())
}
