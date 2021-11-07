import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { SoftwareService } from '../../services/available-versions/available-versions.service';

@Component({
  selector: 'lukso-available-versions',
  templateUrl: './available-versions.component.html',
  styleUrls: ['./available-versions.component.css'],
})
export class AvailableVersionsComponent {
  softwareService: SoftwareService;
  downloadedSoftware$: Observable<any>;
  availableSoftware$: Observable<any>;

  isDownloading = false;

  constructor(softwareService: SoftwareService) {
    this.softwareService = softwareService;
    this.downloadedSoftware$ = softwareService.getDownloadedVersions$();
    this.availableSoftware$ = softwareService.getAvailableVersions$();
  }

  install(client: string, release: any) {
    release.isDownloading = true;
    this.softwareService
      .downloadClient(client, release.tag, release.url)
      .subscribe(() => {
        release.isDownloading = false;
        console.log('success');
      });
  }
}
