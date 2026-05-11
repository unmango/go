package os

func Openwd() (*Root, error) {
	wd, err := System.Getwd()
	if err != nil {
		return nil, err
	}
	return OpenRoot(wd)
}

func OpenInwd(name string) (File, error) {
	wd, err := System.Getwd()
	if err != nil {
		return nil, err
	}
	return OpenInRoot(wd, name)
}
