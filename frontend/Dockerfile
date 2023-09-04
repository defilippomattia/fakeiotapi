# Stage 1: Build the React app
FROM node:20-alpine3.17 

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

RUN npm run build

# Stage 2: Serve the React app from Nginx
FROM nginx:1.21.1-alpine

COPY --from=0 /app/build /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

COPY ./fullchain.pem /etc/nginx/fullchain.pem
COPY ./privkey.pem /etc/nginx/privkey.pem

EXPOSE 80
EXPOSE 443

CMD ["nginx", "-g", "daemon off;"]
