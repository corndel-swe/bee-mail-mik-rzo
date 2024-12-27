package main

type MessageAdapter struct {
	message *ExternalMessage
}

func (m *MessageAdapter) String() string {
	return m.message.String()
}

func (m *MessageAdapter) log() {
	m.message.log()
}

func (m *MessageAdapter) markDelivered() {
	m.message.markDelivered()
}

func (m *MessageAdapter) markRead() {
	if !m.message.read {
		m.message.toggleRead()
	}
}
