package controllers

import (
	"github.com/astaxie/beego"
)

// Receive online message and send application
type NT000013Controller struct {
	beego.Controller
}

// @Title
// @Description Receive online message and send application
// @Param	body		body 	models.NT000013I	"Input parameter"
// @Success 200  {object} models.NT000013O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000013Controller) Post() ()  {
}
