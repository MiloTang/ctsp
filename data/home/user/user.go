package user

import (
	"fmt"
	"strconv"
	"webnet/config"
	"webnet/mysql"
)

func GetUser(user ...interface{}) map[string]string {
	fmt.Println(config.Conf["homegetuser"])
	fmt.Println(user...)
	mysql.Connect()
	res := make(map[string]string, 100)
	defer mysql.Close()
	u := mysql.Select(config.Conf["homegetuser"], user...)
	for _, v := range u {
		for k1, v1 := range v {
			switch t := v1.(type) {
			case int64:
				res[k1] = strconv.FormatInt(v1.(int64), 10)
			case string:
				res[k1] = v1.(string)
			default:
				fmt.Println(t)
			}
		}
	}
	return res
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
