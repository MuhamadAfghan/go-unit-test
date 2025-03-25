package helper

import (
	"runtime"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorldTable(b *testing.B) {
	tests := []struct {
		name    string
		request string
	}{
		{
			name:    "World",
			request: "World",
		},
		{
			name:    "Golang",
			request: "Golang",
		},
		{
			name:    "Muhamad Afghan Alzena",
			request: "Muhamad Afghan Alzena",
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(test.request)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("SubTest1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("World")
		}
	})

	b.Run("SubTest2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Golang")
		}
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test1",
			input:    "World",
			expected: "Hello World!",
		},
		{
			name:     "Test2",
			input:    "Golang",
			expected: "Hello Golang!",
		},
		{
			name:     "Test3",
			input:    "Test",
			expected: "Hello Test!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.input)
			assert.Equal(t, test.expected, result, "they should be equal")
			fmt.Println("Test passed:", test.name)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("SubTest1", func(t *testing.T) {
		fmt.Println("SubTest1")
	})

	t.Run("SubTest2", func(t *testing.T) {
		fmt.Println("SubTest2")
	})
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")

	m.Run()

	//after
	fmt.Println("After Unit Test")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Skipping this test on Windows")
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("World")

	require.Equal(t, "Hello World!", result, "they should be equal")
	// require.Equal(t, "Hello Worldd!", result, "they should be equal") // this will fail the test

	fmt.Println("TestHelloWorld passed")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("World")

	assert.Equal(t, "Hello World!", result, "they should be equal")
	// assert.Equal(t, "Hello Worldd!", result, "they should be equal") // this will fail the test

	fmt.Println("TestHelloWorldAssert passed")
}
