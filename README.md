# SferumChatMembers

Утилита для извлечения и форматирования данных об образовательных профилях участников группы в Sferum. Работает с cookies браузера и выдаёт результаты в виде таблицы и CSV-файла.

---

## 🛠 Установка

### Установите golang:
https://go.dev/doc/install


### Клонируйте репозиторий:
```bash
git clone https://github.com/igro83/sferumChatMembers.git
```
### Смените текущую директорию
```bash
cd sferumChatMembers
```
### Установите зависимости:
```bash
go mod download
```
### Заполните файл .env.example с вашими данными.
Заполните в файле .env.example строку REMIXDSID= данными из браузера - cookies remixdsid с сайта мессенджера web.vk.me
Заполните в файле .env.example строку CHAT= данными из адресной строки браузера, при переходе на сайте web.vk.me в нужный чат

### Переименуйте .env.example в .env

### Запустите сборку приложения
```bash
go build
```

### Запустите приложение
```bash
./sferumMembers
```
