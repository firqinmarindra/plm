package models

import "fmt"

type Membertask struct {
	Email      string `json:"email"`
	Id_member  int    `json:"id_member"`
	Role       string `json:"role"`
	Id_project int    `json:"id_project"`
}

type Membertaskview struct {
	Id_project int `json:"id_project"`
}

type MemberView1 struct {
	Status  bool       `json:"status"`
	ResView MemberView `json:"res_view"`
}

type MemberView struct {
	Id_project int              `json:"id_project"`
	Member     []MemberViewtask `json:"member"`
}

type MemberViewtask struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

//add member
func (ExampleModel Models) AddMember(Add Membertask) bool {

	sqlStatement2 := "INSERT INTO  tbl_member_belongto_project(id_user, id_project, creator, status, create_date) " +
		"VALUES ($1,$2 ,$3, $4, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Email,
		Add.Id_project,
		Add.Email,
		true,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//view member
func (ExampleModel Models) GetId_projectMember(View Membertask) MemberView {

	sqlStatement3 := "SELECT tbl_member_belongto_project.id_project FROM tbl_member_belongto_project " +
		"WHERE tbl_member_belongto_project.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	task := MemberView{}
	for res3.Next() {
		err3 := res3.Scan(&task.Id_project)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		//result.MemberView = append(result.MemberView, task)
	}

	return task

}

func (ExampleModel Models) ViewMember(View Membertask, Id_project int) MemberView1 {
	getview := MemberView1{}

	sqlStatement3 := "SELECT tbl_user.name,tbl_user.position  FROM tbl_member_belongto_project " +
		"INNER JOIN tbl_user ON tbl_member_belongto_project.id_user = tbl_user.email " +
		"WHERE tbl_member_belongto_project.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := MemberView{}

	for res3.Next() {
		task := MemberViewtask{}
		err3 := res3.Scan(&task.Name, &task.Position)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.Member = append(result.Member, task)
	}

	//return result

	result.Id_project = View.Id_project

	getview.Status = true
	getview.ResView = result
	return getview

}

//delete member
func (ExampleModel Models) DelMember(Id_member int) bool {

	sqlStatement2 := "DELETE FROM tbl_member_belongto_project " +
		"WHERE id = $1"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Id_member,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}
