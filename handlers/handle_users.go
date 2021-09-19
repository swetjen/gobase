package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"monkey/crud"
	"net/http"
)

func UsersGetAll(c echo.Context) error {
	// Check Authentication...
	users, err := crud.GetAllUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func LoginAccessToken(c echo.Context) error {
	type User struct {
		Email    string `json:"email" bson:"email"`
		Password string `json:"password,omitempty"`
		Token    string `json:"token,omitempty"`
	}
	u := new(User)
	err := c.Bind(u)
	fmt.Println(u)
	if err != nil {
		log.Println(err)
		return err
	}

	if u.Email != "swetjen@gmail.com" || u.Password != "bummer" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Invalid username or password."}
	}

	return nil

	// create token
	//token := jwt.New(jwt.SigningMethodHS256)
	//
	//// set claims
	//claims := token.Claims.(jwt.MapClaims)
	//claims["id"] = u.Email
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	//
	//// generate encoded token and send it in response
	//u.Token, err = token.SignedString([]byte("secret"))
	//if err != nil {
	//	panic(err)
	//}
	//u.Password = ""
	//return c.JSON(http.StatusOK, u)
}
