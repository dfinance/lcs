package lcs

import "math/big"

// Convert BgiInt to bytes with little endian (required by Libra).
func BigToBytes(val *big.Int, bytesLen int) []byte {
	bytes := val.Bytes()

	if len(bytes) < bytesLen {
		diff := bytesLen - len(bytes)
		for i := 0; i < diff; i++ {
			bytes = append([]byte{0}, bytes...)
		}
	}

	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-i-1] = bytes[len(bytes)-i-1], bytes[i]
	}

	return bytes
}

// Convert little endian bytes to big endian bytes and then to sdk.Int.
func LeToBig(bytes []byte) *big.Int {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[len(bytes)-i-1], bytes[i] = bytes[i], bytes[len(bytes)-i-1]
	}

	bigValue := &big.Int{}
	bigValue.SetBytes(bytes)

	return bigValue
}
