FROM swaggerapi/swagger-ui

# Basic認証
RUN apk add apache2-utils

ARG USERNAME=username
ARG PASSWORD=password

RUN htpasswd -c -b /etc/nginx/.htpasswd ${USERNAME} ${PASSWORD} && cat /etc/nginx/.htpasswd

RUN sed -i "s|location / {|location / {\n\tauth_basic \"Restricted Content\";\n\tauth_basic_user_file /etc/nginx/.htpasswd;\n|g" /etc/nginx/nginx.conf
