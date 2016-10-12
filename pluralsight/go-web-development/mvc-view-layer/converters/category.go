package converters

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/models"
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/viewmodels"
)

// ConvertCategoryToViewModel - Converts the Category model into data consumable by the View Model
func ConvertCategoryToViewModel(category models.Category) viewmodels.Category {
	result := viewmodels.Category{
		ImageURL:      category.ImageURL(),
		Title:         category.Title(),
		Description:   category.Description(),
		IsOrientRight: category.IsOrientRight(),
		ID:            category.ID(),
	}
	return result
}
