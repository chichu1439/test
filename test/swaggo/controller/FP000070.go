package controller

import "github.com/gin-gonic/gin"

// FP000070 godoc
// @Summary     partclearing query
// @Description partclearing query
// @Tags        paymentSwitch
// @Param       FP000070Request body     models.FP000070I true "partclearing query"
// @Success     200                        {object} controller.Response{response=models.FP000070O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/partclearing/query [post]
func (c *Controller) FP000070(ctx *gin.Context) {

}
