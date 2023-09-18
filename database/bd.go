package database

import (
	"cuadrangular/models"
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB inicializa la conexión a la base de datos MySQL
func InitDB() error {
	connectionString := "root:Artes2021*@tcp(127.0.0.1:3306)/conection"
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Conexión a la base de datos MySQL es exitosamente")

	return nil

}

// InsertNewTeam inserta un nuevo equipo en la base de datos
func InsertNewTeam(team *models.Equipos) error {
	var err error
	if db == nil {
		return errors.New("La conexión a la base de datos no está disponible")
	}

	query := "INSERT INTO Equipos (Nombre_equipo) VALUES (?)"

	_, err = db.Exec(query, team.Nombre_equipo)
	if err != nil {
		return err
	}

	return nil
}

func SelectTeams() ([]models.Equipos, error) {
	query := "SELECT * FROM Equipos"
	res, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var equipos []models.Equipos
	for res.Next() {
		var equipo models.Equipos
		err := res.Scan(&equipo.Id_equipo, &equipo.Nombre_equipo)
		if err != nil {
			return nil, err
		}
		equipos = append(equipos, equipo)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return equipos, nil
}
