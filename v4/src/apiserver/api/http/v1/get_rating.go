package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type RatingRequest struct {
	// AuthedRequest `valid:"optional"`
}

type RatingResponse struct {
	Rating
}

func (a *api) GetRating(c echo.Context, _ RatingRequest) error {
	data, err := a.core.GetUserRating(c.Request().Context(), c.Get("username").(string))
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	resp := RatingResponse{
		Rating: Rating(data),
	}

	return c.JSON(http.StatusOK, &resp)
}
