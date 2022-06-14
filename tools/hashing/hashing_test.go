package hashing_test

import (
	"fmt"
	"testing"

	"676f.dev/utilities/tools/hashing"
	"github.com/stretchr/testify/assert"
)

func TestStringMD5(t *testing.T) {
	type test struct {
		inputs string
		want   string
	}

	tests := []test{
		{inputs: "", want: "d41d8cd98f00b204e9800998ecf8427e"},
		{inputs: "test", want: "098f6bcd4621d373cade4e832627b4f6"},
		{inputs: fmt.Sprint("1234", "asdf", "5678"), want: "f6ea411e785b7a48206b25059ddd9755"},
		{inputs: "ilovegolang", want: "47e1be8963058e60557f7813482949bc"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, hashing.StringMD5(tc.inputs))
	}
}

func TestStringSHA256(t *testing.T) {
	type test struct {
		inputs string
		want   string
	}

	tests := []test{
		{inputs: "", want: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"},
		{inputs: "test", want: "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"},
		{inputs: fmt.Sprint("1234", "asdf", "5678"), want: "8763f6b8e647e264d41c1842e48a5be067df517d8b4ec12b4aaff06edfad6cb0"},
		{inputs: "ilovegolang", want: "d9d7946f91920533e7f082de00cd2697ce394f956a28ece26953536b1809a74f"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, hashing.StringSHA256(tc.inputs))
	}
}

func TestStringSHA512(t *testing.T) {
	type test struct {
		inputs string
		want   string
	}

	tests := []test{
		{inputs: "", want: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"},
		{inputs: "test", want: "ee26b0dd4af7e749aa1a8ee3c10ae9923f618980772e473f8819a5d4940e0db27ac185f8a0e1d5f84f88bc887fd67b143732c304cc5fa9ad8e6f57f50028a8ff"},
		{inputs: fmt.Sprint("1234", "asdf", "5678"), want: "24ab8480b8b50b506e04c31773f5a306f2443ab7a51d5c6f3599eda0e076476f51375e42bed61365de854c79d30ff94f6d7f834385a584c40e014c29df466a4c"},
		{inputs: "ilovegolang", want: "e51e2e5a43803a1b64c2d943d7daeef982a55e4569e5450e6dd801b8ade88014b4536b5a7b8bf99602a7c84a514933682e08688b7c870dcbc53e9de6422c6dbe"},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.want, hashing.StringSHA512(tc.inputs))
	}
}
