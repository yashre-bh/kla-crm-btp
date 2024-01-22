# mvc-lecture
- Setup Golang on your system. Clone the repo. From the root directory:
```
- go mod vendor
- go mod tidy
```


- MYSQL:
1. `mysql -u root -p` : and enter password
2. Create a new database 'books': `CREATE DATABASE kla_crm;`
3. Connect to the database: `USE kla_crm;;`
4. Import the sql dump file into your database: `mysql -u root -p kla-crm < dump.sql`


- Running the server:
1. `go build -o mvc ./cmd/main.go`
2.  Run the binary file: `./mvc`

