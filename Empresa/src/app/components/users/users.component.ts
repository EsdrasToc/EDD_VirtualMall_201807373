import { Component, OnInit } from '@angular/core';
import { ServicesService } from './../../services/services.service';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {

  constructor(private service: ServicesService) { }

  ngOnInit(): void {
  }

  GraphUsers(){
    this.service.getGraphUsers().subscribe();
  }

}
