Please specify either:
• a "go_package" option in the .proto source file, or
• a "M" argument on the command line.

解决方案：
在syntax=”proto3″;下一行增加option go_package配置项。

命令：
protoc --go_out=./pb/ ./proto/*.proto
protoc --proto_path=proto proto/*.proto --go_out=pb