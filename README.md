# Golang-CRUD-RESTful-API-with-PostgreSQL
## Initialization
go mod name : gotest

## Package :
gorilla/mux : go get -u github.com/gorilla/mux  
lib/pq : go get github.com/lib/pq  
joho/godotenv : go get github.com/joho/godotenv

## Database
database name : gotest  
owner : postgres

Create Table Query :
```CREATE TABLE car
(
 car_id serial NOT NULL,
 car_name character varying NOT NULL,
 car_brand character varying NOT NULL,
 car_type character varying,
 car_year integer,
 car_desription character varying,
 CONSTRAINT pk_buku PRIMARY KEY (id )
)
WITH (
 OIDS=FALSE
);
ALTER TABLE car
 OWNER TO postgres;
 ```
 
## JSON Body
JSON Body for AddCar :
```
{
	"car_id":1,
	"car_name":"Murcielago",
    	"car_brand":"Lamborghini",
    	"car_type":"Exotic",
	"car_year":2005,
	"car_description":"Super Car"
}
```
JSON Body for Update Car By ID :
```
{
	"car_name":"Murcielago",
    	"car_brand":"Lamborghini",
    	"car_type":"Exotic",
	"car_year":2005,
	"car_description":"Super Car"
}
```
JSON Body for Update Car By Car Name :
```
{
	"car_id":1,
    	"car_brand":"Lamborghini",
    	"car_type":"Exotic",
	"car_year":2005,
	"car_description":"Super Car"
}
```
## API Handling and Features
1. http://127.0.0.1:1111/api/car 		//GET Method => Get all car data
2. http://127.0.0.1:1111/api/car 		//POST Method with JSON Body => Add car
3. http://127.0.0.1:1111/api/car/{id}		//GET Method => Get car data by id
4. http://127.0.0.1:1111/api/car/{id}		//PUT Method with JSON Body => Update existed car data by id
5. http://127.0.0.1:1111/api/car/{id}		//DELETE Method => Delete car data by id
6. http://127.0.0.1:1111/api/carbrand/{brand}	//GET Method => Get car data by brand name
7. http://127.0.0.1:1111/api/carbrand/{brand}	//DELETE Method => Delete car data by brand name
8. http://127.0.0.1:1111/api/carname/{name}	//GET Method => Get car data by car name
9. http://127.0.0.1:1111/api/carname/{name}	//PUT Method with JSON Body => Update car data by car name
10. http://127.0.0.1:1111/api/carname/{name}	//DELETE Method => Delete car data by car name
