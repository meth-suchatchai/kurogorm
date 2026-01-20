package kurogorm

import (
	"crypto/md5"
	"encoding/hex"
	dbmodels "github.com/meth-suchatchai/kurogorm/daos"
	"log"
)

var createPermissions = []dbmodels.Permission{
	{
		Name:        "Create Role",
		Code:        "CREATE_ROLE",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Create Permission",
		Code:        "CREATE_PERMISSION",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Read Role Permission",
		Code:        "READ_ROLE_PERMISSION",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Read Blog",
		Code:        "READ_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Create Blog",
		Code:        "CREATE_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Update Blog",
		Code:        "UPDATE_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Delete Blog",
		Code:        "DELETE_BLOG",
		Description: "",
		IsActive:    true,
	},
	{
		Name:        "Admin Audit",
		Code:        "AUDIT",
		Description: "",
		IsActive:    true,
	},
}

var createKubePermissions = []dbmodels.Permission{
	{
		Name:        "Kubernetes full access control",
		Code:        "KUBE_FULL_ACCESS",
		Description: "",
		IsActive:    true,
	},
}

func (c *defaultClient) Seed() {
	tx := c.orm.Create(createPermissions)

	tx = c.orm.Create(&dbmodels.Role{
		Name:        "Admin",
		Description: "full access control",
		Permission:  createPermissions,
	})
	if tx.Error != nil {
		log.Fatalf("error seed data: %v", tx.Error)
	}

	encryptdPass := EncryptedHash("qwertyuiop")
	tx = c.orm.Create(&dbmodels.User{
		MobileNumber:      "8023736019",
		CountryCode:       "81",
		FullName:          "Kuroshibz",
		IsActive:          true,
		PasswordEncrypted: encryptdPass,
		TFEnable:          false,
		TFCode:            "",
		Permission:        createPermissions,
	})
}

func (c *defaultClient) SeedKubePermissions() error {
	db := c.orm.Create(createKubePermissions)

	if db.Error != nil {
		return db.Error
	}
	return nil
}

func EncryptedHash(data string) string {
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
