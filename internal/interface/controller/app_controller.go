package controller

type (
	AppController interface {
		UserController
	}

	appController struct {
		UserController
	}
)

func NewAppController(userController UserController) AppController {
	return &appController{
		UserController: userController,
	}
}
