package repository

type UserMgmtRepository struct {
	RoleModuleRepository
	UserRepository
}

func NewUserMgmtRepository() *UserMgmtRepository {
	userMgmtRepo := &UserMgmtRepository{}
	userMgmtRepo.RoleModuleRepository = *NewRoleModuleRepository()
	userMgmtRepo.UserRepository = *NewUserRepository()
	return userMgmtRepo
}
