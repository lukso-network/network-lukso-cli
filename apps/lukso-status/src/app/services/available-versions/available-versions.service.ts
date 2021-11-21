import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';
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
  availableSoftware$: Observable<Releases[] | null>;
  downloadedSoftware$: Observable<any>;

  constructor(private httpClient: HttpClient) {
    this.httpClient = httpClient;
    this.downloadedSoftware$ = httpClient.get('/api/downloaded-versions');

    this.availableSoftware$ = httpClient
      .get<AvailableSoftwareBackendResponse>('/api/available-versions')
      .pipe(
        map((availableSoftware) => {
          return Object.entries(availableSoftware).map(([name, release]) => {
            return {
              name,
              humanReadableName: release.humanReadableName,
              downloadInfo: Object.entries(release.downloadInfo)
                .map(([tag, { downloadUrl }]) => {
                  console.log(downloadUrl);
                  return { tag, name, downloadUrl } as DownloadInfo;
                })
                .reverse(),
            } as Releases;
          });
        }),
        catchError((error: any) => {
          return of(null);
        })
      );
  }

  downloadClient(client: string, version: string, url: string) {
    return this.httpClient.post('/api/update-client', {
      client,
      version,
      url,
    });
  }

  startClients(network: string, settings: Settings, clients: string[]) {
    return this.httpClient.post('/api/start-clients', {
      network,
      settings,
      clients,
    }) as Observable<string>;
  }

  stopClients(clients: string[]) {
    return this.httpClient.post('/api/stop-clients', {
      clients,
    });
  }

  getSettings(network: string) {
    return this.httpClient.get('/api/settings', {
      params: {
        network,
      },
    }) as Observable<Settings>;
  }

  setConfig(network: string, settings: Settings) {
    console.log(settings);
    return this.httpClient.post('/api/settings', {
      network,
      settings,
    }) as Observable<Settings>;
  }

  getAvailableVersions$() {
    return this.availableSoftware$;
  }

  getDownloadedVersions$() {
    return this.downloadedSoftware$;
  }
}
