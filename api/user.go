package api

import (
	"encoding/json"
	"github.com/snehal1112/taxcal/model"
	"net/http"
)

func (api *API) InitUser() {
	user := api.BaseRoutes["User"]
	user.Handle("", api.ApiHandler(createUser)).Methods("POST")
	user.Handle("/login", api.ApiHandler(loginUser)).Methods("POST")
}

func createUser(c *Context, w http.ResponseWriter, r *http.Request) {
	user := model.NewUser()
	userRecord := user.FromJson(r.Body)
	err := c.App.CreateUser(userRecord)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{success: false}`))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{success: true}`))
}

func loginUser(c *Context, w http.ResponseWriter, r *http.Request) {
	user := model.NewUser()
	userRecord := user.FromJson(r.Body)

	isAuth, userObj, err := c.App.AuthUser(userRecord)
	if err != nil || isAuth == false {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"success": false, "message":"Authentication failed."}`))
		return
	}

	// FIXME: move this to somewhere app folder 
	// after session create set the session id to cookies.
	// Client side use interceptor to set the access token.
	//http.SetCookie(w,&http.Cookie{
	//	Name: "CODER_ACCESS",
	//	Value: cast.ToString(isAuth),
	//	Domain: "0.0.0.0:8774",
	//	Path: "/",
	//	Expires: time.Now().Add(30 * time.Minute),
	//})

	w.WriteHeader(http.StatusCreated)
	result, _ := json.Marshal(userObj)
	w.Write(result)

}