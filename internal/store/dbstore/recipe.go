package dbstore

import (
	"farmercookbook/internal/store"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type RecipeStore struct {
	db *gorm.DB
}

type NewRecipeStoreParams struct {
	DB *gorm.DB
}

func NewRecipeStore(params NewRecipeStoreParams) *RecipeStore {
	return &RecipeStore{
		db: params.DB,
	}
}

func (s *RecipeStore) CreateRecipe(recipe *store.Recipe) (*store.Recipe, error) {

	result := s.db.Create(recipe)

	if result.Error != nil {
		return nil, result.Error
	}
	return recipe, nil
}

func (s *RecipeStore) GetRecipe(id int) (*store.Recipe, error) {
	log.Println("GetRecipe")
	var recipe store.Recipe

	err := s.db.First(&recipe, id).Error

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (s *RecipeStore) GetRecipes(max int) ([]store.Recipe, error) {
	var recipes []store.Recipe

	err := s.db.Limit(max).Find(&recipes).Error
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (s *RecipeStore) UpdateRecipe(recipe *store.Recipe) (*store.Recipe, error) {
	result := s.db.Save(recipe)

	if result.Error != nil {
		return nil, result.Error
	}

	return recipe, nil
}

func (s *RecipeStore) DeleteRecipe(id int) error {
	result := s.db.Delete(&store.Recipe{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *RecipeStore) FindRecipes(query string) ([]store.Recipe, error) {
	log.Println(query)
	var recipes []store.Recipe
	q := fmt.Sprintf("%%%s%%", query)
	err := s.db.Where("name like ?", q).Find(&recipes).Error
	if err != nil {
		return recipes, err
	}

	return recipes, nil
}
