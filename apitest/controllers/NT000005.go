package controllers

import (
	"github.com/astaxie/beego"
)

// Delete Template
type NT000005Controller struct {
	beego.Controller
}

// @Title
// @Description Delete Template
// @Param	body		body 	models.NT000005I	"Input parameter"
// @Success 200  {object} models.NT000005O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000005Controller) Post() ()  {
}
