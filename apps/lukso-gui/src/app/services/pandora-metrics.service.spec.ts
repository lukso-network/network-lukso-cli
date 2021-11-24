import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { PandoraMetricsService } from './pandora-metrics.service';

describe('PandoraMetricsService', () => {
  let service: PandoraMetricsService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
    });
    service = TestBed.inject(PandoraMetricsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
