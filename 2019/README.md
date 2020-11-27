# advent-of-code-2019
My solutions for adventofcode.com 2019

### How to run
```
docker-compose run php-cli dayX/dayX.php
```

### How to debug on Windows
For debugging on Windows create a `.env` file in your root directory and add the following:
```
XDEBUG_USER_CONFIG=-d xdebug.remote_host=<the ip of your machine>
```
For Linux this step is not needed, it should work out of the box
