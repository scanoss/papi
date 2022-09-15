/*
 SPDX-License-Identifier: MIT

   Copyright (c) 2021, SCANOSS

   Permission is hereby granted, free of charge, to any person obtaining a copy
   of this software and associated documentation files (the "Software"), to deal
   in the Software without restriction, including without limitation the rights
   to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
   copies of the Software, and to permit persons to whom the Software is
   furnished to do so, subject to the following conditions:

   The above copyright notice and this permission notice shall be included in
   all copies or substantial portions of the Software.

   THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
   IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
   AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
   OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
   THE SOFTWARE.
*/

/***
  Provide a test instance of the Dependency Server
***/
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"

	common "github.com/scanoss/papi/api/commonv2"
	deps "github.com/scanoss/papi/api/dependenciesv2"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

var depRespJson = `{
  "files": [
    {
      "file": "audit-workbench-master/package.json",
      "id": "dependency",
      "status": "pending",
      "dependencies": [
        {
          "component": "abort-controller",
          "purl": "abort-controller",
          "version": "",
          "licenses": [
            {
              "name": "MIT"
            }
          ]
        },
        {
          "component": "chart.js",
          "purl": "chart.js",
          "version": "",
          "licenses": [
            {
              "name": "MIT"
            }
          ]
        }
      ]
    }
  ]
}
`

type server struct {
	deps.UnimplementedDependenciesServer
}

// Echo Implement the Echo gRPC
func (s *server) Echo(ctx context.Context, in *common.EchoRequest) (*common.EchoResponse, error) {
	log.Printf("Received: %v", in.GetMessage())

	// Uncomment to simulate long-running requests
	//log.Printf("Context: %#v - Error: %#v", ctx, ctx.Err())
	//for i := 0; i < 120; i++ {
	//	log.Printf("Checking: %v", ctx.Err())
	//	cte := ctx.Err()
	//	if cte == context.Canceled || cte == context.DeadlineExceeded {
	//		log.Printf("Cancelled")
	//		return nil, status.Errorf(codes.Canceled, "DepService.Echo canceled")
	//	}
	//	time.Sleep(2 * time.Second)
	//}
	return &common.EchoResponse{Message: in.GetMessage()}, nil
}

// GetDependencies Implement the gRPC call for getting dependencies
func (s *server) GetDependencies(ctx context.Context, in *deps.DependencyRequest) (*deps.DependencyResponse, error) {
	files := in.GetFiles()
	log.Printf("Received: %v - %v", in.GetDepth(), files)
	for _, file := range files {
		log.Printf("File: %v", file.GetFile())
		for _, purl := range file.GetPurls() {
			log.Printf("  Purl: %v  Req: %v", purl.GetPurl(), purl.GetRequirement())
		}
	}
	data, err := json.Marshal(in)
	if err != nil {
		log.Printf("Failed to marshall JSON from request: %v", err)
	} else {
		fmt.Println("JSON Data:", string(data))
	}
	var depResp deps.DependencyResponse
	err = json.Unmarshal([]byte(depRespJson), &depResp)
	if err != nil {
		log.Printf("Failed to marshall JSON from request: %v", err)
	}
	statusResp := common.StatusResponse{Status: common.StatusCode_SUCCESS, Message: "it worked"}
	return &deps.DependencyResponse{Files: depResp.Files, Status: &statusResp}, nil
}

// Launch the gRPC service and listen for requests
func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	deps.RegisterDependenciesServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
