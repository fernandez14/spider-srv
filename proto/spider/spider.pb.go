// Code generated by protoc-gen-go.
// source: proto/spider/spider.proto
// DO NOT EDIT!

/*
Package spider is a generated protocol buffer package.

It is generated from these files:
	proto/spider/spider.proto

It has these top-level messages:
	Url
	FetchDatasetRequest
	FetchDatasetResponse
	FetchDatasetByRequest
	FetchDatasetByResponse
	TrackSitemapRequest
	TrackSitemapResponse
	TrackURLRequest
	TrackURLResponse
	PrepareDatasetsRequest
	PrepareDatasetsResponse
	FetchPagesRequest
	FetchPagesResponse
*/
package spider

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Url struct {
	Id        uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Url       string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
	Title     string `protobuf:"bytes,3,opt,name=title" json:"title,omitempty"`
	Group     string `protobuf:"bytes,4,opt,name=group" json:"group,omitempty"`
	SitemapId uint64 `protobuf:"varint,5,opt,name=sitemap_id,json=sitemapId" json:"sitemap_id,omitempty"`
}

func (m *Url) Reset()                    { *m = Url{} }
func (m *Url) String() string            { return proto.CompactTextString(m) }
func (*Url) ProtoMessage()               {}
func (*Url) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FetchDatasetRequest struct {
	Id    uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	UrlId uint64 `protobuf:"varint,2,opt,name=url_id,json=urlId" json:"url_id,omitempty"`
}

func (m *FetchDatasetRequest) Reset()                    { *m = FetchDatasetRequest{} }
func (m *FetchDatasetRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchDatasetRequest) ProtoMessage()               {}
func (*FetchDatasetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FetchDatasetResponse struct {
	Data map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FetchDatasetResponse) Reset()                    { *m = FetchDatasetResponse{} }
func (m *FetchDatasetResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchDatasetResponse) ProtoMessage()               {}
func (*FetchDatasetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchDatasetResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type FetchDatasetByRequest struct {
	SelectorId uint64   `protobuf:"varint,1,opt,name=selector_id,json=selectorId" json:"selector_id,omitempty"`
	Conditions []string `protobuf:"bytes,2,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *FetchDatasetByRequest) Reset()                    { *m = FetchDatasetByRequest{} }
func (m *FetchDatasetByRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchDatasetByRequest) ProtoMessage()               {}
func (*FetchDatasetByRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FetchDatasetByResponse struct {
	Data map[string]string `protobuf:"bytes,1,rep,name=data" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *FetchDatasetByResponse) Reset()                    { *m = FetchDatasetByResponse{} }
func (m *FetchDatasetByResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchDatasetByResponse) ProtoMessage()               {}
func (*FetchDatasetByResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FetchDatasetByResponse) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

type TrackSitemapRequest struct {
	From     string                         `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	Depth    uint64                         `protobuf:"varint,2,opt,name=depth" json:"depth,omitempty"`
	Patterns []*TrackSitemapRequest_Pattern `protobuf:"bytes,3,rep,name=patterns" json:"patterns,omitempty"`
	Name     string                         `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Strict   bool                           `protobuf:"varint,5,opt,name=strict" json:"strict,omitempty"`
}

func (m *TrackSitemapRequest) Reset()                    { *m = TrackSitemapRequest{} }
func (m *TrackSitemapRequest) String() string            { return proto.CompactTextString(m) }
func (*TrackSitemapRequest) ProtoMessage()               {}
func (*TrackSitemapRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *TrackSitemapRequest) GetPatterns() []*TrackSitemapRequest_Pattern {
	if m != nil {
		return m.Patterns
	}
	return nil
}

type TrackSitemapRequest_Pattern struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Matches string `protobuf:"bytes,2,opt,name=matches" json:"matches,omitempty"`
}

func (m *TrackSitemapRequest_Pattern) Reset()                    { *m = TrackSitemapRequest_Pattern{} }
func (m *TrackSitemapRequest_Pattern) String() string            { return proto.CompactTextString(m) }
func (*TrackSitemapRequest_Pattern) ProtoMessage()               {}
func (*TrackSitemapRequest_Pattern) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type TrackSitemapResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *TrackSitemapResponse) Reset()                    { *m = TrackSitemapResponse{} }
func (m *TrackSitemapResponse) String() string            { return proto.CompactTextString(m) }
func (*TrackSitemapResponse) ProtoMessage()               {}
func (*TrackSitemapResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type TrackURLRequest struct {
	Url   string `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	Group string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
}

func (m *TrackURLRequest) Reset()                    { *m = TrackURLRequest{} }
func (m *TrackURLRequest) String() string            { return proto.CompactTextString(m) }
func (*TrackURLRequest) ProtoMessage()               {}
func (*TrackURLRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type TrackURLResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *TrackURLResponse) Reset()                    { *m = TrackURLResponse{} }
func (m *TrackURLResponse) String() string            { return proto.CompactTextString(m) }
func (*TrackURLResponse) ProtoMessage()               {}
func (*TrackURLResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type PrepareDatasetsRequest struct {
	SelectorId uint64   `protobuf:"varint,1,opt,name=selector_id,json=selectorId" json:"selector_id,omitempty"`
	Group      string   `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
	Conditions []string `protobuf:"bytes,3,rep,name=conditions" json:"conditions,omitempty"`
}

func (m *PrepareDatasetsRequest) Reset()                    { *m = PrepareDatasetsRequest{} }
func (m *PrepareDatasetsRequest) String() string            { return proto.CompactTextString(m) }
func (*PrepareDatasetsRequest) ProtoMessage()               {}
func (*PrepareDatasetsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type PrepareDatasetsResponse struct {
	Count uint64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *PrepareDatasetsResponse) Reset()                    { *m = PrepareDatasetsResponse{} }
func (m *PrepareDatasetsResponse) String() string            { return proto.CompactTextString(m) }
func (*PrepareDatasetsResponse) ProtoMessage()               {}
func (*PrepareDatasetsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type FetchPagesRequest struct {
	Search string `protobuf:"bytes,1,opt,name=search" json:"search,omitempty"`
	Group  string `protobuf:"bytes,2,opt,name=group" json:"group,omitempty"`
}

func (m *FetchPagesRequest) Reset()                    { *m = FetchPagesRequest{} }
func (m *FetchPagesRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchPagesRequest) ProtoMessage()               {}
func (*FetchPagesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type FetchPagesResponse struct {
	Results []*Url `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *FetchPagesResponse) Reset()                    { *m = FetchPagesResponse{} }
func (m *FetchPagesResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchPagesResponse) ProtoMessage()               {}
func (*FetchPagesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FetchPagesResponse) GetResults() []*Url {
	if m != nil {
		return m.Results
	}
	return nil
}

func init() {
	proto.RegisterType((*Url)(nil), "Url")
	proto.RegisterType((*FetchDatasetRequest)(nil), "FetchDatasetRequest")
	proto.RegisterType((*FetchDatasetResponse)(nil), "FetchDatasetResponse")
	proto.RegisterType((*FetchDatasetByRequest)(nil), "FetchDatasetByRequest")
	proto.RegisterType((*FetchDatasetByResponse)(nil), "FetchDatasetByResponse")
	proto.RegisterType((*TrackSitemapRequest)(nil), "TrackSitemapRequest")
	proto.RegisterType((*TrackSitemapRequest_Pattern)(nil), "TrackSitemapRequest.Pattern")
	proto.RegisterType((*TrackSitemapResponse)(nil), "TrackSitemapResponse")
	proto.RegisterType((*TrackURLRequest)(nil), "TrackURLRequest")
	proto.RegisterType((*TrackURLResponse)(nil), "TrackURLResponse")
	proto.RegisterType((*PrepareDatasetsRequest)(nil), "PrepareDatasetsRequest")
	proto.RegisterType((*PrepareDatasetsResponse)(nil), "PrepareDatasetsResponse")
	proto.RegisterType((*FetchPagesRequest)(nil), "FetchPagesRequest")
	proto.RegisterType((*FetchPagesResponse)(nil), "FetchPagesResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Spider service

type SpiderClient interface {
	TrackSitemap(ctx context.Context, in *TrackSitemapRequest, opts ...client.CallOption) (*TrackSitemapResponse, error)
	TrackURL(ctx context.Context, in *TrackURLRequest, opts ...client.CallOption) (*TrackURLResponse, error)
	FetchDataset(ctx context.Context, in *FetchDatasetRequest, opts ...client.CallOption) (*FetchDatasetResponse, error)
	FetchDatasetBy(ctx context.Context, in *FetchDatasetByRequest, opts ...client.CallOption) (*FetchDatasetByResponse, error)
	PrepareDatasets(ctx context.Context, in *PrepareDatasetsRequest, opts ...client.CallOption) (*PrepareDatasetsResponse, error)
	FetchPages(ctx context.Context, in *FetchPagesRequest, opts ...client.CallOption) (*FetchPagesResponse, error)
}

type spiderClient struct {
	c           client.Client
	serviceName string
}

func NewSpiderClient(serviceName string, c client.Client) SpiderClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "spider"
	}
	return &spiderClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *spiderClient) TrackSitemap(ctx context.Context, in *TrackSitemapRequest, opts ...client.CallOption) (*TrackSitemapResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Spider.TrackSitemap", in)
	out := new(TrackSitemapResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderClient) TrackURL(ctx context.Context, in *TrackURLRequest, opts ...client.CallOption) (*TrackURLResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Spider.TrackURL", in)
	out := new(TrackURLResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderClient) FetchDataset(ctx context.Context, in *FetchDatasetRequest, opts ...client.CallOption) (*FetchDatasetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Spider.FetchDataset", in)
	out := new(FetchDatasetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderClient) FetchDatasetBy(ctx context.Context, in *FetchDatasetByRequest, opts ...client.CallOption) (*FetchDatasetByResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Spider.FetchDatasetBy", in)
	out := new(FetchDatasetByResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderClient) PrepareDatasets(ctx context.Context, in *PrepareDatasetsRequest, opts ...client.CallOption) (*PrepareDatasetsResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Spider.PrepareDatasets", in)
	out := new(PrepareDatasetsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spiderClient) FetchPages(ctx context.Context, in *FetchPagesRequest, opts ...client.CallOption) (*FetchPagesResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Spider.FetchPages", in)
	out := new(FetchPagesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Spider service

type SpiderHandler interface {
	TrackSitemap(context.Context, *TrackSitemapRequest, *TrackSitemapResponse) error
	TrackURL(context.Context, *TrackURLRequest, *TrackURLResponse) error
	FetchDataset(context.Context, *FetchDatasetRequest, *FetchDatasetResponse) error
	FetchDatasetBy(context.Context, *FetchDatasetByRequest, *FetchDatasetByResponse) error
	PrepareDatasets(context.Context, *PrepareDatasetsRequest, *PrepareDatasetsResponse) error
	FetchPages(context.Context, *FetchPagesRequest, *FetchPagesResponse) error
}

func RegisterSpiderHandler(s server.Server, hdlr SpiderHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Spider{hdlr}, opts...))
}

type Spider struct {
	SpiderHandler
}

func (h *Spider) TrackSitemap(ctx context.Context, in *TrackSitemapRequest, out *TrackSitemapResponse) error {
	return h.SpiderHandler.TrackSitemap(ctx, in, out)
}

func (h *Spider) TrackURL(ctx context.Context, in *TrackURLRequest, out *TrackURLResponse) error {
	return h.SpiderHandler.TrackURL(ctx, in, out)
}

func (h *Spider) FetchDataset(ctx context.Context, in *FetchDatasetRequest, out *FetchDatasetResponse) error {
	return h.SpiderHandler.FetchDataset(ctx, in, out)
}

func (h *Spider) FetchDatasetBy(ctx context.Context, in *FetchDatasetByRequest, out *FetchDatasetByResponse) error {
	return h.SpiderHandler.FetchDatasetBy(ctx, in, out)
}

func (h *Spider) PrepareDatasets(ctx context.Context, in *PrepareDatasetsRequest, out *PrepareDatasetsResponse) error {
	return h.SpiderHandler.PrepareDatasets(ctx, in, out)
}

func (h *Spider) FetchPages(ctx context.Context, in *FetchPagesRequest, out *FetchPagesResponse) error {
	return h.SpiderHandler.FetchPages(ctx, in, out)
}

func init() { proto.RegisterFile("proto/spider/spider.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 641 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x55, 0x4b, 0x6f, 0xd3, 0x40,
	0x10, 0xf6, 0x23, 0x49, 0x9b, 0x29, 0xea, 0x63, 0xeb, 0xa4, 0xc6, 0x82, 0x36, 0xec, 0x01, 0xe5,
	0xb4, 0x15, 0x2d, 0xa8, 0x05, 0xc1, 0x81, 0x57, 0xa5, 0x4a, 0x1c, 0x2a, 0x97, 0x48, 0xdc, 0xaa,
	0xc5, 0x5e, 0x1a, 0xab, 0x8e, 0x6d, 0x76, 0xd7, 0x48, 0xb9, 0x73, 0xe0, 0x57, 0xf2, 0x57, 0x40,
	0x5e, 0xaf, 0x53, 0xc7, 0x71, 0x24, 0x0e, 0x9c, 0xbc, 0xf3, 0x69, 0x76, 0x5e, 0xdf, 0x37, 0x6b,
	0x78, 0x98, 0xf1, 0x54, 0xa6, 0xc7, 0x22, 0x8b, 0x42, 0xc6, 0xf5, 0x87, 0x28, 0x0c, 0x73, 0xb0,
	0x27, 0x3c, 0x46, 0xdb, 0x60, 0x45, 0xa1, 0x6b, 0x8e, 0xcc, 0x71, 0xc7, 0xb7, 0xa2, 0x10, 0xed,
	0x82, 0x9d, 0xf3, 0xd8, 0xb5, 0x46, 0xe6, 0xb8, 0xef, 0x17, 0x47, 0xe4, 0x40, 0x57, 0x46, 0x32,
	0x66, 0xae, 0xad, 0xb0, 0xd2, 0x28, 0xd0, 0x5b, 0x9e, 0xe6, 0x99, 0xdb, 0x29, 0x51, 0x65, 0xa0,
	0xc7, 0x00, 0x22, 0x92, 0x6c, 0x46, 0xb3, 0x9b, 0x28, 0x74, 0xbb, 0x2a, 0x6a, 0x5f, 0x23, 0x97,
	0x21, 0x7e, 0x0d, 0xfb, 0x17, 0x4c, 0x06, 0xd3, 0x0f, 0x54, 0x52, 0xc1, 0xa4, 0xcf, 0xbe, 0xe7,
	0x4c, 0xc8, 0x95, 0x1a, 0x06, 0xd0, 0xcb, 0x79, 0x5c, 0x44, 0xb0, 0x14, 0xd6, 0xcd, 0x79, 0x7c,
	0x19, 0xe2, 0x9f, 0x26, 0x38, 0xcb, 0xd7, 0x45, 0x96, 0x26, 0x82, 0xa1, 0x53, 0xe8, 0x84, 0x54,
	0x52, 0xd7, 0x1c, 0xd9, 0xe3, 0xad, 0x93, 0x23, 0xd2, 0xe6, 0x44, 0x0a, 0xfb, 0x63, 0x22, 0xf9,
	0xdc, 0x57, 0xce, 0xde, 0x19, 0xf4, 0x17, 0x50, 0xd1, 0xf5, 0x1d, 0x9b, 0xab, 0x12, 0xfa, 0x7e,
	0x71, 0x2c, 0xfa, 0xfb, 0x41, 0xe3, 0x9c, 0xe9, 0x49, 0x94, 0xc6, 0x2b, 0xeb, 0xdc, 0xc4, 0x5f,
	0x60, 0x50, 0x4f, 0xf0, 0x6e, 0x5e, 0xb5, 0x71, 0x04, 0x5b, 0x82, 0xc5, 0x2c, 0x90, 0x29, 0xbf,
	0x59, 0xf4, 0x03, 0x15, 0x74, 0x19, 0xa2, 0x43, 0x80, 0x20, 0x4d, 0xc2, 0x48, 0x46, 0x69, 0x22,
	0x5c, 0x6b, 0x64, 0x8f, 0xfb, 0x7e, 0x0d, 0xc1, 0xbf, 0x4c, 0x18, 0x36, 0x43, 0xeb, 0x16, 0x5f,
	0x2c, 0xb5, 0xf8, 0x84, 0xb4, 0xbb, 0xfd, 0xbf, 0x26, 0x7f, 0x9b, 0xb0, 0xff, 0x99, 0xd3, 0xe0,
	0xee, 0xba, 0x24, 0xaf, 0xea, 0x11, 0x41, 0xe7, 0x1b, 0x4f, 0x67, 0x3a, 0x88, 0x3a, 0x17, 0x51,
	0x42, 0x96, 0xc9, 0x69, 0xc5, 0x96, 0x32, 0xd0, 0x39, 0x6c, 0x66, 0x54, 0x4a, 0xc6, 0x13, 0xe1,
	0xda, 0xaa, 0xea, 0x47, 0xa4, 0x25, 0x22, 0xb9, 0x2a, 0x9d, 0xfc, 0x85, 0x77, 0x91, 0x23, 0xa1,
	0x33, 0xa6, 0x95, 0xa5, 0xce, 0x68, 0x08, 0x3d, 0x21, 0x79, 0x14, 0x48, 0x25, 0xaa, 0x4d, 0x5f,
	0x5b, 0xde, 0x19, 0x6c, 0xe8, 0x00, 0x8b, 0x6b, 0x66, 0xed, 0x9a, 0x0b, 0x1b, 0x33, 0x2a, 0x83,
	0x29, 0x13, 0xba, 0xc5, 0xca, 0xc4, 0x4f, 0xc1, 0x59, 0xae, 0x46, 0x0f, 0xba, 0xa1, 0x45, 0xfc,
	0x12, 0x76, 0x94, 0xdf, 0xc4, 0xff, 0x54, 0xcd, 0x40, 0xaf, 0x88, 0xb9, 0xb4, 0x22, 0xe5, 0x32,
	0x58, 0xb5, 0x65, 0xc0, 0x18, 0x76, 0xef, 0xaf, 0xae, 0x09, 0x9f, 0xc2, 0xf0, 0x8a, 0xb3, 0x8c,
	0x72, 0xa6, 0xc9, 0x14, 0xff, 0xac, 0xa6, 0xd6, 0xa4, 0x0d, 0x8d, 0xd9, 0x2b, 0x1a, 0x3b, 0x86,
	0x83, 0x95, 0x84, 0xba, 0x36, 0x07, 0xba, 0x41, 0x9a, 0x27, 0x52, 0xe7, 0x2a, 0x0d, 0xfc, 0x16,
	0xf6, 0x94, 0xd8, 0xae, 0xe8, 0x2d, 0x5b, 0x14, 0x57, 0xd0, 0xc1, 0x28, 0x0f, 0xa6, 0x7a, 0x0a,
	0xda, 0x5a, 0x33, 0x88, 0xe7, 0x80, 0xea, 0x21, 0x74, 0xba, 0x43, 0xd8, 0xe0, 0x4c, 0xe4, 0xb1,
	0x14, 0x5a, 0xd5, 0x1d, 0x32, 0xe1, 0xb1, 0x5f, 0x81, 0x27, 0x7f, 0x2c, 0xe8, 0x5d, 0xab, 0x17,
	0x0b, 0xbd, 0x81, 0x07, 0x75, 0xb2, 0x90, 0xd3, 0xa6, 0x24, 0x6f, 0x40, 0xda, 0x18, 0xc5, 0x06,
	0x7a, 0x06, 0x9b, 0x15, 0x11, 0x68, 0x97, 0x34, 0xe8, 0xf4, 0xf6, 0x48, 0x93, 0x25, 0x6c, 0x14,
	0x19, 0xeb, 0x2b, 0x86, 0x1c, 0xd2, 0xf2, 0x70, 0x79, 0x83, 0xd6, 0xa7, 0x06, 0x1b, 0xe8, 0x3d,
	0x6c, 0x2f, 0x6f, 0x28, 0x1a, 0x92, 0xd6, 0x47, 0xc3, 0x3b, 0x58, 0xb3, 0xca, 0xd8, 0x40, 0x17,
	0xb0, 0xd3, 0xa0, 0x0a, 0x1d, 0x90, 0x76, 0xb5, 0x78, 0x2e, 0x59, 0xc3, 0x2a, 0x36, 0xd0, 0x19,
	0xc0, 0xfd, 0xf8, 0x11, 0x22, 0x2b, 0x74, 0x7a, 0xfb, 0x64, 0x95, 0x1f, 0x6c, 0x7c, 0xed, 0xa9,
	0x3f, 0xc5, 0xe9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x0d, 0x0c, 0x71, 0x46, 0x06, 0x00,
	0x00,
}
