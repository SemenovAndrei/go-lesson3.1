package transaction

import "testing"

func TestMapRowToTransaction(t *testing.T) {
	type args struct {
		row []string
	}
	tests := []struct {
		name       string
		args       args
		wantId     string
		wantFrom   string
		wantTo     string
		wantAmount int64
	}{
		{
			name: "ok",
			args: args{row: []string{"0011","0001","0002","10000000"}},
			wantId: "0011",
			wantFrom: "0001",
			wantTo: "0002",
			wantAmount: 10000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotId, gotFrom, gotTo, gotAmount := MapRowToTransaction(tt.args.row)
			if gotId != tt.wantId {
				t.Errorf("MapRowToTransaction() gotId = %v, want %v", gotId, tt.wantId)
			}
			if gotFrom != tt.wantFrom {
				t.Errorf("MapRowToTransaction() gotFrom = %v, want %v", gotFrom, tt.wantFrom)
			}
			if gotTo != tt.wantTo {
				t.Errorf("MapRowToTransaction() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
			if gotAmount != tt.wantAmount {
				t.Errorf("MapRowToTransaction() gotAmount = %v, want %v", gotAmount, tt.wantAmount)
			}
		})
	}
}
