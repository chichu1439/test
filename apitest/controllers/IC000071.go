package controllers

import (
	"github.com/astaxie/beego"
)

// trial repayment plan
type IC000071Controller struct {
	beego.Controller
}

// @Title
// @Description   trial repayment plan
// @Param	body		body 	models.IC000071Request	"Input parameter"
// @Success 200  {object} models.IC000071Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *IC000071Controller) Post() ()  {
}
