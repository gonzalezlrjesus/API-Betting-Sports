package models

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
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
	Timerestante   int64
	Matrix         []MatrixRemates
	Actualposition MatrixRemates
}

// MatrixRemates model MatrixRemates
type MatrixRemates struct {
	Monto         int64
	MatrixRow     int
	NumeroCaballo int
	Seudonimo     string
	idCaballo     int
	Horsename     string
	idRacing      uint
}

// Manager export to controller
var Manager = ClientManager{
	broadcast:  make(chan []byte),
	Register:   make(chan *Clientmodel),
	unregister: make(chan *Clientmodel),
	clients:    make(map[*Clientmodel]bool),
}

var myInt8 int64 = 16
var arrayRemates []MatrixRemates
var actualPosition MatrixRemates
var idCarrera uint
var finalizacion string
var montoTotal int64

// Start export func
func (manager *ClientManager) Start() {

	for {
		select {
		case conn := <-manager.Register:

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

			if actualPosition.idRacing != 0 {
				// fmt.Println("IF ENTRO", actualPosition.idRacing)
				if TimeisEqualStartTime(actualPosition.idRacing) {
					// fmt.Println("APROBADO CERRADO", actualPosition.idRacing)
					arrayRemates = nil
					actualPosition = MatrixRemates{}
					myInt8, _ = strconv.ParseInt(os.Getenv("TimebetweenAuctions"), 10, 64)
					myInt8 = myInt8 + 1
					idCarrera = uint(0)
					montoTotal = 0
				}
			}

			if len(manager.clients) == 0 {
				myInt8, _ = strconv.ParseInt(os.Getenv("TimebetweenAuctions"), 10, 64)
				// arrayRemates = nil
			}

			if len(manager.clients) > 0 {
				myInt8 = myInt8 - 1

				if myInt8 < 0 {
					myInt8, _ = strconv.ParseInt(os.Getenv("TimebetweenAuctions"), 10, 64)

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

		jsonMessage2, _ := json.Marshal(parsedData["content"])
		s, _ := strconv.Unquote(string(jsonMessage2))
		var parsedData2 map[string]interface{}
		errjson := json.Unmarshal([]byte(s), &parsedData2)
		log.Println("errjson: *** ", parsedData2)
		log.Println("errjson: *** ", errjson)

		tempID := parsedData2["idcarrera"].(string)
		u64, err := strconv.ParseUint(tempID, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		idCarrera := uint(u64)
		finalizacion = parsedData2["finalizo"].(string)

		matrixfloat64 := int64(parsedData2["matrix"].(float64))
		matrixRowfloat64 := int(parsedData2["matrixRow"].(float64))
		numeroCaballofloat64 := int(parsedData2["numerocaballo"].(float64))
		seudonimoActual := parsedData2["seudonimo"].(string)
		HorsenameActual := parsedData2["horsename"].(string)
		idhorse := int(parsedData2["idcaballo"].(float64))

		var a MatrixRemates //
		a.Monto = matrixfloat64
		a.MatrixRow = matrixRowfloat64
		a.NumeroCaballo = numeroCaballofloat64
		a.Seudonimo = seudonimoActual
		a.idCaballo = idhorse
		a.Horsename = HorsenameActual
		a.idRacing = idCarrera

		var respaldoActual MatrixRemates //
		respaldoActual.Monto = actualPosition.Monto
		respaldoActual.MatrixRow = actualPosition.MatrixRow
		respaldoActual.NumeroCaballo = actualPosition.NumeroCaballo
		respaldoActual.Seudonimo = actualPosition.Seudonimo
		respaldoActual.Horsename = actualPosition.Horsename
		respaldoActual.idCaballo = actualPosition.idCaballo
		respaldoActual.idRacing = actualPosition.idRacing

		actualPosition.Monto = a.Monto
		actualPosition.MatrixRow = int(parsedData2["matrixRowSiguiente"].(float64))
		actualPosition.NumeroCaballo = a.NumeroCaballo
		actualPosition.Seudonimo = a.Seudonimo
		actualPosition.Horsename = a.Horsename
		actualPosition.idCaballo = a.idCaballo
		actualPosition.idRacing = a.idRacing

		// fmt.Println("actualPosition BEFORE ---", actualPosition)
		// fmt.Println("          a MatrixRemates BEFORE ---", a)
		// fmt.Println("                                     respaldoActual BEFORE ---", respaldoActual)

		if respaldoActual != actualPosition {
			if a.Seudonimo == "CASA" {

				client := &Client{}
				err := GetDB().Table("clients").Where("seudonimo = ?", a.Seudonimo).First(client).Error
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						fmt.Println("err not found seudonimo websocket CASA:", err)

					}

				}

				temp := &Coins{Clientidentificationcard: client.Identificationcard}

				//check client_Coins in DB
				errCoins := GetDB().Table("coins").Where("ClientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
				if errCoins == gorm.ErrRecordNotFound {
					fmt.Println("Client Coins not found ClientIdentificationcard websocket: ", errCoins)

				}

				temp.DecreaseCoins(float64(a.Monto))

				CreateRemates(idCarrera, a.idCaballo, a.NumeroCaballo, a.Seudonimo, a.Monto, a.Horsename)
				montoTotal = montoTotal + a.Monto
				// fmt.Println("Monto CASA:", montoTotal)
				if finalizacion == "finalizo" {
					fmt.Println("Monto finalizo CASA:", montoTotal)
					CreateTablas(idCarrera, montoTotal)
					arrayRemates = nil
					CloseRacing(idCarrera)
					idCarrera = uint(0)
					montoTotal = 0
				}

			} else if a.Seudonimo == "vacio" {

				client := &Client{}
				err := GetDB().Table("clients").Where("seudonimo = ?", respaldoActual.Seudonimo).First(client).Error
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						fmt.Println("err not found seudonimo websocket: CLIENTE ", err)

					}

				}

				temp := &Coins{Clientidentificationcard: client.Identificationcard}

				//check client_Coins in DB
				errCoins := GetDB().Table("coins").Where("ClientIdentificationcard = ?", temp.Clientidentificationcard).First(temp).Error
				if errCoins == gorm.ErrRecordNotFound {
					fmt.Println("Client Coins not found ClientIdentificationcard websocket: ", errCoins)

				}

				temp.DecreaseCoins(float64(respaldoActual.Monto))
				CreateRemates(idCarrera, respaldoActual.idCaballo, respaldoActual.NumeroCaballo, respaldoActual.Seudonimo, respaldoActual.Monto, respaldoActual.Horsename)
				montoTotal = montoTotal + respaldoActual.Monto
				// fmt.Println("Monto VACIO:", montoTotal)
				if finalizacion == "finalizo" {
					fmt.Println("Monto finalizo CLIENTE:", montoTotal)
					CreateTablas(idCarrera, montoTotal)
					arrayRemates = nil
					CloseRacing(idCarrera)
					idCarrera = uint(0)
					montoTotal = 0
				}

			}

		}

		if a.Monto != -1 {
			arrayRemates = append(arrayRemates, a)
		}

		if finalizacion == "finalizo" {
			arrayRemates = nil
			// idCarrera = ""
		}

		// fmt.Println("arrayRemates ", arrayRemates)
		myInt8, _ = strconv.ParseInt(os.Getenv("TimebetweenAuctions"), 10, 64)
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
			c.Socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
