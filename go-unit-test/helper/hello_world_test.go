package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	arg := "Rocket"
	expect := "Hello " + arg
	result := HelloWorld(arg)
	if result != expect {
		panic("Result is not " + expect)
	}
}

// // Fail Test
// func TestFail(t *testing.T) {
// 	arg := "Rocket"
// 	expect := "Hello 2 " + arg
// 	result := HelloWorld(arg)
// 	if result != expect {
// 		// error
// 		t.Fail()
// 	}
// 	fmt.Println("TestFail Done")
// }

// func TestFailNow(t *testing.T) {
// 	arg := "Rocket"
// 	expect := "Hello 3" + arg
// 	result := HelloWorld(arg)
// 	if result != expect {
// 		// error
// 		t.FailNow()
// 	}
// 	fmt.Println("TestFailNow Done")
// }

// func TestError(t *testing.T) {
// 	arg := "Rocket"
// 	expect := "Hello 4 " + arg
// 	result := HelloWorld(arg)
// 	if result != expect {
// 		// error
// 		t.Error("Result must be " + expect)
// 	}
// 	fmt.Println("TestError Done")
// }

// func TestFatal(t *testing.T) {
// 	arg := "Rocket"
// 	expect := "Hello 5" + arg
// 	result := HelloWorld(arg)
// 	if result != expect {
// 		// error
// 		t.Fatal("Result must be " + expect)
// 	}
// 	fmt.Println("TestFatal Done")
// }

// Assertion
func TestAssert(t *testing.T) {
	arg := "Rocket"
	expect := "Hello " + arg
	result := HelloWorld(arg)

	// Success case
	assert.Equal(t, expect, result, "Result mmust be "+expect)

	// // Fail case
	// assert.Equal(t, expect, "Hello Assert Fail", "Result must be "+expect)

	fmt.Println("TestAssert Done")
}

func TestRequire(t *testing.T) {
	arg := "Rocket"
	expect := "Hello " + arg
	result := HelloWorld(arg)

	// Success case
	require.Equal(t, expect, result, "Result mmust be "+expect)

	// Fail case
	require.Equal(t, expect, "Hello Require Fail", "Result must be "+expect)

	fmt.Println("TestRequire Done")
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	arg := "Rocket"
	expect := "Hello " + arg
	result := HelloWorld(arg)

	// Success case
	assert.Equal(t, expect, result, "Result must be "+expect)

	fmt.Println("TestSkip Done")
}

// Before dan After Test
func TestMain(m *testing.M) {
	// before UT
	fmt.Println("Before UT")

	m.Run()

	// after UT
	fmt.Println("After UT")

	fmt.Println("Before After Test Done")
}

// Sub Test
func TestSubTest(t *testing.T) {

	// t.Run(SubTestName, func(t *testing.T)  { UT })
	t.Run("SubUT1", func(t *testing.T) {
		arg := "Rocket"
		expect := "Hello " + arg
		result := HelloWorld(arg)

		assert.Equal(t, expect, result, "Result must be "+expect)
	})

	t.Run("SubUT2", func(t *testing.T) {
		arg := "Racoon"
		expect := "Hello " + arg
		result := HelloWorld(arg)

		assert.Equal(t, expect, result, "Result must be "+expect)
	})

	fmt.Println("TestSubTest Done")
}

// Table Test
// ideal test: sub + table
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(Rocket)",
			request:  "Rocket",
			expected: "Hello Rocket",
		},
		{
			name:     "HelloWorld(Racoon)",
			request:  "Racoon",
			expected: "Hello Racoon",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}


// Benchmark
// benchmark not running when there's fail UT
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Rocket")
	}
}

func BenchmarkHelloWorld2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Racoon")
	}
}

// Sub Benchmark
// ideal for multi bench
func BenchmarkSub(b *testing.B) {
	b.Run("Sub Rocket", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sub Rocket")
		}
	})

	b.Run("Sub Racoon", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Sub Racoon")
		}
	})
}

// Benchmark Table
func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name     string
		request  string
	}{
		{
			name:     "HelloWorld(Rocket)",
			request:  "Rocket",
		},
		{
			name:     "HelloWorld(Racoon)",
			request:  "Racoon",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}
