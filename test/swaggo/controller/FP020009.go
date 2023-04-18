package controller

import "github.com/gin-gonic/gin"

// FP020009 godoc
// @Summary     execl list
// @Description execl list
// @Tags        paymentSwitch
// @Param       FP020009Request body     models.RequestSfp0200009 true "execl list"
// @Success     200                        {object} controller.Response{response=models.ResponseSfp0200009}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/import/execl/list [post]
func (c *Controller) FP020009(ctx *gin.Context) {

}
