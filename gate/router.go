package gate

import (
    "server/game"
    "server/msg"
)

func init() {
    // Route Hello to game
    // All communication are through ChanRPC including the management messages
    msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
