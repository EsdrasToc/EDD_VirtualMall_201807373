export interface Requests {
}

export interface Mes{
    Mes : String,
    Valor : Number,
}

export interface Anio{
    Anio : Number,
    Meses : Mes[]
}

export interface SearchMonth{
    Anio : Number,
    Mes : Number,
}

export interface Departament{
    Nombre : String;
    Dias : CalendarDay[];
}

export interface CalendarDay{
    Existe : Boolean;
    Numero : Number;
}

export interface Calendar{
    Departamentos :  Departament[];
    Dias : Number[];
}