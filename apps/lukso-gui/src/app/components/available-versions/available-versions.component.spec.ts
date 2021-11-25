import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { LetModule } from '@rx-angular/template';

import { AvailableVersionsComponent } from './available-versions.component';

describe('AvailableVersionsComponent', () => {
  let component: AvailableVersionsComponent;
  let fixture: ComponentFixture<AvailableVersionsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [HttpClientTestingModule, LetModule],
      declarations: [AvailableVersionsComponent],
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AvailableVersionsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
