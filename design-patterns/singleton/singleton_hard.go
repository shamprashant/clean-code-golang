package main

import (
	"fmt"
	"sync"
	"time"
)

// DatabaseConnection represents a singleton database connection
type DatabaseConnection struct {
	connectionString string
	isConnected      bool
	mu               sync.RWMutex
}

var (
	dbInstance *DatabaseConnection
	Once       sync.Once
)

// GetDatabaseConnection returns the singleton instance
func GetDatabaseConnection() *DatabaseConnection {
	Once.Do(func() {
		fmt.Println("Creating database connection...")
		dbInstance = &DatabaseConnection{
			connectionString: "postgres://localhost:5432/mydb",
			isConnected:      false,
		}
		// Simulate connection setup
		time.Sleep(100 * time.Millisecond)
		dbInstance.isConnected = true
		fmt.Println("Database connection established")
	})
	return dbInstance
}

// Connect simulates connecting to database
func (db *DatabaseConnection) Connect() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if db.isConnected {
		return nil
	}

	fmt.Println("Connecting to database...")
	time.Sleep(50 * time.Millisecond)
	db.isConnected = true
	return nil
}

// Query simulates a database query
func (db *DatabaseConnection) Query(sql string) ([]string, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	if !db.isConnected {
		return nil, fmt.Errorf("database not connected")
	}

	// Simulate query execution
	return []string{"result1", "result2"}, nil
}

// Close simulates closing the connection
func (db *DatabaseConnection) Close() error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if !db.isConnected {
		return nil
	}

	fmt.Println("Closing database connection...")
	db.isConnected = false
	return nil
}

// Example usage
// func main() {
// 	// Simulate concurrent access
// 	var wg sync.WaitGroup

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()

// 			db := GetDatabaseConnection()
// 			fmt.Printf("Goroutine %d got connection: %p\n", id, db)

// 			results, err := db.Query("SELECT * FROM users")
// 			if err != nil {
// 				fmt.Printf("Goroutine %d error: %v\n", id, err)
// 				return
// 			}

// 			fmt.Printf("Goroutine %d results: %v\n", id, results)
// 		}(i)
// 	}

// 	wg.Wait()

// 	// Clean up
// 	GetDatabaseConnection().Close()
// }
