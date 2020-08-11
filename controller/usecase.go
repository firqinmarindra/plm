package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type Usecget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
	Link_usec  string `json:"link_usec"`
}

type UsecEditget struct {
	Email     string `json:"email"`
	Id        int    `json:"id"`
	Link_usec string `json:"link_usec"`
}

type UsecDescGet struct {
	Email       string `json:"email"`
	Id_usecase  int    `json:"id_usecase"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Index       int    `json:"index"`
}

type UsecDescEditGet struct {
	Email       string `json:"email"`
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Index       int    `json:"index"`
}

type DelUsecDescGet struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
}

type ViewUsecget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

type ViewUsecErdget struct {
	Email  string `json:"email"`
	Id_erd int    `json:"id_erd"`
}

type ViewUsecScenarget struct {
	Email               string `json:"email"`
	Id_usecase_scenario int    `json:"id_usecase_scenario"`
}

type ViewUsecDflowget struct {
	Email       string `json:"email"`
	Id_dataflow int    `json:"id_dataflow"`
}

//add usecase
func (Controller Controller) AddUsec(c echo.Context) error {
	a := new(Usecget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Usecasetask{
		Email:      a.Email,
		Id_project: a.Id_project,
		Link_usec:  a.Link_usec,
	}
	fmt.Println(a.Email)

	addUsec := Controller.ma.AddUsec(ab)

	if addUsec {

		getadd := Usecget{a.Email, a.Id_project, a.Link_usec}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create usecase",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create usecase",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit usecase
func (Controller Controller) EditUsec(c echo.Context) error {
	a := new(UsecEditget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Usecasetask{
		Email:     a.Email,
		Id:        a.Id,
		Link_usec: a.Link_usec,
	}
	fmt.Println(a.Email)

	editUsec := Controller.ma.EditUsec(ab)

	if editUsec {

		getedit := UsecEditget{a.Email, a.Id, a.Link_usec}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update usecase",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update usecase",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//add usecase description
func (Controller Controller) AddUsecDesc(c echo.Context) error {
	a := new(UsecDescGet)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecDescTask{
		Email:       a.Email,
		Id_usecase:  a.Id_usecase,
		Name:        a.Name,
		Description: a.Description,
		Index:       a.Index,
	}
	fmt.Println(a.Email)

	addUsecDesc := Controller.ma.AddUsecDesc(ab)

	if addUsecDesc {

		getadd := UsecDescGet{a.Email, a.Id_usecase, a.Name, a.Description, a.Index}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create usecase description",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create usecase description",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//edit usecase description
func (Controller Controller) EditUsecDesc(c echo.Context) error {
	a := new(UsecDescEditGet)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecDescTask{
		Email:       a.Email,
		Id:          a.Id,
		Name:        a.Name,
		Description: a.Description,
		Index:       a.Index,
	}
	fmt.Println(a.Email)

	editUsecDesc := Controller.ma.EditUsecDesc(ab)

	if editUsecDesc {

		getedit := UsecDescEditGet{a.Email, a.Id, a.Name, a.Description, a.Index}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil update usecase desc",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal update usecase desc",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//delete usecase description
func (Controller Controller) DelUsecDesc(c echo.Context) error {
	a := new(DelUsecDescGet)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Usecasetask{
		Email: a.Email,
		Id:    a.Id,
	}
	fmt.Println(a.Email)

	editUsecDesc := Controller.ma.DelUsecDesc(ab.Id)

	if editUsecDesc {

		getedit := DelUsecDescGet{a.Email, a.Id}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil delete usecase desc",
			Data:    getedit,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal delete usecase desc",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//view usecase
func (Controller Controller) ViewUsec(c echo.Context) error {
	a := new(ViewUsecget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecViewtask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	view := Controller.ma.ViewUsec(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data Usecase",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view usecase erd
func (Controller Controller) ViewUsecErd(c echo.Context) error {
	a := new(ViewUsecErdget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecViewErdtask{
		Email:  a.Email,
		Id_erd: a.Id_erd,
	}
	view := Controller.ma.ViewUsecErd(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data Usecase by erd",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view usecase scenario
func (Controller Controller) ViewUsecScenar(c echo.Context) error {
	a := new(ViewUsecScenarget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecViewScenartask{
		Email:               a.Email,
		Id_usecase_scenario: a.Id_usecase_scenario,
	}
	view := Controller.ma.ViewUsecScenar(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data Usecase by Scenario",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}

//view usecase dataflow
func (Controller Controller) ViewUsecDataflow(c echo.Context) error {
	a := new(ViewUsecDflowget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.UsecViewDflowtask{
		Email:       a.Email,
		Id_dataflow: a.Id_dataflow,
	}
	view := Controller.ma.ViewUsecDataflow(ab)

	res := responsegenr.ResponseGenericGet{
		Status:  "Success",
		Message: "Berhasil dapatkan data Usecase by dataflow",
		Data:    view,
	}
	return c.JSON(http.StatusOK, res)

}
