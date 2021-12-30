FROM node:lts

RUN mkdir -p /buyme/front
WORKDIR /buyme/front

ENV PATH /buyme/front/node_modules/.bin:$PATH

COPY package*.json /buyme/front
RUN npm install

COPY . /buyme/front

CMD ["npm", "run", "dev"]