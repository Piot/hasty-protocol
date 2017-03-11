package opath

import "testing"

func checkPath(t *testing.T, path string, valid bool) {
	_, err := NewFromString(path)
	if (err == nil) && !valid {
		t.Errorf("'%s' is not a valid path and should fail", path)
	} else if (err != nil) && valid {
		t.Errorf("Failed '%s' should be a valid path", path)
	}
}

func TestPath(t *testing.T) {
	checkPath(t, "/zaphod", true)
	checkPath(t, "/zaphod", true)
	checkPath(t, "/32/42numbersallowedatstart", true)
	checkPath(t, "/_underscore_is_allowed", true)
	checkPath(t, "/@zaphod", true)
	checkPath(t, "/zaphod/@1234", true)

	checkPath(t, "/endslashnotallowed/", false)
	checkPath(t, "/something/åaphod", false)
	checkPath(t, "/23#@aa/23", false)
	checkPath(t, "musthaveroot", false)
	checkPath(t, "/root/endslashnotallowed/", false)
	checkPath(t, "/endslashnotallowed_/", false)
	checkPath(t, "*/", false)
	checkPath(t, "/z aphod", false)
	checkPath(t, " /zaphod", false)
}
