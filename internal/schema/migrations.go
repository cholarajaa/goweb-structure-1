package schema

import (
	"github.com/GuiaBolso/darwin"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add Products New table",
		Script: `
		CREATE TABLE products_new (
			product_id UUID,
			name varchar(255),
			price INT,
			quantity INT,

			PRIMARY KEY (product_id)
		);`,
	},
}
