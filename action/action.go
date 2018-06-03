package action

import (
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
	t, _ := template.ParseFiles("templates/home/index.html")
	t.Execute(w, nil)
}

func Admin(w http.ResponseWriter, r *http.Request) {
	user.GetUser("10000000")
	csid, b := gocs.StartCS(w, r, "admin", gocs.CS)
	fmt.Println("------>", b)
	gocs.CS.SetSession(csid, gocs.Key, "test")
	fmt.Fprintln(w, "这是管理页面!")
}
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	p.Title = "城市小站"
	p.KeyWord = "城市小站，为每一个城市生活更加便利"
	t, _ := template.ParseFiles("templates/home/error.html")
	t.Execute(w, p)
}
