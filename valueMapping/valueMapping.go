package valueMapping

import "github.com/K-Phoen/sdk"

type ResultOption func(result *sdk.ValueMappingResult)

func DisplayText(text string) ResultOption {
	return func(result *sdk.ValueMappingResult) {
		result.Text = &text
	}
}

func Color(color string) ResultOption {
	return func(result *sdk.ValueMappingResult) {
		result.Color = &color
	}
}

type SpecialMatchOption func(options *sdk.SpecialValueMapOptions)

func SpecialMatchTrue() SpecialMatchOption {
	return func(options *sdk.SpecialValueMapOptions) {
		options.Match = sdk.SpecialValueMatchTrue
	}
}

func SpecialMatchFalse() SpecialMatchOption {
	return func(options *sdk.SpecialValueMapOptions) {
		options.Match = sdk.SpecialValueMatchFalse
	}
}

func SpecialMatchNull() SpecialMatchOption {
	return func(options *sdk.SpecialValueMapOptions) {
		options.Match = sdk.SpecialValueMatchNull
	}
}

func SpecialMatchNaN() SpecialMatchOption {
	return func(options *sdk.SpecialValueMapOptions) {
		options.Match = sdk.SpecialValueMatchNaN
	}
}

func SpecialMatchNullAndNaN() SpecialMatchOption {
	return func(options *sdk.SpecialValueMapOptions) {
		options.Match = sdk.SpecialValueMatchNullAndNaN
	}
}

func SpecialMatchEmpty() SpecialMatchOption {
	return func(options *sdk.SpecialValueMapOptions) {
		options.Match = sdk.SpecialValueMatchEmpty
	}
}
