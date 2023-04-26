package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/jsonpb"
)

func ProtobufToJSON(message proto.Message) (string, error){
	marshaler := jsonpb.Marshaler{
		EnumsAsInts: false,
		EmitDefaults: true,
		Indent: "  ",
		OrigName: true,
	}

	return marshaler.MarshalToString(message)
}