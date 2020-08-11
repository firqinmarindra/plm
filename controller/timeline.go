package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type CreteTimelget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
	Task       string `json:"task"`
	Id_member  int    `json:"id_member"`
	Weight     int    `json:"weight"`
	Start_date string `json:"start_date"`
	Due_date   string `json:"due_date"`
	Note       string `json:"note"`
}

type UpdateTimelget struct {
	Email      string `json:"email"`
	Id         int    `json:"id"`
	Task       string `json:"task"`
	Id_member  int    `json:"id_member"`
	Weight     int    `json:"weight"`
	Start_date string `json:"start_date"`
	Due_date   string `json:"due_date"`
	Note       string `json:"note"`
}

type ViewTimelget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

//create timeline
func (Controller Controller) CreateTimel(c echo.Context) error {
	a := new(CreteTimelget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TimeL{
		Email:      a.Email,
		Id_project: a.Id_project,
		Task:       a.Task,
		Id_member:  a.Id_member,
		Weight:     a.Weight,
		Start_date: a.Start_date,
		Due_date:   a.Due_date,
		Note:       a.Note,
	}
	fmt.Println(a.Email)

	createTimel := Controller.ma.CreateTimel(ab)

	if createTimel {

		getadd := CreteTimelget{a.Email, a.Id_project, a.Task, a.Id_member, a.Weight, a.Start_date, a.Due_date, a.Note}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create timeline",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create timeline",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit timeline
func (Controller Controller) EditTimel(c echo.Context) error {
	a := new(UpdateTimelget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TimeL{
		Email:      a.Email,
		Id:         a.Id,
		Task:       a.Task,
		Id_member:  a.Id_member,
		Weight:     a.Weight,
		Start_date: a.Start_date,
		Due_date:   a.Due_date,
		Note:       a.Note,
	}
	fmt.Println(a.Email)

	editTimel := Controller.ma.EditTimel(ab)

	if editTimel {

		getedit := UpdateTimelget{a.Email, a.Id, a.Task, a.Id_member, a.Weight, a.Start_date, a.Due_date, a.Note}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update timeline",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update timeline",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view member
func (Controller Controller) ViewTimel(c echo.Context) error {
	a := new(ViewTimelget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TimelViewtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewTimel(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data timeline",
		Data:    view,
		//Data:    view.Email,
	}
	return c.JSON(http.StatusOK, res)

}
