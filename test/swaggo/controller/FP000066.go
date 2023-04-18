package controller

import "github.com/gin-gonic/gin"

// FP000066 godoc
// @Summary     host start
// @Description host start
// @Tags        paymentSwitch
// @Param       FP000066Request body     models.FP000066I true "host start"
// @Success     200                        {object} controller.Response{response=models.FP000066O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/host/start [post]
func (c *Controller) FP000066(ctx *gin.Context) {

}
