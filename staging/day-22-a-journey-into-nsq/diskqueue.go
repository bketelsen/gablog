for msg := range c.incomingMsgChan {
	select {
	case c.memoryMsgChan <- msg:
	default:
		err := WriteMessageToBackend(&msgBuf, msg, c.backend)
		if err != nil {
			// ... handle errors ...
		}
	}
}
