package controller

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"plm/models"
	"plm/responsegenr"
)

type TaskProject struct {
	Id           int    `json:"id"`
	Id_project   int    `json:"id_project"`
	Email        string `json:"email"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
}

//type CreateGet struct {
//	Email 		string `json:"email"`
//	Project_name string `json:"project_name"`
//	Deskripsi    string `json:"deskripsi"`
//}

type EditGet struct {
	Email 			string `json:"email"`
	Project_name 	string `json:"project_name"`
	Deskripsi 		string `json:"deskripsi"`
	Id_project 		int `json:"id_project"`
}

type  CreateGets struct {
	Id_project 		int `json:"id_project"`
	Project_name 	string `json:"project_name"`
	Deskripsi 		string `json:"deskripsi"`

}
//create
func (Controller Controller) PostsCreate(c echo.Context) error {
	a := new(TaskProject)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TaskProject{
		Email:        a.Email,
		Project_name: a.Project_name,
		Deskripsi:    a.Deskripsi,

	}

	fmt.Println(a.Email)



	insertTbl_project := Controller.ma.InsertTbl_Project(ab)
	GetMaxId := Controller.ma.MaxIdProject(ab)
	//cekIdProject := Controller.ma.MaxIdProject(ab)
	insertTbl_member_belongto_project := Controller.ma.InsertTbl_member_belongto_project(ab,GetMaxId)

	fmt.Println(GetMaxId)
	//if cekIdProject {

		if insertTbl_project && insertTbl_member_belongto_project {

			getCreate := CreateGets{Id_project: GetMaxId, Project_name: a.Project_name, Deskripsi: a.Deskripsi}
			res := responsegenr.ResponseGenericGet{
				Status:  "Success",
				Message: "Berhasil create project",
				Data:    getCreate,
			}
			return c.JSON(http.StatusOK, res)

		} else {

			res := responsegenr.ResponseGeneric{
				Status:  "Error",
				Message: "Gagal create project",
			}
			return c.JSON(http.StatusOK, res)

		}
	//} else {
	//
	//	res := responsegenr.ResponseGeneric{
	//		Status:  "Error",
	//		Message: "Id Project tidak ditemukan",
	//	}
	//	return c.JSON(http.StatusOK, res)
	//
	//}
}

//edit
func (Controller Controller) EditProject(c echo.Context) error {
	a := new(TaskProject)
	if err := c.Bind(a); err != nil {
		return err
	}

	ab := models.TaskProject{
		Email:        a.Email,
		Project_name: a.Project_name,
		Deskripsi:    a.Deskripsi,
		Id_project: 	a.Id_project,
	}

	cekId_Project := Controller.ma.CekId_Project(ab)
	editTbl_Project := Controller.ma.EditTbl_Project(ab)

	if cekId_Project{

		if editTbl_Project {

			getCreate := EditGet{a.Email, a.Project_name, a.Deskripsi, a.Id_project}
			res := responsegenr.ResponseGenericGet{
				Status:  "Success",
				Message: "Berhasil update project",
				Data:    getCreate,
			}
			return c.JSON(http.StatusOK, res)

		} else {
			res := responsegenr.ResponseGeneric{
				Status:  "Error",
				Message: "Gagal update project",
			}
			return c.JSON(http.StatusOK, res)
		}
	} else {

		res := responsegenr.ResponseGeneric{
			Status:  "Error",
			Message: "Id_project tidak tersedia",
		}
		return c.JSON(http.StatusOK, res)

	}
}
//view
//func (Controller Controller) ViewProject(c echo.Context) error {
//	a := new(TaskProject)
//	if err := c.Bind(a); err != nil {
//		return err
//	}
//
//	ab := models.TaskProject{
//		Email:    a.Email,
//		Id_project: a.Id_project,
//	}
//	posts := Controller.ma.ViewProject(ab)
//
//	if posts.Status {
//		res := responsegenr.ResponseGenericGet{
//			Status:  "Success",
//			Message: "Login berhasil",
//			Data:    posts.ResView,
//		}
//		return c.JSON(http.StatusOK, res)
//	}
//	res := responsegenr.ResponseGenericGet{
//		Status:  "Error",
//		Message: "Login Gagal",
//		Data:    posts,
//	}
//	return c.JSON(http.StatusOK, res)
//}

