package endpoints

import "auth/core"

type ServiceEndpoints struct {
	UserService UserServiceEndpoints
}

func New(domain *core.Core, config core.Config) *ServiceEndpoints {
	return &ServiceEndpoints{
		MakeUserServiceEndpoints(domain.User, config),
	}
}
