#!/bin/sh

echo 'websocket server starting...'
./WebsocketServer --daemon=true --service_addr=192.168.88.206 --spring.cloud.config.name=ws --spring.cloud.config.profile=rfbak --spring.cloud.config.label=master --spring.cloud.config.uri=http://api.zhangling.link:8084

