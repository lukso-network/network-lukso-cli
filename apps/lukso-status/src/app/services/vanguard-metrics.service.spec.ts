import { HttpClientTestingModule } from '@angular/common/http/testing';
import { TestBed } from '@angular/core/testing';

import { VanguardService } from './vanguard-metrics.service';

describe('VanguardServiceService', () => {
  let service: VanguardService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(VanguardService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
