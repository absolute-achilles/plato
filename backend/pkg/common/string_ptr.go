package common

func StringPtr(v string) *string {
	return new(v)
}
