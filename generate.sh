protoc  --go_out=. -I=server/echopb server/echopb/echo.proto \
        --go-grpc_out=. -I=server/echopb server/echopb/echo.proto \
        