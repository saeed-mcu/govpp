// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package acl

import (
	"context"
	"io"

	api "git.fd.io/govpp.git/api"
)

// RPCService represents RPC service API for acl module.
type RPCService interface {
	DumpACL(ctx context.Context, in *ACLDump) (RPCService_DumpACLClient, error)
	DumpACLInterfaceEtypeWhitelist(ctx context.Context, in *ACLInterfaceEtypeWhitelistDump) (RPCService_DumpACLInterfaceEtypeWhitelistClient, error)
	DumpACLInterfaceList(ctx context.Context, in *ACLInterfaceListDump) (RPCService_DumpACLInterfaceListClient, error)
	DumpMacipACL(ctx context.Context, in *MacipACLDump) (RPCService_DumpMacipACLClient, error)
	DumpMacipACLInterfaceList(ctx context.Context, in *MacipACLInterfaceListDump) (RPCService_DumpMacipACLInterfaceListClient, error)
	ACLAddReplace(ctx context.Context, in *ACLAddReplace) (*ACLAddReplaceReply, error)
	ACLDel(ctx context.Context, in *ACLDel) (*ACLDelReply, error)
	ACLInterfaceAddDel(ctx context.Context, in *ACLInterfaceAddDel) (*ACLInterfaceAddDelReply, error)
	ACLInterfaceSetACLList(ctx context.Context, in *ACLInterfaceSetACLList) (*ACLInterfaceSetACLListReply, error)
	ACLInterfaceSetEtypeWhitelist(ctx context.Context, in *ACLInterfaceSetEtypeWhitelist) (*ACLInterfaceSetEtypeWhitelistReply, error)
	ACLPluginControlPing(ctx context.Context, in *ACLPluginControlPing) (*ACLPluginControlPingReply, error)
	ACLPluginGetConnTableMaxEntries(ctx context.Context, in *ACLPluginGetConnTableMaxEntries) (*ACLPluginGetConnTableMaxEntriesReply, error)
	ACLPluginGetVersion(ctx context.Context, in *ACLPluginGetVersion) (*ACLPluginGetVersionReply, error)
	ACLStatsIntfCountersEnable(ctx context.Context, in *ACLStatsIntfCountersEnable) (*ACLStatsIntfCountersEnableReply, error)
	MacipACLAdd(ctx context.Context, in *MacipACLAdd) (*MacipACLAddReply, error)
	MacipACLAddReplace(ctx context.Context, in *MacipACLAddReplace) (*MacipACLAddReplaceReply, error)
	MacipACLDel(ctx context.Context, in *MacipACLDel) (*MacipACLDelReply, error)
	MacipACLInterfaceAddDel(ctx context.Context, in *MacipACLInterfaceAddDel) (*MacipACLInterfaceAddDelReply, error)
	MacipACLInterfaceGet(ctx context.Context, in *MacipACLInterfaceGet) (*MacipACLInterfaceGetReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpACL(ctx context.Context, in *ACLDump) (RPCService_DumpACLClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpACLClient{stream}
	return x, nil
}

type RPCService_DumpACLClient interface {
	Recv() (*ACLDetails, error)
}

type serviceClient_DumpACLClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpACLClient) Recv() (*ACLDetails, error) {
	m := new(ACLDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpACLInterfaceEtypeWhitelist(ctx context.Context, in *ACLInterfaceEtypeWhitelistDump) (RPCService_DumpACLInterfaceEtypeWhitelistClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpACLInterfaceEtypeWhitelistClient{stream}
	return x, nil
}

type RPCService_DumpACLInterfaceEtypeWhitelistClient interface {
	Recv() (*ACLInterfaceEtypeWhitelistDetails, error)
}

type serviceClient_DumpACLInterfaceEtypeWhitelistClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpACLInterfaceEtypeWhitelistClient) Recv() (*ACLInterfaceEtypeWhitelistDetails, error) {
	m := new(ACLInterfaceEtypeWhitelistDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpACLInterfaceList(ctx context.Context, in *ACLInterfaceListDump) (RPCService_DumpACLInterfaceListClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpACLInterfaceListClient{stream}
	return x, nil
}

type RPCService_DumpACLInterfaceListClient interface {
	Recv() (*ACLInterfaceListDetails, error)
}

type serviceClient_DumpACLInterfaceListClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpACLInterfaceListClient) Recv() (*ACLInterfaceListDetails, error) {
	m := new(ACLInterfaceListDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMacipACL(ctx context.Context, in *MacipACLDump) (RPCService_DumpMacipACLClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMacipACLClient{stream}
	return x, nil
}

type RPCService_DumpMacipACLClient interface {
	Recv() (*MacipACLDetails, error)
}

type serviceClient_DumpMacipACLClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMacipACLClient) Recv() (*MacipACLDetails, error) {
	m := new(MacipACLDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMacipACLInterfaceList(ctx context.Context, in *MacipACLInterfaceListDump) (RPCService_DumpMacipACLInterfaceListClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMacipACLInterfaceListClient{stream}
	return x, nil
}

type RPCService_DumpMacipACLInterfaceListClient interface {
	Recv() (*MacipACLInterfaceListDetails, error)
}

type serviceClient_DumpMacipACLInterfaceListClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMacipACLInterfaceListClient) Recv() (*MacipACLInterfaceListDetails, error) {
	m := new(MacipACLInterfaceListDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) ACLAddReplace(ctx context.Context, in *ACLAddReplace) (*ACLAddReplaceReply, error) {
	out := new(ACLAddReplaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLDel(ctx context.Context, in *ACLDel) (*ACLDelReply, error) {
	out := new(ACLDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLInterfaceAddDel(ctx context.Context, in *ACLInterfaceAddDel) (*ACLInterfaceAddDelReply, error) {
	out := new(ACLInterfaceAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLInterfaceSetACLList(ctx context.Context, in *ACLInterfaceSetACLList) (*ACLInterfaceSetACLListReply, error) {
	out := new(ACLInterfaceSetACLListReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLInterfaceSetEtypeWhitelist(ctx context.Context, in *ACLInterfaceSetEtypeWhitelist) (*ACLInterfaceSetEtypeWhitelistReply, error) {
	out := new(ACLInterfaceSetEtypeWhitelistReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLPluginControlPing(ctx context.Context, in *ACLPluginControlPing) (*ACLPluginControlPingReply, error) {
	out := new(ACLPluginControlPingReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLPluginGetConnTableMaxEntries(ctx context.Context, in *ACLPluginGetConnTableMaxEntries) (*ACLPluginGetConnTableMaxEntriesReply, error) {
	out := new(ACLPluginGetConnTableMaxEntriesReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLPluginGetVersion(ctx context.Context, in *ACLPluginGetVersion) (*ACLPluginGetVersionReply, error) {
	out := new(ACLPluginGetVersionReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) ACLStatsIntfCountersEnable(ctx context.Context, in *ACLStatsIntfCountersEnable) (*ACLStatsIntfCountersEnableReply, error) {
	out := new(ACLStatsIntfCountersEnableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MacipACLAdd(ctx context.Context, in *MacipACLAdd) (*MacipACLAddReply, error) {
	out := new(MacipACLAddReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MacipACLAddReplace(ctx context.Context, in *MacipACLAddReplace) (*MacipACLAddReplaceReply, error) {
	out := new(MacipACLAddReplaceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MacipACLDel(ctx context.Context, in *MacipACLDel) (*MacipACLDelReply, error) {
	out := new(MacipACLDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MacipACLInterfaceAddDel(ctx context.Context, in *MacipACLInterfaceAddDel) (*MacipACLInterfaceAddDelReply, error) {
	out := new(MacipACLInterfaceAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) MacipACLInterfaceGet(ctx context.Context, in *MacipACLInterfaceGet) (*MacipACLInterfaceGetReply, error) {
	out := new(MacipACLInterfaceGetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ = api.RegisterMessage
var _ = context.Background
var _ = io.Copy
