export interface Requests {
}

export interface Shop{
    Nombre : String;
	Descripcion : String;
	Contacto : String;
	Calificacion: Number;
	Logo : String;
	Comentarios ?: Comentario[];
}

export interface Product{
	Nombre ? :String;
	Codigo ? :Number;
	Descripcion ? :String;
	Precio ? :Number;
	Cantidad ? :number;
	Imagen ? :String;
	Almacenamiento ? : String;
	Comentarios ?: Comentario[];
}

export interface CarProduct{
	Tienda ?: Shop;
	Fecha ? : String;
	Producto ?: Product[];
}

export interface User{
	Dpi : Number;
	Nombre : String;
	Correo : String;
	Password :  String;
	Cuenta : String;
}

export interface Authenticate{
	Dpi : Number;
	Password : String;
}

export interface Pedido{
	Fecha : String;
	Tienda : String;
	Departamento : String;
	Calificacion : Number;
	Productos : Product[];
} 

export interface Comentario{
	Usuario : User;
	Contenido : String;
	Comentarios : Comentario[];
}

export interface ForComment{
	Producto : Product;
	Contenido : String;
}

export interface ShopComment{
	Tienda : Shop;
	Contenido : String;
}

export interface ShopComments{
	Tienda : Shop;
	Contenido : String;
	Usuario : User;
}