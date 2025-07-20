package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/VincentLAU5142/FEM_project_GO/internal/api"
	"github.com/VincentLAU5142/FEM_project_GO/internal/store"
)

type Application struct {
	Logger *log.Logger
	//call the api folder's func
	WorkoutHandler *api.WorkoutHandler
	//database
	DB *sql.DB
}

func NewApplication() (*Application, error) {
	//database
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//our stores will go here

	//our handlers will go here
	workoutHandler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (a *Application) HeathCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
