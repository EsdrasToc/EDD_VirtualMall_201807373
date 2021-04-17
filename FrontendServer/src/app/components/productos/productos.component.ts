import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { RequestsService } from './../../services/requests.service'
import { Product, Shop, CarProduct } from './../../interfaces/requests';

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
  cantidad : number = 1;
  stock : boolean = true;

  @ Output() response : EventEmitter<Product> = new EventEmitter();

  constructor() { }

  ngOnInit(){
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

}
