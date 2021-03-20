import { Component, OnInit } from '@angular/core';
import { RequestsService } from './services/requests.service';
import { Shop, Product } from './interfaces/requests';
import { ProductosComponent } from './components/productos/productos.component';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [RequestsService]
})
export class AppComponent implements OnInit{
  title = 'FrontendServer';
  shops : Shop[] = [];
  products : Product[] = [];
  shopName : String = "";
  shopScore : Number = 0;

  constructor(
    private requestService:RequestsService/*,
    private productosComponent:ProductosComponent*/
  ){}

  ngOnInit(){
    this.requestService.getShops().subscribe(data => {
      this.shops = data;
    });

    /*this.requestService.getProducts("Samsung", 3).subscribe(data => {
      this.products = data;
    });*/

  }

  clear(){
    this.shops = [];
  }

  getProducts(name:String, score:Number){
    this.requestService.getProducts(name, score).subscribe(data => {
      console.log(data)
      this.products = data;
    });

    //this.shopName = name;
    //this.shopScore = score;
  }
}
