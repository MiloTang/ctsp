package user

import (
	"fmt"
	"webnet/config"
	"webnet/mysql"
)

func GetUser(user ...interface{}) map[string]interface{} {
	fmt.Println(config.Conf["homegetuser"])
	fmt.Println(user...)
	mysql.Connect()
	res := make(map[string]interface{}, 100)
	defer mysql.Close()
	u := mysql.Select(config.Conf["homegetuser"], user...)
	for _, v := range u {
		for k1, v1 := range v {
			res[k1] = v1
		}
	}
	return res
}
func GetUsers() []map[string]interface{} {
	mysql.Connect()
	ress := make([]map[string]interface{}, 0)
	defer mysql.Close()
	us := mysql.Select(config.Conf["homegetusers"])
	for _, v := range us {
		//		for k1, v1 := range v {
		//			switch t := v1.(type) {
		//			case int64:
		//				res[k1] = strconv.FormatInt(v1.(int64), 10)
		//			case string:
		//				res[k1] = v1.(string)
		//			default:
		//				fmt.Println(t)
		//			}
		//		}
		ress = append(ress, v)
	}
	return ress
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
