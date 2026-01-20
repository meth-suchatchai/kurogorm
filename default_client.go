package kurogorm

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type defaultClient struct {
	orm  *gorm.DB
	name string
}

type Config struct {
	Name     string
	SSLMode  bool
	Host     string
	Username string
	Password string
	Port     int
	Timezone string
	Debug    bool
}

// ConnectSQL connect orm database
func ConnectSQL(cfg *Config) (Client, error) {
	name := cfg.Name
	ssl := "disable"
	if cfg.SSLMode {
		ssl = "enable"
	}

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v timezone=%v",
		cfg.Host, cfg.Username, cfg.Password,
		cfg.Name, cfg.Port, ssl, cfg.Timezone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &defaultClient{
		orm:  db,
		name: name,
	}, nil
}

func (c *defaultClient) ORM() *gorm.DB {
	return c.orm
}
