/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved.
 * Date:
 *     2019/12/7
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package protocol

import (
	"errors"
	"github.com/gogo/protobuf/proto"
)

type MsgProcessor struct {

}

func (*MsgProcessor) NewResponseData() (interface{}, error) {
	return &Response{}, nil
}

func (*MsgProcessor) Decode(buf []byte) (interface{}, error) {
	data := &Request{}
	error := proto.Unmarshal(buf, data)
	return  data, error
}

func (*MsgProcessor) Encode(data interface{}) ([]byte, error) {
	result, ok := data.(*Response)
	if !ok {
		return nil, errors.New("invalid data *protocol.Response")
	}
	return proto.Marshal(result)
}
