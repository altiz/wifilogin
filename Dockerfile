FROM golang:latest
RUN mkdir /app 
WORKDIR /app

COPY . .



RUN apt-get update && apt-get install -y libaio1 wget unzip
RUN go get github.com/githubnemo/CompileDaemon
RUN go get gopkg.in/goracle.v2
RUN go get github.com/nixys/nxs-go-zabbix/v5

WORKDIR /opt/oracle

RUN wget https://download.oracle.com/otn_software/linux/instantclient/instantclient-basiclite-linuxx64.zip && \
    unzip instantclient-basiclite-linuxx64.zip && \
    rm -f instantclient-basiclite-linuxx64.zip && \
    cd instantclient* && \
    rm -f *jdbc* *occi* *mysql* *jar uidrvci genezi adrci && \
    echo /opt/oracle/instantclient* > /etc/ld.so.conf.d/oracle-instantclient.conf && \
    ldconfig

WORKDIR /app/cmd/wifilogin/

#RUN go build -o main .


ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main

#ENTRYPOINT /app/cmd/docktest/main
