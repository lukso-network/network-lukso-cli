import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { map, switchMap, tap } from 'rxjs/operators';

@Injectable({
  providedIn: 'root',
})
export class AvailableVersionsService {
  availableSoftware$: Observable<any>;
  downloadedSoftware$: Observable<any>;
  constructor(private httpClient: HttpClient) {
    this.downloadedSoftware$ = httpClient.get('/api/downloaded-versions').pipe(
      map((versions) => {
        return Object.entries(versions).map(([name, versions]) => {
          return { name, versions };
        });
      }),
      tap((a) => {
        console.log(a);
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
      }),
      tap((a) => {
        console.log(a);
      })
    );
  }

  getAvailableVersions$() {
    return this.availableSoftware$;
  }

  getDownloadedVersions$() {
    return this.downloadedSoftware$;
  }
}
