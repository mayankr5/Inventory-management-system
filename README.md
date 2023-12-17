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

**_Docker:_**
* Download and Install Docker from https://www.docker.com/products/docker-desktop/ 

### Set up Project
---
1. Clone the github repository on your local machine directory using command
```sh
git clone https://github.com/mayankr5/Inventory-management-system.git
```

### Run Project

---
If you are running the programming first time then run following command.
```sh
docker-compose up
```
Otherwise run following command.
```sh
docker-compose up --build
```

Now, Go server is running on **PORT**: 3000 (if you didn't change value of HTTP_PORT in .env file).

All the API's are in postman collections. Import `Inventory_management_system.postman_collection.json` file into your postman and run all test.

---

