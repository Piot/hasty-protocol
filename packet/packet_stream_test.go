package packet

import "testing"

func TestBuffer(t *testing.T) {
	stream := NewPacketStream()
	testOctets := []byte{0x02, 128}

	if stream.buffer.Len() != 0 {
		t.Errorf("Length is wrong %d", stream.buffer.Len())
	}

	secondStream, err := stream.Feed(testOctets)

	if err != nil {
		t.Errorf("Feed returned error:%s", err)
	}

	if secondStream.buffer.Len() != 2 {
		t.Errorf("Length is wrong %d", secondStream.buffer.Len())
	}

	if secondStream.octets()[0] != 0x02 {
		t.Errorf("Contents is wrong %c", secondStream.octets()[0])
	}

	newStream, newPacket, fetchErr := secondStream.FetchPacket()
	if fetchErr != nil {
		t.Errorf("Fetcherror:%s", fetchErr)
	}
	if newStream.buffer.Len() != len(testOctets) {
		t.Errorf("Octets left:%d", newStream.buffer.Len())
	}

	if newPacket.payload != nil {
		t.Errorf("Payload:%d", newPacket.payload)
	}
}

func TestBufferInChunks(t *testing.T) {
	stream := NewPacketStream()
	testOctets := []byte{1, 1}

	if stream.buffer.Len() != 0 {
		t.Errorf("Length is wrong %d", stream.buffer.Len())
	}

	secondStream, err := stream.Feed(testOctets)

	if err != nil {
		t.Errorf("Feed returned error")
	}

	if secondStream.buffer.Len() != len(testOctets) {
		t.Errorf("Length is wrong %d", secondStream.buffer.Len())
	}

	if secondStream.octets()[0] != 0x01 {
		t.Errorf("Contents is wrong %c", secondStream.octets()[0])
	}

	newStream, newPacket, fetchErr := secondStream.FetchPacket()
	if fetchErr != nil {
		t.Errorf("FetchErr:%s", fetchErr)
	}
	if newStream.buffer.Len() != len(testOctets) {
		t.Errorf("Octets left:%d", newStream.buffer.Len())
	}

	if newPacket.payload != nil {
		t.Errorf("Payload:%d", newPacket.payload)
	}

	testOctetsPartTwo := []byte{'h'}
	thirdStream, err2 := secondStream.Feed(testOctetsPartTwo)
	if err2 != nil {
		t.Errorf("streamErr:%s", err2)
	}

	_, foundPacket, fetchErr2 := thirdStream.FetchPacket()
	if fetchErr2 != nil {
		t.Errorf("fetchErr2:%s", fetchErr2)
	}
	if foundPacket.header.headerOctetSize != 2 {
		t.Errorf("Header octet size wrong:%d", foundPacket.header.headerOctetSize)
	}
	if foundPacket.header.packetType != PublishStream {
		t.Errorf("Not publish")
	}
	if foundPacket.header.payloadSize != 1 {
		t.Errorf("payload size wrong")
	}
	if foundPacket.payload[0] != testOctetsPartTwo[0] {
		t.Errorf("payload diff")
	}
}
