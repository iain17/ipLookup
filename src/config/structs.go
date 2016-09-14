package config

type Config struct {
	HTTP struct {
			 BindingAddress string `toml:"bindingAddress"`
		 } `toml:"http"`
}
