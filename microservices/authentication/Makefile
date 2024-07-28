PROJECT := $(shell cat go.mod | grep module | awk -F ' ' '{print $$2}' | awk -F '/' '{print $$NF}')

AUTOGEN := autogen
DOCS := $(AUTOGEN)/docs
VENDOR := vendor

GOLANGCI-LINT := github.com/golangci/golangci-lint/cmd/golangci-lint@latest

help: ## Показать описание основных команд
	@echo "\033[1;34mMakefile $(PROJECT) usage:\033[0m"
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

setup: ## Настроить рабочее окружение
	@echo -n "\033[1;34mSetup...\033[0m\t"
	@mkdir -p $(AUTOGEN)/server
	@echo "\033[0;32mOK\033[0m"

clean: ## Очистить рабочее окружение
	@echo -n "\033[1;34mClean...\033[0m\t"
	@go run $(GOLANGCI_LINT) cache clean
	@rm -rf $(VENDOR)
	@rm -rf $(AUTOGEN)
	@go clean -r -i -testcache -modcache
	@echo "\033[0;32mOK\033[0m"

swagger: setup ## Генерация swagger
# Генерируем swagger спецификацию на основе анотаций
	@echo -n "\033[1;34mSpec...\033[0m\t\t"
	@swag init \
		--parseDependency \
		-d cmd/api,internal/transport \
		-o $(DOCS) \
		-g main.go \
		>/dev/null 2>&1
	@echo "\033[0;32mOK\033[0m"
# Подготавливаем данные для подключения в docs.go
	@echo -n "\033[1;34mPrepare...\033[0m\t"
	@cat README.md | sed 's/"/\\"/g' > $(DOCS)/README.md
	@awk '{print "\"" $$0 "\\n\"+"}' $(DOCS)/README.md > $(DOCS)/README.data
	@sed -i -e '1 s/^/ Description:/;' $(DOCS)/README.data
	@echo "\"\"," >> $(DOCS)/README.data
	@echo "\033[0;32mOK\033[0m"
# Подключаем README.md в docs.go
	@echo -n "\033[1;34mDocs...\033[0m\t\t"
	@sed -e '/%README_FILE%/{' -e "r $(DOCS)/README.data" -e 'd' -e '}' $(DOCS)/docs.go > $(DOCS)/tmp.go
	@mv $(DOCS)/tmp.go $(DOCS)/docs.go
	@echo "\033[0;32mOK\033[0m"
# Генерируем router на основе swagger спецификаций
	@echo -n "\033[1;34mCodegen...\033[0m\t"
	@oapi-codegen --config=configs/codegen.yml $(DOCS)/swagger.yaml
	@echo "\033[0;32mOK\033[0m"

format: setup ## Запуск форматирования исходного кода
	@echo -n "\033[1;34mFormat...\033[0m\t"
	@go run $(GOLANGCI_LINT) run —fix ./...
	@echo "\033[0;32mOK\033[0m"

lint: setup ## Запуск линтера
	@echo -n "\033[1;34mLint...\033[0m\t"
	@go run $(GOLANGCI_LINT) run ./...
	@echo "\033[0;32mOK\033[0m"

all: swagger format lint ## Последовательный запуск основных команд

.DEFAULT_GOAL := all
