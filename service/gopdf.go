package service

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/signintech/gopdf"
)

func CreateResume(c echo.Context) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "../ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
	}
	pdf.Cell(nil, "您好")
	pdf.WritePdf("hello.pdf")
	return c.JSON(http.StatusOK, pdf.Close().Error())
}
