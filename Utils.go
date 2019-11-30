package packet

import (
	"bytes"
	"encoding/binary"
)

// Int64FromBytes converts a Big Endian representation to an int64.
func Int64FromBytes(b []byte) (int64, error) {
	buffer := bytes.NewReader(b)
	var result int64
	err := binary.Read(buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// Int64ToBytes converts an int64 to its Big Endian representation.
func Int64ToBytes(i int64) ([]byte, error) {
	buffer := bytes.Buffer{}
	err := binary.Write(&buffer, binary.BigEndian, i)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}

// Int32FromBytes converts a Big Endian representation to an int32.
func Int32FromBytes(b []byte) (int32, error) {
	buffer := bytes.NewReader(b)
	var result int32
	err := binary.Read(buffer, binary.BigEndian, &result)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// Int32ToBytes converts an int32 to its Big Endian representation.
func Int32ToBytes(i int32) ([]byte, error) {
	buffer := bytes.Buffer{}
	err := binary.Write(&buffer, binary.BigEndian, i)

	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
