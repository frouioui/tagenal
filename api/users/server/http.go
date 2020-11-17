package server

func (usersrv *UserServerAPI) RunServerHTTP() (err error) {
	err = usersrv.ServerHTTP.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
