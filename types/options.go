package types

import (
	"fmt"
	"reflect"

	"github.com/lanxre/go-valo/types/enums"
)

type Validator interface {
	Validate() error
}

type ValorantAccountParams struct {
	Name string `validate:"required"`
	Tag  string `validate:"required"`
}

type PlayerParams struct {
	PUUID string `validate:"required"`
}

type PlayerRegionParams struct {
	PUUID  string       `validate:"required"`
	Region enums.Region `validate:"required"`
}

type PlayerRegionPlatformParams struct {
	PUUID    string         `validate:"required"`
	Region   enums.Region   `validate:"required"`
	Platform enums.Platform `validate:"required"`
}

type LeaderboardParams struct {
	Region   enums.Region   `validate:"required"`
	Platform enums.Platform `validate:"required"`
}

type MatchlistParamV3 struct {
	Name   string       `validate:"required"`
	Tag    string       `validate:"required"`
	Region enums.Region `validate:"required"`
}

type MatchlistParamV4 struct {
	Name     string         `validate:"required"`
	Tag      string         `validate:"required"`
	Region   enums.Region   `validate:"required"`
	Platform enums.Platform `validate:"required"`
}

type MatchParamV2 struct {
	Uuid string `validate:"required"`
}

type MatchParamV4 struct {
	Uuid   string       `validate:"required"`
	Region enums.Region `validate:"required"`
}

type WebsiteNewsParams struct {
	Countrycode enums.Region `validate:"required"`
}

func Validate(p any) error {
	v := reflect.ValueOf(p)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		tag := field.Tag.Get("validate")

		if tag == "required" {
			if value.Kind() == reflect.String && value.String() == "" {
				return fmt.Errorf("%s is required", field.Name)
			}
		}
	}
	return nil
}

func (p ValorantAccountParams) Validate() error {
	return Validate(p)
}

func (p PlayerParams) Validate() error {
	return Validate(p)
}

func (p PlayerRegionParams) Validate() error {
	return Validate(p)
}

func (p PlayerRegionPlatformParams) Validate() error {
	return Validate(p)
}

func (p LeaderboardParams) Validate() error {
	return Validate(p)
}

func (p MatchlistParamV3) Validate() error {
	return Validate(p)
}

func (p MatchlistParamV4) Validate() error {
	return Validate(p)
}

func (p MatchParamV2) Validate() error {
	return Validate(p)
}

func (p MatchParamV4) Validate() error {
	return Validate(p)
}

func (p WebsiteNewsParams) Validate() error {
	return Validate(p)
}
