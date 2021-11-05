import { Component } from '@angular/core';
import { Observable } from 'rxjs';
import { AvailableVersionsService } from '../../services/available-versions/available-versions.service';

@Component({
  selector: 'lukso-available-versions',
  templateUrl: './available-versions.component.html',
  styleUrls: ['./available-versions.component.css'],
})
export class AvailableVersionsComponent {
  downloadedSoftware$: Observable<any>;
  availableSoftware$: Observable<any>;
  constructor(availableVersions: AvailableVersionsService) {
    this.downloadedSoftware$ = availableVersions.getDownloadedVersions$();
    this.availableSoftware$ = availableVersions.getAvailableVersions$();
  }
}
