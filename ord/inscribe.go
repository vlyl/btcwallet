package ord

func Inscribe(option *Option, path string) {

	inscription, err := FromFile(option.ChainParam, path)
	if err != nil {
		panic(err)
	}

	// todo update index

	wallet
}
