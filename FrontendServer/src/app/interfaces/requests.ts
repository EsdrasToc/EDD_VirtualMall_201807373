export interface Requests {
}

export interface Shop{
    Nombre : String;
	Descripcion : String;
	Contacto : String;
	Calificacion: Number;
	Logo : String;
}

export interface Product{
	Nombre:String;
	Codigo:Number;
	Descripcion:String;
	Precio:Number;
	Cantidad:Number;
	Imagen:String;
}

export interface CarProduct{
	Tienda ?: Shop;
	Producto ?: Product[];
}