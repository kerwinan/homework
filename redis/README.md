### redis

#### 1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

**Ubuntu Desktop, 32G, i7-9700K 8Core 3.60GHz**

> get/set
>
> redis-benchmark -h 127.0.0.1 -p 6379 -n 100000 -c 20 -d 10 -t get

**Raspberry Pi 4 Model B Rev 1.1, 8G, 4Core 1.5GHz**

#### 2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。