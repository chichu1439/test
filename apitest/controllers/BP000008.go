package controllers

import (
	"github.com/astaxie/beego"
)

// Repayment
type BP000008Controller struct {
	beego.Controller
}

// @Title
// @Description  Repayment
// @Param	body		body 	models.BP000008Request	"Input parameter"
// @Success 200  {object} models.BP000008Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000008Controller) Post() ()  {
}
