import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AvailableVersionsComponent } from './components/available-versions/available-versions.component';
import { SettingsComponent } from './components/settings/settings.component';
import { StatusComponent } from './components/status/status.component';

const routes: Routes = [
  { path: '', component: StatusComponent },
  { path: 'settings', component: SettingsComponent },
  { path: 'updates', component: AvailableVersionsComponent },
  { path: 'status', component: StatusComponent },
  {
    path: 'launchpad',
    loadChildren: () =>
      import('./modules/launchpad/launchpad.module').then(
        (m) => m.LaunchpadModule
      ),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
