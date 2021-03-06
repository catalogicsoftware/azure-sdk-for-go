package billing

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// InvoiceSectionsWithCreateSubscriptionPermissionClient is the billing client provides access to billing resources for
// Azure subscriptions.
type InvoiceSectionsWithCreateSubscriptionPermissionClient struct {
	BaseClient
}

// NewInvoiceSectionsWithCreateSubscriptionPermissionClient creates an instance of the
// InvoiceSectionsWithCreateSubscriptionPermissionClient client.
func NewInvoiceSectionsWithCreateSubscriptionPermissionClient(subscriptionID string) InvoiceSectionsWithCreateSubscriptionPermissionClient {
	return NewInvoiceSectionsWithCreateSubscriptionPermissionClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewInvoiceSectionsWithCreateSubscriptionPermissionClientWithBaseURI creates an instance of the
// InvoiceSectionsWithCreateSubscriptionPermissionClient client.
func NewInvoiceSectionsWithCreateSubscriptionPermissionClientWithBaseURI(baseURI string, subscriptionID string) InvoiceSectionsWithCreateSubscriptionPermissionClient {
	return InvoiceSectionsWithCreateSubscriptionPermissionClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// List lists all invoiceSections with create subscription permission for a user.
// Parameters:
// billingAccountName - billing Account Id.
// expand - may be used to expand the billingProfiles.
func (client InvoiceSectionsWithCreateSubscriptionPermissionClient) List(ctx context.Context, billingAccountName string, expand string) (result InvoiceSectionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/InvoiceSectionsWithCreateSubscriptionPermissionClient.List")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListPreparer(ctx, billingAccountName, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsWithCreateSubscriptionPermissionClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsWithCreateSubscriptionPermissionClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "billing.InvoiceSectionsWithCreateSubscriptionPermissionClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client InvoiceSectionsWithCreateSubscriptionPermissionClient) ListPreparer(ctx context.Context, billingAccountName string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountName": autorest.Encode("path", billingAccountName),
	}

	const APIVersion = "2018-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/listInvoiceSectionsWithCreateSubscriptionPermission", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client InvoiceSectionsWithCreateSubscriptionPermissionClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client InvoiceSectionsWithCreateSubscriptionPermissionClient) ListResponder(resp *http.Response) (result InvoiceSectionListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
