package router

import (
	"all-my-gear-backend-go/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/user", middleware.GetAllUser).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/newuser", middleware.CreateUser).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/user/{id}", middleware.UpdateUser).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteuser/{id}", middleware.DeleteUser).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/newgearitem", middleware.CreateGearitem).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/gearitem/{id}", middleware.GetGearItem).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/gearitem", middleware.GetAllGearitems).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/gearitem/{id}", middleware.UpdateGearitems).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deletegearitem/{id}", middleware.DeleteGearitem).Methods("DELETE", "OPTIONS")
    
    router.HandleFunc("/api/newchecklist", middleware.CreateChecklist).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/checklist", middleware.GetAllChecklists).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/checklist/{id}", middleware.GetChecklist).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/checklist/{id}", middleware.UpdateChecklist).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deletechecklist/{id}", middleware.DeleteChecklist).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/api/gearcheckrelations", middleware.GetAllGearCheckRelations).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/gearcheckrelations/{id}", middleware.GetRelationsByChecklistID).Methods("GET", "OPTIONS")

    return router
}