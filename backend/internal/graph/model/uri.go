package model

import (
	"fmt"
	"io"
	"net/url"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalURI(u url.URL) graphql.WriterFunc {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, fmt.Sprintf(`"%s"`, u.String()))
	})
}

func UnmarshalURI(v interface{}) (url.URL, error) {
	switch v := v.(type) {
	case string:
		u, err := url.Parse(v)
		if err != nil {
			return url.URL{}, err
		}
		return *u, nil
	case []byte:
		u := &url.URL{}
		if err := u.UnmarshalBinary(v); err != nil {
			return url.URL{}, err
		}
		return *u, nil
	default:
		return url.URL{}, fmt.Errorf("%T is not a url.URL", v)
	}
}
