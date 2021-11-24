import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ButtonComponent } from './button/button.component';
import {
  LetModule,
  PushModule,
  ViewportPrioModule,
} from '@rx-angular/template';

@NgModule({
  declarations: [ButtonComponent],
  imports: [CommonModule, LetModule, PushModule, ViewportPrioModule],
  exports: [ButtonComponent],
})
export class SharedModule {}
