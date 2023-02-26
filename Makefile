swagger:
	@echo Generating swagger.yaml
	swagger generate spec -o ./swagger.json --scan-models

swagger_run: swagger run
	@echo Generating swagger.yaml and run Api
run:
	@go run ./cmd/api