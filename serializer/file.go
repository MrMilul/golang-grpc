package serializer

import(
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToBinary(message proto.Message, filename string) error{
	data, err := proto.Marshal(message)
	if err != nil{
		return fmt.Errorf("Cannot convert data to binary")
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil{
		return fmt.Errorf("Cannot write binary data to file")
	}
	return nil
}

func ReadProtobufFromBinary(filename string, message proto.Message) error{
	data, err := ioutil.ReadFile(filename)
	if err != nil{
		return fmt.Errorf("Cannot read data to binary")
	}
	err = proto.Unmarshal(data, message)
	if err != nil{
		return fmt.Errorf("Cannot convert binary to message")
	}

	return nil

}

func WriteProtobufToJSON(message proto.Message, filename string)error{
	jData, err := ProtobufToJSON(message)
	if err != nil{
		return fmt.Errorf("Cannot convert message to JSON")
	}

	err = ioutil.WriteFile(filename,  []byte(jData), 0644)

	if err != nil{
		return fmt.Errorf("Cannot convert binary to message")
	}
	return nil
}
