package controller

import "github.com/gin-gonic/gin"

// FP000097 godoc
// @Summary     simulato3set message
// @Description simulato3set message
// @Tags        paymentSwitch
// @Param       FP000097Request body     models.FP000097I true "simulato3set message"
// @Success     200                        {object} controller.Response{response=models.FP000097O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/scenario/simulato3/set/message [post]
func (c *Controller) FP000097(ctx *gin.Context) {

}
