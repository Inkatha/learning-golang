package converters

import (
	"learning-golang/pluralsight/go-web-development/mvc-view-layer/models"
	"testing"
)

func Test_ConvertsCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetImageURL("The image URL")
	category.SetTitle("The Title")
	category.SetDescription("The description")
	category.SetIsOrientRight(true)
	category.SetID(42)

	result := ConvertCategoryToViewModel(category)

	if result.ImageURL != category.ImageURL() {
		t.Log("Image URL not converted properly")
		t.Fail()
	}

	if result.Title != category.Title() {
		t.Log("Title not converted properly")
		t.Fail()
	}

	if result.Description != category.Description() {
		t.Log("Description not converted properly")
		t.Fail()
	}

	if result.IsOrientRight != category.IsOrientRight() {
		t.Log("Orientation not converted properly")
		t.Fail()
	}

	if result.ID != category.ID() {
		t.Log("Id nto converted properly")
		t.Fail()
	}
}
