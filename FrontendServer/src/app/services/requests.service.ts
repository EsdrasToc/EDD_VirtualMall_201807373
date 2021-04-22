import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
//import jsonData from '../';

import { Shop, Product, CarProduct, Authenticate, User } from './../interfaces/requests';

@Injectable({
  providedIn: 'root'
})
export class RequestsService {

  constructor(private http:HttpClient) { }

  getShops(){
    const path = 'http://localhost:3000/getshops';
    return this.http.get<Shop[]>(path);
  }

  getProducts(name:String, score:Number){
    const path = 'http://localhost:3000/getProducts/'+name+'/'+score;
    console.log(path)
    return this.http.get<Product[]>(path);
  }

  putPurchase(products : CarProduct[]){
    const path = "http://localhost:3000/putPurchase";
    return this.http.put<CarProduct[]>(path, products);
  }

  authenticate(data : Authenticate){
    const path = "http://localhost:3000/Authenticate";
    console.log("holiiiis")
    console.log(data)
    return this.http.post<User>(path, data);
  }

  putOrder(data : CarProduct[]){
    const path = "http://localhost:3000/putOrder";
    return this.http.put<CarProduct[]>(path, data);
  }

  NewUser(user : User){
    const path = "http://localhost:3000/NewUser"

    return this.http.put(path, user);
  }
}