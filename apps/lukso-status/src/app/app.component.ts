import { HttpClient } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { switchMap } from 'rxjs/operators';
import { Settings } from './interfaces/settings';
import { NETWORKS } from './modules/launchpad/launchpad/helpers/create-keys';
import { SoftwareService } from './services/available-versions/available-versions.service';
import { DataService } from './services/data.service';

@Component({
  selector: 'lukso-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  title = 'lukso-status';
  NETWORKS = NETWORKS;

  dataService: DataService;
  softwareService: SoftwareService;

  form: FormGroup = new FormGroup({
    network: new FormControl('', [Validators.required]),
  });

  constructor(
    private http: HttpClient,
    dataService: DataService,
    softwareService: SoftwareService
  ) {
    this.dataService = dataService;
    this.softwareService = softwareService;
  }

  ngOnInit() {
    const network =
      (localStorage.getItem('network') as NETWORKS) || NETWORKS.L15_DEV;
    this.dataService.setNetwork(network);
    this.form.controls.network.setValue(network);
    this.form.controls.network.valueChanges.subscribe((network) => {
      this.dataService.setNetwork(network);
    });
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

  startClients(network: string) {
    this.softwareService
      .getConfig(network)
      .pipe(
        switchMap((settings: Settings) => {
          console.log(settings, network);
          return this.softwareService.startClients(network, settings);
        })
      )
      .subscribe();
  }

  get networkFormCtrl() {
    return this.form.get('network') as FormControl;
  }
}
