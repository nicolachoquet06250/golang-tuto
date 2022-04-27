package errorHandling

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
)

type error interface {
	Error() string
}

type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }

func Main() {
	func() {
		f, err := os.Open("/test.txt")

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(f.Name(), "opened successfully")
	}()

	func() {
		f, err := os.Open("test.txt")
		if err != nil {
			if pErr, ok := err.(*os.PathError); ok {
				fmt.Println("Failed to open file at path", pErr.Path)
				return
			}
			fmt.Println("Generic error", err)
			return
		}
		fmt.Println(f.Name(), "opened successfully")
	}()

	func() {
		addr, err := net.LookupHost("golangbot123.com")
		if err != nil {
			if dnsErr, ok := err.(*net.DNSError); ok {
				if dnsErr.Timeout() {
					fmt.Println("operation timed out")
					return
				}
				if dnsErr.Temporary() {
					fmt.Println("temporary error")
					return
				}
				fmt.Println("Generic DNS error", err)
				return
			}
			fmt.Println("Generic error", err)
			return
		}
		fmt.Println(addr)
	}()

	func() {
		files, err := filepath.Glob("[")
		if err != nil {
			if err == filepath.ErrBadPattern {
				fmt.Println("Bad pattern error:", err)
				return
			}
			fmt.Println("Generic error:", err)
			return
		}
		fmt.Println("matched files", files)
	}()

	func() {
		files, _ := filepath.Glob("[")
		fmt.Println("matched files", files)
	}()
}
