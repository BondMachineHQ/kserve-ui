FROM node:18-alpine3.15

COPY . /opt/app

WORKDIR /opt/app

RUN npm install

ENTRYPOINT ["node", "main.js"]