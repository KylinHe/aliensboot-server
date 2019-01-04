package object

import (
	"fmt"
	"net"
	"os"
	//"proto_msg"
	"encoding/binary"
	"github.com/KylinHe/aliensboot-core/common/cipher/xxtea"
	"github.com/gogo/protobuf/proto"
	"sync"
	"time"

	"github.com/KylinHe/aliensboot-core/protocol"
	"test/protobuf_client/base"
)

type Player struct {
	sync.RWMutex
	m_Idx int32 //索引

	syncStartTime time.Time

	seq      int32
	Token    string
	Username string
	Password string

	Game      bool
	synctime  int64
	secretkey string

	m_SrvInfo string //游戏服务器信息
	passport  string //
	game      string

	m_pCon *net.TCPConn //tcp 的连接器

	m_bReadyDisCon bool //是否准备断开连接

	m_Channel        chan interface{}
	m_Channel_Closed bool //通道是否 已经关闭了

	m_State int32 //玩家状态
}

/**
*	@brief 获取角色索引
 */
func (this *Player) GetIdx() int32 { return this.m_Idx }

/**
*	@brief 初始化
 */
func (this *Player) Init(idx int32, passportServer string, gameServer string, synctime int64, secretkey string) bool {
	this.m_Idx = idx
	this.synctime = synctime
	this.m_SrvInfo = passportServer
	this.passport = passportServer
	this.secretkey = secretkey
	this.game = gameServer
	this.m_State = base.PLAYER_STATE_NONE //默认设置为 无状态
	this._open_channel()                  //打开通道

	if this.IsConnect() == false {
		//fmt.Printf(">>> try to connect \n");
		if this.Connect() == false { //连接失败
			fmt.Printf(">>> connect failed %v\n", this.m_Idx)
			return false
		}
		fmt.Printf(">>> connect success %v\n", this.m_Idx)
	}
	return true
}

func (this *Player) ReconnectGame() bool {
	//已经连接了这个地址不需要重复连接
	if this.IsConnect() && this.m_SrvInfo == this.game {
		return true
	}
	this.m_SrvInfo = this.game
	this.m_State = base.PLAYER_STATE_NONE //默认设置为 无状态
	this._open_channel()                  //打开通道

	if this.IsConnect() {
		this.DisConnect()
		fmt.Printf(">>> reconnect success %v \n", this.m_Idx)
	}
	fmt.Printf(">>> try to connect %v\n", this.m_Idx)
	if this.Connect() == false { //连接失败
		fmt.Printf(">>> reconnect failed %v\n", this.m_Idx)
		return false
	} else {
		fmt.Printf(">>> reconnect success %v\n", this.m_Idx)
	}
	return true
}

func (this *Player) ReconnectPassport() bool {
	//已经连接了这个地址不需要重复连接
	if this.IsConnect() && this.m_SrvInfo == this.passport {
		return true
	}
	this.m_SrvInfo = this.passport
	this.m_State = base.PLAYER_STATE_NONE //默认设置为 无状态
	this._open_channel()                  //打开通道

	if this.IsConnect() {
		this.DisConnect()
		fmt.Printf(">>> reconnect success %v \n", this.m_Idx)
	}
	fmt.Printf(">>> try to connect %v\n", this.m_Idx)
	if this.Connect() == false { //连接失败
		fmt.Printf(">>> reconnect failed %v\n", this.m_Idx)
		return false
	} else {
		fmt.Printf(">>> reconnect success %v\n", this.m_Idx)
	}
	return true
}

/**
*	@brief 插入操作码
 */
func (this *Player) AcceptOp(op int) {
	if this.m_Channel == nil {
		return
	}
	select {
	case this.m_Channel <- op:
	default:
		fmt.Printf("%d message channel full\n", this.m_Idx)
		//TODO 消息管道满了需要通知客户端消息请求太过频繁
	}
}

func (this *Player) isCipher() bool {
	return this.secretkey != ""
}

/**
*	@brief 请求连接
 */
func (this *Player) Connect() bool {
	this.Lock()
	defer this.Unlock()
	if this.m_SrvInfo == "" {
		return false
	}
	if this.m_pCon != nil {
		return false
	}
	addr, err := net.ResolveTCPAddr("tcp", this.m_SrvInfo)
	if err != nil {
		fmt.Printf("Error: Idx=%d, %s\n", this.m_Idx, err.Error())
		return false
	}
	this.m_pCon, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("Error: Idx=%d, %s\n", this.m_Idx, err.Error())
		this.m_pCon = nil //重新设置为nil
		return false
	}

	buf := make([]byte, 8182)
	go func() {
		for {
			if this.m_pCon == nil {
				break
			}

			len, err := this.m_pCon.Read(buf)
			if err != nil {
				this.Lock()
				defer this.Unlock()
				this.m_pCon = nil
				fmt.Printf("read error: %v\n", err.Error())
				break
			}
			//fmt.Printf("Recv msg = %v\n   %v", len, buf )

			data := buf[4:len]
			if this.isCipher() {
				data = xxtea.Decrypt(data, []byte(this.secretkey))
			}

			recv := &protocol.Response{}
			err = proto.Unmarshal(data, recv)
			if err != nil {
				fmt.Printf("unmarshaling error: %v\n", err.Error())
				continue
				//this.m_pCon = nil
				//break
			}
			this._receive_Message(recv)
			//if code == 11024{
			//	this.m_State = base.PLAYER_STATE_GAME;//表示 玩家已经进入游戏
			//}
		}
	}()

	return true
}

/**
*	@brief 目前是否处于连接中
 */
func (this *Player) IsConnect() bool {
	this.RLock()
	defer this.RUnlock()
	if this.m_pCon == nil {
		return false
	}
	return true
}

/**
*	@brief 断开连接
 */
func (this *Player) DisConnect() bool {
	this.Lock()
	defer this.Unlock()
	if this.m_pCon == nil {
		return false
	}
	err := this.m_pCon.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Idx=%d, %s", this.m_Idx, err.Error())
	}
	this.m_pCon = nil
	return true
}

func (this *Player) _open_channel() { //打开消息管道
	if this.m_Channel != nil {
		return
	}
	this.m_Channel_Closed = false
	this.m_Channel = make(chan interface{}, 10) //10个通道的最大数量
	go func() {
		for {
			//只要消息管道没有关闭，就一直等待消息
			v, open := <-this.m_Channel
			if !open {
				this.m_Channel = nil
				break
			}
			opType, _ := v.(int)
			this._accept_op(opType)
		}
		this._close_channel()
	}()
}

func (this *Player) _close_channel() { //关闭通道
	if this.m_Channel_Closed == true {
		return
	} //已经准备要关闭了
	if this.m_Channel == nil {
		return
	}
	close(this.m_Channel)
	this.m_Channel_Closed = true
}

func (this *Player) _accept_op(opType int) {
	switch opType {
	case base.OP_SYNC:
		this.syncData()
		break
	}
}

func (this *Player) syncData() {
	this.seq++
	sessionID := this.seq*10000 + this.m_Idx
	this.syncStartTime = time.Now()
	this._send_Message(1, BuildLoginRequest(this.Username, "111111", sessionID))
	//this._send_Message(4, BuildRequest2(sessionID, this.Username, strconv.Itoa(int(this.seq))))
}

//func (this *Player)_deal_register(){//进行登录测试
//	if this.m_State == base.PLAYER_STATE_NONE{//无状态, 好吧 那就尝试请求登录吧
//
//
//		//fmt.Printf("=================== Account = %s Login\n", account);
//		//request.Req_Client_Login = &protocol.Req_Client_Login_{ Account: proto.String(account),Platform:proto.String("00") }
//
//		this.m_State =  base.PLAYER_STATE_LOGIN;
//	}else if this.m_State == base.PLAYER_STATE_LOGIN{//登录状态中， 好吧先不进行任何处理
//
//	}else if this.m_State == base.PLAYER_STATE_GAME{//登录成功状态， 好吧进行相应处理
//		this.DisConnect();
//		this.m_State =  base.PLAYER_STATE_NONE;
//	}
//}

func (this *Player) _send_Message(id uint16, message proto.Message) { //发送消息
	this.RLock()
	defer this.RUnlock()
	if this.m_pCon == nil {
		return
	}

	buff, err := proto.Marshal(message)

	if err != nil { //序列化失败
		fmt.Printf(">> send message err = %v\n", err.Error())
		return
	}
	//fmt.Printf("%v => send: %v - %v %v\n", this.Username, id, message, time.Now())
	if this.isCipher() {
		buff = xxtea.Encrypt(buff, []byte(this.secretkey))
		//fmt.Printf("key %v", this.secretkey)
	}

	m := make([]byte, len(buff)+4)
	binary.LittleEndian.PutUint16(m, uint16(len(buff))+2) //
	binary.LittleEndian.PutUint16(m[2:], id)

	copy(m[4:], buff)

	this.m_pCon.Write(m) //把消息发送出去
}

func (this *Player) _receive_Message(message proto.Message) { //接收消息
	fmt.Printf("%v => receive: %v %v\n", this.Username, message, time.Now())
	this.syncData()
	//code := message.GetSequence()
}
