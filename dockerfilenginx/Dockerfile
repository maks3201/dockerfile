FROM alpine:3.15.4

ENV NGINX_VERSION "1.20.2-r0"

RUN apk add --no-cache --repository https://dl-cdn.alpinelinux.org/alpine/v3.9/main/ \
	      nginx=${NGINX_VERSION}

EXPOSE 80

COPY default.conf /etc/nginx/http.d/

COPY nginx.conf /etc/nginx/

COPY index.html /opt/

CMD ["nginx","-g","daemon off;"]
