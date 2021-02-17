// Code generated by protoc-gen-go. DO NOT EDIT.
// source: message/advertisement_dto.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 广告分组
type SAdGroup struct {
	Id   int64  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name"`
	// 是否开放给外部
	Opened               bool     `protobuf:"varint,3,opt,name=Opened,proto3" json:"Opened"`
	Enabled              bool     `protobuf:"varint,4,opt,name=Enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAdGroup) Reset()         { *m = SAdGroup{} }
func (m *SAdGroup) String() string { return proto.CompactTextString(m) }
func (*SAdGroup) ProtoMessage()    {}
func (*SAdGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{0}
}
func (m *SAdGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdGroup.Unmarshal(m, b)
}
func (m *SAdGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdGroup.Marshal(b, m, deterministic)
}
func (dst *SAdGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdGroup.Merge(dst, src)
}
func (m *SAdGroup) XXX_Size() int {
	return xxx_messageInfo_SAdGroup.Size(m)
}
func (m *SAdGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SAdGroup proto.InternalMessageInfo

func (m *SAdGroup) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SAdGroup) GetOpened() bool {
	if m != nil {
		return m.Opened
	}
	return false
}

func (m *SAdGroup) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

// 广告位编号
type AdPositionId struct {
	GroupId              int64    `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId"`
	PositionId           int64    `protobuf:"varint,2,opt,name=PositionId,proto3" json:"PositionId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdPositionId) Reset()         { *m = AdPositionId{} }
func (m *AdPositionId) String() string { return proto.CompactTextString(m) }
func (*AdPositionId) ProtoMessage()    {}
func (*AdPositionId) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{1}
}
func (m *AdPositionId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdPositionId.Unmarshal(m, b)
}
func (m *AdPositionId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdPositionId.Marshal(b, m, deterministic)
}
func (dst *AdPositionId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdPositionId.Merge(dst, src)
}
func (m *AdPositionId) XXX_Size() int {
	return xxx_messageInfo_AdPositionId.Size(m)
}
func (m *AdPositionId) XXX_DiscardUnknown() {
	xxx_messageInfo_AdPositionId.DiscardUnknown(m)
}

var xxx_messageInfo_AdPositionId proto.InternalMessageInfo

func (m *AdPositionId) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *AdPositionId) GetPositionId() int64 {
	if m != nil {
		return m.PositionId
	}
	return 0
}

type AdGroupListResponse struct {
	Value                []*SAdGroup `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AdGroupListResponse) Reset()         { *m = AdGroupListResponse{} }
func (m *AdGroupListResponse) String() string { return proto.CompactTextString(m) }
func (*AdGroupListResponse) ProtoMessage()    {}
func (*AdGroupListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{2}
}
func (m *AdGroupListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdGroupListResponse.Unmarshal(m, b)
}
func (m *AdGroupListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdGroupListResponse.Marshal(b, m, deterministic)
}
func (dst *AdGroupListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdGroupListResponse.Merge(dst, src)
}
func (m *AdGroupListResponse) XXX_Size() int {
	return xxx_messageInfo_AdGroupListResponse.Size(m)
}
func (m *AdGroupListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AdGroupListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AdGroupListResponse proto.InternalMessageInfo

func (m *AdGroupListResponse) GetValue() []*SAdGroup {
	if m != nil {
		return m.Value
	}
	return nil
}

// 广告位
type SAdPosition struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// 分组编号
	GroupId int64 `protobuf:"varint,2,opt,name=GroupId,proto3" json:"GroupId"`
	// 引用键
	Key string `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key"`
	// 名称
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name"`
	// todo:广告位类型限制
	// 广告类型限制,0为无限制
	TypeLimit int32 `protobuf:"varint,5,opt,name=TypeLimit,proto3" json:"TypeLimit"`
	// 是否开放给外部
	Opened bool `protobuf:"varint,6,opt,name=Opened,proto3" json:"Opened"`
	// 是否启用
	Enabled bool `protobuf:"varint,7,opt,name=Enabled,proto3" json:"Enabled"`
	// 默认广告编号
	DefaultId            int64    `protobuf:"varint,8,opt,name=DefaultId,proto3" json:"DefaultId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAdPosition) Reset()         { *m = SAdPosition{} }
func (m *SAdPosition) String() string { return proto.CompactTextString(m) }
func (*SAdPosition) ProtoMessage()    {}
func (*SAdPosition) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{3}
}
func (m *SAdPosition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdPosition.Unmarshal(m, b)
}
func (m *SAdPosition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdPosition.Marshal(b, m, deterministic)
}
func (dst *SAdPosition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdPosition.Merge(dst, src)
}
func (m *SAdPosition) XXX_Size() int {
	return xxx_messageInfo_SAdPosition.Size(m)
}
func (m *SAdPosition) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdPosition.DiscardUnknown(m)
}

var xxx_messageInfo_SAdPosition proto.InternalMessageInfo

func (m *SAdPosition) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdPosition) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *SAdPosition) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SAdPosition) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SAdPosition) GetTypeLimit() int32 {
	if m != nil {
		return m.TypeLimit
	}
	return 0
}

func (m *SAdPosition) GetOpened() bool {
	if m != nil {
		return m.Opened
	}
	return false
}

func (m *SAdPosition) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *SAdPosition) GetDefaultId() int64 {
	if m != nil {
		return m.DefaultId
	}
	return 0
}

// 广告用户设置
type SAdUserSet struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// 广告位编号
	PosId int64 `protobuf:"varint,2,opt,name=PosId,proto3" json:"PosId"`
	// 广告用户编号
	AdUserId int64 `protobuf:"varint,3,opt,name=AdUserId,proto3" json:"AdUserId"`
	// 广告编号
	AdId                 int64    `protobuf:"varint,4,opt,name=AdId,proto3" json:"AdId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAdUserSet) Reset()         { *m = SAdUserSet{} }
func (m *SAdUserSet) String() string { return proto.CompactTextString(m) }
func (*SAdUserSet) ProtoMessage()    {}
func (*SAdUserSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{4}
}
func (m *SAdUserSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdUserSet.Unmarshal(m, b)
}
func (m *SAdUserSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdUserSet.Marshal(b, m, deterministic)
}
func (dst *SAdUserSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdUserSet.Merge(dst, src)
}
func (m *SAdUserSet) XXX_Size() int {
	return xxx_messageInfo_SAdUserSet.Size(m)
}
func (m *SAdUserSet) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdUserSet.DiscardUnknown(m)
}

var xxx_messageInfo_SAdUserSet proto.InternalMessageInfo

func (m *SAdUserSet) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdUserSet) GetPosId() int64 {
	if m != nil {
		return m.PosId
	}
	return 0
}

func (m *SAdUserSet) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SAdUserSet) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

// 广告
type SAd struct {
	// 编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// 广告用户编号
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name"`
	// 广告类型
	Type int32 `protobuf:"varint,4,opt,name=Type,proto3" json:"Type"`
	// 展现次数
	ShowTimes int32 `protobuf:"varint,5,opt,name=ShowTimes,proto3" json:"ShowTimes"`
	// 点击次数
	ClickTimes int32 `protobuf:"varint,6,opt,name=ClickTimes,proto3" json:"ClickTimes"`
	// 展现天数
	ShowDays int32 `protobuf:"varint,7,opt,name=ShowDays,proto3" json:"ShowDays"`
	// 修改时间
	UpdateTime           int64    `protobuf:"varint,8,opt,name=UpdateTime,proto3" json:"UpdateTime"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SAd) Reset()         { *m = SAd{} }
func (m *SAd) String() string { return proto.CompactTextString(m) }
func (*SAd) ProtoMessage()    {}
func (*SAd) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{5}
}
func (m *SAd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAd.Unmarshal(m, b)
}
func (m *SAd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAd.Marshal(b, m, deterministic)
}
func (dst *SAd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAd.Merge(dst, src)
}
func (m *SAd) XXX_Size() int {
	return xxx_messageInfo_SAd.Size(m)
}
func (m *SAd) XXX_DiscardUnknown() {
	xxx_messageInfo_SAd.DiscardUnknown(m)
}

var xxx_messageInfo_SAd proto.InternalMessageInfo

func (m *SAd) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAd) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *SAd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SAd) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *SAd) GetShowTimes() int32 {
	if m != nil {
		return m.ShowTimes
	}
	return 0
}

func (m *SAd) GetClickTimes() int32 {
	if m != nil {
		return m.ClickTimes
	}
	return 0
}

func (m *SAd) GetShowDays() int32 {
	if m != nil {
		return m.ShowDays
	}
	return 0
}

func (m *SAd) GetUpdateTime() int64 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

// 广告数据传输对象
type SAdDto struct {
	Id                   int64             `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	AdType               int32             `protobuf:"varint,2,opt,name=AdType,proto3" json:"AdType"`
	Data                 map[string]string `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SAdDto) Reset()         { *m = SAdDto{} }
func (m *SAdDto) String() string { return proto.CompactTextString(m) }
func (*SAdDto) ProtoMessage()    {}
func (*SAdDto) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{6}
}
func (m *SAdDto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SAdDto.Unmarshal(m, b)
}
func (m *SAdDto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SAdDto.Marshal(b, m, deterministic)
}
func (dst *SAdDto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SAdDto.Merge(dst, src)
}
func (m *SAdDto) XXX_Size() int {
	return xxx_messageInfo_SAdDto.Size(m)
}
func (m *SAdDto) XXX_DiscardUnknown() {
	xxx_messageInfo_SAdDto.DiscardUnknown(m)
}

var xxx_messageInfo_SAdDto proto.InternalMessageInfo

func (m *SAdDto) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SAdDto) GetAdType() int32 {
	if m != nil {
		return m.AdType
	}
	return 0
}

func (m *SAdDto) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

//  超链接
type SHyperLink struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId"`
	Title                string   `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title"`
	LinkUrl              string   `protobuf:"bytes,4,opt,name=LinkUrl,proto3" json:"LinkUrl"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SHyperLink) Reset()         { *m = SHyperLink{} }
func (m *SHyperLink) String() string { return proto.CompactTextString(m) }
func (*SHyperLink) ProtoMessage()    {}
func (*SHyperLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{7}
}
func (m *SHyperLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SHyperLink.Unmarshal(m, b)
}
func (m *SHyperLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SHyperLink.Marshal(b, m, deterministic)
}
func (dst *SHyperLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SHyperLink.Merge(dst, src)
}
func (m *SHyperLink) XXX_Size() int {
	return xxx_messageInfo_SHyperLink.Size(m)
}
func (m *SHyperLink) XXX_DiscardUnknown() {
	xxx_messageInfo_SHyperLink.DiscardUnknown(m)
}

var xxx_messageInfo_SHyperLink proto.InternalMessageInfo

func (m *SHyperLink) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SHyperLink) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SHyperLink) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SHyperLink) GetLinkUrl() string {
	if m != nil {
		return m.LinkUrl
	}
	return ""
}

// 广告图片
type SImage struct {
	// 图片编号
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id"`
	// 广告编号
	AdId int64 `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId"`
	// 图片标题
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title"`
	// 链接
	LinkUrl string `protobuf:"bytes,4,opt,name=LinkUrl,proto3" json:"LinkUrl"`
	// 图片地址
	ImageUrl string `protobuf:"bytes,5,opt,name=ImageUrl,proto3" json:"ImageUrl"`
	// 排列序号
	SortNum int32 `protobuf:"varint,6,opt,name=SortNum,proto3" json:"SortNum"`
	// 是否启用
	Enabled              bool     `protobuf:"varint,7,opt,name=Enabled,proto3" json:"Enabled"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SImage) Reset()         { *m = SImage{} }
func (m *SImage) String() string { return proto.CompactTextString(m) }
func (*SImage) ProtoMessage()    {}
func (*SImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{8}
}
func (m *SImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SImage.Unmarshal(m, b)
}
func (m *SImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SImage.Marshal(b, m, deterministic)
}
func (dst *SImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SImage.Merge(dst, src)
}
func (m *SImage) XXX_Size() int {
	return xxx_messageInfo_SImage.Size(m)
}
func (m *SImage) XXX_DiscardUnknown() {
	xxx_messageInfo_SImage.DiscardUnknown(m)
}

var xxx_messageInfo_SImage proto.InternalMessageInfo

func (m *SImage) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SImage) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SImage) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SImage) GetLinkUrl() string {
	if m != nil {
		return m.LinkUrl
	}
	return ""
}

func (m *SImage) GetImageUrl() string {
	if m != nil {
		return m.ImageUrl
	}
	return ""
}

func (m *SImage) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SImage) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type SetDefaultAdRequest struct {
	GroupId              int64    `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId"`
	PosId                int64    `protobuf:"varint,2,opt,name=PosId,proto3" json:"PosId"`
	AdId                 int64    `protobuf:"varint,3,opt,name=AdId,proto3" json:"AdId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDefaultAdRequest) Reset()         { *m = SetDefaultAdRequest{} }
func (m *SetDefaultAdRequest) String() string { return proto.CompactTextString(m) }
func (*SetDefaultAdRequest) ProtoMessage()    {}
func (*SetDefaultAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{9}
}
func (m *SetDefaultAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDefaultAdRequest.Unmarshal(m, b)
}
func (m *SetDefaultAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDefaultAdRequest.Marshal(b, m, deterministic)
}
func (dst *SetDefaultAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDefaultAdRequest.Merge(dst, src)
}
func (m *SetDefaultAdRequest) XXX_Size() int {
	return xxx_messageInfo_SetDefaultAdRequest.Size(m)
}
func (m *SetDefaultAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDefaultAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDefaultAdRequest proto.InternalMessageInfo

func (m *SetDefaultAdRequest) GetGroupId() int64 {
	if m != nil {
		return m.GroupId
	}
	return 0
}

func (m *SetDefaultAdRequest) GetPosId() int64 {
	if m != nil {
		return m.PosId
	}
	return 0
}

func (m *SetDefaultAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

type SetUserAdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	PosId                int64    `protobuf:"varint,2,opt,name=PosId,proto3" json:"PosId"`
	AdId                 int64    `protobuf:"varint,3,opt,name=AdId,proto3" json:"AdId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetUserAdRequest) Reset()         { *m = SetUserAdRequest{} }
func (m *SetUserAdRequest) String() string { return proto.CompactTextString(m) }
func (*SetUserAdRequest) ProtoMessage()    {}
func (*SetUserAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{10}
}
func (m *SetUserAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetUserAdRequest.Unmarshal(m, b)
}
func (m *SetUserAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetUserAdRequest.Marshal(b, m, deterministic)
}
func (dst *SetUserAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetUserAdRequest.Merge(dst, src)
}
func (m *SetUserAdRequest) XXX_Size() int {
	return xxx_messageInfo_SetUserAdRequest.Size(m)
}
func (m *SetUserAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetUserAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetUserAdRequest proto.InternalMessageInfo

func (m *SetUserAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SetUserAdRequest) GetPosId() int64 {
	if m != nil {
		return m.PosId
	}
	return 0
}

func (m *SetUserAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

type AdIdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdIdRequest) Reset()         { *m = AdIdRequest{} }
func (m *AdIdRequest) String() string { return proto.CompactTextString(m) }
func (*AdIdRequest) ProtoMessage()    {}
func (*AdIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{11}
}
func (m *AdIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdIdRequest.Unmarshal(m, b)
}
func (m *AdIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdIdRequest.Marshal(b, m, deterministic)
}
func (dst *AdIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdIdRequest.Merge(dst, src)
}
func (m *AdIdRequest) XXX_Size() int {
	return xxx_messageInfo_AdIdRequest.Size(m)
}
func (m *AdIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdIdRequest proto.InternalMessageInfo

func (m *AdIdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *AdIdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

type AdKeyRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	AdPosKey             string   `protobuf:"bytes,2,opt,name=AdPosKey,proto3" json:"AdPosKey"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AdKeyRequest) Reset()         { *m = AdKeyRequest{} }
func (m *AdKeyRequest) String() string { return proto.CompactTextString(m) }
func (*AdKeyRequest) ProtoMessage()    {}
func (*AdKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{12}
}
func (m *AdKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AdKeyRequest.Unmarshal(m, b)
}
func (m *AdKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AdKeyRequest.Marshal(b, m, deterministic)
}
func (dst *AdKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AdKeyRequest.Merge(dst, src)
}
func (m *AdKeyRequest) XXX_Size() int {
	return xxx_messageInfo_AdKeyRequest.Size(m)
}
func (m *AdKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AdKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AdKeyRequest proto.InternalMessageInfo

func (m *AdKeyRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *AdKeyRequest) GetAdPosKey() string {
	if m != nil {
		return m.AdPosKey
	}
	return ""
}

type SaveAdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	Value                *SAd     `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveAdRequest) Reset()         { *m = SaveAdRequest{} }
func (m *SaveAdRequest) String() string { return proto.CompactTextString(m) }
func (*SaveAdRequest) ProtoMessage()    {}
func (*SaveAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{13}
}
func (m *SaveAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveAdRequest.Unmarshal(m, b)
}
func (m *SaveAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveAdRequest.Marshal(b, m, deterministic)
}
func (dst *SaveAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveAdRequest.Merge(dst, src)
}
func (m *SaveAdRequest) XXX_Size() int {
	return xxx_messageInfo_SaveAdRequest.Size(m)
}
func (m *SaveAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveAdRequest proto.InternalMessageInfo

func (m *SaveAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SaveAdRequest) GetValue() *SAd {
	if m != nil {
		return m.Value
	}
	return nil
}

type SaveLinkAdRequest struct {
	AdUserId             int64       `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	AdId                 int64       `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId"`
	Value                *SHyperLink `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SaveLinkAdRequest) Reset()         { *m = SaveLinkAdRequest{} }
func (m *SaveLinkAdRequest) String() string { return proto.CompactTextString(m) }
func (*SaveLinkAdRequest) ProtoMessage()    {}
func (*SaveLinkAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{14}
}
func (m *SaveLinkAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveLinkAdRequest.Unmarshal(m, b)
}
func (m *SaveLinkAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveLinkAdRequest.Marshal(b, m, deterministic)
}
func (dst *SaveLinkAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveLinkAdRequest.Merge(dst, src)
}
func (m *SaveLinkAdRequest) XXX_Size() int {
	return xxx_messageInfo_SaveLinkAdRequest.Size(m)
}
func (m *SaveLinkAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveLinkAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveLinkAdRequest proto.InternalMessageInfo

func (m *SaveLinkAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SaveLinkAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SaveLinkAdRequest) GetValue() *SHyperLink {
	if m != nil {
		return m.Value
	}
	return nil
}

type SaveImageAdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId"`
	Value                *SImage  `protobuf:"bytes,3,opt,name=Value,proto3" json:"Value"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SaveImageAdRequest) Reset()         { *m = SaveImageAdRequest{} }
func (m *SaveImageAdRequest) String() string { return proto.CompactTextString(m) }
func (*SaveImageAdRequest) ProtoMessage()    {}
func (*SaveImageAdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{15}
}
func (m *SaveImageAdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SaveImageAdRequest.Unmarshal(m, b)
}
func (m *SaveImageAdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SaveImageAdRequest.Marshal(b, m, deterministic)
}
func (dst *SaveImageAdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SaveImageAdRequest.Merge(dst, src)
}
func (m *SaveImageAdRequest) XXX_Size() int {
	return xxx_messageInfo_SaveImageAdRequest.Size(m)
}
func (m *SaveImageAdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SaveImageAdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SaveImageAdRequest proto.InternalMessageInfo

func (m *SaveImageAdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *SaveImageAdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *SaveImageAdRequest) GetValue() *SImage {
	if m != nil {
		return m.Value
	}
	return nil
}

type ImageAdIdRequest struct {
	AdUserId             int64    `protobuf:"varint,1,opt,name=AdUserId,proto3" json:"AdUserId"`
	AdId                 int64    `protobuf:"varint,2,opt,name=AdId,proto3" json:"AdId"`
	ImageId              int64    `protobuf:"varint,3,opt,name=ImageId,proto3" json:"ImageId"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImageAdIdRequest) Reset()         { *m = ImageAdIdRequest{} }
func (m *ImageAdIdRequest) String() string { return proto.CompactTextString(m) }
func (*ImageAdIdRequest) ProtoMessage()    {}
func (*ImageAdIdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_advertisement_dto_96c8f25a2ce1c925, []int{16}
}
func (m *ImageAdIdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImageAdIdRequest.Unmarshal(m, b)
}
func (m *ImageAdIdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImageAdIdRequest.Marshal(b, m, deterministic)
}
func (dst *ImageAdIdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImageAdIdRequest.Merge(dst, src)
}
func (m *ImageAdIdRequest) XXX_Size() int {
	return xxx_messageInfo_ImageAdIdRequest.Size(m)
}
func (m *ImageAdIdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ImageAdIdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ImageAdIdRequest proto.InternalMessageInfo

func (m *ImageAdIdRequest) GetAdUserId() int64 {
	if m != nil {
		return m.AdUserId
	}
	return 0
}

func (m *ImageAdIdRequest) GetAdId() int64 {
	if m != nil {
		return m.AdId
	}
	return 0
}

func (m *ImageAdIdRequest) GetImageId() int64 {
	if m != nil {
		return m.ImageId
	}
	return 0
}

func init() {
	proto.RegisterType((*SAdGroup)(nil), "SAdGroup")
	proto.RegisterType((*AdPositionId)(nil), "AdPositionId")
	proto.RegisterType((*AdGroupListResponse)(nil), "AdGroupListResponse")
	proto.RegisterType((*SAdPosition)(nil), "SAdPosition")
	proto.RegisterType((*SAdUserSet)(nil), "SAdUserSet")
	proto.RegisterType((*SAd)(nil), "SAd")
	proto.RegisterType((*SAdDto)(nil), "SAdDto")
	proto.RegisterMapType((map[string]string)(nil), "SAdDto.DataEntry")
	proto.RegisterType((*SHyperLink)(nil), "SHyperLink")
	proto.RegisterType((*SImage)(nil), "SImage")
	proto.RegisterType((*SetDefaultAdRequest)(nil), "SetDefaultAdRequest")
	proto.RegisterType((*SetUserAdRequest)(nil), "SetUserAdRequest")
	proto.RegisterType((*AdIdRequest)(nil), "AdIdRequest")
	proto.RegisterType((*AdKeyRequest)(nil), "AdKeyRequest")
	proto.RegisterType((*SaveAdRequest)(nil), "SaveAdRequest")
	proto.RegisterType((*SaveLinkAdRequest)(nil), "SaveLinkAdRequest")
	proto.RegisterType((*SaveImageAdRequest)(nil), "SaveImageAdRequest")
	proto.RegisterType((*ImageAdIdRequest)(nil), "ImageAdIdRequest")
}

func init() {
	proto.RegisterFile("message/advertisement_dto.proto", fileDescriptor_advertisement_dto_96c8f25a2ce1c925)
}

var fileDescriptor_advertisement_dto_96c8f25a2ce1c925 = []byte{
	// 735 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdb, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xe3, 0x5c, 0x27, 0x80, 0x5a, 0x17, 0x55, 0x56, 0x05, 0x6d, 0xb0, 0x84, 0x94, 0x27,
	0x23, 0x15, 0x09, 0x10, 0x88, 0x07, 0x43, 0x4a, 0x1b, 0xb5, 0x2a, 0xd5, 0xba, 0x45, 0x80, 0x90,
	0xa8, 0x93, 0xdd, 0xa6, 0x56, 0xe3, 0xac, 0xb1, 0x37, 0x45, 0xfe, 0x0b, 0x3e, 0x84, 0x3f, 0xe1,
	0x81, 0x5f, 0x42, 0x33, 0xbe, 0x96, 0xa6, 0x28, 0xaa, 0xc4, 0x53, 0xf6, 0xcc, 0x78, 0xcf, 0xec,
	0x9c, 0x39, 0xbb, 0x81, 0xad, 0x40, 0xc4, 0xb1, 0x37, 0x11, 0x4f, 0x3c, 0x7e, 0x29, 0x22, 0xe5,
	0xc7, 0x22, 0x10, 0x33, 0xf5, 0x95, 0x2b, 0x69, 0x87, 0x91, 0x54, 0xd2, 0x3a, 0x85, 0xb6, 0xeb,
	0xf0, 0xdd, 0x48, 0xce, 0x43, 0xe3, 0x1e, 0xd4, 0x86, 0xdc, 0xd4, 0x7a, 0x5a, 0x5f, 0x67, 0xb5,
	0x21, 0x37, 0x0c, 0xa8, 0x1f, 0x7a, 0x81, 0x30, 0x6b, 0x3d, 0xad, 0xdf, 0x61, 0xb4, 0x36, 0xd6,
	0xa1, 0xf9, 0x3e, 0x14, 0x33, 0xc1, 0x4d, 0xbd, 0xa7, 0xf5, 0xdb, 0x2c, 0x43, 0x86, 0x09, 0xad,
	0x9d, 0x99, 0x37, 0x9a, 0x0a, 0x6e, 0xd6, 0x29, 0x91, 0x43, 0x6b, 0x0f, 0xee, 0x38, 0xfc, 0x48,
	0xc6, 0xbe, 0xf2, 0xe5, 0x6c, 0x48, 0x5f, 0x52, 0xb9, 0xa2, 0x54, 0x0e, 0x8d, 0x4d, 0x80, 0xf2,
	0x3b, 0xaa, 0xaa, 0xb3, 0x4a, 0xc4, 0x7a, 0x06, 0x6b, 0xd9, 0x51, 0x0f, 0xfc, 0x58, 0x31, 0x11,
	0x87, 0x72, 0x16, 0x0b, 0x63, 0x0b, 0x1a, 0x1f, 0xbc, 0xe9, 0x5c, 0x98, 0x5a, 0x4f, 0xef, 0x77,
	0xb7, 0x3b, 0x76, 0xde, 0x10, 0x4b, 0xe3, 0xd6, 0x2f, 0x0d, 0xba, 0x6e, 0x79, 0x86, 0x6b, 0x7d,
	0x56, 0x4e, 0x54, 0xbb, 0x7a, 0xa2, 0x15, 0xd0, 0xf7, 0x45, 0x42, 0xad, 0x76, 0x18, 0x2e, 0x0b,
	0x4d, 0xea, 0x15, 0x4d, 0x1e, 0x40, 0xe7, 0x38, 0x09, 0xc5, 0x81, 0x1f, 0xf8, 0xca, 0x6c, 0xf4,
	0xb4, 0x7e, 0x83, 0x95, 0x81, 0x8a, 0x62, 0xcd, 0x9b, 0x14, 0x6b, 0x5d, 0x51, 0x0c, 0xf9, 0x06,
	0xe2, 0xcc, 0x9b, 0x4f, 0xd5, 0x90, 0x9b, 0x6d, 0x3a, 0x51, 0x19, 0xb0, 0x46, 0x00, 0xae, 0xc3,
	0x4f, 0x62, 0x11, 0xb9, 0x42, 0x5d, 0xeb, 0xe5, 0x3e, 0x34, 0x8e, 0x64, 0x5c, 0x74, 0x92, 0x02,
	0x63, 0x03, 0xda, 0xe9, 0x96, 0x61, 0x3a, 0x37, 0x9d, 0x15, 0x18, 0x3b, 0x72, 0xf8, 0x30, 0x1d,
	0x9b, 0xce, 0x68, 0x6d, 0xfd, 0xd6, 0x40, 0x77, 0x1d, 0x7e, 0x8d, 0x7d, 0x1d, 0x9a, 0x19, 0x4b,
	0x4a, 0xdf, 0x2c, 0x39, 0x48, 0x15, 0xbd, 0xa2, 0x8a, 0x01, 0x75, 0x14, 0x81, 0x78, 0x1b, 0x8c,
	0xd6, 0xd8, 0x99, 0x7b, 0x2e, 0xbf, 0x1f, 0xfb, 0x81, 0x88, 0x73, 0xa5, 0x8a, 0x00, 0xce, 0xff,
	0xed, 0xd4, 0x1f, 0x5f, 0xa4, 0xe9, 0x26, 0xa5, 0x2b, 0x11, 0xec, 0x02, 0x3f, 0x1e, 0x78, 0x49,
	0x4c, 0x92, 0x35, 0x58, 0x81, 0x71, 0xef, 0x49, 0xc8, 0x3d, 0x25, 0xf0, 0xd3, 0x4c, 0xb4, 0x4a,
	0xc4, 0xfa, 0xa1, 0x41, 0xd3, 0x75, 0xf8, 0x40, 0xc9, 0x45, 0x4d, 0x39, 0x9c, 0x8e, 0x5a, 0x23,
	0xd2, 0x0c, 0x19, 0x8f, 0xa1, 0x3e, 0xf0, 0x94, 0x67, 0xea, 0x64, 0xab, 0x55, 0x3b, 0xdd, 0x6e,
	0x63, 0x6c, 0x67, 0xa6, 0xa2, 0x84, 0x51, 0x7a, 0xe3, 0x39, 0x74, 0x8a, 0x10, 0x1a, 0xe6, 0x42,
	0x24, 0x44, 0xde, 0x61, 0xb8, 0xc4, 0x81, 0x5c, 0x92, 0x3b, 0xd3, 0x5b, 0x94, 0x82, 0x97, 0xb5,
	0x17, 0x9a, 0x75, 0x0a, 0xe0, 0xee, 0x25, 0xa1, 0x88, 0x0e, 0xfc, 0xd9, 0xc5, 0xa2, 0xcb, 0x47,
	0x63, 0xa9, 0x95, 0x63, 0x41, 0xae, 0x63, 0x5f, 0x4d, 0x73, 0x9d, 0x53, 0x80, 0x46, 0x42, 0x86,
	0x93, 0x68, 0x9a, 0xb9, 0x32, 0x87, 0xd6, 0x4f, 0x6c, 0x7a, 0x18, 0x78, 0x13, 0xf1, 0x3f, 0xe8,
	0x71, 0x1e, 0x44, 0x8e, 0xa9, 0x06, 0xa5, 0x0a, 0x8c, 0xbb, 0x5c, 0x19, 0xa9, 0xc3, 0x79, 0x90,
	0x0d, 0x32, 0x87, 0x37, 0xfb, 0xde, 0xfa, 0x04, 0x6b, 0xae, 0x50, 0x99, 0xd3, 0x1d, 0xce, 0xc4,
	0xb7, 0xb9, 0x88, 0xd5, 0x3f, 0x1e, 0x8c, 0xc5, 0x66, 0xcf, 0x5b, 0xd3, 0x2b, 0x86, 0xfe, 0x08,
	0x2b, 0xae, 0x50, 0xe8, 0xd6, 0x92, 0xb7, 0x7a, 0x29, 0xb4, 0xbf, 0x2e, 0xc5, 0xf2, 0xcc, 0xaf,
	0xa1, 0x8b, 0xbf, 0xcb, 0x90, 0x2e, 0xd0, 0xdc, 0x7a, 0x87, 0xaf, 0xe3, 0xbe, 0x48, 0x96, 0xd9,
	0x4f, 0xb9, 0x23, 0x19, 0xe3, 0x93, 0x94, 0xba, 0xa9, 0xc0, 0xd6, 0x2e, 0xdc, 0x75, 0xbd, 0x4b,
	0xb1, 0x5c, 0x77, 0x1b, 0xf9, 0x8b, 0x89, 0x2c, 0xdd, 0xed, 0x3a, 0x5a, 0x3b, 0x7f, 0x2c, 0xcf,
	0x60, 0x15, 0x89, 0x70, 0xc6, 0xce, 0x6d, 0xbb, 0x32, 0x1e, 0xe5, 0x05, 0x74, 0x2a, 0xd0, 0xb5,
	0x4b, 0xa3, 0xe7, 0x75, 0xc6, 0x60, 0x60, 0x1d, 0x32, 0xcc, 0xed, 0x0b, 0x3d, 0xbc, 0x5a, 0xa8,
	0x65, 0xa7, 0x76, 0xcf, 0x8b, 0x7c, 0x81, 0x95, 0xac, 0xc0, 0xad, 0x27, 0x84, 0xf6, 0x23, 0x8e,
	0x62, 0xee, 0x39, 0x7c, 0xb3, 0x09, 0x6b, 0x63, 0x19, 0xd8, 0x13, 0x5f, 0x9d, 0xcf, 0x47, 0xf6,
	0x44, 0x6e, 0x4b, 0x3b, 0x0a, 0xc7, 0x9f, 0x5b, 0xf6, 0x2b, 0xfa, 0x6f, 0x1d, 0x35, 0xe9, 0xe7,
	0xe9, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x97, 0x93, 0xcd, 0x85, 0x07, 0x00, 0x00,
}
