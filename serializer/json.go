package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) ([]byte, error) {
	// Create a protojson marshaller
	marshaller := protojson.MarshalOptions{
		Multiline:      true,
		Indent:         "  ",
		UseEnumNumbers: false,
	}

	// Marshal the proto message to JSON
	return marshaller.Marshal(message)
}
