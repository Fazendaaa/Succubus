FROM node:lts-alpine3.13

RUN [ "npm", "install", "-g", "npm" ]

USER node

WORKDIR /home/node/app

ENTRYPOINT [ "npx", "@appnest/readme", "generate" ]
