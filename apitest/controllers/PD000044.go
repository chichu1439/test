package controllers

import (
	"github.com/astaxie/beego"
)
// Query Loan interest strategy
type PD000044Controller struct {
	beego.Controller
}

// @Title
// @Description Query Loan interest strategy
// @Param	body		body 	models.SPD0000044I	"Input parameter"
// @Success 200  {object} models.SPD0000044O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000044Controller) Post() {
}
