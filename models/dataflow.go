package models

import "fmt"

type Dataflowtask struct {
	Email           string `json:"email"`
	Id              int    `json:"id"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
	Link_dataflow   string `json:"link_dataflow"`
	Description     string `json:"description"`
	Status          string `json:"status"`
}

type DataflowDescViewtask struct {
	Email           string `json:"email"`
	Id_usecase_desc int    `json:"id_usecase_desc"`
	//Id_usecase  int    `json:"id_usecase"`
	//Index       int    `json:"index"`
	//Name        string `json:"name"`
	//Description string `json:"description"`
	//Creator     int    `json:"creator"`
	//Create_date string `json:"create_date"`
}

type DataflowDescView struct {
	//Email     string          `json:"email"`
	DataflowDescView []DataflowDescViewtask `json:"dataflow_desc_view"`
}

type DataflowUjourViewtask struct {
	Email          string `json:"email"`
	Id_userjourney int    `json:"id_userjourney"`
	//Id_dataflow    int    `json:"id_dataflow"`
	//Name           string `json:"name"`
	//Initiate_state int    `json:"initiate_state"`
	//Userjourney    string `json:"userjourney"`
	//Index          int    `json:"index"`
	//Creator        int    `json:"creator"`
	//Description    string `json:"description"`
	//Status         string `json:"status"`
	//Create_date    string `json:"create_date"`
	//Approval       string `json:"approval"`
}

type DataflowUjourView struct {
	//Email     string          `json:"email"`
	DataflowUserjourView []DataflowUjourViewtask `json:"dataflow_userjour_view"`
}

type DataflowStructViewtask struct {
	Id_dataflow_structure int    `json:"id_dataflow_Structure"`
	Email                 string `json:"email"`
	//Id_dataflow   int    `json:"id_dataflow"`
	//Index         int    `json:"index"`
	//Name          string `json:"name"`
	//Protocol      int    `json:"protocol"`
	//Type          int    `json:"type"`
	//Address       string `json:"address"`
	//Request_type  string `json:"request_type"`
	//Response_type string `json:"response_type"`
	//Creator       int    `json:"creator"`
	//Create_date   string `json:"create_date"`
	//Approval      string `json:"approval"`
}

type DataflowStructView struct {
	//Email     string          `json:"email"`
	DataflowStructView []DataflowStructViewtask `json:"dataflow_struct_view"`
}

type DataflowProjectViewtask struct {
	Id_project int    `json:"id_project"`
	Email      string `json:"email"`
	//Id_dataflow   int    `json:"id_dataflow"`
	//Index         int    `json:"index"`
	//Name          string `json:"name"`
	//Protocol      int    `json:"protocol"`
	//Type          int    `json:"type"`
	//Address       string `json:"address"`
	//Request_type  string `json:"request_type"`
	//Response_type string `json:"response_type"`
	//Creator       int    `json:"creator"`
	//Create_date   string `json:"create_date"`
	//Approval      string `json:"approval"`
}

type DataflowProjectView struct {
	//Email     string          `json:"email"`
	DataflowProjectView []DataflowProjectViewtask `json:"dataflow_project_view"`
}

//add dataflow
func (ExampleModel Models) AddDataflow(Add Dataflowtask) bool {

	sqlStatement2 :=
		"INSERT INTO  tbl_dataflow (id_usecase, dataflow, description, creator, create_date, approval) " +
			"VALUES ($1,$2 ,$3, $4, now()::timestamp, $5)"
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Add.Id_usecase_desc,
		Add.Link_dataflow,
		Add.Description,
		2,
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

//edit dataflow
func (ExampleModel Models) EditDataflow(Edit Dataflowtask) bool {

	sqlStatement2 := "UPDATE tbl_dataflow " +
		"SET id_usecase = $2, dataflow = $1, description = $3, creator = $4, approval = $5 " +
		"WHERE id = $6  "
	res2, err2 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement2,
		Edit.Link_dataflow,
		5,
		Edit.Description,
		2,
		Edit.Status,
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

//view dataflow by description
func (ExampleModel Models) ViewDataflowDesc(View DataflowDescViewtask) DataflowDescView {

	sqlStatement3 := "SELECT tbl_usecase_desc.id, tbl_member_belongto_project.id_user  FROM tbl_dataflow " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_usecase_desc ON tbl_dataflow.id_usecase = tbl_usecase_desc.id " +
		"WHERE tbl_usecase_desc.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_usecase_desc,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DataflowDescView{}

	for res3.Next() {
		task := DataflowDescViewtask{}
		err3 := res3.Scan(&task.Id_usecase_desc, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_usecase, &task.Index, &task.Name, &task.Description, &task.Creator, &task.Create_date)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowDescView = append(result.DataflowDescView, task)
	}

	return result

}

//view dataflow by userjourney
func (ExampleModel Models) ViewUjour(View DataflowUjourViewtask) DataflowUjourView {

	sqlStatement3 := "SELECT tbl_user_journey.id, tbl_member_belongto_project.id_user  FROM tbl_dataflow " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_user_journey ON tbl_dataflow.id = tbl_user_journey.id_dataflow " +
		"WHERE tbl_user_journey.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_userjourney,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DataflowUjourView{}

	for res3.Next() {
		task := DataflowUjourViewtask{}
		err3 := res3.Scan(&task.Id_userjourney, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_dataflow, &task.Index, &task.Name, &task.Initiate_state, &task.Userjourney,
		//	&task.Index, &task.Creator, &task.Description, &task.Status, &task.Create_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowUserjourView = append(result.DataflowUserjourView, task)
	}

	return result

}

//view dataflow by dataflow structure
func (ExampleModel Models) ViewDflowStruct(View DataflowStructViewtask) DataflowStructView {

	sqlStatement3 := "SELECT tbl_dataflow_structure.id, tbl_member_belongto_project.id_user  FROM tbl_dataflow " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow.creator = tbl_member_belongto_project.id " +
		"INNER JOIN tbl_dataflow_structure ON tbl_dataflow.id = tbl_dataflow_structure.id_dataflow " +
		"WHERE tbl_dataflow_structure.id=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_dataflow_structure,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DataflowStructView{}

	for res3.Next() {
		task := DataflowStructViewtask{}
		err3 := res3.Scan(&task.Id_dataflow_structure, &task.Email)
		//err3 := res3.Scan(&task.Id, &task.Id_dataflow, &task.Index, &task.Name, &task.Protocol, &task.Type,
		//	&task.Index, &task.Request_type, &task.Response_type, &task.Creator, &task.Create_date, &task.Approval)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}
		result.DataflowStructView = append(result.DataflowStructView, task)
	}

	return result

}

//view dataflow by project
func (ExampleModel Models) ViewDflowProj(View DataflowProjectViewtask) DataflowProjectView {

	sqlStatement3 := "SELECT  tbl_member_belongto_project.id_project, tbl_member_belongto_project.creator  FROM tbl_dataflow " +
		"INNER JOIN tbl_member_belongto_project ON tbl_dataflow.creator = tbl_member_belongto_project.id  " +
		"WHERE tbl_member_belongto_project.id_project=$1 "
	res3, err3 := ExampleModel.db.GetDatabaseConfig().Query(sqlStatement3,
		View.Id_project,
	)
	if err3 != nil {
		fmt.Println(err3)

	} else {
		fmt.Println(res3)
	}
	result := DataflowProjectView{}

	for res3.Next() {
		task := DataflowProjectViewtask{}
		err3 := res3.Scan(&task.Id_project, &task.Email)
		// Exit if we get an error
		if err3 != nil {
			fmt.Println(err3)

		}

		result.DataflowProjectView = append(result.DataflowProjectView, task)
	}

	return result

}
