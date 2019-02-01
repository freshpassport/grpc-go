/*
 *
 * Copyright 2019 gRPC authors.
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

package test

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	testpb "google.golang.org/grpc/test/grpc_testing"
)

func (s) TestContextCanceled(t *testing.T) {
	log.Println("PID:", os.Getpid())

	ss := &stubServer{
		unaryCall: func(ctx context.Context, in *testpb.SimpleRequest) (*testpb.SimpleResponse, error) {
			return &testpb.SimpleResponse{Payload: &testpb.Payload{
				Body: make([]byte, 5*1024*1024),
			}}, nil
		},
	}
	if err := ss.Start(nil); err != nil {
		t.Fatalf("Error starting endpoint server: %v", err)
	}
	defer ss.Stop()
	ss.client.
	for i := 0; i < 200; i++ {
		_, err := ss.client.UnaryCall(context.Background(), &testpb.SimpleRequest{})
		log.Printf(`client.UnaryCall %v`, err)

		time.Sleep(time.Second)
	}
}
