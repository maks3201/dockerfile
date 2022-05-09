FROM alpine:3.15.4

ENV NGINX_VERSION "1.20.2-r1"

COPY index.html /opt/

RUN apk add --no-cache --repository https://dl-cdn.alpinelinux.org/alpine/v3.9/main/ \
	      nginx=${NGINX_VERSION} && apk --no-cache add curl 

EXPOSE 80

COPY default.conf /etc/nginx/http.d/

COPY nginx.conf /etc/nginx/

ENTRYPOINT echo "hostname: `hostname`" >> /opt/index.html && nginx -g "daemon off;"
