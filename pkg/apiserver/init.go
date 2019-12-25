package apiserver

import (
	lru "github.com/hashicorp/golang-lru"
	xlog "pgxs.io/chassis/log"
)

var log *xlog.Entry

func init() {
	log = xlog.New().Service("gopkg").Category("rest")
	var err error
	pkgCache, err = lru.New(50)
	if err != nil {
		log.Errorf("New Lru cache error")
	}
}
