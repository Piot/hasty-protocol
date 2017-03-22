package asciistring

import "testing"

func TestFromTooSmallBuffer(t *testing.T) {
	octets := []byte{}
	decoded, totalOctetLength, err := FromOctets(octets)
	if err == nil {
		t.Error("Should have reported an error", err)
	}
	if totalOctetLength != 0 {
		t.Errorf("octet length wrong:%d", totalOctetLength)
	}

	if decoded != "" {
		t.Errorf("answer:'%s'", decoded)
	}
}

func TestFromOctets(t *testing.T) {
	octets := []byte{3, 'h', 'i', '!'}
	decoded, totalOctetLength, err := FromOctets(octets)
	if decoded != "hi!" {
		t.Errorf("answer:'%s'", decoded)
	}
	if totalOctetLength != 4 {
		t.Errorf("octet length wrong:%d", totalOctetLength)
	}

	if err != nil {
		t.Errorf("Should not report error")
	}
}

func TestFromOctetsBeyondBuffer(t *testing.T) {
	octets := []byte{4, 'h', 'i', '!'}
	_, _, err := FromOctets(octets)
	if err == nil {
		t.Errorf("Error: buffer is too small")
	}
}

func TestFromOctetsExtendedBuffer(t *testing.T) {
	octets := []byte{2, 'h', 'i', '!'}
	decoded, totalOctetLength, err := FromOctets(octets)
	if decoded != "hi" {
		t.Errorf("Decoded was not hi:%s", decoded)
	}

	if totalOctetLength != 3 {
		t.Errorf("Reported wrong length:%d", totalOctetLength)
	}

	if err != nil {
		t.Errorf("Should not report error")
	}
}

func TestToOctets(t *testing.T) {
	encoded, err := ToOctets("Hello, this is a test")
	if err != nil {
		t.Errorf("Unexpected error")
	}
	if len(encoded) != 22 {
		t.Errorf("Length is wrong:%d", len(encoded))
	}
	encodedLength := encoded[0]
	if encodedLength != 21 {
		t.Errorf("Encoded length wrong:%d", encodedLength)
	}

	if encoded[21] != 't' {
		t.Errorf("End was wrong")
	}
}

func TestToAndFromOctets(t *testing.T) {
	const x = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	encoded, errTo := ToOctets(x)
	const expectedOctetSize = len(x) + 2
	if len(encoded) != expectedOctetSize {
		t.Errorf("Length was wrong:%d", len(encoded))
	}
	if errTo != nil {
		t.Errorf("Error:%s", errTo)
	}
	decoded, totalOctetLength, errFrom := FromOctets(encoded)
	if totalOctetLength != expectedOctetSize {
		t.Errorf("Octetlength:%d", totalOctetLength)
	}
	if decoded != x {
		t.Errorf("Decoded wrong:%s", decoded)
	}
	if errFrom != nil {
		t.Errorf("error:%s", errFrom)
	}
}
