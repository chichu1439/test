package controllers

import (
	"github.com/astaxie/beego"
)

// Update Notification Strategy
type NT000009Controller struct {
	beego.Controller
}

// @Title
// @Description Update Notification Strategy
// @Param	body		body 	models.NT000009I	"Input parameter"
// @Success 200  {object} models.NT000009O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000009Controller) Post() ()  {
}
