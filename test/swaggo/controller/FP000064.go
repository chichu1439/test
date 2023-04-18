package controller

import "github.com/gin-gonic/gin"

// FP000064 godoc
// @Summary     aig serviceID
// @Description aig serviceID
// @Tags        paymentSwitch
// @Param       FP000064Request body     models.FP000064I true "aig serviceID"
// @Success     200                        {object} controller.Response{response=models.FP000064O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/aig/serviceID [post]
func (c *Controller) FP000064(ctx *gin.Context) {

}
