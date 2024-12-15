package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/assertions"
	"github.com/aws/jsii-runtime-go"
	"testing"
)

func TestAwsStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewAwsStack(app, "TestAwsStack", nil)

	// THEN
	template := assertions.Template_FromStack(stack, nil)

	// Check for ECR Repository
	template.HasResourceProperties(jsii.String("AWS::ECR::Repository"), map[string]interface{}{
		"ImageScanningConfiguration": map[string]interface{}{
			"scanOnPush": jsii.Bool(true),
		},
		"RepositoryName": jsii.String(ecrRepositoryName),
	})

	// Check for VPC
	template.HasResourceProperties(jsii.String("AWS::EC2::VPC"), map[string]interface{}{
		"CidrBlock":          jsii.String(vpcCidrDefault),
		"EnableDnsSupport":   jsii.Bool(enableDnsSupportDefault),
		"EnableDnsHostnames": jsii.Bool(enableDnsHostnamesDefault),
	})

	// Check for Subnets
	template.HasResourceProperties(jsii.String("AWS::EC2::Subnet"), map[string]interface{}{
		"MapPublicIpOnLaunch": jsii.Bool(true),
	})

	// Check for Security Group
	template.HasResourceProperties(jsii.String("AWS::EC2::SecurityGroup"), map[string]interface{}{
		"GroupDescription": jsii.String("Allow access to EKS Cluster"),
		"SecurityGroupEgress": []interface{}{
			map[string]interface{}{
				"CidrIp":     jsii.String("0.0.0.0/0"),
				"IpProtocol": jsii.String("-1"),
			},
		},
	})

	// Check for Security Group Ingress Rule
	template.HasResourceProperties(jsii.String("AWS::EC2::SecurityGroupIngress"), map[string]interface{}{
		"CidrIp":     jsii.String("0.0.0.0/0"),
		"IpProtocol": jsii.String("-1"),
	})

	// Check for Outputs
	template.HasOutput(jsii.String("ECRRepositoryName"), map[string]interface{}{
		"Value": jsii.String(ecrRepositoryName),
	})
	template.HasOutput(jsii.String("VPCID"), map[string]interface{}{})
	template.HasOutput(jsii.String("VPCCIDR"), map[string]interface{}{})
	template.HasOutput(jsii.String("SecurityGroupID"), map[string]interface{}{})
}
