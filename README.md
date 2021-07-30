# Filter Products
- A basic Go app with mongo driver for filtering the products.

## To Configure
- Import your data to MongoDB and change `flags/config.json` file. (Don't forget to change `DBNAME`)
- Change collection name which is located in the `db/db.go`
  - Default is `const collectionName = "clothes"`. You can change as you like.

## Usage
- in terminal: `go run .`

## Some Example Filterings
- `products?productName=slim&productColor=gri&productCompany=defacto`
- `products?productName=jean&gender=kadın&productCompany=koton`

##### ps. Parameters are case insensitive.