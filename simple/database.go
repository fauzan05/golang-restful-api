package simple

type Database struct {
	Name string
}

type DatabaseMySQL Database
type DatabaseMongoDB Database

func NewDatabaseMySQL() *DatabaseMySQL {
	database := &Database{
		Name: "Database MySQL",
	}
	return (*DatabaseMySQL)(database)
}

func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{
		Name: "Database MongoDB",
	}
	return (*DatabaseMongoDB)(database)
}

type DatabaseRepository struct {
	DatabaseMySQL   *DatabaseMySQL
	DatabaseMongoDB *DatabaseMongoDB
}

func NewDatabaseRepository(mysql *DatabaseMySQL, mongo *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMySQL: mysql,
		DatabaseMongoDB: mongo,
	}
}
