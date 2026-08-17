package hasher

type Hasher interface {
	HashPassword(password string) (string, error)
	CheckPassword(hashedPassword, password string) error
}
