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
  authenticate : boolean = false;
  DPI: Number = 0;
  Correo: String = "";
  Password: String = "";
  Nombre: String = "";

  time :Number = 1;

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

  Authenticate(){
    if (this.DPI == 1234567890101 && this.Correo == "auxiliar@edd.com" && this.Password == "1234" && this.Nombre == "EDD2021"){
      this.authenticate = true;
    }
  }

  GraphMO(){
    this.requestService.GraphMO().subscribe()
  }

  GraphMP(){
    this.requestService.GraphMP().subscribe()
  }

  GraphMU(){
    this.requestService.GraphMU().subscribe()
  }

  GraphMS(){
    this.requestService.GraphMS().subscribe()
  }

  Time(){
    this.requestService.ChangeTime(this.time).subscribe()
  }
}