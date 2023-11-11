//go:build testing
// +build testing

package core

import (
	"io"
	"log/slog"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/library"
	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/rating"
	"github.com/migregal/bmstu-iu7-ds-lab2/apiserver/core/ports/reservation"
	"github.com/migregal/bmstu-iu7-ds-lab2/pkg/readiness"
)

type TestSuite struct {
	suite.Suite

	core *Core

	mockedLibrary     *library.MockClient
	mockedRating      *rating.MockClient
	mockedReservation *reservation.MockClient
}

func (s *TestSuite) SetupTest() {
	s.mockedLibrary = library.NewMockClient(s.T())
	s.mockedReservation = reservation.NewMockClient(s.T())
	s.mockedRating = rating.NewMockClient(s.T())

	var err error
	s.core, err = New(
		slog.New(slog.NewJSONHandler(io.Discard, nil)), readiness.New(),
		s.mockedLibrary, s.mockedRating, s.mockedReservation,
	)

	require.NoError(s.T(), err, "failed to init core")
}

func (s *TestSuite) TearDownTest() {
}
