// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package fb

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type BlockOffset struct {
	_tab flatbuffers.Table
}

func GetRootAsBlockOffset(buf []byte, offset flatbuffers.UOffsetT) *BlockOffset {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &BlockOffset{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *BlockOffset) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *BlockOffset) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *BlockOffset) Key(j int) int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt8(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *BlockOffset) KeyLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *BlockOffset) MutateKey(j int, n int8) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt8(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *BlockOffset) Offset() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BlockOffset) MutateOffset(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *BlockOffset) Len() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *BlockOffset) MutateLen(n uint32) bool {
	return rcv._tab.MutateUint32Slot(8, n)
}

func BlockOffsetStart(builder *flatbuffers.Builder) {
	builder.StartObject(3)
}
func BlockOffsetAddKey(builder *flatbuffers.Builder, key flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(key), 0)
}
func BlockOffsetStartKeyVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func BlockOffsetAddOffset(builder *flatbuffers.Builder, offset uint32) {
	builder.PrependUint32Slot(1, offset, 0)
}
func BlockOffsetAddLen(builder *flatbuffers.Builder, len uint32) {
	builder.PrependUint32Slot(2, len, 0)
}
func BlockOffsetEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}