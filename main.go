package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/VincentLAU5142/FEM_project_GO/internal/app"
	"github.com/VincentLAU5142/FEM_project_GO/internal/routes"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
	//defer run after the app in everything else is done
	defer app.DB.Close()

	// http.HandleFunc("/health", HeathCheck)
	//call the func in the routes to instead of the above comment
	r := routes.SetupRoutes(app)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("we are running on port %d\n", port)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
