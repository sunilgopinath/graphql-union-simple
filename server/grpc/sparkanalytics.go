/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package main

import (
	"log"
	"net"

	pb "bitbucket.org/ffxblue/golang-demo/generated/spark"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// type AssetReply_Asset struct {
// 	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
// 	Today   string `protobuf:"bytes,2,opt,name=today" json:"today,omitempty"`
// 	Overall string `protobuf:"bytes,3,opt,name=overall" json:"overall,omitempty"`
// }

// SayHello implements helloworld.GreeterServer
func (s *server) AssetStatistics(ctx context.Context, in *pb.AssetRequest) (*pb.AssetReply, error) {
	return &pb.AssetReply{Id: "1000", Current: 2004, Trending: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSparkAnalyticsServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
