package store

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) *MySQLStorage {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &MySQLStorage{db: db}
}

func (m *MySQLStorage) Init() (*sql.DB, error) {
	err := m.CreateUsersTable()
	if err != nil {
		return nil, err
	}

	err = m.CreateProjectsTable()
	if err != nil {
		return nil, err
	}

	err = m.CreateTasksTable()
	if err != nil {
		return nil, err
	}

	return m.db, nil
}

func (m *MySQLStorage) Close() {
	err := m.db.Close()
	if err != nil {
		log.Printf("Error closing the database connection: %v", err)
	}
}

func (m *MySQLStorage) CreateUsersTable() error {
	_, err := m.db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			email VARCHAR(100) NOT NULL,
		    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		    
		    PRIMARY KEY (id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	return err
}

func (m *MySQLStorage) CreateProjectsTable() error {
	_, err := m.db.Exec(`
		CREATE TABLE IF NOT EXISTS projects (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			userId INT UNSIGNED NOT NULL,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			
			PRIMARY KEY (id),
			FOREIGN KEY (userId) REFERENCES users(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	return err
}

func (m *MySQLStorage) CreateTasksTable() error {
	_, err := m.db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			status ENUM('todo', 'pending', 'completed') DEFAULT 'todo',
			projectId INT UNSIGNED NOT NULL,
			assignedToId INT UNSIGNED NOT NULL,
			description TEXT,
			createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			userId INT,
			
			PRIMARY KEY (id),
			FOREIGN KEY (projectId) REFERENCES projects(id),
			FOREIGN KEY (assignedToId) REFERENCES users(id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`)
	return err
}
