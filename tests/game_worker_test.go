package tests

import (
	"testing"

	"rubrumcreation.com/go-paper-scissor/routines"
)

func TestSmallIntervals(t *testing.T) {

	intervals, err := routines.GetIntervals(4)
	if err != nil {
		t.Fatalf("Error in function to get intervals")
	}

	if len(intervals) != 4 {
		t.Fatalf("Did not recieve the required amount of workers")
	}

	if intervals[0].Start != 0 && intervals[0].End != 1 {
		t.Logf("Nr:0 %v", intervals[0])
		t.Fatalf("Start and/or end of the worker is not correct")
	}
	if intervals[1].Start != 1 && intervals[1].End != 2 {
		t.Logf("Nr:0 %v", intervals[0])
		t.Fatalf("Start and/or end of the worker is not correct")
	}
	if intervals[2].Start != 2 && intervals[2].End != 3 {
		t.Logf("Nr:0 %v", intervals[0])
		t.Fatalf("Start and/or end of the worker is not correct")
	}
	if intervals[3].Start != 3 && intervals[3].End != 4 {
		t.Logf("Nr:0 %v", intervals[0])
		t.Fatalf("Start and/or end of the worker is not correct")
	}

}

func TestMediumIntervals(t *testing.T) {

	interval, err := routines.GetIntervals(20000)

	if err != nil {
		t.Fatalf("Error in medium partition test: %v", err)
	}

	if len(interval) != 4 {
		t.Fatalf("Did not recieve the required amount of workers")
	}

	if interval[0].Start != 0 && interval[0].End != 5000 {

		t.Logf("Nr:0 %v", interval[0])
		t.Fatalf("Start and/or end of the worker is not correct")
	}

	if interval[1].Start != 5000 && interval[1].End != 10000 {
		t.Logf("Nr:1 %v", interval[1])
		t.Fatalf("Start and/or end of the worker is not correct")
	}

	if interval[2].Start != 10000 && interval[2].End != 15000 {
		t.Logf("Nr:2 %v", interval[2])
		t.Fatalf("Start and/or end of the worker is not correct")
	}
	if interval[3].Start != 15000 && interval[3].End != 20000 {
		t.Logf("Nr:3 %v", interval[3])
		t.Fatalf("Start and/or end of the worker is not correct")
	}
}

func TestLargeIntervals(t *testing.T) {
	remainder := 1020047 % 4
	chunkSize := 1020047 / 4
	interval, err := routines.GetIntervals(1020047)
	if err != nil {
		t.Fatalf("Error in medium partition test: %v", err)
	}

	if len(interval) != 4 {
		t.Fatalf("Did not recieve the required amount of workers")
	}

	if interval[0].Start != 0 && interval[0].End != chunkSize+remainder {
		t.Logf("Nr:0 %v", interval[0])
		t.Fatalf("Start and/or end of the worker is not correct")
	}

	if interval[1].Start != interval[0].End && interval[1].End != interval[1].Start+chunkSize {
		t.Logf("Nr:1 %v", interval[1])
		t.Fatalf("Start and/or end of the worker is not correct")
	}

	if interval[2].Start != interval[1].End && interval[2].End != interval[2].End+chunkSize {
		t.Logf("Nr:2 %v", interval[2])
		t.Fatalf("Start and/or end of the worker is not correct")
	}

	if interval[3].Start != interval[2].End && interval[3].End != interval[3].Start+chunkSize {
		t.Logf("Nr:3 %v", interval[3])
		t.Fatalf("Start and/or end of the worker is not correct")
	}
}
