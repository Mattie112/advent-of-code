All new years will be added to this repo and can be run by using the `cli.php` file

Windows users need to add a .env file with:
```
XDEBUG_CONFIG="client_host=192.168.100.10"
PHP_IDE_CONFIG="serverName=AOC"
```

Linux users can use (docker internal ip for the host):
```
XDEBUG_CONFIG="client_host=172.17.0.1"
PHP_IDE_CONFIG="serverName=AOC"