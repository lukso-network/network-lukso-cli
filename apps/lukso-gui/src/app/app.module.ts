import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { ReactiveFormsModule } from '@angular/forms';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { RxState } from '@rx-angular/state';
import {
  LetModule,
  PushModule,
  ViewportPrioModule,
} from '@rx-angular/template';
import { NgxChartsModule } from '@swimlane/ngx-charts';
import { TimeagoModule } from 'ngx-timeago';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { AvailableVersionsComponent } from './components/available-versions/available-versions.component';
import { SettingsComponent } from './components/settings/settings.component';
import { SetupComponent } from './components/setup/setup.component';
import { NetworkStatusComponent } from './components/status/network-status/network-status.component';
import { OrchestratorStatusComponent } from './components/status/orchestrator-status/orchestrator-status.component';
import { StatusBoxComponent } from './components/status/status-box/status-box.component';
import { StatusComponent } from './components/status/status.component';
import { ValidatorStatusComponent } from './components/status/validator-status/validator-status.component';
import { SharedModule } from './modules/shared/shared.module';
import { GlobalState, GLOBAL_RX_STATE } from './shared/rx-state';

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
    SetupComponent,
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule,
    ReactiveFormsModule,
    NgxChartsModule,
    BrowserAnimationsModule,
    TimeagoModule.forRoot(),
    LetModule,
    PushModule,
    ViewportPrioModule,
    SharedModule,
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
