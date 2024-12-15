package main

import (
	cdk "github.com/aws/aws-cdk-go/awscdk/v2"
	ec2 "github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	ecr "github.com/aws/aws-cdk-go/awscdk/v2/awsecr"
	"os"

	// "github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

const (
	// Default CIDR block for the VPC.
	// This is the IP address range that will be used for the VPC.
	// As a default, it is class A network.
	vpcCidrDefault = "10.0.0.0/16"

	// Default mask for the subnets.
	// This will affect the number of available IP addresses in the subnet.
	// Recommend you to use it at least /26 for the subnets.
	// The default value is /24.
	subnetsMaskDefault = 24

	// Default name for the EKS cluster. This is the name of the EKS cluster that will be created.
	ecrRepositoryName = "dbbsoftware-test-task-ecr-repository"

	// Default name for	the VPC. This is the name of the VPC that will be created.
	vpcName = "dbbsoftware-test-task-vpc"

	// createInternetGatewayDefault is a flag that indicates whether the default internet gateway will be created.
	createInternetGatewayDefault = false

	// enableDnsHostnamesDefault is a flag that indicates whether the DNS hostnames will be enabled.
	enableDnsHostnamesDefault = true

	// enableDnsSupportDefault is a flag that indicates whether the DNS support will be enabled.
	enableDnsSupportDefault = true
)

type AwsStackProps struct {
	cdk.StackProps
}

func NewAwsStack(scope constructs.Construct, id string, props *AwsStackProps) cdk.Stack {
	var sprops cdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := cdk.NewStack(scope, &id, &sprops)

	ecrRepo := ecr.NewRepository(stack, jsii.String("DBBSoftwareTestTaskECR"), &ecr.RepositoryProps{
		ImageScanOnPush: jsii.Bool(true),
		RepositoryName:  jsii.String(ecrRepositoryName),
	})

	vpc := ec2.NewVpc(stack, jsii.String("DBBSoftwareTestTaskVPC"), &ec2.VpcProps{
		CreateInternetGateway: jsii.Bool(createInternetGatewayDefault),
		EnableDnsHostnames:    jsii.Bool(enableDnsHostnamesDefault),
		EnableDnsSupport:      jsii.Bool(enableDnsSupportDefault),
		IpAddresses:           ec2.IpAddresses_Cidr(jsii.String(vpcCidrDefault)),
		IpProtocol:            ec2.IpProtocol_IPV4_ONLY,
		MaxAzs:                jsii.Number[int](3),
		SubnetConfiguration: &[]*ec2.SubnetConfiguration{
			{
				Name:                jsii.String("PublicSubnet"),
				SubnetType:          ec2.SubnetType_PUBLIC,
				CidrMask:            jsii.Number[int](subnetsMaskDefault),
				MapPublicIpOnLaunch: jsii.Bool(true),
			},
		},
		VpcName: jsii.String(vpcName),
	})

	eksSecurityGroup := ec2.NewSecurityGroup(stack, jsii.String("DBBSoftwareTestTaskVpcSg"), &ec2.SecurityGroupProps{
		Vpc:               vpc,
		Description:       jsii.String("Allow access to EKS Cluster"),
		SecurityGroupName: jsii.String("dbb-software-test-task-vpc-sg"),
	})

	eksSecurityGroup.AddIngressRule(
		ec2.Peer_AnyIpv4(),
		ec2.Port_AllTraffic(),
		jsii.String("Allow all inbound traffic"),
		jsii.Bool(false),
	)

	addOutput(stack, "ECR Repository Name", "ECRRepositoryName", ecrRepo.RepositoryName())
	addOutput(stack, "VPC ID", "VPCID", vpc.VpcId())
	addOutput(stack, "VPC CIDR", "VPCCIDR", vpc.VpcCidrBlock())
	addOutput(stack, "Security Group ID", "SecurityGroupID", eksSecurityGroup.SecurityGroupId())

	return stack
}

func addOutput(stack cdk.Stack, desc string, displayName string, value *string) {
	cdk.NewCfnOutput(stack, jsii.String(desc), &cdk.CfnOutputProps{
		ExportName:  jsii.String(displayName),
		Value:       value,
		Description: jsii.String(desc),
	})
}

func main() {
	defer jsii.Close()

	app := cdk.NewApp(nil)

	NewAwsStack(app, "DBBSoftwareTestTaskAwsStack", &AwsStackProps{
		cdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *cdk.Environment {
	return &cdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
