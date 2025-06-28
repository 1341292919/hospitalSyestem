MODULE = Hospital

SERVICE_NAME = hospital

.PHONY:target
target:
		sh build.sh
		sh output/bootstrap.sh

.PHONY:new
new:
	hz new \
	-module $(MODULE) \
	-service "$(SERVICE_NAME)"
	hz update -idl ./idl/model.thrift
	hz update -idl ./idl/user.thrift
	hz update -idl ./idl/epidemic.thrift
	hz update -idl ./idl/clinical.thrift
	hz update -idl ./idl/visit.thrift
	hz update -idl ./idl/biobank.thrift

.PHONY: gen
gen:
	hz update -idl ./idl/model.thrift
	hz update -idl ./idl/user.thrift
	hz update -idl ./idl/epidemic.thrift
	hz update -idl ./idl/clinical.thrift
	hz update -idl ./idl/visit.thrift
	hz update -idl ./idl/biobank.thrift

# 启动必要的环境，比如 etcd、mysql
.PHONY: env-up
env-up:
	@ docker compose -f ./docker/docker-compose.yml up -d

# 关闭必要的环境，但不清理 data（位于 docker/data 目录中）
.PHONY: env-down
env-down:
	@ cd ./docker && docker compose down

# 清除所有的构建产物
.PHONY: clean
clean:
	@find . -type d -name "output" -exec rm -rf {} + -print

# 清除所有构建产物、compose 环境和它的数据
.PHONY: clean-all
clean-all: clean
	@echo "$(PREFIX) Checking if docker-compose services are running..."
	@docker compose -f ./docker/docker-compose.yml ps -q | grep '.' && docker compose -f ./docker/docker-compose.yml down || echo "$(PREFIX) No services are running."
	@echo "$(PREFIX) Removing docker data..."
	rm -rf ./docker/data
