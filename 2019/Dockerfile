FROM php:7.4-cli
RUN pecl install xdebug-2.8.0 \
    && docker-php-ext-enable xdebug
WORKDIR /opt/project
CMD ["php"]
