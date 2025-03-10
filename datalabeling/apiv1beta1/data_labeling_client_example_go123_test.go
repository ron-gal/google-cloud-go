// Copyright 2025 Google LLC
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

//go:build go1.23

package datalabeling_test

import (
	"context"

	datalabeling "cloud.google.com/go/datalabeling/apiv1beta1"
	datalabelingpb "cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb"
)

func ExampleClient_ListAnnotatedDatasets_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListAnnotatedDatasetsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListAnnotatedDatasetsRequest.
	}
	for resp, err := range c.ListAnnotatedDatasets(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListAnnotationSpecSets_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListAnnotationSpecSetsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListAnnotationSpecSetsRequest.
	}
	for resp, err := range c.ListAnnotationSpecSets(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListDataItems_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListDataItemsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListDataItemsRequest.
	}
	for resp, err := range c.ListDataItems(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListDatasets_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListDatasetsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListDatasetsRequest.
	}
	for resp, err := range c.ListDatasets(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListEvaluationJobs_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListEvaluationJobsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListEvaluationJobsRequest.
	}
	for resp, err := range c.ListEvaluationJobs(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListExamples_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListExamplesRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListExamplesRequest.
	}
	for resp, err := range c.ListExamples(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_ListInstructions_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.ListInstructionsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#ListInstructionsRequest.
	}
	for resp, err := range c.ListInstructions(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_SearchEvaluations_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.SearchEvaluationsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#SearchEvaluationsRequest.
	}
	for resp, err := range c.SearchEvaluations(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleClient_SearchExampleComparisons_all() {
	ctx := context.Background()
	// This snippet has been automatically generated and should be regarded as a code template only.
	// It will require modifications to work:
	// - It may require correct/in-range values for request initialization.
	// - It may require specifying regional endpoints when creating the service client as shown in:
	//   https://pkg.go.dev/cloud.google.com/go#hdr-Client_Options
	c, err := datalabeling.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	defer c.Close()

	req := &datalabelingpb.SearchExampleComparisonsRequest{
		// TODO: Fill request struct fields.
		// See https://pkg.go.dev/cloud.google.com/go/datalabeling/apiv1beta1/datalabelingpb#SearchExampleComparisonsRequest.
	}
	for resp, err := range c.SearchExampleComparisons(ctx, req).All() {
		if err != nil {
			// TODO: Handle error and break/return/continue. Iteration will stop after any error.
		}
		// TODO: Use resp.
		_ = resp
	}
}
