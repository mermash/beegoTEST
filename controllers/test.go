package controllers

import (
	"user/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about object
type TestController struct {
	beego.Controller
}

// @Title TestEndpoint
// @Description Tests the API
// @Success 200 {object} models.Test
// @Failure 403 body is empty
// @router /hello [get]
func (t *TestController) Test() {
	Response := models.TestFunction()
	t.Data["json"] = Response
	t.ServeJSON()
}
