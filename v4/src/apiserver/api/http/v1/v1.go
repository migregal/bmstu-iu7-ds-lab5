package v1

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/library"
	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/rating"
	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/reservation"
	"github.com/migregal/bmstu-iu7-ds-lab2/pkg/httpwrapper"
	"github.com/migregal/bmstu-iu7-ds-lab2/pkg/oauth2/auth0/authenticator"
)

type Core interface {
	GetLibraries(context.Context, string, uint64, uint64) (library.Infos, error)
	GetLibraryBooks(context.Context, string, bool, uint64, uint64) (library.Books, error)
	GetUserRating(ctx context.Context, username string) (rating.Rating, error)
	GetUserReservations(context.Context, string) ([]reservation.FullInfo, error)
	TakeBook(
		ctx context.Context, usename, libraryID, bookID string, end time.Time,
	) (reservation.FullInfo, error)
	ReturnBook(ctx context.Context, username, reservationID, condition string, date time.Time) error
}

type api struct {
	core Core
}

func InitListener(mx *echo.Echo, lg *slog.Logger, core Core, cfg authenticator.Config) error {
	gr := mx.Group("/api/v1")

	a := api{core: core}

	auth, err := authenticator.New(cfg)
	if err != nil {
		return fmt.Errorf("init auth: %w", err)
	}

	gr.Use(authenticator.NewMiddleware(lg, auth).Handle)

	gr.GET("/libraries", httpwrapper.WrapRequest(lg, a.GetLibraries))
	gr.GET("/libraries/:id/books", httpwrapper.WrapRequest(lg, a.GetLibraryBooks))

	gr.GET("/reservations", httpwrapper.WrapRequest(lg, a.GetReservations))
	gr.POST("/reservations", httpwrapper.WrapRequest(lg, a.TakeBook))
	gr.POST("/reservations/:id/return", httpwrapper.WrapRequest(lg, a.ReturnBook))

	gr.GET("/rating", httpwrapper.WrapRequest(lg, a.GetRating))

	return nil
}
