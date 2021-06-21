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

package automl_test

import (
	"context"

	automl "cloud.google.com/go/automl/apiv1"
	"google.golang.org/api/iterator"
	automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"
)

func ExampleNewClient() {
	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleClient_CreateDataset() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.CreateDatasetRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetDataset() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.GetDatasetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListDatasets() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.ListDatasetsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListDatasets(ctx, req)
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

func ExampleClient_UpdateDataset() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.UpdateDatasetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeleteDataset() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.DeleteDatasetRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteDataset(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ImportData() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.ImportDataRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ImportData(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ExportData() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.ExportDataRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportData(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetAnnotationSpec() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.GetAnnotationSpecRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetAnnotationSpec(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_CreateModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.CreateModelRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_GetModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.GetModelRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListModels() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.ListModelsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListModels(ctx, req)
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

func ExampleClient_DeleteModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.DeleteModelRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_UpdateModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.UpdateModelRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_DeployModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.DeployModelRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeployModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_UndeployModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.UndeployModelRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UndeployModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_ExportModel() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.ExportModelRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.ExportModel(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleClient_GetModelEvaluation() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.GetModelEvaluationRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetModelEvaluation(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleClient_ListModelEvaluations() {
	// import automlpb "google.golang.org/genproto/googleapis/cloud/automl/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := automl.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &automlpb.ListModelEvaluationsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListModelEvaluations(ctx, req)
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