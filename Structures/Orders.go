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
	Next       *Year     `json:"-"`
	Previous   *Year     `json:"-"`
}

/*func (this *Year) AddYear(year int) *Year {
	aux := this

	if this == nil {
		this = &Year{Value: year}
		return this
	}

	newYear := &Year{Value: year}
	for aux != nil {
		fmt.Println(year)
		fmt.Println(aux.Previous)
		if aux == this && newYear.Value < this.Value {
			newYear.Next = this
			this.Next = nil
			this.Previous = newYear
			this = newYear
			break
		} else if aux.Next == nil {
			aux.Next = newYear
			newYear.Previous = aux
			break
		} else if aux.Previous.Value < newYear.Value && aux.Value > newYear.Value {
			this.Previous.Next = newYear
			newYear.Previous = this.Previous
			newYear.Next = this
			this.Previous = newYear
			break
		}

		aux = aux.Next
	}

	return this
}

func (this *Year) SearchYear(year int) *Year {
	aux := this

	for aux != nil {
		if aux.Value == year {
			return aux
		}

		aux = aux.Next
	}

	return &Year{}
}

func (this *Year) GraphYears() string {
	var nodes, lines string
	aux := this

	if this == nil {
		return ""
	} else if this.Next == nil {
		return "digraph G{\n nodeA[label=\"" + strconv.Itoa(this.Value) + "\"]\n}"
	}

	for aux.Next != nil {
		if aux == this {
			lines = "\nnode" + strconv.Itoa(this.Value) + "->node" + strconv.Itoa(aux.Next.Value)
		} else {
			lines = "\nnode" + strconv.Itoa(this.Value) + "->node" + strconv.Itoa(aux.Next.Value)
			lines = "\nnode" + strconv.Itoa(this.Value) + "->node" + strconv.Itoa(aux.Previous.Value)
		}

		nodes = nodes + "\nnode" + strconv.Itoa(this.Value) + "[label=\"" + strconv.Itoa(this.Value) + "\"]"

		aux = aux.Next
	}

	nodes = nodes + "\nnode" + strconv.Itoa(this.Value) + "[label=\"" + strconv.Itoa(this.Value) + "\"]"
	lines = "\nnode" + strconv.Itoa(this.Value) + "->node" + strconv.Itoa(aux.Previous.Value)

	return "digraph G{\nrankdir=\"LR\"\n" + nodes + lines + "\n}"
}*/

func (this *Year) GraphMonths() string {
	var lines, nodes string
	aux := this.FirstMonth
	nodes = "nodeAnio[label=\"" + strconv.Itoa(this.Value) + "\"]"

	if aux.Next == nil {
		return "digraph G{\nrankdir=\"LR\"\nnodeHola[label=\"" + aux.Month + "\"]\n}"
	}
	for aux.Next != nil {
		if aux == this.FirstMonth {
			lines = "\nnodeAnio->node" + aux.Month + "\nnode" + aux.Month + "->node" + aux.Next.Month
		} else {
			lines = lines + "\nnode" + aux.Month + "->node" + aux.Previous.Month + "\nnode" + aux.Month + "->node" + aux.Next.Month
		}

		nodes = nodes + "\nnode" + aux.Month + "[label=\"" + aux.Month + "\"]"

		aux = aux.Next
	}

	nodes = nodes + "\nnode" + aux.Month + "[label=\"" + aux.Month + "\"]"
	lines = lines + "\nnode" + aux.Month + "->node" + aux.Previous.Month

	return "digraph G{\nrankdir=\"LR\"\n" + nodes + lines + "\n}"
}

func (this *Year) GraphNodes() string {
	if this == nil {
		return ""
	} else if (this.right == nil && this.left == nil) || (this.right.Value == 0 && this.left.Value == 0) {
		return "node" + strconv.Itoa(this.Value) + "[label=\"" + strconv.Itoa(this.Value) + "\"]\n"
	} else if this.right == nil || this.right.Value == 0 {
		return "node" + strconv.Itoa(this.Value) + "[label=\"" + strconv.Itoa(this.Value) + "\"]\n" + this.left.GraphNodes()
	} else if this.left == nil || this.left.Value == 0 {
		return "node" + strconv.Itoa(this.Value) + "[label=\"" + strconv.Itoa(this.Value) + "\"]\n" + this.right.GraphNodes()
	} else {
		return "node" + strconv.Itoa(this.Value) + "[label=\"" + strconv.Itoa(this.Value) + "\"]\n" + this.right.GraphNodes() + this.left.GraphNodes()
	}
}

func (this *Year) GraphLines() string {
	if this == nil {
		return ""
	} else if (this.right == nil && this.left == nil) || (this.right.Value == 0 && this.left.Value == 0) {
		return ""
	} else if this.right == nil || this.right.Value == 0 {
		return "\nnode" + strconv.Itoa(this.Value) + "->" + "node" + strconv.Itoa(this.left.Value) + this.left.GraphLines()
	} else if this.left == nil || this.left.Value == 0 {
		return "\nnode" + strconv.Itoa(this.Value) + "->" + "node" + strconv.Itoa(this.right.Value) + this.right.GraphLines()
	} else {
		return "\nnode" + strconv.Itoa(this.Value) + "->" + "node" + strconv.Itoa(this.left.Value) + "\nnode" + strconv.Itoa(this.Value) + "->" + "node" + strconv.Itoa(this.right.Value) + this.left.GraphLines() + this.right.GraphLines()
	}
}

func (this *Year) ToJson() string {
	/*if this.left == nil && this.right == nil {
		json = "{\n Anio : " + string(this.Value) + ",\nMeses : [" + this.MonthsToJson() + "]\n"
	}*/

	if this == nil {
		return ""
	} else if (this.right == nil && this.left == nil) || (this.right.Value == 0 && this.left.Value == 0) {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}\n"
	} else if this.right == nil || this.right.Value == 0 {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}" + ",\n" + this.left.ToJson()
	} else if this.left == nil || this.left.Value == 0 {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}" + ",\n" + this.right.ToJson()
	} else {
		return "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}" + ",\n" + this.right.ToJson() + ",\n" + this.left.ToJson()
	}

	/*aux := this
	json := ""

	if this == nil {
		return ""
	}

	for aux.Next != nil {
		json = json + "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n},\n"

		aux = aux.Next
	}
	json = json + "{\n \"Anio\" : " + strconv.Itoa(this.Value) + ",\n\"Meses\" : [" + this.MonthsToJson() + "]\n}\n"

	return json*/
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

	/*fmt.Println(newMonth.Month + " added in ")
	fmt.Println(this.Value)
	fmt.Println(" year")*/
}

func (avl *Year) Insert(newNode Year) *Year { // Insertar valor
	//exist := false
	value := newNode.Value
	if avl == nil { // Se inserta el nodo raíz
		fmt.Println("Insertando el nodo raiz " + strconv.Itoa(value))
		node := &Year{
			Value:      value,
			FirstMonth: newNode.FirstMonth,
			LastMonth:  newNode.LastMonth,
		}
		avl = node
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
		fmt.Println("De aca sali")
		if avl.right.height()-avl.left.height() == 2 {
			if value < avl.right.Value {
				avl = avl.right_left() // derecha izquierda
			} else {

				avl = avl.right_right() // derecha derecha
			}
		}
	} else {
		fmt.Println(strconv.Itoa(avl.Value) + "Ya existe")
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

/*func (avl *Year) Insert(value int) *Year { // Insertar valor
	if value == 2020 {
		fmt.Println("Iteracion")
	}
	if avl == nil { // Se inserta el nodo raíz
		if value == 2020 {
			fmt.Println("Que hace aqui?")
		}
		node := &Year{
			Value: value,
		}
		avl = node
		avl.left = &Year{Value: 0}
		avl.right = &Year{Value: 0}
		fmt.Println("Year " + strconv.Itoa(avl.Value) + " add")
	} else if avl.Value == 0 {
		if value == 2020 {
			fmt.Println("Cero")
			fmt.Println(avl.left.Value)
			fmt.Println(avl.right.Value)
		}
		//avl = nil
		//fmt.Println("Entro aqui :c")
		node := &Year{
			Value: value,
		}
		avl = node
		avl.left = &Year{Value: 0}
		avl.right = &Year{Value: 0}
		//fmt.Println("Year " + strconv.Itoa(avl.Value) + " add")
	} else if value < avl.Value { // Insertar en el subárbol izquierdo
		if value == 2020 {
			fmt.Println("Menor?")
		}
		avl.left = avl.left.Insert(value)              // recorrer para encontrar la posición que se va a insertar
		if avl.left.height()-avl.right.height() == 2 { // Juzgar el equilibrio
			if value < avl.left.Value { // Izquierda Izquierda
				avl = avl.left_left()
			} else { // Acerca de
				avl = avl.left_right()
			}
		}
	} else if value > avl.Value { // Inserta el subárbol derecho
		if value == 2020 {
			fmt.Println("Mayor")
		}
		avl.right = avl.right.Insert(value)
		if avl.right.height()-avl.left.height() == 2 {
			if value < avl.right.Value {
				avl = avl.right_left() // derecha izquierda
			} else {

				avl = avl.right_right() // derecha derecha
			}
		}
	} else {
		if value == 2020 {
			fmt.Println("WTH?")
		}
		fmt.Println("the key", value, "has exists")
	}
	avl.high = max1(avl.right.height(), avl.left.height()) + 1 // Actualizar altura

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
	if avl == nil || avl.Value == 0 {
		return 0
	} else {
		return avl.high
	}
}*/

func (avl *Year) InOrder() {
	if avl == nil {
		return
	}
	avl.left.InOrder()
	fmt.Println(avl.Value)
	avl.right.InOrder()
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

func (this Calendar) NumberOfRows() int {
	aux := this.Start
	var i int
	for aux != nil {
		if aux.Down == nil {
			return aux.Row
		}
		aux = aux.Down
	}

	return i
}

func (this Calendar) NumberOfColumns() int {
	aux := this.Start
	var i int

	for aux != nil {
		i++
		aux = aux.Right
	}

	return i
}

func (this Calendar) RowsToArray() string {
	json := ""

	aux := this.Start

	for aux != nil {
		if aux.RowName != "" {

		}
	}

	return json
}

/*func (this Calendar) ToShow() string {
	json := ""
	rows := this.NumberOfRows()
	columns := this.NumberOfColumns()

	return ""
}*/

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
	if node == nil || node.Value == 0 {
		return
	}

	YearInorden(node.left)
	fmt.Println("=== Año: ")
	fmt.Println(node.Value)
	//node.ViewMonths()
	YearInorden(node.right)
}
