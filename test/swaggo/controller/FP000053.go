package controller

import "github.com/gin-gonic/gin"

// FP000053 godoc
// @Summary     api query
// @Description api query
// @Tags        paymentSwitch
// @Param       FP000053Request body     models.FP000053I true "api query"
// @Success     200                        {object} controller.Response{response=models.FP000053O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/test/api/query [post]
func (c *Controller) FP000053(ctx *gin.Context) {

}
