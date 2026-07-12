package user

type UserHanlder struct {
    userService *UserService
}


func NewUserHandler(userService *UserService) *UserHanlder {
    return &UserHanlder{
        userService: userService,
    }
}