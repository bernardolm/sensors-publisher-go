package mock

import (
	"strings"
	"time"

	"golang.org/x/text/cases"
)

type mock struct {
	caser             cases.Caser
	class             string
	icon              string
	id                string
	manufacturer      string
	model             string
	picture           string
	time              time.Time
	unitOfMeasurement string
}

func (s *mock) Class() string             { return s.class }
func (s *mock) Icon() string              { return s.icon }
func (s *mock) ID() string                { return s.id }
func (s *mock) Manufacturer() string      { return s.manufacturer }
func (s *mock) Model() string             { return s.model }
func (s *mock) Picture() string           { return s.picture }
func (s *mock) Time() time.Time           { return s.time }
func (s *mock) UnitOfMeasurement() string { return s.unitOfMeasurement }

func (s *mock) Name() string {
	return s.caser.String(strings.ReplaceAll(s.class, "_", " ") + " mock")
}
