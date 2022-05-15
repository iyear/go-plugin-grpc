## Go Plugin over gRPC

施工中，非生产可用

你可以自由地分发使用该框架

example: [plus](./examples/plus)

在 `Core` 端可以看到类似如下输出：
```
2022/05/15 23:48:09 core [INFO] bind plugin math.v1
2022/05/15 23:48:10 core [DEBUG] checking health of plugin math.v1: 1652629689/1652629675
2022/05/15 23:48:12 core [DEBUG] checking health of plugin math.v1: 1652629689/1652629677
2022/05/15 23:48:14 core [DEBUG] checking health of plugin math.v1: 1652629689/1652629679
start call: 2 + 3 = ?
2022/05/15 23:48:14 math.v1 [INFO] enter math.v1.plus
2022/05/15 23:48:14 math.v1 [INFO] finish plus func
call finished. result map: map[V:5], err: <nil>
result: 5 ,time: 1.0838ms
2022/05/15 23:48:16 core [DEBUG] checking health of plugin math.v1: 1652629689/1652629681
2022/05/15 23:48:18 core [DEBUG] checking health of plugin math.v1: 1652629689/1652629683
2022/05/15 23:48:20 core [DEBUG] checking health of plugin math.v1: 1652629699/1652629685
```