// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/google/cloudprober/surfacers/prometheus/proto/config.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SurfacerConf struct {
	// How many metrics entries (EventMetrics) to buffer. Incoming metrics
	// processing is paused while serving data to prometheus. This buffer is to
	// make writes to prometheus surfacer non-blocking.
	// NOTE: This field is confusing for users and will be removed from the config
	// after v0.10.3.
	MetricsBufferSize *int64 `protobuf:"varint,1,opt,name=metrics_buffer_size,json=metricsBufferSize,def=10000" json:"metrics_buffer_size,omitempty"`
	// Whether to include timestamps in metrics. If enabled (default) each metric
	// string includes the metric timestamp as recorded in the EventMetric.
	// Prometheus associates the scraped values with this timestamp. If disabled,
	// i.e. timestamps are not exported, prometheus associates scraped values with
	// scrape timestamp.
	IncludeTimestamp *bool `protobuf:"varint,2,opt,name=include_timestamp,json=includeTimestamp,def=1" json:"include_timestamp,omitempty"`
	// URL that prometheus scrapes metrics from.
	MetricsUrl *string `protobuf:"bytes,3,opt,name=metrics_url,json=metricsUrl,def=/metrics" json:"metrics_url,omitempty"`
	// Prefix to add to all metric names. For example setting this field to
	// "cloudprober_" will result in metrics with names:
	// cloudprober_total, cloudprober_success, cloudprober_latency, ..
	MetricsPrefix        *string  `protobuf:"bytes,4,opt,name=metrics_prefix,json=metricsPrefix" json:"metrics_prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SurfacerConf) Reset()         { *m = SurfacerConf{} }
func (m *SurfacerConf) String() string { return proto.CompactTextString(m) }
func (*SurfacerConf) ProtoMessage()    {}
func (*SurfacerConf) Descriptor() ([]byte, []int) {
	return fileDescriptor_c23443a82a936d49, []int{0}
}

func (m *SurfacerConf) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SurfacerConf.Unmarshal(m, b)
}
func (m *SurfacerConf) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SurfacerConf.Marshal(b, m, deterministic)
}
func (m *SurfacerConf) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SurfacerConf.Merge(m, src)
}
func (m *SurfacerConf) XXX_Size() int {
	return xxx_messageInfo_SurfacerConf.Size(m)
}
func (m *SurfacerConf) XXX_DiscardUnknown() {
	xxx_messageInfo_SurfacerConf.DiscardUnknown(m)
}

var xxx_messageInfo_SurfacerConf proto.InternalMessageInfo

const Default_SurfacerConf_MetricsBufferSize int64 = 10000
const Default_SurfacerConf_IncludeTimestamp bool = true
const Default_SurfacerConf_MetricsUrl string = "/metrics"

func (m *SurfacerConf) GetMetricsBufferSize() int64 {
	if m != nil && m.MetricsBufferSize != nil {
		return *m.MetricsBufferSize
	}
	return Default_SurfacerConf_MetricsBufferSize
}

func (m *SurfacerConf) GetIncludeTimestamp() bool {
	if m != nil && m.IncludeTimestamp != nil {
		return *m.IncludeTimestamp
	}
	return Default_SurfacerConf_IncludeTimestamp
}

func (m *SurfacerConf) GetMetricsUrl() string {
	if m != nil && m.MetricsUrl != nil {
		return *m.MetricsUrl
	}
	return Default_SurfacerConf_MetricsUrl
}

func (m *SurfacerConf) GetMetricsPrefix() string {
	if m != nil && m.MetricsPrefix != nil {
		return *m.MetricsPrefix
	}
	return ""
}

func init() {
	proto.RegisterType((*SurfacerConf)(nil), "cloudprober.surfacer.prometheus.SurfacerConf")
}

func init() {
	proto.RegisterFile("github.com/google/cloudprober/surfacers/prometheus/proto/config.proto", fileDescriptor_c23443a82a936d49)
}

var fileDescriptor_c23443a82a936d49 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8d, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0x89, 0xbb, 0xc2, 0x1a, 0xff, 0xe0, 0xc6, 0x4b, 0x6f, 0x16, 0x41, 0xa8, 0x97, 0xa6,
	0x7b, 0xf0, 0xd2, 0xa3, 0xe2, 0x5d, 0xba, 0x7a, 0x2e, 0xdb, 0xec, 0xa4, 0x1b, 0x48, 0x3a, 0x65,
	0x92, 0x80, 0xf4, 0x23, 0xfa, 0xa9, 0xc4, 0x92, 0xa2, 0xb7, 0x37, 0xbf, 0xf7, 0x7b, 0x0c, 0x7f,
	0xeb, 0x4d, 0x38, 0xc5, 0xae, 0x54, 0xe8, 0x64, 0x8f, 0xd8, 0x5b, 0x90, 0xca, 0x62, 0x3c, 0x8e,
	0x84, 0x1d, 0x90, 0xf4, 0x91, 0xf4, 0x41, 0x01, 0x79, 0x39, 0x12, 0x3a, 0x08, 0x27, 0x88, 0x73,
	0x0c, 0x28, 0x15, 0x0e, 0xda, 0xf4, 0xe5, 0x7c, 0x88, 0xfb, 0x7f, 0xa3, 0x72, 0x19, 0x95, 0x7f,
	0x9b, 0x87, 0x6f, 0xc6, 0xaf, 0xf6, 0x89, 0xbf, 0xe2, 0xa0, 0xc5, 0x33, 0xbf, 0x73, 0x10, 0xc8,
	0x28, 0xdf, 0x76, 0x51, 0x6b, 0xa0, 0xd6, 0x9b, 0x09, 0x32, 0x96, 0xb3, 0x62, 0x55, 0x9f, 0xef,
	0xaa, 0xaa, 0xaa, 0x9a, 0x6d, 0x32, 0x5e, 0x66, 0x61, 0x6f, 0x26, 0x10, 0x3b, 0xbe, 0x35, 0x83,
	0xb2, 0xf1, 0x08, 0x6d, 0x30, 0x0e, 0x7c, 0x38, 0xb8, 0x31, 0x3b, 0xcb, 0x59, 0xb1, 0xa9, 0xd7,
	0x81, 0x22, 0x34, 0xb7, 0xa9, 0xfe, 0x58, 0x5a, 0xf1, 0xc4, 0x2f, 0x97, 0x4f, 0x91, 0x6c, 0xb6,
	0xca, 0x59, 0x71, 0x51, 0x6f, 0x64, 0x62, 0x0d, 0x4f, 0xe1, 0x93, 0xac, 0x78, 0xe4, 0x37, 0x8b,
	0x3a, 0x12, 0x68, 0xf3, 0x95, 0xad, 0x7f, 0xed, 0xe6, 0x3a, 0xd1, 0xf7, 0x19, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0x19, 0xe8, 0x1c, 0x35, 0x01, 0x00, 0x00,
}