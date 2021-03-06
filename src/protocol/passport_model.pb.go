// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: passport_model.proto

package protocol

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 服务端不允许login文件名存在,特改为bblogin
// 登录相关通讯协议
type LoginResult int32

const (
	LoginResult_loginSuccess    LoginResult = 0
	LoginResult_invalidUser     LoginResult = 1
	LoginResult_invalidPwd      LoginResult = 2
	LoginResult_forbiddenUser   LoginResult = 3
	LoginResult_tokenExpire     LoginResult = 4
	LoginResult_invalidMaintain LoginResult = 5
)

var LoginResult_name = map[int32]string{
	0: "loginSuccess",
	1: "invalidUser",
	2: "invalidPwd",
	3: "forbiddenUser",
	4: "tokenExpire",
	5: "invalidMaintain",
}
var LoginResult_value = map[string]int32{
	"loginSuccess":    0,
	"invalidUser":     1,
	"invalidPwd":      2,
	"forbiddenUser":   3,
	"tokenExpire":     4,
	"invalidMaintain": 5,
}

func (x LoginResult) String() string {
	return proto.EnumName(LoginResult_name, int32(x))
}
func (LoginResult) EnumDescriptor() ([]byte, []int) { return fileDescriptorPassportModel, []int{0} }

type RegisterResult int32

const (
	RegisterResult_registerSuccess RegisterResult = 0
	RegisterResult_userExists      RegisterResult = 1
	RegisterResult_invalidFormat   RegisterResult = 2
)

var RegisterResult_name = map[int32]string{
	0: "registerSuccess",
	1: "userExists",
	2: "invalidFormat",
}
var RegisterResult_value = map[string]int32{
	"registerSuccess": 0,
	"userExists":      1,
	"invalidFormat":   2,
}

func (x RegisterResult) String() string {
	return proto.EnumName(RegisterResult_name, int32(x))
}
func (RegisterResult) EnumDescriptor() ([]byte, []int) { return fileDescriptorPassportModel, []int{1} }

type User struct {
	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id" gorm:"AUTO_INCREMENT"`
	Username   string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" bson:"username" unique:"true"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty" bson:"password"`
	Salt       string `protobuf:"bytes,4,opt,name=salt,proto3" json:"salt,omitempty" bson:"salt"`
	Channeluid string `protobuf:"bytes,5,opt,name=channeluid,proto3" json:"channeluid,omitempty" bson:"cuid"`
	Channel    string `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty" bson:"channel"`
	Avatar     string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty" bson:"avatar"`
	Mobile     string `protobuf:"bytes,8,opt,name=mobile,proto3" json:"mobile,omitempty" bson:"mobile"  rorm:"-"`
	Openid     string `protobuf:"bytes,9,opt,name=openid,proto3" json:"openid,omitempty" bson:"openid"`
	Status     int32  `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty" bson:"status"`
	RegTime    int64  `protobuf:"varint,11,opt,name=regTime,proto3" json:"regTime,omitempty" bson:"regtime"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptorPassportModel, []int{0} }

func (m *User) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetSalt() string {
	if m != nil {
		return m.Salt
	}
	return ""
}

func (m *User) GetChanneluid() string {
	if m != nil {
		return m.Channeluid
	}
	return ""
}

func (m *User) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetOpenid() string {
	if m != nil {
		return m.Openid
	}
	return ""
}

func (m *User) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *User) GetRegTime() int64 {
	if m != nil {
		return m.RegTime
	}
	return 0
}

func init() {
	proto.RegisterType((*User)(nil), "protocol.User")
	proto.RegisterEnum("protocol.LoginResult", LoginResult_name, LoginResult_value)
	proto.RegisterEnum("protocol.RegisterResult", RegisterResult_name, RegisterResult_value)
}
func (m *User) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *User) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(m.Id))
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.Salt) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Salt)))
		i += copy(dAtA[i:], m.Salt)
	}
	if len(m.Channeluid) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Channeluid)))
		i += copy(dAtA[i:], m.Channeluid)
	}
	if len(m.Channel) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Channel)))
		i += copy(dAtA[i:], m.Channel)
	}
	if len(m.Avatar) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Avatar)))
		i += copy(dAtA[i:], m.Avatar)
	}
	if len(m.Mobile) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Mobile)))
		i += copy(dAtA[i:], m.Mobile)
	}
	if len(m.Openid) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(len(m.Openid)))
		i += copy(dAtA[i:], m.Openid)
	}
	if m.Status != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(m.Status))
	}
	if m.RegTime != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintPassportModel(dAtA, i, uint64(m.RegTime))
	}
	return i, nil
}

func encodeVarintPassportModel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *User) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPassportModel(uint64(m.Id))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Salt)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Channeluid)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Avatar)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Mobile)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	l = len(m.Openid)
	if l > 0 {
		n += 1 + l + sovPassportModel(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovPassportModel(uint64(m.Status))
	}
	if m.RegTime != 0 {
		n += 1 + sovPassportModel(uint64(m.RegTime))
	}
	return n
}

func sovPassportModel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozPassportModel(x uint64) (n int) {
	return sovPassportModel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *User) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPassportModel
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
			return fmt.Errorf("proto: User: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: User: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Salt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Salt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channeluid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channeluid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Avatar", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Avatar = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Mobile", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Mobile = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Openid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
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
				return ErrInvalidLengthPassportModel
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Openid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RegTime", wireType)
			}
			m.RegTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPassportModel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RegTime |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPassportModel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthPassportModel
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
func skipPassportModel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPassportModel
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
					return 0, ErrIntOverflowPassportModel
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
					return 0, ErrIntOverflowPassportModel
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
				return 0, ErrInvalidLengthPassportModel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowPassportModel
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
				next, err := skipPassportModel(dAtA[start:])
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
	ErrInvalidLengthPassportModel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPassportModel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("passport_model.proto", fileDescriptorPassportModel) }

var fileDescriptorPassportModel = []byte{
	// 535 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xd1, 0x6e, 0xd3, 0x3e,
	0x14, 0xc6, 0xff, 0xe9, 0xba, 0xae, 0x3b, 0xdd, 0xda, 0xcc, 0xfb, 0x4b, 0x44, 0x20, 0x96, 0x60,
	0xb8, 0x18, 0x13, 0x5b, 0x2f, 0xc6, 0xd5, 0x24, 0x2e, 0x28, 0x2a, 0x02, 0x89, 0x0d, 0x64, 0xba,
	0xeb, 0xca, 0x89, 0xbd, 0xcc, 0x22, 0xb1, 0x8b, 0xed, 0x6c, 0xe3, 0x3d, 0x78, 0x28, 0x2e, 0x79,
	0x82, 0x08, 0xf5, 0x11, 0xf2, 0x04, 0x28, 0x4e, 0x3a, 0xaa, 0x5d, 0x25, 0xe7, 0xfb, 0x7e, 0xdf,
	0x71, 0x4e, 0x7c, 0xe0, 0xff, 0x05, 0x35, 0x66, 0xa1, 0xb4, 0x9d, 0xe7, 0x8a, 0xf1, 0xec, 0x64,
	0xa1, 0x95, 0x55, 0xa8, 0xef, 0x1e, 0x89, 0xca, 0x1e, 0x1f, 0xa7, 0xc2, 0x5e, 0x17, 0xf1, 0x49,
	0xa2, 0xf2, 0x71, 0xaa, 0x52, 0x35, 0x76, 0x4e, 0x5c, 0x5c, 0xb9, 0xca, 0x15, 0xee, 0xad, 0x09,
	0xe2, 0x9f, 0x5d, 0xe8, 0x5e, 0x1a, 0xae, 0xd1, 0x6b, 0xe8, 0x08, 0x16, 0x78, 0x91, 0x77, 0xb8,
	0x31, 0x79, 0x51, 0x95, 0x61, 0x14, 0x1b, 0x25, 0xcf, 0xf0, 0x5c, 0x30, 0x1c, 0xa5, 0x4a, 0xe7,
	0x67, 0xf8, 0xed, 0xe5, 0xec, 0xf3, 0xfc, 0xe3, 0xc5, 0x3b, 0x32, 0x3d, 0x9f, 0x5e, 0xcc, 0x30,
	0xe9, 0x08, 0x86, 0xde, 0x40, 0xbf, 0x30, 0x5c, 0x4b, 0x9a, 0xf3, 0xa0, 0x13, 0x79, 0x87, 0xdb,
	0x93, 0x67, 0x55, 0x19, 0x3e, 0x6d, 0xb2, 0x2b, 0x07, 0x47, 0x85, 0x14, 0xdf, 0x0b, 0x7e, 0x86,
	0xad, 0x2e, 0x38, 0x26, 0xf7, 0x11, 0x34, 0x86, 0x7e, 0x3d, 0xce, 0xad, 0xd2, 0x2c, 0xd8, 0x70,
	0xf1, 0xfd, 0xaa, 0x0c, 0x47, 0x4d, 0x7c, 0xe5, 0x60, 0x72, 0x0f, 0xa1, 0xe7, 0xd0, 0x35, 0x34,
	0xb3, 0x41, 0xd7, 0xc1, 0xa3, 0xaa, 0x0c, 0x07, 0x0d, 0x5c, 0xab, 0x98, 0x38, 0x13, 0x8d, 0x01,
	0x92, 0x6b, 0x2a, 0x25, 0xcf, 0x0a, 0xc1, 0x82, 0xcd, 0x87, 0x68, 0x52, 0x08, 0x86, 0xc9, 0x1a,
	0x82, 0x5e, 0xc1, 0x56, 0x5b, 0x05, 0x3d, 0x47, 0xa3, 0xaa, 0x0c, 0x87, 0x2d, 0xdd, 0x18, 0x98,
	0xac, 0x10, 0xf4, 0x12, 0x7a, 0xf4, 0x86, 0x5a, 0xaa, 0x83, 0x2d, 0x07, 0xef, 0x55, 0x65, 0xb8,
	0xdb, 0xc0, 0x8d, 0x8e, 0x49, 0x0b, 0xa0, 0x53, 0xe8, 0xe5, 0x2a, 0x16, 0x19, 0x0f, 0xfa, 0x0e,
	0x7d, 0x52, 0x95, 0xe1, 0xa3, 0x06, 0x6d, 0x74, 0x1c, 0x45, 0xda, 0xfd, 0xdc, 0x63, 0x4c, 0x5a,
	0xb4, 0xee, 0xaf, 0x16, 0x5c, 0x0a, 0x16, 0x6c, 0x3f, 0xec, 0xdf, 0xe8, 0x98, 0xb4, 0x40, 0x8d,
	0x1a, 0x4b, 0x6d, 0x61, 0x02, 0x88, 0xbc, 0xc3, 0xcd, 0x75, 0xb4, 0xd1, 0x31, 0x69, 0x81, 0x7a,
	0x46, 0xcd, 0xd3, 0x99, 0xc8, 0x79, 0x30, 0x70, 0x97, 0xbc, 0x36, 0xa3, 0xe6, 0xa9, 0x15, 0x39,
	0xc7, 0x64, 0x85, 0x1c, 0xfd, 0x80, 0xc1, 0x27, 0x95, 0x0a, 0x49, 0xb8, 0x29, 0x32, 0x8b, 0x7c,
	0xd8, 0xc9, 0xea, 0xf2, 0x6b, 0x91, 0x24, 0xdc, 0x18, 0xff, 0x3f, 0x34, 0x82, 0x81, 0x90, 0x37,
	0x34, 0x13, 0xac, 0xde, 0x1e, 0xdf, 0x43, 0x43, 0x80, 0x56, 0xf8, 0x72, 0xcb, 0xfc, 0x0e, 0xda,
	0x83, 0xdd, 0x2b, 0xa5, 0x63, 0xc1, 0x18, 0x97, 0x0e, 0xd9, 0xa8, 0x33, 0x56, 0x7d, 0xe3, 0x72,
	0x7a, 0xb7, 0x10, 0x9a, 0xfb, 0x5d, 0xb4, 0x0f, 0xa3, 0x36, 0x73, 0x4e, 0x85, 0xb4, 0x54, 0x48,
	0x7f, 0xf3, 0xe8, 0x03, 0x0c, 0x09, 0x4f, 0x85, 0xb1, 0x5c, 0xb7, 0xa7, 0xef, 0xc3, 0x48, 0xb7,
	0xca, 0xbf, 0x0f, 0x18, 0x02, 0xd4, 0x6b, 0x34, 0xbd, 0x13, 0xc6, 0x1a, 0xdf, 0xab, 0xcf, 0x6b,
	0x7b, 0xbd, 0x57, 0x3a, 0xa7, 0xd6, 0xef, 0x4c, 0x76, 0x7e, 0x2d, 0x0f, 0xbc, 0xdf, 0xcb, 0x03,
	0xef, 0xcf, 0xf2, 0xc0, 0x8b, 0x7b, 0x6e, 0xe1, 0x4f, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x96,
	0x32, 0x8d, 0x2d, 0x41, 0x03, 0x00, 0x00,
}
