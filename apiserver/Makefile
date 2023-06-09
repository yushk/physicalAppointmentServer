V := 1 # When V is set, print commands and build progress.
M :=  # When M is set, build with -mod vendor.

.PHONY: build
build:
	$(GO) build $(FLAGS)apiserver

tidy:
	go mod tidy

gen: clean server

.PHONY: server
server:
	swagger generate server -P v1.Principal -A sport -t ./ -f ./swagger/swagger.yaml --model-package=v1 && \
	cp ./cmd/sport-server/main.go ./


.PHONY: clean
clean:
	rm -rf ./client && \
	rm -rf ./cmd && \
	rm -rf ./main.go && \
	rm -rf ./restapi/operations && \
	rm -rf ./restapi/doc.go && \
	rm -rf ./restapi/embedded_spec.go && \
	rm -rf ./restapi/server.go && \
	rm -rf ./v1 && \
	rm -rf ./swagger/models.json && \
	rm -rf ./bin && \
	rm -rf ./v1

##### ^^^^^^ EDIT ABOVE ^^^^^^ #####

##### =====> Internals <===== #####

# 版本号 v1.0.3-6-g0c2b1cf-dev
# 1、6:表示自打tag v1.0.3以来有6次提交（commit）
# 2、g0c2b1cf：g 为git的缩写，在多种管理工具并存的环境中很有用处
# 3、0c2b1cf：7位字符表示为最新提交的commit id 前7位
# 4、如果本地仓库有修改，则认为是dirty的，则追加-dev，表示是开发版：v1.0.3-6-g0c2b1cf-dev
VERSION          := $(shell git describe --tags --always --dirty="-dev")

# 时间
DATE             := $(shell date -u '+%Y-%m-%d-%H%M UTC')

# 版本标志  -s -w 缩小可执行文件大小
VERSION_FLAGS    := -ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)" -s -w'

# go arch
GOARCH			 := $(shell go env GOARCH)

# docker tag
DOCKER_TAG       := 0.2.3-dev

# 输出文件夹
OUTPUT_DIR       := -o ./bin/

# 标志
FLAGS            := $(if $V,-v) $(if $M,-mod vendor) $(VERSION_FLAGS) $(OUTPUT_DIR)

GO        		 := CGO_ENABLED=0 GO111MODULE=on go
CGO        		 := CGO_ENABLED=1 GO111MODULE=on go
