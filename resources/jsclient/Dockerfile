FROM node:4.8-alpine

ADD jsclient/jsclient.js /opt/

WORKDIR /opt/

RUN npm install \
  ws \
  chance

ENTRYPOINT ["node", "jsclient.js"]
