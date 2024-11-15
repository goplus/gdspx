.DEFAULT_GOAL := pc

CURRENT_PATH=$(shell pwd)
.PHONY: engine init initweb fmt gen upload

fmt:
	go fmt ./... 

pc:
	./tools/init.sh

web: 
	./tools/init_web.sh 

gen:
	cd ./cmd/codegen && go run . && cd $(CURRENT_PATH) && \
	$(MAKE) fmt && $(MAKE) fmt 
	
upload:
	./webserver/upload.sh 
