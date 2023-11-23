FROM php:8.1-fpm-alpine

RUN docker-php-ext-install \
    pdo \
    pdo_mysql

RUN addgroup -g 1000 phpuser \
    && adduser -G phpuser -g phpuser -s /bin/sh -D phpuser

USER phpuser 

WORKDIR /var/www/html

COPY src/php . 

