package packetserializers

import (
	"fmt"

	"github.com/piot/hasty-protocol/qos"
)

func convertFromOctetToQoSConst(t byte) (qos.Enum, error) {
	qosType := qos.Enum(t)
	if qosType < qos.Normal || qosType > qos.Prioritized {
		return qos.Normal, fmt.Errorf("Illegal qos type:%v", t)
	}

	return qosType, nil
}

// ToQos : Create Qos
func ToQos(in uint8) (out qos.QoS, err error) {
	qosEnum, err := convertFromOctetToQoSConst(in)
	if err != nil {
		return qos.NewQoS(qos.Normal), err
	}
	qos := qos.NewQoS(qosEnum)
	return qos, nil
}
