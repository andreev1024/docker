***Настройка дла Linux***

#Локальная разработка

`*repo*` - папка, куда клонировали репозиторий.

*   Добавляем в файл `/etc/hosts` новый хост (к примеру, `test`);
*   Клонируем репозиторий;
*   В файле `*repo*/docker/local/go/docker-compose.yml` правим пути, на хост-машине (к примеру, `~/www/golang/src/app:/go/src/app`, где `~/www/golang/src/app` путь к файлам на хост-машине), помеченные `#CHANGE IT`;
*   Nginx принимает запросы на `8080` порту и проксирует их на Go приложение, которое слушает на `8090` порту.  Изменить порты можно в `*repo*/docker/local/go/src/nginx/default`.
*  `cd *repo*/docker/local/go && docker-compose up` - запускаем приложение.
*  В браузере открываем `http://test:8080`.

##Миграции
Для миграций мы используем [этот модуль.](https://github.com/mattes/migrate)

*   убедитесь, что папка с миграциями подключена:
```
    golang:
      ...
      volumes:
        - ~/www/golang/migrations:/go/src/migrations   # CHANGE IT
        ...
      links:
        - mysql
```
*   В go-контейнере должен быть установлен модуль миграций. Можно сделать это в `init.sh`
```
    cd /go/src/app && go install -v && go get -u github.com/mattes/migrate && app
```
*   Заходим в контейнер `docker exec -it go_golang_1 bash`
*   Запускаем миграции `migrate -url "mysql://root:root@tcp(mysql:3306)/test" -path /go/src/migrations up`.
*   Имей ввиду, что команда `migrate` будет доступна не сразу после старта контейнера, т.к контейнеру нужно время на установку нужного пакета.