FROM php:8.0-cli
COPY --from=mlocati/php-extension-installer /usr/bin/install-php-extensions /usr/local/bin/
RUN install-php-extensions xdebug
COPY config/xdebug.ini $PHP_INI_DIR/conf.d/xdebug.ini
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer
ENV PATH /root/.composer/vendor/bin:$PATH
