package project

var (
	Hostname = ""
	Port     = 2345
	Username = ""
	Password = ""
	Database = ""
)

func AddZipcode(d Zipcode_CCVI) int {

	db, err := openConnection()
	if err != nil {
		fmt.Println(err)
		return -1
	}
	defer db.Close()

	insertStatement := `insert into "zipcode_ccvi" ("zipcode", "category") values ($1, $2)`
	_, err = db.Exec(insertStatement, d.Community_area_or_zip, d.Ccvi_category)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	i, err := strconv.Atoi(d.Community_area_or_zip)
    if err != nil {
        // ... handle error
        panic(err)
    }

	return i
}

func openConnection() (*sql.DB, error) {
   	Hostname = "localhost"
	Port = 5433
	Username = "postgres"
	Password = "root"
	Database = "project"

	// connection string
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		Hostname, Port, Username, Password, Database)

	// open database
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
