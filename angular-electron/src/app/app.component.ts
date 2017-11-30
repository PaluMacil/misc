import { Component, OnInit } from '@angular/core';
import { ElectronService } from 'ngx-electron';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'app';
  appData = '';

  constructor(private _electronService: ElectronService) { }
electro
  ngOnInit () {
    this.appData = this._electronService.remote.app.getPath('userData');
  }
}
