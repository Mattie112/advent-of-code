# advent-of-code

2021 Is in Golang, please ignore all PHP related stuff, todo clean up this repo so it is more clear on what to use for what year!

My repo for AoC, this one replaces the following:

- https://github.com/Mattie112/advent-of-code-2017
- https://github.com/Mattie112/advent-of-code-2018
- https://github.com/Mattie112/advent-of-code-2019

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
```
