package serializer_test

import (
	"testing"

	"example.com/laptop_store/serializer"
	"example.com/laptop_store/messages"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)


func TestFileSerializer(t *testing.T){

	file := "../tmp/laptop.bin"
	fileJson := "../tmp/laptop.json"
	laptop1 := serializer.NewLaptop()

	err := serializer.WriteProtobufToBinary(laptop1, file)

	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinary(file, laptop2)

	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	err = serializer.WriteProtobufToJSON(laptop1, fileJson)
	require.NoError(t, err)

}