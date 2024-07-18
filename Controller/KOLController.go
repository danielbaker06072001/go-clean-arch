package Controller

import (
	"net/http"
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Service/Interface"
	"wan-api-verify-user/ViewModel"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type KOLController struct {
	KOLService Interface.IKOLService
}

func NewKOLController(echoCtx *echo.Echo, KOLServiceObject Interface.IKOLService) {

	KOLControllerObject := &KOLController{
		KOLService: KOLServiceObject,
	}

	echoCtx.GET("/healthz", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	})

	VerifyUserGroup := echoCtx.Group("/verify-user")
	{
		VerifyUserGroup.POST("/kol", KOLControllerObject.UpdateKol)
		VerifyUserGroup.POST("/client", KOLControllerObject.UpdateClient)
	}
}

func (KolController *KOLController) UpdateKol(echoCtx echo.Context) error {
	var KolVM ViewModel.UpdateKolViewModel

	Guid := uuid.New().String()

	var input DTO.KOLInputDTO
	if err := echoCtx.Bind(&input); err != nil {
		KolVM.CommonUpdateResponse.Result = "Failed"
		KolVM.CommonUpdateResponse.Guid = Guid
		echoCtx.JSON(http.StatusBadRequest, KolVM)
		return nil
	}

	var params DTO.AddedParam = make(DTO.AddedParam)
	params["KolID"] = input.KolID
	params["VerificationStatus"] = *input.VerificationStatus
	if input.ImageUrl != nil {
		for _, p := range *input.ImageUrl {
			params[p.Key] = p.Value
		}
	}

	KolDto, err := KolController.KOLService.UpdateKol(params)
	if err != nil {
		KolVM.CommonUpdateResponse.Result = "Failed"
		KolVM.CommonUpdateResponse.ErrorMessage = err.Error()
		KolVM.CommonUpdateResponse.Guid = Guid
		echoCtx.JSON(http.StatusBadRequest, KolVM)
		return nil
	}

	// * Update successfully
	KolVM.CommonUpdateResponse.Result = "Success"
	KolVM.CommonUpdateResponse.Guid = Guid
	KolVM.Kol = KolDto
	echoCtx.JSON(http.StatusOK, KolVM)
	return nil
}

func (KolController *KOLController) UpdateClient(echoCtx echo.Context) error {
	return nil
}
