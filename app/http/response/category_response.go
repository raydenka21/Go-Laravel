package response

import (
	"goravel/app/models"
	"time"
)

type CategoryResponse struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	//Updated_at time.Time
	//Deleted_at time.Time
}

func ToCategoryResponseAll(category []models.Category) []CategoryResponse {
	var categoryResponses []CategoryResponse
	for i, _ := range category {
		var newCategory CategoryResponse
		newCategory.Id = category[i].Id
		newCategory.Name = category[i].Name
		newCategory.Created_at = category[i].CreatedAt
		categoryResponses = append(categoryResponses, newCategory)
	}
	return categoryResponses

}

func ToCategoryResponseDetail(category models.Category) CategoryResponse {
	return CategoryResponse{
		Id:         category.Id,
		Name:       category.Name,
		Created_at: category.CreatedAt,
	}
}
