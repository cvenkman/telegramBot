Бот может присылать фото кошек собак и лис, работать с базой данных, через YouTube api присылать последнее видео каналов.

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
