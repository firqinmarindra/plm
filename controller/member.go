package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type Memberget struct {
	Email      string `json:"email"`
	Id_member  int    `json:"id_member"`
	Role       string `json:"role"`
	Id_project int    `json:"id_project"`
}
type DelMemberget struct {
	Email      string `json:"email"`
	Id_member  int    `json:"id_member"`
	Id_project int    `json:"id_project"`
}
type ViewMemberget struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
}

//add member
func (Controller Controller) AddMember(c echo.Context) error {
	a := new(Memberget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Membertask{
		Email:      a.Email,
		Id_member:  a.Id_member,
		Role:       a.Role,
		Id_project: a.Id_project,
	}
	fmt.Println(a.Email)

	addMember := Controller.ma.AddMember(ab)

	if addMember {

		getadd := Memberget{a.Email, a.Id_member, a.Role, a.Id_project}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil create member",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal create member",
		}
		return c.JSON(http.StatusOK, res)

	}
}

//veiw member
func (Controller Controller) ViewMember(c echo.Context) error {
	a := new(ViewMemberget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Membertask{
		Email:      a.Email,
		Id_project: a.Id_project,
	}
	//cekLogin := Controller.ma.Ceklogin(ab.Email, ab.Password)
	id_project := Controller.ma.GetId_projectMember(ab)
	view := Controller.ma.ViewMember(ab, id_project.Id_project)

	if view.Status {
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "view berhasil",
			Data:    view.ResView,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGenericGet{
			Status:  "Error",
			Message: "Login Gagal ",
		}
		return c.JSON(http.StatusOK, res)
	}

}

//delete member
func (Controller Controller) DelMember(c echo.Context) error {
	a := new(DelMemberget)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.Membertask{
		Email:      a.Email,
		Id_member:  a.Id_member,
		Id_project: a.Id_project,
	}
	fmt.Println(a.Email)

	delMember := Controller.ma.DelMember(ab.Id_member)

	if delMember {

		getadd := DelMemberget{a.Email, a.Id_member, a.Id_project}
		res := responsegenr.ResponseGenericGet{
			Status:  "Success",
			Message: "Berhasil hapus member",
			Data:    getadd,
		}
		return c.JSON(http.StatusOK, res)

	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Gagal hapus member",
		}
		return c.JSON(http.StatusOK, res)

	}
}
