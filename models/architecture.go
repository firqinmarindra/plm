package models

import "fmt"

type Archtask struct {
	Email        string `json:"email"`
	Id_project   int    `json:"id_project"`
	Link_arch    string `json:"link_arch"`
	Id_arch_diag int    `json:"id_arch_diag"`
	Status       string `json:"status"`
}
type ArchDestask struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	Id_project   int    `json:"id_project"`
	Id_arch_diag int    `json:"id_arch_diag"`
	Description  string `json:"description"`
	Desc_index   string `json:"desc_index"`
}

type ArchViewtask struct {
	Email      string `json:"email"`
	Id_project int    `json:"id_project"`
	//Diagram     string `json:"diagram"`
	//Creator     string `json:"creator"`
	//Create_date string `json:"create_date"`
	//Approval    string `json:"approval"`
}

type ArchView struct {
	//Email     string          `json:"email"`
	ArchView []ArchViewtask `json:"arch_view"`
}

//add arch
func (ExampleModel Models) AddArch(Add Archtask) bool {

	sqlStatement2 := "INSERT INTO  tbl_architecture_diagram (id, diagram, creator, approval) " +
		"VALUES ($1,$2 ,$3, $4)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_project,
		Add.Link_arch,
		136,
		"pending",
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//edit arch
func (ExampleModel Models) EditArch(Edit Archtask) bool {

	sqlStatement2 := "UPDATE tbl_architecture_diagram " +
		"SET diagram = $1, approval = $2 " +
		"WHERE id = $3 AND creator = $4 "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Link_arch,
		Edit.Status,
		Edit.Id_arch_diag,
		8,
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//add des arch
func (ExampleModel Models) AddArchDesc(Add ArchDestask) bool {

	sqlStatement2 := "INSERT INTO  tbl_architecture_diagram_dec (id_architecture_diagram, index, description) " +
		"VALUES ($1,$2 ,$3)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_arch_diag,
		Add.Desc_index,
		Add.Description,
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//edit des arch
func (ExampleModel Models) EditArchDesc(Edit ArchDestask) bool {

	sqlStatement2 := "UPDATE tbl_architecture_diagram_dec " +
		"SET description = $1, index = $2 " +
		"WHERE id_architecture_diagram = $3 AND id = $4 "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Description,
		Edit.Desc_index,
		Edit.Id_arch_diag,
		Edit.Id,
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//delete des arch
func (ExampleModel Models) DelArchDesc(Id int) bool {

	sqlStatement2 := "DELETE FROM tbl_architecture_diagram_dec " +
		"WHERE id = $1"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Id,
	)
	if err2 != nil {
		fmt.Println(err2)
		return false
	} else {
		fmt.Println(res2)
		return true
	}
}

//view arch
func (ExampleModel Models) ViewArch(View ArchViewtask) ArchView {

	sqlStatement3 := "SELECT tbl_architecture_diagram.id, tbl_member_belongto_project.id_user FROM tbl_architecture_diagram " +
		"INNER JOIN tbl_member_belongto_project ON tbl_architecture_diagram.creator = tbl_member_belongto_project.id " +
		"WHERE tbl_architecture_diagram.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := ArchView{}

	for res3.Next() {
		task := ArchViewtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Diagram, &task.Creator, &task.Create_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.ArchView = append(result.ArchView, task)
	}

	return result

}
