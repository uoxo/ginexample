// package user

// type User struct {
// 	Mi_user_id string
// 	Phone      string
// 	Rolename   string //角色名字
// 	Email      string
// 	Password   string
// 	Username   string
// 	Headpic    string
// }

// type UserModel struct{}

// //phone登入获取数据
// func (m UserModel) LoginGetData(account, isphone, account_type string) (data User, err error) {

// 	isphoneS := ""
// 	if isphone == "yes" {
// 		isphoneS = `u.phone = '` + account + `' `
// 	} else {
// 		isphoneS = `u.email = '` + account + `' `
// 	}

// 	err = mysql.GetDB().SelectOne(&data, usersql.GetLoginData_g(account_type, isphoneS))
// }