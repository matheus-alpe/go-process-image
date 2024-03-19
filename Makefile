.PHONY: dev clear

clear:
	@rm -rf ./images/output/*

dev: clear
	@go run ./cmd
