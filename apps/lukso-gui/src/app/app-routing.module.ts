import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AvailableVersionsComponent } from './components/available-versions/available-versions.component';
import { SettingsComponent } from './components/settings/settings.component';
import { SetupComponent } from './components/setup/setup.component';
import { StatusComponent } from './components/status/status.component';
import { SetupGuard } from './guards/setup.guard';

const routes: Routes = [
  { path: '', component: StatusComponent, canActivate: [SetupGuard] },
  { path: 'settings', component: SettingsComponent, canActivate: [SetupGuard] },
  {
    path: 'updates',
    component: AvailableVersionsComponent,
    canActivate: [SetupGuard],
  },
  { path: 'setup', component: SetupComponent },
  { path: 'status', component: StatusComponent, canActivate: [SetupGuard] },
  {
    path: 'launchpad',
    loadChildren: () =>
      import('./modules/launchpad/launchpad.module').then(
        (m) => m.LaunchpadModule
      ),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes, { enableTracing: true })],
  exports: [RouterModule],
})
export class AppRoutingModule {}
