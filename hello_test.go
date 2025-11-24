package hello

import "testing"

// Execute with: go test ./...
func TestSayHello(t *testing.T) {
	// want := "Hello, test!"
	// // got := Say("test")
	// got := Say([]string{"test"})

	// if want != got {
	// 	t.Errorf("Wanted %s, got %s", want, got)
	// }

	subtests := []struct{
		items []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items: []string{"Tony"},
			result: "Hello, Tony!",
		},
		{
			items: []string{"Alice", "Bob"},
			result: "Hello, Alice, Bob!",
		},
	}

	for _, subtest := range subtests {
		if value := Say(subtest.items); value != subtest.result {
			t.Errorf("Wanted %s (%v), got %s", subtest.result, subtest.items, value)
		}
	}
}

