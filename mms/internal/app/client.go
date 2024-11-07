package app

func RunClient() {
	// flag.Parse()

	// cc, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatal("Failed to connect to server ", err)
	// }

	// clientUser := pb.NewUserServiceClient(cc)

	// myhandler := api.NewHandler(clientUser)

	// app := myhandler.Init()

	// go func() {
	// 	if err := app.Listen(":5000"); err != nil {
	// 		log.Fatal("server error ", err)
	// 	}
	// }()

	// stop := make(chan os.Signal, 1)
	// signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	// <-stop

	// log.Fatal("Shutting down the server...")

}
