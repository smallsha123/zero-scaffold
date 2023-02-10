#!/bin/bash

# eg: ./build.sh api-common.test
# eg: ./build.sh api-admin.test
# eg: ./build.sh api-boss.test
# eg: ./build.sh api-user.test

server=`echo $1 | sed -r 's/([^\.]+)?\.([^\.]+)?/\1/'`
cluster=`echo $1 | sed -r 's/([^\.]+)?\.([^\.]+)?/\2/'`
server_tag=`echo $server | sed -r 's/([^\-]+)?\-([^\-]+)?/\2/'`

exe="go.shop."$server_tag
#remote="root@39.106.131.145"
#workspace="/www/wwwroot/go.shop."$server_tag

echo "==========================="
echo "编译"
echo "==========================="
echo ""
#linux系统的
#go mod tidy && GOOS=linux GOARCH=amd64 go build -o ${exe} ${server}/core.go
#苹果系统
go mod tidy && GOOS=darwin GOARCH=amd64 go build -o ${exe} ${server}/core.go

if [ $? -ne 0 ]; then
    echo "go build failed"
    exit 1
fi
#
#echo "==========================="
#echo "部署"
#echo "==========================="
#echo ""
#
#builder_path="/tmp/publish"
#
#rm -rf ${builder_path}
#mkdir -p ${builder_path}/logs ${builder_path}/etc
#
#mv ${exe} ${builder_path}
#cp -r ${server}/etc/${cluster}/* ${builder_path}/etc
#
#cd ${builder_path} && zip -q -r ${exe}.zip *
#
#ssh ${remote} "mkdir -p ${workspace}"
#scp ${exe}.zip ${remote}:${workspace}
#ssh ${remote} "cd ${workspace} && unzip -o ${exe}.zip && chmod +x ${exe}"
#
#echo "==========================="
#echo "启动 ${exe} 服务"
#echo "==========================="
#echo ""
#
#ssh $remote "/www/server/panel/pyenv/bin/supervisorctl restart ${exe}:${exe}_00"