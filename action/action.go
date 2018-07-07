package action

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"webnet/data/home/user"
	"webnet/gocs"
	"webnet/page"
)

var (
	p = &page.PageHead{}
)

func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/home/index.html", "templates/home/header.html")
	t.Execute(w, nil)
}

func Admin(w http.ResponseWriter, r *http.Request) {
	res.Header().Set("Access-Control-Allow-Origin", "*")
	u := user.GetUser("10000000")
	fmt.Println(u)
	csid, b := gocs.StartCS(w, r, "admin", gocs.CS)
	fmt.Println("------>", b)
	gocs.CS.SetSession(csid, gocs.Key, "test")
	us := user.GetUsers()
	fmt.Println(us)
	bss, errs := json.Marshal(us)
	if errs != nil {
		fmt.Println("json.Marshal failed:", errs)
		return
	}
	fmt.Fprintln(w, string(bss))
}
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	p.Title = "城市小站"
	p.KeyWord = "城市小站，为每一个城市生活更加便利"
	t, _ := template.ParseFiles("templates/home/error.html")
	t.Execute(w, p)
}
