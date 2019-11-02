package prepared_protobuf

import (
	"bytes"
	"log"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"

	"github.com/tigrannajaryan/blog-examples/prepared-protobuf/otlp"
)

// Helper functions to generate test data.

func genLabels() *otlp.LabelValues {
	return &otlp.LabelValues{
		Values:                  []string{"val1", "val2", ""},
		UnspecifiedValueIndexes: []int32{2},
	}
}

func genResource() *otlp.Resource {
	return &otlp.Resource{
		Labels: []*otlp.AttributeKeyValue{
			{Key: "StartTimeUnixnano", Int64Value: 12345678},
			{Key: "Pid", Int64Value: 1234},
			{Key: "HostName", StringValue: "fakehost"},
			{Key: "ServiceName", StringValue: "generator"},
		},
	}
}

func genMetricDescriptor(i int) *otlp.MetricDescriptor {
	descr := &otlp.MetricDescriptor{
		Name:        "metric" + strconv.Itoa(i),
		Description: "some description: " + strconv.Itoa(i),
		Type:        otlp.MetricDescriptor_GAUGE_INT64,
		LabelKeys: []string{
			"label1",
			"label2",
		},
	}
	return descr
}

func genInt64DataPoints(offset int) []*otlp.Int64Value {
	var points []*otlp.Int64Value

	const dataPoints = 1
	for k := 0; k < dataPoints; k++ {
		pointTs := (time.Now().Add(time.Duration(k) * time.Millisecond)).UnixNano()

		point := otlp.Int64Value{
			TimestampUnixnano: pointTs,
			Value:             int64(offset * k),
		}

		if k == 0 {
			point.StartTimeUnixnano = pointTs
		}

		points = append(points, &point)
	}

	return points
}

func TestPreparedMetricEncodingIsIdentical(t *testing.T) {

	// This test verifies that full and prepared encodings are compatible on the wire.

	// Create a metric
	descr := &otlp.MetricDescriptor{
		Name:        "regularMetric",
		Description: "some description: 1",
		Type:        otlp.MetricDescriptor_GAUGE_INT64,
		LabelKeys: []string{
			"label1",
			"label2",
		},
	}

	resource := genResource()

	data := []*otlp.Int64Value{
		{
			TimestampUnixnano: time.Now().UnixNano(),
			Value:             123,
		},
	}

	labelValues := genLabels()

	regularMetric := &otlp.Metric{
		MetricDescriptor: descr,
		Resource:         resource,
		Int64Timeseries: []*otlp.Int64TimeSeries{
			{
				LabelValues: labelValues,
				Points:      data,
			},
		},
	}

	// Create a prepared metric. First prepare sub-components by
	// encoding them into byte arrays.
	descrBytes, err := proto.Marshal(descr)
	require.NoError(t, err)

	resourceBytes, err := proto.Marshal(resource)
	require.NoError(t, err)

	labelValuesBytes, err := proto.Marshal(labelValues)
	require.NoError(t, err)

	// Use byte arrays to create MetricPrepared struct.
	preparedMetric := &otlp.MetricPrepared{
		MetricDescriptor: descrBytes,
		Resource:         resourceBytes,
		Int64Timeseries: []*otlp.Int64TimeSeriesPrepared{
			{
				LabelValues: labelValuesBytes,
				Points:      data,
			},
		},
	}

	// Encode regular metric into byte array.
	regularMetricBytes, err := proto.Marshal(regularMetric)
	require.NoError(t, err)

	// Encode prepared metric into byte array.
	preparedMetricBytes, err := proto.Marshal(preparedMetric)
	require.NoError(t, err)

	// Ensure that encoded byte arrays are identical.
	assert.Equal(t, 0, bytes.Compare(regularMetricBytes, preparedMetricBytes),
		"Prepared and full encoding results are different")

	// Ensure it can be decoded fully.
	var fullMetricDecodedFromPrepared otlp.Metric
	err = proto.Unmarshal(preparedMetricBytes, &fullMetricDecodedFromPrepared)
	require.NoError(t, err)

	// Ensure it can be decoded partially.
	var preparedMetricDecodedFromPrepared otlp.MetricPrepared
	err = proto.Unmarshal(preparedMetricBytes, &preparedMetricDecodedFromPrepared)
	require.NoError(t, err)

	// Now decode prepared components and verify they were correctly encoded.

	var decodedResource otlp.Resource
	err = proto.Unmarshal(preparedMetricDecodedFromPrepared.Resource, &decodedResource)
	require.NoError(t, err)

	assert.True(t, proto.Equal(resource, &decodedResource),
		"Resource decoded from prepared byte array differs from original")

	var decodedMetricDescriptor otlp.MetricDescriptor
	err = proto.Unmarshal(preparedMetricDecodedFromPrepared.MetricDescriptor, &decodedMetricDescriptor)
	require.NoError(t, err)

	assert.True(t, proto.Equal(descr, &decodedMetricDescriptor),
		"Metric descriptor decoded from prepared byte array differs from original")

	var decodedLabelValues otlp.LabelValues
	err = proto.Unmarshal(preparedMetricDecodedFromPrepared.Int64Timeseries[0].LabelValues, &decodedLabelValues)
	require.NoError(t, err)

	assert.True(t, proto.Equal(labelValues, &decodedLabelValues),
		"Label values decoded from prepared byte array differs from original")
}

func genFullMetrics(metricCount int) proto.Message {
	// Generate test data.
	batch := &otlp.ResourceMetrics{Resource: genResource()}
	for i := 0; i < metricCount; i++ {
		metric := &otlp.Metric{
			MetricDescriptor: genMetricDescriptor(1),
			Int64Timeseries: []*otlp.Int64TimeSeries{
				{
					LabelValues: genLabels(),
					Points:      genInt64DataPoints(i),
				},
			},
		}
		batch.Metrics = append(batch.Metrics, metric)
	}
	return batch
}

// genPreparedMetrics creates data equivalent to what genFullMetrics
// but using "Prepared" version of ProtoBuf messages.
func genPreparedMetrics(metricCount int) proto.Message {
	// Prepare components and encode them.

	resource := genResource()
	resourceBytes, err := proto.Marshal(resource)
	if err != nil {
		log.Fatal("Cannot marshal resource")
	}

	descr := genMetricDescriptor(1)
	descrBytes, err := proto.Marshal(descr)
	if err != nil {
		log.Fatal("Cannot marshal metric descriptor")
	}

	labelValues := genLabels()
	labelValuesBytes, err := proto.Marshal(labelValues)
	if err != nil {
		log.Fatal("Cannot marshal label values")
	}

	// Create structs which reference prepared components.
	batch := &otlp.ResourceMetricsPrepared{
		Resource: resourceBytes,
	}
	for i := 0; i < metricCount; i++ {
		metric := &otlp.MetricPrepared{
			MetricDescriptor: descrBytes,
			Int64Timeseries: []*otlp.Int64TimeSeriesPrepared{
				{
					LabelValues: labelValuesBytes,
					Points:      genInt64DataPoints(i),
				},
			},
		}
		batch.Metrics = append(batch.Metrics, metric)
	}
	return batch
}

func BenchmarkMetricEncode(b *testing.B) {
	// This benchmark compares performance of regular encoding and prepared encoding of
	// the same metric data.

	tests := []struct {
		name      string
		generator func(metricCount int) proto.Message
	}{
		{
			name:      "Full",
			generator: genFullMetrics,
		},
		{
			name:      "Prepared",
			generator: genPreparedMetrics,
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			// Don't count timer spent on generating data.
			b.StopTimer()
			batch := test.generator(100)
			runtime.GC()
			b.ResetTimer()

			// Measure encoding time.
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				var err error
				bytes, err := proto.Marshal(batch)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot encode batch")
				}
			}
		})
	}
}

func BenchmarkMetricDecode(b *testing.B) {
	// This benchmark compares performance of full decoding and partial encoding of
	// the same metric data.

	tests := []struct {
		name            string
		messageProvider func() proto.Message
	}{
		{
			name:            "Full",
			messageProvider: func() proto.Message { return &otlp.ResourceMetrics{} },
		},
		{
			name:            "Partial",
			messageProvider: func() proto.Message { return &otlp.ResourceMetricsPrepared{} },
		},
	}

	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			// Dont'count time to generate and encode test data.
			b.StopTimer()
			batch := genFullMetrics(100)
			bytes, err := proto.Marshal(batch)
			if err != nil || len(bytes) == 0 {
				log.Fatal("Cannot encode batch")
			}

			runtime.GC()
			b.ResetTimer()

			// Measure decoding time.
			b.StartTimer()
			for i := 0; i < b.N; i++ {
				// Decode into a message.
				decoded := test.messageProvider()
				err := proto.Unmarshal(bytes, decoded)
				if err != nil || len(bytes) == 0 {
					log.Fatal("Cannot decode batch")
				}
			}
		})
	}
}
