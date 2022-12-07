package controllers

import (
	"github.com/astaxie/beego"
)

// Repayment amount calculate
type BP000007Controller struct {
	beego.Controller
}

// @Title
// @Description  Repayment amount calculate
// @Param	body		body 	models.BP000007Request	"Input parameter"
// @Success 200  {object} models.BP000007Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000007Controller) Post() ()  {
}
