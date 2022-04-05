# 最终生成的可执行文件
TARGET := "./output/target"
# 测试覆盖率报告
COVERAGE := "./output/coverage.out"
COVERAGE_HTML := "./output/coverage.html"


.PHONY : all clean generate test compile
all : clean generate  test compile

clean :
	@-rm -rf output
	@-mkdir output

# 一些代码生成的步骤
generate :
	@cd cmd && wire

test :
	@go test -coverprofile=${COVERAGE} ./...
	@go tool cover -html=${COVERAGE} -o=${COVERAGE_HTML}

compile :
	@go build -o ${TARGET} ./cmd


