package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"webnet/action"
	"webnet/common"
	"webnet/config"
	"webnet/gocs"
	"webnet/routes"
)

func init() {
	//git remote add origin git@github.com:MiloTang/ctsp.git
	config.ConfigPath = common.RootPath()
	config.ConfigFile = config.ConfigPath + "/config/config.config"
	fmt.Println(config.ConfigFile)
	config.InitConfig(config.ConfigFile)
	gocs.MaxLT = 3600
	gocs.CookieName = "CityShop"
	gocs.Salt = "CityShopSalt"
	gocs.Key = "CityShopKey"
	gocs.CS = gocs.NewCookieSession()
}
func main() {
	http.Handle("/css/", http.FileServer(http.Dir("static")))
	http.Handle("/js/", http.FileServer(http.Dir("static")))
	http.Handle("/images/", http.FileServer(http.Dir("static")))
	http.Handle("/bootstrap/", http.FileServer(http.Dir("static")))
	http.Handle("/layoutitlib/", http.FileServer(http.Dir("static")))
	http.Handle("/wysiwyg/", http.FileServer(http.Dir("static")))
	http.Handle("/fonts/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/", Entry)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func Entry(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(r.Host)
	config.UpdateConfig(config.ConfigFile)
	urls := strings.Split(r.URL.Path, "/")
	if handler, ok := routes.Routes["/"+urls[1]]; ok {
		handler(w, r)
		return
	} else {
		action.ErrorPage(w, r)
	}
}
