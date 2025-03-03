.DEFAULT_GOAL := pc

CURRENT_PATH=$(shell pwd)
.PHONY: engine init initweb fmt gen upload updatemod template

fmt:
	go fmt ./... 

updatemod:
	gdspx updatemod

ios:
	./tools/init_template.sh -p "ios"

template:
	@if [ "$(word 2,$(MAKECMDGOALS))" != "" ]; then \
		./tools/init_template.sh -p "$(word 2,$(MAKECMDGOALS))"; \
	else \
		./tools/init_template.sh; \
	fi

pc:
	./tools/init.sh

pcpack: 
	./tools/init.sh -a

web: 
	./tools/init_web.sh 

webpack: 
	./tools/init_web.sh -a

gen:
	cd ./cmd/codegen && go run . && cd $(CURRENT_PATH) && \
	$(MAKE) fmt && $(MAKE) fmt && \
	cd ./spx && go fmt ./... && cd $(CURRENT_PATH)
	
upload:
	./webserver/upload.sh 

%:
	@:
