package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	accountLibrary "github.com/osamamosaad/paytabs/pkgs/account-library"
	"github.com/osamamosaad/paytabs/pkgs/contracts"
	"github.com/osamamosaad/paytabs/repositories"
	"github.com/osamamosaad/paytabs/storage"
	"github.com/osamamosaad/paytabs/utils"
)

func getAccountLib(storage *storage.Storage) contracts.IAccountLibrary {
	return accountLibrary.New(repositories.NewAccount(storage))
}

func ListAccounts(storage *storage.Storage) func(echo.Context) error {
	return func(c echo.Context) error {

		accountLib := getAccountLib(storage)
		results, err := accountLib.FindAll()
		if err != nil {
			return utils.WriteError(c, http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, results)
	}
}

func GetAccount(storage *storage.Storage) func(echo.Context) error {
	return func(c echo.Context) error {

		ID := c.Param("id")
		accountLib := getAccountLib(storage)
		result, err := accountLib.FindById(ID)
		if err != nil {
			return utils.WriteError(c, http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, result)
	}
}
