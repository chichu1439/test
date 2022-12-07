package controllers

import (
	"github.com/astaxie/beego"
)

// Get repayment plan
type BP000006Controller struct {
	beego.Controller
}

// @Title
// @Description  Get repayment plan
// @Param	body		body 	models.BP000006Request	"Input parameter"
// @Success 200  {object} models.BP000006Response  successful operation
// Failure 403 {"code":"error code","message":"error message","response":null}
// @router / [post]
func (c *BP000006Controller) Post() ()  {
}
