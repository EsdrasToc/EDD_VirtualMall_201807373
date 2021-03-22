import { Component, Input, Output, EventEmitter } from '@angular/core';
import { RequestsService } from './../../services/requests.service'
import { Product, Shop, CarProduct } from './../../interfaces/requests';

@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css'],
  providers:[RequestsService]
})
export class ProductosComponent{

  /*@ Input() nombre : String;
  @ Input() codigo: Number;
	@ Input() descripcion: String;
	@ Input() precio : Number;
	@ Input() cantidad : Number;
	@ Input() imagen : String;*/
  @ Input() producto : Product;
  @ Input() tienda : Shop;

  @ Output() response : EventEmitter<Product> = new EventEmitter();

  constructor() { }

  addToCar(){

    //console.log(this.finalProduct);
    this.response.emit(this.producto);
  }

}
