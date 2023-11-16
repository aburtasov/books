## Books-app for create and store books and authors in Database

### Prerequisites:
* docker
* docker-compose
* make

### For run application:
``` sudo make build ```

 GRPC-server will start and also DB MYSQL will start with test data (see directory database)

#### As client-app you can use [evans](https://github.com/ktr0731/evans).

#### Run evans in project directory:
``` evans proto/books.proto -p 8000 ```

##### Available 4 methods for call:
* addBook
* addAuthor
* GetBooks
* GetAuthor



