package service

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/signintech/gopdf"
	"github.com/wmcff/serve-data/model/dto"
)

func CreateResume(resumeDto *dto.ResumeDto, c echo.Context) error {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("wts11", "ttf/wts11.ttf")
	if err != nil {
		log.Print(err.Error())
	}

	err = pdf.SetFont("wts11", "", 14)
	if err != nil {
		log.Print(err.Error())
	}
	pdf.Cell(nil, "姓名: "+resumeDto.Person.Name+"  ")
	pdf.Br(20)
	pdf.Cell(nil, "性别: "+resumeDto.Person.Sex+"  ")
	pdf.Br(20)
	pdf.Cell(nil, "年龄: "+resumeDto.Person.Age+"  ")
	pdf.WritePdf("files/hello.pdf")
	var res interface{}
	return c.JSON(http.StatusOK, res)
}
