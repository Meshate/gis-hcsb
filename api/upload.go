package api

import (
	"encoding/json"
	"gis-hcsb/models"
	"gis-hcsb/util"
	"github.com/kataras/iris/v12"
)

type Point [2]float64

type Router struct {
	Start     Point  `json:"start"`
	End       Point  `json:"end"`
	Type      string `json:"type"`
	StartName string `json:"startName,omitempty"`
	EndName   string `json:"endName,omitempty"`
	Ts        int64  `json:"ts"`
}

type uploadReq struct {
	Username string `json:"username"`
	Route    Router `json:"route"`
}

func Upload(c iris.Context) {
	var req uploadReq
	c.ReadJSON(&req)
	rt, _ := json.Marshal(req.Route)
	rd := util.Bytes2String(rt)
	rec := models.Record{
		Username: req.Username,
		Road:     rd,
	}
	if !rec.Insert() {
		c.JSON(util.Error("insert error"))
		return
	}
	c.JSON(util.Ok())
}
