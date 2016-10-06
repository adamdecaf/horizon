package geo

// func insert_raw_data() {
// 	fmt.Println("[data] Starting raw data insert")

// 	// number of workers, and size of job queue
// 	pool := grpool.NewPool(10, 10 * 1000 * 1000)

// 	defer pool.Release()

// 	if pool == nil {
// 		fmt.Println("[storage] worker pool is nil")
// 	} else {
// 		pool.WaitCount(10)

// 		config := utils.NewConfig()

// 		if run := config.Get("INSERT_RAW_STATES"); run == "yes" {
// 			if err := geo.InsertRawStates(*pool); err != nil {
// 				fmt.Printf("[data] Error when inserting raw state data (err=%s)\n", err)
// 			}
// 		}

// 		if run := config.Get("INSERT_RAW_CITIES"); run == "yes" {
// 			if err := geo.InsertRawCitiesFromStates(*pool); err != nil {
// 				fmt.Printf("[data] Error when inserting raw city data (err=%s)\n", err)
// 			}
// 		}

// 		if run := config.Get("INSERT_RAW_COUNTRIES"); run == "yes" {
// 			if err := geo.InsertCountries(*pool); err != nil {
// 				fmt.Printf("[data] Error when inserting country data (err=%s)\n", err)
// 			}
// 		}

// 		pool.WaitAll()
// 	}
// }
