package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/osamamosaad/paytabs/entities"
	accountLibrary "github.com/osamamosaad/paytabs/pkgs/account-library"
	transactionLibrary "github.com/osamamosaad/paytabs/pkgs/transaction-library"
	"github.com/osamamosaad/paytabs/repositories"
	DBStorage "github.com/osamamosaad/paytabs/storage"
	"github.com/osamamosaad/paytabs/utils"
)

func Transaction(storage *DBStorage.Storage) func(echo.Context) error {
	return func(c echo.Context) error {
		var transaction entities.Transaction
		err := c.Bind(&transaction)
		if err != nil {
			return utils.WriteError(c, http.StatusBadRequest, err.Error())
		}

		transLib := transactionLibrary.New(
			repositories.NewTransaction(storage),
			accountLibrary.New(repositories.NewAccount(storage)),
		)

		result, err := transLib.Transfer(transaction)
		if err != nil {
			return utils.WriteError(c, http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, result)
	}
}
