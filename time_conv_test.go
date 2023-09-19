package time_conv

import "testing"

func TestTimeconv(t *testing.T) {
	test_data := []string{"1", "days", "to", "hours"}
	want := 24
	got, err := Convert(&test_data)

	if got != want || err != nil {
		t.Errorf("Convert() = %q, want %q", got, want)
	}
}

func TestUnknownInstruction(t *testing.T) {
	test_data := []string{"1", "days", "to", "nanoseconds"}
	_, err := Convert(&test_data)

	if err == nil {
		t.Errorf("Convert() not detected uknown time unit")
	}
}
