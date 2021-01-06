package client

func InitGRPCClient() (err error) {
	err = InitUsersGRPC()
	// err = InitArticlesGRPC()
	return err
}
