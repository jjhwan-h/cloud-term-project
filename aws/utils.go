package aws

import (
	a "github.com/aws/aws-sdk-go-v2/aws"
)

func ToInt32(n int32) *int32 {
	return a.Int32(n)
}

func ToString(s string) *string {
	return a.String(s)
}

func ToBool(b bool) *bool {
	return a.Bool(b)
}
func ptr(s string) *string {
	return &s
}
