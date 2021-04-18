package api

import (
	"gis-hcsb/models"
	"gis-hcsb/util"
	"github.com/kataras/iris/v12"
	"strings"
)

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c iris.Context) {
	var req LoginReq
	c.ReadJSON(&req)
	req.Username = strings.Trim(req.Username, " ")
	u := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	if !u.LoginCheck() {
		c.JSON(util.Error(iris.Map{
			"msg": "用户名已存在",
		}))
		return
	}
	c.JSON(util.Ok())
}
