package main


import (

	"net/http"

)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}

	err := app.writeJSON(w, http.StatusOK, env, nil)
	if err != nil {
		app.logger.Error("error writing JSON", "error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}