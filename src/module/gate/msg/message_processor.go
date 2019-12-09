/*******************************************************************************
 * Copyright (c) 2015, 2018 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2018/3/30
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package msg

import (
	"encoding/binary"
	"errors"
	"github.com/KylinHe/aliensboot-core/chanrpc"
	"github.com/KylinHe/aliensboot-core/protocol/base"
)

var Processor = NewMsgProcessor() //protobuf.NewProcessor()

type MessageProcessor struct {
	littleEndian bool
	msgRouter    *chanrpc.Server
}

func NewMsgProcessor() *MessageProcessor {
	return &MessageProcessor{}
}

// It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *MessageProcessor) SetByteOrder(littleEndian bool) {
	p.littleEndian = littleEndian
}

//func (this *MessageProcessor) Route(msg interface{}, userData interface{}) error {
//	this.msgRouter.Go(reflect.TypeOf(&base.Any{}), msg, userData)
//	return nil
//}

// must goroutine safe
func (this *MessageProcessor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < 2 {
		return nil, errors.New("data too short")
	}

	var id uint16 = 0
	var seqId uint32 = 0
	if this.littleEndian {
		id = binary.LittleEndian.Uint16(data)
		seqId = binary.LittleEndian.Uint32(data[2:])
	} else {
		id = binary.BigEndian.Uint16(data)
		seqId = binary.BigEndian.Uint32(data[2:])
	}
	//log.Debugf("marshal %v - %v", id, data)
	return &base.Any{Id: id, SeqId: seqId, Value: data[6:]}, nil
}

// must goroutine safe
func (this *MessageProcessor) Marshal(msg interface{}) ([][]byte, error) {
	any, ok := msg.(*base.Any)
	if !ok || any == nil {
		return nil, errors.New("invalid any type")
	}
	id := make([]byte, 2)
	seqId := make([]byte, 4)
	if this.littleEndian {
		binary.LittleEndian.PutUint16(id, any.Id)
		binary.LittleEndian.PutUint32(seqId, any.SeqId)
	} else {
		binary.BigEndian.PutUint16(id, any.Id)
		binary.BigEndian.PutUint32(seqId, any.SeqId)
	}
	//log.Debugf("marshal %v - %v", any.Id, id)
	return [][]byte{id, seqId, any.Value}, nil
}
