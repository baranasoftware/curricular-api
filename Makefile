# credential environment variables are defined in the below file
include baranasoftware.env
.PHONY: terraform 
default:
	cd src;go run main.go --local=true

terraform:
	cd src; GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go; \
	cd ../terraform; \
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=us-east-1 terraform init  -var aws_account_ids=${AWS_ACCOUNT_ID} -var sec_group=${SEC_GROUP};\
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=us-east-1 terraform plan  -var aws_account_ids=${AWS_ACCOUNT_ID} -var sec_group=${SEC_GROUP};\
	AWS_PROFILE=${AWS_PROFILE} AWS_REGION=us-east-1 terraform apply -var aws_account_ids=${AWS_ACCOUNT_ID} -var sec_group=${SEC_GROUP};

test:
	curl -X GET -H 'X-API-Key: ${X_API_KEY}' https://${AWS_API_ID}.execute-api.us-east-1.amazonaws.com/curricular-and-academic-api/students
