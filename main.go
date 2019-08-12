package main

import (
	"csv/pkg/converter"
	"csv/pkg/onboarding"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	lines, err := ioutil.ReadFile("./B2B_GWG_UID.csv")
	if err != nil {
		panic(err)
	}
	encoded := base64.StdEncoding.EncodeToString(lines)
	decoded, _ := base64.StdEncoding.DecodeString(encoded)
	fmt.Println(string(decoded))

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":8080"))
}

type UploadRequest struct {
	Documentbase64 string `json:"documentBase64"`
}

func upload(c echo.Context) error {
	var request UploadRequest
	if err := c.Bind(&request); err != nil {
		panic(err)
	}

	converter := &converter.CSVConverter{}
	companies, err := converter.Convert(request.Documentbase64)
	if err != nil {
		log.Println("Failed to convert csv: ", err)

		return c.JSON(http.StatusInternalServerError, "InternalServerError")
	}

	log.Println(companies)

	for i := 0; i < len(companies); i++ {
		onboarding.OnboardRequest(companies[i].UID)

	}
	return c.JSON(http.StatusOK, companies)
}
