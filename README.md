# proxyTelegramBot

Первое знакомство с golang.

Не бошьая утилита для пересылки сообщений полученных по http в канал или контакт в телеграм.

## Как использовать.
Собирите билд.

Переименуйте config.json.example в config.json

Заполните поля в config.json.

Token токен телеграм бота.

ChatId id канала или контакта.

HttpPort можно оставить по умолчанию

Запустите бинарный файл.

Перейдите в браузере по адресу http://localhost:HttpPort/msg?msg=test

Есле все сделанно правильно в телеграмм должно прийти сообщение от вашего бота Token.

Для чего это нужно.

Не накаждом оборудовании можно зпустить телеграм бота, а уведомления нужны.

Поэтому при помощи curl или другими стредсвами можно делать отправку сообщений через данного посредника.
