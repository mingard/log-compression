package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/mingard/log-compression/logcompressionpb"

	"google.golang.org/grpc"
)

func compressLogMessage(logMessage *pb.LogMessage, mapping *pb.CompressionMapping) (string, string) {
	original := fmt.Sprintf("%+v", *logMessage)

	if enum, ok := mapping.EnumMapping[logMessage.Message]; ok {
		logMessage.Message = fmt.Sprintf("ENUM%d", enum)
	}

	compressed := fmt.Sprintf("%+v", *logMessage)
	return original, compressed
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLogCompressionServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	logs := []*pb.LogMessage{
		// {Name:"dadi-dadi-api",Hostname:"ip-172-31-1-187",PID:3082692,Level:50,Module:"api",Msg:"Error: Validation failed\n    at Model._createValidationError (/app/foundland/api/node_modules/@dadi/api/dadi/lib/model/index.js:204:17)\n    at /app/foundland/api/node_modules/@dadi/api/dadi/lib/model/index.js:714:12\n    at runMicrotasks (<anonymous>)\n    at processTicksAndRejections (node:internal/process/task_queues:96:5)\n    at async Collection.<anonymous> (/app/foundland/api/node_modules/@dadi/api/dadi/lib/controller/documents.js:177:21) {\n  statusCode: 400,\n  success: false,\n  errors: [\n    {\n      field: 'allCategories',\n      message: 'not master and slaveOk=false',\n      code: 13435\n    }\n  ]\n}",Timestamp:"2024-01-23T15:11:59.242Z",V:0},
		{Timestamp: "2024-01-25T12:00:00Z", Level: "ERROR", Message: "Error finding account for user ID: 123"},
		{Timestamp: "2024-01-25T12:01:00Z", Level: "ERROR", Message: "Error finding account for user ID: 456"},
		{Timestamp: "2024-01-25T12:02:00Z", Level: "INFO", Message: "User 123 logged in successfully"},
		{Timestamp: "2024-01-25T12:03:00Z", Level: "WARN", Message: "Missing user data for ID: 789"},
		{Timestamp: "2024-01-25T12:04:00Z", Level: "ERROR", Message: "Failed to connect to database"},
		{Timestamp: "2024-01-25T12:05:00Z", Level: "INFO", Message: "Service started on port 8080"},
		{Timestamp: "2024-01-25T12:06:00Z", Level: "ERROR", Message: "Error finding account for user ID: 321"},
		{Timestamp: "2024-01-25T12:07:00Z", Level: "INFO", Message: "Configuration file loaded"},
		{Timestamp: "2024-01-25T12:08:00Z", Level: "ERROR", Message: "Invalid user credentials provided"},
		{Timestamp: "2024-01-25T12:09:00Z", Level: "INFO", Message: "Cache cleared successfully"},
		{Timestamp: "2024-01-25T12:10:00Z", Level: "INFO", Message: "New user registered: user@example.com"},
		{Timestamp: "2024-01-25T12:11:00Z", Level: "ERROR", Message: "Payment processing failed"},
		{Timestamp: "2024-01-25T12:12:00Z", Level: "ERROR", Message: "Error finding account for user ID: 654"},
		{Timestamp: "2024-01-25T12:13:00Z", Level: "WARN", Message: "Low disk space warning"},
		{Timestamp: "2024-01-25T12:14:00Z", Level: "ERROR", Message: "Database query timeout"},
		{Timestamp: "2024-01-25T12:15:00Z", Level: "INFO", Message: "Email sent to user@example.com"},
		{Timestamp: "2024-01-25T12:16:00Z", Level: "INFO", Message: "Backup completed successfully"},
		{Timestamp: "2024-01-25T12:17:00Z", Level: "ERROR", Message: "Failed to load external API data"},
		{Timestamp: "2024-01-25T12:18:00Z", Level: "INFO", Message: "System health check OK"},
		{Timestamp: "2024-01-25T12:19:00Z", Level: "ERROR", Message: "User authentication failed"},
	}

	for _, logMessage := range logs {
		r, err := c.CompressLog(ctx, logMessage)
		if err != nil {
			log.Fatalf("could not compress log: %v", err)
		}

		original, compressed := compressLogMessage(logMessage, r)

		// Calculate the size of the original message
		originalSize := len(original)

		// Calculate the size of the compressed message
		compressedSize := len(compressed)

		// Log the sizes
		log.Printf("Original size: %d bytes, Compressed size: %d bytes", originalSize, compressedSize)

		fmt.Printf("Original: %s\n", original)
		fmt.Printf("Compressed: %s\n", compressed)
		fmt.Printf("Enum Response: %+v\n", r)
	}
}
