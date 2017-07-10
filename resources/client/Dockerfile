FROM node:4.8-alpine

ADD client/client.js /opt/

WORKDIR /opt/

RUN npm install \
  ws \
  chance

ENTRYPOINT ["node", "client.js"]
