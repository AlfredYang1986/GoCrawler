package Config

type CGConfig struct {

	Details []struct{
		Name string
		Method string
		Args map[string]string
	}

	Lists []struct{
		Name string
		Method string
		Args map[string]string
	}

	Entry []struct{
		Name string
		Title string
		Method string
		Url string
		Handler string
	}
}
