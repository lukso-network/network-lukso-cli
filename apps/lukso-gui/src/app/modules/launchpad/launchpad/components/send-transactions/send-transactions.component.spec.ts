import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SendTransactionsComponent } from './send-transactions.component';

describe('SendTransactionsComponent', () => {
  let component: SendTransactionsComponent;
  let fixture: ComponentFixture<SendTransactionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SendTransactionsComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SendTransactionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
