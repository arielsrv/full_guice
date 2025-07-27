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
	}, "Register should not panic with valid constructor")
}

func TestContainer_Provide_Panic(t *testing.T) {
	// Arrange
	container := New()

	// Act & Assert (should panic)
	assert.Panics(t, func() {
		// Providing the same dependency twice causes a panic
		container.Provide(func() string { return "test" })
		container.Provide(func() string { return "test" })
	}, "Register should panic when there's an error")
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

func TestGet(t *testing.T) {
	// Arrange
	Reset() // Reset global container
	testValue := "test value"
	Container.Provide(func() string { return testValue })

	// Act
	result := GetInstance[string]()

	// Assert
	assert.Equal(t, testValue, result, "GetInstance should return the provided dependency")
}

func TestGet_Panic(t *testing.T) {
	// Arrange
	Reset() // Reset global container to ensure no dependencies

	// Act & Assert (should panic)
	assert.Panics(t, func() {
		GetInstance[string]() // No string dependency provided
	}, "GetInstance should panic when dependency is missing")
}

func TestGet_ComplexType(t *testing.T) {
	// Arrange
	type TestStruct struct {
		Value string
	}

	Reset() // Reset global container
	testStruct := &TestStruct{Value: "complex test"}
	Container.Provide(func() *TestStruct { return testStruct })

	// Act
	result := GetInstance[*TestStruct]()

	// Assert
	assert.Equal(t, testStruct, result, "GetInstance should work with complex types")
	assert.Equal(t, "complex test", result.Value, "GetInstance should preserve struct values")
}

func TestGetNamed_NotImplemented(t *testing.T) {
	// Arrange
	container := New()
	container.Provide(func() string { return "test" }, Named("test_name"))

	// Act & Assert (should panic with specific message)
	assert.PanicsWithValue(t, "GetNamed is not yet implemented. Use specific getter functions for named dependencies or regular GetInstance[T]() for non-named ones.", func() {
		GetNamed[string](container, "test_name")
	}, "GetNamed should panic with not implemented message")
}
