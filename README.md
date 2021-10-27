# OnlineVideoPlayer
### 需要根据自己的情况改的点有以下几个：
1. 如果需要对外暴露服务且自己的ip是私有ip，则需要通过nat端口映射或者直接把服务部署到公有ip上去；如果自己的本身就是公有ip则不需要
2. 视频的处理方式，就需要把sample-mp4-file.mp4替换为自己录制的mp4文件，然后按照
```
ffmpeg -i 你的视频名字.mp4 -profile:v baseline -level 3.0 -s 640x360 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls -strict -2 index.m3u8
并将的到的文件 全部剪切到videos/video1目录下面去
```
