import { ComponentFixture, TestBed } from '@angular/core/testing';
import { providers } from 'ethers';
import { SendTransactionsComponent } from './send-transactions.component';

jest.mock('ethers');

// eslint-disable-next-line @typescript-eslint/no-explicit-any
declare let window: any;

window.ethereum = {
  provider: {},
};

describe('SendTransactionsComponent', () => {
  let component: SendTransactionsComponent;
  let fixture: ComponentFixture<SendTransactionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [SendTransactionsComponent],
    }).compileComponents();
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
