import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';

@Component({
  selector: 'lukso-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent {
  title = 'lukso-status';
  constructor(private http: HttpClient) {
    this.http = http;
  }

  updateClient() {
    this.http
      .post('/api/update-client', {
        client: 'lukso-status',
        version: 'v0.0.1-alpha.9',
      })
      .subscribe(() => {
        console.log('success');
      });
  }
}
