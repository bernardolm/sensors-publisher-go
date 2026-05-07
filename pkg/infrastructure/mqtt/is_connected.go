package mqtt

func (c *Client) IsConnected() bool {
	return c.client != nil && c.client.IsConnected()
}
