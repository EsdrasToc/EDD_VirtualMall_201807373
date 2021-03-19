import { Component, OnInit } from '@angular/core';
import { RequestsService } from './services/requests.service';
import { Shop } from './interfaces/requests';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
  providers: [RequestsService]
})
export class AppComponent implements OnInit{
  title = 'FrontendServer';
  shops : Shop[] = [];

  constructor(
    private requestService:RequestsService
  ){}

  ngOnInit(){
    this.requestService.getShops().subscribe(data => {
      this.shops = data;
    });
  }
}
