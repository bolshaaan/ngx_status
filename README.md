# ngx_status


just make

 - go get  -u github.com/bolshaaan/ngx_status
 - go install github.com/bolshaaan/ngx_status

and use it in your path $GOPATH/bin/ngx_status

##

this is nginx status format for telegraf to put stat to influxdb


vim /etc/telegraf/telegraf.conf


    - [[inputs.exec]]
    - commands = ["/home/bolshakov/go/bin/ngx_status"]