FROM alpine:3.11
LABEL maintainer="drmfly.liw@gmail.com"
RUN echo -e  "http://mirrors.aliyun.com/alpine/v3.11/main\nhttp://mirrors.aliyun.com/alpine/v3.11/community" >  /etc/apk/repositories \
&& apk update && apk add tzdata \
&& cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
&& echo "Shanghai/Asia" > /etc/timezone \
&& apk del tzdata
WORKDIR /app
COPY ["gopkg","./configs/app.example.yml", "./"]
ENTRYPOINT ["./gopkg"]

