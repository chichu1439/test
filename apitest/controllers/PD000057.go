package controllers

import (
	"github.com/astaxie/beego"
)
// Delete Loan interest strategy
type PD000057Controller struct {
	beego.Controller
}

// @Title
// @Description  Delete Loan interest strategy
// @Param	body		body 	models.SPD0000057I	"Input parameter"
// @Success 200  {object} models.SPD0000057O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000057Controller) Post() {
}
