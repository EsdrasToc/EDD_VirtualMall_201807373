import { Component, OnInit } from '@angular/core';
import { Input, Output } from '@angular/core';
import { CarProduct } from './../../interfaces/requests';

@Component({
  selector: 'app-carrito',
  templateUrl: './carrito.component.html',
  styleUrls: ['./carrito.component.css']
})
export class CarritoComponent {

  @ Input() car : CarProduct[];

  constructor() { }

}
