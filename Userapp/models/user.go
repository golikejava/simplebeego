package models

//"errors"
// "strconv"
// "time"

// var (
// 	UserList map[string]*User
// )

// func init() {
// 	UserList = make(map[string]*User)
// 	u := User{"user_11111", "astaxie"}
// 	UserList["user_11111"] = &u
// }

type User struct {
	Name    string
	Address string
}

type CTX struct {
	Name    string
	Address string
}

// type Profile struct {
// 	Gender  string
// 	Age     int
// 	Address string
// 	Email   string
// }

// func AddUser(u User) string {
// 	u.name = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.name] = &u
// 	return u.name

// 	u.address = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
// 	UserList[u.address] = &u
// 	return u.address
// }

// func GetAllUsers() map[string]*User {
// 	return UserList
// }

// func GetUser(uid string) (u *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		return u, nil
// 	}
// 	return nil, errors.New("User not exists")
// }

// func UpdateUser(uid string, uu *User) (a *User, err error) {
// 	if u, ok := UserList[uid]; ok {
// 		if uu.Username != "" {
// 			u.Username = uu.Username
// 		}
// 		if uu.Password != "" {
// 			u.Password = uu.Password
// 		}
// 		if uu.Profile.Age != 0 {
// 			u.Profile.Age = uu.Profile.Age
// 		}
// 		if uu.Profile.Address != "" {
// 			u.Profile.Address = uu.Profile.Address
// 		}
// 		if uu.Profile.Gender != "" {
// 			u.Profile.Gender = uu.Profile.Gender
// 		}
// 		if uu.Profile.Email != "" {
// 			u.Profile.Email = uu.Profile.Email
// 		}
// 		return u, nil
// 	}
// 	return nil, errors.New("User Not Exist")
// }

// func Login(username, password string) bool {
// 	for _, u := range UserList {
// 		if u.Username == username && u.Password == password {
// 			return true
// 		}
// 	}
// 	return false
// }

// func DeleteUser(uid string) {
// 	delete(UserList, uid)
// }
