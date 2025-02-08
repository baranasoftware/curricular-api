# credential enviornment variables are defined in below file
include baranasoftware.env
default:
	cd src;go run main.go --local=true

terraform:
	cd src; GOOS=linux GOARCH=arm64 go build -tags lambda.norpc -o bootstrap main.go; \
	cd terraform; \
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=us-east-1 terraform init  -var account_id=${AWS_ACCOUNT_ID}; \
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=us-east-1 terraform plan  -var version_tag=0.0.1; \
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=us-east-1 terraform apply -var version_tag=0.0.1;
