package controller

import "github.com/gin-gonic/gin"

// FP020007 godoc
// @Summary     execl add
// @Description execl add
// @Tags        paymentSwitch
// @Param       FP020007Request body     models.FP020007I true "execl add"
// @Success     200                        {object} controller.Response{response=string}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/import/execl/add [post]
func (c *Controller) FP020007(ctx *gin.Context) {

}
