package application

var BizCardServiceBean BizCardService
var UserServiceBean UserService

func RegisterServiceBeans() {
	SetupUserService()
	SetupBizCardService()
}
