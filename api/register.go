package api

import (
	"gis-hcsb/models"
	"gis-hcsb/util"
	"github.com/kataras/iris/v12"
	"strings"
)

type RegisterReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c iris.Context) {
	var req RegisterReq
	_ = c.ReadJSON(&req)
	req.Username = strings.Trim(req.Username, " ")
	if req.Username == "" {
		c.JSON(util.Error(iris.Map{
			"msg": "fuck it",
		}))
		return
	}
	u := models.User{
		Username: req.Username,
		Password: req.Password,
	}
	if u.CheckUsername() {
		c.JSON(util.Error(iris.Map{
			"msg": "用户名已存在",
		}))
		return
	}
	if !u.Insert() {
		c.JSON(util.Error(iris.Map{
			"msg": "register fail",
		}))
		return
	}
	c.JSON(util.Ok())
}
