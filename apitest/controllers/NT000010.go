package controllers

import (
	"github.com/astaxie/beego"
)

// Delete Notification Strategy
type NT000010Controller struct {
	beego.Controller
}

// @Title
// @Description Delete Notification Strategy
// @Param	body		body 	models.NT000010I	"Input parameter"
// @Success 200  {object} models.NT000010O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000010Controller) Post() ()  {
}
