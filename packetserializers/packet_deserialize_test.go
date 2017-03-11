package packetserializers

import (
	"log"
	"testing"

	"github.com/piot/hasty-protocol/base64packet"
	"github.com/piot/hasty-protocol/commands"
)

type DummyHandler struct {
	publish commands.Publish
}

func (t *DummyHandler) HandlePublish(p commands.Publish) {
	log.Println("Dummy Handling:", p)
	t.publish = p
}

func (t *DummyHandler) HandleSubscribe(p commands.Subscribe) {
	log.Println("Dummy Handlingx:", p)
}

func (t *DummyHandler) HandleUnsubscribe(p commands.Unsubscribe) {
	log.Println("Dummy Handlingy:", p)
}

func TestDeserializeBuffer(t *testing.T) {
	const packetBase64 = "AUwQL2hlbGxvL3dvcmxkLzEyMzsAAAADc3ViACcAAAABc3ViZmllbGQAmpmZmZmZKEACc3VicwAGAAAAaGVsbG8AABB0ZW1wAAwAAAAA"
	packet, err := base64packet.Base64ToPacket(packetBase64)
	if err != nil {
		t.Errorf("Err:%s", err)
	}

	log.Println("packet:", packet)
	handler := DummyHandler{}
	Deserialize(packet, &handler)
	if handler.publish.Path().ToString() != "/hello/world/123" {
		t.Errorf("Path is wrong:%s", handler.publish)
	}
}

func TestDeserializeRandomGarbageBuffer(t *testing.T) {
	const packetBase64 = "yq/cWFTIQuiJKafyMDl189uPMt2vrnnXVAktoz0Nqch7AAsAJVIo0By4PS5mWpSO+lu74Oe7RhALvIz4yICIqu8wY35te/cvlx3gJyov"
	_, err := base64packet.Base64ToPacket(packetBase64)
	if err != nil {
		log.Println("base64 was not a packet, good")
		return
	}
	t.Errorf("It shouldn't be accepted as a valid packet")
}
