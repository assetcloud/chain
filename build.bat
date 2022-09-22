@echo on

set BUILDTIME=%date:~3%-%time:~0,2%:%time:~3,2%:%time:~6,2%
echo %BUILDTIME%

for /F "delims=" %%i in ('git rev-parse --short HEAD') do ( set commitid=%%i)

set BUILD_FLAGS=" -X github.com/assetcloud/chain/common/version.GitCommit=%commitid% -X github.com/assetcloud/chain/common/version.BuildTime=%BUILDTIME% -w -s"

go env -w CGO_ENABLED=1
go build  -ldflags  %BUILD_FLAGS% -v -o build/chain.exe github.com/assetcloud/chain/cmd/chain
go build  -ldflags  %BUILD_FLAGS% -v -o build/chain-cli.exe github.com/assetcloud/chain/cmd/cli
