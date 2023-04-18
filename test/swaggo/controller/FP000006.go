package controller

import "github.com/gin-gonic/gin"

// FP000006 godoc
// @Summary     participant creation
// @Description participant creation
// @Tags        paymentSwitch
// @Param       FP000006Request body     models.FP000006I true "participant creation"
// @Success     200                        {object} controller.Response{response=models.FP000006O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant/creation [post]
func (c *Controller) FP000006(ctx *gin.Context) {

}
