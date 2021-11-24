import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { KeygenService } from './keygen.service';

describe('KeygenService', () => {
  let service: KeygenService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(KeygenService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
