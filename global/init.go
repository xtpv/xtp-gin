package global

func InitResource() (err error) {
	err = initConfig()
	if err != nil {
		return
	}

	err = initDB()
	if err != nil {
		return
	}

	err = initLogger()
	if err != nil {
		return
	}

	return
}
