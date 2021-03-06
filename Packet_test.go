package packet_test

import (
	"encoding/hex"
	"github.com/aerogo/packet"
	"github.com/akyoto/assert"
	"testing"
)

func TestPacketValid(t *testing.T) {
	builder := packet.NewPacketBuilder()

	err := builder.AddBool(true)
	assert.Nil(t, err)
	err = builder.AddByte(16)
	assert.Nil(t, err)
	err = builder.AddRune('別')
	assert.Nil(t, err)
	err = builder.AddInt8(-32)
	assert.Nil(t, err)
	err = builder.AddInt16(-23536)
	assert.Nil(t, err)
	err = builder.AddInt32(-2045968102)
	assert.Nil(t, err)
	err = builder.AddInt64(-8932684869)
	assert.Nil(t, err)
	err = builder.AddUint8(248)
	assert.Nil(t, err)
	err = builder.AddUint16(65532)
	assert.Nil(t, err)
	err = builder.AddUint32(4294967294)
	assert.Nil(t, err)
	err = builder.AddUint64(8294967294)
	assert.Nil(t, err)
	err = builder.AddFloat32(2.396)
	assert.Nil(t, err)
	err = builder.AddFloat64(1282387.3748)
	assert.Nil(t, err)
	err = builder.AddComplex64(64.113 + 11.28i)
	assert.Nil(t, err)
	err = builder.AddComplex128(34875634875.2349587 + 543957.23459843i)
	assert.Nil(t, err)
	err = builder.AddString("Fuck you, you little prick.")
	assert.Nil(t, err)
	err = builder.AddBytes([]byte("私はお尻であなたをファックしたい"))

	pkt := builder.BuildPacket(238)
	assert.Equal(t, pkt.Opcode, int32(238))
	assert.Equal(t, pkt.Length, int64(155))

	decomposer := packet.NewPacketDecomposer(pkt)

	boolVal, err := decomposer.ReadBool()
	assert.Nil(t, err)
	assert.Equal(t, boolVal, true)
	byteVal, err := decomposer.ReadByte()
	assert.Nil(t, err)
	assert.Equal(t, byteVal, byte(16))
	runeVal, err := decomposer.ReadRune()
	assert.Nil(t, err)
	assert.Equal(t, runeVal, '別')
	int8Val, err := decomposer.ReadInt8()
	assert.Nil(t, err)
	assert.Equal(t, int8Val, int8(-32))
	int16Val, err := decomposer.ReadInt16()
	assert.Nil(t, err)
	assert.Equal(t, int16Val, int16(-23536))
	int32Val, err := decomposer.ReadInt32()
	assert.Nil(t, err)
	assert.Equal(t, int32Val, int32(-2045968102))
	int64Val, err := decomposer.ReadInt64()
	assert.Nil(t, err)
	assert.Equal(t, int64Val, int64(-8932684869))
	uint8Val, err := decomposer.ReadUint8()
	assert.Nil(t, err)
	assert.Equal(t, uint8Val, uint8(248))
	uint16Val, err := decomposer.ReadUint16()
	assert.Nil(t, err)
	assert.Equal(t, uint16Val, uint16(65532))
	uint32Val, err := decomposer.ReadUint32()
	assert.Nil(t, err)
	assert.Equal(t, uint32Val, uint32(4294967294))
	uint64Val, err := decomposer.ReadUint64()
	assert.Nil(t, err)
	assert.Equal(t, uint64Val, uint64(8294967294))
	float32Val, err := decomposer.ReadFloat32()
	assert.Nil(t, err)
	assert.Equal(t, float32Val, float32(2.396))
	float64Val, err := decomposer.ReadFloat64()
	assert.Nil(t, err)
	assert.Equal(t, float64Val, float64(1282387.3748))
	complex64Val, err := decomposer.ReadComplex64()
	assert.Nil(t, err)
	assert.Equal(t, complex64Val, complex64(64.113+11.28i))
	complex128Val, err := decomposer.ReadComplex128()
	assert.Nil(t, err)
	assert.Equal(t, complex128Val, complex128(34875634875.2349587+543957.23459843i))
	strVal, err := decomposer.ReadString()
	assert.Nil(t, err)
	assert.Equal(t, strVal, "Fuck you, you little prick.")
	bytesVal, err := decomposer.ReadBytes()
	assert.Nil(t, err)
	assert.Equal(t, string(bytesVal), "私はお尻であなたをファックしたい")

	bytes, err := pkt.Bytes()
	assert.Nil(t, err)
	t.Log(hex.Dump(bytes))
}
