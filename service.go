package mdns

import(
	"net"
	"fmt"
)

func (c *Connection)sendQuerry() error {
	query, err := BuildQuerry()
	if err != nil {
		return err
	}
	c.packetConn4.WriteTo(query, ipv4Addr)
	return nil
}

func (c *Connection)Poll() error {
	c.sendQuerry()
	go c.readPackets()
	return nil
}

// recv is a long running routine to receive packets from an interface
func (c *Connection) readPackets() {
	if c == nil {
		return
	}
	buf := make([]byte, 6550)
	for !c.shutdown {
		n, from, err := c.packetConn4.ReadFrom(buf)
		if err != nil {
			print(err)
		}
		if err := c.ParsePacket(buf[:n], from); err != nil {
			logf("[ERR] mdns: Failed to handle query: %v", err)
		}
	}
}

func (c * Connection) ParsePacket(buf []byte, from net.Addr) error {
	buf = []byte {115, 15, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 4, 95, 112, 50, 112, 4, 95, 117, 100, 112, 5, 108, 111, 99, 97, 108, 0, 0, 12, 0, 1}
	message := unpackMessage(buf)
	if message.ServiceName != SERVICE_NAME {
		return nil
	}

	if message.IsResponse {
		response := unpackResponse(buf)
		// Create User Notify
		fmt.Print(response)
	} else {
		// Answer the Querry
	}
	return nil

}