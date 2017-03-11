package packet

import "testing"

func TestCompletePacketHeaderWithout(t *testing.T) {
	octets := []byte{2, 129, 1}
	packet, packetIsDone, err := CheckIfWeHavePacketHeader(octets)
	if err != nil {
		t.Errorf("Packet was in error")
	}
	if !packetIsDone {
		t.Errorf("Packet is not done")
	}
	if packet.payloadSize != 1 {
		t.Errorf("Payload size is wrong:%d", packet.payloadSize)
	}

	if packet.packetType != PublishStream {
		t.Errorf("Packet type is not nop:%d", packet.packetType)
	}
}

func TestIncompletePacketHeader(t *testing.T) {
	octets := []byte{0x02, 128}
	packet, packetIsDone, err := CheckIfWeHavePacketHeader(octets)
	if err != nil {
		t.Errorf("error:%s", err)
	}
	if packetIsDone {
		t.Errorf("Packet should not be done")
	}
	if packet.payloadSize != 0 {
		t.Errorf("Packet should be zero")
	}
}

func TestCompletePacketHeaderSmallSize(t *testing.T) {
	octets := []byte{1, 128}
	packet, packetIsDone, err := CheckIfWeHavePacketHeader(octets)
	if err != nil {
		t.Errorf("Error:%s", err)
	}
	if !packetIsDone {
		t.Errorf("Packet header should be done")
	}
	if packet.payloadSize != 0 {
		t.Errorf("Packet should be zero")
	}
}
