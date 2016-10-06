package internet

// func insert_gzipped_sql() {
// 	fmt.Printf("[storage] Inserting gzipped sql files\n")

// 	config := utils.NewConfig()

// 	if run := config.Get("INSERT_HOSTNAMES"); run == "yes" {
//                 fmt.Printf("[storage] Starting to insert 1m hostnames in .sql file\n")
// 		rows, err := postgres.ExecuteGzippedSQL("data/raw-data/top-1m-hostnames.sql.gz")
// 		if err != nil {
// 			fmt.Printf("[storage] Error when inserting top 1m hostnames from .sql err=%s\n", err)
// 		}
// 		fmt.Printf("[storage] Inserted %d hostname rows.\n", rows)
// 	}
// }
