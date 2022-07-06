package router

import (
	"echo-aws-sdk-playground/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func GetPreSignedUrlForUpload(c echo.Context) error {
	param := c.QueryParam("key")
	log.Info("requested with + ", param)
	upload, err := service.GetPreSignedUrlForUpload(param)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, upload)
}

func Bar(c echo.Context) error {
	//jsonBody := make(map[string]interface{})
	fmt.Println(c.Request().Body)
	abc := make([]byte, 1024*10)
	read, err := c.Request().Body.Read(abc)
	if err != nil {
		//return err
	}
	fmt.Println(read)
	fmt.Println(string(abc))
	//err = json.NewDecoder(c.Request().Body).Decode(&jsonBody)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil
	//}
	//fmt.Println(jsonBody)
	return c.JSON(http.StatusOK, string(abc))
}
