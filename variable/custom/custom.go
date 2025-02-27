package custom

import (
	"sort"
	"strings"

	"github.com/K-Phoen/sdk"
)

// Option represents an option that can be used to configure a custom variable.
type Option func(constant *Custom)

// ValuesMap represent a "label" to "value" map of options for a custom variable.
type ValuesMap map[string]string

func (values ValuesMap) asQuery() string {
	valuesList := make([]string, 0, len(values))

	for _, value := range values {
		valuesList = append(valuesList, value)
	}

	sort.Strings(valuesList)

	return strings.Join(valuesList, ",")
}

func (values ValuesMap) getRev(value string) (label string, ok bool) {
	for label, val := range values {
		if val == value {
			return label, true
		}
	}
	return "", false
}

func (values ValuesMap) labelsFor(vals ...string) *sdk.StringSliceString {
	out := make([]string, len(vals))
	for i, v := range vals {
		label, ok := values.getRev(v)
		if ok {
			out[i] = label
		} else {
			out[i] = v
		}
	}

	return &sdk.StringSliceString{Value: out, Valid: true}
}

// Custom represents a "custom" templated variable.
type Custom struct {
	Builder sdk.TemplateVar
	values  ValuesMap
}

// New creates a new "custom" templated variable.
func New(name string, options ...Option) *Custom {
	custom := &Custom{Builder: sdk.TemplateVar{
		Name:    name,
		Label:   name,
		Type:    "custom",
		Options: []sdk.Option{},
	}}

	for _, opt := range options {
		opt(custom)
	}

	return custom
}

// Values sets the possible values for the variable.
func Values(values ValuesMap) Option {
	return func(custom *Custom) {
		for label, value := range values {
			custom.Builder.Options = append(custom.Builder.Options, sdk.Option{
				Text:  label,
				Value: value,
			})
		}

		sort.Slice(custom.Builder.Options, func(i, j int) bool {
			return custom.Builder.Options[i].Text < custom.Builder.Options[j].Text
		})

		custom.values = values
		custom.Builder.Query = values.asQuery()
	}
}

// Default sets the default value of the variable.
func Default(values ...string) Option {
	return func(custom *Custom) {
		custom.Builder.Current = sdk.Current{
			Text:  custom.values.labelsFor(values...),
			Value: values,
		}
	}
}

// Label sets the label of the variable.
func Label(label string) Option {
	return func(custom *Custom) {
		custom.Builder.Label = label
	}
}

// HideLabel ensures that this variable's label will not be displayed.
func HideLabel() Option {
	return func(custom *Custom) {
		custom.Builder.Hide = 1
	}
}

// Hide ensures that the variable will not be displayed.
func Hide() Option {
	return func(custom *Custom) {
		custom.Builder.Hide = 2
	}
}

// Multi allows several values to be selected.
func Multi() Option {
	return func(custom *Custom) {
		custom.Builder.Multi = true
	}
}

// IncludeAll adds an option to allow all values to be selected.
func IncludeAll() Option {
	return func(custom *Custom) {
		custom.Builder.IncludeAll = true
		custom.Builder.Options = append(custom.Builder.Options, sdk.Option{
			Text:  "All",
			Value: "$__all",
		})
	}
}

// AllValue define the value used when selecting the "All" option.
func AllValue(value string) Option {
	return func(custom *Custom) {
		custom.Builder.AllValue = value
	}
}
