health/health.pb.go: health/health.proto
	protoc health/health.proto --go_out=plugins=grpc:.

exchange.pb.go: exchange.proto
	protoc exchange.proto --go_out=plugins=grpc:.

generate: exchange.pb.go health/health.pb.go
