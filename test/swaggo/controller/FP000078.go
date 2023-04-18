package controller

import "github.com/gin-gonic/gin"

// FP000078 godoc
// @Summary     alert balance set
// @Description alert balance set
// @Tags        paymentSwitch
// @Param       FP000078Request body     models.FP000073I true "alert balance set"
// @Success     200                        {object} controller.Response{response=models.FP000073I}
// @Failure     500                {object} controller.Error
// @Router     /v2/b2b/nfps/alert-balance/setting [post]
func (c *Controller) FP000078(ctx *gin.Context) {

}
