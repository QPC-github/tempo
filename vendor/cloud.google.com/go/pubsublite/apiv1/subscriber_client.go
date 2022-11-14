// Copyright 2022 Google LLC
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

package pubsublite

import (
	"context"
	"math"

	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	pubsublitepb "google.golang.org/genproto/googleapis/cloud/pubsublite/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var newSubscriberClientHook clientHook

// SubscriberCallOptions contains the retry settings for each method of SubscriberClient.
type SubscriberCallOptions struct {
	Subscribe []gax.CallOption
}

func defaultSubscriberGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("pubsublite.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("pubsublite.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://pubsublite.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultSubscriberCallOptions() *SubscriberCallOptions {
	return &SubscriberCallOptions{
		Subscribe: []gax.CallOption{},
	}
}

// internalSubscriberClient is an interface that defines the methods availaible from Pub/Sub Lite API.
type internalSubscriberClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	Subscribe(context.Context, ...gax.CallOption) (pubsublitepb.SubscriberService_SubscribeClient, error)
}

// SubscriberClient is a client for interacting with Pub/Sub Lite API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// The service that a subscriber client application uses to receive messages
// from subscriptions.
type SubscriberClient struct {
	// The internal transport-dependent client.
	internalClient internalSubscriberClient

	// The call options for this service.
	CallOptions *SubscriberCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *SubscriberClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *SubscriberClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *SubscriberClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// Subscribe establishes a stream with the server for receiving messages.
func (c *SubscriberClient) Subscribe(ctx context.Context, opts ...gax.CallOption) (pubsublitepb.SubscriberService_SubscribeClient, error) {
	return c.internalClient.Subscribe(ctx, opts...)
}

// subscriberGRPCClient is a client for interacting with Pub/Sub Lite API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type subscriberGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing SubscriberClient
	CallOptions **SubscriberCallOptions

	// The gRPC API client.
	subscriberClient pubsublitepb.SubscriberServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewSubscriberClient creates a new subscriber service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// The service that a subscriber client application uses to receive messages
// from subscriptions.
func NewSubscriberClient(ctx context.Context, opts ...option.ClientOption) (*SubscriberClient, error) {
	clientOpts := defaultSubscriberGRPCClientOptions()
	if newSubscriberClientHook != nil {
		hookOpts, err := newSubscriberClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := SubscriberClient{CallOptions: defaultSubscriberCallOptions()}

	c := &subscriberGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		subscriberClient: pubsublitepb.NewSubscriberServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *subscriberGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *subscriberGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", versionClient, "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *subscriberGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *subscriberGRPCClient) Subscribe(ctx context.Context, opts ...gax.CallOption) (pubsublitepb.SubscriberService_SubscribeClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	var resp pubsublitepb.SubscriberService_SubscribeClient
	opts = append((*c.CallOptions).Subscribe[0:len((*c.CallOptions).Subscribe):len((*c.CallOptions).Subscribe)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.subscriberClient.Subscribe(ctx, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}