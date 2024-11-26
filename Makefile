.DEFAULT_GOAL := pc

CURRENT_PATH=$(shell pwd)
.PHONY: engine init initweb fmt gen upload updatemod

fmt:
	go fmt ./... 

updatemod:
	gdspx updatemod

pc:
	./tools/init.sh

web: 
	./tools/init_web.sh 

webpack: 
	./tools/init_web.sh -a

gen:
	cd ./cmd/codegen && go run . && cd $(CURRENT_PATH) && \
	$(MAKE) fmt && $(MAKE) fmt 
	
upload:
	./webserver/upload.sh 
