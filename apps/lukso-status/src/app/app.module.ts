import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { AvailableVersionsComponent } from './components/available-versions/available-versions.component';
import { AppRoutingModule } from './app-routing.module';
import { StatusComponent } from './components/status/status.component';
import { OrchestratorStatusComponent } from './components/status/orchestrator-status/orchestrator-status.component';
import { ValidatorStatusComponent } from './components/status/validator-status/validator-status.component';
import { SettingsComponent } from './components/settings/settings.component';
import { ReactiveFormsModule } from '@angular/forms';
import { NgxChartsModule } from '@swimlane/ngx-charts';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { NetworkStatusComponent } from './components/status/network-status/network-status.component';
import { TimeagoModule } from 'ngx-timeago';
import { StatusBoxComponent } from './components/status/status-box/status-box.component';

import { GLOBAL_RX_STATE, GlobalState } from './shared/rx-state';
import { RxState } from '@rx-angular/state';

@NgModule({
  declarations: [
    AppComponent,
    AvailableVersionsComponent,
    StatusComponent,
    OrchestratorStatusComponent,
    ValidatorStatusComponent,
    SettingsComponent,
    NetworkStatusComponent,
    StatusBoxComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule,
    ReactiveFormsModule,
    NgxChartsModule,
    BrowserAnimationsModule,
    TimeagoModule.forRoot(),
  ],
  providers: [
    {
      provide: GLOBAL_RX_STATE,
      useFactory: () => new RxState<GlobalState>(),
    },
  ],
  bootstrap: [AppComponent],
})
export class AppModule {}
