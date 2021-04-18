package api

import (
	. "gis-hcsb"
	"gis-hcsb/util"
	"github.com/kataras/iris/v12"
)

func GpAss(c iris.Context) {
	b, _ := c.GetBody()
	Ass = util.Bytes2String(b)
	c.JSON(util.Ok())
}
