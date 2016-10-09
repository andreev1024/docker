#Docker containers

**Note**
* If you use Linux add `sudo` to all Docker commands.

##Local installation

###Prepare

* Install Docker. Read more details here: https://docs.docker.com

* Install Docker Compose (If you use Linux install v.1.3.0 because v.1.4.1 has a bug);

* Our application will be accessible at the `rereca`, `rereca-backend`, `rereca-api` domains. You should map this domains to your Docker host IP address. If you run Docker natively on your Linux operating system, use the IP address of your own computer. If you rely on Boot2Docker, find your Docker host IP address with the boot2docker ip bash command. Let's assume your Docker host IP address is 192.168.0.100. You can map out domains name to the 192.168.0.100 IP address by appending these lines to your local computer's `/etc/hosts` file:

```
    192.168.0.100 rereca   
    192.168.0.100 rereca-backend   
    192.168.0.100 rereca-api   
```
 
* Go into to `rebecca-devops/docker/local` directory. Lets see what we've got:

```
    - src                       //  dir with configuration files
    - www                       //  there we must place our app
    - docker-compose.yml        //  docker-compose file for PHP app
```

* Copy directory with your project (for e.g `rereca`) in `www` . So, now path to your project will be `www/rereca`.

###PHP

*  Open `docker-compose.yml` and configure it:

```
php:
  ...
  environment:
      MY_NAME: "local-user"
      MY_ID: 1000
      HOST_IP: "192.168.0.101"
      HOST_NAMES: "rereca rereca-backend rereca-api"
      AFTER_SCRIPT_CMD: "/sbin/my_init -- php5-fpm --nodaemonize"
  ...
```

`HOST_IP` - there IP from section above (in our exaple it's 192.168.0.100).

`HOST_NAMES` - there host names separated by a space from section above.

`MY_ID` - there must be your UID on host machine.

* Next, we must configure `composer`. Go into `src/php/home/local-user/.composer`. This directory will be mount in container. We must place in it `auth.json`. This file contains token which give you access to repo.

* Next, we must configure `ssh` because `composer` uses it. Add in `src/php/home/local-user/.ssh` keys `id_rsa` and `id_rsa.pub`.

###MYSQL

* Open `docker-compose.yml` and search `mysql` configuration block. There we have:

```
    environment:
        - DB_NAME=_rebecca  // create new database with name '_rebecca'
        - DB_USER=root      // create new user with name 'root'   
        - DB_PASS=root      // ...and password 'root'
    ports:
        - "33060:3306"      // from your localhost you can connect to container
                            // database by `localhost:33060`
```

###Application configuration 

If you use docker with Yii2, please, read `yii2-configure.md`

##Run

* Go into directory where is `docker-compose.yml` and execute:

```
    docker-compose up
```

* Connect to mysql container by `localhost:33060` and import your database (I recommended use Mysql Workbench but you can use your favorite tools).

* Now, if your web-server port is `8080` 

```
    nginx:
      image: andreev/nginx
      ports:
        + "8080:80"
```

you can get your application in browser by `http://rereca:8080`

##You must know

* In your app you must reference on your services by `service name` which
specified in `docker-compose.yml` (mysql, redis, gearman etc.)

```
    'db' => [       
        'dsn' => 'mysql:host=mysql;dbname=_rebecca;', //  service named `mysql`
        'username' => 'root',
        'password' => 'root',
    ],
    'redis' => [       
        'class' => 'yii\redis\Connection',
        'hostname' => 'redis',   //  service named `redis`
        'port' => 6379,
        'database' => 0,
    ],
```

* You always can dig into container:

You can go into PHP continer like a `local-user`. `local-user` has same UID/GID as you (see PHP section). So, all files which `local-user` will create in container will be belong to you on host machine.

```
    docker exec -u local-user -it [container-id] bash     //  like local-user
```

Also you can dig like a root:

```
    docker exec -it [container-id] bash     //  like root
```

##Node

All *Node* application managed by [Strong](http://strong-pm.io/). We have script `docker/src/node/strong.sh` which adds our node applications to `Strong`. This script will executed when container started. Also `docker/src/node/strong.sh` must be executable.

If you want to get *Strong* logs you must log in *Node* (see FAQ) and use Strong log command (see Strong documentation).
```
for example u can use 
$ slc ctl log-dump express-example-app --follow
```


##FAQ

* *I've got an error when I run migrations.*

    You must run migrations from the container which contain `php`.

* *How to restart Docker containers?*

    Go into docker directory and run `docker-compose stop; docker-compose up -d`.

* *On runtime I've got message:*

    `Warning: World-writable config file '/etc/mysql/debian.cnf' is ignored`

    Open `docker/src/mysql/etc/mysql` and execute `chmod 600 debian.cnf`

* *What about logging ?*

    By default many containers logging in `stdout` and `stderr`. You can view logs by

```
        docker logs -f container_name
```

[More...](https://docs.docker.com/userguide/usingdocker/)

Node container use *Strong* (see Node section).


MAINTAINERS
-----------
 * Andreev <andreev1024@gmail.com>
 * Domi Besedin <dmitriy@reseed-s.com>
