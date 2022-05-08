FROM alpine:3.15.4

ENV NGINX_VERSION "1.20.2-r1"

RUN apk add --no-cache --repository https://dl-cdn.alpinelinux.org/alpine/v3.9/main/ \
	      nginx=${NGINX_VERSION}

EXPOSE 80

COPY default.conf /etc/nginx/http.d/

COPY nginx.conf /etc/nginx/

COPY index.html /opt/

RUN echo "<p> hostname: `hostname`" >> /opt/index.html

CMD ["nginx","-g","daemon off;"]
