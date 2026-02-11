package handler

import (
	"mampu-demo/model"
	"mampu-demo/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type WalletHandler struct {
	svc *service.WalletService
}

func NewWalletHandler(svc *service.WalletService) *WalletHandler {
	return &WalletHandler{svc: svc}
}

func (h *WalletHandler) GetBalance(c echo.Context) error {
	userID := c.Param("userId")

	wallet, err := h.svc.GetBalance(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"massage": "User not found",
		})
	}
	return c.JSON(http.StatusOK, wallet)
}

func (h *WalletHandler) Withdraw(c echo.Context) error {
	req := new(model.WithdrawRequest)

	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format inputan harus berupa json",
		})
	}

	err := h.svc.Withdraw(req.UserID, req.Amount)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Tarik tunai berhasil",
	})
}
