// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protob.proto

package protob

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type MatchType int32

const (
	MatchType_NONE            MatchType = 0
	MatchType_EXACT           MatchType = 1
	MatchType_CANONICAL_EXACT MatchType = 2
	MatchType_CANONICAL_FUZZY MatchType = 3
	MatchType_PARTIAL_EXACT   MatchType = 4
	MatchType_PARTIAL_FUZZY   MatchType = 5
)

var MatchType_name = map[int32]string{
	0: "NONE",
	1: "EXACT",
	2: "CANONICAL_EXACT",
	3: "CANONICAL_FUZZY",
	4: "PARTIAL_EXACT",
	5: "PARTIAL_FUZZY",
}
var MatchType_value = map[string]int32{
	"NONE":            0,
	"EXACT":           1,
	"CANONICAL_EXACT": 2,
	"CANONICAL_FUZZY": 3,
	"PARTIAL_EXACT":   4,
	"PARTIAL_FUZZY":   5,
}

func (x MatchType) String() string {
	return proto.EnumName(MatchType_name, int32(x))
}
func (MatchType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{0}
}

type Version struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Version) Reset()         { *m = Version{} }
func (m *Version) String() string { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()    {}
func (*Version) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{0}
}
func (m *Version) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Version.Unmarshal(m, b)
}
func (m *Version) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Version.Marshal(b, m, deterministic)
}
func (dst *Version) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Version.Merge(dst, src)
}
func (m *Version) XXX_Size() int {
	return xxx_messageInfo_Version.Size(m)
}
func (m *Version) XXX_DiscardUnknown() {
	xxx_messageInfo_Version.DiscardUnknown(m)
}

var xxx_messageInfo_Version proto.InternalMessageInfo

func (m *Version) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{1}
}
func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (dst *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(dst, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type Title struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ArchiveId            string   `protobuf:"bytes,2,opt,name=archive_id,json=archiveId,proto3" json:"archive_id,omitempty"`
	Path                 string   `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Lang                 string   `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Title) Reset()         { *m = Title{} }
func (m *Title) String() string { return proto.CompactTextString(m) }
func (*Title) ProtoMessage()    {}
func (*Title) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{2}
}
func (m *Title) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Title.Unmarshal(m, b)
}
func (m *Title) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Title.Marshal(b, m, deterministic)
}
func (dst *Title) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Title.Merge(dst, src)
}
func (m *Title) XXX_Size() int {
	return xxx_messageInfo_Title.Size(m)
}
func (m *Title) XXX_DiscardUnknown() {
	xxx_messageInfo_Title.DiscardUnknown(m)
}

var xxx_messageInfo_Title proto.InternalMessageInfo

func (m *Title) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Title) GetArchiveId() string {
	if m != nil {
		return m.ArchiveId
	}
	return ""
}

func (m *Title) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Title) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

type TitlesOpt struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TitlesOpt) Reset()         { *m = TitlesOpt{} }
func (m *TitlesOpt) String() string { return proto.CompactTextString(m) }
func (*TitlesOpt) ProtoMessage()    {}
func (*TitlesOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{3}
}
func (m *TitlesOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TitlesOpt.Unmarshal(m, b)
}
func (m *TitlesOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TitlesOpt.Marshal(b, m, deterministic)
}
func (dst *TitlesOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TitlesOpt.Merge(dst, src)
}
func (m *TitlesOpt) XXX_Size() int {
	return xxx_messageInfo_TitlesOpt.Size(m)
}
func (m *TitlesOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_TitlesOpt.DiscardUnknown(m)
}

var xxx_messageInfo_TitlesOpt proto.InternalMessageInfo

type Page struct {
	Id                   string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Offset               int32         `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Text                 []byte        `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	TitleId              string        `protobuf:"bytes,4,opt,name=title_id,json=titleId,proto3" json:"title_id,omitempty"`
	TitlePath            string        `protobuf:"bytes,5,opt,name=title_path,json=titlePath,proto3" json:"title_path,omitempty"`
	Names                []*NameString `protobuf:"bytes,6,rep,name=names,proto3" json:"names,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{4}
}
func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (dst *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(dst, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Page) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Page) GetText() []byte {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *Page) GetTitleId() string {
	if m != nil {
		return m.TitleId
	}
	return ""
}

func (m *Page) GetTitlePath() string {
	if m != nil {
		return m.TitlePath
	}
	return ""
}

func (m *Page) GetNames() []*NameString {
	if m != nil {
		return m.Names
	}
	return nil
}

type PagesOpt struct {
	WithText             bool     `protobuf:"varint,1,opt,name=with_text,json=withText,proto3" json:"with_text,omitempty"`
	TitleIds             []int32  `protobuf:"varint,2,rep,packed,name=title_ids,json=titleIds,proto3" json:"title_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagesOpt) Reset()         { *m = PagesOpt{} }
func (m *PagesOpt) String() string { return proto.CompactTextString(m) }
func (*PagesOpt) ProtoMessage()    {}
func (*PagesOpt) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{5}
}
func (m *PagesOpt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagesOpt.Unmarshal(m, b)
}
func (m *PagesOpt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagesOpt.Marshal(b, m, deterministic)
}
func (dst *PagesOpt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagesOpt.Merge(dst, src)
}
func (m *PagesOpt) XXX_Size() int {
	return xxx_messageInfo_PagesOpt.Size(m)
}
func (m *PagesOpt) XXX_DiscardUnknown() {
	xxx_messageInfo_PagesOpt.DiscardUnknown(m)
}

var xxx_messageInfo_PagesOpt proto.InternalMessageInfo

func (m *PagesOpt) GetWithText() bool {
	if m != nil {
		return m.WithText
	}
	return false
}

func (m *PagesOpt) GetTitleIds() []int32 {
	if m != nil {
		return m.TitleIds
	}
	return nil
}

type NameString struct {
	Value                string    `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Odds                 float32   `protobuf:"fixed32,2,opt,name=odds,proto3" json:"odds,omitempty"`
	Path                 string    `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	Curated              bool      `protobuf:"varint,4,opt,name=curated,proto3" json:"curated,omitempty"`
	EditDistance         int32     `protobuf:"varint,5,opt,name=edit_distance,json=editDistance,proto3" json:"edit_distance,omitempty"`
	EditDistanceStem     int32     `protobuf:"varint,6,opt,name=edit_distance_stem,json=editDistanceStem,proto3" json:"edit_distance_stem,omitempty"`
	SourceId             int32     `protobuf:"varint,7,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	Match                MatchType `protobuf:"varint,8,opt,name=match,proto3,enum=protob.MatchType" json:"match,omitempty"`
	OffsetStart          int32     `protobuf:"varint,9,opt,name=offset_start,json=offsetStart,proto3" json:"offset_start,omitempty"`
	OffsetEnd            int32     `protobuf:"varint,10,opt,name=offset_end,json=offsetEnd,proto3" json:"offset_end,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *NameString) Reset()         { *m = NameString{} }
func (m *NameString) String() string { return proto.CompactTextString(m) }
func (*NameString) ProtoMessage()    {}
func (*NameString) Descriptor() ([]byte, []int) {
	return fileDescriptor_protob_85767d7f15f16203, []int{6}
}
func (m *NameString) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameString.Unmarshal(m, b)
}
func (m *NameString) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameString.Marshal(b, m, deterministic)
}
func (dst *NameString) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameString.Merge(dst, src)
}
func (m *NameString) XXX_Size() int {
	return xxx_messageInfo_NameString.Size(m)
}
func (m *NameString) XXX_DiscardUnknown() {
	xxx_messageInfo_NameString.DiscardUnknown(m)
}

var xxx_messageInfo_NameString proto.InternalMessageInfo

func (m *NameString) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *NameString) GetOdds() float32 {
	if m != nil {
		return m.Odds
	}
	return 0
}

func (m *NameString) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *NameString) GetCurated() bool {
	if m != nil {
		return m.Curated
	}
	return false
}

func (m *NameString) GetEditDistance() int32 {
	if m != nil {
		return m.EditDistance
	}
	return 0
}

func (m *NameString) GetEditDistanceStem() int32 {
	if m != nil {
		return m.EditDistanceStem
	}
	return 0
}

func (m *NameString) GetSourceId() int32 {
	if m != nil {
		return m.SourceId
	}
	return 0
}

func (m *NameString) GetMatch() MatchType {
	if m != nil {
		return m.Match
	}
	return MatchType_NONE
}

func (m *NameString) GetOffsetStart() int32 {
	if m != nil {
		return m.OffsetStart
	}
	return 0
}

func (m *NameString) GetOffsetEnd() int32 {
	if m != nil {
		return m.OffsetEnd
	}
	return 0
}

func init() {
	proto.RegisterType((*Version)(nil), "protob.Version")
	proto.RegisterType((*Void)(nil), "protob.Void")
	proto.RegisterType((*Title)(nil), "protob.Title")
	proto.RegisterType((*TitlesOpt)(nil), "protob.TitlesOpt")
	proto.RegisterType((*Page)(nil), "protob.Page")
	proto.RegisterType((*PagesOpt)(nil), "protob.PagesOpt")
	proto.RegisterType((*NameString)(nil), "protob.NameString")
	proto.RegisterEnum("protob.MatchType", MatchType_name, MatchType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BHLIndexClient is the client API for BHLIndex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BHLIndexClient interface {
	Ver(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Version, error)
	Pages(ctx context.Context, in *PagesOpt, opts ...grpc.CallOption) (BHLIndex_PagesClient, error)
	Titles(ctx context.Context, in *TitlesOpt, opts ...grpc.CallOption) (BHLIndex_TitlesClient, error)
}

type bHLIndexClient struct {
	cc *grpc.ClientConn
}

func NewBHLIndexClient(cc *grpc.ClientConn) BHLIndexClient {
	return &bHLIndexClient{cc}
}

func (c *bHLIndexClient) Ver(ctx context.Context, in *Void, opts ...grpc.CallOption) (*Version, error) {
	out := new(Version)
	err := c.cc.Invoke(ctx, "/protob.BHLIndex/Ver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bHLIndexClient) Pages(ctx context.Context, in *PagesOpt, opts ...grpc.CallOption) (BHLIndex_PagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BHLIndex_serviceDesc.Streams[0], "/protob.BHLIndex/Pages", opts...)
	if err != nil {
		return nil, err
	}
	x := &bHLIndexPagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BHLIndex_PagesClient interface {
	Recv() (*Page, error)
	grpc.ClientStream
}

type bHLIndexPagesClient struct {
	grpc.ClientStream
}

func (x *bHLIndexPagesClient) Recv() (*Page, error) {
	m := new(Page)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bHLIndexClient) Titles(ctx context.Context, in *TitlesOpt, opts ...grpc.CallOption) (BHLIndex_TitlesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BHLIndex_serviceDesc.Streams[1], "/protob.BHLIndex/Titles", opts...)
	if err != nil {
		return nil, err
	}
	x := &bHLIndexTitlesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BHLIndex_TitlesClient interface {
	Recv() (*Title, error)
	grpc.ClientStream
}

type bHLIndexTitlesClient struct {
	grpc.ClientStream
}

func (x *bHLIndexTitlesClient) Recv() (*Title, error) {
	m := new(Title)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BHLIndexServer is the server API for BHLIndex service.
type BHLIndexServer interface {
	Ver(context.Context, *Void) (*Version, error)
	Pages(*PagesOpt, BHLIndex_PagesServer) error
	Titles(*TitlesOpt, BHLIndex_TitlesServer) error
}

func RegisterBHLIndexServer(s *grpc.Server, srv BHLIndexServer) {
	s.RegisterService(&_BHLIndex_serviceDesc, srv)
}

func _BHLIndex_Ver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BHLIndexServer).Ver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protob.BHLIndex/Ver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BHLIndexServer).Ver(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _BHLIndex_Pages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PagesOpt)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BHLIndexServer).Pages(m, &bHLIndexPagesServer{stream})
}

type BHLIndex_PagesServer interface {
	Send(*Page) error
	grpc.ServerStream
}

type bHLIndexPagesServer struct {
	grpc.ServerStream
}

func (x *bHLIndexPagesServer) Send(m *Page) error {
	return x.ServerStream.SendMsg(m)
}

func _BHLIndex_Titles_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TitlesOpt)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BHLIndexServer).Titles(m, &bHLIndexTitlesServer{stream})
}

type BHLIndex_TitlesServer interface {
	Send(*Title) error
	grpc.ServerStream
}

type bHLIndexTitlesServer struct {
	grpc.ServerStream
}

func (x *bHLIndexTitlesServer) Send(m *Title) error {
	return x.ServerStream.SendMsg(m)
}

var _BHLIndex_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protob.BHLIndex",
	HandlerType: (*BHLIndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ver",
			Handler:    _BHLIndex_Ver_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pages",
			Handler:       _BHLIndex_Pages_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Titles",
			Handler:       _BHLIndex_Titles_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "protob.proto",
}

func init() { proto.RegisterFile("protob.proto", fileDescriptor_protob_85767d7f15f16203) }

var fileDescriptor_protob_85767d7f15f16203 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xe1, 0x6e, 0xd3, 0x30,
	0x10, 0xc7, 0x97, 0x34, 0x4e, 0x93, 0x5b, 0xb6, 0x75, 0x07, 0x42, 0x61, 0x13, 0xa2, 0x04, 0x09,
	0x22, 0x40, 0x13, 0x1a, 0x4f, 0x50, 0xb6, 0x21, 0x22, 0x8d, 0x6e, 0xf2, 0xca, 0x04, 0xfb, 0x40,
	0x94, 0xd5, 0xde, 0x1a, 0xa9, 0x4d, 0xaa, 0xc4, 0x1d, 0xe3, 0x29, 0x78, 0x0b, 0x5e, 0x83, 0x57,
	0x43, 0x3e, 0x27, 0xa5, 0x95, 0xf6, 0xa9, 0xbe, 0x9f, 0xff, 0xbd, 0xfb, 0xdf, 0x39, 0x07, 0xc1,
	0xbc, 0x2a, 0x55, 0x79, 0x7d, 0x40, 0x3f, 0xe8, 0x9a, 0x28, 0x7a, 0x0e, 0xdd, 0x4b, 0x59, 0xd5,
	0x79, 0x59, 0xe0, 0x63, 0x60, 0x77, 0xd9, 0x74, 0x21, 0x43, 0xab, 0x6f, 0xc5, 0x3e, 0x37, 0x41,
	0xe4, 0x82, 0x73, 0x59, 0xe6, 0x22, 0xfa, 0x01, 0x6c, 0x94, 0xab, 0xa9, 0xc4, 0x6d, 0xb0, 0x73,
	0x41, 0x1a, 0xc6, 0xed, 0x5c, 0xe0, 0x33, 0x80, 0xac, 0x1a, 0x4f, 0xf2, 0x3b, 0x99, 0xe6, 0x22,
	0xb4, 0xe9, 0xbf, 0x7e, 0x43, 0x12, 0x81, 0x08, 0xce, 0x3c, 0x53, 0x93, 0xb0, 0x43, 0x17, 0x74,
	0xd6, 0x6c, 0x9a, 0x15, 0xb7, 0xa1, 0x63, 0x98, 0x3e, 0x47, 0x9b, 0xe0, 0x53, 0xfe, 0xfa, 0x6c,
	0xae, 0xa2, 0x3f, 0x16, 0x38, 0xe7, 0xd9, 0xed, 0x6a, 0x31, 0x9f, 0x8a, 0x3d, 0x01, 0xb7, 0xbc,
	0xb9, 0xa9, 0xa5, 0xa2, 0x42, 0x8c, 0x37, 0x91, 0xce, 0xa8, 0xe4, 0xbd, 0xa2, 0x2a, 0x01, 0xa7,
	0x33, 0x3e, 0x05, 0x4f, 0xe9, 0x8c, 0xda, 0x96, 0xa9, 0xd4, 0xa5, 0x38, 0x21, 0xcf, 0xe6, 0x8a,
	0xac, 0x31, 0xe3, 0x99, 0xc8, 0xb9, 0xf6, 0x17, 0x03, 0x2b, 0xb2, 0x99, 0xac, 0x43, 0xb7, 0xdf,
	0x89, 0x37, 0x0f, 0xf1, 0xa0, 0x19, 0xdd, 0x30, 0x9b, 0xc9, 0x0b, 0x55, 0xe5, 0xc5, 0x2d, 0x37,
	0x82, 0xe8, 0x18, 0x3c, 0xed, 0x53, 0x9b, 0xc6, 0x7d, 0xf0, 0x7f, 0xe6, 0x6a, 0x92, 0x92, 0x11,
	0x6d, 0xd9, 0xe3, 0x9e, 0x06, 0x23, 0x6d, 0x66, 0x1f, 0xfc, 0xd6, 0x4c, 0x1d, 0xda, 0xfd, 0x4e,
	0xcc, 0xb8, 0xd7, 0xb8, 0xa9, 0xa3, 0xbf, 0x36, 0xc0, 0xff, 0xdc, 0x0f, 0x3f, 0x84, 0x6e, 0xb1,
	0x14, 0xf4, 0x67, 0x2b, 0xb6, 0x39, 0x9d, 0x1f, 0x1c, 0x6e, 0x08, 0xdd, 0xf1, 0xa2, 0xca, 0x94,
	0x34, 0x5d, 0x7b, 0xbc, 0x0d, 0xf1, 0x25, 0x6c, 0x49, 0x91, 0xab, 0x54, 0xe4, 0xb5, 0xca, 0x8a,
	0xb1, 0xa4, 0xc6, 0x19, 0x0f, 0x34, 0x3c, 0x6e, 0x18, 0xbe, 0x03, 0x5c, 0x13, 0xa5, 0xb5, 0x92,
	0xb3, 0xd0, 0x25, 0x65, 0x6f, 0x55, 0x79, 0xa1, 0xe4, 0x4c, 0xb7, 0x55, 0x97, 0x8b, 0x6a, 0x4c,
	0x43, 0xee, 0x92, 0xc8, 0x33, 0x20, 0x11, 0xf8, 0x1a, 0xd8, 0x2c, 0x53, 0xe3, 0x49, 0xe8, 0xf5,
	0xad, 0x78, 0xfb, 0x70, 0xb7, 0x1d, 0xe3, 0x17, 0x0d, 0x47, 0xbf, 0xe6, 0x92, 0x9b, 0x7b, 0x7c,
	0x01, 0x81, 0x79, 0xc7, 0xb4, 0x56, 0x59, 0xa5, 0x42, 0x9f, 0x12, 0x6d, 0x1a, 0x76, 0xa1, 0x91,
	0x7e, 0xb1, 0x46, 0x22, 0x0b, 0x11, 0x02, 0x09, 0x7c, 0x43, 0x4e, 0x0a, 0xf1, 0x66, 0x0e, 0xfe,
	0x32, 0x2b, 0x7a, 0xe0, 0x0c, 0xcf, 0x86, 0x27, 0xbd, 0x0d, 0xf4, 0x81, 0x9d, 0x7c, 0x1b, 0x1c,
	0x8d, 0x7a, 0x16, 0x3e, 0x82, 0x9d, 0xa3, 0xc1, 0xf0, 0x6c, 0x98, 0x1c, 0x0d, 0x4e, 0x53, 0x03,
	0xed, 0x75, 0xf8, 0xe9, 0xeb, 0xd5, 0xd5, 0xf7, 0x5e, 0x07, 0x77, 0x61, 0xeb, 0x7c, 0xc0, 0x47,
	0xc9, 0x52, 0xe7, 0xac, 0x22, 0xa3, 0x62, 0x87, 0xbf, 0x2d, 0xf0, 0x3e, 0x7e, 0x3e, 0x4d, 0x0a,
	0x21, 0xef, 0xf1, 0x15, 0x74, 0x2e, 0x65, 0x85, 0x41, 0xdb, 0xa1, 0xde, 0x98, 0xbd, 0x9d, 0x65,
	0x64, 0x16, 0x2c, 0xda, 0xc0, 0xb7, 0xc0, 0xe8, 0x73, 0xc1, 0x5e, 0x7b, 0xd7, 0x7e, 0x3d, 0x7b,
	0xc1, 0x2a, 0x89, 0x36, 0xde, 0x5b, 0x78, 0x00, 0xae, 0xd9, 0x08, 0x5c, 0x4e, 0x6e, 0xb9, 0x21,
	0x7b, 0x5b, 0x6b, 0x48, 0xeb, 0xaf, 0xcd, 0x4a, 0x7f, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0x40,
	0xe2, 0xb1, 0xc5, 0xe9, 0x03, 0x00, 0x00,
}
