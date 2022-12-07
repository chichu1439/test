package controllers

import (
	"github.com/astaxie/beego"
)

// Query Notification Strategy List
type NT000007Controller struct {
	beego.Controller
}

// @Title
// @Description Query Notification Strategy List
// @Param	body		body 	models.NT000007I	"Input parameter"
// @Success 200  {object} models.NT000007O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *NT000007Controller) Post() ()  {
}
