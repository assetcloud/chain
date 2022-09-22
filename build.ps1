echo "building chain.exe chain-cli.exe"
$commitid = git rev-parse --short=8 HEAD
echo $commitid

$BUILDTIME=get-date -format "yyyy-MM-dd/HH:mm:ss"
echo $BUILDTIME

$BUILD_FLAGS='''-X "github.com/assetcloud/chain/common/version.GitCommit={0}" -X "github.com/assetcloud/chain/common/version.BuildTime={1}" -w -s''' -f $commitid,$BUILDTIME
echo $BUILD_FLAGS


go env -w CGO_ENABLED=1
go build  -ldflags  $BUILD_FLAGS  -v -o build/chain.exe github.com/assetcloud/chain/cmd/chain
go build  -ldflags  $BUILD_FLAGS  -v -o build/chain-cli.exe github.com/assetcloud/chain/cmd/cli

echo "build end"
