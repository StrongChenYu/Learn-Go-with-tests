package main

import "testing"

func TestHello(t *testing.T)  {
	assertCorrectMessage := func(t testing.TB, actual, expected string) {
		//如果没有这个，就不会报告具体的行号，这会导致很难找到问题
		t.Helper()
		if actual != expected {
			t.Errorf("expected: %q, but actual: %q", expected, actual)
		}
	}

	t.Run("say hello to people in chinese", func(t *testing.T) {
		actual := Hello("chenyu", "Chinese")
		expect := "你好, chenyu"
		assertCorrectMessage(t, actual, expect)
	})

	t.Run("say hello to people in spanish", func(t *testing.T) {
		actual := Hello("", "Spanish")
		expect := "Hola, world"
		assertCorrectMessage(t, actual, expect)
	})

	t.Run("say hello to people in default", func(t *testing.T) {
		actual := Hello("chenyu", "")
		expect := "Hello, chenyu"
		assertCorrectMessage(t, actual, expect)
	})

	//got := Hello()
	//expect := "hello world"
	//
	//if got != expect {
	//	t.Errorf("got:%q but expect: %q", got, expect)
	//}
}
