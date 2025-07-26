package di

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	// Act
	container := New()

	// Assert
	assert.NotNil(t, container, "New should return a non-nil container")
	assert.NotNil(t, container.container, "New should initialize the underlying container")
}

func TestContainer_Provide(t *testing.T) {
	// Arrange
	container := New()

	// Act & Assert (no panic)
	assert.NotPanics(t, func() {
		container.Provide(func() string { return "test" })
	}, "Provide should not panic with valid constructor")
}

func TestContainer_Provide_Panic(t *testing.T) {
	// Arrange
	container := New()

	// Act & Assert (should panic)
	assert.Panics(t, func() {
		// Providing the same dependency twice causes a panic
		container.Provide(func() string { return "test" })
		container.Provide(func() string { return "test" })
	}, "Provide should panic when there's an error")
}

func TestContainer_Invoke(t *testing.T) {
	// Arrange
	container := New()
	testValue := "test value"
	container.Provide(func() string { return testValue })

	// Act
	var result string
	err := container.Invoke(func(s string) {
		result = s
	})

	// Assert
	require.NoError(t, err, "Invoke should not return an error")
	assert.Equal(t, testValue, result, "Invoke should pass the provided dependency to the function")
}

func TestContainer_Invoke_Error(t *testing.T) {
	// Arrange
	container := New()

	// Act
	err := container.Invoke(func(_ string) {
		// This function requires a string dependency that hasn't been provided
	})

	// Assert
	assert.Error(t, err, "Invoke should return an error when dependencies are missing")
}

func TestNamed(t *testing.T) {
	// Arrange
	container := New()

	// Act & Assert (no panic)
	assert.NotPanics(t, func() {
		container.Provide(func() string { return "test" }, Named("test_name"))
	}, "Named should return a valid ProvideOption")

	// Verify the named dependency can be retrieved
	var result string
	err := container.Invoke(func(in struct {
		In

		Value string `name:"test_name"`
	},
	) {
		result = in.Value
	})

	require.NoError(t, err, "Invoke should not return an error")
	assert.Equal(t, "test", result, "Named dependency should be retrievable")
}
