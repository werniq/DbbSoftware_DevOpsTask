## DBB Software Test Task

In this repository you can find three branches, each one for the part of the test task.
1. **feature/python-health-app**: This branch contains the code for the Python Health App, with a Dockerfile to build the image
2. **feature/aws-cdk-setup**: Branch for the AWS CDK setup, with the required infrastructure to deploy the Python Health App
3. **feature/deploy-ci-cd**: Branch with the CI/CD pipeline, which build Docker image, pushes it to the ECR, and deploys application to the AWS Elastic Beanstalk, creating new environment

## Pre-requisites

Before running the code, You need to create an application and a role for the Elastic Beanstalk environment, with following permissions:
![img.png](images/img.png)

I've tried to create this Role using CDK, but even with specifying same policies - I could not create environments with that role. 

## Results