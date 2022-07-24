FROM nginx:1.23.1-alpine

COPY ./cert.pem /etc/nginx/ssl/cert.pem
COPY ./key.key /etc/nginx/ssl/key.key
COPY ./go_app.conf /etc/nginx/conf.d/default.conf

ENTRYPOINT ["nginx", "-g", "daemon off;"]