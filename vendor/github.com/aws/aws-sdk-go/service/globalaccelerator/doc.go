// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package globalaccelerator provides the client and types for making API
// requests to AWS Global Accelerator.
//
// This is the AWS Global Accelerator API Reference. This guide is for developers
// who need detailed information about AWS Global Accelerator API actions, data
// types, and errors. For more information about Global Accelerator features,
// see the AWS Global Accelerator Developer Guide (https://docs.awa.amazon.com/global-accelerator/latest/dg/Welcome.html).
//
// AWS Global Accelerator is a network layer service in which you create accelerators
// to improve availability and performance for internet applications used by
// a global audience.
//
// Global Accelerator provides you with static IP addresses that you associate
// with your accelerator. These IP addresses are anycast from the AWS edge network
// and distribute incoming application traffic across multiple endpoint resources
// in multiple AWS Regions, which increases the availability of your applications.
// Endpoints can be Elastic IP addresses, Network Load Balancers, and Application
// Load Balancers that are located in one AWS Region or multiple Regions.
//
// Global Accelerator uses the AWS global network to route traffic to the optimal
// regional endpoint based on health, client location, and policies that you
// configure. The service reacts instantly to changes in health or configuration
// to ensure that internet traffic from clients is directed to only healthy
// endpoints.
//
// Global Accelerator includes components that work together to help you improve
// performance and availability for your applications:
//
// Static IP addressAWS Global Accelerator provides you with a set of static
// IP addresses which are anycast from the AWS edge network and serve as the
// single fixed points of contact for your clients. If you already have Elastic
// Load Balancing or Elastic IP address resources set up for your applications,
// you can easily add those to Global Accelerator to allow the resources to
// be accessed by a Global Accelerator static IP address.
//
// AcceleratorAn accelerator directs traffic to optimal endpoints over the AWS
// global network to improve availability and performance for your internet
// applications that have a global audience. Each accelerator includes one or
// more listeners.
//
// Network zoneA network zone services the static IP addresses for your accelerator
// from a unique IP subnet. Similar to an AWS Availability Zone, a network zone
// is an isolated unit with its own set of physical infrastructure. When you
// configure an accelerator, Global Accelerator allocates two IPv4 addresses
// for it. If one IP address from a network zone becomes unavailable due to
// IP address blocking by certain client networks, or network disruptions, then
// client applications can retry on the healthy static IP address from the other
// isolated network zone.
//
// ListenerA listener processes inbound connections from clients to Global Accelerator,
// based on the protocol and port that you configure. Each listener has one
// or more endpoint groups associated with it, and traffic is forwarded to endpoints
// in one of the groups. You associate endpoint groups with listeners by specifying
// the Regions that you want to distribute traffic to. Traffic is distributed
// to optimal endpoints within the endpoint groups associated with a listener.
//
// Endpoint groupEach endpoint group is associated with a specific AWS Region.
// Endpoint groups include one or more endpoints in the Region. You can increase
// or reduce the percentage of traffic that would be otherwise directed to an
// endpoint group by adjusting a setting called a traffic dial. The traffic
// dial lets you easily do performance testing or blue/green deployment testing
// for new releases across different AWS Regions, for example.
//
// EndpointAn endpoint is an Elastic IP address, Network Load Balancer, or Application
// Load Balancer. Traffic is routed to endpoints based on several factors, including
// the geo-proximity to the user, the health of the endpoint, and the configuration
// options that you choose, such as endpoint weights. You can configure weights
// for each endpoint, which are numbers that you can use to specify the proportion
// of traffic to route to each one. This can be useful, for example, to do performance
// testing within a Region.
//
// See https://docs.aws.amazon.com/goto/WebAPI/globalaccelerator-2018-08-08 for more information on this service.
//
// See globalaccelerator package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/globalaccelerator/
//
// Using the Client
//
// To contact AWS Global Accelerator with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS Global Accelerator client GlobalAccelerator for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/globalaccelerator/#New
package globalaccelerator
