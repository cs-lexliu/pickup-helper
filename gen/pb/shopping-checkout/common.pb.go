// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.2
// source: shopping-checkout/common.proto

package shoppingcheckoutpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Country  string `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	ZipCode  string `protobuf:"bytes,2,opt,name=zip_code,json=zipCode,proto3" json:"zip_code,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	County   string `protobuf:"bytes,4,opt,name=county,proto3" json:"county,omitempty"`
	Address1 string `protobuf:"bytes,5,opt,name=address1,proto3" json:"address1,omitempty"`
	Address2 string `protobuf:"bytes,6,opt,name=address2,proto3" json:"address2,omitempty"`
	UnitNo   string `protobuf:"bytes,7,opt,name=unit_no,json=unitNo,proto3" json:"unit_no,omitempty"`
	State    string `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *Address) GetAddress1() string {
	if x != nil {
		return x.Address1
	}
	return ""
}

func (x *Address) GetAddress2() string {
	if x != nil {
		return x.Address2
	}
	return ""
}

func (x *Address) GetUnitNo() string {
	if x != nil {
		return x.UnitNo
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type IconPath struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Android *IconPath_Android `protobuf:"bytes,1,opt,name=android,proto3" json:"android,omitempty"`
	Ios     *IconPath_IOS     `protobuf:"bytes,2,opt,name=ios,proto3" json:"ios,omitempty"`
	Web     string            `protobuf:"bytes,3,opt,name=web,proto3" json:"web,omitempty"`
}

func (x *IconPath) Reset() {
	*x = IconPath{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IconPath) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IconPath) ProtoMessage() {}

func (x *IconPath) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IconPath.ProtoReflect.Descriptor instead.
func (*IconPath) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{1}
}

func (x *IconPath) GetAndroid() *IconPath_Android {
	if x != nil {
		return x.Android
	}
	return nil
}

func (x *IconPath) GetIos() *IconPath_IOS {
	if x != nil {
		return x.Ios
	}
	return nil
}

func (x *IconPath) GetWeb() string {
	if x != nil {
		return x.Web
	}
	return ""
}

type PromoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string                       `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Formatting map[string]*PromoInfo_Format `protobuf:"bytes,2,rep,name=formatting,proto3" json:"formatting,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	IconPath   *IconPath                    `protobuf:"bytes,3,opt,name=icon_path,json=iconPath,proto3" json:"icon_path,omitempty"`
}

func (x *PromoInfo) Reset() {
	*x = PromoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromoInfo) ProtoMessage() {}

func (x *PromoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromoInfo.ProtoReflect.Descriptor instead.
func (*PromoInfo) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{2}
}

func (x *PromoInfo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *PromoInfo) GetFormatting() map[string]*PromoInfo_Format {
	if x != nil {
		return x.Formatting
	}
	return nil
}

func (x *PromoInfo) GetIconPath() *IconPath {
	if x != nil {
		return x.IconPath
	}
	return nil
}

type FormattedText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text       string                           `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Formatting map[string]*FormattedText_Format `protobuf:"bytes,2,rep,name=formatting,proto3" json:"formatting,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FormattedText) Reset() {
	*x = FormattedText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormattedText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormattedText) ProtoMessage() {}

func (x *FormattedText) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormattedText.ProtoReflect.Descriptor instead.
func (*FormattedText) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{3}
}

func (x *FormattedText) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *FormattedText) GetFormatting() map[string]*FormattedText_Format {
	if x != nil {
		return x.Formatting
	}
	return nil
}

type IconPath_Android struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hdpi    string `protobuf:"bytes,1,opt,name=hdpi,proto3" json:"hdpi,omitempty"`
	Mdpi    string `protobuf:"bytes,2,opt,name=mdpi,proto3" json:"mdpi,omitempty"`
	Xhdpi   string `protobuf:"bytes,3,opt,name=xhdpi,proto3" json:"xhdpi,omitempty"`
	Xxhdpi  string `protobuf:"bytes,4,opt,name=xxhdpi,proto3" json:"xxhdpi,omitempty"`
	Xxxhdpi string `protobuf:"bytes,5,opt,name=xxxhdpi,proto3" json:"xxxhdpi,omitempty"`
}

func (x *IconPath_Android) Reset() {
	*x = IconPath_Android{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IconPath_Android) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IconPath_Android) ProtoMessage() {}

func (x *IconPath_Android) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IconPath_Android.ProtoReflect.Descriptor instead.
func (*IconPath_Android) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{1, 0}
}

func (x *IconPath_Android) GetHdpi() string {
	if x != nil {
		return x.Hdpi
	}
	return ""
}

func (x *IconPath_Android) GetMdpi() string {
	if x != nil {
		return x.Mdpi
	}
	return ""
}

func (x *IconPath_Android) GetXhdpi() string {
	if x != nil {
		return x.Xhdpi
	}
	return ""
}

func (x *IconPath_Android) GetXxhdpi() string {
	if x != nil {
		return x.Xxhdpi
	}
	return ""
}

func (x *IconPath_Android) GetXxxhdpi() string {
	if x != nil {
		return x.Xxxhdpi
	}
	return ""
}

type IconPath_IOS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TwoX   string `protobuf:"bytes,1,opt,name=twoX,proto3" json:"twoX,omitempty"`
	ThreeX string `protobuf:"bytes,2,opt,name=threeX,proto3" json:"threeX,omitempty"`
}

func (x *IconPath_IOS) Reset() {
	*x = IconPath_IOS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IconPath_IOS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IconPath_IOS) ProtoMessage() {}

func (x *IconPath_IOS) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IconPath_IOS.ProtoReflect.Descriptor instead.
func (*IconPath_IOS) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{1, 1}
}

func (x *IconPath_IOS) GetTwoX() string {
	if x != nil {
		return x.TwoX
	}
	return ""
}

func (x *IconPath_IOS) GetThreeX() string {
	if x != nil {
		return x.ThreeX
	}
	return ""
}

type PromoInfo_Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PromoInfo_Format) Reset() {
	*x = PromoInfo_Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromoInfo_Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromoInfo_Format) ProtoMessage() {}

func (x *PromoInfo_Format) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromoInfo_Format.ProtoReflect.Descriptor instead.
func (*PromoInfo_Format) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{2, 0}
}

func (x *PromoInfo_Format) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

type FormattedText_Format struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FormattedText_Format) Reset() {
	*x = FormattedText_Format{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shopping_checkout_common_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FormattedText_Format) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FormattedText_Format) ProtoMessage() {}

func (x *FormattedText_Format) ProtoReflect() protoreflect.Message {
	mi := &file_shopping_checkout_common_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FormattedText_Format.ProtoReflect.Descriptor instead.
func (*FormattedText_Format) Descriptor() ([]byte, []int) {
	return file_shopping_checkout_common_proto_rawDescGZIP(), []int{3, 0}
}

func (x *FormattedText_Format) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

var File_shopping_checkout_common_proto protoreflect.FileDescriptor

var file_shopping_checkout_common_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x70, 0x62, 0x22, 0xd1, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x08, 0x7a, 0x69,
	0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69,
	0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x31, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x31, 0x12, 0x1a, 0x0a,
	0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69,
	0x74, 0x5f, 0x6e, 0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74,
	0x4e, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xbe, 0x02, 0x0a, 0x08, 0x49, 0x63, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3e, 0x0a, 0x07, 0x61, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e,
	0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x63, 0x6f, 0x6e,
	0x50, 0x61, 0x74, 0x68, 0x2e, 0x41, 0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x52, 0x07, 0x61, 0x6e,
	0x64, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x32, 0x0a, 0x03, 0x69, 0x6f, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70, 0x62, 0x2e, 0x49, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68,
	0x2e, 0x49, 0x4f, 0x53, 0x52, 0x03, 0x69, 0x6f, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x65, 0x62,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x65, 0x62, 0x1a, 0x79, 0x0a, 0x07, 0x41,
	0x6e, 0x64, 0x72, 0x6f, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x64, 0x70, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x64, 0x70, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x64,
	0x70, 0x69, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6d, 0x64, 0x70, 0x69, 0x12, 0x14,
	0x0a, 0x05, 0x78, 0x68, 0x64, 0x70, 0x69, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x78,
	0x68, 0x64, 0x70, 0x69, 0x12, 0x16, 0x0a, 0x06, 0x78, 0x78, 0x68, 0x64, 0x70, 0x69, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x78, 0x78, 0x68, 0x64, 0x70, 0x69, 0x12, 0x18, 0x0a, 0x07,
	0x78, 0x78, 0x78, 0x68, 0x64, 0x70, 0x69, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x78,
	0x78, 0x78, 0x68, 0x64, 0x70, 0x69, 0x1a, 0x31, 0x0a, 0x03, 0x49, 0x4f, 0x53, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x77, 0x6f, 0x58, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x77, 0x6f,
	0x58, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x68, 0x72, 0x65, 0x65, 0x58, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x68, 0x72, 0x65, 0x65, 0x58, 0x22, 0x96, 0x03, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x4d, 0x0a, 0x0a, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f,
	0x75, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x39, 0x0a, 0x09, 0x69, 0x63,
	0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x70, 0x62, 0x2e, 0x49, 0x63, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x52, 0x08, 0x69, 0x63, 0x6f,
	0x6e, 0x50, 0x61, 0x74, 0x68, 0x1a, 0x85, 0x01, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x42, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x46, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x63, 0x0a,
	0x0f, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x3a, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x6f, 0x49, 0x6e, 0x66, 0x6f,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xeb, 0x02, 0x0a, 0x0d, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x51, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70,
	0x62, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x2e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x1a, 0x89, 0x01, 0x0a, 0x06,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x46, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x73, 0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70, 0x62, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x1a, 0x37,
	0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x67, 0x0a, 0x0f, 0x46, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x73, 0x68,
	0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70, 0x62,
	0x2e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x2e, 0x46,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x61, 0x72, 0x6f, 0x75, 0x73, 0x65, 0x6c, 0x6c, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x68, 0x6f,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x3b, 0x73,
	0x68, 0x6f, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shopping_checkout_common_proto_rawDescOnce sync.Once
	file_shopping_checkout_common_proto_rawDescData = file_shopping_checkout_common_proto_rawDesc
)

func file_shopping_checkout_common_proto_rawDescGZIP() []byte {
	file_shopping_checkout_common_proto_rawDescOnce.Do(func() {
		file_shopping_checkout_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_shopping_checkout_common_proto_rawDescData)
	})
	return file_shopping_checkout_common_proto_rawDescData
}

var file_shopping_checkout_common_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_shopping_checkout_common_proto_goTypes = []interface{}{
	(*Address)(nil),              // 0: shoppingcheckoutpb.Address
	(*IconPath)(nil),             // 1: shoppingcheckoutpb.IconPath
	(*PromoInfo)(nil),            // 2: shoppingcheckoutpb.PromoInfo
	(*FormattedText)(nil),        // 3: shoppingcheckoutpb.FormattedText
	(*IconPath_Android)(nil),     // 4: shoppingcheckoutpb.IconPath.Android
	(*IconPath_IOS)(nil),         // 5: shoppingcheckoutpb.IconPath.IOS
	(*PromoInfo_Format)(nil),     // 6: shoppingcheckoutpb.PromoInfo.Format
	nil,                          // 7: shoppingcheckoutpb.PromoInfo.FormattingEntry
	nil,                          // 8: shoppingcheckoutpb.PromoInfo.Format.MetaEntry
	(*FormattedText_Format)(nil), // 9: shoppingcheckoutpb.FormattedText.Format
	nil,                          // 10: shoppingcheckoutpb.FormattedText.FormattingEntry
	nil,                          // 11: shoppingcheckoutpb.FormattedText.Format.MetaEntry
}
var file_shopping_checkout_common_proto_depIdxs = []int32{
	4,  // 0: shoppingcheckoutpb.IconPath.android:type_name -> shoppingcheckoutpb.IconPath.Android
	5,  // 1: shoppingcheckoutpb.IconPath.ios:type_name -> shoppingcheckoutpb.IconPath.IOS
	7,  // 2: shoppingcheckoutpb.PromoInfo.formatting:type_name -> shoppingcheckoutpb.PromoInfo.FormattingEntry
	1,  // 3: shoppingcheckoutpb.PromoInfo.icon_path:type_name -> shoppingcheckoutpb.IconPath
	10, // 4: shoppingcheckoutpb.FormattedText.formatting:type_name -> shoppingcheckoutpb.FormattedText.FormattingEntry
	8,  // 5: shoppingcheckoutpb.PromoInfo.Format.meta:type_name -> shoppingcheckoutpb.PromoInfo.Format.MetaEntry
	6,  // 6: shoppingcheckoutpb.PromoInfo.FormattingEntry.value:type_name -> shoppingcheckoutpb.PromoInfo.Format
	11, // 7: shoppingcheckoutpb.FormattedText.Format.meta:type_name -> shoppingcheckoutpb.FormattedText.Format.MetaEntry
	9,  // 8: shoppingcheckoutpb.FormattedText.FormattingEntry.value:type_name -> shoppingcheckoutpb.FormattedText.Format
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_shopping_checkout_common_proto_init() }
func file_shopping_checkout_common_proto_init() {
	if File_shopping_checkout_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shopping_checkout_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IconPath); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromoInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormattedText); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IconPath_Android); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IconPath_IOS); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromoInfo_Format); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_shopping_checkout_common_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FormattedText_Format); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shopping_checkout_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shopping_checkout_common_proto_goTypes,
		DependencyIndexes: file_shopping_checkout_common_proto_depIdxs,
		MessageInfos:      file_shopping_checkout_common_proto_msgTypes,
	}.Build()
	File_shopping_checkout_common_proto = out.File
	file_shopping_checkout_common_proto_rawDesc = nil
	file_shopping_checkout_common_proto_goTypes = nil
	file_shopping_checkout_common_proto_depIdxs = nil
}
