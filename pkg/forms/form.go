package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

var EmailRX = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) MinLength(field string, minLength int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < minLength {
		f.Errors.Add(field, fmt.Sprintf("This field is too short. Min lenght %d", minLength))
	}
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}

	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

func (f *Form) Require(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
