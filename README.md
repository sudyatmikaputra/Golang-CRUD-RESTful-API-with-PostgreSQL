# Golang-CRUD-RESTful-API-with-PostgreSQL
## Initialization
go mod name : gotest
<hr style="border:1px solid gray">

## Package :
gorilla/mux : go get -u github.com/gorilla/mux  
lib/pq : go get github.com/lib/pq  
joho/godotenv : go get github.com/joho/godotenv
<hr style="border:1px solid gray">

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
 <hr style="border:1px solid gray">

## JSON Body
JSON Body for AddCar :
```{
	"car_id":1,
	"car_name":"Murcielago",
    	"car_brand":"Lamborghini",
    	"car_type":"Exotic",
	"car_year":2005,
	"car_description":"Super Car"
}
```
JSON Body for Update Car By ID :
```{
	"car_name":"Murcielago",
    	"car_brand":"Lamborghini",
    	"car_type":"Exotic",
	"car_year":2005,
	"car_description":"Super Car"
}
```
JSON Body for Update Car By Car Name :
```{
	"car_id":1,
    	"car_brand":"Lamborghini",
    	"car_type":"Exotic",
	"car_year":2005,
	"car_description":"Super Car"
}
```
