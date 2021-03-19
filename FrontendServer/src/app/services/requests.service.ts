import { Injectable } from '@angular/core';
import { HttpClient, HttpClientModule } from '@angular/common/http';
//import jsonData from '../';

import { Shop } from './../interfaces/requests';

@Injectable({
  providedIn: 'root'
})
export class RequestsService {

  constructor(private http:HttpClient) { }

  getShops(){
    const path = 'http://localhost:3000/getshops';
    console.log(this.http.get<Shop[]>(path));
    return this.http.get<Shop[]>(path);
  }
}