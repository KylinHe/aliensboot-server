// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: room_model.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 游戏结算分为1v1、2v2结算，data.resultDisplay用于客户端结算面板的显示，data.resultData用户后台上报。游戏分胜利，失败，平局，逃跑四种结果。
type GameResult int32

const (
	GameResult_Win    GameResult = 0
	GameResult_Lose   GameResult = 1
	GameResult_Equal  GameResult = 2
	GameResult_Escape GameResult = 3
)

var GameResult_name = map[int32]string{
	0: "Win",
	1: "Lose",
	2: "Equal",
	3: "Escape",
}
var GameResult_value = map[string]int32{
	"Win":    0,
	"Lose":   1,
	"Equal":  2,
	"Escape": 3,
}

func (x GameResult) String() string {
	return proto.EnumName(GameResult_name, int32(x))
}
func (GameResult) EnumDescriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{0} }

type PlayerResult struct {
	Playerid int64   `protobuf:"varint,1,opt,name=playerid,proto3" json:"playerid,omitempty"`
	Record   *Record `protobuf:"bytes,3,opt,name=record" json:"record,omitempty"`
}

func (m *PlayerResult) Reset()                    { *m = PlayerResult{} }
func (m *PlayerResult) String() string            { return proto.CompactTextString(m) }
func (*PlayerResult) ProtoMessage()               {}
func (*PlayerResult) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{0} }

func (m *PlayerResult) GetPlayerid() int64 {
	if m != nil {
		return m.Playerid
	}
	return 0
}

func (m *PlayerResult) GetRecord() *Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type Record struct {
	Result GameResult `protobuf:"varint,1,opt,name=result,proto3,enum=protocol.GameResult" json:"result,omitempty"`
	Score  int32      `protobuf:"varint,2,opt,name=score,proto3" json:"score,omitempty"`
	Unit   string     `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
}

func (m *Record) Reset()                    { *m = Record{} }
func (m *Record) String() string            { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()               {}
func (*Record) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{1} }

func (m *Record) GetResult() GameResult {
	if m != nil {
		return m.Result
	}
	return GameResult_Win
}

func (m *Record) GetScore() int32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Record) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

// 玩家加入座位请求
type JoinRequest struct {
	AppID    string `protobuf:"bytes,1,opt,name=appID,proto3" json:"appID,omitempty"`
	RoomID   string `protobuf:"bytes,2,opt,name=roomID,proto3" json:"roomID,omitempty"`
	PlayerID int64  `protobuf:"varint,3,opt,name=playerID,proto3" json:"playerID,omitempty"`
	SeatID   int32  `protobuf:"varint,4,opt,name=seatID,proto3" json:"seatID,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{2} }

func (m *JoinRequest) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *JoinRequest) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *JoinRequest) GetPlayerID() int64 {
	if m != nil {
		return m.PlayerID
	}
	return 0
}

func (m *JoinRequest) GetSeatID() int32 {
	if m != nil {
		return m.SeatID
	}
	return 0
}

type Room struct {
	RoomID string  `protobuf:"bytes,1,opt,name=roomID,proto3" json:"roomID,omitempty"`
	Mode   int32   `protobuf:"varint,2,opt,name=mode,proto3" json:"mode,omitempty"`
	Seats  []*Seat `protobuf:"bytes,3,rep,name=seats" json:"seats,omitempty"`
}

func (m *Room) Reset()                    { *m = Room{} }
func (m *Room) String() string            { return proto.CompactTextString(m) }
func (*Room) ProtoMessage()               {}
func (*Room) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{3} }

func (m *Room) GetRoomID() string {
	if m != nil {
		return m.RoomID
	}
	return ""
}

func (m *Room) GetMode() int32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Room) GetSeats() []*Seat {
	if m != nil {
		return m.Seats
	}
	return nil
}

type Seat struct {
	Id     int32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Lock   bool    `protobuf:"varint,2,opt,name=lock,proto3" json:"lock,omitempty"`
	Player *Player `protobuf:"bytes,3,opt,name=player" json:"player,omitempty"`
}

func (m *Seat) Reset()                    { *m = Seat{} }
func (m *Seat) String() string            { return proto.CompactTextString(m) }
func (*Seat) ProtoMessage()               {}
func (*Seat) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{4} }

func (m *Seat) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Seat) GetLock() bool {
	if m != nil {
		return m.Lock
	}
	return false
}

func (m *Seat) GetPlayer() *Player {
	if m != nil {
		return m.Player
	}
	return nil
}

type Player struct {
	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Group    int32  `protobuf:"varint,2,opt,name=group,proto3" json:"group,omitempty"`
	Seat     int32  `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
	Nickname string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Avatar   string `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender   string `protobuf:"bytes,6,opt,name=gender,proto3" json:"gender,omitempty"`
	Role     int32  `protobuf:"varint,7,opt,name=role,proto3" json:"role,omitempty"`
}

func (m *Player) Reset()                    { *m = Player{} }
func (m *Player) String() string            { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()               {}
func (*Player) Descriptor() ([]byte, []int) { return fileDescriptorRoomModel, []int{5} }

func (m *Player) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Player) GetGroup() int32 {
	if m != nil {
		return m.Group
	}
	return 0
}

func (m *Player) GetSeat() int32 {
	if m != nil {
		return m.Seat
	}
	return 0
}

func (m *Player) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *Player) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *Player) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *Player) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func init() {
	proto.RegisterType((*PlayerResult)(nil), "protocol.PlayerResult")
	proto.RegisterType((*Record)(nil), "protocol.Record")
	proto.RegisterType((*JoinRequest)(nil), "protocol.JoinRequest")
	proto.RegisterType((*Room)(nil), "protocol.Room")
	proto.RegisterType((*Seat)(nil), "protocol.Seat")
	proto.RegisterType((*Player)(nil), "protocol.Player")
	proto.RegisterEnum("protocol.GameResult", GameResult_name, GameResult_value)
}
func (m *PlayerResult) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PlayerResult) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Playerid != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Playerid))
	}
	if m.Record != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Record.Size()))
		n1, err := m.Record.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *Record) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Record) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Result != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Result))
	}
	if m.Score != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Score))
	}
	if len(m.Unit) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Unit)))
		i += copy(dAtA[i:], m.Unit)
	}
	return i, nil
}

func (m *JoinRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JoinRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.AppID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.AppID)))
		i += copy(dAtA[i:], m.AppID)
	}
	if len(m.RoomID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.RoomID)))
		i += copy(dAtA[i:], m.RoomID)
	}
	if m.PlayerID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.PlayerID))
	}
	if m.SeatID != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.SeatID))
	}
	return i, nil
}

func (m *Room) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Room) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RoomID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.RoomID)))
		i += copy(dAtA[i:], m.RoomID)
	}
	if m.Mode != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Mode))
	}
	if len(m.Seats) > 0 {
		for _, msg := range m.Seats {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintRoomModel(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *Seat) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Seat) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Id))
	}
	if m.Lock {
		dAtA[i] = 0x10
		i++
		if m.Lock {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Player != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Player.Size()))
		n2, err := m.Player.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *Player) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Player) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Id))
	}
	if m.Group != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Group))
	}
	if m.Seat != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Seat))
	}
	if len(m.Nickname) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Nickname)))
		i += copy(dAtA[i:], m.Nickname)
	}
	if len(m.Avatar) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Avatar)))
		i += copy(dAtA[i:], m.Avatar)
	}
	if len(m.Gender) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(len(m.Gender)))
		i += copy(dAtA[i:], m.Gender)
	}
	if m.Role != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintRoomModel(dAtA, i, uint64(m.Role))
	}
	return i, nil
}

func encodeVarintRoomModel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *PlayerResult) Size() (n int) {
	var l int
	_ = l
	if m.Playerid != 0 {
		n += 1 + sovRoomModel(uint64(m.Playerid))
	}
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovRoomModel(uint64(l))
	}
	return n
}

func (m *Record) Size() (n int) {
	var l int
	_ = l
	if m.Result != 0 {
		n += 1 + sovRoomModel(uint64(m.Result))
	}
	if m.Score != 0 {
		n += 1 + sovRoomModel(uint64(m.Score))
	}
	l = len(m.Unit)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	return n
}

func (m *JoinRequest) Size() (n int) {
	var l int
	_ = l
	l = len(m.AppID)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	l = len(m.RoomID)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	if m.PlayerID != 0 {
		n += 1 + sovRoomModel(uint64(m.PlayerID))
	}
	if m.SeatID != 0 {
		n += 1 + sovRoomModel(uint64(m.SeatID))
	}
	return n
}

func (m *Room) Size() (n int) {
	var l int
	_ = l
	l = len(m.RoomID)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	if m.Mode != 0 {
		n += 1 + sovRoomModel(uint64(m.Mode))
	}
	if len(m.Seats) > 0 {
		for _, e := range m.Seats {
			l = e.Size()
			n += 1 + l + sovRoomModel(uint64(l))
		}
	}
	return n
}

func (m *Seat) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRoomModel(uint64(m.Id))
	}
	if m.Lock {
		n += 2
	}
	if m.Player != nil {
		l = m.Player.Size()
		n += 1 + l + sovRoomModel(uint64(l))
	}
	return n
}

func (m *Player) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRoomModel(uint64(m.Id))
	}
	if m.Group != 0 {
		n += 1 + sovRoomModel(uint64(m.Group))
	}
	if m.Seat != 0 {
		n += 1 + sovRoomModel(uint64(m.Seat))
	}
	l = len(m.Nickname)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	l = len(m.Gender)
	if l > 0 {
		n += 1 + l + sovRoomModel(uint64(l))
	}
	if m.Role != 0 {
		n += 1 + sovRoomModel(uint64(m.Role))
	}
	return n
}

func sovRoomModel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozRoomModel(x uint64) (n int) {
	return sovRoomModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PlayerResult) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PlayerResult: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PlayerResult: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Playerid", wireType)
			}
			m.Playerid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Playerid |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Record == nil {
				m.Record = &Record{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Record) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Record: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Record: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Result", wireType)
			}
			m.Result = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Result |= (GameResult(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Score", wireType)
			}
			m.Score = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Score |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unit", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unit = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JoinRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JoinRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JoinRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PlayerID", wireType)
			}
			m.PlayerID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PlayerID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SeatID", wireType)
			}
			m.SeatID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SeatID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Room) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Room: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Room: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoomID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RoomID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mode", wireType)
			}
			m.Mode = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Mode |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seats = append(m.Seats, &Seat{})
			if err := m.Seats[len(m.Seats)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Seat) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Seat: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Seat: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lock", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lock = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Player == nil {
				m.Player = &Player{}
			}
			if err := m.Player.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Player) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Player: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Player: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Group", wireType)
			}
			m.Group = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Group |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seat", wireType)
			}
			m.Seat = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seat |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nickname", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nickname = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthRoomModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRoomModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRoomModel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRoomModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoomModel
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRoomModel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthRoomModel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowRoomModel
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipRoomModel(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthRoomModel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoomModel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("room_model.proto", fileDescriptorRoomModel) }

var fileDescriptorRoomModel = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0xdf, 0xaa, 0xd3, 0x40,
	0x10, 0xc6, 0xdd, 0xe6, 0xcf, 0x69, 0xa6, 0x87, 0x12, 0x96, 0x22, 0xc1, 0x8b, 0x52, 0x82, 0x17,
	0x41, 0xa4, 0x17, 0x15, 0x7c, 0x00, 0xe9, 0x41, 0x2a, 0x5e, 0xc8, 0x7a, 0x40, 0xef, 0x74, 0x4d,
	0x86, 0x43, 0x38, 0x49, 0x26, 0x67, 0x93, 0x08, 0xbe, 0x8e, 0x4f, 0xe3, 0xa5, 0x8f, 0x20, 0x7d,
	0x12, 0xd9, 0xd9, 0x3d, 0x6d, 0xc5, 0xab, 0xce, 0x37, 0xdb, 0xf9, 0xcd, 0xcc, 0x37, 0x81, 0xd4,
	0x10, 0xb5, 0x5f, 0x5a, 0xaa, 0xb0, 0xd9, 0xf6, 0x86, 0x46, 0x92, 0x73, 0xfe, 0x29, 0xa9, 0xc9,
	0x6f, 0xe1, 0xfa, 0x43, 0xa3, 0x7f, 0xa0, 0x51, 0x38, 0x4c, 0xcd, 0x28, 0x9f, 0xc1, 0xbc, 0x67,
	0x5d, 0x57, 0x99, 0xd8, 0x88, 0x22, 0x50, 0x27, 0x2d, 0x0b, 0x88, 0x0d, 0x96, 0x64, 0xaa, 0x2c,
	0xd8, 0x88, 0x62, 0xb1, 0x4b, 0xb7, 0x8f, 0x98, 0xad, 0xe2, 0xbc, 0xf2, 0xef, 0xf9, 0x57, 0x88,
	0x5d, 0x46, 0xbe, 0xb4, 0x35, 0x96, 0xcc, 0xb4, 0xe5, 0x6e, 0x75, 0xae, 0x79, 0xab, 0x5b, 0x74,
	0x5d, 0x95, 0xff, 0x8f, 0x5c, 0x41, 0x34, 0x94, 0x64, 0x30, 0x9b, 0x6d, 0x44, 0x11, 0x29, 0x27,
	0xa4, 0x84, 0x70, 0xea, 0xea, 0x91, 0xbb, 0x26, 0x8a, 0xe3, 0x9c, 0x60, 0xf1, 0x8e, 0xea, 0x4e,
	0xe1, 0xc3, 0x84, 0x03, 0x17, 0xea, 0xbe, 0x3f, 0xec, 0xb9, 0x4b, 0xa2, 0x9c, 0x90, 0x4f, 0x21,
	0xb6, 0xab, 0x1f, 0xf6, 0xcc, 0x4b, 0x94, 0x57, 0xe7, 0x25, 0x0f, 0x7b, 0x86, 0x9e, 0x96, 0x74,
	0x35, 0x03, 0xea, 0xf1, 0xb0, 0xcf, 0x42, 0x9e, 0xc1, 0xab, 0xfc, 0x33, 0x84, 0x8a, 0xa8, 0xbd,
	0x60, 0x8a, 0x7f, 0x98, 0x12, 0x42, 0xeb, 0xb0, 0x9f, 0x9c, 0x63, 0xf9, 0x1c, 0x22, 0x5b, 0x3d,
	0x64, 0xc1, 0x26, 0x28, 0x16, 0xbb, 0xe5, 0x79, 0xf7, 0x8f, 0xa8, 0x47, 0xe5, 0x1e, 0xf3, 0x5b,
	0x08, 0xad, 0x94, 0x4b, 0x98, 0x79, 0xd3, 0x23, 0x35, 0xab, 0x2b, 0x4b, 0x6c, 0xa8, 0xbc, 0x67,
	0xe2, 0x5c, 0x71, 0x6c, 0x4f, 0xe0, 0x26, 0xfd, 0xff, 0x04, 0xfe, 0x8c, 0xfe, 0x3d, 0xff, 0x29,
	0x20, 0x76, 0xa9, 0x0b, 0x70, 0xc0, 0xe0, 0x15, 0x44, 0x77, 0x86, 0xa6, 0xfe, 0xd1, 0x65, 0x16,
	0xb6, 0x9d, 0x9d, 0x87, 0xc1, 0x91, 0xe2, 0xd8, 0x1a, 0xd5, 0xd5, 0xe5, 0x7d, 0xa7, 0x5b, 0x64,
	0x3b, 0x12, 0x75, 0xd2, 0xd6, 0x08, 0xfd, 0x5d, 0x8f, 0xda, 0x64, 0x91, 0x33, 0xc2, 0x29, 0x9b,
	0xbf, 0xc3, 0xae, 0x42, 0x93, 0xc5, 0x2e, 0xef, 0x94, 0xe5, 0x1b, 0x6a, 0x30, 0xbb, 0x72, 0x7c,
	0x1b, 0xbf, 0x78, 0x0d, 0x70, 0xfe, 0x0a, 0xe4, 0x15, 0x04, 0x9f, 0xea, 0x2e, 0x7d, 0x22, 0xe7,
	0x10, 0xbe, 0xa7, 0x01, 0x53, 0x21, 0x13, 0x88, 0x6e, 0x1e, 0x26, 0xdd, 0xa4, 0x33, 0x09, 0x10,
	0xdf, 0x0c, 0xa5, 0xee, 0x31, 0x0d, 0xde, 0x5c, 0xff, 0x3a, 0xae, 0xc5, 0xef, 0xe3, 0x5a, 0xfc,
	0x39, 0xae, 0xc5, 0xb7, 0x98, 0x3d, 0x78, 0xf5, 0x37, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xa4, 0x7e,
	0x3b, 0xe8, 0x02, 0x00, 0x00,
}