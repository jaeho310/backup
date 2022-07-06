package main

import (
	"echo-aws-sdk-playground/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.GET("/s3/presigned-url/upload", router.GetPreSignedUrlForUpload)
	e.POST("/foo", router.Bar)
	e.Logger.Fatal(e.Start(":8395"))
}
