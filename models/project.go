package models

import (
	"fmt"
)

type TaskProject struct {
	Name         string `json:"name"`
	Id_project   int    `json:"id_project"`
	Email        string `json:"email"`
	Project_name string `json:"project_name"`
	Deskripsi    string `json:"deskripsi"`
}

type ProjectView struct {
	Status  bool       `json:"status"`
	ResView ProjectRes `json:"resView"`
}
type ProjectRes struct {
	Id_project   string       `json:"id_project"`
	Project_name string       `json:"project_name"`
	Deskripsi    string       `json:"deskripsi"`
	Member       ProjectsView `json:"member"`
}

type ProjectsView struct {
	Id_member string `json:"id_member"`
}

type ProjectTask struct {
	Id_project   string `json:"id_project"`
	Project_name string `json:"project_name"`
	Creator      string `json:"creator"`
	Deskripsi    string `json:"deskripsi"`
}
type CekId struct {
	Max int `json:"max"`
}


//create
func (ExampleModel Models) InsertTbl_Project(Create TaskProject) bool {

	sqlStatement2 := "INSERT INTO tbl_project (name, creator, description, create_date) " +
		"VALUES ($1,$2 ,$3, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Create.Project_name,
		Create.Email,
		Create.Deskripsi,
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}
func (ExampleModel Models) MaxIdProject(Create TaskProject) int {

	//var MaxId int
	//MaxId = Create.Id_project

	Maxx := CekId{}
	sqlStatement3 := " SELECT  max(id) FROM tbl_project "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
	)

	if err3 != nil {
		fmt.Println(err3)
	} else {
		fmt.Println(res3)
	}


	//task := TaskProject{}
	for res3.Next() {
		err := res3.Scan(&Maxx.Max)
		// Exit if we get an error
		if err != nil {
			fmt.Println(err)
		}

	}
	fmt.Println(Maxx)

	return Maxx.Max
}
func (ExampleModel Models) InsertTbl_member_belongto_project(Create TaskProject, MaxId int) bool {
	sqlStatement := "INSERT INTO tbl_member_belongto_project (id_user, id_project, creator, status, create_date) " +
		"VALUES ($1,$2 ,$3, $4, now()::timestamp)"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Create.Email,
		MaxId,
		Create.Email,
		true,
	)
	fmt.Println(MaxId)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(res)
		return true
	}

}

//edit
func (ExampleModel Models) CekId_Project(Edit TaskProject) bool {
	//cek id_project
	sqlStatement3 := "SELECT tbl_project.creator FROM tbl_project " +
		"WHERE tbl_project.id =$1"
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		Edit.Id_project,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	task := ProjectTask{}
	for res3.Next() {
		err3 := res3.Scan(&task.Id_project)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)
		}
	}

	if task.Id_project != task.Id_project {
		return false
	}
	return true
}
func (ExampleModel Models) EditTbl_Project(Edit TaskProject) bool {

	sqlStatement := " UPDATE tbl_project " +
		"SET  name = $3, description = $4 " +
		"WHERE tbl_project.id =$1 AND tbl_project.creator =$2"
	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
		Edit.Id_project,
		Edit.Email,
		Edit.Project_name,
		Edit.Deskripsi,
	)
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		fmt.Println(res)
		return true
	}

}

//view
//func (ExampleModel Models) ViewProject(View TaskProject) ProjectView {
//	getView := ProjectView{}
//
//	sqlStatement3 := "SELECT tbl_project.name FROM tbl_project " +
//		"WHERE tbl_project.id_project =$1 "
//	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
//		View.Id_project,
//	)
//	if err3 != nil {
//		fmt.Println(err3)
//
//	} else {
//		fmt.Println(res3)
//	}
//	task := ProjectTask{}
//	for res3.Next() {
//		err3 := res3.Scan(&task.Id_project, &task.Project_name, &task.Deskripsi)
//		// Exit if we get an error
//		if err3 != nil {
//			fmt.Println(err3)
//
//		}
//	}
//	if task.Id_project != "" {
//		getView.Status = false
//		return getView
//	}
//
//	sqlStatement := "SELECT * FROM tbl_member_belongto_project " +
//		"INNER JOIN tbl_project ON tbl_member_belongto_project.id_project = tbl_project.id "+
//		"WHERE tbl_member_belongto_project.id =$1 "
//	res, err := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement,
//		View.Id_project,
//	)
//	if err3 != nil {
//		fmt.Println(err)
//
//	} else {
//		fmt.Println(res)
//	}
//
//	ViewResponse := ProjectView{}
//	for res.Next() {
//		projects := ProjectRes{}
//		err2 := res.Scan(&projects.Member.Id_member)
//		// Exit if we get an error
//		if err2 != nil {
//			fmt.Println(err2)
//		}
//		projects.Id_project = task.Id_project
//
//		ViewResponse.Projects = append(ViewResponse.Projects, projects)
//	}
//
//	loginResponse.Email = Login.Email
//
//	getlogin.Status = true
//	getlogin.ResLogin = loginResponse
//
//	return getView

//}
