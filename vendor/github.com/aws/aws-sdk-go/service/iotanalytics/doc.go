// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotanalytics provides the client and types for making API
// requests to AWS IoT Analytics.
//
// AWS IoT Analytics allows you to collect large amounts of device data, process
// messages, and store them. You can then query the data and run sophisticated
// analytics on it. AWS IoT Analytics enables advanced data exploration through
// integration with Jupyter Notebooks and data visualization through integration
// with Amazon QuickSight.
//
// Traditional analytics and business intelligence tools are designed to process
// structured data. IoT data often comes from devices that record noisy processes
// (such as temperature, motion, or sound). As a result the data from these
// devices can have significant gaps, corrupted messages, and false readings
// that must be cleaned up before analysis can occur. Also, IoT data is often
// only meaningful in the context of other data from external sources.
//
// AWS IoT Analytics automates the steps required to analyze data from IoT devices.
// AWS IoT Analytics filters, transforms, and enriches IoT data before storing
// it in a time-series data store for analysis. You can set up the service to
// collect only the data you need from your devices, apply mathematical transforms
// to process the data, and enrich the data with device-specific metadata such
// as device type and location before storing it. Then, you can analyze your
// data by running queries using the built-in SQL query engine, or perform more
// complex analytics and machine learning inference. AWS IoT Analytics includes
// pre-built models for common IoT use cases so you can answer questions like
// which devices are about to fail or which customers are at risk of abandoning
// their wearable devices.
//
// See https://docs.aws.amazon.com/goto/WebAPI/iotanalytics-2017-11-27 for more information on this service.
//
// See iotanalytics package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/iotanalytics/
//
// Using the Client
//
// To contact AWS IoT Analytics with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS IoT Analytics client IoTAnalytics for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/iotanalytics/#New
package iotanalytics
