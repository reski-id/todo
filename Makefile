swag-init:
	swag init --parseInternal --exclude build,deploy,docs,scripts,vendor -g main.go

swag-fmt:
	swag fmt --exclude build,developments,docs,scripts -g main.go 

run:
	go run main.go