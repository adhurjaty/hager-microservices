// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user/user.proto

package go_micro_srv_user

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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Company              string   `protobuf:"bytes,3,opt,name=company" json:"company,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password" json:"password,omitempty"`
	AuthLevel            string   `protobuf:"bytes,6,opt,name=authLevel" json:"authLevel,omitempty"`
	FirstName            string   `protobuf:"bytes,7,opt,name=firstName" json:"firstName,omitempty"`
	LastName             string   `protobuf:"bytes,8,opt,name=lastName" json:"lastName,omitempty"`
	JobTitle             string   `protobuf:"bytes,9,opt,name=jobTitle" json:"jobTitle,omitempty"`
	PrimaryAddress       string   `protobuf:"bytes,10,opt,name=primaryAddress" json:"primaryAddress,omitempty"`
	SecondaryAddress     string   `protobuf:"bytes,11,opt,name=secondaryAddress" json:"secondaryAddress,omitempty"`
	City                 string   `protobuf:"bytes,12,opt,name=city" json:"city,omitempty"`
	State                string   `protobuf:"bytes,13,opt,name=state" json:"state,omitempty"`
	PostalCode           string   `protobuf:"bytes,14,opt,name=postalCode" json:"postalCode,omitempty"`
	Country              string   `protobuf:"bytes,15,opt,name=country" json:"country,omitempty"`
	Phone                string   `protobuf:"bytes,16,opt,name=phone" json:"phone,omitempty"`
	Fax                  string   `protobuf:"bytes,17,opt,name=fax" json:"fax,omitempty"`
	WebSite              string   `protobuf:"bytes,18,opt,name=webSite" json:"webSite,omitempty"`
	Comments             string   `protobuf:"bytes,19,opt,name=comments" json:"comments,omitempty"`
	EmailNews            bool     `protobuf:"varint,20,opt,name=emailNews" json:"emailNews,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_a318b1fea88de74f, []int{0}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetCompany() string {
	if m != nil {
		return m.Company
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *User) GetAuthLevel() string {
	if m != nil {
		return m.AuthLevel
	}
	return ""
}

func (m *User) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *User) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *User) GetJobTitle() string {
	if m != nil {
		return m.JobTitle
	}
	return ""
}

func (m *User) GetPrimaryAddress() string {
	if m != nil {
		return m.PrimaryAddress
	}
	return ""
}

func (m *User) GetSecondaryAddress() string {
	if m != nil {
		return m.SecondaryAddress
	}
	return ""
}

func (m *User) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *User) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *User) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *User) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *User) GetFax() string {
	if m != nil {
		return m.Fax
	}
	return ""
}

func (m *User) GetWebSite() string {
	if m != nil {
		return m.WebSite
	}
	return ""
}

func (m *User) GetComments() string {
	if m != nil {
		return m.Comments
	}
	return ""
}

func (m *User) GetEmailNews() bool {
	if m != nil {
		return m.EmailNews
	}
	return false
}

type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_a318b1fea88de74f, []int{1}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

type Response struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Users                []*User  `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_a318b1fea88de74f, []int{2}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Response) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *Response) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_a318b1fea88de74f, []int{3}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_a318b1fea88de74f, []int{4}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "go.micro.srv.user.User")
	proto.RegisterType((*Request)(nil), "go.micro.srv.user.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.user.Response")
	proto.RegisterType((*Token)(nil), "go.micro.srv.user.Token")
	proto.RegisterType((*Error)(nil), "go.micro.srv.user.Error")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for UserService service

type UserServiceClient interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type userServiceClient struct {
	c           client.Client
	serviceName string
}

func NewUserServiceClient(serviceName string, c client.Client) UserServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.user"
	}
	return &userServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *userServiceClient) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "UserService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&UserService{hdlr}, opts...))
}

type UserService struct {
	UserServiceHandler
}

func (h *UserService) Create(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *UserService) Get(ctx context.Context, in *User, out *Response) error {
	return h.UserServiceHandler.Get(ctx, in, out)
}

func (h *UserService) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.UserServiceHandler.GetAll(ctx, in, out)
}

func (h *UserService) Auth(ctx context.Context, in *User, out *Token) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *UserService) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}

func init() { proto.RegisterFile("proto/user/user.proto", fileDescriptor_user_a318b1fea88de74f) }

var fileDescriptor_user_a318b1fea88de74f = []byte{
	// 564 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdf, 0x4e, 0x13, 0x4f,
	0x14, 0xa6, 0xed, 0x6e, 0x69, 0x4f, 0x7f, 0xf0, 0x83, 0x11, 0xe3, 0x04, 0x8d, 0x69, 0xf6, 0xc2,
	0x10, 0x8d, 0xab, 0xc1, 0x4b, 0x43, 0x22, 0x21, 0x84, 0x1b, 0xc3, 0xc5, 0x82, 0xde, 0x0f, 0xbb,
	0x07, 0x19, 0xdd, 0xdd, 0x59, 0x67, 0xa6, 0xc5, 0xbe, 0x89, 0x4f, 0xe0, 0x2b, 0xf9, 0x3a, 0xe6,
	0x9c, 0xe9, 0x42, 0x23, 0x2d, 0x26, 0xdc, 0xb4, 0xf3, 0xfd, 0xe9, 0xb7, 0xd3, 0xef, 0x9c, 0x16,
	0x1e, 0x37, 0xd6, 0x78, 0xf3, 0x66, 0xe2, 0xd0, 0xf2, 0x4b, 0xca, 0x58, 0x6c, 0x7f, 0x31, 0x69,
	0xa5, 0x73, 0x6b, 0x52, 0x67, 0xa7, 0x29, 0x09, 0xc9, 0xaf, 0x08, 0xa2, 0x4f, 0x0e, 0xad, 0xd8,
	0x84, 0xae, 0x2e, 0x64, 0x67, 0xdc, 0xd9, 0x1b, 0x66, 0x5d, 0x5d, 0x88, 0x5d, 0x18, 0x90, 0xa1,
	0x56, 0x15, 0xca, 0x2e, 0xb3, 0x37, 0x58, 0x48, 0x58, 0xcf, 0x4d, 0xd5, 0xa8, 0x7a, 0x26, 0x7b,
	0x2c, 0xb5, 0x50, 0xec, 0x40, 0x8c, 0x95, 0xd2, 0xa5, 0x8c, 0x98, 0x0f, 0x80, 0xb2, 0x1a, 0xe5,
	0xdc, 0xb5, 0xb1, 0x85, 0x8c, 0x43, 0x56, 0x8b, 0xc5, 0x33, 0x18, 0xaa, 0x89, 0xbf, 0xfa, 0x88,
	0x53, 0x2c, 0x65, 0x9f, 0xc5, 0x5b, 0x82, 0xd4, 0x4b, 0x6d, 0x9d, 0x3f, 0xa5, 0x6b, 0xac, 0x07,
	0xf5, 0x86, 0xa0, 0xdc, 0x52, 0xcd, 0xc5, 0x41, 0xc8, 0x6d, 0x31, 0x69, 0x5f, 0xcd, 0xc5, 0xb9,
	0xf6, 0x25, 0xca, 0x61, 0xd0, 0x5a, 0x2c, 0x5e, 0xc0, 0x66, 0x63, 0x75, 0xa5, 0xec, 0xec, 0xb0,
	0x28, 0x2c, 0x3a, 0x27, 0x81, 0x1d, 0x7f, 0xb1, 0xe2, 0x25, 0x6c, 0x39, 0xcc, 0x4d, 0x5d, 0x2c,
	0x38, 0x47, 0xec, 0xbc, 0xc3, 0x0b, 0x01, 0x51, 0xae, 0xfd, 0x4c, 0xfe, 0xc7, 0x3a, 0x9f, 0xa9,
	0x0d, 0xe7, 0x95, 0x47, 0xb9, 0x11, 0xda, 0x60, 0x20, 0x9e, 0x03, 0x34, 0xc6, 0x79, 0x55, 0x1e,
	0x99, 0x02, 0xe5, 0x26, 0x4b, 0x0b, 0x4c, 0x68, 0x77, 0x52, 0x7b, 0x3b, 0x93, 0xff, 0xb7, 0xed,
	0x32, 0xa4, 0xbc, 0xe6, 0xca, 0xd4, 0x28, 0xb7, 0x42, 0x1e, 0x03, 0xb1, 0x05, 0xbd, 0x4b, 0xf5,
	0x43, 0x6e, 0x33, 0x47, 0x47, 0x4a, 0xb8, 0xc6, 0x8b, 0x33, 0xed, 0x51, 0x8a, 0x90, 0x30, 0x87,
	0xd4, 0x4a, 0x6e, 0xaa, 0x0a, 0x6b, 0xef, 0xe4, 0xa3, 0xd0, 0x4a, 0x8b, 0xa9, 0x6b, 0x1e, 0xd7,
	0x29, 0x5e, 0x3b, 0xb9, 0x33, 0xee, 0xec, 0x0d, 0xb2, 0x5b, 0x22, 0x19, 0xc2, 0x7a, 0x86, 0xdf,
	0x27, 0xe8, 0x7c, 0xf2, 0xb3, 0x03, 0x83, 0x0c, 0x5d, 0x63, 0x6a, 0x87, 0xe2, 0x15, 0x44, 0xb4,
	0x17, 0xbc, 0x39, 0xa3, 0xfd, 0x27, 0xe9, 0x9d, 0x15, 0x4b, 0x69, 0xbd, 0x32, 0x36, 0x89, 0xd7,
	0x10, 0xd3, 0xbb, 0x93, 0xdd, 0x71, 0xef, 0x3e, 0x77, 0x70, 0x89, 0xb7, 0xd0, 0x47, 0x6b, 0x8d,
	0x75, 0xb2, 0xc7, 0x7e, 0xb9, 0xc4, 0x7f, 0x4c, 0x86, 0x6c, 0xee, 0x4b, 0x10, 0xe2, 0x73, 0xf3,
	0x0d, 0x6b, 0xaa, 0xca, 0xd3, 0x61, 0xbe, 0xd1, 0x01, 0x10, 0x3b, 0x55, 0xa5, 0x2e, 0x78, 0xa3,
	0x07, 0x59, 0x00, 0x0f, 0x78, 0xcc, 0x01, 0xc4, 0x4c, 0xf0, 0xd4, 0x69, 0x8a, 0xf4, 0x94, 0x38,
	0xe3, 0xb3, 0x18, 0xc3, 0xa8, 0x40, 0x97, 0x5b, 0xdd, 0x78, 0x6d, 0xea, 0xf9, 0x8f, 0x67, 0x91,
	0xda, 0xff, 0xdd, 0x85, 0x11, 0x7d, 0xcf, 0x33, 0xb4, 0x53, 0x9d, 0xa3, 0xf8, 0x00, 0xfd, 0x23,
	0x8b, 0xb4, 0x1b, 0xab, 0x1a, 0xd9, 0x7d, 0xba, 0x44, 0x68, 0x67, 0x90, 0xac, 0x89, 0x03, 0xe8,
	0x9d, 0xa0, 0x7f, 0xf0, 0xc7, 0x8f, 0xa0, 0x7f, 0x82, 0xfe, 0xb0, 0x2c, 0xc5, 0xee, 0x52, 0x23,
	0xcf, 0xfd, 0x5f, 0x21, 0xef, 0x21, 0x3a, 0x9c, 0xf8, 0xab, 0xd5, 0x97, 0x58, 0xd6, 0x2b, 0x4f,
	0x2b, 0x59, 0x13, 0xc7, 0xb0, 0xf1, 0x99, 0x86, 0xa1, 0x3c, 0x86, 0x01, 0xae, 0x34, 0xdf, 0x17,
	0x73, 0xd1, 0xe7, 0x3f, 0xba, 0x77, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb6, 0x31, 0xe4, 0x29,
	0x01, 0x05, 0x00, 0x00,
}