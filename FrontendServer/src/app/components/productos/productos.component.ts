import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { RequestsService } from './../../services/requests.service'
import { Product, Shop, Comentario, User, ForComment } from './../../interfaces/requests';

@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css'],
  providers:[RequestsService]
})
export class ProductosComponent implements OnInit{

  /*@ Input() nombre : String;
  @ Input() codigo: Number;
	@ Input() descripcion: String;
	@ Input() precio : Number;
	@ Input() cantidad : Number;
	@ Input() imagen : String;*/
  @ Input() producto : Product;
  @ Input() tienda : Shop;
  @ Input() user: User;
  //@ Input() comentarios : Comentario[];
  cantidad : number = 1;
  stock : boolean = true;

  comentario : String = "";

  @ Output() response : EventEmitter<Product> = new EventEmitter();

  constructor(private request : RequestsService) { }

  ngOnInit(){
    console.log("HOLA MUNDO, S I ESTOY CARGANDO")
  }

  addToCar(){

    //console.log(this.finalProduct);
    this.response.emit(
      {
        Nombre : this.producto.Nombre,
        Codigo : this.producto.Codigo,
        Descripcion : this.producto.Descripcion,
        Precio : this.producto.Precio,
        Cantidad : this.cantidad,
        Imagen : this.producto.Imagen,
        Almacenamiento : this.producto.Almacenamiento
      }
    );

    this.cantidad = 1;
  }

  comentar(){
    this.request.ProductComment(
      {
        Producto : this.producto,
        Contenido : this.comentario,
      }
    ).subscribe();
  }

}
