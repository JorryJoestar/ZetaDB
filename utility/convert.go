package utility

import (
	"encoding/binary"
	"errors"
	"math"
)

//convert one byte from a byte slice to CHAR(string)
//error if this byte out of range of ascii
//error if this byte slice length is not 1
func BytesToCHAR(bytes []byte) (string, error) {
	if len(bytes) != 1 {
		return "", errors.New("length of byte slice invalid")
	}
	if bytes[0] > 0b01111111 {
		return "", errors.New("out of range of ascii")
	}
	return string(bytes), nil
}

//convert a string whose length is 1 into a byte slice
//error if string length is not 1
func CHARToBytes(c string) ([]byte, error) {
	if len(c) != 1 {
		return nil, errors.New("length of string invalid")
	}
	return []byte(c), nil
}

//convert one byte to CHAR(string)
//error if this byte out of range of ascii
func ByteToCHAR(b byte) (string, error) {
	if b > 0b01111111 {
		return "", errors.New("out of range of ascii")
	}
	var bytes []byte
	bytes = append(bytes, b)
	return string(bytes), nil
}

//convert CHAR(string) to a byte
//error if string length is not 1
func CHARToByte(c string) (byte, error) {
	if len(c) != 1 {
		return 0, errors.New("length of string invalid")
	}
	bytes := []byte(c)
	return bytes[0], nil
}

//convert byte slice to VARCHAR(string)
//error if any one byte in byte slice out of range of ascii
func BytesToVARCHAR(bytes []byte) (string, error) {
	for _, b := range bytes {
		if b > 0b01111111 {
			return "", errors.New("out of range of ascii")
		}
	}
	return string(bytes), nil
}

//convert VARCHAR(string) to byte slice
//error if any one character of VARCHAR out of range of ascii
func VARCHARToBytes(s string) ([]byte, error) {
	bytes := []byte(s)
	for _, b := range bytes {
		if b > 0b01111111 {
			return nil, errors.New("out of range of ascii")
		}
	}
	return bytes, nil
}

//convert a byte to bool
//if byte is 0, return false, else return true
func ByteToBOOLEAN(b byte) bool {
	if b == 0b00000000 {
		return false
	} else {
		return true

	}
}

//convert bool type to a byte
func BOOLEANToByte(b bool) byte {
	if b {
		return 0b00000001
	} else {
		return 0b00000000
	}
}

//convert 4 bytes to an int32, little-endian
//error if byte slice length is not 4
func BytesToINT(bytes []byte) (int32, error) {
	if len(bytes) != 4 {
		return 0, errors.New("length of byte slice invalid")
	}
	i := int32(bytes[0]) + int32(bytes[1])<<8 + int32(bytes[2])<<16 + int32(bytes[3])<<24
	return i, nil
}

//convert an int32 to 4 bytes, little-endian
func INTToBytes(i int32) []byte {
	var bytes []byte
	bytes = append(bytes, byte(i))
	bytes = append(bytes, byte(i>>8))
	bytes = append(bytes, byte(i>>16))
	bytes = append(bytes, byte(i>>24))
	return bytes
}

//convert 4 bytes to an int32, little-endian
func BytesToInteger(bytes []byte) (int32, error) {
	return BytesToINT(bytes)
}

//convert an int32 to 4 bytes, little-endian
func IntegerToBytes(i int32) []byte {
	return INTToBytes(i)
}

//convert 2 bytes to a int16, little-endian
func BytesToSHORTINT(bytes []byte) (int16, error) {
	if len(bytes) != 2 {
		return 0, errors.New("length of byte slice invalid")
	}
	i := int16(bytes[0]) + int16(bytes[1])<<8
	return i, nil
}

//convert an int16 to 2 bytes, little-endian
func SHORTINTToBytes(i int16) []byte {
	var bytes []byte
	bytes = append(bytes, byte(i))
	bytes = append(bytes, byte(i>>8))
	return bytes
}

//convert 4 bytes to FLOAT(float32), little-endian
func BytesToFLOAT(bytes []byte) (float32, error) {
	if len(bytes) != 4 {
		return 0, errors.New("length of byte slice invalid")
	}
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits), nil
}

//convert FLOAT(float32) to 4 bytes, little-endian
func FLOATToBytes(f float32) []byte {
	bits := math.Float32bits(f)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

//convert 4 bytes to REAL(float32), little-endian
func BytesToREAL(bytes []byte) (float32, error) {
	return BytesToFLOAT(bytes)
}

//convert REAL(float32) to 4 bytes, little-endian
func REALToBytes(f float32) []byte {
	return FLOATToBytes(f)
}

//convert 8 bytes to DOUBLEPRECISION(float64), little-endian
//error if byte slice length is not 8
func BytesToDOUBLEPRECISION(bytes []byte) (float64, error) {
	if len(bytes) != 8 {
		return 0, errors.New("length of byte slice invalid")
	}
	bits := binary.LittleEndian.Uint64(bytes)
	return math.Float64frombits(bits), nil
}

//convert DOUBLEPRECISION(float64) to 8 bytes, little-endian
func DOUBLEPRECISIONToBytes(f float64) []byte {
	bits := math.Float64bits(f)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

//convert bytes to DECIMAL(float64), sign byte at head
func BytesToDECIMAL(bytes []byte, n int, d int) (float64, error) {
	if n < d || n < 1 || d < 0 {
		return 0, errors.New("n or d invalid")
	}

	//byteNum is the number of bytes to store n digits plus a sign
	byteNum := 1 //sign byte
	if n%2 == 0 {
		byteNum += n / 2
	} else {
		byteNum += n/2 + 1
	}
	if len(bytes) != byteNum {
		return 0, errors.New("length of byte slice invalid")
	}

	//identify sign(+/-), if signByte is 0, sign is +
	signByte := bytes[0]
	bytes = bytes[1:]
	negtive := true
	if signByte == 0 {
		negtive = false
	} else {
		negtive = true
	}

	var f float64 = 0
	for i, b := range bytes {
		nLow := int(b & 0b00001111)
		nHigh := int((b & 0b11110000) >> 4)
		if nLow >= 10 || nHigh >= 10 {
			return 0, errors.New("byte content invalid")
		}
		base := math.Pow(100, float64(i))
		f += base * float64(nLow+10*nHigh)
	}

	decimal := math.Pow(0.1, float64(d))
	f *= decimal

	if negtive {
		f = -f
	}

	return f, nil
}

//convert DECIMAL(float64) to byte slice
func DECIMALToBytes(f float64, n int, d int) ([]byte, error) {
	if n < d || n < 1 || d < 0 {
		return nil, errors.New("n or d invalid")
	}

	//byteNum is the number of bytes to store n digits plus a sign
	byteNum := 1 //sign byte
	if n%2 == 0 {
		byteNum += n / 2
	} else {
		byteNum += n/2 + 1
	}
	bytes := make([]byte, byteNum)

	//assign sign byte
	if f >= 0 {
		bytes[0] = 0b00000000
	} else {
		bytes[0] = 0b00000001
	}

	//fetch value
	for i := 0; i < d; i++ {
		f *= 10
	}
	if f < 0 {
		f = -f
	}
	var value int = int(f)

	for i := 1; i < byteNum; i++ {
		nLow := value % 10
		value = value / 10
		nHigh := value % 10
		value = value / 10
		b := (byte(nHigh) << 4) | byte(nLow)

		bytes[i] = b
	}

	return bytes, nil
}

//convert bytes to NUMERIC(float64)
func BytesToNUMERIC(bytes []byte, n int, d int) (float64, error) {
	return BytesToDECIMAL(bytes, n, d)
}

//convert NUMERIC(float64) to bytes
func NUMERICToBytes(f float64, n int, d int) ([]byte, error) {
	return DECIMALToBytes(f, n, d)
}
