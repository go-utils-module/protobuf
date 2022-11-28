/**
 * Created by PhpStorm.
 * @file   build.go
 * @author 李锦 <Lijin@cavemanstudio.net>
 * @date   2022/11/27 21:14
 * @desc   build.go
 */

package notification

//go:generate protoc -I.  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out=. --grpc-gateway_opt=paths=source_relative --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=generate_unbound_methods=true --openapiv2_out=. --openapiv2_opt=logtostderr=true message.proto
