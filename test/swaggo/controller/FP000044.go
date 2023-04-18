package controller

import "github.com/gin-gonic/gin"

// FP000044 godoc
// @Summary     participant list
// @Description participant list
// @Tags        paymentSwitch
// @Param       FP000044Request body     models.FP000044I true "participant list"
// @Success     200                        {object} controller.Response{response=models.FP000044O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/alert/participant/list [post]
func (c *Controller) FP000044(ctx *gin.Context) {

}
