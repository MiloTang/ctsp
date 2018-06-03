package user

import (
	"fmt"
	"webnet/config"
	"webnet/mysql"
)

func GetUser(user ...interface{}) []string {
	fmt.Println(config.Conf["homegetuser"])
	fmt.Println(user...)
	mysql.Connect()
	defer mysql.Close()
	u := mysql.Select(config.Conf["homegetuser"], user...)
	fmt.Println(u)
	return nil
}
func GetUsers(user ...string) map[string][]string {
	return nil
}
func SetUser(user ...string) bool {
	return true
}
func UpdateUser(user ...string) bool {
	return true
}
func DeleteUser(user ...string) bool {
	return true
}
