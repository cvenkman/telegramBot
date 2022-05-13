
![bot](https://github.com/cvenkman/telegramBot/raw/master/bot.png)



структура проекта

* /cmd/botServer - точка входа
* /configs - Шаблоны файлов конфигураций и файлы настроек по-умолчанию.
* /pkg/storage - база данных
    - интерфейс
    - /boltDB - реализация
* /internal
    - /telegram
        * /telegram - работа с telegram-bot-api и handler функции
        * /models - json модели для взаимодействия с telegram api
    - /youtube
        * /youtube - работа с youtube api
        * /models - json модели для взаимодействия с youtube api