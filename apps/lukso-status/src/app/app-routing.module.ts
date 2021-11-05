import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { AvailableVersionsComponent } from './components/available-versions/available-versions.component';
import { InstallComponent } from './components/install/install.component';
import { StatusComponent } from './components/status/status.component';

const routes: Routes = [
  { path: '', component: InstallComponent },
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
