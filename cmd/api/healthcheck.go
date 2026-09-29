package main


import (

	"context"

)
type healthCheck struct  {
	Status string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
} 
type healthCheckInput struct {
}

type healthCheckOutput struct {
	Body healthCheck 
}


func (app *application) healthcheck(ctx context.Context, input *healthCheckInput) (*healthCheckOutput, error) {
	return &healthCheckOutput{
		Body: healthCheck{
			Status:      "available",
			Environment: app.config.env,
			Version:     version,
		},
	}, nil
}

// func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
// 	env := envelope{
// 		"status": "available",
// 		"system_info": map[string]string{
// 			"environment": app.config.env,
// 			"version": version,
// 		},
// 	}

// 	err := app.writeJSON(w, http.StatusOK, env, nil)
// 	if err != nil {
// 		app.logger.Error("error writing JSON", "error", err)
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// }