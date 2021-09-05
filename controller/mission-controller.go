package controller

type MissionRequestHandler interface {
	GetAllMissions()
}

type MissionController struct {
}

func (controller *MissionController) GetAllMissions() {
}
