import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '../../../../src/environments/environment';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import {
  AvailableSoftwareBackendResponse,
  DownloadInfo,
  Releases,
} from '../../interfaces/available-software';
import { Settings } from '../../interfaces/settings';

@Injectable({
  providedIn: 'root',
})
export class SoftwareService {
  availableSoftware$: Observable<Releases[]>;
  downloadedSoftware$: Observable<any>;

  constructor(private httpClient: HttpClient) {
    this.httpClient = httpClient;
    this.downloadedSoftware$ = httpClient.get(
      environment.API + '/downloaded-versions'
    );

    this.availableSoftware$ = httpClient
      .get<AvailableSoftwareBackendResponse>(
        environment.API + '/available-versions'
      )
      .pipe(
        map((availableSoftware) => {
          return Object.entries(availableSoftware).map(([name, release]) => {
            return {
              name,
              humanReadableName: release.humanReadableName,
              downloadInfo: Object.entries(release.downloadInfo)
                .map(([tag, { downloadUrl }]) => {
                  return { tag, name, downloadUrl } as DownloadInfo;
                })
                .reverse(),
            } as Releases;
          });
        })
      );
  }

  downloadClient(client: string, version: string, url: string) {
    return this.httpClient.post(environment.API + '/update-client', {
      client,
      version,
      url,
    });
  }

  startClients(network: string, settings: Settings, clients: string[]) {
    return this.httpClient.post(environment.API + '/start-clients', {
      network,
      settings,
      clients,
    }) as Observable<string>;
  }

  stopClients(clients: string[]) {
    return this.httpClient.post(environment.API + '/stop-clients', {
      clients,
    });
  }

  getSettings(network: string) {
    return this.httpClient.get(environment.API + '/settings', {
      params: {
        network,
      },
    }) as Observable<Settings>;
  }

  setSettings(network: string, settings: Settings) {
    return this.httpClient.post(environment.API + '/settings', {
      network,
      settings,
    });
  }

  getAvailableVersions$() {
    return this.availableSoftware$;
  }

  getDownloadedVersions$() {
    return this.downloadedSoftware$;
  }
}
