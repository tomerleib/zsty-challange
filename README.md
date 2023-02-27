# zsty-challange

This is an example app written in Golang that use AWS SDK to read keys from DynamoDB.

The router itself is built on [Gin Web Framework](https://gin-gonic.com/)

## Note
Since this is my first ever project in Golang, I did not included any tests.

## AWS Credentials
In order to use the AWS SDK, the user should use a shared configuration file under `~/.aws/credentials`.
In case such file does not exists, the user should create it by running the following command:
```bash
$ aws configure
```

## Test
To test the code here, use `verification.sh` code.