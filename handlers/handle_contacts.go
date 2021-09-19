package handlers

import (
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"monkey/crud"
	"monkey/models"
	"net/http"
	"time"
)

func ContactsCreate(c echo.Context) (err error) {
	type LeadRequest struct {
		Message string `json:"message"`
		Email   string `json:"email"`
	}
	i := new(LeadRequest)

	if err = c.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate

	// Save
	newContact := models.Contact{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Message:   i.Message,
		Email:     i.Email,
	}

	crud.ContactsCreateOne(newContact)
	return c.JSON(http.StatusOK, newContact)

}

func ContactsGetAll(c echo.Context) error {
	contacts, err := crud.ContactsGetAll()
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, contacts, " ")

}
