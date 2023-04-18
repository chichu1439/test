package controller

import "github.com/gin-gonic/gin"

// FP000063 godoc
// @Summary     apis chema
// @Description apis chema
// @Tags        paymentSwitch
// @Param       FP000063Request body     models.FP000063I true "apis chema"
// @Success     200                        {object} controller.Response{response=models.FP000063O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/api/schema [post]
func (c *Controller) FP000063(ctx *gin.Context) {

}
