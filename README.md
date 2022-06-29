# mycli
mycli is one tools , for unix \ net \ k8s \ container \ vm


### User

##### ping

example: `mycli linux ping www.baidu.com `



```
# dns-ip解析：
PING www.baidu.com (110.242.68.3):
32 bytes from 110.242.68.3: icmp_seq=0 time=95.177ms
32 bytes from 110.242.68.3: icmp_seq=1 time=14.267ms
32 bytes from 110.242.68.3: icmp_seq=2 time=13.295ms
32 bytes from 110.242.68.3: icmp_seq=3 time=12.995ms
32 bytes from 110.242.68.3: icmp_seq=4 time=13.248ms

--- www.baidu.com ping statistics ---
# 丢包率
5 packets transmitted, 5 packets received, 0% packet loss

# 路由响应时常
round-trip min/avg/max/stddev = 12.995ms/29.7964ms/95.177ms/32.693173ms

```

### 编译命令

```

bash build/build.sh Builder   
```