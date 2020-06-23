// Code generated by GoVPP's binapi-generator. DO NOT EDIT.
// versions:
//  binapi-generator: v0.4.0-dev
//  VPP:              20.01-45~g7a071e370~b63
// source: /usr/share/vpp/api/core/memclnt.api.json

/*
Package memclnt contains generated code for VPP binary API defined by memclnt.api (version 2.1.0).

It consists of:
	 22 messages
	  2 types
*/
package memclnt

import (
	"bytes"
	"context"
	"encoding/binary"
	"io"
	"math"
	"strconv"

	api "git.fd.io/govpp.git/api"
	codec "git.fd.io/govpp.git/codec"
	struc "github.com/lunixbochs/struc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2 // please upgrade the GoVPP api package

const (
	// ModuleName is the name of this module.
	ModuleName = "memclnt"
	// APIVersion is the API version of this module.
	APIVersion = "2.1.0"
	// VersionCrc is the CRC of this module.
	VersionCrc = 0x8d3dd881
)

// MessageTableEntry represents VPP binary API type 'message_table_entry'.
type MessageTableEntry struct {
	Index uint16 `binapi:"u16,name=index" json:"index,omitempty"`
	Name  string `binapi:"string[64],name=name" json:"name,omitempty" struc:"[64]byte"`
}

func (*MessageTableEntry) GetTypeName() string { return "message_table_entry" }

// ModuleVersion represents VPP binary API type 'module_version'.
type ModuleVersion struct {
	Major uint32 `binapi:"u32,name=major" json:"major,omitempty"`
	Minor uint32 `binapi:"u32,name=minor" json:"minor,omitempty"`
	Patch uint32 `binapi:"u32,name=patch" json:"patch,omitempty"`
	Name  string `binapi:"string[64],name=name" json:"name,omitempty" struc:"[64]byte"`
}

func (*ModuleVersion) GetTypeName() string { return "module_version" }

// APIVersions represents VPP binary API message 'api_versions'.
type APIVersions struct{}

func (m *APIVersions) Reset()                        { *m = APIVersions{} }
func (*APIVersions) GetMessageName() string          { return "api_versions" }
func (*APIVersions) GetCrcString() string            { return "51077d14" }
func (*APIVersions) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *APIVersions) Size() int {
	if m == nil {
		return 0
	}
	var size int
	return size
}
func (m *APIVersions) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	return buf, nil
}
func (m *APIVersions) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	return nil
}

// APIVersionsReply represents VPP binary API message 'api_versions_reply'.
type APIVersionsReply struct {
	Retval      int32           `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count       uint32          `binapi:"u32,name=count" json:"count,omitempty" struc:"sizeof=APIVersions"`
	APIVersions []ModuleVersion `binapi:"module_version[count],name=api_versions" json:"api_versions,omitempty"`
}

func (m *APIVersionsReply) Reset()                        { *m = APIVersionsReply{} }
func (*APIVersionsReply) GetMessageName() string          { return "api_versions_reply" }
func (*APIVersionsReply) GetCrcString() string            { return "5f0d99d6" }
func (*APIVersionsReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *APIVersionsReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Retval
	size += 4
	// field[1] m.Count
	size += 4
	// field[1] m.APIVersions
	for j1 := 0; j1 < len(m.APIVersions); j1++ {
		var s1 ModuleVersion
		_ = s1
		if j1 < len(m.APIVersions) {
			s1 = m.APIVersions[j1]
		}
		// field[2] s1.Major
		size += 4
		// field[2] s1.Minor
		size += 4
		// field[2] s1.Patch
		size += 4
		// field[2] s1.Name
		size += 64
	}
	return size
}
func (m *APIVersionsReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Retval
	o.PutUint32(buf[pos:pos+4], uint32(m.Retval))
	pos += 4
	// field[1] m.Count
	o.PutUint32(buf[pos:pos+4], uint32(len(m.APIVersions)))
	pos += 4
	// field[1] m.APIVersions
	for j1 := 0; j1 < len(m.APIVersions); j1++ {
		var v1 ModuleVersion
		if j1 < len(m.APIVersions) {
			v1 = m.APIVersions[j1]
		}
		// field[2] v1.Major
		o.PutUint32(buf[pos:pos+4], uint32(v1.Major))
		pos += 4
		// field[2] v1.Minor
		o.PutUint32(buf[pos:pos+4], uint32(v1.Minor))
		pos += 4
		// field[2] v1.Patch
		o.PutUint32(buf[pos:pos+4], uint32(v1.Patch))
		pos += 4
		// field[2] v1.Name
		copy(buf[pos:pos+64], v1.Name)
		pos += 64
	}
	return buf, nil
}
func (m *APIVersionsReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Retval
	m.Retval = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Count
	m.Count = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.APIVersions
	m.APIVersions = make([]ModuleVersion, int(m.Count))
	for j1 := 0; j1 < int(m.Count); j1++ {
		// field[2] m.APIVersions[j1].Major
		m.APIVersions[j1].Major = uint32(o.Uint32(tmp[pos : pos+4]))
		pos += 4
		// field[2] m.APIVersions[j1].Minor
		m.APIVersions[j1].Minor = uint32(o.Uint32(tmp[pos : pos+4]))
		pos += 4
		// field[2] m.APIVersions[j1].Patch
		m.APIVersions[j1].Patch = uint32(o.Uint32(tmp[pos : pos+4]))
		pos += 4
		// field[2] m.APIVersions[j1].Name
		{
			nul := bytes.Index(tmp[pos:pos+64], []byte{0x00})
			m.APIVersions[j1].Name = codec.DecodeString(tmp[pos : pos+nul])
			pos += 64
		}
	}
	return nil
}

// GetFirstMsgID represents VPP binary API message 'get_first_msg_id'.
type GetFirstMsgID struct {
	Name string `binapi:"string[64],name=name" json:"name,omitempty" struc:"[64]byte"`
}

func (m *GetFirstMsgID) Reset()                        { *m = GetFirstMsgID{} }
func (*GetFirstMsgID) GetMessageName() string          { return "get_first_msg_id" }
func (*GetFirstMsgID) GetCrcString() string            { return "ebf79a66" }
func (*GetFirstMsgID) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *GetFirstMsgID) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Name
	size += 64
	return size
}
func (m *GetFirstMsgID) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Name
	copy(buf[pos:pos+64], m.Name)
	pos += 64
	return buf, nil
}
func (m *GetFirstMsgID) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Name
	{
		nul := bytes.Index(tmp[pos:pos+64], []byte{0x00})
		m.Name = codec.DecodeString(tmp[pos : pos+nul])
		pos += 64
	}
	return nil
}

// GetFirstMsgIDReply represents VPP binary API message 'get_first_msg_id_reply'.
type GetFirstMsgIDReply struct {
	Retval     int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	FirstMsgID uint16 `binapi:"u16,name=first_msg_id" json:"first_msg_id,omitempty"`
}

func (m *GetFirstMsgIDReply) Reset()                        { *m = GetFirstMsgIDReply{} }
func (*GetFirstMsgIDReply) GetMessageName() string          { return "get_first_msg_id_reply" }
func (*GetFirstMsgIDReply) GetCrcString() string            { return "7d337472" }
func (*GetFirstMsgIDReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *GetFirstMsgIDReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Retval
	size += 4
	// field[1] m.FirstMsgID
	size += 2
	return size
}
func (m *GetFirstMsgIDReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Retval
	o.PutUint32(buf[pos:pos+4], uint32(m.Retval))
	pos += 4
	// field[1] m.FirstMsgID
	o.PutUint16(buf[pos:pos+2], uint16(m.FirstMsgID))
	pos += 2
	return buf, nil
}
func (m *GetFirstMsgIDReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Retval
	m.Retval = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.FirstMsgID
	m.FirstMsgID = uint16(o.Uint16(tmp[pos : pos+2]))
	pos += 2
	return nil
}

// MemclntCreate represents VPP binary API message 'memclnt_create'.
type MemclntCreate struct {
	CtxQuota    int32    `binapi:"i32,name=ctx_quota" json:"ctx_quota,omitempty"`
	InputQueue  uint64   `binapi:"u64,name=input_queue" json:"input_queue,omitempty"`
	Name        string   `binapi:"string[64],name=name" json:"name,omitempty" struc:"[64]byte"`
	APIVersions []uint32 `binapi:"u32[8],name=api_versions" json:"api_versions,omitempty" struc:"[8]uint32"`
}

func (m *MemclntCreate) Reset()                        { *m = MemclntCreate{} }
func (*MemclntCreate) GetMessageName() string          { return "memclnt_create" }
func (*MemclntCreate) GetCrcString() string            { return "9c5e1c2f" }
func (*MemclntCreate) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *MemclntCreate) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.CtxQuota
	size += 4
	// field[1] m.InputQueue
	size += 8
	// field[1] m.Name
	size += 64
	// field[1] m.APIVersions
	size += 32
	return size
}
func (m *MemclntCreate) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.CtxQuota
	o.PutUint32(buf[pos:pos+4], uint32(m.CtxQuota))
	pos += 4
	// field[1] m.InputQueue
	o.PutUint64(buf[pos:pos+8], uint64(m.InputQueue))
	pos += 8
	// field[1] m.Name
	copy(buf[pos:pos+64], m.Name)
	pos += 64
	// field[1] m.APIVersions
	for i := 0; i < 8; i++ {
		var x uint32
		if i < len(m.APIVersions) {
			x = uint32(m.APIVersions[i])
		}
		o.PutUint32(buf[pos:pos+4], uint32(x))
		pos += 4
	}
	return buf, nil
}
func (m *MemclntCreate) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.CtxQuota
	m.CtxQuota = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.InputQueue
	m.InputQueue = uint64(o.Uint64(tmp[pos : pos+8]))
	pos += 8
	// field[1] m.Name
	{
		nul := bytes.Index(tmp[pos:pos+64], []byte{0x00})
		m.Name = codec.DecodeString(tmp[pos : pos+nul])
		pos += 64
	}
	// field[1] m.APIVersions
	m.APIVersions = make([]uint32, 8)
	for i := 0; i < len(m.APIVersions); i++ {
		m.APIVersions[i] = uint32(o.Uint32(tmp[pos : pos+4]))
		pos += 4
	}
	return nil
}

// MemclntCreateReply represents VPP binary API message 'memclnt_create_reply'.
type MemclntCreateReply struct {
	Response     int32  `binapi:"i32,name=response" json:"response,omitempty"`
	Handle       uint64 `binapi:"u64,name=handle" json:"handle,omitempty"`
	Index        uint32 `binapi:"u32,name=index" json:"index,omitempty"`
	MessageTable uint64 `binapi:"u64,name=message_table" json:"message_table,omitempty"`
}

func (m *MemclntCreateReply) Reset()                        { *m = MemclntCreateReply{} }
func (*MemclntCreateReply) GetMessageName() string          { return "memclnt_create_reply" }
func (*MemclntCreateReply) GetCrcString() string            { return "42ec4560" }
func (*MemclntCreateReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *MemclntCreateReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Response
	size += 4
	// field[1] m.Handle
	size += 8
	// field[1] m.Index
	size += 4
	// field[1] m.MessageTable
	size += 8
	return size
}
func (m *MemclntCreateReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Response
	o.PutUint32(buf[pos:pos+4], uint32(m.Response))
	pos += 4
	// field[1] m.Handle
	o.PutUint64(buf[pos:pos+8], uint64(m.Handle))
	pos += 8
	// field[1] m.Index
	o.PutUint32(buf[pos:pos+4], uint32(m.Index))
	pos += 4
	// field[1] m.MessageTable
	o.PutUint64(buf[pos:pos+8], uint64(m.MessageTable))
	pos += 8
	return buf, nil
}
func (m *MemclntCreateReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Response
	m.Response = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Handle
	m.Handle = uint64(o.Uint64(tmp[pos : pos+8]))
	pos += 8
	// field[1] m.Index
	m.Index = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.MessageTable
	m.MessageTable = uint64(o.Uint64(tmp[pos : pos+8]))
	pos += 8
	return nil
}

// MemclntDelete represents VPP binary API message 'memclnt_delete'.
type MemclntDelete struct {
	Index     uint32 `binapi:"u32,name=index" json:"index,omitempty"`
	Handle    uint64 `binapi:"u64,name=handle" json:"handle,omitempty"`
	DoCleanup bool   `binapi:"bool,name=do_cleanup" json:"do_cleanup,omitempty"`
}

func (m *MemclntDelete) Reset()                        { *m = MemclntDelete{} }
func (*MemclntDelete) GetMessageName() string          { return "memclnt_delete" }
func (*MemclntDelete) GetCrcString() string            { return "7e1c04e3" }
func (*MemclntDelete) GetMessageType() api.MessageType { return api.OtherMessage }

func (m *MemclntDelete) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Index
	size += 4
	// field[1] m.Handle
	size += 8
	// field[1] m.DoCleanup
	size += 1
	return size
}
func (m *MemclntDelete) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Index
	o.PutUint32(buf[pos:pos+4], uint32(m.Index))
	pos += 4
	// field[1] m.Handle
	o.PutUint64(buf[pos:pos+8], uint64(m.Handle))
	pos += 8
	// field[1] m.DoCleanup
	if m.DoCleanup {
		buf[pos] = 1
	}
	pos += 1
	return buf, nil
}
func (m *MemclntDelete) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Index
	m.Index = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Handle
	m.Handle = uint64(o.Uint64(tmp[pos : pos+8]))
	pos += 8
	// field[1] m.DoCleanup
	m.DoCleanup = tmp[pos] != 0
	pos += 1
	return nil
}

// MemclntDeleteReply represents VPP binary API message 'memclnt_delete_reply'.
type MemclntDeleteReply struct {
	Response int32  `binapi:"i32,name=response" json:"response,omitempty"`
	Handle   uint64 `binapi:"u64,name=handle" json:"handle,omitempty"`
}

func (m *MemclntDeleteReply) Reset()                        { *m = MemclntDeleteReply{} }
func (*MemclntDeleteReply) GetMessageName() string          { return "memclnt_delete_reply" }
func (*MemclntDeleteReply) GetCrcString() string            { return "3d3b6312" }
func (*MemclntDeleteReply) GetMessageType() api.MessageType { return api.OtherMessage }

func (m *MemclntDeleteReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Response
	size += 4
	// field[1] m.Handle
	size += 8
	return size
}
func (m *MemclntDeleteReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Response
	o.PutUint32(buf[pos:pos+4], uint32(m.Response))
	pos += 4
	// field[1] m.Handle
	o.PutUint64(buf[pos:pos+8], uint64(m.Handle))
	pos += 8
	return buf, nil
}
func (m *MemclntDeleteReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Response
	m.Response = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Handle
	m.Handle = uint64(o.Uint64(tmp[pos : pos+8]))
	pos += 8
	return nil
}

// MemclntKeepalive represents VPP binary API message 'memclnt_keepalive'.
type MemclntKeepalive struct{}

func (m *MemclntKeepalive) Reset()                        { *m = MemclntKeepalive{} }
func (*MemclntKeepalive) GetMessageName() string          { return "memclnt_keepalive" }
func (*MemclntKeepalive) GetCrcString() string            { return "51077d14" }
func (*MemclntKeepalive) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *MemclntKeepalive) Size() int {
	if m == nil {
		return 0
	}
	var size int
	return size
}
func (m *MemclntKeepalive) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	return buf, nil
}
func (m *MemclntKeepalive) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	return nil
}

// MemclntKeepaliveReply represents VPP binary API message 'memclnt_keepalive_reply'.
type MemclntKeepaliveReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MemclntKeepaliveReply) Reset()                        { *m = MemclntKeepaliveReply{} }
func (*MemclntKeepaliveReply) GetMessageName() string          { return "memclnt_keepalive_reply" }
func (*MemclntKeepaliveReply) GetCrcString() string            { return "e8d4e804" }
func (*MemclntKeepaliveReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *MemclntKeepaliveReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Retval
	size += 4
	return size
}
func (m *MemclntKeepaliveReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Retval
	o.PutUint32(buf[pos:pos+4], uint32(m.Retval))
	pos += 4
	return buf, nil
}
func (m *MemclntKeepaliveReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Retval
	m.Retval = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	return nil
}

// MemclntReadTimeout represents VPP binary API message 'memclnt_read_timeout'.
type MemclntReadTimeout struct {
	Dummy uint8 `binapi:"u8,name=dummy" json:"dummy,omitempty"`
}

func (m *MemclntReadTimeout) Reset()                        { *m = MemclntReadTimeout{} }
func (*MemclntReadTimeout) GetMessageName() string          { return "memclnt_read_timeout" }
func (*MemclntReadTimeout) GetCrcString() string            { return "c3a3a452" }
func (*MemclntReadTimeout) GetMessageType() api.MessageType { return api.OtherMessage }

func (m *MemclntReadTimeout) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Dummy
	size += 1
	return size
}
func (m *MemclntReadTimeout) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Dummy
	buf[pos] = uint8(m.Dummy)
	pos += 1
	return buf, nil
}
func (m *MemclntReadTimeout) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Dummy
	m.Dummy = uint8(tmp[pos])
	pos += 1
	return nil
}

// MemclntRxThreadSuspend represents VPP binary API message 'memclnt_rx_thread_suspend'.
type MemclntRxThreadSuspend struct {
	Dummy uint8 `binapi:"u8,name=dummy" json:"dummy,omitempty"`
}

func (m *MemclntRxThreadSuspend) Reset()                        { *m = MemclntRxThreadSuspend{} }
func (*MemclntRxThreadSuspend) GetMessageName() string          { return "memclnt_rx_thread_suspend" }
func (*MemclntRxThreadSuspend) GetCrcString() string            { return "c3a3a452" }
func (*MemclntRxThreadSuspend) GetMessageType() api.MessageType { return api.OtherMessage }

func (m *MemclntRxThreadSuspend) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Dummy
	size += 1
	return size
}
func (m *MemclntRxThreadSuspend) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Dummy
	buf[pos] = uint8(m.Dummy)
	pos += 1
	return buf, nil
}
func (m *MemclntRxThreadSuspend) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Dummy
	m.Dummy = uint8(tmp[pos])
	pos += 1
	return nil
}

// RPCCall represents VPP binary API message 'rpc_call'.
type RPCCall struct {
	Function        uint64 `binapi:"u64,name=function" json:"function,omitempty"`
	Multicast       uint8  `binapi:"u8,name=multicast" json:"multicast,omitempty"`
	NeedBarrierSync uint8  `binapi:"u8,name=need_barrier_sync" json:"need_barrier_sync,omitempty"`
	SendReply       uint8  `binapi:"u8,name=send_reply" json:"send_reply,omitempty"`
	DataLen         uint32 `binapi:"u32,name=data_len" json:"data_len,omitempty" struc:"sizeof=Data"`
	Data            []byte `binapi:"u8[data_len],name=data" json:"data,omitempty"`
}

func (m *RPCCall) Reset()                        { *m = RPCCall{} }
func (*RPCCall) GetMessageName() string          { return "rpc_call" }
func (*RPCCall) GetCrcString() string            { return "7e8a2c95" }
func (*RPCCall) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *RPCCall) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Function
	size += 8
	// field[1] m.Multicast
	size += 1
	// field[1] m.NeedBarrierSync
	size += 1
	// field[1] m.SendReply
	size += 1
	// field[1] m.DataLen
	size += 4
	// field[1] m.Data
	size += 1 * len(m.Data)
	return size
}
func (m *RPCCall) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Function
	o.PutUint64(buf[pos:pos+8], uint64(m.Function))
	pos += 8
	// field[1] m.Multicast
	buf[pos] = uint8(m.Multicast)
	pos += 1
	// field[1] m.NeedBarrierSync
	buf[pos] = uint8(m.NeedBarrierSync)
	pos += 1
	// field[1] m.SendReply
	buf[pos] = uint8(m.SendReply)
	pos += 1
	// field[1] m.DataLen
	o.PutUint32(buf[pos:pos+4], uint32(len(m.Data)))
	pos += 4
	// field[1] m.Data
	for i := 0; i < len(m.Data); i++ {
		var x uint8
		if i < len(m.Data) {
			x = uint8(m.Data[i])
		}
		buf[pos] = uint8(x)
		pos += 1
	}
	return buf, nil
}
func (m *RPCCall) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Function
	m.Function = uint64(o.Uint64(tmp[pos : pos+8]))
	pos += 8
	// field[1] m.Multicast
	m.Multicast = uint8(tmp[pos])
	pos += 1
	// field[1] m.NeedBarrierSync
	m.NeedBarrierSync = uint8(tmp[pos])
	pos += 1
	// field[1] m.SendReply
	m.SendReply = uint8(tmp[pos])
	pos += 1
	// field[1] m.DataLen
	m.DataLen = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Data
	m.Data = make([]uint8, m.DataLen)
	for i := 0; i < len(m.Data); i++ {
		m.Data[i] = uint8(tmp[pos])
		pos += 1
	}
	return nil
}

// RPCCallReply represents VPP binary API message 'rpc_call_reply'.
type RPCCallReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *RPCCallReply) Reset()                        { *m = RPCCallReply{} }
func (*RPCCallReply) GetMessageName() string          { return "rpc_call_reply" }
func (*RPCCallReply) GetCrcString() string            { return "e8d4e804" }
func (*RPCCallReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *RPCCallReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Retval
	size += 4
	return size
}
func (m *RPCCallReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Retval
	o.PutUint32(buf[pos:pos+4], uint32(m.Retval))
	pos += 4
	return buf, nil
}
func (m *RPCCallReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Retval
	m.Retval = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	return nil
}

// RxThreadExit represents VPP binary API message 'rx_thread_exit'.
type RxThreadExit struct {
	Dummy uint8 `binapi:"u8,name=dummy" json:"dummy,omitempty"`
}

func (m *RxThreadExit) Reset()                        { *m = RxThreadExit{} }
func (*RxThreadExit) GetMessageName() string          { return "rx_thread_exit" }
func (*RxThreadExit) GetCrcString() string            { return "c3a3a452" }
func (*RxThreadExit) GetMessageType() api.MessageType { return api.OtherMessage }

func (m *RxThreadExit) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Dummy
	size += 1
	return size
}
func (m *RxThreadExit) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Dummy
	buf[pos] = uint8(m.Dummy)
	pos += 1
	return buf, nil
}
func (m *RxThreadExit) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Dummy
	m.Dummy = uint8(tmp[pos])
	pos += 1
	return nil
}

// SockInitShm represents VPP binary API message 'sock_init_shm'.
type SockInitShm struct {
	RequestedSize uint32   `binapi:"u32,name=requested_size" json:"requested_size,omitempty"`
	Nitems        uint8    `binapi:"u8,name=nitems" json:"nitems,omitempty" struc:"sizeof=Configs"`
	Configs       []uint64 `binapi:"u64[nitems],name=configs" json:"configs,omitempty"`
}

func (m *SockInitShm) Reset()                        { *m = SockInitShm{} }
func (*SockInitShm) GetMessageName() string          { return "sock_init_shm" }
func (*SockInitShm) GetCrcString() string            { return "51646d92" }
func (*SockInitShm) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *SockInitShm) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.RequestedSize
	size += 4
	// field[1] m.Nitems
	size += 1
	// field[1] m.Configs
	size += 8 * len(m.Configs)
	return size
}
func (m *SockInitShm) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.RequestedSize
	o.PutUint32(buf[pos:pos+4], uint32(m.RequestedSize))
	pos += 4
	// field[1] m.Nitems
	buf[pos] = uint8(len(m.Configs))
	pos += 1
	// field[1] m.Configs
	for i := 0; i < len(m.Configs); i++ {
		var x uint64
		if i < len(m.Configs) {
			x = uint64(m.Configs[i])
		}
		o.PutUint64(buf[pos:pos+8], uint64(x))
		pos += 8
	}
	return buf, nil
}
func (m *SockInitShm) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.RequestedSize
	m.RequestedSize = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Nitems
	m.Nitems = uint8(tmp[pos])
	pos += 1
	// field[1] m.Configs
	m.Configs = make([]uint64, m.Nitems)
	for i := 0; i < len(m.Configs); i++ {
		m.Configs[i] = uint64(o.Uint64(tmp[pos : pos+8]))
		pos += 8
	}
	return nil
}

// SockInitShmReply represents VPP binary API message 'sock_init_shm_reply'.
type SockInitShmReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *SockInitShmReply) Reset()                        { *m = SockInitShmReply{} }
func (*SockInitShmReply) GetMessageName() string          { return "sock_init_shm_reply" }
func (*SockInitShmReply) GetCrcString() string            { return "e8d4e804" }
func (*SockInitShmReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *SockInitShmReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Retval
	size += 4
	return size
}
func (m *SockInitShmReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Retval
	o.PutUint32(buf[pos:pos+4], uint32(m.Retval))
	pos += 4
	return buf, nil
}
func (m *SockInitShmReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Retval
	m.Retval = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	return nil
}

// SockclntCreate represents VPP binary API message 'sockclnt_create'.
type SockclntCreate struct {
	Name string `binapi:"string[64],name=name" json:"name,omitempty" struc:"[64]byte"`
}

func (m *SockclntCreate) Reset()                        { *m = SockclntCreate{} }
func (*SockclntCreate) GetMessageName() string          { return "sockclnt_create" }
func (*SockclntCreate) GetCrcString() string            { return "455fb9c4" }
func (*SockclntCreate) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *SockclntCreate) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Name
	size += 64
	return size
}
func (m *SockclntCreate) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Name
	copy(buf[pos:pos+64], m.Name)
	pos += 64
	return buf, nil
}
func (m *SockclntCreate) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Name
	{
		nul := bytes.Index(tmp[pos:pos+64], []byte{0x00})
		m.Name = codec.DecodeString(tmp[pos : pos+nul])
		pos += 64
	}
	return nil
}

// SockclntCreateReply represents VPP binary API message 'sockclnt_create_reply'.
type SockclntCreateReply struct {
	Response     int32               `binapi:"i32,name=response" json:"response,omitempty"`
	Index        uint32              `binapi:"u32,name=index" json:"index,omitempty"`
	Count        uint16              `binapi:"u16,name=count" json:"count,omitempty" struc:"sizeof=MessageTable"`
	MessageTable []MessageTableEntry `binapi:"message_table_entry[count],name=message_table" json:"message_table,omitempty"`
}

func (m *SockclntCreateReply) Reset()                        { *m = SockclntCreateReply{} }
func (*SockclntCreateReply) GetMessageName() string          { return "sockclnt_create_reply" }
func (*SockclntCreateReply) GetCrcString() string            { return "35166268" }
func (*SockclntCreateReply) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *SockclntCreateReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Response
	size += 4
	// field[1] m.Index
	size += 4
	// field[1] m.Count
	size += 2
	// field[1] m.MessageTable
	for j1 := 0; j1 < len(m.MessageTable); j1++ {
		var s1 MessageTableEntry
		_ = s1
		if j1 < len(m.MessageTable) {
			s1 = m.MessageTable[j1]
		}
		// field[2] s1.Index
		size += 2
		// field[2] s1.Name
		size += 64
	}
	return size
}
func (m *SockclntCreateReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Response
	o.PutUint32(buf[pos:pos+4], uint32(m.Response))
	pos += 4
	// field[1] m.Index
	o.PutUint32(buf[pos:pos+4], uint32(m.Index))
	pos += 4
	// field[1] m.Count
	o.PutUint16(buf[pos:pos+2], uint16(len(m.MessageTable)))
	pos += 2
	// field[1] m.MessageTable
	for j1 := 0; j1 < len(m.MessageTable); j1++ {
		var v1 MessageTableEntry
		if j1 < len(m.MessageTable) {
			v1 = m.MessageTable[j1]
		}
		// field[2] v1.Index
		o.PutUint16(buf[pos:pos+2], uint16(v1.Index))
		pos += 2
		// field[2] v1.Name
		copy(buf[pos:pos+64], v1.Name)
		pos += 64
	}
	return buf, nil
}
func (m *SockclntCreateReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Response
	m.Response = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Index
	m.Index = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	// field[1] m.Count
	m.Count = uint16(o.Uint16(tmp[pos : pos+2]))
	pos += 2
	// field[1] m.MessageTable
	m.MessageTable = make([]MessageTableEntry, int(m.Count))
	for j1 := 0; j1 < int(m.Count); j1++ {
		// field[2] m.MessageTable[j1].Index
		m.MessageTable[j1].Index = uint16(o.Uint16(tmp[pos : pos+2]))
		pos += 2
		// field[2] m.MessageTable[j1].Name
		{
			nul := bytes.Index(tmp[pos:pos+64], []byte{0x00})
			m.MessageTable[j1].Name = codec.DecodeString(tmp[pos : pos+nul])
			pos += 64
		}
	}
	return nil
}

// SockclntDelete represents VPP binary API message 'sockclnt_delete'.
type SockclntDelete struct {
	Index uint32 `binapi:"u32,name=index" json:"index,omitempty"`
}

func (m *SockclntDelete) Reset()                        { *m = SockclntDelete{} }
func (*SockclntDelete) GetMessageName() string          { return "sockclnt_delete" }
func (*SockclntDelete) GetCrcString() string            { return "8ac76db6" }
func (*SockclntDelete) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *SockclntDelete) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Index
	size += 4
	return size
}
func (m *SockclntDelete) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Index
	o.PutUint32(buf[pos:pos+4], uint32(m.Index))
	pos += 4
	return buf, nil
}
func (m *SockclntDelete) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Index
	m.Index = uint32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	return nil
}

// SockclntDeleteReply represents VPP binary API message 'sockclnt_delete_reply'.
type SockclntDeleteReply struct {
	Response int32 `binapi:"i32,name=response" json:"response,omitempty"`
}

func (m *SockclntDeleteReply) Reset()                        { *m = SockclntDeleteReply{} }
func (*SockclntDeleteReply) GetMessageName() string          { return "sockclnt_delete_reply" }
func (*SockclntDeleteReply) GetCrcString() string            { return "8f38b1ee" }
func (*SockclntDeleteReply) GetMessageType() api.MessageType { return api.ReplyMessage }

func (m *SockclntDeleteReply) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.Response
	size += 4
	return size
}
func (m *SockclntDeleteReply) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.Response
	o.PutUint32(buf[pos:pos+4], uint32(m.Response))
	pos += 4
	return buf, nil
}
func (m *SockclntDeleteReply) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.Response
	m.Response = int32(o.Uint32(tmp[pos : pos+4]))
	pos += 4
	return nil
}

// TracePluginMsgIds represents VPP binary API message 'trace_plugin_msg_ids'.
type TracePluginMsgIds struct {
	PluginName string `binapi:"string[128],name=plugin_name" json:"plugin_name,omitempty" struc:"[128]byte"`
	FirstMsgID uint16 `binapi:"u16,name=first_msg_id" json:"first_msg_id,omitempty"`
	LastMsgID  uint16 `binapi:"u16,name=last_msg_id" json:"last_msg_id,omitempty"`
}

func (m *TracePluginMsgIds) Reset()                        { *m = TracePluginMsgIds{} }
func (*TracePluginMsgIds) GetMessageName() string          { return "trace_plugin_msg_ids" }
func (*TracePluginMsgIds) GetCrcString() string            { return "f476d3ce" }
func (*TracePluginMsgIds) GetMessageType() api.MessageType { return api.RequestMessage }

func (m *TracePluginMsgIds) Size() int {
	if m == nil {
		return 0
	}
	var size int
	// field[1] m.PluginName
	size += 128
	// field[1] m.FirstMsgID
	size += 2
	// field[1] m.LastMsgID
	size += 2
	return size
}
func (m *TracePluginMsgIds) Marshal(b []byte) ([]byte, error) {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	var buf []byte
	if b == nil {
		buf = make([]byte, m.Size())
	} else {
		buf = b
	}
	// field[1] m.PluginName
	copy(buf[pos:pos+128], m.PluginName)
	pos += 128
	// field[1] m.FirstMsgID
	o.PutUint16(buf[pos:pos+2], uint16(m.FirstMsgID))
	pos += 2
	// field[1] m.LastMsgID
	o.PutUint16(buf[pos:pos+2], uint16(m.LastMsgID))
	pos += 2
	return buf, nil
}
func (m *TracePluginMsgIds) Unmarshal(tmp []byte) error {
	o := binary.BigEndian
	_ = o
	pos := 0
	_ = pos
	// field[1] m.PluginName
	{
		nul := bytes.Index(tmp[pos:pos+128], []byte{0x00})
		m.PluginName = codec.DecodeString(tmp[pos : pos+nul])
		pos += 128
	}
	// field[1] m.FirstMsgID
	m.FirstMsgID = uint16(o.Uint16(tmp[pos : pos+2]))
	pos += 2
	// field[1] m.LastMsgID
	m.LastMsgID = uint16(o.Uint16(tmp[pos : pos+2]))
	pos += 2
	return nil
}

func init() { file_memclnt_binapi_init() }
func file_memclnt_binapi_init() {
	api.RegisterMessage((*APIVersions)(nil), "memclnt.APIVersions")
	api.RegisterMessage((*APIVersionsReply)(nil), "memclnt.APIVersionsReply")
	api.RegisterMessage((*GetFirstMsgID)(nil), "memclnt.GetFirstMsgID")
	api.RegisterMessage((*GetFirstMsgIDReply)(nil), "memclnt.GetFirstMsgIDReply")
	api.RegisterMessage((*MemclntCreate)(nil), "memclnt.MemclntCreate")
	api.RegisterMessage((*MemclntCreateReply)(nil), "memclnt.MemclntCreateReply")
	api.RegisterMessage((*MemclntDelete)(nil), "memclnt.MemclntDelete")
	api.RegisterMessage((*MemclntDeleteReply)(nil), "memclnt.MemclntDeleteReply")
	api.RegisterMessage((*MemclntKeepalive)(nil), "memclnt.MemclntKeepalive")
	api.RegisterMessage((*MemclntKeepaliveReply)(nil), "memclnt.MemclntKeepaliveReply")
	api.RegisterMessage((*MemclntReadTimeout)(nil), "memclnt.MemclntReadTimeout")
	api.RegisterMessage((*MemclntRxThreadSuspend)(nil), "memclnt.MemclntRxThreadSuspend")
	api.RegisterMessage((*RPCCall)(nil), "memclnt.RPCCall")
	api.RegisterMessage((*RPCCallReply)(nil), "memclnt.RPCCallReply")
	api.RegisterMessage((*RxThreadExit)(nil), "memclnt.RxThreadExit")
	api.RegisterMessage((*SockInitShm)(nil), "memclnt.SockInitShm")
	api.RegisterMessage((*SockInitShmReply)(nil), "memclnt.SockInitShmReply")
	api.RegisterMessage((*SockclntCreate)(nil), "memclnt.SockclntCreate")
	api.RegisterMessage((*SockclntCreateReply)(nil), "memclnt.SockclntCreateReply")
	api.RegisterMessage((*SockclntDelete)(nil), "memclnt.SockclntDelete")
	api.RegisterMessage((*SockclntDeleteReply)(nil), "memclnt.SockclntDeleteReply")
	api.RegisterMessage((*TracePluginMsgIds)(nil), "memclnt.TracePluginMsgIds")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*APIVersions)(nil),
		(*APIVersionsReply)(nil),
		(*GetFirstMsgID)(nil),
		(*GetFirstMsgIDReply)(nil),
		(*MemclntCreate)(nil),
		(*MemclntCreateReply)(nil),
		(*MemclntDelete)(nil),
		(*MemclntDeleteReply)(nil),
		(*MemclntKeepalive)(nil),
		(*MemclntKeepaliveReply)(nil),
		(*MemclntReadTimeout)(nil),
		(*MemclntRxThreadSuspend)(nil),
		(*RPCCall)(nil),
		(*RPCCallReply)(nil),
		(*RxThreadExit)(nil),
		(*SockInitShm)(nil),
		(*SockInitShmReply)(nil),
		(*SockclntCreate)(nil),
		(*SockclntCreateReply)(nil),
		(*SockclntDelete)(nil),
		(*SockclntDeleteReply)(nil),
		(*TracePluginMsgIds)(nil),
	}
}

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = codec.DecodeString
var _ = bytes.NewBuffer
var _ = context.Background
var _ = io.Copy
var _ = strconv.Itoa
var _ = struc.Pack
var _ = binary.BigEndian
var _ = math.Float32bits
