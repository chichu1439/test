package controller

import "github.com/gin-gonic/gin"

// FP000113 godoc
// @Summary     payment creat list
// @Description payment creat list
// @Tags        paymentSwitch
// @Param       FP000113Request body     models.FP000113I true "payment creat list"
// @Success     200                        {object} controller.Response{response=models.FP000113O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/payment/creat/list [post]
func (c *Controller) FP000113(ctx *gin.Context) {

}
