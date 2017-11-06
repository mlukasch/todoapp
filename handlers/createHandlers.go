package handlers

import (
	"errors"
	"fmt"
	"html"
	"net/http"
	"templatetest22/db"
	"templatetest22/utils"
	"templatetest22/validators"
)

func (this *HandlerConfig) UserNew(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UserNew")
	err := r.ParseMultipartForm(10000000)
	if err != nil {
		fmt.Println("error parsing form")
		utils.HandleAsNotFound(err, w)
		return
	}

	// Validation:
	validator := validators.NewForm()
	validator.AddField("firstName", r.FormValue("firstName")).Add(validators.NotEmptyValidator)
	validator.AddField("lastName", r.FormValue("lastName")).Add(validators.NotEmptyValidator)
	validator.AddField("email", r.FormValue("email")).Add(validators.NotEmptyValidator)
	validator.AddField("password", r.FormValue("password")).Add(validators.NotEmptyValidator)
	ok, errs := validator.Execute()
	response := map[string]interface{}{
		"errors": errs,
		"ok":     ok,
	}

	if !ok {
		utils.HandleJsonResponse(response, w)
		return
	}

	todoUser := db.TodoUser{
		FirstName: html.EscapeString(r.FormValue("firstName")),
		LastName:  html.EscapeString(r.FormValue("lastName")),
		Email:     html.EscapeString(r.FormValue("email")),
		Password:  r.FormValue("password"),
	}
	_, err = db.InsertTodoUser(this.db, &todoUser)
	if err != nil {
		response["ok"] = false
		response["errors"] = errors.New("DB Operation failed")
		utils.HandleJsonResponse(response, w)
		return
	}
	utils.HandleJsonResponse(response, w)
}
