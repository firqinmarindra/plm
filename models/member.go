package models

import "fmt"

type Membertask struct {
	Email      string `json:"email"`
	Id_member  int    `json:"id_member"`
	Role       string `json:"role"`
	Id_project int    `json:"id_project"`
}

type MemberViewtask struct {
	Id          int    `json:"id"`
	Id_user     string `json:"id_user"`
	Id_project  int    `json:"id_project"`
	Creator     string `json:"creator"`
	Status      string `json:"status"`
	Create_date string `json:"create_date"`
}

type MemberView struct {
	MemberView []MemberViewtask `json:"MemberView"`
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
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//view member
func (ExampleModel Models) ViewMember(View MemberViewtask) MemberView {

	sqlStatement3 := "SELECT *  FROM tbl_member_belongto_project " +
		"WHERE tbl_member_belongto_project.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := MemberView{}

	for res3.Next() {
		task := MemberViewtask{}
		err3 := res3.Scan(&task.Id, &task.Id_user, &task.Id_project, &task.Creator, &task.Status, &task.Create_date)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.MemberView = append(result.MemberView, task)
	}

	return result

}

//delete member
func (ExampleModel Models) DelMember(Id_member int) bool {

	sqlStatement2 := "DELETE FROM tbl_member_belongto_project " +
		"WHERE id = $1"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Id_member,
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}
