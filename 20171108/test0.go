package main

import (
	"fmt"
	"net/http"
	"time"
)

/*
net/http包中的SetCookie来设置:
写
http.SetCookie(w ResponseWriter, cookie *Cookie)
type Cookie struct {
	Name       string
	Value      string
	Path       string
	Domain     string
	Expires    time.Time
	RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
	MaxAge   int
	Secure   bool
	HttpOnly bool
	Raw      string
	Unparsed []string // Raw text of unparsed attribute-value pairs
}
读：cookie, _ := r.Cookie("username")
for _, cookie := range r.Cookies() {
	fmt.Fprint(w, cookie.Name)
}
*/

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie)
	fmt.Fprintln(w, cookie)
	fmt.Fprintln(w, "------------------------------")
	for _, cookie := range r.Cookies() {
		fmt.Fprintln(w, cookie)
	}
}

func saveCookie(w http.ResponseWriter, r *http.Request) {
	exp := time.Now()
	exp = exp.AddDate(1, 0, 0) //有效期一年
	cookie := http.Cookie{Name: "username", Value: "hello", Expires: exp}
	http.SetCookie(w, &cookie)
}
func main() {
	http.HandleFunc("/getCookie", getCookie)
	http.HandleFunc("/saveCookie", saveCookie)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
