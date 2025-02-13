import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TransactionHomeComponent } from './transactionHome.component';

describe('HomeComponent', () => {
  let component: TransactionHomeComponent;
  let fixture: ComponentFixture<TransactionHomeComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TransactionHomeComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TransactionHomeComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
