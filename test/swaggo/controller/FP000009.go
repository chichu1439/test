package controller

import "github.com/gin-gonic/gin"

// FP000009 godoc
// @Summary     participant agency relation creation
// @Description participant agency relation creation
// @Tags        paymentSwitch
// @Param       FP000009Request body     models.FP000009I true "participant agency relation creation"
// @Success     200                        {object} controller.Response{response=models.FP000009O}
// @Failure     500                {object} controller.Error
// @Router     /v1/nfps/participant-agency-relation/creation [post]
func (c *Controller) FP000009(ctx *gin.Context) {

}
