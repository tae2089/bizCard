// Code generated by entc, DO NOT EDIT.

package bizcard

const (
	// Label holds the string label denoting the bizcard type in the database.
	Label = "biz_card"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the bizcard in the database.
	Table = "biz_cards"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "biz_cards"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_id"
)

// Columns holds all SQL columns for bizcard fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldPhoneNumber,
	FieldEmail,
	FieldAge,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "biz_cards"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
)
