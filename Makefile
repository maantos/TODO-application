swagger:
	@echo Generating swagger.json
	swagger generate spec -o ./pkg/docs/swagger.json --scan-models

swagger_run: swagger run
	@echo Generating swagger.yaml and run Api
run:
	@go run ./cmd/api