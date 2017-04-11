package main

func main() {

	//	waiter := waiter.NewSignalWaiter(make(chan bool))
	//
	//	ln, err := net.Listen("tcp", ":8080")
	//	if err != nil {
	//		// handle error
	//	}
	//	for {
	//		conn, err := ln.Accept()
	//		if err != nil {
	//			// handle error
	//		}
	//		close := make(chan bool, 5)
	//		moderator := moderator.New(
	//			talker.NewCTTalker(conn, message.NewClimateTalkMessageUtil(), close, waiter),
	//			talker.NewMQTTTalker(message.NewMQTTMessageUtil(), mqttClientFactory, close, waiter),
	//			talker.NewForkTalker(forkerClientFactory, close, waiter),
	//			close,
	//			waiter)
	//		go handleConnection(moderator)
	//	}

}
