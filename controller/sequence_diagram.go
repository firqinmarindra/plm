package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type SeqDiagget struct {
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Sequence_link         string `json:"sequence_link"`
	Description           int    `json:"description"`
}

type SeqDiagEditget struct {
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
	Id                    int    `json:"id"`
	Sequence_link         string `json:"sequence_link"`
	Description           int    `json:"description"`
	Status                string `json:"status"`
}

type ViewSeqDiagget struct {
	Email                 string `json:"email"`
	Id_dataflow_structure int    `json:"id_dataflow_structure"`
}

type ViewSeqprojget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

//add sequence diagram
func (Controller Controller) AddSeqDiag(c echo.Context) error {
	a := new(SeqDiagget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.SeqDiagtask{
		Email:                 a.Email,
		Id_dataflow_structure: a.Id_dataflow_structure,
		Sequence_link:         a.Sequence_link,
		Description:           a.Description,
	}
	fmt.Println(a.Email)

	addSeqDiag := Controller.ma.AddSeqDiags(ab)

	if addSeqDiag {

		getadd := SeqDiagget{a.Email, a.Id_dataflow_structure, a.Sequence_link, a.Description}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create sequence diagram",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create sequence diagram",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit sequence diagram
func (Controller Controller) EditSeqDiag(c echo.Context) error {
	a := new(SeqDiagEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.SeqDiagtask{
		Email:                 a.Email,
		Id_dataflow_structure: a.Id_dataflow_structure,
		Id:                    a.Id,
		Sequence_link:         a.Sequence_link,
		Description:           a.Description,
		Status:                a.Status,
	}
	fmt.Println(a.Email)

	editSeqDiag := Controller.ma.EditSeqDiags(ab)

	if editSeqDiag {

		getadd := SeqDiagEditget{a.Email, a.Id_dataflow_structure, a.Id, a.Sequence_link, a.Description, a.Status}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update sequence diagram",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update sequence diagram",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view sequence diagram
func (Controller Controller) ViewSeqDiag(c echo.Context) error {
	a := new(ViewSeqDiagget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ViewSeqDiagtask{
		Email:                 a.Email,
		Id_dataflow_structure: a.Id_dataflow_structure,
	}
	view := Controller.ma.ViewSeqDiag(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data sequence diagram",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view sequence diagram by project
func (Controller Controller) ViewSeqDiagProj(c echo.Context) error {
	a := new(ViewSeqprojget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.ViewSeqDiagProjtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewSeqDiagProj(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data sequence diagram bye project",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}
