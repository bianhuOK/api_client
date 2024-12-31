package demo

import (
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

// RestFulHello is a struct used for implementation of restfull hello program
type RestFulHello struct {
}

// SayHello is a method used to reply user with hello in json format
func (r *RestFulHello) SayHello(b *rf.Context) {
	result := struct {
		Message string `json:"message"`
	}{
		Message: "hello",
	}
	b.WriteJSON(result, "application/json")
}

// URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns() []rf.Route {
	return []rf.Route{
		{Method: "GET", Path: "/hello", ResourceFunc: r.SayHello,
			Returns: []*rf.Returns{{Code: 200}}},
	}
}
