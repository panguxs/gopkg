package mapper

import (
	xlog "pgxs.io/chassis/log"
)

var log *xlog.Entry

func init() {
	log = xlog.New().Service("gopkg").Category("mapper")

}
