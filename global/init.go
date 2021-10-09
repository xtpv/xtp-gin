package global

func InitResource() error {
	err := initConfig()
	if err != nil {
		return err
	}

	err = initDB()
	if err != nil {
		return err
	}
	return nil
}
