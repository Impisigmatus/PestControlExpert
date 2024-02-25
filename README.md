# PestControlExpert
## Состав репозитория
[Пользовательский интерфейс](www/README.md)

Микросервисы:
* [notification](microservices/notification/README.md)
* [prices](microservices/prices/README.md)

## Окружение
Для настройки окружения используется утилита `task`
```
$ go install github.com/go-task/task/v3/cmd/task@latest
```
Основные команд
```
$ task default # Список команд
$ task update  # Обновление сервисов
$ task cicd    # Локальный CI/CD сервисов
```
> Для просмотра полного списка команд:`$ task --list`
