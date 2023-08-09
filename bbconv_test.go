package main

import (
	"reflect"
	"testing"
)

func Test_sanity(t *testing.T) {
	want := true

	if got := true; got != want {
		t.Errorf("We're insane: got %v, want %v", got, want)
	}
}

func Test_fileNamePrompt(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Figure out how to implement tests for this or refactor.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fileNamePrompt(); got != tt.want {
				t.Errorf("fileNamePrompt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_outFileName(t *testing.T) {
	type args struct {
		inFileName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "base_case", args: args{inFileName: "test_file.csv"}, want: "test_file-fixed.csv"},
		{name: "no_extension", args: args{inFileName: "test_file"}, want: "test_file-fixed"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := outFileName(tt.args.inFileName); got != tt.want {
				t.Errorf("outFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readCsv(t *testing.T) {
	basicCsvWant := [][]string{
		{"foo", "bar", "baz"},
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	monerisCsvWant := [][]string{
		{"Date", "Order ID", "Manual discounts", "Subtotal", "Trx discount", "Taxes", "Amount paid", "Total", "Cash penny rounding", "Payment Method", "Customer ID", "Employee ID"},
		{"2023/08/05 2:40:44 PM", "7271", "0.00", "4.75", "0.00", "0.24", "5.00", "5.00", "0.01", "Cash", "Anonymous", "username"},
		{"2023/08/05 2:26:35 PM", "7270", "0.00", "4.75", "0.00", "0.24", "4.99", "4.99", "0.00", "Debit", "Anonymous", "username"},
		{"2023/08/05 2:25:50 PM", "7269", "0.00", "4.40", "0.00", "0.22", "4.62", "4.62", "0.00", "Visa", "Anonymous", "username"},
		{"2023/08/05 2:25:04 PM", "7268", "0.00", "43.75", "0.00", "1.99", "45.74", "45.74", "0.00", "MasterCard", "Anonymous", "username"},
	}

	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{name: "base_case", args: args{fileName: "testdata/basic.csv"}, want: basicCsvWant},
		{name: "moneris_export", args: args{fileName: "testdata/moneris.csv"}, want: monerisCsvWant},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := readCsv(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readCsv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "base_case", args: args{date: "2023/07/08"}, want: "fizzbuzz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertDate(tt.args.date); got != tt.want {
				t.Errorf("convertDate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_dateIndex(t *testing.T) {
	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{name: "base_case", args: args{record: []string{"foo", "bar", "Date"}}, want: 2, wantErr: false},
		{name: "no_date", args: args{record: []string{"foo", "bar"}}, want: -1, wantErr: true},
		{name: "double_date", args: args{record: []string{"Date", "Date"}}, want: -1, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dateIndex(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("dateIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("dateIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
