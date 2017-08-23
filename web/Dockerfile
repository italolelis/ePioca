FROM node:8.4-alpine

RUN mkdir -p /server/log

ARG APP_DIR='/server/http'

ADD . $APP_DIR

RUN npm install

RUN npm install --global nodemon

EXPOSE 8080

CMD ["npm", "start"]

WORKDIR $APP_DIR
