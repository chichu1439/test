package controller

import "github.com/gin-gonic/gin"

// FP000106 godoc
// @Summary     adjustment update
// @Description adjustment update
// @Tags        paymentSwitch
// @Param       FP000106Request body     models.FP000106I true "adjustment update"
// @Success     200                        {object} controller.Response{response=models.FP000106O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/balance/adjustment/update [post]
func (c *Controller) FP000106(ctx *gin.Context) {

}
