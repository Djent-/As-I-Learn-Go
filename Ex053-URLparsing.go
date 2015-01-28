// This work is copyright Mark McGranaghan and licensed under a Creative Commons Attribution 3.0 Unported License.
// http://creativecommons.org/licenses/by/3.0/

package main

import "fmt"
import "net/url"
import "strings"

func main() {
	//example URL contains scheme, auth info, host, port, path, query params,
	//and query fragment
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(u.Scheme)
	
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)
	
	//the Host contains both the hostname and the port
	//Split the Host manually to extract the port
	fmt.Println(u.Host)
	h := strings.Split(u.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])
	
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}