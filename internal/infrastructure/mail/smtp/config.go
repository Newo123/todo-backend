package smtp

type Config struct {
	Host     string `envconfig:"HOST" required:"true"`
	Port     string `envconfig:"PORT" default:"587"`
	User     string `envconfig:"USER" required:"true"`
	Password string `envconfig:"PASSWORD" required:"true"`
	From     string `envconfig:"From" required:"true"`
}

func NewConfig()
