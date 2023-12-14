<h1 align="center">Inventory Management System</h1>

## About the Project

It is basic inventory management system utilising basic crud operations. It provides following features:
1. Allow to add product or category of product in database which utilize create operation
2. Allow to retrieve products or category of product and a product with the id which utilize the retrieve operation.
3. Allow to update the product or category 
4. Allow to delete the product or category

---
**_NOTE:_**
  * When you a delete a category then it will delete all the product of that category.
  * If you add product with new category it automatically add category to categories database.
---

## Getting Started

### Prerequisites

**_Golang:_**
* Download and Install Golang from https://go.dev/dl/ or if you using a mac os then use following command in homebrew environment.
```sh
brew install golang
```

**_Postgresql:_**
* Download and Install Postgresql from https://www.postgresql.org/download/ or if you using a mac os then use following command in homebrew environment.
```sh
brew install postgresql
```

### Set up Database
1. Open postgresql shell which looks like this
```sh
postgres=#
```
2. Now Create Database in postgres by following command
```sh
CREATE DATABSE inventory_management_system;
```
3. Select and use that database
```sh
\c inventory_management_system;
```
4. Now Create Product and Categories tables
```sh
CREATE TABLE categories (category_id INT GENERATED ALWAYS AS IDENTITY, name VARCHAR(255) NOT NULL, PRIMARY KEY(category_id)); 
```
```sh
CREATE TABLE product (Product_id INT GENERATED ALWAYS AS IDENTITY, name VARCHAR(255) NOT NULL, quantity INT NOT NULL, price INT NOT NULL, category_id INT NOT NULL, PRIMARY KEY(product_id),CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE);
```

All done. Now close you can close postgres shell.

### Set up Project
1. Clone the github repository on your local machine directory using command
```sh
git clone https://github.com/mayankr5/Inventory-management-system.git
```
2. Inside that directory run:
```sh
go mod init github.com/inventory-management-system
```
```sh
go mod tidy
```
3. Set up the _/config/.env_ file: <br>
Set **DB_USER** to _user_ and **DB_PASSWORD** to _password_ of that correspond user of postgres
```
HTTP_PORT=3000

DB_HOST=localhost
DB_USER={user}
DB_PASSWORD={password}
DB_NAME=inventory_management_system
DB_PORT=5432
DB_DIALECT=postgres
```
Now we are ready to run the project.

### Run Project
In the project directory run the command:
```sh
go run main.go
```

Now, Go server is running on **PORT**: 3000 (if you didn't change value of HTTP_PORT in .env file).

All the API's are in postman collections. Import `Inventory_management_system.postman_collection` file into your postman and run all test.





