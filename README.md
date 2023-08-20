## Books-app для создания и хранения книг и авторов в БД

### Prerequisites:
* docker
* docker-compose
* make

### Для запуска приложения:
``` sudo make build ```

Запустится GRPC-сервер и БД MYSQL с тестовыми данными (смотреть директорию database)

#### В качестве клиента можно использовать [evans](https://github.com/ktr0731/evans).

#### Запустите evans в директории проекта:
``` evans proto/books.proto -p 8000 ```

##### Доступно 4 метода для вызова:
* addBook
* addAuthor
* GetBooks
* GetAuthor



