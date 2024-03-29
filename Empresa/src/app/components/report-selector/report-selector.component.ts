import { Component, OnInit, Input } from '@angular/core';
import { Anio, SearchMonth } from './../../interfaces/requests';
import { ServicesService } from './../../services/services.service';

@Component({
  selector: 'app-report-selector',
  templateUrl: './report-selector.component.html',
  styleUrls: ['./report-selector.component.css']
})
export class ReportSelectorComponent implements OnInit {

  constructor(private request : ServicesService) { }

  @ Input() years : Anio[] = [];
  year : Number = 0;
  month : Number = 0;
  searchMonth : SearchMonth = {Anio : 0, Mes : 0};

  ngOnInit(): void {
  }

  getGraphYears(){
    this.request.getGraphYears().subscribe();
  }

  getGraphMonths(){
    this.request.getGraphMonths(this.year).subscribe();
  }

  getGraphCalendar(){

  }

  setYear(year:Number){
    this.year = year;
    console.log(this.year);
  }

  setMonth(year:Number, Month:Number){
    this.year = year;
    this.month = Month;
    this.searchMonth.Anio = this.year;
    this.searchMonth.Mes = this.month;
    console.log(this.year+"          "+this.month);
  }
}
