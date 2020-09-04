// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01
// source: .vppapi/core/p2p_ethernet.api.json

// Package p2p_ethernet contains generated bindings for API file p2p_ethernet.api.
//
// Contents:
//   2 aliases
//   6 enums
//   4 messages
//
package p2p_ethernet

import (
	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	"net"
	"strconv"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "p2p_ethernet"
	APIVersion = "1.0.0"
	VersionCrc = 0x282f1ed1
)

// IfStatusFlags defines enum 'if_status_flags'.
type IfStatusFlags uint32

const (
	IF_STATUS_API_FLAG_ADMIN_UP IfStatusFlags = 1
	IF_STATUS_API_FLAG_LINK_UP  IfStatusFlags = 2
)

var (
	IfStatusFlags_name = map[uint32]string{
		1: "IF_STATUS_API_FLAG_ADMIN_UP",
		2: "IF_STATUS_API_FLAG_LINK_UP",
	}
	IfStatusFlags_value = map[string]uint32{
		"IF_STATUS_API_FLAG_ADMIN_UP": 1,
		"IF_STATUS_API_FLAG_LINK_UP":  2,
	}
)

func (x IfStatusFlags) String() string {
	s, ok := IfStatusFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := IfStatusFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "IfStatusFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// IfType defines enum 'if_type'.
type IfType uint32

const (
	IF_API_TYPE_HARDWARE IfType = 1
	IF_API_TYPE_SUB      IfType = 2
	IF_API_TYPE_P2P      IfType = 3
	IF_API_TYPE_PIPE     IfType = 4
)

var (
	IfType_name = map[uint32]string{
		1: "IF_API_TYPE_HARDWARE",
		2: "IF_API_TYPE_SUB",
		3: "IF_API_TYPE_P2P",
		4: "IF_API_TYPE_PIPE",
	}
	IfType_value = map[string]uint32{
		"IF_API_TYPE_HARDWARE": 1,
		"IF_API_TYPE_SUB":      2,
		"IF_API_TYPE_P2P":      3,
		"IF_API_TYPE_PIPE":     4,
	}
)

func (x IfType) String() string {
	s, ok := IfType_name[uint32(x)]
	if ok {
		return s
	}
	return "IfType(" + strconv.Itoa(int(x)) + ")"
}

// LinkDuplex defines enum 'link_duplex'.
type LinkDuplex uint32

const (
	LINK_DUPLEX_API_UNKNOWN LinkDuplex = 0
	LINK_DUPLEX_API_HALF    LinkDuplex = 1
	LINK_DUPLEX_API_FULL    LinkDuplex = 2
)

var (
	LinkDuplex_name = map[uint32]string{
		0: "LINK_DUPLEX_API_UNKNOWN",
		1: "LINK_DUPLEX_API_HALF",
		2: "LINK_DUPLEX_API_FULL",
	}
	LinkDuplex_value = map[string]uint32{
		"LINK_DUPLEX_API_UNKNOWN": 0,
		"LINK_DUPLEX_API_HALF":    1,
		"LINK_DUPLEX_API_FULL":    2,
	}
)

func (x LinkDuplex) String() string {
	s, ok := LinkDuplex_name[uint32(x)]
	if ok {
		return s
	}
	return "LinkDuplex(" + strconv.Itoa(int(x)) + ")"
}

// MtuProto defines enum 'mtu_proto'.
type MtuProto uint32

const (
	MTU_PROTO_API_L3   MtuProto = 1
	MTU_PROTO_API_IP4  MtuProto = 2
	MTU_PROTO_API_IP6  MtuProto = 3
	MTU_PROTO_API_MPLS MtuProto = 4
	MTU_PROTO_API_N    MtuProto = 5
)

var (
	MtuProto_name = map[uint32]string{
		1: "MTU_PROTO_API_L3",
		2: "MTU_PROTO_API_IP4",
		3: "MTU_PROTO_API_IP6",
		4: "MTU_PROTO_API_MPLS",
		5: "MTU_PROTO_API_N",
	}
	MtuProto_value = map[string]uint32{
		"MTU_PROTO_API_L3":   1,
		"MTU_PROTO_API_IP4":  2,
		"MTU_PROTO_API_IP6":  3,
		"MTU_PROTO_API_MPLS": 4,
		"MTU_PROTO_API_N":    5,
	}
)

func (x MtuProto) String() string {
	s, ok := MtuProto_name[uint32(x)]
	if ok {
		return s
	}
	return "MtuProto(" + strconv.Itoa(int(x)) + ")"
}

// RxMode defines enum 'rx_mode'.
type RxMode uint32

const (
	RX_MODE_API_UNKNOWN   RxMode = 0
	RX_MODE_API_POLLING   RxMode = 1
	RX_MODE_API_INTERRUPT RxMode = 2
	RX_MODE_API_ADAPTIVE  RxMode = 3
	RX_MODE_API_DEFAULT   RxMode = 4
)

var (
	RxMode_name = map[uint32]string{
		0: "RX_MODE_API_UNKNOWN",
		1: "RX_MODE_API_POLLING",
		2: "RX_MODE_API_INTERRUPT",
		3: "RX_MODE_API_ADAPTIVE",
		4: "RX_MODE_API_DEFAULT",
	}
	RxMode_value = map[string]uint32{
		"RX_MODE_API_UNKNOWN":   0,
		"RX_MODE_API_POLLING":   1,
		"RX_MODE_API_INTERRUPT": 2,
		"RX_MODE_API_ADAPTIVE":  3,
		"RX_MODE_API_DEFAULT":   4,
	}
)

func (x RxMode) String() string {
	s, ok := RxMode_name[uint32(x)]
	if ok {
		return s
	}
	return "RxMode(" + strconv.Itoa(int(x)) + ")"
}

// SubIfFlags defines enum 'sub_if_flags'.
type SubIfFlags uint32

const (
	SUB_IF_API_FLAG_NO_TAGS           SubIfFlags = 1
	SUB_IF_API_FLAG_ONE_TAG           SubIfFlags = 2
	SUB_IF_API_FLAG_TWO_TAGS          SubIfFlags = 4
	SUB_IF_API_FLAG_DOT1AD            SubIfFlags = 8
	SUB_IF_API_FLAG_EXACT_MATCH       SubIfFlags = 16
	SUB_IF_API_FLAG_DEFAULT           SubIfFlags = 32
	SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY SubIfFlags = 64
	SUB_IF_API_FLAG_INNER_VLAN_ID_ANY SubIfFlags = 128
	SUB_IF_API_FLAG_MASK_VNET         SubIfFlags = 254
	SUB_IF_API_FLAG_DOT1AH            SubIfFlags = 256
)

var (
	SubIfFlags_name = map[uint32]string{
		1:   "SUB_IF_API_FLAG_NO_TAGS",
		2:   "SUB_IF_API_FLAG_ONE_TAG",
		4:   "SUB_IF_API_FLAG_TWO_TAGS",
		8:   "SUB_IF_API_FLAG_DOT1AD",
		16:  "SUB_IF_API_FLAG_EXACT_MATCH",
		32:  "SUB_IF_API_FLAG_DEFAULT",
		64:  "SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY",
		128: "SUB_IF_API_FLAG_INNER_VLAN_ID_ANY",
		254: "SUB_IF_API_FLAG_MASK_VNET",
		256: "SUB_IF_API_FLAG_DOT1AH",
	}
	SubIfFlags_value = map[string]uint32{
		"SUB_IF_API_FLAG_NO_TAGS":           1,
		"SUB_IF_API_FLAG_ONE_TAG":           2,
		"SUB_IF_API_FLAG_TWO_TAGS":          4,
		"SUB_IF_API_FLAG_DOT1AD":            8,
		"SUB_IF_API_FLAG_EXACT_MATCH":       16,
		"SUB_IF_API_FLAG_DEFAULT":           32,
		"SUB_IF_API_FLAG_OUTER_VLAN_ID_ANY": 64,
		"SUB_IF_API_FLAG_INNER_VLAN_ID_ANY": 128,
		"SUB_IF_API_FLAG_MASK_VNET":         254,
		"SUB_IF_API_FLAG_DOT1AH":            256,
	}
)

func (x SubIfFlags) String() string {
	s, ok := SubIfFlags_name[uint32(x)]
	if ok {
		return s
	}
	str := func(n uint32) string {
		s, ok := SubIfFlags_name[uint32(n)]
		if ok {
			return s
		}
		return "SubIfFlags(" + strconv.Itoa(int(n)) + ")"
	}
	for i := uint32(0); i <= 32; i++ {
		val := uint32(x)
		if val&(1<<i) != 0 {
			if s != "" {
				s += "|"
			}
			s += str(1 << i)
		}
	}
	if s == "" {
		return str(uint32(x))
	}
	return s
}

// InterfaceIndex defines alias 'interface_index'.
type InterfaceIndex uint32

// MacAddress defines alias 'mac_address'.
type MacAddress [6]uint8

func ParseMacAddress(s string) (MacAddress, error) {
	var macaddr MacAddress
	mac, err := net.ParseMAC(s)
	if err != nil {
		return macaddr, err
	}
	copy(macaddr[:], mac[:])
	return macaddr, nil
}
func (x MacAddress) ToMAC() net.HardwareAddr {
	return net.HardwareAddr(x[:])
}
func (x MacAddress) String() string {
	return x.ToMAC().String()
}
func (x *MacAddress) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}
func (x *MacAddress) UnmarshalText(text []byte) error {
	mac, err := ParseMacAddress(string(text))
	if err != nil {
		return err
	}
	*x = mac
	return nil
}

// P2pEthernetAdd defines message 'p2p_ethernet_add'.
type P2pEthernetAdd struct {
	ParentIfIndex InterfaceIndex `binapi:"interface_index,name=parent_if_index" json:"parent_if_index,omitempty"`
	SubifID       uint32         `binapi:"u32,name=subif_id" json:"subif_id,omitempty"`
	RemoteMac     MacAddress     `binapi:"mac_address,name=remote_mac" json:"remote_mac,omitempty"`
}

func (m *P2pEthernetAdd) Reset()               { *m = P2pEthernetAdd{} }
func (*P2pEthernetAdd) GetMessageName() string { return "p2p_ethernet_add" }
func (*P2pEthernetAdd) GetCrcString() string   { return "eeb8e717" }
func (*P2pEthernetAdd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *P2pEthernetAdd) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.ParentIfIndex
	size += 4     // m.SubifID
	size += 1 * 6 // m.RemoteMac
	return size
}
func (m *P2pEthernetAdd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.ParentIfIndex))
	buf.EncodeUint32(m.SubifID)
	buf.EncodeBytes(m.RemoteMac[:], 6)
	return buf.Bytes(), nil
}
func (m *P2pEthernetAdd) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ParentIfIndex = InterfaceIndex(buf.DecodeUint32())
	m.SubifID = buf.DecodeUint32()
	copy(m.RemoteMac[:], buf.DecodeBytes(6))
	return nil
}

// P2pEthernetAddReply defines message 'p2p_ethernet_add_reply'.
type P2pEthernetAddReply struct {
	Retval    int32          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *P2pEthernetAddReply) Reset()               { *m = P2pEthernetAddReply{} }
func (*P2pEthernetAddReply) GetMessageName() string { return "p2p_ethernet_add_reply" }
func (*P2pEthernetAddReply) GetCrcString() string   { return "5383d31f" }
func (*P2pEthernetAddReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *P2pEthernetAddReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *P2pEthernetAddReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *P2pEthernetAddReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = InterfaceIndex(buf.DecodeUint32())
	return nil
}

// P2pEthernetDel defines message 'p2p_ethernet_del'.
type P2pEthernetDel struct {
	ParentIfIndex InterfaceIndex `binapi:"interface_index,name=parent_if_index" json:"parent_if_index,omitempty"`
	RemoteMac     MacAddress     `binapi:"mac_address,name=remote_mac" json:"remote_mac,omitempty"`
}

func (m *P2pEthernetDel) Reset()               { *m = P2pEthernetDel{} }
func (*P2pEthernetDel) GetMessageName() string { return "p2p_ethernet_del" }
func (*P2pEthernetDel) GetCrcString() string   { return "0b62c386" }
func (*P2pEthernetDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *P2pEthernetDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.ParentIfIndex
	size += 1 * 6 // m.RemoteMac
	return size
}
func (m *P2pEthernetDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.ParentIfIndex))
	buf.EncodeBytes(m.RemoteMac[:], 6)
	return buf.Bytes(), nil
}
func (m *P2pEthernetDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.ParentIfIndex = InterfaceIndex(buf.DecodeUint32())
	copy(m.RemoteMac[:], buf.DecodeBytes(6))
	return nil
}

// P2pEthernetDelReply defines message 'p2p_ethernet_del_reply'.
type P2pEthernetDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *P2pEthernetDelReply) Reset()               { *m = P2pEthernetDelReply{} }
func (*P2pEthernetDelReply) GetMessageName() string { return "p2p_ethernet_del_reply" }
func (*P2pEthernetDelReply) GetCrcString() string   { return "e8d4e804" }
func (*P2pEthernetDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *P2pEthernetDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *P2pEthernetDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *P2pEthernetDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_p2p_ethernet_binapi_init() }
func file_p2p_ethernet_binapi_init() {
	api.RegisterMessage((*P2pEthernetAdd)(nil), "p2p_ethernet_add_eeb8e717")
	api.RegisterMessage((*P2pEthernetAddReply)(nil), "p2p_ethernet_add_reply_5383d31f")
	api.RegisterMessage((*P2pEthernetDel)(nil), "p2p_ethernet_del_0b62c386")
	api.RegisterMessage((*P2pEthernetDelReply)(nil), "p2p_ethernet_del_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*P2pEthernetAdd)(nil),
		(*P2pEthernetAddReply)(nil),
		(*P2pEthernetDel)(nil),
		(*P2pEthernetDelReply)(nil),
	}
}