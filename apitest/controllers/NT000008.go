package controllers

import (
	"github.com/astaxie/beego"
)

// Query Notification Strategy Detail
type NT000008Controller struct {
	beego.Controller
}

// @Title
// @Description Query Notification Strategy Detail
// @Param	body		body 	models.NT000008I	"Input parameter"
// @Success 200  {object} models.NT000008O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000008Controller) Post() ()  {
}
