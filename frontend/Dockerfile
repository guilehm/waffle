FROM node:20.11-alpine

WORKDIR /usr/app

COPY ./package.json ./
COPY ./yarn.lock ./

RUN yarn install

COPY . .
