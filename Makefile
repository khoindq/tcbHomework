default: help
BLACK        := $(shell tput -Txterm setaf 0)
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
PURPLE       := $(shell tput -Txterm setaf 5)
BLUE         := $(shell tput -Txterm setaf 6)
WHITE        := $(shell tput -Txterm setaf 7)

RESET := $(shell tput -Txterm sgr0)


# set target color
TARGET_COLOR := $(BLUE)

colors: ## - Show all the colors
	@echo "${BLACK}BLACK${RESET}"
	@echo "${RED}RED${RESET}"
	@echo "${GREEN}GREEN${RESET}"
	@echo "${YELLOW}YELLOW${RESET}"
	@echo "${LIGHTPURPLE}LIGHTPURPLE${RESET}"
	@echo "${PURPLE}PURPLE${RESET}"
	@echo "${BLUE}BLUE${RESET}"
	@echo "${WHITE}WHITE${RESET}"


.PHONY: help
help: ## - Show help message
	@printf "${TARGET_COLOR}Usage: make\n${RESET}"
	@awk 'BEGIN {FS = ":.*##";} /^[$$()% a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
	@printf "\n"
.DEFAULT_GOAL := help

check_defined = \
    $(strip $(foreach 1,$1, \
        $(call __check_defined,$1,$(strip $(value 2)))))
__check_defined = \
    $(if $(value $1),, \
      $(error ${RED} $1$(if $2, ($2)) is required ${RESET}))



#======================= Commands =======================#

# Docs
api-doc: ## - Generate docs
	@echo "${TARGET_COLOR} Generating api docs... !${RESET}"
	swag init  -g ./cmd/poolservice/main.go -o ./docs/poolservice 

# Run
poolservice-run: ## - Run pool service 
	@echo "${TARGET_COLOR} Running poolservice... !${RESET}"
	go run cmd/poolservice/main.go

# Gen mock interface
gen-mock: ## Generate mock interface
	find . -name 'mock_*' -delete
	go install github.com/vektra/mockery/v2/.../ && \
	mockery --all --case underscore --inpackage

# Tesing
test:  ## run test, and get coverage
	go test -cover ./...



