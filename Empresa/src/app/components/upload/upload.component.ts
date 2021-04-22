import { Component, OnInit } from '@angular/core';
import { ServicesService } from './../../services/services.service';

@Component({
  selector: 'app-upload',
  templateUrl: './upload.component.html',
  styleUrls: ['./upload.component.css']
})
export class UploadComponent implements OnInit {

  Tiendas : String = "";
  constructor(private request : ServicesService) { }

  ngOnInit(): void {
  }

  PutShops(){
    this.request.PostShops(this.Tiendas).subscribe;
  }

}
