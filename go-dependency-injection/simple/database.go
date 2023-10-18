package simple

type Database struct {
	Name string
}

type DatabaseRepository struct {
	DatabasePostgreSQL *Database
	DatabaseMongoDB    *Database
}

func NewDatabasePostgreSQL() *Database {
	return &Database{Name: "PostgreSQL"}

}

func NewDatabaseMongoDB() *Database {
	return &Database{Name: "MongoDB"}

}

func NewDatabaseRepository(postgreSQL *Database, mongodb *Database) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongodb}

}
