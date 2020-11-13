package db

import (
	"database/sql"
	"log"
)

// Conecta com bando de dados
func ConectaComBancoDeDados() *sql.DB {
	conexao := "user=hpfcgdek dbname=hpfcgdek password=zRlFTcrycjQ9wJZKUKA_i25wgMJ7tz6o host=lallah.db.elephantsql.com"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}

	return db

}
