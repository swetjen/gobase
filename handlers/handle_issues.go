package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"monkey/crud"
	"monkey/models"
	"net/http"
	"time"
)

func IssuesGetAll(c echo.Context) error {
	issues, err := crud.GetAllIssues()
	if err != nil {
		return err
	}
	return c.JSONPretty(http.StatusOK, issues, " ")
}

func IssueGetOneByCode(c echo.Context) error {
	code := c.Param("code")
	fmt.Println(code)
	issue, err := crud.GetOneIssueByCode(code)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Issue not found.")
	}
	return c.JSONPretty(http.StatusOK, issue, " ")
}

func IssueGetOneById(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)
	issue, err := crud.GetOneIssueById(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Issue with ID "+id+" not found.")
	}
	return c.JSONPretty(http.StatusOK, issue, " ")
}

func IssueCreate(c echo.Context) (err error) {
	type IssueRequest struct {
		Title     string `json:"title"`
		Code      string `json:"code"`
		Completed bool   `json:"false"`
	}
	i := new(IssueRequest)
	// Bind to type
	if err = c.Bind(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// Validate

	// Do something
	newIssue := models.Issue{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Title:     i.Title,
		Code:      i.Code,
		Completed: i.Completed,
	}
	crud.CreateOne(newIssue)
	return c.JSON(http.StatusOK, newIssue)
}
