/**
 * Created by PhpStorm.
 * @file   build.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/27 21:14
 * @desc   build.go
 */

package hello

//go:generate protoc -I.  --go_out=impl --go_opt=paths=source_relative --go-grpc_out=impl --go-grpc_opt=paths=source_relative --grpc-gateway_out=impl --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=generate_unbound_methods=true --openapiv2_out=impl --openapiv2_opt=logtostderr=true message.proto
