// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

package ip

import (
	"context"
	"io"

	api "git.fd.io/govpp.git/api"
)

// RPCService represents RPC service API for ip module.
type RPCService interface {
	DumpIPAddress(ctx context.Context, in *IPAddressDump) (RPCService_DumpIPAddressClient, error)
	DumpIPContainerProxy(ctx context.Context, in *IPContainerProxyDump) (RPCService_DumpIPContainerProxyClient, error)
	DumpIP(ctx context.Context, in *IPDump) (RPCService_DumpIPClient, error)
	DumpIPMroute(ctx context.Context, in *IPMrouteDump) (RPCService_DumpIPMrouteClient, error)
	DumpIPMtable(ctx context.Context, in *IPMtableDump) (RPCService_DumpIPMtableClient, error)
	DumpIPPuntRedirect(ctx context.Context, in *IPPuntRedirectDump) (RPCService_DumpIPPuntRedirectClient, error)
	DumpIPRoute(ctx context.Context, in *IPRouteDump) (RPCService_DumpIPRouteClient, error)
	DumpIPTable(ctx context.Context, in *IPTableDump) (RPCService_DumpIPTableClient, error)
	DumpIPUnnumbered(ctx context.Context, in *IPUnnumberedDump) (RPCService_DumpIPUnnumberedClient, error)
	DumpMfibSignal(ctx context.Context, in *MfibSignalDump) (RPCService_DumpMfibSignalClient, error)
	IoamDisable(ctx context.Context, in *IoamDisable) (*IoamDisableReply, error)
	IoamEnable(ctx context.Context, in *IoamEnable) (*IoamEnableReply, error)
	IPContainerProxyAddDel(ctx context.Context, in *IPContainerProxyAddDel) (*IPContainerProxyAddDelReply, error)
	IPMrouteAddDel(ctx context.Context, in *IPMrouteAddDel) (*IPMrouteAddDelReply, error)
	IPPuntPolice(ctx context.Context, in *IPPuntPolice) (*IPPuntPoliceReply, error)
	IPPuntRedirect(ctx context.Context, in *IPPuntRedirect) (*IPPuntRedirectReply, error)
	IPReassemblyEnableDisable(ctx context.Context, in *IPReassemblyEnableDisable) (*IPReassemblyEnableDisableReply, error)
	IPReassemblyGet(ctx context.Context, in *IPReassemblyGet) (*IPReassemblyGetReply, error)
	IPReassemblySet(ctx context.Context, in *IPReassemblySet) (*IPReassemblySetReply, error)
	IPRouteAddDel(ctx context.Context, in *IPRouteAddDel) (*IPRouteAddDelReply, error)
	IPSourceAndPortRangeCheckAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckAddDel) (*IPSourceAndPortRangeCheckAddDelReply, error)
	IPSourceAndPortRangeCheckInterfaceAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckInterfaceAddDel) (*IPSourceAndPortRangeCheckInterfaceAddDelReply, error)
	IPSourceCheckInterfaceAddDel(ctx context.Context, in *IPSourceCheckInterfaceAddDel) (*IPSourceCheckInterfaceAddDelReply, error)
	IPTableAddDel(ctx context.Context, in *IPTableAddDel) (*IPTableAddDelReply, error)
	IPTableFlush(ctx context.Context, in *IPTableFlush) (*IPTableFlushReply, error)
	IPTableReplaceBegin(ctx context.Context, in *IPTableReplaceBegin) (*IPTableReplaceBeginReply, error)
	IPTableReplaceEnd(ctx context.Context, in *IPTableReplaceEnd) (*IPTableReplaceEndReply, error)
	SetIPFlowHash(ctx context.Context, in *SetIPFlowHash) (*SetIPFlowHashReply, error)
	SwInterfaceIP6EnableDisable(ctx context.Context, in *SwInterfaceIP6EnableDisable) (*SwInterfaceIP6EnableDisableReply, error)
	SwInterfaceIP6SetLinkLocalAddress(ctx context.Context, in *SwInterfaceIP6SetLinkLocalAddress) (*SwInterfaceIP6SetLinkLocalAddressReply, error)
}

type serviceClient struct {
	ch api.Channel
}

func NewServiceClient(ch api.Channel) RPCService {
	return &serviceClient{ch}
}

func (c *serviceClient) DumpIPAddress(ctx context.Context, in *IPAddressDump) (RPCService_DumpIPAddressClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPAddressClient{stream}
	return x, nil
}

type RPCService_DumpIPAddressClient interface {
	Recv() (*IPAddressDetails, error)
}

type serviceClient_DumpIPAddressClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPAddressClient) Recv() (*IPAddressDetails, error) {
	m := new(IPAddressDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPContainerProxy(ctx context.Context, in *IPContainerProxyDump) (RPCService_DumpIPContainerProxyClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPContainerProxyClient{stream}
	return x, nil
}

type RPCService_DumpIPContainerProxyClient interface {
	Recv() (*IPContainerProxyDetails, error)
}

type serviceClient_DumpIPContainerProxyClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPContainerProxyClient) Recv() (*IPContainerProxyDetails, error) {
	m := new(IPContainerProxyDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIP(ctx context.Context, in *IPDump) (RPCService_DumpIPClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPClient{stream}
	return x, nil
}

type RPCService_DumpIPClient interface {
	Recv() (*IPDetails, error)
}

type serviceClient_DumpIPClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPClient) Recv() (*IPDetails, error) {
	m := new(IPDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPMroute(ctx context.Context, in *IPMrouteDump) (RPCService_DumpIPMrouteClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPMrouteClient{stream}
	return x, nil
}

type RPCService_DumpIPMrouteClient interface {
	Recv() (*IPMrouteDetails, error)
}

type serviceClient_DumpIPMrouteClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPMrouteClient) Recv() (*IPMrouteDetails, error) {
	m := new(IPMrouteDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPMtable(ctx context.Context, in *IPMtableDump) (RPCService_DumpIPMtableClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPMtableClient{stream}
	return x, nil
}

type RPCService_DumpIPMtableClient interface {
	Recv() (*IPMtableDetails, error)
}

type serviceClient_DumpIPMtableClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPMtableClient) Recv() (*IPMtableDetails, error) {
	m := new(IPMtableDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPPuntRedirect(ctx context.Context, in *IPPuntRedirectDump) (RPCService_DumpIPPuntRedirectClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPPuntRedirectClient{stream}
	return x, nil
}

type RPCService_DumpIPPuntRedirectClient interface {
	Recv() (*IPPuntRedirectDetails, error)
}

type serviceClient_DumpIPPuntRedirectClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPPuntRedirectClient) Recv() (*IPPuntRedirectDetails, error) {
	m := new(IPPuntRedirectDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPRoute(ctx context.Context, in *IPRouteDump) (RPCService_DumpIPRouteClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPRouteClient{stream}
	return x, nil
}

type RPCService_DumpIPRouteClient interface {
	Recv() (*IPRouteDetails, error)
}

type serviceClient_DumpIPRouteClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPRouteClient) Recv() (*IPRouteDetails, error) {
	m := new(IPRouteDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPTable(ctx context.Context, in *IPTableDump) (RPCService_DumpIPTableClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPTableClient{stream}
	return x, nil
}

type RPCService_DumpIPTableClient interface {
	Recv() (*IPTableDetails, error)
}

type serviceClient_DumpIPTableClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPTableClient) Recv() (*IPTableDetails, error) {
	m := new(IPTableDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpIPUnnumbered(ctx context.Context, in *IPUnnumberedDump) (RPCService_DumpIPUnnumberedClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpIPUnnumberedClient{stream}
	return x, nil
}

type RPCService_DumpIPUnnumberedClient interface {
	Recv() (*IPUnnumberedDetails, error)
}

type serviceClient_DumpIPUnnumberedClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpIPUnnumberedClient) Recv() (*IPUnnumberedDetails, error) {
	m := new(IPUnnumberedDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) DumpMfibSignal(ctx context.Context, in *MfibSignalDump) (RPCService_DumpMfibSignalClient, error) {
	stream := c.ch.SendMultiRequest(in)
	x := &serviceClient_DumpMfibSignalClient{stream}
	return x, nil
}

type RPCService_DumpMfibSignalClient interface {
	Recv() (*MfibSignalDetails, error)
}

type serviceClient_DumpMfibSignalClient struct {
	api.MultiRequestCtx
}

func (c *serviceClient_DumpMfibSignalClient) Recv() (*MfibSignalDetails, error) {
	m := new(MfibSignalDetails)
	stop, err := c.MultiRequestCtx.ReceiveReply(m)
	if err != nil {
		return nil, err
	}
	if stop {
		return nil, io.EOF
	}
	return m, nil
}

func (c *serviceClient) IoamDisable(ctx context.Context, in *IoamDisable) (*IoamDisableReply, error) {
	out := new(IoamDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IoamEnable(ctx context.Context, in *IoamEnable) (*IoamEnableReply, error) {
	out := new(IoamEnableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPContainerProxyAddDel(ctx context.Context, in *IPContainerProxyAddDel) (*IPContainerProxyAddDelReply, error) {
	out := new(IPContainerProxyAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPMrouteAddDel(ctx context.Context, in *IPMrouteAddDel) (*IPMrouteAddDelReply, error) {
	out := new(IPMrouteAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPPuntPolice(ctx context.Context, in *IPPuntPolice) (*IPPuntPoliceReply, error) {
	out := new(IPPuntPoliceReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPPuntRedirect(ctx context.Context, in *IPPuntRedirect) (*IPPuntRedirectReply, error) {
	out := new(IPPuntRedirectReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPReassemblyEnableDisable(ctx context.Context, in *IPReassemblyEnableDisable) (*IPReassemblyEnableDisableReply, error) {
	out := new(IPReassemblyEnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPReassemblyGet(ctx context.Context, in *IPReassemblyGet) (*IPReassemblyGetReply, error) {
	out := new(IPReassemblyGetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPReassemblySet(ctx context.Context, in *IPReassemblySet) (*IPReassemblySetReply, error) {
	out := new(IPReassemblySetReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPRouteAddDel(ctx context.Context, in *IPRouteAddDel) (*IPRouteAddDelReply, error) {
	out := new(IPRouteAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPSourceAndPortRangeCheckAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckAddDel) (*IPSourceAndPortRangeCheckAddDelReply, error) {
	out := new(IPSourceAndPortRangeCheckAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPSourceAndPortRangeCheckInterfaceAddDel(ctx context.Context, in *IPSourceAndPortRangeCheckInterfaceAddDel) (*IPSourceAndPortRangeCheckInterfaceAddDelReply, error) {
	out := new(IPSourceAndPortRangeCheckInterfaceAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPSourceCheckInterfaceAddDel(ctx context.Context, in *IPSourceCheckInterfaceAddDel) (*IPSourceCheckInterfaceAddDelReply, error) {
	out := new(IPSourceCheckInterfaceAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableAddDel(ctx context.Context, in *IPTableAddDel) (*IPTableAddDelReply, error) {
	out := new(IPTableAddDelReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableFlush(ctx context.Context, in *IPTableFlush) (*IPTableFlushReply, error) {
	out := new(IPTableFlushReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableReplaceBegin(ctx context.Context, in *IPTableReplaceBegin) (*IPTableReplaceBeginReply, error) {
	out := new(IPTableReplaceBeginReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) IPTableReplaceEnd(ctx context.Context, in *IPTableReplaceEnd) (*IPTableReplaceEndReply, error) {
	out := new(IPTableReplaceEndReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SetIPFlowHash(ctx context.Context, in *SetIPFlowHash) (*SetIPFlowHashReply, error) {
	out := new(SetIPFlowHashReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceIP6EnableDisable(ctx context.Context, in *SwInterfaceIP6EnableDisable) (*SwInterfaceIP6EnableDisableReply, error) {
	out := new(SwInterfaceIP6EnableDisableReply)
	err := c.ch.SendRequest(in).ReceiveReply(out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) SwInterfaceIP6SetLinkLocalAddress(ctx context.Context, in *SwInterfaceIP6SetLinkLocalAddress) (*SwInterfaceIP6SetLinkLocalAddressReply, error) {
	out := new(SwInterfaceIP6SetLinkLocalAddressReply)
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
