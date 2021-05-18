package goleafcore_test

// func testFiber(t *testing.T) {
// 	glinit.InitLog()
// 	db := glinit.InitDb()
// 	server := glinit.InitServer()

// 	go glinit.StartServer(server)
// 	defer glinit.EndServer(db)

// 	doneTestFiber := make(chan bool)

// 	go func() {
// 		time.Sleep(10 * time.Second)
// 		doneTestFiber <- true
// 	}()

// 	logrus.Debug("Done testing ", <-doneTestFiber)
// }
