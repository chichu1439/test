package controller

import "github.com/gin-gonic/gin"

// FP000104 godoc
// @Summary     netting param update
// @Description netting param update
// @Tags        paymentSwitch
// @Param       FP000104Request body     models.FP000104I true "netting param update"
// @Success     200                        {object} controller.Response{response=models.FP000104O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/netting/param/update [post]
func (c *Controller) FP000104(ctx *gin.Context) {

}
