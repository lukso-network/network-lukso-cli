import { Injectable } from '@angular/core';
import { CanActivate, Router, UrlTree } from '@angular/router';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { SoftwareService } from '../services/available-versions/available-versions.service';

@Injectable({
  providedIn: 'root',
})
export class SetupGuard implements CanActivate {
  constructor(
    private softwareService: SoftwareService,
    private router: Router
  ) {}
  canActivate(): Observable<boolean | UrlTree> {
    return this.softwareService.getDownloadedVersions$().pipe(
      map((a) => {
        return !(
          a &&
          Object.keys(a).length === 0 &&
          Object.getPrototypeOf(a) === Object.prototype
        );
      }),
      map((setupPerformed) => {
        if (!setupPerformed) {
          return this.router.parseUrl('/setup');
        }
        return true;
      })
    );
  }
}
