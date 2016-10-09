There I explain how you can configure your Yii2 application to use with Docker containers. It is assumed that you configure Docker and its containers.

Database and storage
--------------------

In application you must reference on your services by `service name` which
specified in `docker-compose.yml` (mysql, redis, gearman etc.). Open your DB config and change it like below:

```
    'db' => [
        'dsn' => 'mysql:host=mysql;dbname=_rebecca;',  //  service named `mysql`  
        'username' => 'root',   //  your mysql username 
        'password' => 'root',   //  your mysql password
    ],
    'redis' => [
        'class' => 'yii\redis\Connection',
        'hostname' => 'redis',  //  service named `redis`
        'port' => 6379,
        'database' => 0,
    ],
```

Hosts and urls
--------------

Add in your `common/config/params-local.php` code from section below and configure `endpoint` values:

```
//  common/config/params-local.php

$values = [
    'host.api' => 'http://rereca-api:8080', //  your local API endpoint
];

return [
    ...
    'api.service.getUrl' => $values['host.api'] . '/v1/service_get-url',

    'host.backend' => 'http://rereca-backend:8080',     //  your local Backend endpoint
    'host.api' => $values['host.api'],
    ...
];
```







