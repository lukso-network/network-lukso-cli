import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { RxState } from '@rx-angular/state';
import { GLOBAL_RX_STATE } from '../../../../../../app/shared/rx-state';

import { CreateKeysComponent } from './create-keys.component';

describe('CreateKeysComponent', () => {
  let component: CreateKeysComponent;
  let fixture: ComponentFixture<CreateKeysComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [CreateKeysComponent],
      imports: [ReactiveFormsModule],
      providers: [{ provide: GLOBAL_RX_STATE, useClass: RxState }],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(CreateKeysComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
