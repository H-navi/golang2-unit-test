package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	//befor
	fmt.Println("Before unit test")

	m.Run()

	//after
	fmt.Println("After unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Tono")
	if result != "Hello Tono" {
		t.FailNow()
	}
	fmt.Println("TestHelloWorld Done")
}

func TestHelloWorldJean(t *testing.T) {
	result := HelloWorld("Jean")
	if result != "Hello Jean" {
		t.Fail()
	}
	fmt.Println("TestHelloWorldJean Done")
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("Jean")
	if result != "Hello Jean" {
		t.Error("Should be Hello Jean")
	}
	fmt.Println("Executed if fail")
}

func TestHelloWorldJeanFatal(t *testing.T) {
	result := HelloWorld("Jean")
	if result != "Hello Jean2" {
		t.Fatal("Should be Hello Jean")
	}
	fmt.Println("Not executed if fail")
}

func TestHelloWorldAssertion(t *testing.T) {
	result := HelloWorld("Jean")
	assert.Equal(t, "Hello Jean", result, "Should be Hello Jean")
	fmt.Println("Executed")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Jean")
	require.Equal(t, "Hello Jean", result, "Should be Hello Jean")
	fmt.Println("Not Executed")
}

func TestSkip(t *testing.T) {
	//check OS
	if runtime.GOOS == "windows" {
		t.Skip("This unit test isn't for windows")
	}

	result := HelloWorld("Tono")
	require.Equal(t, "Hello Tono", result, "Should be Hello Tono")
}

func TestSubTest(t *testing.T) {
	t.Run("Tono", func(t *testing.T) {
		result := HelloWorld("Tono")
		require.Equal(t, "Hello Tono", result)
	})
	t.Run("Jean", func(t *testing.T) {
		result := HelloWorld("Jean")
		require.Equal(t, "Hello Jean", result)
	})
}
