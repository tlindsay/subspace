// TEMPORARY AUTOGENERATED FILE: easyjson stub code to make the package
// compilable during generation.

package  main

import (
  "github.com/mailru/easyjson/jwriter"
  "github.com/mailru/easyjson/jlexer"
)

func ( JsonMeta ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* JsonMeta ) UnmarshalJSON([]byte) error { return nil }
func ( JsonMeta ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* JsonMeta ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_JsonMeta *JsonMeta

func ( JsonResponse ) MarshalJSON() ([]byte, error) { return nil, nil }
func (* JsonResponse ) UnmarshalJSON([]byte) error { return nil }
func ( JsonResponse ) MarshalEasyJSON(w *jwriter.Writer) {}
func (* JsonResponse ) UnmarshalEasyJSON(l *jlexer.Lexer) {}

type EasyJSON_exporter_JsonResponse *JsonResponse