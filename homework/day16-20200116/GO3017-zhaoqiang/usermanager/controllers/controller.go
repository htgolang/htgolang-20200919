package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
	"zhao/models"
	"zhao/monitoring"
	"zhao/server"
)

type Response struct {
	Writer http.ResponseWriter
	Code   int
}

func (w *Response) Header() http.Header {
	return w.Writer.Header()
}

func (w *Response) Write(n []byte) (int, error) {
	return w.Writer.Write(n)
}

func (w *Response) WriteHeader(code int) {
	w.Code = code
	w.Writer.WriteHeader(code)
}

type Context struct {
	Output *Response
	Input  *http.Request
}

type HandlerrFunc func(ctx *Context)

func WrapHandler(action HandlerrFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := r.Header.Get("Authorization")
		if password == "" || strings.TrimSpace(password) == "Basic" {
			w.Header().Set("www-authenticate", "basic")
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Header()
		if server.AuthLogin(password) {

			resp := &Response{
				Writer: w,
				Code:   200,
			}
			ctx := &Context{
				Output: resp,
				Input:  r,
			}

			monitoring.RequestCountTotal.Inc()
			monitoring.RequestPathCounter.WithLabelValues(r.URL.EscapedPath()).Inc()
			starttime := time.Now()
			action(ctx)
			resuTime := time.Now().Sub(starttime)
			fmt.Println(resuTime)
			monitoring.RequestTime.Observe(float64(resuTime.Seconds()) / 1000)
			monitoring.RequestStatusCode.WithLabelValues(strconv.Itoa(ctx.Output.Code)).Inc()
			return
		}
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		w.Write([]byte("authtication error 2"))
	}
}

func GetUsers(ctx *Context) {
	tpl, err := template.ParseFiles("template/user.html")
	if err != nil {
		fmt.Println(err)
		ctx.Output.WriteHeader(500)
		return
	}

	err = tpl.ExecuteTemplate(ctx.Output, "user.html", server.GetUsers())
	if err != nil {
		fmt.Println(err)
		ctx.Output.WriteHeader(500)
		return
	}
}
func AddUser(ctx *Context) {
	if ctx.Input.Method == http.MethodGet {
		tpl, err := template.ParseFiles("template/create.html")
		if err != nil {
			ctx.Output.WriteHeader(500)
			fmt.Println("/create", err)
			return
		}
		err = tpl.ExecuteTemplate(ctx.Output.Writer, "create.html", nil)
		if err != nil {
			ctx.Output.WriteHeader(500)
			fmt.Println("/create", err)
			return
		}
	}
	if ctx.Input.Method == http.MethodPost {
		err := server.AddUser(
			ctx.Input.FormValue("name"),
			ctx.Input.FormValue("sex") == "true",
			ctx.Input.FormValue("addr"),
			ctx.Input.FormValue("tel"),
			ctx.Input.FormValue("brithday"),
			ctx.Input.FormValue("passwd"),
		)
		if err != nil {
			fmt.Println(err)
			ctx.Output.WriteHeader(500)
			ctx.Output.Write([]byte(err.Error()))
			return
		}
		http.Redirect(ctx.Output, ctx.Input, "/", 301)
	}
}

func DeleteUser(ctx *Context) {
	id, err := strconv.ParseInt(ctx.Input.FormValue("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.Output.WriteHeader(500)
		return
	}
	err = server.DeleteUser(id)
	if err != nil {
		fmt.Println(err)
		ctx.Output.WriteHeader(500)
		return
	}
	http.Redirect(ctx.Output, ctx.Input, "/", 301)
}

func ModifyUser(ctx *Context) {
	if ctx.Input.Method == http.MethodGet {
		user, err := server.ParseUser(ctx.Input)
		if err != nil {
			fmt.Println(err)
			ctx.Output.WriteHeader(500)
			ctx.Output.Write([]byte(err.Error()))
			return
		}

		// fmt.Printf("GET方法获取的用户信息====》 %#v\n", user)
		tpl, err := template.ParseFiles("template/modify.html")
		if err != nil {
			fmt.Println(err)
			ctx.Output.WriteHeader(500)
			return
		}

		err = tpl.ExecuteTemplate(ctx.Output, "modify.html",
			struct {
				*models.User
				Err string
			}{user, ctx.Input.FormValue("err")})
		if err != nil {
			fmt.Println(err)
			ctx.Output.WriteHeader(500)
			return
		}
	}

	if ctx.Input.Method == http.MethodPost {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		user, err := server.ParseUser(ctx.Input)
		if err != nil {
			fmt.Println(err)
			http.Redirect(ctx.Output, ctx.Input, fmt.Sprintf("/modify?id=%v&name=%v&sex=%v&tel=%v&brithday=%v&addr=%v&passwd=%v&err=%v", user.ID, user.Name, user.Sex, user.Tel, user.Brithday, user.Addr, user.Passwd, err.Error()), 302)
			return
		}
		fmt.Printf("POST方法获取的用户信息: %#v\n", user)
		if err := server.ModifyAuth(user); err != nil {
			fmt.Println(err)
			http.Redirect(ctx.Output, ctx.Input, fmt.Sprintf("/modify?id=%v&name=%v&sex=%v&tel=%v&brithday=%v&addr=%v&passwd=%v&err=%v", user.ID, user.Name, user.Sex, user.Tel, user.Brithday, user.Addr, user.Passwd, err.Error()), 302)
			return
		}

		err = server.ModifyUser(user)
		if err != nil {
			fmt.Println(err)
			http.Redirect(ctx.Output, ctx.Input, fmt.Sprintf("/modify?id=%v&name=%v&sex=%v&tel=%v&brithday=%v&addr=%v&passwd=%v&err=%v", user.ID, user.Name, user.Sex, user.Tel, user.Brithday, user.Addr, user.Passwd, err.Error()), 302)
			return
		}
		http.Redirect(ctx.Output, ctx.Input, "/", 301)
	}
}
