package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientDependency_ParseUrl(t *testing.T) {
	dependencyWithSprintf := &ClientDependency{
		baseUrl: "https://something.com/%s/dummy-%s-%s-%s",
		name:    "dummy",
	}
	dependencyWithOutSprintf := &ClientDependency{
		baseUrl: "https://something.com/",
		name:    "dummy",
	}
	tagName := "v1.0.0"
	t.Run("should parse macos flag", func(t *testing.T) {
		systemOs = macos
		systemArch = "amd64"
		assert.Equal(
			t,
			"https://something.com/v1.0.0/dummy-v1.0.0-darwin-amd64",
			dependencyWithSprintf.ParseUrl(tagName),
		)
		assert.Equal(
			t,
			"https://something.com/",
			dependencyWithOutSprintf.ParseUrl(tagName),
		)
	})
	t.Run("should work without flag flag", func(t *testing.T) {
		systemOs = ubuntu
		systemArch = "amd64"
		assert.Equal(
			t,
			"https://something.com/v1.0.0/dummy-v1.0.0-linux-amd64",
			dependencyWithSprintf.ParseUrl(tagName),
		)
		assert.Equal(
			t,
			"https://something.com/",
			dependencyWithOutSprintf.ParseUrl(tagName),
		)
	})
	t.Run("should work without flag flag", func(t *testing.T) {
		systemOs = windows
		systemArch = "amd64"
		assert.Equal(
			t,
			"https://something.com/v1.0.0/dummy-v1.0.0-windows-amd64",
			dependencyWithSprintf.ParseUrl(tagName),
		)
		assert.Equal(
			t,
			"https://something.com/",
			dependencyWithOutSprintf.ParseUrl(tagName),
		)
	})
	t.Run("should parse ubuntu as a fallback with no parse", func(t *testing.T) {
		systemOs = ""
		assert.Equal(
			t,
			"https://something.com/",
			dependencyWithOutSprintf.ParseUrl(tagName),
		)
	})
}
