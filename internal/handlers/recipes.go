package handlers

import (
	"farmercookbook/internal/store"
	"farmercookbook/internal/utils"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type RecipesHandler struct {
	tmpl  *template.Template
	store store.RecipeStore
}

type RecipesHandlerDeps struct {
	Tmpl  *template.Template
	Store store.RecipeStore
}

func NewRecipesHandler(deps RecipesHandlerDeps) *RecipesHandler {
	return &RecipesHandler{
		tmpl:  deps.Tmpl,
		store: deps.Store,
	}
}

func (h *RecipesHandler) GetRecipes(w http.ResponseWriter, r *http.Request) {
	log.Println("GetRecipes")
	recipes, err := h.store.GetRecipes(1000)
	if err != nil {
		log.Fatalln(err)
	}

	err = h.tmpl.ExecuteTemplate(w, "recipeList", recipes)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *RecipesHandler) NewRecipe(w http.ResponseWriter, r *http.Request) {
	err := h.tmpl.ExecuteTemplate(w, "addRecipeForm", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {

	log.Println("CreateRecipe")
	recipe := &store.Recipe{
		Name:        r.FormValue("name"),
		Image:       r.FormValue("image"),
		Description: r.FormValue("description"),
		PrepTime:    r.FormValue("prepTime"),
		CookTime:    r.FormValue("cookTime"),
		Ingredients: r.FormValue("ingredients"),
		Steps:       r.FormValue("steps"),
		CreatedDate: r.FormValue("createdDate"),
	}

	result, err := h.store.CreateRecipe(recipe)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

	http.Redirect(w, r, "/recipes", http.StatusSeeOther)
}

func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	recipeId, _ := strconv.Atoi(vars["id"])

	utils.PrintDebugJSON("vars", vars)

	recipe, err := h.store.GetRecipe(recipeId)
	if err != nil {
		log.Println(err)
	}

	err = h.tmpl.ExecuteTemplate(w, "recipeDetail", recipe)
}

func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	utils.PrintDebugJSON("vars", vars)
	recipeId, _ := strconv.Atoi(vars["id"])

	recipe, err := h.store.GetRecipe(recipeId)
	if err != nil {
		log.Println(err)
	}

	for key, value := range vars {
		switch key {
		case "name":
			recipe.Name = value
		case "image":
			recipe.Image = value
		case "description":
			recipe.Description = value
		case "prepTime":
			recipe.PrepTime = value
		case "cookTime":
			recipe.CookTime = value
		case "ingredients":
			recipe.Ingredients = value
		case "steps":
			recipe.Steps = value
		case "createdDate":
			recipe.CreatedDate = value
		}
	}

	_, err = h.store.UpdateRecipe(recipe)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/recipes", http.StatusSeeOther)

}

func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	log.Println("DeleteRecipe")
	vars := mux.Vars(r)
	recipeId, _ := strconv.Atoi(vars["id"])
	err := h.store.DeleteRecipe(recipeId)
	if err != nil {
		log.Println(err)
	}

	http.Redirect(w, r, "/recipes", http.StatusSeeOther)
}

func (h *RecipesHandler) FindRecipes(w http.ResponseWriter, r *http.Request) {
	log.Println("FindRecipes")
	err := r.ParseForm()
	utils.PrintDebugJSON("form", r.Form)

	recipes, err := h.store.FindRecipes(r.Form["query"][0])
	if err != nil {
		log.Fatal(err)
	}

	utils.PrintDebugJSON("recipes", recipes)

	err = h.tmpl.ExecuteTemplate(w, "recipeList", recipes)
}
