package models

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

// ClientManager model ClientManager
type ClientManager struct {
	clients    map[*Clientmodel]bool
	broadcast  chan []byte
	Register   chan *Clientmodel
	unregister chan *Clientmodel
}

// Clientmodel model Clientmodel
type Clientmodel struct {
	Id     string
	Socket *websocket.Conn
	Send   chan []byte
}

// Message model message
type Message struct {
	Sender         string `json:"sender,omitempty"`
	Recipient      string `json:"recipient,omitempty"`
	Content        string `json:"content,omitempty"`
	Timerestante   int8
	Matrix         []MatrixRemates
	Actualposition MatrixRemates
}

// MatrixRemates model MatrixRemates
type MatrixRemates struct {
	Monto     int64
	MatrixRow int
	MatrixCol int
	Seudonimo string
	idCaballo int
}

// Manager export to controller
var Manager = ClientManager{
	broadcast:  make(chan []byte),
	Register:   make(chan *Clientmodel),
	unregister: make(chan *Clientmodel),
	clients:    make(map[*Clientmodel]bool),
}

var myInt8 int8 = 16
var arrayRemates []MatrixRemates
var actualPosition MatrixRemates

// Debo hallar una forma de almacenar todos los eventos que han transcurrido
// para que cuando ingrese un cliente se le refleje todo lo que ha ocurrido
// y tener la row y col actual osea que el size de la matrix debe ser
// extraido de conectar esta parte de websocket con la bdd y model de racing and horse

// Start export func
func (manager *ClientManager) Start() {

	for {
		select {
		case conn := <-manager.Register:

			if len(manager.clients) > 0 {

			}

			manager.clients[conn] = true
			jsonMessage, _ := json.Marshal(&Message{Sender: "NewConnection", Content: "/A socket NEW.", Matrix: arrayRemates, Actualposition: actualPosition})
			conn.Send <- jsonMessage
			manager.Send(jsonMessage, conn)
		case conn := <-manager.unregister:

			if _, ok := manager.clients[conn]; ok {
				close(conn.Send)
				delete(manager.clients, conn)
				jsonMessage, _ := json.Marshal(&Message{Content: "/A socket has disconnected."})
				manager.Send(jsonMessage, conn)
			}
		case message := <-manager.broadcast:

			for conn := range manager.clients {
				select {
				case conn.Send <- message:
				default:
					close(conn.Send)
					delete(manager.clients, conn)
				}
			}
		case <-time.After(1 * time.Second):

			// Creo que el countdown deberia empezar a correr cuando
			// el tiempo de remate coincida con la hora transcurriendo
			// sin importar si hay cliente o no
			// de la forma como esta actualmente si nungun cliente ingresa al
			//  remate jamas va a comenzar a o terminar el remate
			// osea el remate deberia comenzar haya o no cliente conectado
			// tambien implica que debo conectarlo con el resto de el backedn
			// para extraer todo estos datos

			if len(manager.clients) == 0 {
				myInt8 = 15
				arrayRemates = nil
			}

			if len(manager.clients) > 0 {
				myInt8 = myInt8 - 1

				if myInt8 < 0 {
					myInt8 = 15

				}

				jsonMessage, _ := json.Marshal(&Message{Sender: "time", Timerestante: (myInt8), Matrix: arrayRemates})
				for conn := range manager.clients {
					select {
					case conn.Send <- jsonMessage:
					default:
						close(conn.Send)
						delete(manager.clients, conn)
					}
				}
			}

		}
	}
}

// Send send message to websocket
func (manager *ClientManager) Send(message []byte, ignore *Clientmodel) {
	for conn := range manager.clients {
		if conn != ignore {
			// fmt.Println("Send Send :", string(message))
			conn.Send <- message
		}
	}
}

func (c *Clientmodel) Read() {
	defer func() {
		Manager.unregister <- c
		c.Socket.Close()
	}()

	for {
		_, message, err := c.Socket.ReadMessage()
		if err != nil {
			Manager.unregister <- c
			c.Socket.Close()
			break
		}

		jsonMessage, _ := json.Marshal(&Message{Sender: c.Id, Content: string(message)})
		var parsedData map[string]interface{}
		json.Unmarshal(jsonMessage, &parsedData)
		// fmt.Println("Resulting parsedData : ", parsedData["content"])

		jsonMessage2, _ := json.Marshal(parsedData["content"])
		s, _ := strconv.Unquote(string(jsonMessage2))
		var parsedData2 map[string]interface{}
		errjson := json.Unmarshal([]byte(s), &parsedData2)
		fmt.Println("errjson: *** ", errjson)

		matrixfloat64 := int64(parsedData2["matrix"].(float64))
		matrixRowfloat64 := int(parsedData2["matrixRow"].(float64))
		matrixColfloat64 := int(parsedData2["matrixCol"].(float64))
		seudonimoActual := parsedData2["seudonimo"].(string)
		idhorse := int(parsedData2["idcaballo"].(float64))

		var a MatrixRemates // a == Student{"", 0}
		a.Monto = matrixfloat64
		a.MatrixRow = matrixRowfloat64
		a.MatrixCol = matrixColfloat64
		a.Seudonimo = seudonimoActual
		a.idCaballo = idhorse

		var respaldoActual MatrixRemates // a == Student{"", 0}
		respaldoActual.Monto = actualPosition.Monto
		respaldoActual.MatrixRow = actualPosition.MatrixRow
		respaldoActual.MatrixCol = actualPosition.MatrixCol
		respaldoActual.Seudonimo = actualPosition.Seudonimo
		respaldoActual.idCaballo = actualPosition.idCaballo

		actualPosition.Monto = a.Monto
		actualPosition.MatrixRow = int(parsedData2["matrixRowSiguiente"].(float64))
		actualPosition.MatrixCol = int(parsedData2["matrixColSiguiente"].(float64))
		actualPosition.Seudonimo = a.Seudonimo
		actualPosition.idCaballo = a.idCaballo

		// actualPosition.MatrixCol = int(parsedData2["matrixColSiguiente"].(float64))
		fmt.Println("actualPosition BEFORE ---", actualPosition)
		fmt.Println("          a MatrixRemates BEFORE ---", a)
		fmt.Println("                                     respaldoActual BEFORE ---", respaldoActual)
		// TEST Reducir cantidad de coins si se termina la seccion del remate del caballo
		// osea si la siguiente row es porque termino de rematar el caballo

		if respaldoActual != actualPosition {
			if a.Seudonimo == "CASA" {
				fmt.Println("CASA se guarda a A")
				CreateRemates(109, a.idCaballo, a.Seudonimo, a.Monto)

			} else if a.Seudonimo == "vacio" {
				fmt.Println("vacio se guarda a actual")
				client := &Client{}
				err := GetDB().Table("clients").Where("seudonimo = ?", respaldoActual.Seudonimo).First(client).Error
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						fmt.Println("err not found seudonimo websocket:", err)

					}

				}
				fmt.Println("client get", client)
				temp := &Coins{Clientidentificationcard: client.Identificationcard}

				//check client_Coins in DB
				errCoins := GetDB().Table("coins").Where("ClientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
				if errCoins == gorm.ErrRecordNotFound {
					fmt.Println("Client Coins not found ClientIdentificationcard websocket: ", errCoins)

				}

				temp.DecreaseCoins(float64(respaldoActual.Monto))
				CreateRemates(109, respaldoActual.idCaballo, respaldoActual.Seudonimo, respaldoActual.Monto)
			} else if a.MatrixCol == 2 && a.Monto != -1 {
				fmt.Println("TERCERA CASILLA FULL se guarda a MatrixRemates")
				client := &Client{}
				err := GetDB().Table("clients").Where("seudonimo = ?", a.Seudonimo).First(client).Error
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						fmt.Println("err not found seudonimo websocket:", err)

					}

				}
				fmt.Println("client get", client)
				temp := &Coins{Clientidentificationcard: client.Identificationcard}

				//check client_Coins in DB
				errCoins := GetDB().Table("coins").Where("ClientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
				if errCoins == gorm.ErrRecordNotFound {
					fmt.Println("Client Coins not found ClientIdentificationcard websocket: ", errCoins)

				}

				temp.DecreaseCoins(float64(a.Monto))
				CreateRemates(109, a.idCaballo, a.Seudonimo, a.Monto)
			}
			// }
		}

		// actualPosition.MatrixRow = int(parsedData2["matrixRowSiguiente"].(float64))

		// actualPosition.Monto = a.Monto
		// actualPosition.MatrixRow = int(parsedData2["matrixRowSiguiente"].(float64))
		// actualPosition.MatrixCol = int(parsedData2["matrixColSiguiente"].(float64))
		// actualPosition.Seudonimo = a.Seudonimo
		// actualPosition.idCaballo = a.idCaballo
		// fmt.Println("actualPosition AFTER ---: ", actualPosition)

		if a.Monto != -1 {
			arrayRemates = append(arrayRemates, a)
		}

		fmt.Println("arrayRemates:", arrayRemates)

		myInt8 = 15
		Manager.broadcast <- jsonMessage
	}
}

func (c *Clientmodel) Write() {
	defer func() {
		c.Socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.Socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			// fmt.Println("message write :", string(message))
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
