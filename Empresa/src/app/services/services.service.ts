import { Injectable } from '@angular/core';
import { Anio, Calendar, SearchMonth } from '../interfaces/requests';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root'
})
export class ServicesService {
  
  constructor(private http:HttpClient) { }

  getYears(){
    const path = 'http://localhost:3000/getYears';
    return this.http.get<Anio[]>(path);
  }

  getGraphYears(){
    const path = 'http://localhost:3000/getGraphYears';
    return this.http.get<any>(path);
  }

  getGraphMonths(year:Number){
    const path = 'http://localhost:3000/getGraphMonths/'+year;
    console.log(path);
    return this.http.get<any>(path);
  }

  getDaysOfMonth(search : SearchMonth){
    const path = 'http://localhost:3000/Month';
    
    return this.http.put<Calendar>(path, search)
  }

  getGraphUsers(){
    const path = 'http://localhost:3000/GraphAccounts'

    return this.http.get(path);
  }

  PostShops(body : String){
    const path = 'http://localhost:3000/cargartienda'

    return this.http.post(path, body);
  }
}
