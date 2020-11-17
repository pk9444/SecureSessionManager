package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/gorilla/sessions"
)

var store3 = sessions.NewCookieStore([]byte("sessionkey"))

var roles = []string{"admin", "assistant", "alien"}
var assignedRole string

//--------------------------------------------------------------------------------------//

func index3(response http.ResponseWriter, request *http.Request) {
	//io.WriteString(response, "ffdfdf")
	t, _ := template.ParseFiles("components/login.html")
	t.Execute(response, nil)

}

func login3(response http.ResponseWriter, request *http.Request) {
	//io.WriteString(response, "ffdfdf")
	//-------------------------set credentials : make it global anyhow---------------------------//
	var credentials map[string]string
	/* create a map*/
	credentials = make(map[string]string)

	/* insert key-value pairs in the map*/
	credentials["sherlock"] = "123"
	credentials["watson"] = "456"
	credentials["client"] = "789"

	var usersArray []string
	for i := range credentials {
		usersArray = append(usersArray, i)
	}

	//---------------------------------------------------------------//
	request.ParseForm()
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	user, ok := credentials[username]
	if ok && password == user {
		session, _ := store3.Get(request, "sessionkey")
		session.Values["username"] = username
		session.Save(request, response)
		http.Redirect(response, request, "/actions", http.StatusSeeOther)

		//role assignment
		if username == usersArray[0] && password == credentials[usersArray[0]] {
			assignedRole = roles[0]
		} else if username == usersArray[1] && password == credentials[usersArray[1]] {
			assignedRole = roles[1]
		} else if username == usersArray[2] && password == credentials[usersArray[2]] {
			assignedRole = roles[2]
		}

	} else {

		data := map[string]interface{}{
			"err": "Invalid",
		}
		t, _ := template.ParseFiles("components/login.html")
		t.Execute(response, data)
	}

	//---------uncomment later if not tested ----------//
	// if username == usersArray[0] && password == credentials["sherlock"] {

	// 	session, _ := store3.Get(request, "sessionkey")
	// 	session.Values["username"] = username
	// 	session.Save(request, response)
	// 	http.Redirect(response, request, "/actions", http.StatusSeeOther)
	// } else {
	// 	data := map[string]interface{}{
	// 		"err": "Invalid",
	// 	}
	// 	t, _ := template.ParseFiles("components/login.gohtml")
	// 	t.Execute(response, data)
	// }

}

func actions3(response http.ResponseWriter, request *http.Request) {
	//io.WriteString(response, "ffdfdf")
	session, _ := store3.Get(request, "sessionkey")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	t, _ := template.ParseFiles("components/actions.html")
	t.Execute(response, data)

}

//##########################################################################################//
func public3(response http.ResponseWriter, request *http.Request) {

	//---------------------------------------------------------------//
	session, _ := store3.Get(request, "sessionkey")
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	t, _ := template.ParseFiles("components/public.html")
	t.Execute(response, data)

}

func private3(response http.ResponseWriter, request *http.Request) {

	// request.ParseForm()
	// username := request.Form.Get("username")
	// password := request.Form.Get("password")
	//---------------------------------------------------------------//
	if assignedRole == "admin" {
		session, _ := store3.Get(request, "sessionkey")
		username2 := session.Values["username"]
		data := map[string]interface{}{
			"username": username2,
		}
		t, _ := template.ParseFiles("components/private.html")
		t.Execute(response, data)
	} else if assignedRole == "assistant" || assignedRole == "alien" {
		io.WriteString(response, "Access Denied")
	}

}

func protected3(response http.ResponseWriter, request *http.Request) {

	//---------------------------------------------------------------//
	// request.ParseForm()
	// username := request.Form.Get("username")
	// password := request.Form.Get("password")

	if assignedRole == "alien" {

		io.WriteString(response, "Access Denied")

	} else if assignedRole == "admin" || assignedRole == "assistant" {
		session, _ := store3.Get(request, "sessionkey")
		username2 := session.Values["username"]
		data := map[string]interface{}{
			"username": username2,
		}
		t, _ := template.ParseFiles("components/protected.html")
		t.Execute(response, data)
	}
	// session, _ := store3.Get(request, "sessionkey")
	// username2 := session.Values["username"]
	// data := map[string]interface{}{
	// 	"username": username2,
	// }
	// t, _ := template.ParseFiles("components/protected.html")
	// t.Execute(response, data)
}

//#############################################################################################//
func logout3(response http.ResponseWriter, request *http.Request) {
	//io.WriteString(response, "6546454")
	session, _ := store3.Get(request, "sessionkey")
	session.Options.MaxAge = -1 //immediately expire the cookie when its saved
	session.Save(request, response)
	//http.Redirect(response, request, "/logout", http.StatusSeeOther)

	//session, _ := store3.Get(request, "sessionkey")

	//----uncomment here for reverting --------//
	username := session.Values["username"]
	data := map[string]interface{}{
		"username": username,
	}
	t, _ := template.ParseFiles("components/logout.html")
	t.Execute(response, data)

}

func main() {

	http.HandleFunc("/", index3)
	http.HandleFunc("/login", login3)
	http.HandleFunc("/private", private3)
	http.HandleFunc("/public", public3)
	http.HandleFunc("/protected", protected3)
	http.HandleFunc("/actions", actions3)
	http.HandleFunc("/logout", logout3)
	http.ListenAndServe(":3002", nil)
}
