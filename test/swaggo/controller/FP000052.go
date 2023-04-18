package controller

import "github.com/gin-gonic/gin"

// FP000052 godoc
// @Summary     step detail
// @Description step detail
// @Tags        paymentSwitch
// @Param       FP000052Request body     models.FP000052I true "step detail"
// @Success     200                        {object} controller.Response{response=models.FP000052O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/step/detail [post]
func (c *Controller) FP000052(ctx *gin.Context) {

}
