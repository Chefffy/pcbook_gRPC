package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func WriteProtobufToJSONFile(message proto.Message,filename string)error{
	data,err := ProtobufToJSON(message)
	if err != nil{
		return fmt.Errorf("can not marshal proto message to JSON: %w",err)
	}

	err = ioutil.WriteFile(filename,[]byte(data),0644)
	if err != nil{
		return fmt.Errorf("can not write JSON data to file: %w",err)
	}

	return nil
}

func WriteProtobufToBinaryFile(message proto.Message,filename string)error{
	data,err := proto.Marshal(message)
	if err != nil{
		return fmt.Errorf("can not marshal proto message to binary: %w",err)
	}

	err = ioutil.WriteFile(filename,data,0644)
	if err != nil{
		return fmt.Errorf("can not wtire binary data to file: %w",err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string,message proto.Message)error{
	data,err := ioutil.ReadFile(filename)
	if err != nil{
		return fmt.Errorf("can not read binary data from file: %w",err)
	}

	err = proto.Unmarshal(data,message)
	if err != nil{
		return fmt.Errorf("can not unmarshal binary to proto message: %w",err)
	}

	return nil
}