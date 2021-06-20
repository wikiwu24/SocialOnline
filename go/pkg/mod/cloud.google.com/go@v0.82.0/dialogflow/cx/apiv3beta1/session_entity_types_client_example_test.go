// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package cx_test

import (
	"context"

	cx "cloud.google.com/go/dialogflow/cx/apiv3beta1"
	"google.golang.org/api/iterator"
	cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
)

func ExampleNewSessionEntityTypesClient() {
	ctx := context.Background()
	c, err := cx.NewSessionEntityTypesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleSessionEntityTypesClient_ListSessionEntityTypes() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := cx.NewSessionEntityTypesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.ListSessionEntityTypesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListSessionEntityTypes(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleSessionEntityTypesClient_GetSessionEntityType() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewSessionEntityTypesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.GetSessionEntityTypeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetSessionEntityType(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSessionEntityTypesClient_CreateSessionEntityType() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewSessionEntityTypesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.CreateSessionEntityTypeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateSessionEntityType(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSessionEntityTypesClient_UpdateSessionEntityType() {
	// import cxpb "google.golang.org/genproto/googleapis/cloud/dialogflow/cx/v3beta1"

	ctx := context.Background()
	c, err := cx.NewSessionEntityTypesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.UpdateSessionEntityTypeRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateSessionEntityType(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleSessionEntityTypesClient_DeleteSessionEntityType() {
	ctx := context.Background()
	c, err := cx.NewSessionEntityTypesClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &cxpb.DeleteSessionEntityTypeRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteSessionEntityType(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}