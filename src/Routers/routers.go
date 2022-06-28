package Routes

import (
	"fmt"
	"net/http"

	controler "github.com/Noexperience-Team/carsrent/src/Controlers"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"*"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	appRouter := r.PathPrefix("/api").Subrouter()

	appRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

	})

	appRouter.HandleFunc("/cars", controler.GetCars).Methods("GET")
	appRouter.HandleFunc("/cars", controler.AddCars).Methods("POST", "OPTION")
	appRouter.HandleFunc("/cars/{registration}/{rent}", controler.UpdateCarRentStatus).Methods("POST", "OPTION")
	r.Use(cors)
	return r
}
