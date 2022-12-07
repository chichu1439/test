package controllers

import (
	"github.com/astaxie/beego"
)

// Get order and repayment plan
type BP000001Controller struct {
	beego.Controller
}

// @Title
// @Description  Get order and repayment plan
// @Param	body		body 	models.BP000001Request	"Input parameter"
// @Success 200  {object} models.BP000001Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000001Controller) Post() {
}
