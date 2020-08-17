package models

import (
	"fmt"
)

type TimeL struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
	Task       string `json:"task"`
	Id_member  int    `json:"id_member"`
	Weight     int    `json:"weight"`
	Start_date string `json:"start_date"`
	Due_date   string `json:"due_date"`
	Note       string `json:"note"`
}

type TimelViewtask struct {
	//Id          int    `json:"id"`
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
	//Task        string `json:"task"`
	//Id_member   int    `json:"id_member"`
	//Weight      int    `json:"weight"`
	//Start_date  string `json:"start_date"`
	//End_date    string `json:"end_date"`
	//Due_date    string `json:"due_date"`
	//Note        string `json:"note"`
	//Creator     int    `json:"creator"`
	//Create_date string `json:"create_date"`
}

type TimelView struct {
	TimelView []TimelViewtask `json:"timel_view"`
}

//create timeline
func (ExampleModel Models) CreateTimel(Create TimeL) bool {

	sqlStatement2 := "INSERT INTO  tbl_project_task(id_project, task, member, weight, start_date, end_date, due_date, note, creator, create_date) " +
		"VALUES ($1,$2 ,$3, $4, $5, $6, $7, $8, $9, now()::timestamp)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Create.Id_project,
		Create.Task,
		Create.Id_member,
		Create.Weight,
		Create.Start_date,
		Create.Due_date,
		Create.Due_date,
		Create.Note,
		Create.Id_member,
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

//update timeline
func (ExampleModel Models) EditTimel(Edit TimeL) bool {

	sqlStatement2 := "UPDATE tbl_project_task " +
		"SET creator = $1, task = $2, member = $3, note = $4, start_date = $5, end_date = $6, due_date = $7, weight = $8 " +
		"WHERE tbl_project_task.id = $9  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Id_member,
		Edit.Task,
		Edit.Id_member,
		Edit.Note,
		Edit.Start_date,
		Edit.Due_date,
		Edit.Due_date,
		Edit.Weight,
		Edit.Id,
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

//view timeline
func (ExampleModel Models) ViewTimel(View TimelViewtask) TimelView {

	sqlStatement3 := "SELECT tbl_project_task.id_project, tbl_member_belongto_project.id_user FROM tbl_project_task " +
		"INNER JOIN tbl_member_belongto_project ON tbl_project_task.member = tbl_member_belongto_project.id " +
		"WHERE tbl_project_task.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	defer ExampleModel.db.GetDatabaseConfig().Close()
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := TimelView{}

	for res3.Next() {
		task := TimelViewtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_project, &task.Task, &task.Id_member, &task.Weight, &task.Start_date, &task.End_date, &task.Due_date,
		//&task.Note, &task.Creator, &task.Create_date, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.TimelView = append(result.TimelView, task)
	}

	return result

}
