.PHONY: all genprotobuf benchmark

GO=$(shell which go)

all: genprotobuf benchmark

benchmark:
	${GO} test -v
	sudo nice -n -5 ${GO} test -bench . -benchtime 5s

genprotobuf:
	protoc -I/usr/local/include -I otlp/ otlp/metric_data.proto --go_out=plugins=grpc:otlp

