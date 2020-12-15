package general

import "testing"

func TestTimelib_Format(t *testing.T) {
	str := NewTime().Format()
	t.Log(str)
}
