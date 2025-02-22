package dbstore

import (
	"farmercookbook/internal/store"
	"farmercookbook/internal/utils"
	"gorm.io/gorm"
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
	var recipe store.Recipe

	err := s.db.First(&recipe, id).Error

	if err != nil {
		return nil, err
	}

	return &recipe, nil
}

func (s *RecipeStore) GetRecipes(max int) ([]*store.Recipe, error) {
	var recipes []*store.Recipe

	err := s.db.Limit(max).Find(&recipes).Error
	if err != nil {
		return nil, err
	}

	utils.PrintDebugJSON("recipes", recipes)
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
