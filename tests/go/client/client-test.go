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
  Run simple go client tests against the registered services
 ***/
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	common "github.com/scanoss/papi/api/commonv2"
	deps "github.com/scanoss/papi/api/dependenciesv2"
	scan "github.com/scanoss/papi/api/scanningv2"
)

// github.com/scanoss/papi/commonv1  - public
// github.com/scanoss/sapi/scannerv1 - private

const (
	defaultName = "go client"
)

var (
	d_addr = flag.String("d_addr", "localhost:50051", "dependency server to connec to")
	s_addr = flag.String("s_addr", "localhost:50052", "scanning server to connect to")
	name   = flag.String("name", defaultName, "Echo message")
)

func main() {
	flag.Parse()

	fmt.Println("Hello, World!")

	conn, err := grpc.Dial(*d_addr, grpc.WithInsecure()) // Set up a connection to the server.
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := deps.NewDependenciesClient(conn)
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Echo(ctx, &common.EchoRequest{Message: *name})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Dependency Echo: %s", r.GetMessage())

	conn2, err := grpc.Dial(*s_addr, grpc.WithInsecure()) // Set up a connection to the server.
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn2.Close()
	c2 := scan.NewScanningClient(conn2)
	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	r2, err := c2.Echo(ctx2, &common.EchoRequest{Message: *name})
	if err != nil {
		log.Fatalf("could not echo: %v", err)
	}
	log.Printf("Scanning Echo: %s", r2.GetMessage())

}
