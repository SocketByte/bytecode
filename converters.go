package bytecode

import (
    "encoding/binary"
    "math"
    "strings"
)

func Float32ToBinary(value float32) []byte {
    bytes := make([]byte, 4)
    bits := math.Float32bits(value)
    binary.BigEndian.PutUint32(bytes, bits)

    return bytes
}

func Float64ToBinary(value float64) []byte {
    bytes := make([]byte, 8)
    bits := math.Float64bits(value)
    binary.BigEndian.PutUint64(bytes, bits)

    return bytes
}

func Int16ToBinary(value uint16) []byte {
    return SInt16ToBinary(int16(value))
}

func SInt16ToBinary(value int16) []byte {
    bytes := make([]byte, 2)
    bytes[0] = byte(value >> 8)
    bytes[1] = byte(value)
    return bytes
}

func Int32ToBinary(value uint32) []byte {
    return SInt32ToBinary(int(value))
}

func SInt32ToBinary(value int) []byte {
    bytes := make([]byte, 4)
    bytes[0] = byte(value >> 24)
    bytes[1] = byte(value >> 16)
    bytes[2] = byte(value >> 8)
    bytes[3] = byte(value)

    return bytes
}

func Int64ToBinary(value uint64) []byte {
    return SInt64ToBinary(int64(value))
}

func SInt64ToBinary(value int64) []byte {
    bytes := make([]byte, 8)
    bytes[0] = byte(value >> 56)
    bytes[1] = byte(value >> 48)
    bytes[2] = byte(value >> 40)
    bytes[3] = byte(value >> 32)
    bytes[4] = byte(value >> 24)
    bytes[5] = byte(value >> 16)
    bytes[6] = byte(value >> 8)
    bytes[7] = byte(value)

    return bytes
}

func DescriptorToStackSize(descriptor string) []uint16 {
    beginIndex := strings.LastIndex(descriptor, "(")
    endIndex := strings.LastIndex(descriptor, ")")

    sizes := make([]uint16, 2)
    for i := beginIndex; i < endIndex; i++ {
        char := []rune(descriptor)[i]

        if char == '[' {
            continue
        }

        if char == 'L' {
            sizes[0] += 1

            end := strings.Index(descriptor, ";")

            i = end
            if i > endIndex {
                break
            }
            continue
        }

        sizes[0] += 1
    }
    sizes[0] -= 1
    sub := strings.Split(descriptor, ")")[1]
    if sub == "V" {
        return sizes
    }
    sizes[1] = 1
    return sizes
}
