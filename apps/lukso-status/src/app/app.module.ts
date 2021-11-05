import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { PandoraStatusComponent } from './components/status/pandora-status/pandora-status.component';
import { VanguardStatusComponent } from './components/status/vanguard-status/vanguard-status.component';
import { AvailableVersionsComponent } from './components/available-versions/available-versions.component';
import { AppRoutingModule } from './app-routing.module';
import { StatusComponent } from './components/status/status.component';
import { InstallComponent } from './components/install/install.component';
import { OrchestratorStatusComponent } from './components/status/orchestrator-status/orchestrator-status.component';
import { ValidatorStatusComponent } from './components/status/validator-status/validator-status.component';

@NgModule({
  declarations: [
    AppComponent,
    PandoraStatusComponent,
    VanguardStatusComponent,
    AvailableVersionsComponent,
    StatusComponent,
    InstallComponent,
    OrchestratorStatusComponent,
    ValidatorStatusComponent,
  ],
  imports: [AppRoutingModule, BrowserModule, HttpClientModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
