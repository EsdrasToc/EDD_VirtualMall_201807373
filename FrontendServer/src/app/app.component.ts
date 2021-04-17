import { Component, OnInit } from '@angular/core';
import { RequestsService } from './services/requests.service';
import { Shop, Product, CarProduct, Authenticate, User } from './interfaces/requests';

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
  car :  CarProduct[] = [];
  currentShop : Shop;
  tempProduct : CarProduct = {}; 

  currentUser :  User;
  authenticate : boolean = false;
  aut : Authenticate = {Dpi : 0, Password:""};

  constructor(
    private requestService:RequestsService/*,
    private productosComponent:ProductosComponent*/
  ){}

  ngOnInit(){


    this.requestService.getShops().subscribe(data => {
      //console.log(dd+" "+mm+" "+yyyy)
      console.log(data);
      this.shops = data;
    });
    /*this.requestService.getProducts("Samsung", 3).subscribe(data => {
      this.products = data;
    });*/

  }

  Autenticar(){
    console.log(this.aut)
    this.requestService.authenticate(this.aut).subscribe(data =>{
      console.log(data)
      this.currentUser = data;
      this.currentUser.Dpi = data.Dpi

      if (data.Dpi != 0 && data.Correo != ""){
        console.log("Lo estoy dejando pasar")
        this.authenticate = true
      }
    });
  }

  clear(){
    this.shops = [];
  }

  getProducts(name:String, score:Number, shop : Shop){
    this.requestService.getProducts(name, score).subscribe(data => {
      console.log(data)
      this.products = data;
      this.currentShop = shop;
      console.log(this.currentShop);
    });
    //this.shopName = name;
    //this.shopScore = score;
  }

  addToCar($event:Product){
    //this.tempProduct.Producto = $event;
    var today = new Date();
    var dd = String(today.getDate()).padStart(2, '0');
    var mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
    var yyyy = today.getFullYear();

    if(this.car.length != 0){
      var find = false;
      for (let i = 0; i < this.car.length; i++) {
        if(this.car[i].Tienda?.Calificacion == this.currentShop.Calificacion && this.car[i].Tienda?.Nombre == this.currentShop.Nombre){
          console.log("Ya entre");
          this.car[i].Producto?.push({
            Nombre : $event.Nombre,
            Codigo : $event.Codigo,
            Descripcion : $event.Descripcion,
            Precio : $event.Precio,
            Cantidad : $event.Cantidad,
	          Imagen : $event.Imagen
          });
          find = true;
          break;
        }
      }

      if(!find){
        /*this.tempProduct.Producto = [];
        this.tempProduct.Producto.push($event);
        this.tempProduct.Tienda = this.currentShop;
        this.car.push(this.tempProduct);*/
        this.car.push(
          {
            Tienda : {
              Nombre : this.currentShop.Nombre,
              Descripcion : this.currentShop.Descripcion,
              Contacto : this.currentShop.Contacto,
              Calificacion: this.currentShop.Calificacion,
              Logo : this.currentShop.Logo,
            },
            Fecha : dd+"-"+mm+"-"+yyyy,
            Producto : [
              {
                Nombre : $event.Nombre,
                Codigo : $event.Codigo,
                Descripcion : $event.Descripcion,
                Precio : $event.Precio,
                Cantidad : $event.Cantidad,
                Imagen : $event.Imagen
              }
            ]
          }
        );
      }
    }else{
      /*this.tempProduct.Producto = [];
      this.tempProduct.Producto.push($event);
      this.tempProduct.Tienda = this.currentShop;
      this.car.push(this.tempProduct);*/

      this.car.push(
        {
          Tienda : {
            Nombre : this.currentShop.Nombre,
            Descripcion : this.currentShop.Descripcion,
            Contacto : this.currentShop.Contacto,
            Calificacion: this.currentShop.Calificacion,
            Logo : this.currentShop.Logo,
          },
          Fecha  : dd+"-"+mm+"-"+yyyy,
          Producto : [
            {
              Nombre : $event.Nombre,
              Codigo : $event.Codigo,
              Descripcion : $event.Descripcion,
              Precio : $event.Precio,
              Cantidad : $event.Cantidad,
              Imagen : $event.Imagen
            }
          ]
        }
      );
    }

    console.log(this.car);
  }

  Buy(){
    console.log("HOLA MUNDO");
    console.log(this.car);
    this.requestService.putPurchase(this.car).subscribe();
    this.requestService.putOrder(this.car).subscribe();
  }
}
