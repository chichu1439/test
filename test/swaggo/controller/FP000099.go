package controller

import "github.com/gin-gonic/gin"

// FP000099 godoc
// @Summary     get messages info
// @Description get messages info
// @Tags        paymentSwitch
// @Param       FP000099Request body     models.FP000099I true "get messages info"
// @Success     200                        {object} controller.Response{response=models.FP000099O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/simulato1/get/messages/info [post]
func (c *Controller) FP000099(ctx *gin.Context) {

}
