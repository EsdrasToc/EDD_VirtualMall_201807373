import { Component, OnInit } from '@angular/core';
import { Mes, Anio } from './interfaces/requests';
import { ServicesService } from './services/services.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  years : Anio[] = [];

  constructor(
    private requestService:ServicesService
  ){}

  ngOnInit(){
   this.getYears() 
  }

  getYears(){
    this.requestService.getYears().subscribe(data => {
      this.years = data;
    });
  }
}