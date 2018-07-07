package authentication

type Strategy interface {
	Authenticate(id, password string) error
}
