import { Component, Input } from '@angular/core';
import { RequestsService } from './../../services/requests.service'

@Component({
  selector: 'app-productos',
  templateUrl: './productos.component.html',
  styleUrls: ['./productos.component.css'],
  providers:[RequestsService]
})
export class ProductosComponent{

  @ Input() nombre : String;
  @ Input() codigo: Number;
	@ Input() descripcion: String;
	@ Input() precio : Number;
	@ Input() cantidad : Number;
	@ Input() imagen : String; 

  constructor() { }

}
