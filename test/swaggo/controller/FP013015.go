package controller

import "github.com/gin-gonic/gin"

// FP013015 godoc
// @Summary     20022 account verify v3
// @Description 20022 account verify v3
// @Tags        paymentSwitch
// @Param       FP013015Request body     models.FpsAcmt023 true "20022 account verify v3"
// @Success     200                        {object} controller.Response{response=models.FpsAcmt024}
// @Failure     500                {object} controller.Error
// @Router     /v3/nfps/iso20022/lookup [post]
func (c *Controller) FP013015(ctx *gin.Context) {

}
