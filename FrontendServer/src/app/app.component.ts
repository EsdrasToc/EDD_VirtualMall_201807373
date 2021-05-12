import { Component, OnDestroy, OnInit } from '@angular/core';
import { RequestsService } from './services/requests.service';
import { Shop, Product, CarProduct, Authenticate, User, Comentario } from './interfaces/requests';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [RequestsService]
})
export class AppComponent implements OnDestroy{
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
  comment :string ="";

  newUser : User = {Dpi:0, Nombre:"", Password:"", Correo:"", Cuenta:""};

  constructor(
    private requestService:RequestsService/*,
    private productosComponent:ProductosComponent*/
  ){}

  ngOnDestroy(){
    this.requestService.Blocks().subscribe();
  }


  Autenticar(){
    this.requestService.getShops().subscribe(data => {
      //console.log(dd+" "+mm+" "+yyyy)
      console.log(data);
      this.shops = data;
    });
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
      this.products = data;
      this.currentShop = shop;
      console.log(this.currentShop);
      console.log(this.products);
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

  Register(){
    this.requestService.NewUser(this.newUser).subscribe();
  }

  CommentShop(shop : Shop){
    this.requestService.CommentShop({
      Tienda : shop,
      Contenido : this.comment,
    }).subscribe()

    console.log({
      Tienda : shop,
      Contenido : this.comment,
    })
  }

  SCommentShop(shop : Shop, comment0 : Comentario){
    this.requestService.SCommentShop(
      [
        {
          Tienda : shop,
          Contenido : comment0.Contenido,
          Usuario : comment0.Usuario
        },
        {
          Tienda : shop,
          Contenido : this.comment,
          Usuario : this.currentUser
        }
      ]
    ).subscribe()
  }

  SSCommentShop(shop : Shop, comment0 : Comentario, comment1 : Comentario){
    this.requestService.SSCommentShop(
      [
        {
          Tienda : shop,
          Contenido : comment0.Contenido,
          Usuario : comment0.Usuario
        },{
          Tienda : shop,
          Contenido : comment1.Contenido,
          Usuario : comment1.Usuario
        },
        {
          Tienda : shop,
          Contenido : this.comment,
          Usuario : this.currentUser
        }
      ]
    ).subscribe()
  }
}
