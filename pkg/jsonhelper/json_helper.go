package jsonhelper

import "time"

// encoder func for encoding from A to B.
type Encoder[A, B comparable] func(A) B

func EncodeElement[A, B comparable](input A, encode func(A) B) B {
	var (
		defaultA A
		defaultB B
	)

	if input == defaultA {
		return defaultB
	}

	return encode(input)
}

// EncodeSlice decodes slice of domain objects to a slice of json objects.
func EncodeSlice[A, B comparable](slice []A, encode Encoder[A, B]) []B {
	if slice == nil {
		return nil
	}

	result := make([]B, len(slice))
	for i, element := range slice {
		result[i] = EncodeElement(element, encode)
	}

	return result
}

func EncodeTimeNillable(input time.Time) *time.Time {
	if input.IsZero() {
		return nil
	}
	return &input
}

func Value[T any](value *T) T {
	var defaultValue T

	if value == nil {
		return defaultValue
	}
	return *value
}
