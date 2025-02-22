package main

import (
	"database/sql"
	"farmercookbook/internal/handlers"
	database "farmercookbook/internal/store/db"
	"farmercookbook/internal/store/dbstore"
	"farmercookbook/internal/utils"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template
var db *sql.DB

func init() {
	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db := database.MustOpen("cookbook")

	userStore := dbstore.NewUserStore(
		dbstore.NewUserStoreParams{
			DB: db,
		})

	sessionStore := dbstore.NewSessionStore(
		dbstore.NewSessionStoreParams{
			DB: db,
		})

	recipeStore := dbstore.NewRecipeStore(
		dbstore.NewRecipeStoreParams{
			DB: db,
		})

	utils.PrintDebugJSON("user store", userStore)
	utils.PrintDebugJSON("session", sessionStore)
	utils.PrintDebugJSON("recipe", recipeStore)

	homeHandler := handlers.NewHomeHandler(handlers.HomeHandlerDeps{Tmpl: tmpl}).ServeHTTP
	recipesHandler := handlers.NewRecipesHandler(handlers.RecipesHandlerDeps{Store: recipeStore, Tmpl: tmpl})

	gRouter := mux.NewRouter()
	gRouter.HandleFunc("/", homeHandler).Methods("GET")
	gRouter.HandleFunc("/recipes", recipesHandler.GetRecipes).Methods("GET")
	gRouter.HandleFunc("/recipes", recipesHandler.CreateRecipe).Methods("POST")
	gRouter.HandleFunc("/recipes/new", recipesHandler.NewRecipe).Methods("GET")
	gRouter.HandleFunc("/recipes/{id}", recipesHandler.DeleteRecipe).Methods("DELETE")

	http.ListenAndServe(":4000", gRouter)
}
