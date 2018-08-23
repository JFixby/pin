package pin

import "testing"

func TestLog(t *testing.T) {
	D("D tag message")
	E("E tag message")

	Debug("Debug tag message")
	Error("Error tag message")

	D("D tag", "message")
	E("E tag", "message")

	Debug("Debug tag", "message")
	Error("Error tag", "message")

	D("D array", [5]int{14234, 42, -1, 1000, 5})
	E("E araray", [5]int{14234, 42, -1, 1000, 5})

	/* Output:
	      pin_test.go:6 D tag message
          pin_test.go:7 E tag message
          pin_test.go:9 Debug tag message
         pin_test.go:10 Error tag message
         pin_test.go:12 D tag > message
         pin_test.go:13 E tag > message
         pin_test.go:15 Debug tag > message
         pin_test.go:16 Error tag > message
         pin_test.go:18 ---(D array)-[5]---
                                 (0) 14234
                                 (1) 42
                                 (2) -1
                                 (3) 1000
                                 (4) 5
         pin_test.go:19 ---(E araray)-[5]---
                                  (0) 14234
                                  (1) 42
                                  (2) -1
                                  (3) 1000
                                  (4) 5F
	 */
}
