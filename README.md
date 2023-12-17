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
---

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
<!-- 
### Set up Database
---
1. Create database by following command. (_Change user_name in following command with postgres user_name_)
```sh
psql template1 -c "CREATE DATABASE inventory_management_system WITH OWNER user_name;";
```
2. Now restore the dumpfile using following command. (_Make sure you are in current working directory_)
```sh
psql inventory_management_system < inventory_management_system
```
Now Database is all setup. -->

### Set up Project
---
1. Clone the github repository on your local machine directory using command
```sh
git clone https://github.com/mayankr5/Inventory-management-system.git
```
2. Set up the _/config/.env_ file: <br>
Set **DB_USER** to _user_ and **DB_PASSWORD** to _password_ of that correspond user of postgres
```
HTTP_PORT=3000

DB_HOST=localhost
DB_USER={USER_NAME}
DB_PASSWORD={USER_PASSWORD}
DB_NAME=inventory_management_system
DB_PORT=5432
DB_DIALECT=postgres
```
3. Now we have to setup Database
- Run the following command in terminal to create a database. (Make sure you are currently in this directory)
```sh
source configs/.env && psql template1 -c "CREATE DATABASE $DB_NAME WITH OWNER $DB_USER;";
```
- Now dump schema in that database
```sh
source configs/.env && pg_restore -U $DB_USER -h $DB_HOST -d $DB_NAME db.dump
```

### Run Project

---
In the project directory run the command:
```sh
go run main.go
```

Now, Go server is running on **PORT**: 3000 (if you didn't change value of HTTP_PORT in .env file).

All the API's are in postman collections. Import `Inventory_management_system.postman_collection` file into your postman and run all test.

---

