package codec_test

import (
	"bytes"
	"testing"

	"git.fd.io/govpp.git/api"
	"git.fd.io/govpp.git/codec"
	"git.fd.io/govpp.git/examples/binapi/ip"
	"git.fd.io/govpp.git/examples/binapi/vpe"
)

type MyMsg struct {
	Index uint16
	Label []byte `struc:"[16]byte"`
	Port  uint16
}

func (*MyMsg) GetMessageName() string {
	return "my_msg"
}
func (*MyMsg) GetCrcString() string {
	return "xxxxx"
}
func (*MyMsg) GetMessageType() api.MessageType {
	return api.OtherMessage
}

func TestEncode(t *testing.T) {
	tests := []struct {
		name    string
		msg     api.Message
		msgID   uint16
		expData []byte
	}{
		{name: "basic",
			msg:     &MyMsg{Index: 1, Label: []byte("Abcdef"), Port: 1000},
			msgID:   100,
			expData: []byte{0x00, 0x64, 0x00, 0x01, 0x41, 0x62, 0x63, 0x64, 0x65, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8},
		},
		{name: "show version",
			msg:     &vpe.ShowVersion{},
			msgID:   743,
			expData: []byte{0x02, 0xE7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		{name: "ip route",
			msg: &ip.IPRouteAddDel{
				IsAdd:       true,
				IsMultipath: true,
				Route: ip.IPRoute{
					TableID:    0,
					StatsIndex: 0,
					Prefix:     ip.Prefix{},
					NPaths:     0,
				},
			},
			msgID:   743,
			expData: []byte{0x02, 0xE7, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00},
		},
		/*{name: "sr",
			msg: &sr.SrPolicyAdd{
				BsidAddr: sr.IP6Address{},
				Weight:   0,
				IsEncap:  false,
				IsSpray:  false,
				FibTable: 0,
				Sids:     sr.Srv6SidList{},
			},
			msgID:   99,
			expData: []byte{0x00, 0x64, 0x00, 0x01, 0x41, 0x62, 0x63, 0x64, 0x65, 0x66, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x03, 0xE8},
		},*/
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := codec.DefaultCodec

			data, err := c.EncodeMsg(test.msg, test.msgID)
			if err != nil {
				t.Fatalf("expected nil error, got: %v", err)
			}
			if !bytes.Equal(data, test.expData) {
				t.Fatalf("expected data: % 0X, got: % 0X", test.expData, data)
			}
		})
	}
}

func TestEncodePanic(t *testing.T) {
	c := codec.DefaultCodec

	msg := &MyMsg{Index: 1, Label: []byte("thisIsLongerThan16Bytes"), Port: 1000}

	_, err := c.EncodeMsg(msg, 100)
	if err == nil {
		t.Fatalf("expected non-nil error, got: %v", err)
	}
}
