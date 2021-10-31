import { HttpClientModule } from '@angular/common/http';
import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { PandoraStatusComponent } from './components/pandora-status/pandora-status.component';
import { VanguardStatusComponent } from './components/vanguard-status/vanguard-status.component';

@NgModule({
  declarations: [AppComponent, PandoraStatusComponent, VanguardStatusComponent],
  imports: [BrowserModule, HttpClientModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
