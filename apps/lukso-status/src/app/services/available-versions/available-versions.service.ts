import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map, switchMap, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class SoftwareService {
  availableSoftware$: Observable<any>;
  downloadedSoftware$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.httpClient = httpClient;
    this.downloadedSoftware$ = httpClient.get('/api/downloaded-versions').pipe(
      map((versions) => {
        return Object.entries(versions).map(([name, versions]) => {
          return { name, versions };
        });
      })
    );

    this.availableSoftware$ = httpClient.get('/api/available-versions').pipe(
      map((versions) => {
        console.log(versions);
        return Object.entries(versions).map(([name, versions]) => {
          return {
            name,
            versions: Object.entries(versions)
              .map(([tag, url]) => {
                return { tag, url };
              })
              .reverse(),
          };
        });
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

  startClients() {
    return this.httpClient.post('/api/start-clients', {});
  }

  stopClients() {
    return this.httpClient.post('/api/stop-clients', {});
  }

  getAvailableVersions$() {
    return this.availableSoftware$;
  }

  getDownloadedVersions$() {
    return this.downloadedSoftware$;
  }
}
