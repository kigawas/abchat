package services

// not useful if via redis

// import (
// 	"sync"

// 	"github.com/kigawas/abchat/websocket"
// )

// var wsClients sync.Map

// func SaveUserConnection(c *websocket.Conn, userID string) {
// 	wsClients.Store(userID, c)
// }

// func DeleteUserConnection(userID string) {
// 	wsClients.Delete(userID)
// }

// func GetUserConnection(userID string) *websocket.Conn {
// 	if conn, ok := wsClients.Load(userID); ok {
// 		return conn.(*websocket.Conn)
// 	}
// 	return nil
// }
