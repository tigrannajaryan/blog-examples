// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metric_data.proto

package otlp

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

// ValueType is the enumeration of possible types that value can have.
type AttributeKeyValue_ValueType int32

const (
	AttributeKeyValue_STRING AttributeKeyValue_ValueType = 0
	AttributeKeyValue_BOOL   AttributeKeyValue_ValueType = 1
	AttributeKeyValue_INT64  AttributeKeyValue_ValueType = 2
	AttributeKeyValue_DOUBLE AttributeKeyValue_ValueType = 3
)

var AttributeKeyValue_ValueType_name = map[int32]string{
	0: "STRING",
	1: "BOOL",
	2: "INT64",
	3: "DOUBLE",
}

var AttributeKeyValue_ValueType_value = map[string]int32{
	"STRING": 0,
	"BOOL":   1,
	"INT64":  2,
	"DOUBLE": 3,
}

func (x AttributeKeyValue_ValueType) String() string {
	return proto.EnumName(AttributeKeyValue_ValueType_name, int32(x))
}

func (AttributeKeyValue_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{3, 0}
}

type MetricDescriptor_Type int32

const (
	MetricDescriptor_UNSPECIFIED          MetricDescriptor_Type = 0
	MetricDescriptor_GAUGE_INT64          MetricDescriptor_Type = 1
	MetricDescriptor_GAUGE_DOUBLE         MetricDescriptor_Type = 2
	MetricDescriptor_GAUGE_HISTOGRAM      MetricDescriptor_Type = 3
	MetricDescriptor_COUNTER_INT64        MetricDescriptor_Type = 4
	MetricDescriptor_COUNTER_DOUBLE       MetricDescriptor_Type = 5
	MetricDescriptor_CUMULATIVE_HISTOGRAM MetricDescriptor_Type = 6
	MetricDescriptor_SUMMARY              MetricDescriptor_Type = 7
)

var MetricDescriptor_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "GAUGE_INT64",
	2: "GAUGE_DOUBLE",
	3: "GAUGE_HISTOGRAM",
	4: "COUNTER_INT64",
	5: "COUNTER_DOUBLE",
	6: "CUMULATIVE_HISTOGRAM",
	7: "SUMMARY",
}

var MetricDescriptor_Type_value = map[string]int32{
	"UNSPECIFIED":          0,
	"GAUGE_INT64":          1,
	"GAUGE_DOUBLE":         2,
	"GAUGE_HISTOGRAM":      3,
	"COUNTER_INT64":        4,
	"COUNTER_DOUBLE":       5,
	"CUMULATIVE_HISTOGRAM": 6,
	"SUMMARY":              7,
}

func (x MetricDescriptor_Type) String() string {
	return proto.EnumName(MetricDescriptor_Type_name, int32(x))
}

func (MetricDescriptor_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{6, 0}
}

// A list of metrics from a Resource.
type ResourceMetrics struct {
	Resource             *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Metrics              []*Metric `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ResourceMetrics) Reset()         { *m = ResourceMetrics{} }
func (m *ResourceMetrics) String() string { return proto.CompactTextString(m) }
func (*ResourceMetrics) ProtoMessage()    {}
func (*ResourceMetrics) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{0}
}

func (m *ResourceMetrics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMetrics.Unmarshal(m, b)
}
func (m *ResourceMetrics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMetrics.Marshal(b, m, deterministic)
}
func (m *ResourceMetrics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMetrics.Merge(m, src)
}
func (m *ResourceMetrics) XXX_Size() int {
	return xxx_messageInfo_ResourceMetrics.Size(m)
}
func (m *ResourceMetrics) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMetrics.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMetrics proto.InternalMessageInfo

func (m *ResourceMetrics) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceMetrics) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// A list of prepared metrics from a Resource.
type ResourceMetricsPrepared struct {
	Resource             []byte            `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Metrics              []*MetricPrepared `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ResourceMetricsPrepared) Reset()         { *m = ResourceMetricsPrepared{} }
func (m *ResourceMetricsPrepared) String() string { return proto.CompactTextString(m) }
func (*ResourceMetricsPrepared) ProtoMessage()    {}
func (*ResourceMetricsPrepared) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{1}
}

func (m *ResourceMetricsPrepared) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourceMetricsPrepared.Unmarshal(m, b)
}
func (m *ResourceMetricsPrepared) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourceMetricsPrepared.Marshal(b, m, deterministic)
}
func (m *ResourceMetricsPrepared) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourceMetricsPrepared.Merge(m, src)
}
func (m *ResourceMetricsPrepared) XXX_Size() int {
	return xxx_messageInfo_ResourceMetricsPrepared.Size(m)
}
func (m *ResourceMetricsPrepared) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourceMetricsPrepared.DiscardUnknown(m)
}

var xxx_messageInfo_ResourceMetricsPrepared proto.InternalMessageInfo

func (m *ResourceMetricsPrepared) GetResource() []byte {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *ResourceMetricsPrepared) GetMetrics() []*MetricPrepared {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// Resource information. This describes the source of telemetry data.
type Resource struct {
	Labels               []*AttributeKeyValue `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	DroppedLabelsCount   int32                `protobuf:"varint,2,opt,name=dropped_labels_count,json=droppedLabelsCount,proto3" json:"dropped_labels_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{2}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetLabels() []*AttributeKeyValue {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *Resource) GetDroppedLabelsCount() int32 {
	if m != nil {
		return m.DroppedLabelsCount
	}
	return 0
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Resource
// labels, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// type of the value.
	Type                 AttributeKeyValue_ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=otlp.AttributeKeyValue_ValueType" json:"type,omitempty"`
	StringValue          string                      `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	Int64Value           int64                       `protobuf:"varint,4,opt,name=int64_value,json=int64Value,proto3" json:"int64_value,omitempty"`
	BoolValue            bool                        `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	DoubleValue          float64                     `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{3}
}

func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeKeyValue.Unmarshal(m, b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return xxx_messageInfo_AttributeKeyValue.Size(m)
}
func (m *AttributeKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeKeyValue proto.InternalMessageInfo

func (m *AttributeKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AttributeKeyValue) GetType() AttributeKeyValue_ValueType {
	if m != nil {
		return m.Type
	}
	return AttributeKeyValue_STRING
}

func (m *AttributeKeyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeKeyValue) GetInt64Value() int64 {
	if m != nil {
		return m.Int64Value
	}
	return 0
}

func (m *AttributeKeyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AttributeKeyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

// Defines a Metric which has one or more timeseries.
type Metric struct {
	MetricDescriptor     *MetricDescriptor  `protobuf:"bytes,1,opt,name=metric_descriptor,json=metricDescriptor,proto3" json:"metric_descriptor,omitempty"`
	Resource             *Resource          `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Int64Timeseries      []*Int64TimeSeries `protobuf:"bytes,3,rep,name=int64_timeseries,json=int64Timeseries,proto3" json:"int64_timeseries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{4}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetMetricDescriptor() *MetricDescriptor {
	if m != nil {
		return m.MetricDescriptor
	}
	return nil
}

func (m *Metric) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Metric) GetInt64Timeseries() []*Int64TimeSeries {
	if m != nil {
		return m.Int64Timeseries
	}
	return nil
}

// MetricPrepared is the prepared version of Metric.
//
// MetricPrepared is byte-level compatible with Metric. A byte-array encoding of
// MetricPrepared can be decoded as Metric and vice-versa. This allows senders that
// need to continuously create and encode instances of Metric with unchanging
// MetricDescriptor and Resource values to prepare and encode metric_descriptor and
// resource fields once and then create instances MetricPrepared messages where only
// timeseries data changes.
type MetricPrepared struct {
	// metric_descriptor is byte array representation of MetricDescriptor encoded
	// in ProtoBuf format using proto.Marshal().
	MetricDescriptor []byte `protobuf:"bytes,1,opt,name=metric_descriptor,json=metricDescriptor,proto3" json:"metric_descriptor,omitempty"`
	// resource is byte array representation of Resource encoded
	// in ProtoBuf format using proto.Marshal().
	Resource             []byte                     `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Int64Timeseries      []*Int64TimeSeriesPrepared `protobuf:"bytes,3,rep,name=int64_timeseries,json=int64Timeseries,proto3" json:"int64_timeseries,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MetricPrepared) Reset()         { *m = MetricPrepared{} }
func (m *MetricPrepared) String() string { return proto.CompactTextString(m) }
func (*MetricPrepared) ProtoMessage()    {}
func (*MetricPrepared) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{5}
}

func (m *MetricPrepared) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricPrepared.Unmarshal(m, b)
}
func (m *MetricPrepared) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricPrepared.Marshal(b, m, deterministic)
}
func (m *MetricPrepared) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricPrepared.Merge(m, src)
}
func (m *MetricPrepared) XXX_Size() int {
	return xxx_messageInfo_MetricPrepared.Size(m)
}
func (m *MetricPrepared) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricPrepared.DiscardUnknown(m)
}

var xxx_messageInfo_MetricPrepared proto.InternalMessageInfo

func (m *MetricPrepared) GetMetricDescriptor() []byte {
	if m != nil {
		return m.MetricDescriptor
	}
	return nil
}

func (m *MetricPrepared) GetResource() []byte {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *MetricPrepared) GetInt64Timeseries() []*Int64TimeSeriesPrepared {
	if m != nil {
		return m.Int64Timeseries
	}
	return nil
}

// Defines a metric type and its schema.
type MetricDescriptor struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string                `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Unit                 string                `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Type                 MetricDescriptor_Type `protobuf:"varint,4,opt,name=type,proto3,enum=otlp.MetricDescriptor_Type" json:"type,omitempty"`
	LabelKeys            []string              `protobuf:"bytes,5,rep,name=label_keys,json=labelKeys,proto3" json:"label_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MetricDescriptor) Reset()         { *m = MetricDescriptor{} }
func (m *MetricDescriptor) String() string { return proto.CompactTextString(m) }
func (*MetricDescriptor) ProtoMessage()    {}
func (*MetricDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{6}
}

func (m *MetricDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDescriptor.Unmarshal(m, b)
}
func (m *MetricDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDescriptor.Marshal(b, m, deterministic)
}
func (m *MetricDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDescriptor.Merge(m, src)
}
func (m *MetricDescriptor) XXX_Size() int {
	return xxx_messageInfo_MetricDescriptor.Size(m)
}
func (m *MetricDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDescriptor proto.InternalMessageInfo

func (m *MetricDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MetricDescriptor) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *MetricDescriptor) GetType() MetricDescriptor_Type {
	if m != nil {
		return m.Type
	}
	return MetricDescriptor_UNSPECIFIED
}

func (m *MetricDescriptor) GetLabelKeys() []string {
	if m != nil {
		return m.LabelKeys
	}
	return nil
}

// Int64TimeSeries is a list of data points that describes the time-varying values
// of a int64 metric.
type Int64TimeSeries struct {
	LabelValues          *LabelValues  `protobuf:"bytes,1,opt,name=label_values,json=labelValues,proto3" json:"label_values,omitempty"`
	Points               []*Int64Value `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Int64TimeSeries) Reset()         { *m = Int64TimeSeries{} }
func (m *Int64TimeSeries) String() string { return proto.CompactTextString(m) }
func (*Int64TimeSeries) ProtoMessage()    {}
func (*Int64TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{7}
}

func (m *Int64TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64TimeSeries.Unmarshal(m, b)
}
func (m *Int64TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64TimeSeries.Marshal(b, m, deterministic)
}
func (m *Int64TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64TimeSeries.Merge(m, src)
}
func (m *Int64TimeSeries) XXX_Size() int {
	return xxx_messageInfo_Int64TimeSeries.Size(m)
}
func (m *Int64TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_Int64TimeSeries proto.InternalMessageInfo

func (m *Int64TimeSeries) GetLabelValues() *LabelValues {
	if m != nil {
		return m.LabelValues
	}
	return nil
}

func (m *Int64TimeSeries) GetPoints() []*Int64Value {
	if m != nil {
		return m.Points
	}
	return nil
}

// Int64TimeSeriesPrepared is the prepared version of Int64TimeSeries.
type Int64TimeSeriesPrepared struct {
	LabelValues          []byte        `protobuf:"bytes,1,opt,name=label_values,json=labelValues,proto3" json:"label_values,omitempty"`
	Points               []*Int64Value `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Int64TimeSeriesPrepared) Reset()         { *m = Int64TimeSeriesPrepared{} }
func (m *Int64TimeSeriesPrepared) String() string { return proto.CompactTextString(m) }
func (*Int64TimeSeriesPrepared) ProtoMessage()    {}
func (*Int64TimeSeriesPrepared) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{8}
}

func (m *Int64TimeSeriesPrepared) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64TimeSeriesPrepared.Unmarshal(m, b)
}
func (m *Int64TimeSeriesPrepared) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64TimeSeriesPrepared.Marshal(b, m, deterministic)
}
func (m *Int64TimeSeriesPrepared) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64TimeSeriesPrepared.Merge(m, src)
}
func (m *Int64TimeSeriesPrepared) XXX_Size() int {
	return xxx_messageInfo_Int64TimeSeriesPrepared.Size(m)
}
func (m *Int64TimeSeriesPrepared) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64TimeSeriesPrepared.DiscardUnknown(m)
}

var xxx_messageInfo_Int64TimeSeriesPrepared proto.InternalMessageInfo

func (m *Int64TimeSeriesPrepared) GetLabelValues() []byte {
	if m != nil {
		return m.LabelValues
	}
	return nil
}

func (m *Int64TimeSeriesPrepared) GetPoints() []*Int64Value {
	if m != nil {
		return m.Points
	}
	return nil
}

// LabelValues is a list of label string values.
type LabelValues struct {
	Values                  []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	UnspecifiedValueIndexes []int32  `protobuf:"varint,2,rep,packed,name=unspecified_value_indexes,json=unspecifiedValueIndexes,proto3" json:"unspecified_value_indexes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *LabelValues) Reset()         { *m = LabelValues{} }
func (m *LabelValues) String() string { return proto.CompactTextString(m) }
func (*LabelValues) ProtoMessage()    {}
func (*LabelValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{9}
}

func (m *LabelValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelValues.Unmarshal(m, b)
}
func (m *LabelValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelValues.Marshal(b, m, deterministic)
}
func (m *LabelValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelValues.Merge(m, src)
}
func (m *LabelValues) XXX_Size() int {
	return xxx_messageInfo_LabelValues.Size(m)
}
func (m *LabelValues) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelValues.DiscardUnknown(m)
}

var xxx_messageInfo_LabelValues proto.InternalMessageInfo

func (m *LabelValues) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *LabelValues) GetUnspecifiedValueIndexes() []int32 {
	if m != nil {
		return m.UnspecifiedValueIndexes
	}
	return nil
}

// Int64Value is a timestamped measurement of int64 value.
type Int64Value struct {
	StartTimeUnixnano    int64    `protobuf:"fixed64,1,opt,name=start_time_unixnano,json=startTimeUnixnano,proto3" json:"start_time_unixnano,omitempty"`
	TimestampUnixnano    int64    `protobuf:"fixed64,2,opt,name=timestamp_unixnano,json=timestampUnixnano,proto3" json:"timestamp_unixnano,omitempty"`
	Value                int64    `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Value) Reset()         { *m = Int64Value{} }
func (m *Int64Value) String() string { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()    {}
func (*Int64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_9a96a8671bd66caf, []int{10}
}

func (m *Int64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Value.Unmarshal(m, b)
}
func (m *Int64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Value.Marshal(b, m, deterministic)
}
func (m *Int64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Value.Merge(m, src)
}
func (m *Int64Value) XXX_Size() int {
	return xxx_messageInfo_Int64Value.Size(m)
}
func (m *Int64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Value proto.InternalMessageInfo

func (m *Int64Value) GetStartTimeUnixnano() int64 {
	if m != nil {
		return m.StartTimeUnixnano
	}
	return 0
}

func (m *Int64Value) GetTimestampUnixnano() int64 {
	if m != nil {
		return m.TimestampUnixnano
	}
	return 0
}

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterEnum("otlp.AttributeKeyValue_ValueType", AttributeKeyValue_ValueType_name, AttributeKeyValue_ValueType_value)
	proto.RegisterEnum("otlp.MetricDescriptor_Type", MetricDescriptor_Type_name, MetricDescriptor_Type_value)
	proto.RegisterType((*ResourceMetrics)(nil), "otlp.ResourceMetrics")
	proto.RegisterType((*ResourceMetricsPrepared)(nil), "otlp.ResourceMetricsPrepared")
	proto.RegisterType((*Resource)(nil), "otlp.Resource")
	proto.RegisterType((*AttributeKeyValue)(nil), "otlp.AttributeKeyValue")
	proto.RegisterType((*Metric)(nil), "otlp.Metric")
	proto.RegisterType((*MetricPrepared)(nil), "otlp.MetricPrepared")
	proto.RegisterType((*MetricDescriptor)(nil), "otlp.MetricDescriptor")
	proto.RegisterType((*Int64TimeSeries)(nil), "otlp.Int64TimeSeries")
	proto.RegisterType((*Int64TimeSeriesPrepared)(nil), "otlp.Int64TimeSeriesPrepared")
	proto.RegisterType((*LabelValues)(nil), "otlp.LabelValues")
	proto.RegisterType((*Int64Value)(nil), "otlp.Int64Value")
}

func init() { proto.RegisterFile("metric_data.proto", fileDescriptor_9a96a8671bd66caf) }

var fileDescriptor_9a96a8671bd66caf = []byte{
	// 818 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0xdd, 0x6e, 0xdb, 0x36,
	0x18, 0xad, 0xfc, 0xa3, 0xc4, 0x9f, 0xb4, 0x58, 0xfe, 0x9a, 0x35, 0x5a, 0x87, 0x60, 0x8a, 0x2e,
	0x06, 0x61, 0xc3, 0xdc, 0x21, 0xcb, 0x8a, 0x61, 0x57, 0x73, 0x1d, 0x2f, 0x15, 0x1a, 0xc7, 0x05,
	0x6d, 0x15, 0xd8, 0x95, 0x20, 0xdb, 0xec, 0x40, 0xd4, 0x96, 0x34, 0x89, 0x1a, 0xea, 0xcb, 0xbd,
	0xc7, 0x76, 0xb5, 0xc7, 0xd8, 0x2b, 0xec, 0xa1, 0x06, 0x91, 0x94, 0xad, 0x3a, 0x31, 0x06, 0xf4,
	0xc6, 0x10, 0xcf, 0x39, 0x3c, 0xdf, 0x47, 0xf2, 0x90, 0x86, 0xde, 0x9a, 0xf2, 0x8c, 0x2d, 0xc2,
	0x65, 0xc4, 0xa3, 0x7e, 0x9a, 0x25, 0x3c, 0xc1, 0x56, 0xc2, 0x57, 0xa9, 0x4b, 0xa1, 0x4b, 0x68,
	0x9e, 0x14, 0xd9, 0x82, 0x8e, 0x85, 0x24, 0xc7, 0xaf, 0xe0, 0x38, 0x53, 0x90, 0xad, 0x39, 0x9a,
	0x67, 0x5c, 0x9e, 0xf4, 0x4b, 0x6d, 0xbf, 0x12, 0x92, 0x2d, 0x8f, 0x5f, 0xc2, 0x91, 0x74, 0xce,
	0xed, 0x86, 0xd3, 0xf4, 0x8c, 0x4b, 0x53, 0x4a, 0xa5, 0x17, 0xa9, 0x48, 0x97, 0xc2, 0xd9, 0x5e,
	0x99, 0xd7, 0x19, 0x4d, 0xa3, 0x8c, 0x2e, 0xf1, 0xe9, 0x5e, 0x39, 0xb3, 0x66, 0xdf, 0xdf, 0xb7,
	0x3f, 0xad, 0xdb, 0x57, 0x16, 0xbb, 0x32, 0x6b, 0x38, 0xae, 0xca, 0xe0, 0x33, 0xd0, 0x57, 0xd1,
	0x9c, 0xae, 0x72, 0x5b, 0x13, 0x53, 0xcf, 0xe4, 0xd4, 0x01, 0xe7, 0x19, 0x9b, 0x17, 0x9c, 0xbe,
	0xa2, 0x9b, 0x37, 0xd1, 0xaa, 0xa0, 0x44, 0xc9, 0xf0, 0x5b, 0x38, 0x5d, 0x66, 0x49, 0x9a, 0xd2,
	0x65, 0x28, 0x91, 0x70, 0x91, 0x14, 0x31, 0xb7, 0x1b, 0x8e, 0xe6, 0xb5, 0x09, 0x2a, 0xee, 0x56,
	0x50, 0xc3, 0x92, 0x71, 0xff, 0x6c, 0x40, 0xef, 0x9e, 0x1f, 0x5a, 0xd0, 0x7c, 0x47, 0x37, 0x62,
	0x2d, 0x1d, 0x52, 0x7e, 0xe2, 0xf7, 0xd0, 0xe2, 0x9b, 0x94, 0x0a, 0xa7, 0x93, 0xcb, 0x8b, 0x03,
	0x8d, 0xf4, 0xc5, 0xef, 0x6c, 0x93, 0x52, 0x22, 0xe4, 0x78, 0x01, 0x66, 0xce, 0x33, 0x16, 0xff,
	0x1a, 0xfe, 0x5e, 0x32, 0x76, 0x53, 0x38, 0x1a, 0x12, 0x93, 0xb5, 0xbe, 0x00, 0x83, 0xc5, 0xfc,
	0xf9, 0x95, 0x52, 0xb4, 0x1c, 0xcd, 0x6b, 0x12, 0x10, 0x90, 0x14, 0x9c, 0x03, 0xcc, 0x93, 0x64,
	0xa5, 0xf8, 0xb6, 0xa3, 0x79, 0xc7, 0xa4, 0x53, 0x22, 0x92, 0xbe, 0x00, 0x73, 0x99, 0x14, 0xf3,
	0x15, 0x55, 0x02, 0xdd, 0xd1, 0x3c, 0x8d, 0x18, 0x12, 0x13, 0x12, 0xf7, 0x07, 0xe8, 0x6c, 0x1b,
	0x43, 0x00, 0x7d, 0x3a, 0x23, 0xfe, 0xdd, 0x8d, 0xf5, 0x08, 0x8f, 0xa1, 0xf5, 0x62, 0x32, 0xb9,
	0xb5, 0x34, 0xec, 0x40, 0xdb, 0xbf, 0x9b, 0x3d, 0xbf, 0xb2, 0x1a, 0xa5, 0xe0, 0x7a, 0x12, 0xbc,
	0xb8, 0x1d, 0x59, 0x4d, 0xf7, 0x1f, 0x0d, 0x74, 0x79, 0x52, 0x38, 0xdc, 0x25, 0x90, 0xe6, 0x8b,
	0x8c, 0xa5, 0x3c, 0xc9, 0x54, 0xb8, 0x9e, 0xd4, 0x8f, 0xf4, 0x7a, 0xcb, 0x12, 0x6b, 0xbd, 0x87,
	0x7c, 0x10, 0xcc, 0xc6, 0xff, 0x04, 0xf3, 0x27, 0xb0, 0xe4, 0xc6, 0x70, 0xb6, 0xa6, 0x39, 0xcd,
	0x18, 0xcd, 0xed, 0xa6, 0xc8, 0xc1, 0xa7, 0x72, 0x8e, 0x5f, 0xb2, 0x33, 0xb6, 0xa6, 0x53, 0x41,
	0x92, 0x2e, 0xab, 0x00, 0xa9, 0x76, 0xff, 0xd6, 0xe0, 0xe4, 0xc3, 0x9c, 0xe1, 0xd7, 0x87, 0x56,
	0x61, 0x3e, 0xd0, 0xed, 0xd3, 0xbd, 0x6e, 0xeb, 0xb9, 0x7e, 0x79, 0xb0, 0xbb, 0xf3, 0x07, 0xbb,
	0xdb, 0x26, 0xfd, 0x5e, 0x97, 0xff, 0x36, 0xc0, 0xda, 0xdf, 0x3a, 0x44, 0x68, 0xc5, 0xd1, 0x9a,
	0xaa, 0x08, 0x8a, 0x6f, 0x74, 0xc0, 0xa8, 0x9a, 0x66, 0x49, 0x2c, 0x3a, 0xea, 0x90, 0x3a, 0x54,
	0xce, 0x2a, 0x62, 0xc6, 0x55, 0xcc, 0xc4, 0x37, 0x3e, 0x53, 0xc9, 0x6d, 0x89, 0xe4, 0x7e, 0xfe,
	0xf0, 0x51, 0xf5, 0x6b, 0x99, 0x3d, 0x07, 0x10, 0x97, 0x27, 0x7c, 0x47, 0x37, 0xb9, 0xdd, 0x76,
	0x9a, 0x5e, 0x87, 0x74, 0x04, 0xf2, 0x8a, 0x6e, 0x72, 0xf7, 0x2f, 0x0d, 0x5a, 0x22, 0x48, 0x5d,
	0x30, 0x82, 0xbb, 0xe9, 0xeb, 0xd1, 0xd0, 0xff, 0xd9, 0x1f, 0x5d, 0x5b, 0x8f, 0x4a, 0xe0, 0x66,
	0x10, 0xdc, 0x8c, 0x42, 0x99, 0x24, 0x0d, 0x2d, 0x30, 0x25, 0xa0, 0xf2, 0xd4, 0xc0, 0xc7, 0xd0,
	0x95, 0xc8, 0x4b, 0x7f, 0x3a, 0x9b, 0xdc, 0x90, 0xc1, 0xd8, 0x6a, 0x62, 0x0f, 0x3e, 0x19, 0x4e,
	0x82, 0xbb, 0xd9, 0x88, 0xa8, 0x99, 0x2d, 0x44, 0x38, 0xa9, 0x20, 0x35, 0xb7, 0x8d, 0x36, 0x9c,
	0x0e, 0x83, 0x71, 0x70, 0x3b, 0x98, 0xf9, 0x6f, 0xea, 0x06, 0x3a, 0x1a, 0x70, 0x34, 0x0d, 0xc6,
	0xe3, 0x01, 0xf9, 0xc5, 0x3a, 0x72, 0x7f, 0x83, 0xee, 0xde, 0xd6, 0xe3, 0x15, 0x98, 0x72, 0x45,
	0xe2, 0x86, 0xe4, 0x2a, 0xb5, 0x3d, 0xb9, 0x15, 0xe2, 0x35, 0x10, 0xd7, 0x23, 0x27, 0xc6, 0x6a,
	0x37, 0x40, 0x0f, 0xf4, 0x34, 0x61, 0x31, 0xaf, 0x1e, 0x2e, 0xab, 0x76, 0xae, 0xea, 0xd9, 0x91,
	0xbc, 0xfb, 0x16, 0xce, 0x0e, 0x9c, 0x76, 0x79, 0x3b, 0xef, 0x95, 0x36, 0x3f, 0xb6, 0x4e, 0x04,
	0x46, 0xad, 0x5b, 0x7c, 0x02, 0xfa, 0xd6, 0xb5, 0x3c, 0x24, 0x35, 0xc2, 0x1f, 0xe1, 0xb3, 0x22,
	0xce, 0x53, 0xba, 0x60, 0x6f, 0x19, 0x5d, 0xca, 0xca, 0x21, 0x8b, 0x97, 0xf4, 0x3d, 0x95, 0x35,
	0xda, 0xe4, 0xac, 0x26, 0x10, 0x6e, 0xbe, 0xa4, 0xdd, 0x3f, 0x34, 0x80, 0x5d, 0x65, 0xec, 0xc3,
	0xe3, 0x9c, 0x47, 0x19, 0x17, 0x29, 0x0f, 0x8b, 0x98, 0xbd, 0x8f, 0xa3, 0x38, 0x11, 0xab, 0xb0,
	0x48, 0x4f, 0x50, 0xe5, 0xa2, 0x03, 0x45, 0xe0, 0x37, 0x80, 0xe2, 0x3e, 0xf0, 0x68, 0x9d, 0xee,
	0xe4, 0x0d, 0x29, 0xdf, 0x32, 0x5b, 0xf9, 0x29, 0xb4, 0x77, 0xef, 0xa2, 0x45, 0xe4, 0x60, 0xae,
	0x8b, 0x7f, 0xb7, 0xef, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x35, 0xdb, 0xfa, 0xf2, 0x06,
	0x00, 0x00,
}
