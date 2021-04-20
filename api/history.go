package api

import (
	"encoding/json"
	"gis-hcsb/models"
	"gis-hcsb/util"
	"github.com/kataras/iris/v12"
)

func History(c iris.Context) {
	user := c.FormValue("username")
	u := models.Record{
		Username: user,
	}
	ret := u.GetByUsername()
	var roads []Router
	for _, i := range ret {
		var r Router
		json.Unmarshal([]byte(i.Road), &r)
		roads = append(roads, r)
	}
	c.JSON(util.OkData(roads))
}
