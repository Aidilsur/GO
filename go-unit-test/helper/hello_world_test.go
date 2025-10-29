package helper

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Aidil")
	}
}


func BenchmarkHelloWorldSurya(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Surya")
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Aidil", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
		HelloWorld("Aidil")
	}
	})
	b.Run("Surya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
		HelloWorld("Surya")
	}
	})
}

func TestMain(m *testing.M) {
	fmt.Println("Before Unit Test")

	m.Run()

	fmt.Println("After Unit Test")
}

// func TestSkip(t *testing.T) {
// 	if runtime.GOOS == "windows" {
// 		t.Skip("can not run on window")
// 	}

// 	result := HelloWorld("Aidil")
// 	assert.Equal(t, "Hello Aidil", result, "result must be 'Hello Aidil'")
// 	fmt.Println("Dieksekusi")
// }

func TestHelloWord1(t *testing.T) {
	result := HelloWorld("Aidil")

	if result != "Hello Aidil" {
		// unit test failed
		// t.Fail()
		t.Error("Result must be Hello")
	} 

		fmt.Println("test done")
}

func TestHelloWord2(t *testing.T) {
	result := HelloWorld("Aidil")

	if result != "Hello Aidil" {
		// unit test failed
		// t.FailNow()
		t.Fatal("Result must be Hello")
	} 

	fmt.Println("test done")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Aidil")
	assert.Equal(t, "Hello Aidil", result, "result must be 'Hello Aidil'")
	fmt.Println("Dieksekusi")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Aidil")
	require.Equal(t, "Hello Aidil", result, "result must be 'Hello Aidil'")
	fmt.Println("Dieksekusi")
}

func TestSubTest(t *testing.T) {
	t.Run("Aidil", func(t *testing.T) {
		result := HelloWorld("Aidil")
		require.Equal(t, "Hello Aidil", result, "result must be 'Hello Aidil'")
	})
	t.Run("Surya", func(t *testing.T) {
		result := HelloWorld("Surya")
		require.Equal(t, "Hello Surya", result, "result must be 'Hello Surya'")
	})
}

type helloWorldTestCase struct {
	name string
	request string
	expected string
}


func TestTableHelloWorld(t *testing.T) {
	tests := []helloWorldTestCase{
		{
			name: "Hello Aidil(Aidil)",
			request: "Aidil",
			expected: "Hello Aidil",
		},
		{
			name: "Hello Surya(Surya)",
			request: "Surya",
			expected: "Hello Surya",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := HelloWorld(tc.request)
			require.Equal(t, tc.expected, result)
		})
	}
}

func BenchmarkTable(b *testing.B) {
	tests := []helloWorldTestCase{
		{
			name: "Hello Aidil(Aidil)",
			request: "Aidil",
			expected: "Hello Aidil",
		},
		{
			name: "Hello Surya(Surya)",
			request: "Surya",
			expected: "Hello Surya",
		},
	}

	for _, tb := range tests {
		b.Run(tb.name, func(b *testing.B) {
			HelloWorld(tb.request)
		})
	}
}

