# customerAPI
**customerAPI** is a RESTful API which handles customer data. Currently it can:
* Receive data for a new customer
* Return all customers' data
* Return data of a single customer

### Features
The API does not utilize a persistent data storage, but one can be easily implemented, it simply has to realize the `Database` interface.

POST data is validated with simple filters.

Basic Auth is implemented (currently hardcoded).

### Usage

Clone the repository and compile the application:
```
make install
```
Run the application: `./customer_api`

It is now accepting GET and POST requests on `/v1/users`. For simple API usage examples, refer to [tests](app_test.go).

**Optional:** export `CUSTOMER_API_PORT` to specify the port of the endpoint. Otherwise the default value `8080` will be used.

### Future Work
* Implement RESTful response messages for successful operations and errors (currently only posting proper response codes)
* Add the ability to update and delete customer data
* Add persistent data store