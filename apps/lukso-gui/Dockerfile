FROM node:16.13.0-alpine3.14 AS compile-image
WORKDIR /app
COPY package.json ./
RUN yarn install
ENV PATH="./node_modules/.bin:$PATH" 
COPY . ./
RUN nx build --prod

FROM nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY --from=compile-image /app/dist/ui /usr/share/nginx/html