package controllers

import (
	"github.com/astaxie/beego"
)
// Query Micro Loan product full info
type PD000053Controller struct {
	beego.Controller
}

// @Title
// @Description  Query Micro Loan product full info
// @Param	body		body 	models.SPD0000053I	"Input parameter"
// @Success 200  {object} models.SPD0000053O  successful operation
// @Failure 403 body is empty
// @router / [post]
func (c *PD000053Controller) Post() {
}
