package ws

import (
	"fmt"

	"github.com/Nykenik24/enkrypted/internal/middleware"
)

func (c *Client) AddMiddleware(midware middleware.Middleware) error {
	name := midware.Info().Name
	if _, ok := c.midware[name]; ok {
		return fmt.Errorf("middleware '%s' already in client", name)
	}
	c.midware[name] = midware

	return nil
}

func (c *Client) RemoveMiddleware(midware middleware.Middleware) {
	delete(c.midware, midware.Info().Name)
}

func (c *Client) RemoveMidwareByName(name string) {
	delete(c.midware, name)
}
