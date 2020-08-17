package routes

import (
	"github.com/labstack/echo"
	"plm/controller"
)

type Routing struct {
	routes controller.Controller
}

type RoutingInterface interface {
	GetRoutes() *echo.Echo
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	//user
	e.POST("/tasks/user/register/", Routing.routes.PostsRegis)
	e.POST("/tasks/user/login/", Routing.routes.GetLogin)
	e.PUT("/tasks/user/edit/", Routing.routes.UpEdit)

	//project
	e.POST("/tasks/project/create/", Routing.routes.PostsCreate)
	e.POST("/tasks/project/view/", Routing.routes.ViewProject)
	e.PUT("/tasks/project/edit/", Routing.routes.EditProject)

	//member
	e.POST("/tasks/member/add/", Routing.routes.AddMember)
	e.POST("/tasks/member/view/", Routing.routes.ViewMember)
	e.POST("/tasks/member/delete/", Routing.routes.DelMember)

	//timeline
	e.POST("/tasks/timeline/create/", Routing.routes.CreateTimel)
	e.POST("/tasks/timeline/view/", Routing.routes.ViewTimel)
	e.PUT("/tasks/timeline/update/", Routing.routes.EditTimel)

	//arch
	e.POST("/tasks/arch/add/", Routing.routes.AddArch)
	e.PUT("/tasks/arch/edit/", Routing.routes.EditArch)
	e.POST("/tasks/arch/view/", Routing.routes.ViewArch)
	e.POST("/tasks/arch/adddesc/", Routing.routes.AddArchDes)
	e.PUT("/tasks/arch/editdesc/", Routing.routes.EditArchDes)
	e.POST("/tasks/arch/deldesc/", Routing.routes.DelArchDes)

	//usecase
	e.POST("/tasks/usecase/add/", Routing.routes.AddUsec)
	e.PUT("/tasks/usecase/edit/", Routing.routes.EditUsec)
	e.POST("/tasks/usecase/view/", Routing.routes.ViewUsec)
	e.POST("/tasks/usecase/view/erd", Routing.routes.ViewUsecErd)
	e.POST("/tasks/usecase/adddesc/", Routing.routes.AddUsecDesc)
	e.PUT("/tasks/usecase/editdesc/", Routing.routes.EditUsecDesc)
	e.POST("/tasks/usecase/viewdesc/scenario", Routing.routes.ViewUsecScenar)
	e.POST("/tasks/usecase/viewdesc/dataflow", Routing.routes.ViewUsecDataflow)
	e.POST("/tasks/usecase/deldesc/", Routing.routes.DelUsecDesc)

	//dataflow
	e.POST("/tasks/dataflow/add/", Routing.routes.AddDataflow)
	e.PUT("/tasks/dataflow/edit/", Routing.routes.EditDataflow)
	e.POST("/tasks/dataflow/view/project/", Routing.routes.ViewDflowProj)
	e.POST("/tasks/dataflow/view/description/", Routing.routes.ViewDataflowDesc)
	e.POST("/tasks/dataflow/view/userjourney/", Routing.routes.ViewUJourDesc)
	e.POST("/tasks/dataflow/view/dataflowstruct/", Routing.routes.ViewDflowStruct)

	//erd
	e.POST("/tasks/erd/add/", Routing.routes.AddErd)
	e.PUT("/tasks/erd/edit/", Routing.routes.EditErd)
	e.POST("/tasks/erd/addtable/", Routing.routes.AddErdTable)
	e.PUT("/tasks/erd/edittable/", Routing.routes.EditErdTable)
	e.POST("/tasks/erd/deltable/", Routing.routes.DelErdTable)
	e.POST("/tasks/erd/addfield/", Routing.routes.AddErdField)
	e.PUT("/tasks/erd/editfield/", Routing.routes.EditErdField)
	e.POST("/tasks/erd/delfield/", Routing.routes.DelErdField)
	e.POST("/tasks/erd/viewtable/", Routing.routes.ViewErdTable)
	e.POST("/tasks/erd/view/usecase", Routing.routes.ViewErdUsec)
	e.POST("/tasks/erd/view/field", Routing.routes.ViewErdFeild)
	e.POST("/tasks/erd/viewtable/all", Routing.routes.ViewErdTableAll)
	e.POST("/tasks/erd/view/project", Routing.routes.ViewErdProj)

	//scenario
	e.POST("/tasks/usecase_scenario/add/", Routing.routes.AddScenario)
	e.PUT("/tasks/usecase_scenario/edit/", Routing.routes.EditScenario)
	e.POST("/tasks/usecase_scenario/view/", Routing.routes.ViewScen)
	e.POST("/tasks/usecase_scenario/view/project", Routing.routes.ViewScenProj)

	//dataflow structure
	e.POST("/tasks/dataflow_structure/add/", Routing.routes.AddDataFStruc)
	e.PUT("/tasks/dataflow_structure/edit/", Routing.routes.EditDataFStruc)
	e.POST("/tasks/dataflow_structure/view/dataflow", Routing.routes.ViewDflow)
	e.POST("/tasks/dataflow_structure/view/sequence", Routing.routes.ViewDflowSequence)
	e.POST("/tasks/dataflow_structure/view/design", Routing.routes.ViewDflowDesign)
	e.POST("/tasks/dataflow_structure/view/project", Routing.routes.ViewDflowStructProj)

	//userjourney
	e.POST("/tasks/userjourney/add/", Routing.routes.AddUserjrny)
	e.PUT("/tasks/userjourney/edit/", Routing.routes.EditUserjrny)
	e.POST("/tasks/userjourney/view/dataflow", Routing.routes.ViewUserjrnyDflow)
	e.POST("/tasks/userjourney/view/dataflow_structure", Routing.routes.ViewUserjrnyDflowStruct)
	e.POST("/tasks/userjourney/view/project", Routing.routes.ViewUserjrnyProj)

	//sequence diagram
	e.POST("/tasks/sequence_diagram/add/", Routing.routes.AddSeqDiag)
	e.PUT("/tasks/sequence_diagram/edit/", Routing.routes.EditSeqDiag)
	e.POST("/tasks/sequence_diagram/view/", Routing.routes.ViewSeqDiag)
	e.POST("/tasks/sequence_diagram/view/project", Routing.routes.ViewSeqDiagProj)

	//UI UX
	e.POST("/tasks/uiux/add/", Routing.routes.AddUi)
	e.PUT("/tasks/uiux/edit/", Routing.routes.EditUi)
	e.POST("/tasks/uiux/view/", Routing.routes.ViewUi)

	return e
}
