package marsh

import (
	"bytes"
	"encoding/binary"
)

const (
	maxStringLength uint64 = 0x000000000000FFFF
)

func (fielder *FieldObject) Marshal() []byte {
	/*
		Index	4		0x0000		0x0003
		Random	8		0x0004		0x000C
		Count	2		0x000D		0x000E
		Array	?		0x000F		0x????
		** Array - each entry contains a string and is laid out as follows **
		Length	2		0x0000		0x0001
		String	N		0x0002		0x0002+N
		nil		1		0x0002+N+1
	*/
	var expectedLength int = 12
	for _, str := range fielder.Attributes {
		expectedLength += len(str) + 1 + 2
	}

	var offset uint16 = 0
	result := make([]byte, expectedLength)

	// 32-bit index
	binary.BigEndian.PutUint32(result[offset:], fielder.Index)
	offset += 4

	// 64-bit Random Number
	binary.BigEndian.PutUint64(result[offset:], fielder.Random)
	offset += 8

	// 16-bit counter for the number of attributes
	binary.BigEndian.PutUint16(result[offset:], uint16(len(fielder.Attributes)))
	offset += 2

	for _, attrib := range fielder.Attributes {
		// 16-bit counter for the length of the string
		binary.BigEndian.PutUint16(result[offset:], uint16(len(attrib)))
		offset += 2

		// convert the string to a byte array
		byteAttrib := []byte(attrib)

		// limit the string length to the maximum allowed
		var maxStringData uint64 = uint64(len(byteAttrib))
		if maxStringLength < maxStringData {
			maxStringData = maxStringLength
		}

		// N-byte string
		for index := uint64(0); index < maxStringData; index++ {
			result[offset] = byteAttrib[index]

			offset++
		}

		// adjust the pointer by the length of the string and an additional one for one for null termination
		offset += uint16(len(attrib)) + uint16(1)
	}

	// return the resulting buffer
	return result
}

func Unmarshal(data []byte) *FieldObject {
	fielder := &FieldObject{}

	var offset int = 0
	// first 4 byte
	fielder.Index = binary.BigEndian.Uint32(data[offset : offset+4])
	offset += 4
	// next 8 bytes
	fielder.Random = binary.BigEndian.Uint64(data[offset : offset+8])
	offset += 8
	// next 2 bytes
	attributeCount := binary.BigEndian.Uint16(data[offset : offset+2])

	for index := uint16(0); index < attributeCount; index++ {
		// next 2 bytes - length
		attributeDataLength := binary.BigEndian.Uint16(data[offset : offset+2])
		offset += 2

		// extract the attribute string
		data := data[offset : offset+int(attributeDataLength)]
		// add the length and an extra for the null terminator
		offset += int(attributeDataLength) + 1

		// convert the data to a string and save it
		strData := string(data)
		fielder.AddAttribute(strData)
	}

	return fielder
}

func (fielder *FieldObject) StreamMarshal() ([]byte, error) {
	result := new(bytes.Buffer)
	err := binary.Write(result, binary.BigEndian, fielder)
	if err != nil {
		return result.Bytes(), err
	} else {
		return nil, err
	}
}

func (fielder *FieldObject) StreamUnmarshal(data []byte) (*FieldObject, error) {
	fielder = &FieldObject{}

	buffer := bytes.NewBuffer(data)

	err := binary.Read(buffer, binary.BigEndian, fielder)
	return fielder, err
}
