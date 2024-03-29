package chat

import (
	"bufio"
	"io"
)

type client struct {
	*bufio.Reader
	*bufio.Writer
	wc chan string
}

// StartClient - start messaging
func StartClient(msg chan<- string, cn io.ReadWriteCloser, quit chan struct{}) (chan<- string, chan struct{}) {
	c := new(client)
	c.Reader = bufio.NewReader(cn)
	c.Writer = bufio.NewWriter(cn)
	c.wc = make(chan string)
	done := make(chan struct{})

	go func() {
		scanner := bufio.NewScanner(c.Reader)
		for scanner.Scan() {
			logger.Println(scanner.Text())
			msg <- scanner.Text()
		}
		done <- struct{}{}
	}()

	c.writeMonitor()

	go func() {
		select {
		case <-quit:
			cn.Close()
		case <-done:
		}
	}()
	return c.wc, done
}

func (c *client) writeMonitor() {
	go func() {
		for s := range c.wc {
			c.WriteString(s + "\n")
			c.Flush()
		}
	}()
}
