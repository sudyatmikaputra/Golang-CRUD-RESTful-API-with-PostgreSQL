package models

import (
	"database/sql"
	"fmt"
	"gotest/config"
	"log"
	_ "github.com/lib/pq"
)

type Car struct {
	Car_id				int64  	`json:"car_id"`
	Car_name			string 	`json:"car_name"`
	Car_brand			string 	`json:"car_brand"`
	Car_type			string 	`json:"car_type"`
	Car_year			int64	`json:"car_year"`
	Car_description		string	`json:"car_description"`	
}

// CRUD BY ID

func AddCar(car Car) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO car (car_id, car_name, car_brand, car_type, car_year, car_description) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(sqlStatement, car.Car_id, car.Car_name, car.Car_brand, car.Car_type, car.Car_year, car.Car_description)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	fmt.Printf("Insert data single record %v", car.Car_id)
	return car.Car_id
}

func GetAllCars() ([]Car, error) {
	db := config.CreateConnection()
	defer db.Close()

	var cars []Car

	sqlStatement := `SELECT * FROM car`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var car Car

		err = rows.Scan(&car.Car_id, &car.Car_name, &car.Car_brand, &car.Car_type, &car.Car_year, &car.Car_description)
		if err != nil {
			log.Fatalf("Fail to get data. %v", err)
		}
		cars = append(cars, car)

	}
	return cars, err
}

func GetCar(id int64) (Car, error) {
	db := config.CreateConnection()
	defer db.Close()

	var car Car

	sqlStatement := `SELECT * FROM car WHERE car_id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&car.Car_id, &car.Car_name, &car.Car_brand, &car.Car_type, &car.Car_year, &car.Car_description)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No data found!")
		return car, nil
	case nil:
		return car, nil
	default:
		log.Fatalf("Fail to get data. %v", err)
	}
	return car, err
}

func UpdateCar(id int64, car Car) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE car SET car_name=$2, car_brand=$3, car_type=$4, car_year=$5, car_description=$6 WHERE car_id=$1`
	res, err := db.Exec(sqlStatement, id, car.Car_name, car.Car_brand, car.Car_type, car.Car_year, car.Car_description)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating data. %v", err)
	}
	fmt.Printf("Total rows/record updated is %v\n", rowsAffected)

	return rowsAffected
}

func DeleteCar(id int64) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM car WHERE car_id=$1`
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Cannot find the data. %v", err)
	}
	fmt.Printf("Total deleted data is %v", rowsAffected)

	return rowsAffected
}

//-------------------------------------------------------------------------------------------------------------//

// CRUD BY BRAND

func GetCarByBrand(brand string) ([]Car, error) {
	db := config.CreateConnection()
	defer db.Close()

	var cars []Car

	sqlStatement := `SELECT * FROM car WHERE car_brand=$1`
	rows, err := db.Query(sqlStatement, brand)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var car Car

		err := rows.Scan(&car.Car_id, &car.Car_name, &car.Car_brand, &car.Car_type, &car.Car_year, &car.Car_description)
		if err != nil {
			log.Fatalf("Fail to get data. %v", err)
		}
		cars = append(cars, car)

	}
	return cars, err
}

func DeleteCarByBrand(brand string) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM car WHERE car_brand=$1`
	res, err := db.Exec(sqlStatement, brand)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Cannot find the data. %v", err)
	}
	fmt.Printf("Total deleted data is %v", rowsAffected)

	return rowsAffected
}


//-------------------------------------------------------------------------------------------------------------//

// CRUD BY NAME

func GetCarByName(carname string) (Car, error) {
	db := config.CreateConnection()
	defer db.Close()

	var car Car

	sqlStatement := `SELECT * FROM car WHERE car_name=$1`
	row := db.QueryRow(sqlStatement, carname)
	err := row.Scan(&car.Car_id, &car.Car_name, &car.Car_brand, &car.Car_type, &car.Car_year, &car.Car_description)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No data found!")
		return car, nil
	case nil:
		return car, nil
	default:
		log.Fatalf("Fail to get data. %v", err)
	}

	return car, err
}

func UpdateCarByName(carname string, car Car) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `UPDATE car SET car_id=$2, car_brand=$3, car_type=$4, car_year=$5, car_description=$6 WHERE car_name=$1`
	res, err := db.Exec(sqlStatement, carname, car.Car_id, car.Car_brand, car.Car_type, car.Car_year, car.Car_description)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while updating data. %v", err)
	}
	fmt.Printf("Total rows/record updated is %v\n", rowsAffected)

	return rowsAffected
}

func DeleteCarByName(carname string) int64 {
	db := config.CreateConnection()
	defer db.Close()

	sqlStatement := `DELETE FROM car WHERE car_name=$1`
	res, err := db.Exec(sqlStatement, carname)

	if err != nil {
		log.Fatalf("Fail to execute query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Cannot find the data. %v", err)
	}
	fmt.Printf("Total deleted data is %v", rowsAffected)

	return rowsAffected
}