package controller

import "github.com/gin-gonic/gin"

// FP000107 godoc
// @Summary     adjustment delete
// @Description adjustment delete
// @Tags        paymentSwitch
// @Param       FP000107Request body     models.FP000107I true "adjustment delete"
// @Success     200                        {object} controller.Response{response=models.FP000107O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/delete [post]
func (c *Controller) FP000107(ctx *gin.Context) {

}
