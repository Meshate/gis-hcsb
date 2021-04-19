package api

import (
	"gis-hcsb/util"
	"github.com/kataras/iris/v12"
)

var ass struct {
	Utc         float64 `json:"UTC"`
	Latitude    float64 `json:"latitude"`
	Ns          string  `json:"N_S"`
	Longitude   float64 `json:"longitude"`
	Ew          string  `json:"E_W"`
	GPS_status  int     `json:"GPS_status"`
	Satellite_n string  `json:"satellite_n"`
	Hasl        float64 `json:"HASL"`
}

func GpAss(c iris.Context) {
	c.ReadJSON(&ass)
	c.JSON(util.Ok())
}

func Gpass(c iris.Context) {
	c.JSON(util.OkData(ass))
}
