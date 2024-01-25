package main

import (
    "context"
    "log"
    "net"
    "time"
    "google.golang.org/grpc"
    pb "github.com/mingard/log-compression/logcompressionpb"
)

type server struct {
    pb.UnimplementedLogCompressionServiceServer
    enumMap     map[string]int32
    keyInfoMap  map[string]*keyInfo
    windowSize  time.Duration
    totalSize   int
    totalCount  int
}

type keyInfo struct {
    frequency  int
    size       int
    timestamps []time.Time
}

func (s *server) CompressLog(ctx context.Context, in *pb.LogMessage) (*pb.CompressionMapping, error) {
    currentTime := time.Now()
    messageSize := len(in.Message)

    s.totalCount++
    s.totalSize += messageSize

    if info, exists := s.keyInfoMap[in.Message]; exists {
        info.frequency++
        info.timestamps = append(info.timestamps, currentTime)
    } else {
        s.keyInfoMap[in.Message] = &keyInfo{
            frequency:  1,
            size:       messageSize,
            timestamps: []time.Time{currentTime},
        }
    }

    response := &pb.CompressionMapping{
        EnumMapping: make(map[string]int32),
    }

    for value, info := range s.keyInfoMap {
        for len(info.timestamps) > 0 && currentTime.Sub(info.timestamps[0]) > s.windowSize {
            info.timestamps = info.timestamps[1:]
            info.frequency--
        }

        if shouldCompress(info) {
            if _, exists := s.enumMap[value]; !exists {
                s.enumMap[value] = int32(len(s.enumMap) + 1)
            }
            response.EnumMapping[value] = s.enumMap[value]
        }
    }

    avgSize := s.totalSize / s.totalCount
    s.windowSize = calculateWindowSize(s.totalCount, avgSize)

    return response, nil
}

func shouldCompress(info *keyInfo) bool {
    // Implement your compression logic here
    return true
}

func calculateWindowSize(totalCount int, avgSize int) time.Duration {
    // Example logic to calculate window size
    return time.Duration(avgSize*totalCount) * time.Millisecond
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterLogCompressionServiceServer(s, &server{
        enumMap:    make(map[string]int32),
        keyInfoMap: make(map[string]*keyInfo),
        windowSize: 1 * time.Minute, // Initial window size
    })
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}