package Structures

import (
	"fmt"
	"strconv"
)

/*=================================*/
/*   Inicia el arbol avl de años   */
/*=================================*/
type Year struct {
	Value      int       `json:"Anio"`
	FirstMonth *Calendar `json:"-"`
	LastMonth  *Calendar `json:"-"`
	left       *Year     `json:"-"`
	right      *Year     `json:"-"`
	high       int       `json:"-"`
}

func (this *Year) ToJson() string {
	/*if this.left == nil && this.right == nil {
		json = "{\n Anio : " + string(this.Value) + ",\nMeses : [" + this.MonthsToJson() + "]\n"
	}*/

	if this == nil {
		return ""
	} else if this.right == nil && this.left == nil {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}\n"
	} else if this.right == nil {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}" + ",\n" + this.left.ToJson()
	} else if this.left == nil {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}" + ",\n" + this.right.ToJson()
	} else {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}" + ",\n" + this.right.ToJson() + ",\n" + this.left.ToJson()
	}
}

func (this *Year) MonthsToJson() string {
	json := ""
	aux := this.FirstMonth

	if this.FirstMonth == nil {
		return ""
	}

	for aux.Next != nil {
		json = json + "{\n \"Mes\" : \"" + aux.Month + "\",\n\"Valor\" : " + strconv.Itoa(aux.Value) + "\n},\n"

		aux = aux.Next
	}
	json = json + "{\n \"Mes\" : \"" + aux.Month + "\",\n\"Valor\" : " + strconv.Itoa(aux.Value) + "\n}"

	return json
}

func (this *Year) ViewMonths() {
	aux := this.FirstMonth

	for aux != nil {
		fmt.Println("===== Mes:")
		fmt.Println(aux.Month)
		aux.ViewCalendar()
		aux = aux.Next
	}
}
func (this *Year) SearchMonth(monthValue int) *Calendar {
	aux := this.FirstMonth

	for aux != nil {
		if aux.Value == monthValue {
			return aux
		}

		aux = aux.Next
	}

	return &Calendar{}
}

func (this *Year) SearchYear(year int) *Year {

	if this == nil {
		return &Year{}
	}

	if this.Value == year {
		return this
	}

	if year > this.Value {
		if this.right != nil {
			return this.right.SearchYear(year)
		} else {
			return &Year{}
		}
	}

	if year < this.Value {
		if this.left != nil {
			return this.left.SearchYear(year)
		} else {
			return &Year{}
		}
	}

	return &Year{}
}

func (this *Year) AddMonth(newMonth *Calendar) {
	if this.FirstMonth == nil {
		this.FirstMonth = newMonth
	}

	if this.LastMonth == nil {
		this.LastMonth = newMonth
	} else {
		this.LastMonth.Next = newMonth
		newMonth.Previous = this.LastMonth
		this.LastMonth = newMonth
	}

	fmt.Println(newMonth.Month + " added in ")
	fmt.Print(this.Value)
	fmt.Print(" year")
}

func (avl *Year) Insert(newNode Year) *Year { // Insertar valor
	//exist := false
	value := newNode.Value
	if avl == nil { // Se inserta el nodo raíz
		fmt.Println("Insertando el nodo raiz")
		node := Year{
			Value:      value,
			FirstMonth: newNode.FirstMonth,
			LastMonth:  newNode.LastMonth,
		}
		avl = &node
		avl.left = nil
		avl.right = nil
	} else if value < avl.Value { // Insertar en el subárbol izquierdo
		avl.left = avl.left.Insert(newNode)            // recorrer para encontrar la posición que se va a insertar
		if avl.left.height()-avl.right.height() == 2 { // Juzgar el equilibrio
			if value < avl.left.Value { // Izquierda Izquierda
				avl = avl.left_left()
			} else { // Acerca de
				avl = avl.left_right()
			}
		}
	} else if value > avl.Value { // Inserta el subárbol derecho
		avl.right = avl.right.Insert(newNode)
		if avl.right.height()-avl.left.height() == 2 {
			if value < avl.right.Value {
				avl = avl.right_left() // derecha izquierda
			} else {

				avl = avl.right_right() // derecha derecha
			}
		}
	} else {
		//exist = true
	}
	avl.high = max1(avl.left.height(), avl.right.height()) + 1 // Actualizar altura

	return avl
}

func (avl *Year) left_left() *Year {

	k := avl.left      // Determine la nueva raíz
	avl.left = k.right // rotar
	k.right = avl
	avl.high = max1(avl.left.height(), avl.right.height()) + 1
	k.high = max1(k.left.height(), avl.high) + 1
	return k
}

func (avl *Year) left_right() *Year {
	avl.left = avl.left.left_left() // Primero gira a la izquierda el subárbol izquierdo
	return avl.right_right()        // Rota a la derecha este nodo
}

func (avl *Year) right_left() *Year {
	avl.right = avl.right.right_right()
	return avl.left_left()
}

func (avl *Year) right_right() *Year {

	k := avl.right
	avl.right = k.left
	k.left = avl

	avl.high = max1(avl.left.height(), avl.right.height()) + 1
	k.high = max1(avl.high, k.right.height()) + 1

	return k
}

func (avl *Year) height() int {
	if avl == nil {
		return 0
	} else {
		return avl.high
	}
}

/*=================================*/
/*  Finaliza el arbol avl de años  */
/*=================================*/

/*=================================*/
/*Inicia matriz dispersa de pedidos*/
/*=================================*/

type Calendar struct {
	Start    *Day
	Month    string
	Value    int
	Previous *Calendar
	Next     *Calendar
}

func (m *Calendar) ViewCalendar() {
	tmprow := m.Start
	for tmprow != nil {
		tmpcol := m.Start
		fmt.Print(tmprow.Row, ",", tmpcol.Column, " ")
		tmpcol = tmprow.Right
		if tmprow.Row == 0 {
			for tmpcol != nil {
				fmt.Print(tmprow.Row, ",", tmpcol.Column, " ")
				tmpcol = tmpcol.Right
			}
		} else {
			for tmpcol != nil {
				fmt.Print(tmprow.Row, ",", getCol(tmpcol), "(", tmpcol.Column, ") ")
				tmpcol = tmpcol.Right
			}
		}
		tmprow = tmprow.Down
		fmt.Println()
	}
}

func getCol(tmp *Day) int {
	var col int
	for tmp.Up != nil {
		tmp = tmp.Up
		col = tmp.Column
	}
	return col
}

func (this *Calendar) Init(value int) {
	month := ""
	switch value {
	case 1:
		month = "Enero"
	case 2:
		month = "Febrero"
	case 3:
		month = "Marzo"
	case 4:
		month = "Abril"
	case 5:
		month = "Mayo"
	case 6:
		month = "Junio"
	case 7:
		month = "Julio"
	case 8:
		month = "Agosto"
	case 9:
		month = "Septiembre"
	case 10:
		month = "Octubre"
	case 11:
		month = "Noviembre"
	default:
		month = "Diciembre"
	}

	this.Month = month
	this.Value = value
	this.Start = &Day{Row: 0, Column: 0}
}

func (this *Calendar) AddRow(departament string) (bool, *Day) {
	day := this.Start
	newRow := &Day{RowName: departament}
	if this.Start.Down == nil {
		this.Start.Down = newRow
		newRow.Up = this.Start
		newRow.Row = newRow.Up.Row + 1
		return true, newRow
	} else {
		day = day.Down
		for &day != nil {
			if day.Down == nil && day.RowName != departament {
				day.Down = newRow
				newRow.Up = day
				newRow.Row = newRow.Up.Row + 1
				return true, newRow
			} else if day.RowName == departament {
				break
			}

			day = day.Down
		}
	}
	return false, day
}

func (this *Calendar) AddColumn(dayN int) (bool, *Day) {
	day := this.Start
	newColumn := &Day{Column: dayN}

	if this.Start.Right == nil {
		this.Start.Right = newColumn
		newColumn.Left = this.Start
		return true, newColumn
	} else {
		day = day.Right
		for &day != nil {
			if day.Column > dayN && day.Left.Column < dayN {
				day.Left.Right = newColumn
				newColumn.Left = day.Left
				day.Left = newColumn
				newColumn.Right = day
				return true, newColumn
			} else if day.Right == nil && day.Column != dayN {
				day.Right = newColumn
				newColumn.Left = day
				return true, newColumn
			} else if day.Column == dayN {
				break
			}

			day = day.Right
		}
	}

	return false, day
}

func (this *Calendar) AddDay(row *Day, column *Day, firstElementR bool, firstElementC bool) *Day {
	newDay := &Day{Row: row.Row, Column: column.Column}
	tempRow := row
	tempCol := column
	var auxRow *Day
	//var auxCol *Day

	//lo añadimos primero en la fila
	if firstElementR {
		row.Right = newDay
		newDay.Left = row
		auxRow = newDay
	} else {
		tempRow = tempRow.Right
		for tempRow != nil {
			if tempRow.Column > newDay.Column {
				tempRow.Left.Right = newDay
				newDay.Left = tempRow.Left
				newDay.Right = tempRow
				tempRow.Left = newDay

				auxRow = newDay
			} else if tempRow.Right == nil && tempRow.Column != newDay.Column {
				tempRow.Right = newDay
				newDay.Left = tempRow

				auxRow = newDay
			} else if tempRow.Column == newDay.Column {
				auxRow = tempRow
				break
			}

			tempRow = tempRow.Right
		}
	}

	//Ahora, añadimos en la columna
	if firstElementC {
		column.Down = newDay
		newDay.Up = column
		newDay.Column = getCol(newDay)
		//auxCol = newDay
	} else {
		tempCol = tempCol.Down
		for tempCol != nil {
			if tempCol.Row > newDay.Row {
				tempCol.Up.Down = newDay
				newDay.Up = tempCol.Up
				newDay.Down = tempCol
				tempCol.Up = newDay
				newDay.Column = getCol(newDay)

				//auxCol = newDay
			} else if tempCol.Down == nil && tempCol.Row != newDay.Row {
				tempCol.Down = newDay
				newDay.Up = tempCol
				newDay.Column = getCol(newDay)

				//auxCol = newDay
			} else if tempCol.Row == newDay.Row {
				//auxCol = tempCol
				break
			}

			tempCol = tempCol.Down
		}
	}

	return auxRow
}

func (this *Calendar) AddOrder(departament string, dayN int, order *Order) {
	var firstElementC, firstElementR bool
	var col, row, auxRow *Day

	firstElementC, col = this.AddColumn(dayN)
	firstElementR, row = this.AddRow(departament)
	auxRow = this.AddDay(row, col, firstElementR, firstElementC)

	auxRow.AddOrder(order)
}

type Day struct {
	Up         *Day
	Down       *Day
	Left       *Day
	Right      *Day
	FirstOrder Order
	Row        int
	RowName    string
	Column     int
}

func (this *Day) AddOrder(newOrder *Order) {
	auxOrder := &this.FirstOrder

	if &this.FirstOrder == nil {
		this.FirstOrder = *newOrder
	} else {
		for auxOrder != nil {
			if auxOrder.Next == nil {
				auxOrder.Next = newOrder
				break
			}
			auxOrder = auxOrder.Next
		}
	}
}

type Order struct {
	Date        string    `json:"Fecha"`
	Shop        string    `json:"Tienda"`
	Departament string    `json:"Departamento"`
	Score       int       `json:"Calificacion"`
	Products    []Product `json:"Productos"`
	Next        *Order    `json:"-"`
}

/*===================================*/
/*Finaliza matriz dispersa de pedidos*/
/*===================================*/

type Orders struct {
	Orders []Order `json:"Pedidos"`
}

func YearInorden(node *Year) {
	if node == nil {
		return
	}

	YearInorden(node.left)
	fmt.Println("=== Año: ")
	fmt.Println(node.Value)
	node.ViewMonths()
	YearInorden(node.right)
}
